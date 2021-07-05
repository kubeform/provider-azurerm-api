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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/backup/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePolicyFileShares implements PolicyFileShareInterface
type FakePolicyFileShares struct {
	Fake *FakeBackupV1alpha1
	ns   string
}

var policyfilesharesResource = schema.GroupVersionResource{Group: "backup.azurerm.kubeform.com", Version: "v1alpha1", Resource: "policyfileshares"}

var policyfilesharesKind = schema.GroupVersionKind{Group: "backup.azurerm.kubeform.com", Version: "v1alpha1", Kind: "PolicyFileShare"}

// Get takes name of the policyFileShare, and returns the corresponding policyFileShare object, and an error if there is any.
func (c *FakePolicyFileShares) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PolicyFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(policyfilesharesResource, c.ns, name), &v1alpha1.PolicyFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyFileShare), err
}

// List takes label and field selectors, and returns the list of PolicyFileShares that match those selectors.
func (c *FakePolicyFileShares) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PolicyFileShareList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(policyfilesharesResource, policyfilesharesKind, c.ns, opts), &v1alpha1.PolicyFileShareList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PolicyFileShareList{ListMeta: obj.(*v1alpha1.PolicyFileShareList).ListMeta}
	for _, item := range obj.(*v1alpha1.PolicyFileShareList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested policyFileShares.
func (c *FakePolicyFileShares) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(policyfilesharesResource, c.ns, opts))

}

// Create takes the representation of a policyFileShare and creates it.  Returns the server's representation of the policyFileShare, and an error, if there is any.
func (c *FakePolicyFileShares) Create(ctx context.Context, policyFileShare *v1alpha1.PolicyFileShare, opts v1.CreateOptions) (result *v1alpha1.PolicyFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(policyfilesharesResource, c.ns, policyFileShare), &v1alpha1.PolicyFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyFileShare), err
}

// Update takes the representation of a policyFileShare and updates it. Returns the server's representation of the policyFileShare, and an error, if there is any.
func (c *FakePolicyFileShares) Update(ctx context.Context, policyFileShare *v1alpha1.PolicyFileShare, opts v1.UpdateOptions) (result *v1alpha1.PolicyFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(policyfilesharesResource, c.ns, policyFileShare), &v1alpha1.PolicyFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyFileShare), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePolicyFileShares) UpdateStatus(ctx context.Context, policyFileShare *v1alpha1.PolicyFileShare, opts v1.UpdateOptions) (*v1alpha1.PolicyFileShare, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(policyfilesharesResource, "status", c.ns, policyFileShare), &v1alpha1.PolicyFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyFileShare), err
}

// Delete takes name of the policyFileShare and deletes it. Returns an error if one occurs.
func (c *FakePolicyFileShares) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(policyfilesharesResource, c.ns, name), &v1alpha1.PolicyFileShare{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePolicyFileShares) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(policyfilesharesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PolicyFileShareList{})
	return err
}

// Patch applies the patch and returns the patched policyFileShare.
func (c *FakePolicyFileShares) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PolicyFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(policyfilesharesResource, c.ns, name, pt, data, subresources...), &v1alpha1.PolicyFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyFileShare), err
}