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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/servicebus/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSubscriptionRules implements SubscriptionRuleInterface
type FakeSubscriptionRules struct {
	Fake *FakeServicebusV1alpha1
	ns   string
}

var subscriptionrulesResource = schema.GroupVersionResource{Group: "servicebus.azurerm.kubeform.com", Version: "v1alpha1", Resource: "subscriptionrules"}

var subscriptionrulesKind = schema.GroupVersionKind{Group: "servicebus.azurerm.kubeform.com", Version: "v1alpha1", Kind: "SubscriptionRule"}

// Get takes name of the subscriptionRule, and returns the corresponding subscriptionRule object, and an error if there is any.
func (c *FakeSubscriptionRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SubscriptionRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(subscriptionrulesResource, c.ns, name), &v1alpha1.SubscriptionRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubscriptionRule), err
}

// List takes label and field selectors, and returns the list of SubscriptionRules that match those selectors.
func (c *FakeSubscriptionRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SubscriptionRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(subscriptionrulesResource, subscriptionrulesKind, c.ns, opts), &v1alpha1.SubscriptionRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SubscriptionRuleList{ListMeta: obj.(*v1alpha1.SubscriptionRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.SubscriptionRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested subscriptionRules.
func (c *FakeSubscriptionRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(subscriptionrulesResource, c.ns, opts))

}

// Create takes the representation of a subscriptionRule and creates it.  Returns the server's representation of the subscriptionRule, and an error, if there is any.
func (c *FakeSubscriptionRules) Create(ctx context.Context, subscriptionRule *v1alpha1.SubscriptionRule, opts v1.CreateOptions) (result *v1alpha1.SubscriptionRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(subscriptionrulesResource, c.ns, subscriptionRule), &v1alpha1.SubscriptionRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubscriptionRule), err
}

// Update takes the representation of a subscriptionRule and updates it. Returns the server's representation of the subscriptionRule, and an error, if there is any.
func (c *FakeSubscriptionRules) Update(ctx context.Context, subscriptionRule *v1alpha1.SubscriptionRule, opts v1.UpdateOptions) (result *v1alpha1.SubscriptionRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(subscriptionrulesResource, c.ns, subscriptionRule), &v1alpha1.SubscriptionRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubscriptionRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSubscriptionRules) UpdateStatus(ctx context.Context, subscriptionRule *v1alpha1.SubscriptionRule, opts v1.UpdateOptions) (*v1alpha1.SubscriptionRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(subscriptionrulesResource, "status", c.ns, subscriptionRule), &v1alpha1.SubscriptionRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubscriptionRule), err
}

// Delete takes name of the subscriptionRule and deletes it. Returns an error if one occurs.
func (c *FakeSubscriptionRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(subscriptionrulesResource, c.ns, name), &v1alpha1.SubscriptionRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSubscriptionRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(subscriptionrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SubscriptionRuleList{})
	return err
}

// Patch applies the patch and returns the patched subscriptionRule.
func (c *FakeSubscriptionRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SubscriptionRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(subscriptionrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SubscriptionRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubscriptionRule), err
}