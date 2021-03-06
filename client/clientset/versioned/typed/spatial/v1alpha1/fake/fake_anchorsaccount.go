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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/spatial/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAnchorsAccounts implements AnchorsAccountInterface
type FakeAnchorsAccounts struct {
	Fake *FakeSpatialV1alpha1
	ns   string
}

var anchorsaccountsResource = schema.GroupVersionResource{Group: "spatial.azurerm.kubeform.com", Version: "v1alpha1", Resource: "anchorsaccounts"}

var anchorsaccountsKind = schema.GroupVersionKind{Group: "spatial.azurerm.kubeform.com", Version: "v1alpha1", Kind: "AnchorsAccount"}

// Get takes name of the anchorsAccount, and returns the corresponding anchorsAccount object, and an error if there is any.
func (c *FakeAnchorsAccounts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AnchorsAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(anchorsaccountsResource, c.ns, name), &v1alpha1.AnchorsAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnchorsAccount), err
}

// List takes label and field selectors, and returns the list of AnchorsAccounts that match those selectors.
func (c *FakeAnchorsAccounts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AnchorsAccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(anchorsaccountsResource, anchorsaccountsKind, c.ns, opts), &v1alpha1.AnchorsAccountList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AnchorsAccountList{ListMeta: obj.(*v1alpha1.AnchorsAccountList).ListMeta}
	for _, item := range obj.(*v1alpha1.AnchorsAccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested anchorsAccounts.
func (c *FakeAnchorsAccounts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(anchorsaccountsResource, c.ns, opts))

}

// Create takes the representation of a anchorsAccount and creates it.  Returns the server's representation of the anchorsAccount, and an error, if there is any.
func (c *FakeAnchorsAccounts) Create(ctx context.Context, anchorsAccount *v1alpha1.AnchorsAccount, opts v1.CreateOptions) (result *v1alpha1.AnchorsAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(anchorsaccountsResource, c.ns, anchorsAccount), &v1alpha1.AnchorsAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnchorsAccount), err
}

// Update takes the representation of a anchorsAccount and updates it. Returns the server's representation of the anchorsAccount, and an error, if there is any.
func (c *FakeAnchorsAccounts) Update(ctx context.Context, anchorsAccount *v1alpha1.AnchorsAccount, opts v1.UpdateOptions) (result *v1alpha1.AnchorsAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(anchorsaccountsResource, c.ns, anchorsAccount), &v1alpha1.AnchorsAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnchorsAccount), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAnchorsAccounts) UpdateStatus(ctx context.Context, anchorsAccount *v1alpha1.AnchorsAccount, opts v1.UpdateOptions) (*v1alpha1.AnchorsAccount, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(anchorsaccountsResource, "status", c.ns, anchorsAccount), &v1alpha1.AnchorsAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnchorsAccount), err
}

// Delete takes name of the anchorsAccount and deletes it. Returns an error if one occurs.
func (c *FakeAnchorsAccounts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(anchorsaccountsResource, c.ns, name), &v1alpha1.AnchorsAccount{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAnchorsAccounts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(anchorsaccountsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AnchorsAccountList{})
	return err
}

// Patch applies the patch and returns the patched anchorsAccount.
func (c *FakeAnchorsAccounts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AnchorsAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(anchorsaccountsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AnchorsAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnchorsAccount), err
}
