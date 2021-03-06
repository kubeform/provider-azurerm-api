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

// FakeProperties implements PropertyInterface
type FakeProperties struct {
	Fake *FakeApimanagementV1alpha1
	ns   string
}

var propertiesResource = schema.GroupVersionResource{Group: "apimanagement.azurerm.kubeform.com", Version: "v1alpha1", Resource: "properties"}

var propertiesKind = schema.GroupVersionKind{Group: "apimanagement.azurerm.kubeform.com", Version: "v1alpha1", Kind: "Property"}

// Get takes name of the property, and returns the corresponding property object, and an error if there is any.
func (c *FakeProperties) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Property, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(propertiesResource, c.ns, name), &v1alpha1.Property{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Property), err
}

// List takes label and field selectors, and returns the list of Properties that match those selectors.
func (c *FakeProperties) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PropertyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(propertiesResource, propertiesKind, c.ns, opts), &v1alpha1.PropertyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PropertyList{ListMeta: obj.(*v1alpha1.PropertyList).ListMeta}
	for _, item := range obj.(*v1alpha1.PropertyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested properties.
func (c *FakeProperties) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(propertiesResource, c.ns, opts))

}

// Create takes the representation of a property and creates it.  Returns the server's representation of the property, and an error, if there is any.
func (c *FakeProperties) Create(ctx context.Context, property *v1alpha1.Property, opts v1.CreateOptions) (result *v1alpha1.Property, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(propertiesResource, c.ns, property), &v1alpha1.Property{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Property), err
}

// Update takes the representation of a property and updates it. Returns the server's representation of the property, and an error, if there is any.
func (c *FakeProperties) Update(ctx context.Context, property *v1alpha1.Property, opts v1.UpdateOptions) (result *v1alpha1.Property, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(propertiesResource, c.ns, property), &v1alpha1.Property{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Property), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProperties) UpdateStatus(ctx context.Context, property *v1alpha1.Property, opts v1.UpdateOptions) (*v1alpha1.Property, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(propertiesResource, "status", c.ns, property), &v1alpha1.Property{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Property), err
}

// Delete takes name of the property and deletes it. Returns an error if one occurs.
func (c *FakeProperties) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(propertiesResource, c.ns, name), &v1alpha1.Property{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProperties) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(propertiesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PropertyList{})
	return err
}

// Patch applies the patch and returns the patched property.
func (c *FakeProperties) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Property, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(propertiesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Property{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Property), err
}
