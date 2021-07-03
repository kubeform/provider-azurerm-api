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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/media/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServicesAccounts implements ServicesAccountInterface
type FakeServicesAccounts struct {
	Fake *FakeMediaV1alpha1
	ns   string
}

var servicesaccountsResource = schema.GroupVersionResource{Group: "media.azurerm.kubeform.com", Version: "v1alpha1", Resource: "servicesaccounts"}

var servicesaccountsKind = schema.GroupVersionKind{Group: "media.azurerm.kubeform.com", Version: "v1alpha1", Kind: "ServicesAccount"}

// Get takes name of the servicesAccount, and returns the corresponding servicesAccount object, and an error if there is any.
func (c *FakeServicesAccounts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServicesAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicesaccountsResource, c.ns, name), &v1alpha1.ServicesAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicesAccount), err
}

// List takes label and field selectors, and returns the list of ServicesAccounts that match those selectors.
func (c *FakeServicesAccounts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServicesAccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicesaccountsResource, servicesaccountsKind, c.ns, opts), &v1alpha1.ServicesAccountList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServicesAccountList{ListMeta: obj.(*v1alpha1.ServicesAccountList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServicesAccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested servicesAccounts.
func (c *FakeServicesAccounts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicesaccountsResource, c.ns, opts))

}

// Create takes the representation of a servicesAccount and creates it.  Returns the server's representation of the servicesAccount, and an error, if there is any.
func (c *FakeServicesAccounts) Create(ctx context.Context, servicesAccount *v1alpha1.ServicesAccount, opts v1.CreateOptions) (result *v1alpha1.ServicesAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicesaccountsResource, c.ns, servicesAccount), &v1alpha1.ServicesAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicesAccount), err
}

// Update takes the representation of a servicesAccount and updates it. Returns the server's representation of the servicesAccount, and an error, if there is any.
func (c *FakeServicesAccounts) Update(ctx context.Context, servicesAccount *v1alpha1.ServicesAccount, opts v1.UpdateOptions) (result *v1alpha1.ServicesAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicesaccountsResource, c.ns, servicesAccount), &v1alpha1.ServicesAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicesAccount), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServicesAccounts) UpdateStatus(ctx context.Context, servicesAccount *v1alpha1.ServicesAccount, opts v1.UpdateOptions) (*v1alpha1.ServicesAccount, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicesaccountsResource, "status", c.ns, servicesAccount), &v1alpha1.ServicesAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicesAccount), err
}

// Delete takes name of the servicesAccount and deletes it. Returns an error if one occurs.
func (c *FakeServicesAccounts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicesaccountsResource, c.ns, name), &v1alpha1.ServicesAccount{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServicesAccounts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicesaccountsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServicesAccountList{})
	return err
}

// Patch applies the patch and returns the patched servicesAccount.
func (c *FakeServicesAccounts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServicesAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicesaccountsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ServicesAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicesAccount), err
}
