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

// FakeIdentityProviderTwitters implements IdentityProviderTwitterInterface
type FakeIdentityProviderTwitters struct {
	Fake *FakeApimanagementV1alpha1
	ns   string
}

var identityprovidertwittersResource = schema.GroupVersionResource{Group: "apimanagement.azurerm.kubeform.com", Version: "v1alpha1", Resource: "identityprovidertwitters"}

var identityprovidertwittersKind = schema.GroupVersionKind{Group: "apimanagement.azurerm.kubeform.com", Version: "v1alpha1", Kind: "IdentityProviderTwitter"}

// Get takes name of the identityProviderTwitter, and returns the corresponding identityProviderTwitter object, and an error if there is any.
func (c *FakeIdentityProviderTwitters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IdentityProviderTwitter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(identityprovidertwittersResource, c.ns, name), &v1alpha1.IdentityProviderTwitter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityProviderTwitter), err
}

// List takes label and field selectors, and returns the list of IdentityProviderTwitters that match those selectors.
func (c *FakeIdentityProviderTwitters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IdentityProviderTwitterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(identityprovidertwittersResource, identityprovidertwittersKind, c.ns, opts), &v1alpha1.IdentityProviderTwitterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IdentityProviderTwitterList{ListMeta: obj.(*v1alpha1.IdentityProviderTwitterList).ListMeta}
	for _, item := range obj.(*v1alpha1.IdentityProviderTwitterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested identityProviderTwitters.
func (c *FakeIdentityProviderTwitters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(identityprovidertwittersResource, c.ns, opts))

}

// Create takes the representation of a identityProviderTwitter and creates it.  Returns the server's representation of the identityProviderTwitter, and an error, if there is any.
func (c *FakeIdentityProviderTwitters) Create(ctx context.Context, identityProviderTwitter *v1alpha1.IdentityProviderTwitter, opts v1.CreateOptions) (result *v1alpha1.IdentityProviderTwitter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(identityprovidertwittersResource, c.ns, identityProviderTwitter), &v1alpha1.IdentityProviderTwitter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityProviderTwitter), err
}

// Update takes the representation of a identityProviderTwitter and updates it. Returns the server's representation of the identityProviderTwitter, and an error, if there is any.
func (c *FakeIdentityProviderTwitters) Update(ctx context.Context, identityProviderTwitter *v1alpha1.IdentityProviderTwitter, opts v1.UpdateOptions) (result *v1alpha1.IdentityProviderTwitter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(identityprovidertwittersResource, c.ns, identityProviderTwitter), &v1alpha1.IdentityProviderTwitter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityProviderTwitter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIdentityProviderTwitters) UpdateStatus(ctx context.Context, identityProviderTwitter *v1alpha1.IdentityProviderTwitter, opts v1.UpdateOptions) (*v1alpha1.IdentityProviderTwitter, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(identityprovidertwittersResource, "status", c.ns, identityProviderTwitter), &v1alpha1.IdentityProviderTwitter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityProviderTwitter), err
}

// Delete takes name of the identityProviderTwitter and deletes it. Returns an error if one occurs.
func (c *FakeIdentityProviderTwitters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(identityprovidertwittersResource, c.ns, name), &v1alpha1.IdentityProviderTwitter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIdentityProviderTwitters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(identityprovidertwittersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IdentityProviderTwitterList{})
	return err
}

// Patch applies the patch and returns the patched identityProviderTwitter.
func (c *FakeIdentityProviderTwitters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IdentityProviderTwitter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(identityprovidertwittersResource, c.ns, name, pt, data, subresources...), &v1alpha1.IdentityProviderTwitter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityProviderTwitter), err
}
