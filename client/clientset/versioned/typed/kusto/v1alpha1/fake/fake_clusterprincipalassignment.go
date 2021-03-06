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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/kusto/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterPrincipalAssignments implements ClusterPrincipalAssignmentInterface
type FakeClusterPrincipalAssignments struct {
	Fake *FakeKustoV1alpha1
	ns   string
}

var clusterprincipalassignmentsResource = schema.GroupVersionResource{Group: "kusto.azurerm.kubeform.com", Version: "v1alpha1", Resource: "clusterprincipalassignments"}

var clusterprincipalassignmentsKind = schema.GroupVersionKind{Group: "kusto.azurerm.kubeform.com", Version: "v1alpha1", Kind: "ClusterPrincipalAssignment"}

// Get takes name of the clusterPrincipalAssignment, and returns the corresponding clusterPrincipalAssignment object, and an error if there is any.
func (c *FakeClusterPrincipalAssignments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterPrincipalAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clusterprincipalassignmentsResource, c.ns, name), &v1alpha1.ClusterPrincipalAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPrincipalAssignment), err
}

// List takes label and field selectors, and returns the list of ClusterPrincipalAssignments that match those selectors.
func (c *FakeClusterPrincipalAssignments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterPrincipalAssignmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clusterprincipalassignmentsResource, clusterprincipalassignmentsKind, c.ns, opts), &v1alpha1.ClusterPrincipalAssignmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterPrincipalAssignmentList{ListMeta: obj.(*v1alpha1.ClusterPrincipalAssignmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterPrincipalAssignmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterPrincipalAssignments.
func (c *FakeClusterPrincipalAssignments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clusterprincipalassignmentsResource, c.ns, opts))

}

// Create takes the representation of a clusterPrincipalAssignment and creates it.  Returns the server's representation of the clusterPrincipalAssignment, and an error, if there is any.
func (c *FakeClusterPrincipalAssignments) Create(ctx context.Context, clusterPrincipalAssignment *v1alpha1.ClusterPrincipalAssignment, opts v1.CreateOptions) (result *v1alpha1.ClusterPrincipalAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clusterprincipalassignmentsResource, c.ns, clusterPrincipalAssignment), &v1alpha1.ClusterPrincipalAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPrincipalAssignment), err
}

// Update takes the representation of a clusterPrincipalAssignment and updates it. Returns the server's representation of the clusterPrincipalAssignment, and an error, if there is any.
func (c *FakeClusterPrincipalAssignments) Update(ctx context.Context, clusterPrincipalAssignment *v1alpha1.ClusterPrincipalAssignment, opts v1.UpdateOptions) (result *v1alpha1.ClusterPrincipalAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clusterprincipalassignmentsResource, c.ns, clusterPrincipalAssignment), &v1alpha1.ClusterPrincipalAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPrincipalAssignment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterPrincipalAssignments) UpdateStatus(ctx context.Context, clusterPrincipalAssignment *v1alpha1.ClusterPrincipalAssignment, opts v1.UpdateOptions) (*v1alpha1.ClusterPrincipalAssignment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clusterprincipalassignmentsResource, "status", c.ns, clusterPrincipalAssignment), &v1alpha1.ClusterPrincipalAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPrincipalAssignment), err
}

// Delete takes name of the clusterPrincipalAssignment and deletes it. Returns an error if one occurs.
func (c *FakeClusterPrincipalAssignments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clusterprincipalassignmentsResource, c.ns, name), &v1alpha1.ClusterPrincipalAssignment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterPrincipalAssignments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clusterprincipalassignmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterPrincipalAssignmentList{})
	return err
}

// Patch applies the patch and returns the patched clusterPrincipalAssignment.
func (c *FakeClusterPrincipalAssignments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterPrincipalAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clusterprincipalassignmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ClusterPrincipalAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPrincipalAssignment), err
}
