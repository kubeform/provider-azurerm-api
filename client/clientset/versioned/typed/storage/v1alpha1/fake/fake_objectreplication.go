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

// FakeObjectReplications implements ObjectReplicationInterface
type FakeObjectReplications struct {
	Fake *FakeStorageV1alpha1
	ns   string
}

var objectreplicationsResource = schema.GroupVersionResource{Group: "storage.azurerm.kubeform.com", Version: "v1alpha1", Resource: "objectreplications"}

var objectreplicationsKind = schema.GroupVersionKind{Group: "storage.azurerm.kubeform.com", Version: "v1alpha1", Kind: "ObjectReplication"}

// Get takes name of the objectReplication, and returns the corresponding objectReplication object, and an error if there is any.
func (c *FakeObjectReplications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ObjectReplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(objectreplicationsResource, c.ns, name), &v1alpha1.ObjectReplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectReplication), err
}

// List takes label and field selectors, and returns the list of ObjectReplications that match those selectors.
func (c *FakeObjectReplications) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ObjectReplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(objectreplicationsResource, objectreplicationsKind, c.ns, opts), &v1alpha1.ObjectReplicationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ObjectReplicationList{ListMeta: obj.(*v1alpha1.ObjectReplicationList).ListMeta}
	for _, item := range obj.(*v1alpha1.ObjectReplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested objectReplications.
func (c *FakeObjectReplications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(objectreplicationsResource, c.ns, opts))

}

// Create takes the representation of a objectReplication and creates it.  Returns the server's representation of the objectReplication, and an error, if there is any.
func (c *FakeObjectReplications) Create(ctx context.Context, objectReplication *v1alpha1.ObjectReplication, opts v1.CreateOptions) (result *v1alpha1.ObjectReplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(objectreplicationsResource, c.ns, objectReplication), &v1alpha1.ObjectReplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectReplication), err
}

// Update takes the representation of a objectReplication and updates it. Returns the server's representation of the objectReplication, and an error, if there is any.
func (c *FakeObjectReplications) Update(ctx context.Context, objectReplication *v1alpha1.ObjectReplication, opts v1.UpdateOptions) (result *v1alpha1.ObjectReplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(objectreplicationsResource, c.ns, objectReplication), &v1alpha1.ObjectReplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectReplication), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeObjectReplications) UpdateStatus(ctx context.Context, objectReplication *v1alpha1.ObjectReplication, opts v1.UpdateOptions) (*v1alpha1.ObjectReplication, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(objectreplicationsResource, "status", c.ns, objectReplication), &v1alpha1.ObjectReplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectReplication), err
}

// Delete takes name of the objectReplication and deletes it. Returns an error if one occurs.
func (c *FakeObjectReplications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(objectreplicationsResource, c.ns, name), &v1alpha1.ObjectReplication{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeObjectReplications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(objectreplicationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ObjectReplicationList{})
	return err
}

// Patch applies the patch and returns the patched objectReplication.
func (c *FakeObjectReplications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ObjectReplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(objectreplicationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ObjectReplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectReplication), err
}