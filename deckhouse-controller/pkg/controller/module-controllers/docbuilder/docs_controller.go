// Copyright 2024 Flant JSC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docbuilder

import (
	"archive/tar"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/flant/addon-operator/pkg/utils/logger"
	log "github.com/sirupsen/logrus"
	coordv1 "k8s.io/api/coordination/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/retry"
	"k8s.io/utils/pointer"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/deckhouse/deckhouse/deckhouse-controller/pkg/apis/deckhouse.io/v1alpha1"
	d8env "github.com/deckhouse/deckhouse/go_lib/deckhouse-config/env"
	"github.com/deckhouse/deckhouse/go_lib/dependency"
	"github.com/deckhouse/deckhouse/go_lib/module"
	docs_builder "github.com/deckhouse/deckhouse/go_lib/module/docs-builder"
)

const defaultDocumentationCheckInterval = 10 * time.Second

type moduleDocumentationReconciler struct {
	client               client.Client
	downloadedModulesDir string

	dc          dependency.Container
	docsBuilder *docs_builder.Client

	logger logger.Logger
}

func NewModuleDocumentationController(mgr manager.Manager, dc dependency.Container) error {
	lg := log.WithField("component", "ModuleDocumentation")

	c := &moduleDocumentationReconciler{
		mgr.GetClient(),
		d8env.GetDownloadedModulesDir(),
		dependency.NewDependencyContainer(),
		docs_builder.NewClient(dc.GetHTTPClient()),
		lg,
	}

	ctr, err := controller.New("module-documentation", mgr, controller.Options{
		MaxConcurrentReconciles: 1, // don't use concurrent reconciles here, because docs-builder doesn't support multiply requests at once
		CacheSyncTimeout:        15 * time.Minute,
		NeedLeaderElection:      pointer.Bool(false),
		Reconciler:              c,
	})
	if err != nil {
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.ModuleDocumentation{}).
		Watches(&coordv1.Lease{}, handler.EnqueueRequestsFromMapFunc(c.enqueueLeaseMapFunc), builder.WithPredicates(predicate.Funcs{
			CreateFunc: func(event event.CreateEvent) bool {
				ns := event.Object.GetNamespace()
				if ns != "d8-system" {
					return false
				}

				var hasLabel bool
				for label := range event.Object.GetLabels() {
					if label == "deckhouse.io/documentation-builder-sync" {
						hasLabel = true
						break
					}
				}

				return hasLabel
			},
		})).
		WithEventFilter(predicate.GenerationChangedPredicate{}).
		Complete(ctr)
}

func (mdr *moduleDocumentationReconciler) enqueueLeaseMapFunc(ctx context.Context, _ client.Object) []reconcile.Request {
	requests := make([]reconcile.Request, 0)

	err := retry.OnError(retry.DefaultRetry, apierrors.IsServiceUnavailable, func() error {
		var mdl = new(v1alpha1.ModuleDocumentationList)

		err := mdr.client.List(ctx, mdl)
		if err != nil {
			return err
		}

		requests = make([]reconcile.Request, 0, len(mdl.Items))

		for _, md := range mdl.Items {
			requests = append(requests, reconcile.Request{NamespacedName: types.NamespacedName{Name: md.GetName()}})
		}

		return nil
	})
	if err != nil {
		log.Errorf("create mapping for lease failed: %s", err.Error())
	}

	return requests
}

const documentationExistsFinalizer = "modules.deckhouse.io/documentation-exists"

func (mdr *moduleDocumentationReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	res := ctrl.Result{}

	var md = new(v1alpha1.ModuleDocumentation)

	err := mdr.client.Get(ctx, req.NamespacedName, md)
	if err != nil {
		return res, client.IgnoreNotFound(err)
	}

	if md.DeletionTimestamp.IsZero() {
		if !controllerutil.ContainsFinalizer(md, documentationExistsFinalizer) {
			controllerutil.AddFinalizer(md, documentationExistsFinalizer)
			if err := mdr.client.Update(ctx, md); err != nil {
				mdr.logger.Errorf("update finalizer: %v", err)

				return res, err
			}
		}
	} else {
		if !controllerutil.ContainsFinalizer(md, documentationExistsFinalizer) {
			return res, nil
		}

		// get addresses from cluster, not status, because them more actual
		addrs, err := mdr.getDocsBuilderAddresses(ctx)
		if err != nil {
			return res, fmt.Errorf("get docs builder addresses: %w", err)
		}

		if len(addrs) == 0 {
			// no endpoints for doc builder
			return res, nil
		}

		now := metav1.NewTime(mdr.dc.GetClock().Now().UTC())

		for _, addr := range addrs {
			err := mdr.deleteDocumentation(ctx, addr, md.Name)
			if err == nil {
				continue
			}

			delErr := fmt.Errorf("delete documentation: %w", err)

			_, idx := md.GetConditionByAddress(addr)
			if idx < 0 {
				continue
			}

			md.Status.Conditions[idx].Type = v1alpha1.TypeError
			md.Status.Conditions[idx].Message = delErr.Error()
			md.Status.Conditions[idx].LastTransitionTime = now

			if err := mdr.client.Status().Update(ctx, md); err != nil {
				mdr.logger.Errorf("update status when delete documentation: %v", err)

				return res, fmt.Errorf("update status when delete documentation: %w", errors.Join(delErr, err))
			}

			return res, delErr
		}

		controllerutil.RemoveFinalizer(md, documentationExistsFinalizer)
		if err := mdr.client.Update(ctx, md); err != nil {
			mdr.logger.Errorf("update finalizer: %v", err)

			return res, fmt.Errorf("update finalizer: %w", err)
		}

		return res, nil
	}

	return mdr.createOrUpdateReconcile(ctx, md)
}

func (mdr *moduleDocumentationReconciler) createOrUpdateReconcile(ctx context.Context, md *v1alpha1.ModuleDocumentation) (ctrl.Result, error) {
	res := ctrl.Result{}

	moduleName := md.Name
	mdr.logger.Infof("Updating documentation for %s module", moduleName)
	addrs, err := mdr.getDocsBuilderAddresses(ctx)
	if err != nil {
		return res, fmt.Errorf("get docs builder addresses: %w", err)
	}

	if len(addrs) == 0 {
		// no endpoints for doc builder
		return res, nil
	}

	pr, pw := io.Pipe()
	defer pr.Close()

	mdr.logger.Debugf("Getting the %s module's documentation locally", moduleName)
	fetchModuleErr := mdr.getDocumentationFromModuleDir(md.Spec.Path, pw)

	var rendered int
	now := metav1.NewTime(mdr.dc.GetClock().Now().UTC())

	mdCopy := md.DeepCopy()
	mdCopy.Status.Conditions = make([]v1alpha1.ModuleDocumentationCondition, 0, len(addrs))

	for _, addr := range addrs {
		cond, condIdx := md.GetConditionByAddress(addr)
		if !(condIdx < 0) && cond.Version == md.Spec.Version && cond.Checksum == md.Spec.Checksum && cond.Type == v1alpha1.TypeRendered {
			// documentation is rendered for this builder
			mdCopy.Status.Conditions = append(mdCopy.Status.Conditions, cond)
			rendered++
			continue
		}

		cond = v1alpha1.ModuleDocumentationCondition{
			Address:            addr,
			Version:            md.Spec.Version,
			Checksum:           md.Spec.Checksum,
			LastTransitionTime: now,
		}

		if fetchModuleErr != nil {
			cond.Type = v1alpha1.TypeError
			cond.Message = fmt.Sprintf("Error occurred while fetching the documentation: %s. Please fix the module's docs or restart the Deckhouse to restore the module", fetchModuleErr)
			mdCopy.Status.Conditions = append(mdCopy.Status.Conditions, cond)
			continue
		}

		err = mdr.buildDocumentation(ctx, pr, addr, moduleName, md.Spec.Version)
		if err != nil {
			cond.Type = v1alpha1.TypeError
			cond.Message = err.Error()
		} else {
			rendered++
			cond.Type = v1alpha1.TypeRendered
			cond.Message = ""
		}

		mdCopy.Status.Conditions = append(mdCopy.Status.Conditions, cond)
	}

	switch {
	case rendered == 0:
		mdCopy.Status.RenderResult = v1alpha1.ResultError

	case rendered == len(addrs):
		mdCopy.Status.RenderResult = v1alpha1.ResultRendered

	default:
		mdCopy.Status.RenderResult = v1alpha1.ResultPartially
	}

	err = mdr.client.Status().Patch(ctx, mdCopy, client.MergeFrom(md))
	if err != nil {
		return res, err
	}

	if mdCopy.Status.RenderResult != v1alpha1.ResultRendered {
		return ctrl.Result{RequeueAfter: defaultDocumentationCheckInterval}, nil
	}

	return res, nil
}

func (mdr *moduleDocumentationReconciler) getDocsBuilderAddresses(ctx context.Context) (addresses []string, err error) {
	var leasesList coordv1.LeaseList
	err = mdr.client.List(ctx, &leasesList, client.InNamespace("d8-system"), client.HasLabels{"deckhouse.io/documentation-builder-sync"})
	if err != nil {
		return nil, fmt.Errorf("list leases: %w", err)
	}

	for _, lease := range leasesList.Items {
		if lease.Spec.HolderIdentity == nil {
			continue
		}

		// a stale lease found
		if lease.Spec.RenewTime.Add(time.Duration(*lease.Spec.LeaseDurationSeconds) * time.Second).Before(mdr.dc.GetClock().Now()) {
			continue
		}

		addresses = append(addresses, "http://"+*lease.Spec.HolderIdentity)
	}

	return
}

func (mdr *moduleDocumentationReconciler) getDocumentationFromModuleDir(modulePath string, pw *io.PipeWriter) error {
	moduleDir := path.Join(mdr.downloadedModulesDir, modulePath) + "/"

	dir, err := os.Stat(moduleDir)
	if err != nil {
		return err
	}

	if !dir.IsDir() {
		return fmt.Errorf("%s isn't a directory", moduleDir)
	}

	go func() {
		tw := tar.NewWriter(pw)
		defer tw.Close()

		_ = pw.CloseWithError(filepath.Walk(moduleDir, func(file string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !module.IsDocsPath(strings.TrimPrefix(file, moduleDir)) {
				return nil
			}

			header, err := tar.FileInfoHeader(info, info.Name())
			if err != nil {
				return err
			}

			header.Name = strings.TrimPrefix(file, moduleDir)

			if err := tw.WriteHeader(header); err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			f, err := os.Open(file)
			if err != nil {
				return err
			}
			defer f.Close()

			if _, err := io.Copy(tw, f); err != nil {
				return err
			}

			return nil
		}))
	}()

	return nil
}

func (mdr *moduleDocumentationReconciler) buildDocumentation(ctx context.Context, docsArchive io.Reader, baseAddr, moduleName, moduleVersion string) error {
	err := mdr.docsBuilder.SendDocumentation(ctx, baseAddr, moduleName, moduleVersion, docsArchive)
	if err != nil {
		return fmt.Errorf("send documentation: %w", err)
	}

	err = mdr.docsBuilder.BuildDocumentation(ctx, baseAddr)
	if err != nil {
		return fmt.Errorf("build documentation: %w", err)
	}

	return nil
}

func (mdr *moduleDocumentationReconciler) deleteDocumentation(ctx context.Context, baseAddr, moduleName string) error {
	err := mdr.docsBuilder.DeleteDocumentation(ctx, baseAddr, moduleName)
	if err != nil {
		return fmt.Errorf("delete documentation: %w", err)
	}

	err = mdr.docsBuilder.BuildDocumentation(ctx, baseAddr)
	if err != nil {
		return fmt.Errorf("build documentation: %w", err)
	}

	return nil
}
