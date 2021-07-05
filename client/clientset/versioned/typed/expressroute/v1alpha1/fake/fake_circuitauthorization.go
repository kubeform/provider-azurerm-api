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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/expressroute/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCircuitAuthorizations implements CircuitAuthorizationInterface
type FakeCircuitAuthorizations struct {
	Fake *FakeExpressrouteV1alpha1
	ns   string
}

var circuitauthorizationsResource = schema.GroupVersionResource{Group: "expressroute.azurerm.kubeform.com", Version: "v1alpha1", Resource: "circuitauthorizations"}

var circuitauthorizationsKind = schema.GroupVersionKind{Group: "expressroute.azurerm.kubeform.com", Version: "v1alpha1", Kind: "CircuitAuthorization"}

// Get takes name of the circuitAuthorization, and returns the corresponding circuitAuthorization object, and an error if there is any.
func (c *FakeCircuitAuthorizations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CircuitAuthorization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(circuitauthorizationsResource, c.ns, name), &v1alpha1.CircuitAuthorization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CircuitAuthorization), err
}

// List takes label and field selectors, and returns the list of CircuitAuthorizations that match those selectors.
func (c *FakeCircuitAuthorizations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CircuitAuthorizationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(circuitauthorizationsResource, circuitauthorizationsKind, c.ns, opts), &v1alpha1.CircuitAuthorizationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CircuitAuthorizationList{ListMeta: obj.(*v1alpha1.CircuitAuthorizationList).ListMeta}
	for _, item := range obj.(*v1alpha1.CircuitAuthorizationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested circuitAuthorizations.
func (c *FakeCircuitAuthorizations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(circuitauthorizationsResource, c.ns, opts))

}

// Create takes the representation of a circuitAuthorization and creates it.  Returns the server's representation of the circuitAuthorization, and an error, if there is any.
func (c *FakeCircuitAuthorizations) Create(ctx context.Context, circuitAuthorization *v1alpha1.CircuitAuthorization, opts v1.CreateOptions) (result *v1alpha1.CircuitAuthorization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(circuitauthorizationsResource, c.ns, circuitAuthorization), &v1alpha1.CircuitAuthorization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CircuitAuthorization), err
}

// Update takes the representation of a circuitAuthorization and updates it. Returns the server's representation of the circuitAuthorization, and an error, if there is any.
func (c *FakeCircuitAuthorizations) Update(ctx context.Context, circuitAuthorization *v1alpha1.CircuitAuthorization, opts v1.UpdateOptions) (result *v1alpha1.CircuitAuthorization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(circuitauthorizationsResource, c.ns, circuitAuthorization), &v1alpha1.CircuitAuthorization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CircuitAuthorization), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCircuitAuthorizations) UpdateStatus(ctx context.Context, circuitAuthorization *v1alpha1.CircuitAuthorization, opts v1.UpdateOptions) (*v1alpha1.CircuitAuthorization, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(circuitauthorizationsResource, "status", c.ns, circuitAuthorization), &v1alpha1.CircuitAuthorization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CircuitAuthorization), err
}

// Delete takes name of the circuitAuthorization and deletes it. Returns an error if one occurs.
func (c *FakeCircuitAuthorizations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(circuitauthorizationsResource, c.ns, name), &v1alpha1.CircuitAuthorization{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCircuitAuthorizations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(circuitauthorizationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CircuitAuthorizationList{})
	return err
}

// Patch applies the patch and returns the patched circuitAuthorization.
func (c *FakeCircuitAuthorizations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CircuitAuthorization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(circuitauthorizationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CircuitAuthorization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CircuitAuthorization), err
}