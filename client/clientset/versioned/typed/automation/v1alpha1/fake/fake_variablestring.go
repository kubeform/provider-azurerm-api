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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/automation/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVariableStrings implements VariableStringInterface
type FakeVariableStrings struct {
	Fake *FakeAutomationV1alpha1
	ns   string
}

var variablestringsResource = schema.GroupVersionResource{Group: "automation.azurerm.kubeform.com", Version: "v1alpha1", Resource: "variablestrings"}

var variablestringsKind = schema.GroupVersionKind{Group: "automation.azurerm.kubeform.com", Version: "v1alpha1", Kind: "VariableString"}

// Get takes name of the variableString, and returns the corresponding variableString object, and an error if there is any.
func (c *FakeVariableStrings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VariableString, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(variablestringsResource, c.ns, name), &v1alpha1.VariableString{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VariableString), err
}

// List takes label and field selectors, and returns the list of VariableStrings that match those selectors.
func (c *FakeVariableStrings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VariableStringList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(variablestringsResource, variablestringsKind, c.ns, opts), &v1alpha1.VariableStringList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VariableStringList{ListMeta: obj.(*v1alpha1.VariableStringList).ListMeta}
	for _, item := range obj.(*v1alpha1.VariableStringList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested variableStrings.
func (c *FakeVariableStrings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(variablestringsResource, c.ns, opts))

}

// Create takes the representation of a variableString and creates it.  Returns the server's representation of the variableString, and an error, if there is any.
func (c *FakeVariableStrings) Create(ctx context.Context, variableString *v1alpha1.VariableString, opts v1.CreateOptions) (result *v1alpha1.VariableString, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(variablestringsResource, c.ns, variableString), &v1alpha1.VariableString{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VariableString), err
}

// Update takes the representation of a variableString and updates it. Returns the server's representation of the variableString, and an error, if there is any.
func (c *FakeVariableStrings) Update(ctx context.Context, variableString *v1alpha1.VariableString, opts v1.UpdateOptions) (result *v1alpha1.VariableString, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(variablestringsResource, c.ns, variableString), &v1alpha1.VariableString{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VariableString), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVariableStrings) UpdateStatus(ctx context.Context, variableString *v1alpha1.VariableString, opts v1.UpdateOptions) (*v1alpha1.VariableString, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(variablestringsResource, "status", c.ns, variableString), &v1alpha1.VariableString{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VariableString), err
}

// Delete takes name of the variableString and deletes it. Returns an error if one occurs.
func (c *FakeVariableStrings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(variablestringsResource, c.ns, name), &v1alpha1.VariableString{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVariableStrings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(variablestringsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VariableStringList{})
	return err
}

// Patch applies the patch and returns the patched variableString.
func (c *FakeVariableStrings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VariableString, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(variablestringsResource, c.ns, name, pt, data, subresources...), &v1alpha1.VariableString{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VariableString), err
}