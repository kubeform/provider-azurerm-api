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
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/clientset/versioned/typed/eventgrid/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeEventgridV1alpha1 struct {
	*testing.Fake
}

func (c *FakeEventgridV1alpha1) Domains(namespace string) v1alpha1.DomainInterface {
	return &FakeDomains{c, namespace}
}

func (c *FakeEventgridV1alpha1) DomainTopics(namespace string) v1alpha1.DomainTopicInterface {
	return &FakeDomainTopics{c, namespace}
}

func (c *FakeEventgridV1alpha1) EventSubscriptions(namespace string) v1alpha1.EventSubscriptionInterface {
	return &FakeEventSubscriptions{c, namespace}
}

func (c *FakeEventgridV1alpha1) SystemTopics(namespace string) v1alpha1.SystemTopicInterface {
	return &FakeSystemTopics{c, namespace}
}

func (c *FakeEventgridV1alpha1) SystemTopicEventSubscriptions(namespace string) v1alpha1.SystemTopicEventSubscriptionInterface {
	return &FakeSystemTopicEventSubscriptions{c, namespace}
}

func (c *FakeEventgridV1alpha1) Topics(namespace string) v1alpha1.TopicInterface {
	return &FakeTopics{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeEventgridV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}