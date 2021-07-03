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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/notificationhub/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAuthorizationRules implements AuthorizationRuleInterface
type FakeAuthorizationRules struct {
	Fake *FakeNotificationhubV1alpha1
	ns   string
}

var authorizationrulesResource = schema.GroupVersionResource{Group: "notificationhub.azurerm.kubeform.com", Version: "v1alpha1", Resource: "authorizationrules"}

var authorizationrulesKind = schema.GroupVersionKind{Group: "notificationhub.azurerm.kubeform.com", Version: "v1alpha1", Kind: "AuthorizationRule"}

// Get takes name of the authorizationRule, and returns the corresponding authorizationRule object, and an error if there is any.
func (c *FakeAuthorizationRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AuthorizationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(authorizationrulesResource, c.ns, name), &v1alpha1.AuthorizationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuthorizationRule), err
}

// List takes label and field selectors, and returns the list of AuthorizationRules that match those selectors.
func (c *FakeAuthorizationRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AuthorizationRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(authorizationrulesResource, authorizationrulesKind, c.ns, opts), &v1alpha1.AuthorizationRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AuthorizationRuleList{ListMeta: obj.(*v1alpha1.AuthorizationRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.AuthorizationRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested authorizationRules.
func (c *FakeAuthorizationRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(authorizationrulesResource, c.ns, opts))

}

// Create takes the representation of a authorizationRule and creates it.  Returns the server's representation of the authorizationRule, and an error, if there is any.
func (c *FakeAuthorizationRules) Create(ctx context.Context, authorizationRule *v1alpha1.AuthorizationRule, opts v1.CreateOptions) (result *v1alpha1.AuthorizationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(authorizationrulesResource, c.ns, authorizationRule), &v1alpha1.AuthorizationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuthorizationRule), err
}

// Update takes the representation of a authorizationRule and updates it. Returns the server's representation of the authorizationRule, and an error, if there is any.
func (c *FakeAuthorizationRules) Update(ctx context.Context, authorizationRule *v1alpha1.AuthorizationRule, opts v1.UpdateOptions) (result *v1alpha1.AuthorizationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(authorizationrulesResource, c.ns, authorizationRule), &v1alpha1.AuthorizationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuthorizationRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAuthorizationRules) UpdateStatus(ctx context.Context, authorizationRule *v1alpha1.AuthorizationRule, opts v1.UpdateOptions) (*v1alpha1.AuthorizationRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(authorizationrulesResource, "status", c.ns, authorizationRule), &v1alpha1.AuthorizationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuthorizationRule), err
}

// Delete takes name of the authorizationRule and deletes it. Returns an error if one occurs.
func (c *FakeAuthorizationRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(authorizationrulesResource, c.ns, name), &v1alpha1.AuthorizationRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAuthorizationRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(authorizationrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AuthorizationRuleList{})
	return err
}

// Patch applies the patch and returns the patched authorizationRule.
func (c *FakeAuthorizationRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AuthorizationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(authorizationrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AuthorizationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuthorizationRule), err
}
