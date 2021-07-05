/*
Copyright AppsCode Inc. and Contributors

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
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/clientset/versioned/typed/synapse/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSynapseV1alpha1 struct {
	*testing.Fake
}

func (c *FakeSynapseV1alpha1) FirewallRules(namespace string) v1alpha1.FirewallRuleInterface {
	return &FakeFirewallRules{c, namespace}
}

func (c *FakeSynapseV1alpha1) ManagedPrivateEndpoints(namespace string) v1alpha1.ManagedPrivateEndpointInterface {
	return &FakeManagedPrivateEndpoints{c, namespace}
}

func (c *FakeSynapseV1alpha1) RoleAssignments(namespace string) v1alpha1.RoleAssignmentInterface {
	return &FakeRoleAssignments{c, namespace}
}

func (c *FakeSynapseV1alpha1) SparkPools(namespace string) v1alpha1.SparkPoolInterface {
	return &FakeSparkPools{c, namespace}
}

func (c *FakeSynapseV1alpha1) SqlPools(namespace string) v1alpha1.SqlPoolInterface {
	return &FakeSqlPools{c, namespace}
}

func (c *FakeSynapseV1alpha1) Workspaces(namespace string) v1alpha1.WorkspaceInterface {
	return &FakeWorkspaces{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSynapseV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}