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

// FakeAzureEndpoints implements AzureEndpointInterface
type FakeAzureEndpoints struct {
	Fake *FakeTrafficmanagerV1alpha1
	ns   string
}

var azureendpointsResource = schema.GroupVersionResource{Group: "trafficmanager.azurerm.kubeform.com", Version: "v1alpha1", Resource: "azureendpoints"}

var azureendpointsKind = schema.GroupVersionKind{Group: "trafficmanager.azurerm.kubeform.com", Version: "v1alpha1", Kind: "AzureEndpoint"}

// Get takes name of the azureEndpoint, and returns the corresponding azureEndpoint object, and an error if there is any.
func (c *FakeAzureEndpoints) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AzureEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(azureendpointsResource, c.ns, name), &v1alpha1.AzureEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureEndpoint), err
}

// List takes label and field selectors, and returns the list of AzureEndpoints that match those selectors.
func (c *FakeAzureEndpoints) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AzureEndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(azureendpointsResource, azureendpointsKind, c.ns, opts), &v1alpha1.AzureEndpointList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzureEndpointList{ListMeta: obj.(*v1alpha1.AzureEndpointList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzureEndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azureEndpoints.
func (c *FakeAzureEndpoints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(azureendpointsResource, c.ns, opts))

}

// Create takes the representation of a azureEndpoint and creates it.  Returns the server's representation of the azureEndpoint, and an error, if there is any.
func (c *FakeAzureEndpoints) Create(ctx context.Context, azureEndpoint *v1alpha1.AzureEndpoint, opts v1.CreateOptions) (result *v1alpha1.AzureEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(azureendpointsResource, c.ns, azureEndpoint), &v1alpha1.AzureEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureEndpoint), err
}

// Update takes the representation of a azureEndpoint and updates it. Returns the server's representation of the azureEndpoint, and an error, if there is any.
func (c *FakeAzureEndpoints) Update(ctx context.Context, azureEndpoint *v1alpha1.AzureEndpoint, opts v1.UpdateOptions) (result *v1alpha1.AzureEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(azureendpointsResource, c.ns, azureEndpoint), &v1alpha1.AzureEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureEndpoint), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzureEndpoints) UpdateStatus(ctx context.Context, azureEndpoint *v1alpha1.AzureEndpoint, opts v1.UpdateOptions) (*v1alpha1.AzureEndpoint, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(azureendpointsResource, "status", c.ns, azureEndpoint), &v1alpha1.AzureEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureEndpoint), err
}

// Delete takes name of the azureEndpoint and deletes it. Returns an error if one occurs.
func (c *FakeAzureEndpoints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(azureendpointsResource, c.ns, name), &v1alpha1.AzureEndpoint{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzureEndpoints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(azureendpointsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzureEndpointList{})
	return err
}

// Patch applies the patch and returns the patched azureEndpoint.
func (c *FakeAzureEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AzureEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(azureendpointsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AzureEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureEndpoint), err
}
