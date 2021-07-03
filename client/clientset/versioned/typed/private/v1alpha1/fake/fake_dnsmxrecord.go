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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/private/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDnsMxRecords implements DnsMxRecordInterface
type FakeDnsMxRecords struct {
	Fake *FakePrivateV1alpha1
	ns   string
}

var dnsmxrecordsResource = schema.GroupVersionResource{Group: "private.azurerm.kubeform.com", Version: "v1alpha1", Resource: "dnsmxrecords"}

var dnsmxrecordsKind = schema.GroupVersionKind{Group: "private.azurerm.kubeform.com", Version: "v1alpha1", Kind: "DnsMxRecord"}

// Get takes name of the dnsMxRecord, and returns the corresponding dnsMxRecord object, and an error if there is any.
func (c *FakeDnsMxRecords) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DnsMxRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dnsmxrecordsResource, c.ns, name), &v1alpha1.DnsMxRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsMxRecord), err
}

// List takes label and field selectors, and returns the list of DnsMxRecords that match those selectors.
func (c *FakeDnsMxRecords) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DnsMxRecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dnsmxrecordsResource, dnsmxrecordsKind, c.ns, opts), &v1alpha1.DnsMxRecordList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DnsMxRecordList{ListMeta: obj.(*v1alpha1.DnsMxRecordList).ListMeta}
	for _, item := range obj.(*v1alpha1.DnsMxRecordList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dnsMxRecords.
func (c *FakeDnsMxRecords) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dnsmxrecordsResource, c.ns, opts))

}

// Create takes the representation of a dnsMxRecord and creates it.  Returns the server's representation of the dnsMxRecord, and an error, if there is any.
func (c *FakeDnsMxRecords) Create(ctx context.Context, dnsMxRecord *v1alpha1.DnsMxRecord, opts v1.CreateOptions) (result *v1alpha1.DnsMxRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dnsmxrecordsResource, c.ns, dnsMxRecord), &v1alpha1.DnsMxRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsMxRecord), err
}

// Update takes the representation of a dnsMxRecord and updates it. Returns the server's representation of the dnsMxRecord, and an error, if there is any.
func (c *FakeDnsMxRecords) Update(ctx context.Context, dnsMxRecord *v1alpha1.DnsMxRecord, opts v1.UpdateOptions) (result *v1alpha1.DnsMxRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dnsmxrecordsResource, c.ns, dnsMxRecord), &v1alpha1.DnsMxRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsMxRecord), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDnsMxRecords) UpdateStatus(ctx context.Context, dnsMxRecord *v1alpha1.DnsMxRecord, opts v1.UpdateOptions) (*v1alpha1.DnsMxRecord, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dnsmxrecordsResource, "status", c.ns, dnsMxRecord), &v1alpha1.DnsMxRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsMxRecord), err
}

// Delete takes name of the dnsMxRecord and deletes it. Returns an error if one occurs.
func (c *FakeDnsMxRecords) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dnsmxrecordsResource, c.ns, name), &v1alpha1.DnsMxRecord{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDnsMxRecords) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dnsmxrecordsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DnsMxRecordList{})
	return err
}

// Patch applies the patch and returns the patched dnsMxRecord.
func (c *FakeDnsMxRecords) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DnsMxRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dnsmxrecordsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DnsMxRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsMxRecord), err
}
