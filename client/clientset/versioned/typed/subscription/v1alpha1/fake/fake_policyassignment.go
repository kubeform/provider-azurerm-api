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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/subscription/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePolicyAssignments implements PolicyAssignmentInterface
type FakePolicyAssignments struct {
	Fake *FakeSubscriptionV1alpha1
	ns   string
}

var policyassignmentsResource = schema.GroupVersionResource{Group: "subscription.azurerm.kubeform.com", Version: "v1alpha1", Resource: "policyassignments"}

var policyassignmentsKind = schema.GroupVersionKind{Group: "subscription.azurerm.kubeform.com", Version: "v1alpha1", Kind: "PolicyAssignment"}

// Get takes name of the policyAssignment, and returns the corresponding policyAssignment object, and an error if there is any.
func (c *FakePolicyAssignments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PolicyAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(policyassignmentsResource, c.ns, name), &v1alpha1.PolicyAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyAssignment), err
}

// List takes label and field selectors, and returns the list of PolicyAssignments that match those selectors.
func (c *FakePolicyAssignments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PolicyAssignmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(policyassignmentsResource, policyassignmentsKind, c.ns, opts), &v1alpha1.PolicyAssignmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PolicyAssignmentList{ListMeta: obj.(*v1alpha1.PolicyAssignmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.PolicyAssignmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested policyAssignments.
func (c *FakePolicyAssignments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(policyassignmentsResource, c.ns, opts))

}

// Create takes the representation of a policyAssignment and creates it.  Returns the server's representation of the policyAssignment, and an error, if there is any.
func (c *FakePolicyAssignments) Create(ctx context.Context, policyAssignment *v1alpha1.PolicyAssignment, opts v1.CreateOptions) (result *v1alpha1.PolicyAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(policyassignmentsResource, c.ns, policyAssignment), &v1alpha1.PolicyAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyAssignment), err
}

// Update takes the representation of a policyAssignment and updates it. Returns the server's representation of the policyAssignment, and an error, if there is any.
func (c *FakePolicyAssignments) Update(ctx context.Context, policyAssignment *v1alpha1.PolicyAssignment, opts v1.UpdateOptions) (result *v1alpha1.PolicyAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(policyassignmentsResource, c.ns, policyAssignment), &v1alpha1.PolicyAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyAssignment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePolicyAssignments) UpdateStatus(ctx context.Context, policyAssignment *v1alpha1.PolicyAssignment, opts v1.UpdateOptions) (*v1alpha1.PolicyAssignment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(policyassignmentsResource, "status", c.ns, policyAssignment), &v1alpha1.PolicyAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyAssignment), err
}

// Delete takes name of the policyAssignment and deletes it. Returns an error if one occurs.
func (c *FakePolicyAssignments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(policyassignmentsResource, c.ns, name), &v1alpha1.PolicyAssignment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePolicyAssignments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(policyassignmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PolicyAssignmentList{})
	return err
}

// Patch applies the patch and returns the patched policyAssignment.
func (c *FakePolicyAssignments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PolicyAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(policyassignmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PolicyAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyAssignment), err
}