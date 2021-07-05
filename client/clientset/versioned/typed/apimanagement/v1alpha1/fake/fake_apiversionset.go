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

// FakeApiVersionSets implements ApiVersionSetInterface
type FakeApiVersionSets struct {
	Fake *FakeApimanagementV1alpha1
	ns   string
}

var apiversionsetsResource = schema.GroupVersionResource{Group: "apimanagement.azurerm.kubeform.com", Version: "v1alpha1", Resource: "apiversionsets"}

var apiversionsetsKind = schema.GroupVersionKind{Group: "apimanagement.azurerm.kubeform.com", Version: "v1alpha1", Kind: "ApiVersionSet"}

// Get takes name of the apiVersionSet, and returns the corresponding apiVersionSet object, and an error if there is any.
func (c *FakeApiVersionSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApiVersionSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apiversionsetsResource, c.ns, name), &v1alpha1.ApiVersionSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiVersionSet), err
}

// List takes label and field selectors, and returns the list of ApiVersionSets that match those selectors.
func (c *FakeApiVersionSets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApiVersionSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apiversionsetsResource, apiversionsetsKind, c.ns, opts), &v1alpha1.ApiVersionSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiVersionSetList{ListMeta: obj.(*v1alpha1.ApiVersionSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiVersionSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiVersionSets.
func (c *FakeApiVersionSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apiversionsetsResource, c.ns, opts))

}

// Create takes the representation of a apiVersionSet and creates it.  Returns the server's representation of the apiVersionSet, and an error, if there is any.
func (c *FakeApiVersionSets) Create(ctx context.Context, apiVersionSet *v1alpha1.ApiVersionSet, opts v1.CreateOptions) (result *v1alpha1.ApiVersionSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apiversionsetsResource, c.ns, apiVersionSet), &v1alpha1.ApiVersionSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiVersionSet), err
}

// Update takes the representation of a apiVersionSet and updates it. Returns the server's representation of the apiVersionSet, and an error, if there is any.
func (c *FakeApiVersionSets) Update(ctx context.Context, apiVersionSet *v1alpha1.ApiVersionSet, opts v1.UpdateOptions) (result *v1alpha1.ApiVersionSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apiversionsetsResource, c.ns, apiVersionSet), &v1alpha1.ApiVersionSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiVersionSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiVersionSets) UpdateStatus(ctx context.Context, apiVersionSet *v1alpha1.ApiVersionSet, opts v1.UpdateOptions) (*v1alpha1.ApiVersionSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apiversionsetsResource, "status", c.ns, apiVersionSet), &v1alpha1.ApiVersionSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiVersionSet), err
}

// Delete takes name of the apiVersionSet and deletes it. Returns an error if one occurs.
func (c *FakeApiVersionSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apiversionsetsResource, c.ns, name), &v1alpha1.ApiVersionSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiVersionSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apiversionsetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiVersionSetList{})
	return err
}

// Patch applies the patch and returns the patched apiVersionSet.
func (c *FakeApiVersionSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiVersionSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apiversionsetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiVersionSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiVersionSet), err
}