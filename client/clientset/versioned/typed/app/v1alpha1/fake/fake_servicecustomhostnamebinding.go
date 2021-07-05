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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/app/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceCustomHostnameBindings implements ServiceCustomHostnameBindingInterface
type FakeServiceCustomHostnameBindings struct {
	Fake *FakeAppV1alpha1
	ns   string
}

var servicecustomhostnamebindingsResource = schema.GroupVersionResource{Group: "app.azurerm.kubeform.com", Version: "v1alpha1", Resource: "servicecustomhostnamebindings"}

var servicecustomhostnamebindingsKind = schema.GroupVersionKind{Group: "app.azurerm.kubeform.com", Version: "v1alpha1", Kind: "ServiceCustomHostnameBinding"}

// Get takes name of the serviceCustomHostnameBinding, and returns the corresponding serviceCustomHostnameBinding object, and an error if there is any.
func (c *FakeServiceCustomHostnameBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServiceCustomHostnameBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicecustomhostnamebindingsResource, c.ns, name), &v1alpha1.ServiceCustomHostnameBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceCustomHostnameBinding), err
}

// List takes label and field selectors, and returns the list of ServiceCustomHostnameBindings that match those selectors.
func (c *FakeServiceCustomHostnameBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServiceCustomHostnameBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicecustomhostnamebindingsResource, servicecustomhostnamebindingsKind, c.ns, opts), &v1alpha1.ServiceCustomHostnameBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServiceCustomHostnameBindingList{ListMeta: obj.(*v1alpha1.ServiceCustomHostnameBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServiceCustomHostnameBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceCustomHostnameBindings.
func (c *FakeServiceCustomHostnameBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicecustomhostnamebindingsResource, c.ns, opts))

}

// Create takes the representation of a serviceCustomHostnameBinding and creates it.  Returns the server's representation of the serviceCustomHostnameBinding, and an error, if there is any.
func (c *FakeServiceCustomHostnameBindings) Create(ctx context.Context, serviceCustomHostnameBinding *v1alpha1.ServiceCustomHostnameBinding, opts v1.CreateOptions) (result *v1alpha1.ServiceCustomHostnameBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicecustomhostnamebindingsResource, c.ns, serviceCustomHostnameBinding), &v1alpha1.ServiceCustomHostnameBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceCustomHostnameBinding), err
}

// Update takes the representation of a serviceCustomHostnameBinding and updates it. Returns the server's representation of the serviceCustomHostnameBinding, and an error, if there is any.
func (c *FakeServiceCustomHostnameBindings) Update(ctx context.Context, serviceCustomHostnameBinding *v1alpha1.ServiceCustomHostnameBinding, opts v1.UpdateOptions) (result *v1alpha1.ServiceCustomHostnameBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicecustomhostnamebindingsResource, c.ns, serviceCustomHostnameBinding), &v1alpha1.ServiceCustomHostnameBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceCustomHostnameBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceCustomHostnameBindings) UpdateStatus(ctx context.Context, serviceCustomHostnameBinding *v1alpha1.ServiceCustomHostnameBinding, opts v1.UpdateOptions) (*v1alpha1.ServiceCustomHostnameBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicecustomhostnamebindingsResource, "status", c.ns, serviceCustomHostnameBinding), &v1alpha1.ServiceCustomHostnameBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceCustomHostnameBinding), err
}

// Delete takes name of the serviceCustomHostnameBinding and deletes it. Returns an error if one occurs.
func (c *FakeServiceCustomHostnameBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicecustomhostnamebindingsResource, c.ns, name), &v1alpha1.ServiceCustomHostnameBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceCustomHostnameBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicecustomhostnamebindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServiceCustomHostnameBindingList{})
	return err
}

// Patch applies the patch and returns the patched serviceCustomHostnameBinding.
func (c *FakeServiceCustomHostnameBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceCustomHostnameBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicecustomhostnamebindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ServiceCustomHostnameBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceCustomHostnameBinding), err
}