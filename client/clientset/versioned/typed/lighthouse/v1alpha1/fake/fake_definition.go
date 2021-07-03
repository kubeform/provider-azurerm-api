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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/lighthouse/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDefinitions implements DefinitionInterface
type FakeDefinitions struct {
	Fake *FakeLighthouseV1alpha1
	ns   string
}

var definitionsResource = schema.GroupVersionResource{Group: "lighthouse.azurerm.kubeform.com", Version: "v1alpha1", Resource: "definitions"}

var definitionsKind = schema.GroupVersionKind{Group: "lighthouse.azurerm.kubeform.com", Version: "v1alpha1", Kind: "Definition"}

// Get takes name of the definition, and returns the corresponding definition object, and an error if there is any.
func (c *FakeDefinitions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Definition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(definitionsResource, c.ns, name), &v1alpha1.Definition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Definition), err
}

// List takes label and field selectors, and returns the list of Definitions that match those selectors.
func (c *FakeDefinitions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DefinitionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(definitionsResource, definitionsKind, c.ns, opts), &v1alpha1.DefinitionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DefinitionList{ListMeta: obj.(*v1alpha1.DefinitionList).ListMeta}
	for _, item := range obj.(*v1alpha1.DefinitionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested definitions.
func (c *FakeDefinitions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(definitionsResource, c.ns, opts))

}

// Create takes the representation of a definition and creates it.  Returns the server's representation of the definition, and an error, if there is any.
func (c *FakeDefinitions) Create(ctx context.Context, definition *v1alpha1.Definition, opts v1.CreateOptions) (result *v1alpha1.Definition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(definitionsResource, c.ns, definition), &v1alpha1.Definition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Definition), err
}

// Update takes the representation of a definition and updates it. Returns the server's representation of the definition, and an error, if there is any.
func (c *FakeDefinitions) Update(ctx context.Context, definition *v1alpha1.Definition, opts v1.UpdateOptions) (result *v1alpha1.Definition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(definitionsResource, c.ns, definition), &v1alpha1.Definition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Definition), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDefinitions) UpdateStatus(ctx context.Context, definition *v1alpha1.Definition, opts v1.UpdateOptions) (*v1alpha1.Definition, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(definitionsResource, "status", c.ns, definition), &v1alpha1.Definition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Definition), err
}

// Delete takes name of the definition and deletes it. Returns an error if one occurs.
func (c *FakeDefinitions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(definitionsResource, c.ns, name), &v1alpha1.Definition{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDefinitions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(definitionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DefinitionList{})
	return err
}

// Patch applies the patch and returns the patched definition.
func (c *FakeDefinitions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Definition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(definitionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Definition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Definition), err
}
