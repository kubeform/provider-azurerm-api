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
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/clientset/versioned/typed/kusto/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeKustoV1alpha1 struct {
	*testing.Fake
}

func (c *FakeKustoV1alpha1) AttachedDatabaseConfigurations(namespace string) v1alpha1.AttachedDatabaseConfigurationInterface {
	return &FakeAttachedDatabaseConfigurations{c, namespace}
}

func (c *FakeKustoV1alpha1) Clusters(namespace string) v1alpha1.ClusterInterface {
	return &FakeClusters{c, namespace}
}

func (c *FakeKustoV1alpha1) ClusterCustomerManagedKeys(namespace string) v1alpha1.ClusterCustomerManagedKeyInterface {
	return &FakeClusterCustomerManagedKeys{c, namespace}
}

func (c *FakeKustoV1alpha1) ClusterPrincipalAssignments(namespace string) v1alpha1.ClusterPrincipalAssignmentInterface {
	return &FakeClusterPrincipalAssignments{c, namespace}
}

func (c *FakeKustoV1alpha1) Databases(namespace string) v1alpha1.DatabaseInterface {
	return &FakeDatabases{c, namespace}
}

func (c *FakeKustoV1alpha1) DatabasePrincipals(namespace string) v1alpha1.DatabasePrincipalInterface {
	return &FakeDatabasePrincipals{c, namespace}
}

func (c *FakeKustoV1alpha1) DatabasePrincipalAssignments(namespace string) v1alpha1.DatabasePrincipalAssignmentInterface {
	return &FakeDatabasePrincipalAssignments{c, namespace}
}

func (c *FakeKustoV1alpha1) EventgridDataConnections(namespace string) v1alpha1.EventgridDataConnectionInterface {
	return &FakeEventgridDataConnections{c, namespace}
}

func (c *FakeKustoV1alpha1) EventhubDataConnections(namespace string) v1alpha1.EventhubDataConnectionInterface {
	return &FakeEventhubDataConnections{c, namespace}
}

func (c *FakeKustoV1alpha1) IothubDataConnections(namespace string) v1alpha1.IothubDataConnectionInterface {
	return &FakeIothubDataConnections{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKustoV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
