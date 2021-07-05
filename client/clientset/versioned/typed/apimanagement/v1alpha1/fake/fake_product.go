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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/apimanagement/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeProducts implements ProductInterface
type FakeProducts struct {
	Fake *FakeApimanagementV1alpha1
	ns   string
}

var productsResource = schema.GroupVersionResource{Group: "apimanagement.azurerm.kubeform.com", Version: "v1alpha1", Resource: "products"}

var productsKind = schema.GroupVersionKind{Group: "apimanagement.azurerm.kubeform.com", Version: "v1alpha1", Kind: "Product"}

// Get takes name of the product, and returns the corresponding product object, and an error if there is any.
func (c *FakeProducts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Product, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(productsResource, c.ns, name), &v1alpha1.Product{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Product), err
}

// List takes label and field selectors, and returns the list of Products that match those selectors.
func (c *FakeProducts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ProductList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(productsResource, productsKind, c.ns, opts), &v1alpha1.ProductList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ProductList{ListMeta: obj.(*v1alpha1.ProductList).ListMeta}
	for _, item := range obj.(*v1alpha1.ProductList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested products.
func (c *FakeProducts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(productsResource, c.ns, opts))

}

// Create takes the representation of a product and creates it.  Returns the server's representation of the product, and an error, if there is any.
func (c *FakeProducts) Create(ctx context.Context, product *v1alpha1.Product, opts v1.CreateOptions) (result *v1alpha1.Product, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(productsResource, c.ns, product), &v1alpha1.Product{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Product), err
}

// Update takes the representation of a product and updates it. Returns the server's representation of the product, and an error, if there is any.
func (c *FakeProducts) Update(ctx context.Context, product *v1alpha1.Product, opts v1.UpdateOptions) (result *v1alpha1.Product, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(productsResource, c.ns, product), &v1alpha1.Product{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Product), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProducts) UpdateStatus(ctx context.Context, product *v1alpha1.Product, opts v1.UpdateOptions) (*v1alpha1.Product, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(productsResource, "status", c.ns, product), &v1alpha1.Product{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Product), err
}

// Delete takes name of the product and deletes it. Returns an error if one occurs.
func (c *FakeProducts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(productsResource, c.ns, name), &v1alpha1.Product{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProducts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(productsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ProductList{})
	return err
}

// Patch applies the patch and returns the patched product.
func (c *FakeProducts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Product, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(productsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Product{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Product), err
}