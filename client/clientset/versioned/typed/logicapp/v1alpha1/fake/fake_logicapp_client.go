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
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/clientset/versioned/typed/logicapp/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeLogicappV1alpha1 struct {
	*testing.Fake
}

func (c *FakeLogicappV1alpha1) ActionCustoms(namespace string) v1alpha1.ActionCustomInterface {
	return &FakeActionCustoms{c, namespace}
}

func (c *FakeLogicappV1alpha1) ActionHTTPs(namespace string) v1alpha1.ActionHTTPInterface {
	return &FakeActionHTTPs{c, namespace}
}

func (c *FakeLogicappV1alpha1) IntegrationAccounts(namespace string) v1alpha1.IntegrationAccountInterface {
	return &FakeIntegrationAccounts{c, namespace}
}

func (c *FakeLogicappV1alpha1) TriggerCustoms(namespace string) v1alpha1.TriggerCustomInterface {
	return &FakeTriggerCustoms{c, namespace}
}

func (c *FakeLogicappV1alpha1) TriggerHTTPRequests(namespace string) v1alpha1.TriggerHTTPRequestInterface {
	return &FakeTriggerHTTPRequests{c, namespace}
}

func (c *FakeLogicappV1alpha1) TriggerRecurrences(namespace string) v1alpha1.TriggerRecurrenceInterface {
	return &FakeTriggerRecurrences{c, namespace}
}

func (c *FakeLogicappV1alpha1) Workflows(namespace string) v1alpha1.WorkflowInterface {
	return &FakeWorkflows{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeLogicappV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
