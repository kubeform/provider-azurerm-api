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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/trafficmanager/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeExternalEndpoints implements ExternalEndpointInterface
type FakeExternalEndpoints struct {
	Fake *FakeTrafficmanagerV1alpha1
	ns   string
}

var externalendpointsResource = schema.GroupVersionResource{Group: "trafficmanager.azurerm.kubeform.com", Version: "v1alpha1", Resource: "externalendpoints"}

var externalendpointsKind = schema.GroupVersionKind{Group: "trafficmanager.azurerm.kubeform.com", Version: "v1alpha1", Kind: "ExternalEndpoint"}

// Get takes name of the externalEndpoint, and returns the corresponding externalEndpoint object, and an error if there is any.
func (c *FakeExternalEndpoints) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ExternalEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(externalendpointsResource, c.ns, name), &v1alpha1.ExternalEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ExternalEndpoint), err
}

// List takes label and field selectors, and returns the list of ExternalEndpoints that match those selectors.
func (c *FakeExternalEndpoints) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ExternalEndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(externalendpointsResource, externalendpointsKind, c.ns, opts), &v1alpha1.ExternalEndpointList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ExternalEndpointList{ListMeta: obj.(*v1alpha1.ExternalEndpointList).ListMeta}
	for _, item := range obj.(*v1alpha1.ExternalEndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested externalEndpoints.
func (c *FakeExternalEndpoints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(externalendpointsResource, c.ns, opts))

}

// Create takes the representation of a externalEndpoint and creates it.  Returns the server's representation of the externalEndpoint, and an error, if there is any.
func (c *FakeExternalEndpoints) Create(ctx context.Context, externalEndpoint *v1alpha1.ExternalEndpoint, opts v1.CreateOptions) (result *v1alpha1.ExternalEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(externalendpointsResource, c.ns, externalEndpoint), &v1alpha1.ExternalEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ExternalEndpoint), err
}

// Update takes the representation of a externalEndpoint and updates it. Returns the server's representation of the externalEndpoint, and an error, if there is any.
func (c *FakeExternalEndpoints) Update(ctx context.Context, externalEndpoint *v1alpha1.ExternalEndpoint, opts v1.UpdateOptions) (result *v1alpha1.ExternalEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(externalendpointsResource, c.ns, externalEndpoint), &v1alpha1.ExternalEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ExternalEndpoint), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeExternalEndpoints) UpdateStatus(ctx context.Context, externalEndpoint *v1alpha1.ExternalEndpoint, opts v1.UpdateOptions) (*v1alpha1.ExternalEndpoint, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(externalendpointsResource, "status", c.ns, externalEndpoint), &v1alpha1.ExternalEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ExternalEndpoint), err
}

// Delete takes name of the externalEndpoint and deletes it. Returns an error if one occurs.
func (c *FakeExternalEndpoints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(externalendpointsResource, c.ns, name), &v1alpha1.ExternalEndpoint{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeExternalEndpoints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(externalendpointsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ExternalEndpointList{})
	return err
}

// Patch applies the patch and returns the patched externalEndpoint.
func (c *FakeExternalEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ExternalEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(externalendpointsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ExternalEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ExternalEndpoint), err
}
