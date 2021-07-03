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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/container/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRegistryTokens implements RegistryTokenInterface
type FakeRegistryTokens struct {
	Fake *FakeContainerV1alpha1
	ns   string
}

var registrytokensResource = schema.GroupVersionResource{Group: "container.azurerm.kubeform.com", Version: "v1alpha1", Resource: "registrytokens"}

var registrytokensKind = schema.GroupVersionKind{Group: "container.azurerm.kubeform.com", Version: "v1alpha1", Kind: "RegistryToken"}

// Get takes name of the registryToken, and returns the corresponding registryToken object, and an error if there is any.
func (c *FakeRegistryTokens) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RegistryToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(registrytokensResource, c.ns, name), &v1alpha1.RegistryToken{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegistryToken), err
}

// List takes label and field selectors, and returns the list of RegistryTokens that match those selectors.
func (c *FakeRegistryTokens) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RegistryTokenList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(registrytokensResource, registrytokensKind, c.ns, opts), &v1alpha1.RegistryTokenList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RegistryTokenList{ListMeta: obj.(*v1alpha1.RegistryTokenList).ListMeta}
	for _, item := range obj.(*v1alpha1.RegistryTokenList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested registryTokens.
func (c *FakeRegistryTokens) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(registrytokensResource, c.ns, opts))

}

// Create takes the representation of a registryToken and creates it.  Returns the server's representation of the registryToken, and an error, if there is any.
func (c *FakeRegistryTokens) Create(ctx context.Context, registryToken *v1alpha1.RegistryToken, opts v1.CreateOptions) (result *v1alpha1.RegistryToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(registrytokensResource, c.ns, registryToken), &v1alpha1.RegistryToken{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegistryToken), err
}

// Update takes the representation of a registryToken and updates it. Returns the server's representation of the registryToken, and an error, if there is any.
func (c *FakeRegistryTokens) Update(ctx context.Context, registryToken *v1alpha1.RegistryToken, opts v1.UpdateOptions) (result *v1alpha1.RegistryToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(registrytokensResource, c.ns, registryToken), &v1alpha1.RegistryToken{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegistryToken), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRegistryTokens) UpdateStatus(ctx context.Context, registryToken *v1alpha1.RegistryToken, opts v1.UpdateOptions) (*v1alpha1.RegistryToken, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(registrytokensResource, "status", c.ns, registryToken), &v1alpha1.RegistryToken{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegistryToken), err
}

// Delete takes name of the registryToken and deletes it. Returns an error if one occurs.
func (c *FakeRegistryTokens) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(registrytokensResource, c.ns, name), &v1alpha1.RegistryToken{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRegistryTokens) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(registrytokensResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RegistryTokenList{})
	return err
}

// Patch applies the patch and returns the patched registryToken.
func (c *FakeRegistryTokens) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RegistryToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(registrytokensResource, c.ns, name, pt, data, subresources...), &v1alpha1.RegistryToken{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegistryToken), err
}
