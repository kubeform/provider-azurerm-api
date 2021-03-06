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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/security/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCenterAutomations implements CenterAutomationInterface
type FakeCenterAutomations struct {
	Fake *FakeSecurityV1alpha1
	ns   string
}

var centerautomationsResource = schema.GroupVersionResource{Group: "security.azurerm.kubeform.com", Version: "v1alpha1", Resource: "centerautomations"}

var centerautomationsKind = schema.GroupVersionKind{Group: "security.azurerm.kubeform.com", Version: "v1alpha1", Kind: "CenterAutomation"}

// Get takes name of the centerAutomation, and returns the corresponding centerAutomation object, and an error if there is any.
func (c *FakeCenterAutomations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CenterAutomation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(centerautomationsResource, c.ns, name), &v1alpha1.CenterAutomation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CenterAutomation), err
}

// List takes label and field selectors, and returns the list of CenterAutomations that match those selectors.
func (c *FakeCenterAutomations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CenterAutomationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(centerautomationsResource, centerautomationsKind, c.ns, opts), &v1alpha1.CenterAutomationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CenterAutomationList{ListMeta: obj.(*v1alpha1.CenterAutomationList).ListMeta}
	for _, item := range obj.(*v1alpha1.CenterAutomationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested centerAutomations.
func (c *FakeCenterAutomations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(centerautomationsResource, c.ns, opts))

}

// Create takes the representation of a centerAutomation and creates it.  Returns the server's representation of the centerAutomation, and an error, if there is any.
func (c *FakeCenterAutomations) Create(ctx context.Context, centerAutomation *v1alpha1.CenterAutomation, opts v1.CreateOptions) (result *v1alpha1.CenterAutomation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(centerautomationsResource, c.ns, centerAutomation), &v1alpha1.CenterAutomation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CenterAutomation), err
}

// Update takes the representation of a centerAutomation and updates it. Returns the server's representation of the centerAutomation, and an error, if there is any.
func (c *FakeCenterAutomations) Update(ctx context.Context, centerAutomation *v1alpha1.CenterAutomation, opts v1.UpdateOptions) (result *v1alpha1.CenterAutomation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(centerautomationsResource, c.ns, centerAutomation), &v1alpha1.CenterAutomation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CenterAutomation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCenterAutomations) UpdateStatus(ctx context.Context, centerAutomation *v1alpha1.CenterAutomation, opts v1.UpdateOptions) (*v1alpha1.CenterAutomation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(centerautomationsResource, "status", c.ns, centerAutomation), &v1alpha1.CenterAutomation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CenterAutomation), err
}

// Delete takes name of the centerAutomation and deletes it. Returns an error if one occurs.
func (c *FakeCenterAutomations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(centerautomationsResource, c.ns, name), &v1alpha1.CenterAutomation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCenterAutomations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(centerautomationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CenterAutomationList{})
	return err
}

// Patch applies the patch and returns the patched centerAutomation.
func (c *FakeCenterAutomations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CenterAutomation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(centerautomationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CenterAutomation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CenterAutomation), err
}
