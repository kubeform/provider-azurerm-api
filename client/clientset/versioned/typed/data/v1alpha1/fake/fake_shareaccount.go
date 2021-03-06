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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/data/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeShareAccounts implements ShareAccountInterface
type FakeShareAccounts struct {
	Fake *FakeDataV1alpha1
	ns   string
}

var shareaccountsResource = schema.GroupVersionResource{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Resource: "shareaccounts"}

var shareaccountsKind = schema.GroupVersionKind{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Kind: "ShareAccount"}

// Get takes name of the shareAccount, and returns the corresponding shareAccount object, and an error if there is any.
func (c *FakeShareAccounts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ShareAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(shareaccountsResource, c.ns, name), &v1alpha1.ShareAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShareAccount), err
}

// List takes label and field selectors, and returns the list of ShareAccounts that match those selectors.
func (c *FakeShareAccounts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ShareAccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(shareaccountsResource, shareaccountsKind, c.ns, opts), &v1alpha1.ShareAccountList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ShareAccountList{ListMeta: obj.(*v1alpha1.ShareAccountList).ListMeta}
	for _, item := range obj.(*v1alpha1.ShareAccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested shareAccounts.
func (c *FakeShareAccounts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(shareaccountsResource, c.ns, opts))

}

// Create takes the representation of a shareAccount and creates it.  Returns the server's representation of the shareAccount, and an error, if there is any.
func (c *FakeShareAccounts) Create(ctx context.Context, shareAccount *v1alpha1.ShareAccount, opts v1.CreateOptions) (result *v1alpha1.ShareAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(shareaccountsResource, c.ns, shareAccount), &v1alpha1.ShareAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShareAccount), err
}

// Update takes the representation of a shareAccount and updates it. Returns the server's representation of the shareAccount, and an error, if there is any.
func (c *FakeShareAccounts) Update(ctx context.Context, shareAccount *v1alpha1.ShareAccount, opts v1.UpdateOptions) (result *v1alpha1.ShareAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(shareaccountsResource, c.ns, shareAccount), &v1alpha1.ShareAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShareAccount), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeShareAccounts) UpdateStatus(ctx context.Context, shareAccount *v1alpha1.ShareAccount, opts v1.UpdateOptions) (*v1alpha1.ShareAccount, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(shareaccountsResource, "status", c.ns, shareAccount), &v1alpha1.ShareAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShareAccount), err
}

// Delete takes name of the shareAccount and deletes it. Returns an error if one occurs.
func (c *FakeShareAccounts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(shareaccountsResource, c.ns, name), &v1alpha1.ShareAccount{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeShareAccounts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(shareaccountsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ShareAccountList{})
	return err
}

// Patch applies the patch and returns the patched shareAccount.
func (c *FakeShareAccounts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ShareAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(shareaccountsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ShareAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShareAccount), err
}
