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
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/clientset/versioned/typed/iothub/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeIothubV1alpha1 struct {
	*testing.Fake
}

func (c *FakeIothubV1alpha1) ConsumerGroups(namespace string) v1alpha1.ConsumerGroupInterface {
	return &FakeConsumerGroups{c, namespace}
}

func (c *FakeIothubV1alpha1) Dpses(namespace string) v1alpha1.DpsInterface {
	return &FakeDpses{c, namespace}
}

func (c *FakeIothubV1alpha1) DpsCertificates(namespace string) v1alpha1.DpsCertificateInterface {
	return &FakeDpsCertificates{c, namespace}
}

func (c *FakeIothubV1alpha1) DpsSharedAccessPolicies(namespace string) v1alpha1.DpsSharedAccessPolicyInterface {
	return &FakeDpsSharedAccessPolicies{c, namespace}
}

func (c *FakeIothubV1alpha1) EndpointEventhubs(namespace string) v1alpha1.EndpointEventhubInterface {
	return &FakeEndpointEventhubs{c, namespace}
}

func (c *FakeIothubV1alpha1) EndpointServicebusQueues(namespace string) v1alpha1.EndpointServicebusQueueInterface {
	return &FakeEndpointServicebusQueues{c, namespace}
}

func (c *FakeIothubV1alpha1) EndpointServicebusTopics(namespace string) v1alpha1.EndpointServicebusTopicInterface {
	return &FakeEndpointServicebusTopics{c, namespace}
}

func (c *FakeIothubV1alpha1) EndpointStorageContainers(namespace string) v1alpha1.EndpointStorageContainerInterface {
	return &FakeEndpointStorageContainers{c, namespace}
}

func (c *FakeIothubV1alpha1) Enrichments(namespace string) v1alpha1.EnrichmentInterface {
	return &FakeEnrichments{c, namespace}
}

func (c *FakeIothubV1alpha1) FallbackRoutes(namespace string) v1alpha1.FallbackRouteInterface {
	return &FakeFallbackRoutes{c, namespace}
}

func (c *FakeIothubV1alpha1) Iothubs(namespace string) v1alpha1.IothubInterface {
	return &FakeIothubs{c, namespace}
}

func (c *FakeIothubV1alpha1) Routes(namespace string) v1alpha1.RouteInterface {
	return &FakeRoutes{c, namespace}
}

func (c *FakeIothubV1alpha1) SharedAccessPolicies(namespace string) v1alpha1.SharedAccessPolicyInterface {
	return &FakeSharedAccessPolicies{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeIothubV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
