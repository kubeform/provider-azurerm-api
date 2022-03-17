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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/logicapp/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStandards implements StandardInterface
type FakeStandards struct {
	Fake *FakeLogicappV1alpha1
	ns   string
}

var standardsResource = schema.GroupVersionResource{Group: "logicapp.azurerm.kubeform.com", Version: "v1alpha1", Resource: "standards"}

var standardsKind = schema.GroupVersionKind{Group: "logicapp.azurerm.kubeform.com", Version: "v1alpha1", Kind: "Standard"}

// Get takes name of the standard, and returns the corresponding standard object, and an error if there is any.
func (c *FakeStandards) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Standard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(standardsResource, c.ns, name), &v1alpha1.Standard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Standard), err
}

// List takes label and field selectors, and returns the list of Standards that match those selectors.
func (c *FakeStandards) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.StandardList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(standardsResource, standardsKind, c.ns, opts), &v1alpha1.StandardList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StandardList{ListMeta: obj.(*v1alpha1.StandardList).ListMeta}
	for _, item := range obj.(*v1alpha1.StandardList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested standards.
func (c *FakeStandards) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(standardsResource, c.ns, opts))

}

// Create takes the representation of a standard and creates it.  Returns the server's representation of the standard, and an error, if there is any.
func (c *FakeStandards) Create(ctx context.Context, standard *v1alpha1.Standard, opts v1.CreateOptions) (result *v1alpha1.Standard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(standardsResource, c.ns, standard), &v1alpha1.Standard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Standard), err
}

// Update takes the representation of a standard and updates it. Returns the server's representation of the standard, and an error, if there is any.
func (c *FakeStandards) Update(ctx context.Context, standard *v1alpha1.Standard, opts v1.UpdateOptions) (result *v1alpha1.Standard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(standardsResource, c.ns, standard), &v1alpha1.Standard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Standard), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStandards) UpdateStatus(ctx context.Context, standard *v1alpha1.Standard, opts v1.UpdateOptions) (*v1alpha1.Standard, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(standardsResource, "status", c.ns, standard), &v1alpha1.Standard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Standard), err
}

// Delete takes name of the standard and deletes it. Returns an error if one occurs.
func (c *FakeStandards) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(standardsResource, c.ns, name), &v1alpha1.Standard{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStandards) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(standardsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.StandardList{})
	return err
}

// Patch applies the patch and returns the patched standard.
func (c *FakeStandards) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Standard, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(standardsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Standard{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Standard), err
}