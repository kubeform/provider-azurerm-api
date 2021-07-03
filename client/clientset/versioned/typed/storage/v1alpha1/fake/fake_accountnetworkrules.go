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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/storage/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAccountNetworkRuleses implements AccountNetworkRulesInterface
type FakeAccountNetworkRuleses struct {
	Fake *FakeStorageV1alpha1
	ns   string
}

var accountnetworkrulesesResource = schema.GroupVersionResource{Group: "storage.azurerm.kubeform.com", Version: "v1alpha1", Resource: "accountnetworkruleses"}

var accountnetworkrulesesKind = schema.GroupVersionKind{Group: "storage.azurerm.kubeform.com", Version: "v1alpha1", Kind: "AccountNetworkRules"}

// Get takes name of the accountNetworkRules, and returns the corresponding accountNetworkRules object, and an error if there is any.
func (c *FakeAccountNetworkRuleses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AccountNetworkRules, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(accountnetworkrulesesResource, c.ns, name), &v1alpha1.AccountNetworkRules{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountNetworkRules), err
}

// List takes label and field selectors, and returns the list of AccountNetworkRuleses that match those selectors.
func (c *FakeAccountNetworkRuleses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AccountNetworkRulesList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(accountnetworkrulesesResource, accountnetworkrulesesKind, c.ns, opts), &v1alpha1.AccountNetworkRulesList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AccountNetworkRulesList{ListMeta: obj.(*v1alpha1.AccountNetworkRulesList).ListMeta}
	for _, item := range obj.(*v1alpha1.AccountNetworkRulesList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested accountNetworkRuleses.
func (c *FakeAccountNetworkRuleses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(accountnetworkrulesesResource, c.ns, opts))

}

// Create takes the representation of a accountNetworkRules and creates it.  Returns the server's representation of the accountNetworkRules, and an error, if there is any.
func (c *FakeAccountNetworkRuleses) Create(ctx context.Context, accountNetworkRules *v1alpha1.AccountNetworkRules, opts v1.CreateOptions) (result *v1alpha1.AccountNetworkRules, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(accountnetworkrulesesResource, c.ns, accountNetworkRules), &v1alpha1.AccountNetworkRules{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountNetworkRules), err
}

// Update takes the representation of a accountNetworkRules and updates it. Returns the server's representation of the accountNetworkRules, and an error, if there is any.
func (c *FakeAccountNetworkRuleses) Update(ctx context.Context, accountNetworkRules *v1alpha1.AccountNetworkRules, opts v1.UpdateOptions) (result *v1alpha1.AccountNetworkRules, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(accountnetworkrulesesResource, c.ns, accountNetworkRules), &v1alpha1.AccountNetworkRules{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountNetworkRules), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAccountNetworkRuleses) UpdateStatus(ctx context.Context, accountNetworkRules *v1alpha1.AccountNetworkRules, opts v1.UpdateOptions) (*v1alpha1.AccountNetworkRules, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(accountnetworkrulesesResource, "status", c.ns, accountNetworkRules), &v1alpha1.AccountNetworkRules{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountNetworkRules), err
}

// Delete takes name of the accountNetworkRules and deletes it. Returns an error if one occurs.
func (c *FakeAccountNetworkRuleses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(accountnetworkrulesesResource, c.ns, name), &v1alpha1.AccountNetworkRules{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAccountNetworkRuleses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(accountnetworkrulesesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AccountNetworkRulesList{})
	return err
}

// Patch applies the patch and returns the patched accountNetworkRules.
func (c *FakeAccountNetworkRuleses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AccountNetworkRules, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(accountnetworkrulesesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AccountNetworkRules{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountNetworkRules), err
}
