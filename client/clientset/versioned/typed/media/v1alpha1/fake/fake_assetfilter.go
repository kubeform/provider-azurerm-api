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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/media/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAssetFilters implements AssetFilterInterface
type FakeAssetFilters struct {
	Fake *FakeMediaV1alpha1
	ns   string
}

var assetfiltersResource = schema.GroupVersionResource{Group: "media.azurerm.kubeform.com", Version: "v1alpha1", Resource: "assetfilters"}

var assetfiltersKind = schema.GroupVersionKind{Group: "media.azurerm.kubeform.com", Version: "v1alpha1", Kind: "AssetFilter"}

// Get takes name of the assetFilter, and returns the corresponding assetFilter object, and an error if there is any.
func (c *FakeAssetFilters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AssetFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(assetfiltersResource, c.ns, name), &v1alpha1.AssetFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AssetFilter), err
}

// List takes label and field selectors, and returns the list of AssetFilters that match those selectors.
func (c *FakeAssetFilters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AssetFilterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(assetfiltersResource, assetfiltersKind, c.ns, opts), &v1alpha1.AssetFilterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AssetFilterList{ListMeta: obj.(*v1alpha1.AssetFilterList).ListMeta}
	for _, item := range obj.(*v1alpha1.AssetFilterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested assetFilters.
func (c *FakeAssetFilters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(assetfiltersResource, c.ns, opts))

}

// Create takes the representation of a assetFilter and creates it.  Returns the server's representation of the assetFilter, and an error, if there is any.
func (c *FakeAssetFilters) Create(ctx context.Context, assetFilter *v1alpha1.AssetFilter, opts v1.CreateOptions) (result *v1alpha1.AssetFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(assetfiltersResource, c.ns, assetFilter), &v1alpha1.AssetFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AssetFilter), err
}

// Update takes the representation of a assetFilter and updates it. Returns the server's representation of the assetFilter, and an error, if there is any.
func (c *FakeAssetFilters) Update(ctx context.Context, assetFilter *v1alpha1.AssetFilter, opts v1.UpdateOptions) (result *v1alpha1.AssetFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(assetfiltersResource, c.ns, assetFilter), &v1alpha1.AssetFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AssetFilter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAssetFilters) UpdateStatus(ctx context.Context, assetFilter *v1alpha1.AssetFilter, opts v1.UpdateOptions) (*v1alpha1.AssetFilter, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(assetfiltersResource, "status", c.ns, assetFilter), &v1alpha1.AssetFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AssetFilter), err
}

// Delete takes name of the assetFilter and deletes it. Returns an error if one occurs.
func (c *FakeAssetFilters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(assetfiltersResource, c.ns, name), &v1alpha1.AssetFilter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAssetFilters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(assetfiltersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AssetFilterList{})
	return err
}

// Patch applies the patch and returns the patched assetFilter.
func (c *FakeAssetFilters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AssetFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(assetfiltersResource, c.ns, name, pt, data, subresources...), &v1alpha1.AssetFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AssetFilter), err
}
