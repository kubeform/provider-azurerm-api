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

// FakeShareFiles implements ShareFileInterface
type FakeShareFiles struct {
	Fake *FakeStorageV1alpha1
	ns   string
}

var sharefilesResource = schema.GroupVersionResource{Group: "storage.azurerm.kubeform.com", Version: "v1alpha1", Resource: "sharefiles"}

var sharefilesKind = schema.GroupVersionKind{Group: "storage.azurerm.kubeform.com", Version: "v1alpha1", Kind: "ShareFile"}

// Get takes name of the shareFile, and returns the corresponding shareFile object, and an error if there is any.
func (c *FakeShareFiles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ShareFile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sharefilesResource, c.ns, name), &v1alpha1.ShareFile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShareFile), err
}

// List takes label and field selectors, and returns the list of ShareFiles that match those selectors.
func (c *FakeShareFiles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ShareFileList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sharefilesResource, sharefilesKind, c.ns, opts), &v1alpha1.ShareFileList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ShareFileList{ListMeta: obj.(*v1alpha1.ShareFileList).ListMeta}
	for _, item := range obj.(*v1alpha1.ShareFileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested shareFiles.
func (c *FakeShareFiles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sharefilesResource, c.ns, opts))

}

// Create takes the representation of a shareFile and creates it.  Returns the server's representation of the shareFile, and an error, if there is any.
func (c *FakeShareFiles) Create(ctx context.Context, shareFile *v1alpha1.ShareFile, opts v1.CreateOptions) (result *v1alpha1.ShareFile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sharefilesResource, c.ns, shareFile), &v1alpha1.ShareFile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShareFile), err
}

// Update takes the representation of a shareFile and updates it. Returns the server's representation of the shareFile, and an error, if there is any.
func (c *FakeShareFiles) Update(ctx context.Context, shareFile *v1alpha1.ShareFile, opts v1.UpdateOptions) (result *v1alpha1.ShareFile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sharefilesResource, c.ns, shareFile), &v1alpha1.ShareFile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShareFile), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeShareFiles) UpdateStatus(ctx context.Context, shareFile *v1alpha1.ShareFile, opts v1.UpdateOptions) (*v1alpha1.ShareFile, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sharefilesResource, "status", c.ns, shareFile), &v1alpha1.ShareFile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShareFile), err
}

// Delete takes name of the shareFile and deletes it. Returns an error if one occurs.
func (c *FakeShareFiles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sharefilesResource, c.ns, name), &v1alpha1.ShareFile{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeShareFiles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sharefilesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ShareFileList{})
	return err
}

// Patch applies the patch and returns the patched shareFile.
func (c *FakeShareFiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ShareFile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sharefilesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ShareFile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShareFile), err
}
