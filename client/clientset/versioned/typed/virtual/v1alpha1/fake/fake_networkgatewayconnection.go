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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/virtual/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkGatewayConnections implements NetworkGatewayConnectionInterface
type FakeNetworkGatewayConnections struct {
	Fake *FakeVirtualV1alpha1
	ns   string
}

var networkgatewayconnectionsResource = schema.GroupVersionResource{Group: "virtual.azurerm.kubeform.com", Version: "v1alpha1", Resource: "networkgatewayconnections"}

var networkgatewayconnectionsKind = schema.GroupVersionKind{Group: "virtual.azurerm.kubeform.com", Version: "v1alpha1", Kind: "NetworkGatewayConnection"}

// Get takes name of the networkGatewayConnection, and returns the corresponding networkGatewayConnection object, and an error if there is any.
func (c *FakeNetworkGatewayConnections) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkGatewayConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkgatewayconnectionsResource, c.ns, name), &v1alpha1.NetworkGatewayConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkGatewayConnection), err
}

// List takes label and field selectors, and returns the list of NetworkGatewayConnections that match those selectors.
func (c *FakeNetworkGatewayConnections) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetworkGatewayConnectionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkgatewayconnectionsResource, networkgatewayconnectionsKind, c.ns, opts), &v1alpha1.NetworkGatewayConnectionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkGatewayConnectionList{ListMeta: obj.(*v1alpha1.NetworkGatewayConnectionList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkGatewayConnectionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkGatewayConnections.
func (c *FakeNetworkGatewayConnections) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkgatewayconnectionsResource, c.ns, opts))

}

// Create takes the representation of a networkGatewayConnection and creates it.  Returns the server's representation of the networkGatewayConnection, and an error, if there is any.
func (c *FakeNetworkGatewayConnections) Create(ctx context.Context, networkGatewayConnection *v1alpha1.NetworkGatewayConnection, opts v1.CreateOptions) (result *v1alpha1.NetworkGatewayConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkgatewayconnectionsResource, c.ns, networkGatewayConnection), &v1alpha1.NetworkGatewayConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkGatewayConnection), err
}

// Update takes the representation of a networkGatewayConnection and updates it. Returns the server's representation of the networkGatewayConnection, and an error, if there is any.
func (c *FakeNetworkGatewayConnections) Update(ctx context.Context, networkGatewayConnection *v1alpha1.NetworkGatewayConnection, opts v1.UpdateOptions) (result *v1alpha1.NetworkGatewayConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkgatewayconnectionsResource, c.ns, networkGatewayConnection), &v1alpha1.NetworkGatewayConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkGatewayConnection), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkGatewayConnections) UpdateStatus(ctx context.Context, networkGatewayConnection *v1alpha1.NetworkGatewayConnection, opts v1.UpdateOptions) (*v1alpha1.NetworkGatewayConnection, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkgatewayconnectionsResource, "status", c.ns, networkGatewayConnection), &v1alpha1.NetworkGatewayConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkGatewayConnection), err
}

// Delete takes name of the networkGatewayConnection and deletes it. Returns an error if one occurs.
func (c *FakeNetworkGatewayConnections) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networkgatewayconnectionsResource, c.ns, name), &v1alpha1.NetworkGatewayConnection{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkGatewayConnections) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkgatewayconnectionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkGatewayConnectionList{})
	return err
}

// Patch applies the patch and returns the patched networkGatewayConnection.
func (c *FakeNetworkGatewayConnections) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkGatewayConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkgatewayconnectionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetworkGatewayConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkGatewayConnection), err
}
