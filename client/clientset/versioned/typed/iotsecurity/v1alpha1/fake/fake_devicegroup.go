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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/iotsecurity/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDeviceGroups implements DeviceGroupInterface
type FakeDeviceGroups struct {
	Fake *FakeIotsecurityV1alpha1
	ns   string
}

var devicegroupsResource = schema.GroupVersionResource{Group: "iotsecurity.azurerm.kubeform.com", Version: "v1alpha1", Resource: "devicegroups"}

var devicegroupsKind = schema.GroupVersionKind{Group: "iotsecurity.azurerm.kubeform.com", Version: "v1alpha1", Kind: "DeviceGroup"}

// Get takes name of the deviceGroup, and returns the corresponding deviceGroup object, and an error if there is any.
func (c *FakeDeviceGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DeviceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(devicegroupsResource, c.ns, name), &v1alpha1.DeviceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeviceGroup), err
}

// List takes label and field selectors, and returns the list of DeviceGroups that match those selectors.
func (c *FakeDeviceGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DeviceGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(devicegroupsResource, devicegroupsKind, c.ns, opts), &v1alpha1.DeviceGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DeviceGroupList{ListMeta: obj.(*v1alpha1.DeviceGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.DeviceGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested deviceGroups.
func (c *FakeDeviceGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(devicegroupsResource, c.ns, opts))

}

// Create takes the representation of a deviceGroup and creates it.  Returns the server's representation of the deviceGroup, and an error, if there is any.
func (c *FakeDeviceGroups) Create(ctx context.Context, deviceGroup *v1alpha1.DeviceGroup, opts v1.CreateOptions) (result *v1alpha1.DeviceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(devicegroupsResource, c.ns, deviceGroup), &v1alpha1.DeviceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeviceGroup), err
}

// Update takes the representation of a deviceGroup and updates it. Returns the server's representation of the deviceGroup, and an error, if there is any.
func (c *FakeDeviceGroups) Update(ctx context.Context, deviceGroup *v1alpha1.DeviceGroup, opts v1.UpdateOptions) (result *v1alpha1.DeviceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(devicegroupsResource, c.ns, deviceGroup), &v1alpha1.DeviceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeviceGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDeviceGroups) UpdateStatus(ctx context.Context, deviceGroup *v1alpha1.DeviceGroup, opts v1.UpdateOptions) (*v1alpha1.DeviceGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(devicegroupsResource, "status", c.ns, deviceGroup), &v1alpha1.DeviceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeviceGroup), err
}

// Delete takes name of the deviceGroup and deletes it. Returns an error if one occurs.
func (c *FakeDeviceGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(devicegroupsResource, c.ns, name), &v1alpha1.DeviceGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDeviceGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(devicegroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DeviceGroupList{})
	return err
}

// Patch applies the patch and returns the patched deviceGroup.
func (c *FakeDeviceGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DeviceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(devicegroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DeviceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeviceGroup), err
}
