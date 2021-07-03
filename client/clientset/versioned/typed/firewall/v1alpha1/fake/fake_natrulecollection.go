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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/firewall/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNatRuleCollections implements NatRuleCollectionInterface
type FakeNatRuleCollections struct {
	Fake *FakeFirewallV1alpha1
	ns   string
}

var natrulecollectionsResource = schema.GroupVersionResource{Group: "firewall.azurerm.kubeform.com", Version: "v1alpha1", Resource: "natrulecollections"}

var natrulecollectionsKind = schema.GroupVersionKind{Group: "firewall.azurerm.kubeform.com", Version: "v1alpha1", Kind: "NatRuleCollection"}

// Get takes name of the natRuleCollection, and returns the corresponding natRuleCollection object, and an error if there is any.
func (c *FakeNatRuleCollections) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NatRuleCollection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(natrulecollectionsResource, c.ns, name), &v1alpha1.NatRuleCollection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NatRuleCollection), err
}

// List takes label and field selectors, and returns the list of NatRuleCollections that match those selectors.
func (c *FakeNatRuleCollections) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NatRuleCollectionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(natrulecollectionsResource, natrulecollectionsKind, c.ns, opts), &v1alpha1.NatRuleCollectionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NatRuleCollectionList{ListMeta: obj.(*v1alpha1.NatRuleCollectionList).ListMeta}
	for _, item := range obj.(*v1alpha1.NatRuleCollectionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested natRuleCollections.
func (c *FakeNatRuleCollections) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(natrulecollectionsResource, c.ns, opts))

}

// Create takes the representation of a natRuleCollection and creates it.  Returns the server's representation of the natRuleCollection, and an error, if there is any.
func (c *FakeNatRuleCollections) Create(ctx context.Context, natRuleCollection *v1alpha1.NatRuleCollection, opts v1.CreateOptions) (result *v1alpha1.NatRuleCollection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(natrulecollectionsResource, c.ns, natRuleCollection), &v1alpha1.NatRuleCollection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NatRuleCollection), err
}

// Update takes the representation of a natRuleCollection and updates it. Returns the server's representation of the natRuleCollection, and an error, if there is any.
func (c *FakeNatRuleCollections) Update(ctx context.Context, natRuleCollection *v1alpha1.NatRuleCollection, opts v1.UpdateOptions) (result *v1alpha1.NatRuleCollection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(natrulecollectionsResource, c.ns, natRuleCollection), &v1alpha1.NatRuleCollection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NatRuleCollection), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNatRuleCollections) UpdateStatus(ctx context.Context, natRuleCollection *v1alpha1.NatRuleCollection, opts v1.UpdateOptions) (*v1alpha1.NatRuleCollection, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(natrulecollectionsResource, "status", c.ns, natRuleCollection), &v1alpha1.NatRuleCollection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NatRuleCollection), err
}

// Delete takes name of the natRuleCollection and deletes it. Returns an error if one occurs.
func (c *FakeNatRuleCollections) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(natrulecollectionsResource, c.ns, name), &v1alpha1.NatRuleCollection{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNatRuleCollections) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(natrulecollectionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NatRuleCollectionList{})
	return err
}

// Patch applies the patch and returns the patched natRuleCollection.
func (c *FakeNatRuleCollections) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NatRuleCollection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(natrulecollectionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.NatRuleCollection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NatRuleCollection), err
}
