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

// FakeShareDatasetDataLakeGen2s implements ShareDatasetDataLakeGen2Interface
type FakeShareDatasetDataLakeGen2s struct {
	Fake *FakeDataV1alpha1
	ns   string
}

var sharedatasetdatalakegen2sResource = schema.GroupVersionResource{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Resource: "sharedatasetdatalakegen2s"}

var sharedatasetdatalakegen2sKind = schema.GroupVersionKind{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Kind: "ShareDatasetDataLakeGen2"}

// Get takes name of the shareDatasetDataLakeGen2, and returns the corresponding shareDatasetDataLakeGen2 object, and an error if there is any.
func (c *FakeShareDatasetDataLakeGen2s) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ShareDatasetDataLakeGen2, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sharedatasetdatalakegen2sResource, c.ns, name), &v1alpha1.ShareDatasetDataLakeGen2{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShareDatasetDataLakeGen2), err
}

// List takes label and field selectors, and returns the list of ShareDatasetDataLakeGen2s that match those selectors.
func (c *FakeShareDatasetDataLakeGen2s) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ShareDatasetDataLakeGen2List, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sharedatasetdatalakegen2sResource, sharedatasetdatalakegen2sKind, c.ns, opts), &v1alpha1.ShareDatasetDataLakeGen2List{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ShareDatasetDataLakeGen2List{ListMeta: obj.(*v1alpha1.ShareDatasetDataLakeGen2List).ListMeta}
	for _, item := range obj.(*v1alpha1.ShareDatasetDataLakeGen2List).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested shareDatasetDataLakeGen2s.
func (c *FakeShareDatasetDataLakeGen2s) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sharedatasetdatalakegen2sResource, c.ns, opts))

}

// Create takes the representation of a shareDatasetDataLakeGen2 and creates it.  Returns the server's representation of the shareDatasetDataLakeGen2, and an error, if there is any.
func (c *FakeShareDatasetDataLakeGen2s) Create(ctx context.Context, shareDatasetDataLakeGen2 *v1alpha1.ShareDatasetDataLakeGen2, opts v1.CreateOptions) (result *v1alpha1.ShareDatasetDataLakeGen2, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sharedatasetdatalakegen2sResource, c.ns, shareDatasetDataLakeGen2), &v1alpha1.ShareDatasetDataLakeGen2{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShareDatasetDataLakeGen2), err
}

// Update takes the representation of a shareDatasetDataLakeGen2 and updates it. Returns the server's representation of the shareDatasetDataLakeGen2, and an error, if there is any.
func (c *FakeShareDatasetDataLakeGen2s) Update(ctx context.Context, shareDatasetDataLakeGen2 *v1alpha1.ShareDatasetDataLakeGen2, opts v1.UpdateOptions) (result *v1alpha1.ShareDatasetDataLakeGen2, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sharedatasetdatalakegen2sResource, c.ns, shareDatasetDataLakeGen2), &v1alpha1.ShareDatasetDataLakeGen2{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShareDatasetDataLakeGen2), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeShareDatasetDataLakeGen2s) UpdateStatus(ctx context.Context, shareDatasetDataLakeGen2 *v1alpha1.ShareDatasetDataLakeGen2, opts v1.UpdateOptions) (*v1alpha1.ShareDatasetDataLakeGen2, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sharedatasetdatalakegen2sResource, "status", c.ns, shareDatasetDataLakeGen2), &v1alpha1.ShareDatasetDataLakeGen2{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShareDatasetDataLakeGen2), err
}

// Delete takes name of the shareDatasetDataLakeGen2 and deletes it. Returns an error if one occurs.
func (c *FakeShareDatasetDataLakeGen2s) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sharedatasetdatalakegen2sResource, c.ns, name), &v1alpha1.ShareDatasetDataLakeGen2{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeShareDatasetDataLakeGen2s) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sharedatasetdatalakegen2sResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ShareDatasetDataLakeGen2List{})
	return err
}

// Patch applies the patch and returns the patched shareDatasetDataLakeGen2.
func (c *FakeShareDatasetDataLakeGen2s) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ShareDatasetDataLakeGen2, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sharedatasetdatalakegen2sResource, c.ns, name, pt, data, subresources...), &v1alpha1.ShareDatasetDataLakeGen2{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShareDatasetDataLakeGen2), err
}
