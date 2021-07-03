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

// FakeNamedValues implements NamedValueInterface
type FakeNamedValues struct {
	Fake *FakeApimanagementV1alpha1
	ns   string
}

var namedvaluesResource = schema.GroupVersionResource{Group: "apimanagement.azurerm.kubeform.com", Version: "v1alpha1", Resource: "namedvalues"}

var namedvaluesKind = schema.GroupVersionKind{Group: "apimanagement.azurerm.kubeform.com", Version: "v1alpha1", Kind: "NamedValue"}

// Get takes name of the namedValue, and returns the corresponding namedValue object, and an error if there is any.
func (c *FakeNamedValues) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NamedValue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(namedvaluesResource, c.ns, name), &v1alpha1.NamedValue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NamedValue), err
}

// List takes label and field selectors, and returns the list of NamedValues that match those selectors.
func (c *FakeNamedValues) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NamedValueList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(namedvaluesResource, namedvaluesKind, c.ns, opts), &v1alpha1.NamedValueList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NamedValueList{ListMeta: obj.(*v1alpha1.NamedValueList).ListMeta}
	for _, item := range obj.(*v1alpha1.NamedValueList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested namedValues.
func (c *FakeNamedValues) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(namedvaluesResource, c.ns, opts))

}

// Create takes the representation of a namedValue and creates it.  Returns the server's representation of the namedValue, and an error, if there is any.
func (c *FakeNamedValues) Create(ctx context.Context, namedValue *v1alpha1.NamedValue, opts v1.CreateOptions) (result *v1alpha1.NamedValue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(namedvaluesResource, c.ns, namedValue), &v1alpha1.NamedValue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NamedValue), err
}

// Update takes the representation of a namedValue and updates it. Returns the server's representation of the namedValue, and an error, if there is any.
func (c *FakeNamedValues) Update(ctx context.Context, namedValue *v1alpha1.NamedValue, opts v1.UpdateOptions) (result *v1alpha1.NamedValue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(namedvaluesResource, c.ns, namedValue), &v1alpha1.NamedValue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NamedValue), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNamedValues) UpdateStatus(ctx context.Context, namedValue *v1alpha1.NamedValue, opts v1.UpdateOptions) (*v1alpha1.NamedValue, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(namedvaluesResource, "status", c.ns, namedValue), &v1alpha1.NamedValue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NamedValue), err
}

// Delete takes name of the namedValue and deletes it. Returns an error if one occurs.
func (c *FakeNamedValues) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(namedvaluesResource, c.ns, name), &v1alpha1.NamedValue{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNamedValues) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(namedvaluesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NamedValueList{})
	return err
}

// Patch applies the patch and returns the patched namedValue.
func (c *FakeNamedValues) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NamedValue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(namedvaluesResource, c.ns, name, pt, data, subresources...), &v1alpha1.NamedValue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NamedValue), err
}
