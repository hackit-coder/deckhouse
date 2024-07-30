/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/deckhouse/deckhouse/deckhouse-controller/pkg/client/clientset/versioned/typed/deckhouse.io/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeDeckhouseV1alpha1 struct {
	*testing.Fake
}

func (c *FakeDeckhouseV1alpha1) DeckhouseReleases() v1alpha1.DeckhouseReleaseInterface {
	return &FakeDeckhouseReleases{c}
}

func (c *FakeDeckhouseV1alpha1) Modules() v1alpha1.ModuleInterface {
	return &FakeModules{c}
}

func (c *FakeDeckhouseV1alpha1) ModuleConfigs() v1alpha1.ModuleConfigInterface {
	return &FakeModuleConfigs{c}
}

func (c *FakeDeckhouseV1alpha1) ModuleDocumentations() v1alpha1.ModuleDocumentationInterface {
	return &FakeModuleDocumentations{c}
}

func (c *FakeDeckhouseV1alpha1) ModulePullOverrides() v1alpha1.ModulePullOverrideInterface {
	return &FakeModulePullOverrides{c}
}

func (c *FakeDeckhouseV1alpha1) ModuleReleases() v1alpha1.ModuleReleaseInterface {
	return &FakeModuleReleases{c}
}

func (c *FakeDeckhouseV1alpha1) ModuleSources() v1alpha1.ModuleSourceInterface {
	return &FakeModuleSources{c}
}

func (c *FakeDeckhouseV1alpha1) ModuleUpdatePolicies() v1alpha1.ModuleUpdatePolicyInterface {
	return &FakeModuleUpdatePolicies{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeDeckhouseV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
