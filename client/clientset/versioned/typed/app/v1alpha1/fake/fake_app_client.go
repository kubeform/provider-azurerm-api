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
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/clientset/versioned/typed/app/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeAppV1alpha1 struct {
	*testing.Fake
}

func (c *FakeAppV1alpha1) Configurations(namespace string) v1alpha1.ConfigurationInterface {
	return &FakeConfigurations{c, namespace}
}

func (c *FakeAppV1alpha1) Services(namespace string) v1alpha1.ServiceInterface {
	return &FakeServices{c, namespace}
}

func (c *FakeAppV1alpha1) ServiceActiveSlots(namespace string) v1alpha1.ServiceActiveSlotInterface {
	return &FakeServiceActiveSlots{c, namespace}
}

func (c *FakeAppV1alpha1) ServiceCertificates(namespace string) v1alpha1.ServiceCertificateInterface {
	return &FakeServiceCertificates{c, namespace}
}

func (c *FakeAppV1alpha1) ServiceCertificateBindings(namespace string) v1alpha1.ServiceCertificateBindingInterface {
	return &FakeServiceCertificateBindings{c, namespace}
}

func (c *FakeAppV1alpha1) ServiceCertificateOrders(namespace string) v1alpha1.ServiceCertificateOrderInterface {
	return &FakeServiceCertificateOrders{c, namespace}
}

func (c *FakeAppV1alpha1) ServiceCustomHostnameBindings(namespace string) v1alpha1.ServiceCustomHostnameBindingInterface {
	return &FakeServiceCustomHostnameBindings{c, namespace}
}

func (c *FakeAppV1alpha1) ServiceEnvironments(namespace string) v1alpha1.ServiceEnvironmentInterface {
	return &FakeServiceEnvironments{c, namespace}
}

func (c *FakeAppV1alpha1) ServiceEnvironmentV3s(namespace string) v1alpha1.ServiceEnvironmentV3Interface {
	return &FakeServiceEnvironmentV3s{c, namespace}
}

func (c *FakeAppV1alpha1) ServiceHybridConnections(namespace string) v1alpha1.ServiceHybridConnectionInterface {
	return &FakeServiceHybridConnections{c, namespace}
}

func (c *FakeAppV1alpha1) ServiceManagedCertificates(namespace string) v1alpha1.ServiceManagedCertificateInterface {
	return &FakeServiceManagedCertificates{c, namespace}
}

func (c *FakeAppV1alpha1) ServicePlans(namespace string) v1alpha1.ServicePlanInterface {
	return &FakeServicePlans{c, namespace}
}

func (c *FakeAppV1alpha1) ServiceSlots(namespace string) v1alpha1.ServiceSlotInterface {
	return &FakeServiceSlots{c, namespace}
}

func (c *FakeAppV1alpha1) ServiceSlotVirtualNetworkSwiftConnections(namespace string) v1alpha1.ServiceSlotVirtualNetworkSwiftConnectionInterface {
	return &FakeServiceSlotVirtualNetworkSwiftConnections{c, namespace}
}

func (c *FakeAppV1alpha1) ServiceSourceControlTokens(namespace string) v1alpha1.ServiceSourceControlTokenInterface {
	return &FakeServiceSourceControlTokens{c, namespace}
}

func (c *FakeAppV1alpha1) ServiceVirtualNetworkSwiftConnections(namespace string) v1alpha1.ServiceVirtualNetworkSwiftConnectionInterface {
	return &FakeServiceVirtualNetworkSwiftConnections{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAppV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}