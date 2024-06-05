/*
Copyright 2024 Flant JSC
Licensed under the Deckhouse Platform Enterprise Edition (EE) license. See https://github.com/deckhouse/deckhouse/blob/main/ee/LICENSE
*/

package master

import (
	"context"
	"fmt"
	"system-registry-manager/internal/master/workflow"
	pkg_logs "system-registry-manager/pkg/logs"
	"time"
)

const (
	workInterval = 10 * time.Second
)

func startMasterWorkflow(ctx context.Context, m *Master) {
	log := pkg_logs.GetLoggerFromContext(ctx)
	m.k8sHandler.Start()
	defer m.k8sHandler.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(workInterval)
			err := masterWorkflow(ctx, m)
			if err != nil {
				log.Error(err)
				continue
			}
		}
	}
}

func masterWorkflow(ctx context.Context, m *Master) error {
	// log := pkg_logs.GetLoggerFromContext(ctx)

	workerCount, err := m.k8sHandler.WaitAllWorkers()
	if err != nil {
		return err
	}

	masters := m.k8sHandler.GetMasterNodeNameList()
	if len(masters) != workerCount {
		return fmt.Errorf("len(masters) != workerCount")
	}

	nodeManagers := make([]workflow.NodeManager, 0, len(masters))

	for _, master := range masters {
		nodeManagers = append(nodeManagers, NewNodeManager(ctx, master, m.k8sHandler))
	}

	seaweedfsCaCertsWorkflow := workflow.NewSeaweedfsCaCertsWorkflow(ctx, nodeManagers, len(nodeManagers))
	err = seaweedfsCaCertsWorkflow.Start()
	if err != nil {
		return err
	}

	seaweedfsScaleWorkflow := workflow.NewSeaweedfsScaleWorkflow(ctx, nodeManagers, len(nodeManagers))
	err = seaweedfsScaleWorkflow.Start()
	if err != nil {
		return err
	}
	return nil
}
