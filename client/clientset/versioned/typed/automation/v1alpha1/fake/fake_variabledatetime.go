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

// FakeVariableDatetimes implements VariableDatetimeInterface
type FakeVariableDatetimes struct {
	Fake *FakeAutomationV1alpha1
	ns   string
}

var variabledatetimesResource = schema.GroupVersionResource{Group: "automation.azurerm.kubeform.com", Version: "v1alpha1", Resource: "variabledatetimes"}

var variabledatetimesKind = schema.GroupVersionKind{Group: "automation.azurerm.kubeform.com", Version: "v1alpha1", Kind: "VariableDatetime"}

// Get takes name of the variableDatetime, and returns the corresponding variableDatetime object, and an error if there is any.
func (c *FakeVariableDatetimes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VariableDatetime, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(variabledatetimesResource, c.ns, name), &v1alpha1.VariableDatetime{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VariableDatetime), err
}

// List takes label and field selectors, and returns the list of VariableDatetimes that match those selectors.
func (c *FakeVariableDatetimes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VariableDatetimeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(variabledatetimesResource, variabledatetimesKind, c.ns, opts), &v1alpha1.VariableDatetimeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VariableDatetimeList{ListMeta: obj.(*v1alpha1.VariableDatetimeList).ListMeta}
	for _, item := range obj.(*v1alpha1.VariableDatetimeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested variableDatetimes.
func (c *FakeVariableDatetimes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(variabledatetimesResource, c.ns, opts))

}

// Create takes the representation of a variableDatetime and creates it.  Returns the server's representation of the variableDatetime, and an error, if there is any.
func (c *FakeVariableDatetimes) Create(ctx context.Context, variableDatetime *v1alpha1.VariableDatetime, opts v1.CreateOptions) (result *v1alpha1.VariableDatetime, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(variabledatetimesResource, c.ns, variableDatetime), &v1alpha1.VariableDatetime{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VariableDatetime), err
}

// Update takes the representation of a variableDatetime and updates it. Returns the server's representation of the variableDatetime, and an error, if there is any.
func (c *FakeVariableDatetimes) Update(ctx context.Context, variableDatetime *v1alpha1.VariableDatetime, opts v1.UpdateOptions) (result *v1alpha1.VariableDatetime, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(variabledatetimesResource, c.ns, variableDatetime), &v1alpha1.VariableDatetime{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VariableDatetime), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVariableDatetimes) UpdateStatus(ctx context.Context, variableDatetime *v1alpha1.VariableDatetime, opts v1.UpdateOptions) (*v1alpha1.VariableDatetime, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(variabledatetimesResource, "status", c.ns, variableDatetime), &v1alpha1.VariableDatetime{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VariableDatetime), err
}

// Delete takes name of the variableDatetime and deletes it. Returns an error if one occurs.
func (c *FakeVariableDatetimes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(variabledatetimesResource, c.ns, name), &v1alpha1.VariableDatetime{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVariableDatetimes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(variabledatetimesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VariableDatetimeList{})
	return err
}

// Patch applies the patch and returns the patched variableDatetime.
func (c *FakeVariableDatetimes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VariableDatetime, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(variabledatetimesResource, c.ns, name, pt, data, subresources...), &v1alpha1.VariableDatetime{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VariableDatetime), err
}