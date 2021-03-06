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

// FakeBlobs implements BlobInterface
type FakeBlobs struct {
	Fake *FakeStorageV1alpha1
	ns   string
}

var blobsResource = schema.GroupVersionResource{Group: "storage.azurerm.kubeform.com", Version: "v1alpha1", Resource: "blobs"}

var blobsKind = schema.GroupVersionKind{Group: "storage.azurerm.kubeform.com", Version: "v1alpha1", Kind: "Blob"}

// Get takes name of the blob, and returns the corresponding blob object, and an error if there is any.
func (c *FakeBlobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Blob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(blobsResource, c.ns, name), &v1alpha1.Blob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Blob), err
}

// List takes label and field selectors, and returns the list of Blobs that match those selectors.
func (c *FakeBlobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BlobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(blobsResource, blobsKind, c.ns, opts), &v1alpha1.BlobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BlobList{ListMeta: obj.(*v1alpha1.BlobList).ListMeta}
	for _, item := range obj.(*v1alpha1.BlobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested blobs.
func (c *FakeBlobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(blobsResource, c.ns, opts))

}

// Create takes the representation of a blob and creates it.  Returns the server's representation of the blob, and an error, if there is any.
func (c *FakeBlobs) Create(ctx context.Context, blob *v1alpha1.Blob, opts v1.CreateOptions) (result *v1alpha1.Blob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(blobsResource, c.ns, blob), &v1alpha1.Blob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Blob), err
}

// Update takes the representation of a blob and updates it. Returns the server's representation of the blob, and an error, if there is any.
func (c *FakeBlobs) Update(ctx context.Context, blob *v1alpha1.Blob, opts v1.UpdateOptions) (result *v1alpha1.Blob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(blobsResource, c.ns, blob), &v1alpha1.Blob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Blob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBlobs) UpdateStatus(ctx context.Context, blob *v1alpha1.Blob, opts v1.UpdateOptions) (*v1alpha1.Blob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(blobsResource, "status", c.ns, blob), &v1alpha1.Blob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Blob), err
}

// Delete takes name of the blob and deletes it. Returns an error if one occurs.
func (c *FakeBlobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(blobsResource, c.ns, name), &v1alpha1.Blob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBlobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(blobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BlobList{})
	return err
}

// Patch applies the patch and returns the patched blob.
func (c *FakeBlobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Blob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(blobsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Blob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Blob), err
}
