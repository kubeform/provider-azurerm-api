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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/devtest/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVirtualNetworks implements VirtualNetworkInterface
type FakeVirtualNetworks struct {
	Fake *FakeDevtestV1alpha1
	ns   string
}

var virtualnetworksResource = schema.GroupVersionResource{Group: "devtest.azurerm.kubeform.com", Version: "v1alpha1", Resource: "virtualnetworks"}

var virtualnetworksKind = schema.GroupVersionKind{Group: "devtest.azurerm.kubeform.com", Version: "v1alpha1", Kind: "VirtualNetwork"}

// Get takes name of the virtualNetwork, and returns the corresponding virtualNetwork object, and an error if there is any.
func (c *FakeVirtualNetworks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VirtualNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualnetworksResource, c.ns, name), &v1alpha1.VirtualNetwork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualNetwork), err
}

// List takes label and field selectors, and returns the list of VirtualNetworks that match those selectors.
func (c *FakeVirtualNetworks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VirtualNetworkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualnetworksResource, virtualnetworksKind, c.ns, opts), &v1alpha1.VirtualNetworkList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VirtualNetworkList{ListMeta: obj.(*v1alpha1.VirtualNetworkList).ListMeta}
	for _, item := range obj.(*v1alpha1.VirtualNetworkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualNetworks.
func (c *FakeVirtualNetworks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualnetworksResource, c.ns, opts))

}

// Create takes the representation of a virtualNetwork and creates it.  Returns the server's representation of the virtualNetwork, and an error, if there is any.
func (c *FakeVirtualNetworks) Create(ctx context.Context, virtualNetwork *v1alpha1.VirtualNetwork, opts v1.CreateOptions) (result *v1alpha1.VirtualNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualnetworksResource, c.ns, virtualNetwork), &v1alpha1.VirtualNetwork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualNetwork), err
}

// Update takes the representation of a virtualNetwork and updates it. Returns the server's representation of the virtualNetwork, and an error, if there is any.
func (c *FakeVirtualNetworks) Update(ctx context.Context, virtualNetwork *v1alpha1.VirtualNetwork, opts v1.UpdateOptions) (result *v1alpha1.VirtualNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualnetworksResource, c.ns, virtualNetwork), &v1alpha1.VirtualNetwork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualNetwork), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualNetworks) UpdateStatus(ctx context.Context, virtualNetwork *v1alpha1.VirtualNetwork, opts v1.UpdateOptions) (*v1alpha1.VirtualNetwork, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualnetworksResource, "status", c.ns, virtualNetwork), &v1alpha1.VirtualNetwork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualNetwork), err
}

// Delete takes name of the virtualNetwork and deletes it. Returns an error if one occurs.
func (c *FakeVirtualNetworks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualnetworksResource, c.ns, name), &v1alpha1.VirtualNetwork{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualNetworks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualnetworksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VirtualNetworkList{})
	return err
}

// Patch applies the patch and returns the patched virtualNetwork.
func (c *FakeVirtualNetworks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VirtualNetwork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualnetworksResource, c.ns, name, pt, data, subresources...), &v1alpha1.VirtualNetwork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualNetwork), err
}