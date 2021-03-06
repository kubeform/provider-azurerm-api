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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/apimanagement/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIdentityProviderAadb2cs implements IdentityProviderAadb2cInterface
type FakeIdentityProviderAadb2cs struct {
	Fake *FakeApimanagementV1alpha1
	ns   string
}

var identityprovideraadb2csResource = schema.GroupVersionResource{Group: "apimanagement.azurerm.kubeform.com", Version: "v1alpha1", Resource: "identityprovideraadb2cs"}

var identityprovideraadb2csKind = schema.GroupVersionKind{Group: "apimanagement.azurerm.kubeform.com", Version: "v1alpha1", Kind: "IdentityProviderAadb2c"}

// Get takes name of the identityProviderAadb2c, and returns the corresponding identityProviderAadb2c object, and an error if there is any.
func (c *FakeIdentityProviderAadb2cs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IdentityProviderAadb2c, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(identityprovideraadb2csResource, c.ns, name), &v1alpha1.IdentityProviderAadb2c{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityProviderAadb2c), err
}

// List takes label and field selectors, and returns the list of IdentityProviderAadb2cs that match those selectors.
func (c *FakeIdentityProviderAadb2cs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IdentityProviderAadb2cList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(identityprovideraadb2csResource, identityprovideraadb2csKind, c.ns, opts), &v1alpha1.IdentityProviderAadb2cList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IdentityProviderAadb2cList{ListMeta: obj.(*v1alpha1.IdentityProviderAadb2cList).ListMeta}
	for _, item := range obj.(*v1alpha1.IdentityProviderAadb2cList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested identityProviderAadb2cs.
func (c *FakeIdentityProviderAadb2cs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(identityprovideraadb2csResource, c.ns, opts))

}

// Create takes the representation of a identityProviderAadb2c and creates it.  Returns the server's representation of the identityProviderAadb2c, and an error, if there is any.
func (c *FakeIdentityProviderAadb2cs) Create(ctx context.Context, identityProviderAadb2c *v1alpha1.IdentityProviderAadb2c, opts v1.CreateOptions) (result *v1alpha1.IdentityProviderAadb2c, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(identityprovideraadb2csResource, c.ns, identityProviderAadb2c), &v1alpha1.IdentityProviderAadb2c{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityProviderAadb2c), err
}

// Update takes the representation of a identityProviderAadb2c and updates it. Returns the server's representation of the identityProviderAadb2c, and an error, if there is any.
func (c *FakeIdentityProviderAadb2cs) Update(ctx context.Context, identityProviderAadb2c *v1alpha1.IdentityProviderAadb2c, opts v1.UpdateOptions) (result *v1alpha1.IdentityProviderAadb2c, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(identityprovideraadb2csResource, c.ns, identityProviderAadb2c), &v1alpha1.IdentityProviderAadb2c{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityProviderAadb2c), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIdentityProviderAadb2cs) UpdateStatus(ctx context.Context, identityProviderAadb2c *v1alpha1.IdentityProviderAadb2c, opts v1.UpdateOptions) (*v1alpha1.IdentityProviderAadb2c, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(identityprovideraadb2csResource, "status", c.ns, identityProviderAadb2c), &v1alpha1.IdentityProviderAadb2c{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityProviderAadb2c), err
}

// Delete takes name of the identityProviderAadb2c and deletes it. Returns an error if one occurs.
func (c *FakeIdentityProviderAadb2cs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(identityprovideraadb2csResource, c.ns, name), &v1alpha1.IdentityProviderAadb2c{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIdentityProviderAadb2cs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(identityprovideraadb2csResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IdentityProviderAadb2cList{})
	return err
}

// Patch applies the patch and returns the patched identityProviderAadb2c.
func (c *FakeIdentityProviderAadb2cs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IdentityProviderAadb2c, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(identityprovideraadb2csResource, c.ns, name, pt, data, subresources...), &v1alpha1.IdentityProviderAadb2c{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityProviderAadb2c), err
}
