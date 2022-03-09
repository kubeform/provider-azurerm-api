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

// FakeManagedStorageAccounts implements ManagedStorageAccountInterface
type FakeManagedStorageAccounts struct {
	Fake *FakeKeyvaultV1alpha1
	ns   string
}

var managedstorageaccountsResource = schema.GroupVersionResource{Group: "keyvault.azurerm.kubeform.com", Version: "v1alpha1", Resource: "managedstorageaccounts"}

var managedstorageaccountsKind = schema.GroupVersionKind{Group: "keyvault.azurerm.kubeform.com", Version: "v1alpha1", Kind: "ManagedStorageAccount"}

// Get takes name of the managedStorageAccount, and returns the corresponding managedStorageAccount object, and an error if there is any.
func (c *FakeManagedStorageAccounts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ManagedStorageAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(managedstorageaccountsResource, c.ns, name), &v1alpha1.ManagedStorageAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedStorageAccount), err
}

// List takes label and field selectors, and returns the list of ManagedStorageAccounts that match those selectors.
func (c *FakeManagedStorageAccounts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ManagedStorageAccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(managedstorageaccountsResource, managedstorageaccountsKind, c.ns, opts), &v1alpha1.ManagedStorageAccountList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ManagedStorageAccountList{ListMeta: obj.(*v1alpha1.ManagedStorageAccountList).ListMeta}
	for _, item := range obj.(*v1alpha1.ManagedStorageAccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managedStorageAccounts.
func (c *FakeManagedStorageAccounts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(managedstorageaccountsResource, c.ns, opts))

}

// Create takes the representation of a managedStorageAccount and creates it.  Returns the server's representation of the managedStorageAccount, and an error, if there is any.
func (c *FakeManagedStorageAccounts) Create(ctx context.Context, managedStorageAccount *v1alpha1.ManagedStorageAccount, opts v1.CreateOptions) (result *v1alpha1.ManagedStorageAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(managedstorageaccountsResource, c.ns, managedStorageAccount), &v1alpha1.ManagedStorageAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedStorageAccount), err
}

// Update takes the representation of a managedStorageAccount and updates it. Returns the server's representation of the managedStorageAccount, and an error, if there is any.
func (c *FakeManagedStorageAccounts) Update(ctx context.Context, managedStorageAccount *v1alpha1.ManagedStorageAccount, opts v1.UpdateOptions) (result *v1alpha1.ManagedStorageAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(managedstorageaccountsResource, c.ns, managedStorageAccount), &v1alpha1.ManagedStorageAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedStorageAccount), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagedStorageAccounts) UpdateStatus(ctx context.Context, managedStorageAccount *v1alpha1.ManagedStorageAccount, opts v1.UpdateOptions) (*v1alpha1.ManagedStorageAccount, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(managedstorageaccountsResource, "status", c.ns, managedStorageAccount), &v1alpha1.ManagedStorageAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedStorageAccount), err
}

// Delete takes name of the managedStorageAccount and deletes it. Returns an error if one occurs.
func (c *FakeManagedStorageAccounts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(managedstorageaccountsResource, c.ns, name), &v1alpha1.ManagedStorageAccount{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagedStorageAccounts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(managedstorageaccountsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ManagedStorageAccountList{})
	return err
}

// Patch applies the patch and returns the patched managedStorageAccount.
func (c *FakeManagedStorageAccounts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagedStorageAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(managedstorageaccountsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ManagedStorageAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedStorageAccount), err
}
