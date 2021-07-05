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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/function/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAppSlots implements AppSlotInterface
type FakeAppSlots struct {
	Fake *FakeFunctionV1alpha1
	ns   string
}

var appslotsResource = schema.GroupVersionResource{Group: "function.azurerm.kubeform.com", Version: "v1alpha1", Resource: "appslots"}

var appslotsKind = schema.GroupVersionKind{Group: "function.azurerm.kubeform.com", Version: "v1alpha1", Kind: "AppSlot"}

// Get takes name of the appSlot, and returns the corresponding appSlot object, and an error if there is any.
func (c *FakeAppSlots) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AppSlot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(appslotsResource, c.ns, name), &v1alpha1.AppSlot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppSlot), err
}

// List takes label and field selectors, and returns the list of AppSlots that match those selectors.
func (c *FakeAppSlots) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AppSlotList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(appslotsResource, appslotsKind, c.ns, opts), &v1alpha1.AppSlotList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AppSlotList{ListMeta: obj.(*v1alpha1.AppSlotList).ListMeta}
	for _, item := range obj.(*v1alpha1.AppSlotList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appSlots.
func (c *FakeAppSlots) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(appslotsResource, c.ns, opts))

}

// Create takes the representation of a appSlot and creates it.  Returns the server's representation of the appSlot, and an error, if there is any.
func (c *FakeAppSlots) Create(ctx context.Context, appSlot *v1alpha1.AppSlot, opts v1.CreateOptions) (result *v1alpha1.AppSlot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(appslotsResource, c.ns, appSlot), &v1alpha1.AppSlot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppSlot), err
}

// Update takes the representation of a appSlot and updates it. Returns the server's representation of the appSlot, and an error, if there is any.
func (c *FakeAppSlots) Update(ctx context.Context, appSlot *v1alpha1.AppSlot, opts v1.UpdateOptions) (result *v1alpha1.AppSlot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(appslotsResource, c.ns, appSlot), &v1alpha1.AppSlot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppSlot), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAppSlots) UpdateStatus(ctx context.Context, appSlot *v1alpha1.AppSlot, opts v1.UpdateOptions) (*v1alpha1.AppSlot, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(appslotsResource, "status", c.ns, appSlot), &v1alpha1.AppSlot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppSlot), err
}

// Delete takes name of the appSlot and deletes it. Returns an error if one occurs.
func (c *FakeAppSlots) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(appslotsResource, c.ns, name), &v1alpha1.AppSlot{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAppSlots) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(appslotsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AppSlotList{})
	return err
}

// Patch applies the patch and returns the patched appSlot.
func (c *FakeAppSlots) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AppSlot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(appslotsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AppSlot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppSlot), err
}