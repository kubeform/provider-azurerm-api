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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/mssql/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFirewallRules implements FirewallRuleInterface
type FakeFirewallRules struct {
	Fake *FakeMssqlV1alpha1
	ns   string
}

var firewallrulesResource = schema.GroupVersionResource{Group: "mssql.azurerm.kubeform.com", Version: "v1alpha1", Resource: "firewallrules"}

var firewallrulesKind = schema.GroupVersionKind{Group: "mssql.azurerm.kubeform.com", Version: "v1alpha1", Kind: "FirewallRule"}

// Get takes name of the firewallRule, and returns the corresponding firewallRule object, and an error if there is any.
func (c *FakeFirewallRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(firewallrulesResource, c.ns, name), &v1alpha1.FirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallRule), err
}

// List takes label and field selectors, and returns the list of FirewallRules that match those selectors.
func (c *FakeFirewallRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FirewallRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(firewallrulesResource, firewallrulesKind, c.ns, opts), &v1alpha1.FirewallRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FirewallRuleList{ListMeta: obj.(*v1alpha1.FirewallRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.FirewallRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested firewallRules.
func (c *FakeFirewallRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(firewallrulesResource, c.ns, opts))

}

// Create takes the representation of a firewallRule and creates it.  Returns the server's representation of the firewallRule, and an error, if there is any.
func (c *FakeFirewallRules) Create(ctx context.Context, firewallRule *v1alpha1.FirewallRule, opts v1.CreateOptions) (result *v1alpha1.FirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(firewallrulesResource, c.ns, firewallRule), &v1alpha1.FirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallRule), err
}

// Update takes the representation of a firewallRule and updates it. Returns the server's representation of the firewallRule, and an error, if there is any.
func (c *FakeFirewallRules) Update(ctx context.Context, firewallRule *v1alpha1.FirewallRule, opts v1.UpdateOptions) (result *v1alpha1.FirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(firewallrulesResource, c.ns, firewallRule), &v1alpha1.FirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFirewallRules) UpdateStatus(ctx context.Context, firewallRule *v1alpha1.FirewallRule, opts v1.UpdateOptions) (*v1alpha1.FirewallRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(firewallrulesResource, "status", c.ns, firewallRule), &v1alpha1.FirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallRule), err
}

// Delete takes name of the firewallRule and deletes it. Returns an error if one occurs.
func (c *FakeFirewallRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(firewallrulesResource, c.ns, name), &v1alpha1.FirewallRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFirewallRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(firewallrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FirewallRuleList{})
	return err
}

// Patch applies the patch and returns the patched firewallRule.
func (c *FakeFirewallRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(firewallrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.FirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallRule), err
}
