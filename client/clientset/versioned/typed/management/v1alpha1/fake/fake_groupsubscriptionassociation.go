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

// FakeGroupSubscriptionAssociations implements GroupSubscriptionAssociationInterface
type FakeGroupSubscriptionAssociations struct {
	Fake *FakeManagementV1alpha1
	ns   string
}

var groupsubscriptionassociationsResource = schema.GroupVersionResource{Group: "management.azurerm.kubeform.com", Version: "v1alpha1", Resource: "groupsubscriptionassociations"}

var groupsubscriptionassociationsKind = schema.GroupVersionKind{Group: "management.azurerm.kubeform.com", Version: "v1alpha1", Kind: "GroupSubscriptionAssociation"}

// Get takes name of the groupSubscriptionAssociation, and returns the corresponding groupSubscriptionAssociation object, and an error if there is any.
func (c *FakeGroupSubscriptionAssociations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GroupSubscriptionAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(groupsubscriptionassociationsResource, c.ns, name), &v1alpha1.GroupSubscriptionAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupSubscriptionAssociation), err
}

// List takes label and field selectors, and returns the list of GroupSubscriptionAssociations that match those selectors.
func (c *FakeGroupSubscriptionAssociations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GroupSubscriptionAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(groupsubscriptionassociationsResource, groupsubscriptionassociationsKind, c.ns, opts), &v1alpha1.GroupSubscriptionAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GroupSubscriptionAssociationList{ListMeta: obj.(*v1alpha1.GroupSubscriptionAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.GroupSubscriptionAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested groupSubscriptionAssociations.
func (c *FakeGroupSubscriptionAssociations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(groupsubscriptionassociationsResource, c.ns, opts))

}

// Create takes the representation of a groupSubscriptionAssociation and creates it.  Returns the server's representation of the groupSubscriptionAssociation, and an error, if there is any.
func (c *FakeGroupSubscriptionAssociations) Create(ctx context.Context, groupSubscriptionAssociation *v1alpha1.GroupSubscriptionAssociation, opts v1.CreateOptions) (result *v1alpha1.GroupSubscriptionAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(groupsubscriptionassociationsResource, c.ns, groupSubscriptionAssociation), &v1alpha1.GroupSubscriptionAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupSubscriptionAssociation), err
}

// Update takes the representation of a groupSubscriptionAssociation and updates it. Returns the server's representation of the groupSubscriptionAssociation, and an error, if there is any.
func (c *FakeGroupSubscriptionAssociations) Update(ctx context.Context, groupSubscriptionAssociation *v1alpha1.GroupSubscriptionAssociation, opts v1.UpdateOptions) (result *v1alpha1.GroupSubscriptionAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(groupsubscriptionassociationsResource, c.ns, groupSubscriptionAssociation), &v1alpha1.GroupSubscriptionAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupSubscriptionAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGroupSubscriptionAssociations) UpdateStatus(ctx context.Context, groupSubscriptionAssociation *v1alpha1.GroupSubscriptionAssociation, opts v1.UpdateOptions) (*v1alpha1.GroupSubscriptionAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(groupsubscriptionassociationsResource, "status", c.ns, groupSubscriptionAssociation), &v1alpha1.GroupSubscriptionAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupSubscriptionAssociation), err
}

// Delete takes name of the groupSubscriptionAssociation and deletes it. Returns an error if one occurs.
func (c *FakeGroupSubscriptionAssociations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(groupsubscriptionassociationsResource, c.ns, name), &v1alpha1.GroupSubscriptionAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGroupSubscriptionAssociations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(groupsubscriptionassociationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.GroupSubscriptionAssociationList{})
	return err
}

// Patch applies the patch and returns the patched groupSubscriptionAssociation.
func (c *FakeGroupSubscriptionAssociations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GroupSubscriptionAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(groupsubscriptionassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.GroupSubscriptionAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupSubscriptionAssociation), err
}
