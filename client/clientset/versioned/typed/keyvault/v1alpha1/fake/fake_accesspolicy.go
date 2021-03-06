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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/keyvault/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAccessPolicies implements AccessPolicyInterface
type FakeAccessPolicies struct {
	Fake *FakeKeyvaultV1alpha1
	ns   string
}

var accesspoliciesResource = schema.GroupVersionResource{Group: "keyvault.azurerm.kubeform.com", Version: "v1alpha1", Resource: "accesspolicies"}

var accesspoliciesKind = schema.GroupVersionKind{Group: "keyvault.azurerm.kubeform.com", Version: "v1alpha1", Kind: "AccessPolicy"}

// Get takes name of the accessPolicy, and returns the corresponding accessPolicy object, and an error if there is any.
func (c *FakeAccessPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AccessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(accesspoliciesResource, c.ns, name), &v1alpha1.AccessPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessPolicy), err
}

// List takes label and field selectors, and returns the list of AccessPolicies that match those selectors.
func (c *FakeAccessPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AccessPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(accesspoliciesResource, accesspoliciesKind, c.ns, opts), &v1alpha1.AccessPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AccessPolicyList{ListMeta: obj.(*v1alpha1.AccessPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.AccessPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested accessPolicies.
func (c *FakeAccessPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(accesspoliciesResource, c.ns, opts))

}

// Create takes the representation of a accessPolicy and creates it.  Returns the server's representation of the accessPolicy, and an error, if there is any.
func (c *FakeAccessPolicies) Create(ctx context.Context, accessPolicy *v1alpha1.AccessPolicy, opts v1.CreateOptions) (result *v1alpha1.AccessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(accesspoliciesResource, c.ns, accessPolicy), &v1alpha1.AccessPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessPolicy), err
}

// Update takes the representation of a accessPolicy and updates it. Returns the server's representation of the accessPolicy, and an error, if there is any.
func (c *FakeAccessPolicies) Update(ctx context.Context, accessPolicy *v1alpha1.AccessPolicy, opts v1.UpdateOptions) (result *v1alpha1.AccessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(accesspoliciesResource, c.ns, accessPolicy), &v1alpha1.AccessPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAccessPolicies) UpdateStatus(ctx context.Context, accessPolicy *v1alpha1.AccessPolicy, opts v1.UpdateOptions) (*v1alpha1.AccessPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(accesspoliciesResource, "status", c.ns, accessPolicy), &v1alpha1.AccessPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessPolicy), err
}

// Delete takes name of the accessPolicy and deletes it. Returns an error if one occurs.
func (c *FakeAccessPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(accesspoliciesResource, c.ns, name), &v1alpha1.AccessPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAccessPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(accesspoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AccessPolicyList{})
	return err
}

// Patch applies the patch and returns the patched accessPolicy.
func (c *FakeAccessPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AccessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(accesspoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AccessPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessPolicy), err
}
