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
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/clientset/versioned/typed/spring/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSpringV1alpha1 struct {
	*testing.Fake
}

func (c *FakeSpringV1alpha1) CloudActiveDeployments(namespace string) v1alpha1.CloudActiveDeploymentInterface {
	return &FakeCloudActiveDeployments{c, namespace}
}

func (c *FakeSpringV1alpha1) CloudApps(namespace string) v1alpha1.CloudAppInterface {
	return &FakeCloudApps{c, namespace}
}

func (c *FakeSpringV1alpha1) CloudAppCosmosdbAssociations(namespace string) v1alpha1.CloudAppCosmosdbAssociationInterface {
	return &FakeCloudAppCosmosdbAssociations{c, namespace}
}

func (c *FakeSpringV1alpha1) CloudAppMysqlAssociations(namespace string) v1alpha1.CloudAppMysqlAssociationInterface {
	return &FakeCloudAppMysqlAssociations{c, namespace}
}

func (c *FakeSpringV1alpha1) CloudAppRedisAssociations(namespace string) v1alpha1.CloudAppRedisAssociationInterface {
	return &FakeCloudAppRedisAssociations{c, namespace}
}

func (c *FakeSpringV1alpha1) CloudCertificates(namespace string) v1alpha1.CloudCertificateInterface {
	return &FakeCloudCertificates{c, namespace}
}

func (c *FakeSpringV1alpha1) CloudCustomDomains(namespace string) v1alpha1.CloudCustomDomainInterface {
	return &FakeCloudCustomDomains{c, namespace}
}

func (c *FakeSpringV1alpha1) CloudJavaDeployments(namespace string) v1alpha1.CloudJavaDeploymentInterface {
	return &FakeCloudJavaDeployments{c, namespace}
}

func (c *FakeSpringV1alpha1) CloudServices(namespace string) v1alpha1.CloudServiceInterface {
	return &FakeCloudServices{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSpringV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
