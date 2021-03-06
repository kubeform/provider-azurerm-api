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

// FakeDnsTxtRecords implements DnsTxtRecordInterface
type FakeDnsTxtRecords struct {
	Fake *FakePrivateV1alpha1
	ns   string
}

var dnstxtrecordsResource = schema.GroupVersionResource{Group: "private.azurerm.kubeform.com", Version: "v1alpha1", Resource: "dnstxtrecords"}

var dnstxtrecordsKind = schema.GroupVersionKind{Group: "private.azurerm.kubeform.com", Version: "v1alpha1", Kind: "DnsTxtRecord"}

// Get takes name of the dnsTxtRecord, and returns the corresponding dnsTxtRecord object, and an error if there is any.
func (c *FakeDnsTxtRecords) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DnsTxtRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dnstxtrecordsResource, c.ns, name), &v1alpha1.DnsTxtRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsTxtRecord), err
}

// List takes label and field selectors, and returns the list of DnsTxtRecords that match those selectors.
func (c *FakeDnsTxtRecords) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DnsTxtRecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dnstxtrecordsResource, dnstxtrecordsKind, c.ns, opts), &v1alpha1.DnsTxtRecordList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DnsTxtRecordList{ListMeta: obj.(*v1alpha1.DnsTxtRecordList).ListMeta}
	for _, item := range obj.(*v1alpha1.DnsTxtRecordList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dnsTxtRecords.
func (c *FakeDnsTxtRecords) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dnstxtrecordsResource, c.ns, opts))

}

// Create takes the representation of a dnsTxtRecord and creates it.  Returns the server's representation of the dnsTxtRecord, and an error, if there is any.
func (c *FakeDnsTxtRecords) Create(ctx context.Context, dnsTxtRecord *v1alpha1.DnsTxtRecord, opts v1.CreateOptions) (result *v1alpha1.DnsTxtRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dnstxtrecordsResource, c.ns, dnsTxtRecord), &v1alpha1.DnsTxtRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsTxtRecord), err
}

// Update takes the representation of a dnsTxtRecord and updates it. Returns the server's representation of the dnsTxtRecord, and an error, if there is any.
func (c *FakeDnsTxtRecords) Update(ctx context.Context, dnsTxtRecord *v1alpha1.DnsTxtRecord, opts v1.UpdateOptions) (result *v1alpha1.DnsTxtRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dnstxtrecordsResource, c.ns, dnsTxtRecord), &v1alpha1.DnsTxtRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsTxtRecord), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDnsTxtRecords) UpdateStatus(ctx context.Context, dnsTxtRecord *v1alpha1.DnsTxtRecord, opts v1.UpdateOptions) (*v1alpha1.DnsTxtRecord, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dnstxtrecordsResource, "status", c.ns, dnsTxtRecord), &v1alpha1.DnsTxtRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsTxtRecord), err
}

// Delete takes name of the dnsTxtRecord and deletes it. Returns an error if one occurs.
func (c *FakeDnsTxtRecords) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dnstxtrecordsResource, c.ns, name), &v1alpha1.DnsTxtRecord{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDnsTxtRecords) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dnstxtrecordsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DnsTxtRecordList{})
	return err
}

// Patch applies the patch and returns the patched dnsTxtRecord.
func (c *FakeDnsTxtRecords) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DnsTxtRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dnstxtrecordsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DnsTxtRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsTxtRecord), err
}
