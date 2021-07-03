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

// FakeFactoryLinkedServiceKustos implements FactoryLinkedServiceKustoInterface
type FakeFactoryLinkedServiceKustos struct {
	Fake *FakeDataV1alpha1
	ns   string
}

var factorylinkedservicekustosResource = schema.GroupVersionResource{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Resource: "factorylinkedservicekustos"}

var factorylinkedservicekustosKind = schema.GroupVersionKind{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Kind: "FactoryLinkedServiceKusto"}

// Get takes name of the factoryLinkedServiceKusto, and returns the corresponding factoryLinkedServiceKusto object, and an error if there is any.
func (c *FakeFactoryLinkedServiceKustos) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FactoryLinkedServiceKusto, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(factorylinkedservicekustosResource, c.ns, name), &v1alpha1.FactoryLinkedServiceKusto{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceKusto), err
}

// List takes label and field selectors, and returns the list of FactoryLinkedServiceKustos that match those selectors.
func (c *FakeFactoryLinkedServiceKustos) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FactoryLinkedServiceKustoList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(factorylinkedservicekustosResource, factorylinkedservicekustosKind, c.ns, opts), &v1alpha1.FactoryLinkedServiceKustoList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FactoryLinkedServiceKustoList{ListMeta: obj.(*v1alpha1.FactoryLinkedServiceKustoList).ListMeta}
	for _, item := range obj.(*v1alpha1.FactoryLinkedServiceKustoList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested factoryLinkedServiceKustos.
func (c *FakeFactoryLinkedServiceKustos) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(factorylinkedservicekustosResource, c.ns, opts))

}

// Create takes the representation of a factoryLinkedServiceKusto and creates it.  Returns the server's representation of the factoryLinkedServiceKusto, and an error, if there is any.
func (c *FakeFactoryLinkedServiceKustos) Create(ctx context.Context, factoryLinkedServiceKusto *v1alpha1.FactoryLinkedServiceKusto, opts v1.CreateOptions) (result *v1alpha1.FactoryLinkedServiceKusto, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(factorylinkedservicekustosResource, c.ns, factoryLinkedServiceKusto), &v1alpha1.FactoryLinkedServiceKusto{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceKusto), err
}

// Update takes the representation of a factoryLinkedServiceKusto and updates it. Returns the server's representation of the factoryLinkedServiceKusto, and an error, if there is any.
func (c *FakeFactoryLinkedServiceKustos) Update(ctx context.Context, factoryLinkedServiceKusto *v1alpha1.FactoryLinkedServiceKusto, opts v1.UpdateOptions) (result *v1alpha1.FactoryLinkedServiceKusto, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(factorylinkedservicekustosResource, c.ns, factoryLinkedServiceKusto), &v1alpha1.FactoryLinkedServiceKusto{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceKusto), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFactoryLinkedServiceKustos) UpdateStatus(ctx context.Context, factoryLinkedServiceKusto *v1alpha1.FactoryLinkedServiceKusto, opts v1.UpdateOptions) (*v1alpha1.FactoryLinkedServiceKusto, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(factorylinkedservicekustosResource, "status", c.ns, factoryLinkedServiceKusto), &v1alpha1.FactoryLinkedServiceKusto{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceKusto), err
}

// Delete takes name of the factoryLinkedServiceKusto and deletes it. Returns an error if one occurs.
func (c *FakeFactoryLinkedServiceKustos) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(factorylinkedservicekustosResource, c.ns, name), &v1alpha1.FactoryLinkedServiceKusto{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFactoryLinkedServiceKustos) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(factorylinkedservicekustosResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FactoryLinkedServiceKustoList{})
	return err
}

// Patch applies the patch and returns the patched factoryLinkedServiceKusto.
func (c *FakeFactoryLinkedServiceKustos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FactoryLinkedServiceKusto, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(factorylinkedservicekustosResource, c.ns, name, pt, data, subresources...), &v1alpha1.FactoryLinkedServiceKusto{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceKusto), err
}
