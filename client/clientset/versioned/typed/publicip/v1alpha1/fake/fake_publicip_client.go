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
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/clientset/versioned/typed/publicip/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakePublicipV1alpha1 struct {
	*testing.Fake
}

func (c *FakePublicipV1alpha1) Prefixes(namespace string) v1alpha1.PrefixInterface {
	return &FakePrefixes{c, namespace}
}

func (c *FakePublicipV1alpha1) PublicIPs(namespace string) v1alpha1.PublicIPInterface {
	return &FakePublicIPs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakePublicipV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
