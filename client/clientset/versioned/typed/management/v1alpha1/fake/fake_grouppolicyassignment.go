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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/management/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGroupPolicyAssignments implements GroupPolicyAssignmentInterface
type FakeGroupPolicyAssignments struct {
	Fake *FakeManagementV1alpha1
	ns   string
}

var grouppolicyassignmentsResource = schema.GroupVersionResource{Group: "management.azurerm.kubeform.com", Version: "v1alpha1", Resource: "grouppolicyassignments"}

var grouppolicyassignmentsKind = schema.GroupVersionKind{Group: "management.azurerm.kubeform.com", Version: "v1alpha1", Kind: "GroupPolicyAssignment"}

// Get takes name of the groupPolicyAssignment, and returns the corresponding groupPolicyAssignment object, and an error if there is any.
func (c *FakeGroupPolicyAssignments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GroupPolicyAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(grouppolicyassignmentsResource, c.ns, name), &v1alpha1.GroupPolicyAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupPolicyAssignment), err
}

// List takes label and field selectors, and returns the list of GroupPolicyAssignments that match those selectors.
func (c *FakeGroupPolicyAssignments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GroupPolicyAssignmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(grouppolicyassignmentsResource, grouppolicyassignmentsKind, c.ns, opts), &v1alpha1.GroupPolicyAssignmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GroupPolicyAssignmentList{ListMeta: obj.(*v1alpha1.GroupPolicyAssignmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.GroupPolicyAssignmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested groupPolicyAssignments.
func (c *FakeGroupPolicyAssignments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(grouppolicyassignmentsResource, c.ns, opts))

}

// Create takes the representation of a groupPolicyAssignment and creates it.  Returns the server's representation of the groupPolicyAssignment, and an error, if there is any.
func (c *FakeGroupPolicyAssignments) Create(ctx context.Context, groupPolicyAssignment *v1alpha1.GroupPolicyAssignment, opts v1.CreateOptions) (result *v1alpha1.GroupPolicyAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(grouppolicyassignmentsResource, c.ns, groupPolicyAssignment), &v1alpha1.GroupPolicyAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupPolicyAssignment), err
}

// Update takes the representation of a groupPolicyAssignment and updates it. Returns the server's representation of the groupPolicyAssignment, and an error, if there is any.
func (c *FakeGroupPolicyAssignments) Update(ctx context.Context, groupPolicyAssignment *v1alpha1.GroupPolicyAssignment, opts v1.UpdateOptions) (result *v1alpha1.GroupPolicyAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(grouppolicyassignmentsResource, c.ns, groupPolicyAssignment), &v1alpha1.GroupPolicyAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupPolicyAssignment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGroupPolicyAssignments) UpdateStatus(ctx context.Context, groupPolicyAssignment *v1alpha1.GroupPolicyAssignment, opts v1.UpdateOptions) (*v1alpha1.GroupPolicyAssignment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(grouppolicyassignmentsResource, "status", c.ns, groupPolicyAssignment), &v1alpha1.GroupPolicyAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupPolicyAssignment), err
}

// Delete takes name of the groupPolicyAssignment and deletes it. Returns an error if one occurs.
func (c *FakeGroupPolicyAssignments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(grouppolicyassignmentsResource, c.ns, name), &v1alpha1.GroupPolicyAssignment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGroupPolicyAssignments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(grouppolicyassignmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.GroupPolicyAssignmentList{})
	return err
}

// Patch applies the patch and returns the patched groupPolicyAssignment.
func (c *FakeGroupPolicyAssignments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GroupPolicyAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(grouppolicyassignmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.GroupPolicyAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupPolicyAssignment), err
}
