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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/dns/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePtrRecords implements PtrRecordInterface
type FakePtrRecords struct {
	Fake *FakeDnsV1alpha1
	ns   string
}

var ptrrecordsResource = schema.GroupVersionResource{Group: "dns.azurerm.kubeform.com", Version: "v1alpha1", Resource: "ptrrecords"}

var ptrrecordsKind = schema.GroupVersionKind{Group: "dns.azurerm.kubeform.com", Version: "v1alpha1", Kind: "PtrRecord"}

// Get takes name of the ptrRecord, and returns the corresponding ptrRecord object, and an error if there is any.
func (c *FakePtrRecords) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PtrRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ptrrecordsResource, c.ns, name), &v1alpha1.PtrRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PtrRecord), err
}

// List takes label and field selectors, and returns the list of PtrRecords that match those selectors.
func (c *FakePtrRecords) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PtrRecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ptrrecordsResource, ptrrecordsKind, c.ns, opts), &v1alpha1.PtrRecordList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PtrRecordList{ListMeta: obj.(*v1alpha1.PtrRecordList).ListMeta}
	for _, item := range obj.(*v1alpha1.PtrRecordList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ptrRecords.
func (c *FakePtrRecords) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ptrrecordsResource, c.ns, opts))

}

// Create takes the representation of a ptrRecord and creates it.  Returns the server's representation of the ptrRecord, and an error, if there is any.
func (c *FakePtrRecords) Create(ctx context.Context, ptrRecord *v1alpha1.PtrRecord, opts v1.CreateOptions) (result *v1alpha1.PtrRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ptrrecordsResource, c.ns, ptrRecord), &v1alpha1.PtrRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PtrRecord), err
}

// Update takes the representation of a ptrRecord and updates it. Returns the server's representation of the ptrRecord, and an error, if there is any.
func (c *FakePtrRecords) Update(ctx context.Context, ptrRecord *v1alpha1.PtrRecord, opts v1.UpdateOptions) (result *v1alpha1.PtrRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ptrrecordsResource, c.ns, ptrRecord), &v1alpha1.PtrRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PtrRecord), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePtrRecords) UpdateStatus(ctx context.Context, ptrRecord *v1alpha1.PtrRecord, opts v1.UpdateOptions) (*v1alpha1.PtrRecord, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ptrrecordsResource, "status", c.ns, ptrRecord), &v1alpha1.PtrRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PtrRecord), err
}

// Delete takes name of the ptrRecord and deletes it. Returns an error if one occurs.
func (c *FakePtrRecords) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ptrrecordsResource, c.ns, name), &v1alpha1.PtrRecord{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePtrRecords) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ptrrecordsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PtrRecordList{})
	return err
}

// Patch applies the patch and returns the patched ptrRecord.
func (c *FakePtrRecords) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PtrRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ptrrecordsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PtrRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PtrRecord), err
}
