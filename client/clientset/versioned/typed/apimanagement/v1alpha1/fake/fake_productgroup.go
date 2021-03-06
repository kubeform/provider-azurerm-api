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

// FakeProductGroups implements ProductGroupInterface
type FakeProductGroups struct {
	Fake *FakeApimanagementV1alpha1
	ns   string
}

var productgroupsResource = schema.GroupVersionResource{Group: "apimanagement.azurerm.kubeform.com", Version: "v1alpha1", Resource: "productgroups"}

var productgroupsKind = schema.GroupVersionKind{Group: "apimanagement.azurerm.kubeform.com", Version: "v1alpha1", Kind: "ProductGroup"}

// Get takes name of the productGroup, and returns the corresponding productGroup object, and an error if there is any.
func (c *FakeProductGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ProductGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(productgroupsResource, c.ns, name), &v1alpha1.ProductGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProductGroup), err
}

// List takes label and field selectors, and returns the list of ProductGroups that match those selectors.
func (c *FakeProductGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ProductGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(productgroupsResource, productgroupsKind, c.ns, opts), &v1alpha1.ProductGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ProductGroupList{ListMeta: obj.(*v1alpha1.ProductGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.ProductGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested productGroups.
func (c *FakeProductGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(productgroupsResource, c.ns, opts))

}

// Create takes the representation of a productGroup and creates it.  Returns the server's representation of the productGroup, and an error, if there is any.
func (c *FakeProductGroups) Create(ctx context.Context, productGroup *v1alpha1.ProductGroup, opts v1.CreateOptions) (result *v1alpha1.ProductGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(productgroupsResource, c.ns, productGroup), &v1alpha1.ProductGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProductGroup), err
}

// Update takes the representation of a productGroup and updates it. Returns the server's representation of the productGroup, and an error, if there is any.
func (c *FakeProductGroups) Update(ctx context.Context, productGroup *v1alpha1.ProductGroup, opts v1.UpdateOptions) (result *v1alpha1.ProductGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(productgroupsResource, c.ns, productGroup), &v1alpha1.ProductGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProductGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProductGroups) UpdateStatus(ctx context.Context, productGroup *v1alpha1.ProductGroup, opts v1.UpdateOptions) (*v1alpha1.ProductGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(productgroupsResource, "status", c.ns, productGroup), &v1alpha1.ProductGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProductGroup), err
}

// Delete takes name of the productGroup and deletes it. Returns an error if one occurs.
func (c *FakeProductGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(productgroupsResource, c.ns, name), &v1alpha1.ProductGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProductGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(productgroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ProductGroupList{})
	return err
}

// Patch applies the patch and returns the patched productGroup.
func (c *FakeProductGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ProductGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(productgroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ProductGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProductGroup), err
}
