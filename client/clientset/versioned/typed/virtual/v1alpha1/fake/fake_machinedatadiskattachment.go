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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/virtual/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMachineDataDiskAttachments implements MachineDataDiskAttachmentInterface
type FakeMachineDataDiskAttachments struct {
	Fake *FakeVirtualV1alpha1
	ns   string
}

var machinedatadiskattachmentsResource = schema.GroupVersionResource{Group: "virtual.azurerm.kubeform.com", Version: "v1alpha1", Resource: "machinedatadiskattachments"}

var machinedatadiskattachmentsKind = schema.GroupVersionKind{Group: "virtual.azurerm.kubeform.com", Version: "v1alpha1", Kind: "MachineDataDiskAttachment"}

// Get takes name of the machineDataDiskAttachment, and returns the corresponding machineDataDiskAttachment object, and an error if there is any.
func (c *FakeMachineDataDiskAttachments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MachineDataDiskAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(machinedatadiskattachmentsResource, c.ns, name), &v1alpha1.MachineDataDiskAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineDataDiskAttachment), err
}

// List takes label and field selectors, and returns the list of MachineDataDiskAttachments that match those selectors.
func (c *FakeMachineDataDiskAttachments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MachineDataDiskAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(machinedatadiskattachmentsResource, machinedatadiskattachmentsKind, c.ns, opts), &v1alpha1.MachineDataDiskAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MachineDataDiskAttachmentList{ListMeta: obj.(*v1alpha1.MachineDataDiskAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.MachineDataDiskAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested machineDataDiskAttachments.
func (c *FakeMachineDataDiskAttachments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(machinedatadiskattachmentsResource, c.ns, opts))

}

// Create takes the representation of a machineDataDiskAttachment and creates it.  Returns the server's representation of the machineDataDiskAttachment, and an error, if there is any.
func (c *FakeMachineDataDiskAttachments) Create(ctx context.Context, machineDataDiskAttachment *v1alpha1.MachineDataDiskAttachment, opts v1.CreateOptions) (result *v1alpha1.MachineDataDiskAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(machinedatadiskattachmentsResource, c.ns, machineDataDiskAttachment), &v1alpha1.MachineDataDiskAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineDataDiskAttachment), err
}

// Update takes the representation of a machineDataDiskAttachment and updates it. Returns the server's representation of the machineDataDiskAttachment, and an error, if there is any.
func (c *FakeMachineDataDiskAttachments) Update(ctx context.Context, machineDataDiskAttachment *v1alpha1.MachineDataDiskAttachment, opts v1.UpdateOptions) (result *v1alpha1.MachineDataDiskAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(machinedatadiskattachmentsResource, c.ns, machineDataDiskAttachment), &v1alpha1.MachineDataDiskAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineDataDiskAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMachineDataDiskAttachments) UpdateStatus(ctx context.Context, machineDataDiskAttachment *v1alpha1.MachineDataDiskAttachment, opts v1.UpdateOptions) (*v1alpha1.MachineDataDiskAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(machinedatadiskattachmentsResource, "status", c.ns, machineDataDiskAttachment), &v1alpha1.MachineDataDiskAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineDataDiskAttachment), err
}

// Delete takes name of the machineDataDiskAttachment and deletes it. Returns an error if one occurs.
func (c *FakeMachineDataDiskAttachments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(machinedatadiskattachmentsResource, c.ns, name), &v1alpha1.MachineDataDiskAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMachineDataDiskAttachments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(machinedatadiskattachmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MachineDataDiskAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched machineDataDiskAttachment.
func (c *FakeMachineDataDiskAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MachineDataDiskAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(machinedatadiskattachmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MachineDataDiskAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MachineDataDiskAttachment), err
}