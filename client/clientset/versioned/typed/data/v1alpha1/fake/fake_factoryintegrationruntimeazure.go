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
	"context"

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/data/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFactoryIntegrationRuntimeAzures implements FactoryIntegrationRuntimeAzureInterface
type FakeFactoryIntegrationRuntimeAzures struct {
	Fake *FakeDataV1alpha1
	ns   string
}

var factoryintegrationruntimeazuresResource = schema.GroupVersionResource{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Resource: "factoryintegrationruntimeazures"}

var factoryintegrationruntimeazuresKind = schema.GroupVersionKind{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Kind: "FactoryIntegrationRuntimeAzure"}

// Get takes name of the factoryIntegrationRuntimeAzure, and returns the corresponding factoryIntegrationRuntimeAzure object, and an error if there is any.
func (c *FakeFactoryIntegrationRuntimeAzures) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FactoryIntegrationRuntimeAzure, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(factoryintegrationruntimeazuresResource, c.ns, name), &v1alpha1.FactoryIntegrationRuntimeAzure{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryIntegrationRuntimeAzure), err
}

// List takes label and field selectors, and returns the list of FactoryIntegrationRuntimeAzures that match those selectors.
func (c *FakeFactoryIntegrationRuntimeAzures) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FactoryIntegrationRuntimeAzureList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(factoryintegrationruntimeazuresResource, factoryintegrationruntimeazuresKind, c.ns, opts), &v1alpha1.FactoryIntegrationRuntimeAzureList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FactoryIntegrationRuntimeAzureList{ListMeta: obj.(*v1alpha1.FactoryIntegrationRuntimeAzureList).ListMeta}
	for _, item := range obj.(*v1alpha1.FactoryIntegrationRuntimeAzureList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested factoryIntegrationRuntimeAzures.
func (c *FakeFactoryIntegrationRuntimeAzures) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(factoryintegrationruntimeazuresResource, c.ns, opts))

}

// Create takes the representation of a factoryIntegrationRuntimeAzure and creates it.  Returns the server's representation of the factoryIntegrationRuntimeAzure, and an error, if there is any.
func (c *FakeFactoryIntegrationRuntimeAzures) Create(ctx context.Context, factoryIntegrationRuntimeAzure *v1alpha1.FactoryIntegrationRuntimeAzure, opts v1.CreateOptions) (result *v1alpha1.FactoryIntegrationRuntimeAzure, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(factoryintegrationruntimeazuresResource, c.ns, factoryIntegrationRuntimeAzure), &v1alpha1.FactoryIntegrationRuntimeAzure{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryIntegrationRuntimeAzure), err
}

// Update takes the representation of a factoryIntegrationRuntimeAzure and updates it. Returns the server's representation of the factoryIntegrationRuntimeAzure, and an error, if there is any.
func (c *FakeFactoryIntegrationRuntimeAzures) Update(ctx context.Context, factoryIntegrationRuntimeAzure *v1alpha1.FactoryIntegrationRuntimeAzure, opts v1.UpdateOptions) (result *v1alpha1.FactoryIntegrationRuntimeAzure, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(factoryintegrationruntimeazuresResource, c.ns, factoryIntegrationRuntimeAzure), &v1alpha1.FactoryIntegrationRuntimeAzure{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryIntegrationRuntimeAzure), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFactoryIntegrationRuntimeAzures) UpdateStatus(ctx context.Context, factoryIntegrationRuntimeAzure *v1alpha1.FactoryIntegrationRuntimeAzure, opts v1.UpdateOptions) (*v1alpha1.FactoryIntegrationRuntimeAzure, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(factoryintegrationruntimeazuresResource, "status", c.ns, factoryIntegrationRuntimeAzure), &v1alpha1.FactoryIntegrationRuntimeAzure{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryIntegrationRuntimeAzure), err
}

// Delete takes name of the factoryIntegrationRuntimeAzure and deletes it. Returns an error if one occurs.
func (c *FakeFactoryIntegrationRuntimeAzures) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(factoryintegrationruntimeazuresResource, c.ns, name), &v1alpha1.FactoryIntegrationRuntimeAzure{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFactoryIntegrationRuntimeAzures) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(factoryintegrationruntimeazuresResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FactoryIntegrationRuntimeAzureList{})
	return err
}

// Patch applies the patch and returns the patched factoryIntegrationRuntimeAzure.
func (c *FakeFactoryIntegrationRuntimeAzures) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FactoryIntegrationRuntimeAzure, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(factoryintegrationruntimeazuresResource, c.ns, name, pt, data, subresources...), &v1alpha1.FactoryIntegrationRuntimeAzure{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryIntegrationRuntimeAzure), err
}
