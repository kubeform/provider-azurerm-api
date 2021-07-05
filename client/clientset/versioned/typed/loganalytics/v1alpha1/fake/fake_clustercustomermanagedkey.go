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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/loganalytics/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterCustomerManagedKeys implements ClusterCustomerManagedKeyInterface
type FakeClusterCustomerManagedKeys struct {
	Fake *FakeLoganalyticsV1alpha1
	ns   string
}

var clustercustomermanagedkeysResource = schema.GroupVersionResource{Group: "loganalytics.azurerm.kubeform.com", Version: "v1alpha1", Resource: "clustercustomermanagedkeys"}

var clustercustomermanagedkeysKind = schema.GroupVersionKind{Group: "loganalytics.azurerm.kubeform.com", Version: "v1alpha1", Kind: "ClusterCustomerManagedKey"}

// Get takes name of the clusterCustomerManagedKey, and returns the corresponding clusterCustomerManagedKey object, and an error if there is any.
func (c *FakeClusterCustomerManagedKeys) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterCustomerManagedKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clustercustomermanagedkeysResource, c.ns, name), &v1alpha1.ClusterCustomerManagedKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterCustomerManagedKey), err
}

// List takes label and field selectors, and returns the list of ClusterCustomerManagedKeys that match those selectors.
func (c *FakeClusterCustomerManagedKeys) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterCustomerManagedKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clustercustomermanagedkeysResource, clustercustomermanagedkeysKind, c.ns, opts), &v1alpha1.ClusterCustomerManagedKeyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterCustomerManagedKeyList{ListMeta: obj.(*v1alpha1.ClusterCustomerManagedKeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterCustomerManagedKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterCustomerManagedKeys.
func (c *FakeClusterCustomerManagedKeys) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clustercustomermanagedkeysResource, c.ns, opts))

}

// Create takes the representation of a clusterCustomerManagedKey and creates it.  Returns the server's representation of the clusterCustomerManagedKey, and an error, if there is any.
func (c *FakeClusterCustomerManagedKeys) Create(ctx context.Context, clusterCustomerManagedKey *v1alpha1.ClusterCustomerManagedKey, opts v1.CreateOptions) (result *v1alpha1.ClusterCustomerManagedKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clustercustomermanagedkeysResource, c.ns, clusterCustomerManagedKey), &v1alpha1.ClusterCustomerManagedKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterCustomerManagedKey), err
}

// Update takes the representation of a clusterCustomerManagedKey and updates it. Returns the server's representation of the clusterCustomerManagedKey, and an error, if there is any.
func (c *FakeClusterCustomerManagedKeys) Update(ctx context.Context, clusterCustomerManagedKey *v1alpha1.ClusterCustomerManagedKey, opts v1.UpdateOptions) (result *v1alpha1.ClusterCustomerManagedKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clustercustomermanagedkeysResource, c.ns, clusterCustomerManagedKey), &v1alpha1.ClusterCustomerManagedKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterCustomerManagedKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterCustomerManagedKeys) UpdateStatus(ctx context.Context, clusterCustomerManagedKey *v1alpha1.ClusterCustomerManagedKey, opts v1.UpdateOptions) (*v1alpha1.ClusterCustomerManagedKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clustercustomermanagedkeysResource, "status", c.ns, clusterCustomerManagedKey), &v1alpha1.ClusterCustomerManagedKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterCustomerManagedKey), err
}

// Delete takes name of the clusterCustomerManagedKey and deletes it. Returns an error if one occurs.
func (c *FakeClusterCustomerManagedKeys) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clustercustomermanagedkeysResource, c.ns, name), &v1alpha1.ClusterCustomerManagedKey{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterCustomerManagedKeys) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clustercustomermanagedkeysResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterCustomerManagedKeyList{})
	return err
}

// Patch applies the patch and returns the patched clusterCustomerManagedKey.
func (c *FakeClusterCustomerManagedKeys) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterCustomerManagedKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clustercustomermanagedkeysResource, c.ns, name, pt, data, subresources...), &v1alpha1.ClusterCustomerManagedKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterCustomerManagedKey), err
}