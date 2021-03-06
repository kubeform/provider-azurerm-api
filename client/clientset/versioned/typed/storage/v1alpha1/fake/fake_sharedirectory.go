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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/storage/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeShareDirectories implements ShareDirectoryInterface
type FakeShareDirectories struct {
	Fake *FakeStorageV1alpha1
	ns   string
}

var sharedirectoriesResource = schema.GroupVersionResource{Group: "storage.azurerm.kubeform.com", Version: "v1alpha1", Resource: "sharedirectories"}

var sharedirectoriesKind = schema.GroupVersionKind{Group: "storage.azurerm.kubeform.com", Version: "v1alpha1", Kind: "ShareDirectory"}

// Get takes name of the shareDirectory, and returns the corresponding shareDirectory object, and an error if there is any.
func (c *FakeShareDirectories) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ShareDirectory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sharedirectoriesResource, c.ns, name), &v1alpha1.ShareDirectory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShareDirectory), err
}

// List takes label and field selectors, and returns the list of ShareDirectories that match those selectors.
func (c *FakeShareDirectories) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ShareDirectoryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sharedirectoriesResource, sharedirectoriesKind, c.ns, opts), &v1alpha1.ShareDirectoryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ShareDirectoryList{ListMeta: obj.(*v1alpha1.ShareDirectoryList).ListMeta}
	for _, item := range obj.(*v1alpha1.ShareDirectoryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested shareDirectories.
func (c *FakeShareDirectories) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sharedirectoriesResource, c.ns, opts))

}

// Create takes the representation of a shareDirectory and creates it.  Returns the server's representation of the shareDirectory, and an error, if there is any.
func (c *FakeShareDirectories) Create(ctx context.Context, shareDirectory *v1alpha1.ShareDirectory, opts v1.CreateOptions) (result *v1alpha1.ShareDirectory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sharedirectoriesResource, c.ns, shareDirectory), &v1alpha1.ShareDirectory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShareDirectory), err
}

// Update takes the representation of a shareDirectory and updates it. Returns the server's representation of the shareDirectory, and an error, if there is any.
func (c *FakeShareDirectories) Update(ctx context.Context, shareDirectory *v1alpha1.ShareDirectory, opts v1.UpdateOptions) (result *v1alpha1.ShareDirectory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sharedirectoriesResource, c.ns, shareDirectory), &v1alpha1.ShareDirectory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShareDirectory), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeShareDirectories) UpdateStatus(ctx context.Context, shareDirectory *v1alpha1.ShareDirectory, opts v1.UpdateOptions) (*v1alpha1.ShareDirectory, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sharedirectoriesResource, "status", c.ns, shareDirectory), &v1alpha1.ShareDirectory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShareDirectory), err
}

// Delete takes name of the shareDirectory and deletes it. Returns an error if one occurs.
func (c *FakeShareDirectories) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sharedirectoriesResource, c.ns, name), &v1alpha1.ShareDirectory{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeShareDirectories) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sharedirectoriesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ShareDirectoryList{})
	return err
}

// Patch applies the patch and returns the patched shareDirectory.
func (c *FakeShareDirectories) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ShareDirectory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sharedirectoriesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ShareDirectory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShareDirectory), err
}
