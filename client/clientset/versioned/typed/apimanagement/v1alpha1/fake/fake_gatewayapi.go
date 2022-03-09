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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/apimanagement/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGatewayAPIs implements GatewayAPIInterface
type FakeGatewayAPIs struct {
	Fake *FakeApimanagementV1alpha1
	ns   string
}

var gatewayapisResource = schema.GroupVersionResource{Group: "apimanagement.azurerm.kubeform.com", Version: "v1alpha1", Resource: "gatewayapis"}

var gatewayapisKind = schema.GroupVersionKind{Group: "apimanagement.azurerm.kubeform.com", Version: "v1alpha1", Kind: "GatewayAPI"}

// Get takes name of the gatewayAPI, and returns the corresponding gatewayAPI object, and an error if there is any.
func (c *FakeGatewayAPIs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GatewayAPI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(gatewayapisResource, c.ns, name), &v1alpha1.GatewayAPI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GatewayAPI), err
}

// List takes label and field selectors, and returns the list of GatewayAPIs that match those selectors.
func (c *FakeGatewayAPIs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GatewayAPIList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(gatewayapisResource, gatewayapisKind, c.ns, opts), &v1alpha1.GatewayAPIList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GatewayAPIList{ListMeta: obj.(*v1alpha1.GatewayAPIList).ListMeta}
	for _, item := range obj.(*v1alpha1.GatewayAPIList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gatewayAPIs.
func (c *FakeGatewayAPIs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(gatewayapisResource, c.ns, opts))

}

// Create takes the representation of a gatewayAPI and creates it.  Returns the server's representation of the gatewayAPI, and an error, if there is any.
func (c *FakeGatewayAPIs) Create(ctx context.Context, gatewayAPI *v1alpha1.GatewayAPI, opts v1.CreateOptions) (result *v1alpha1.GatewayAPI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(gatewayapisResource, c.ns, gatewayAPI), &v1alpha1.GatewayAPI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GatewayAPI), err
}

// Update takes the representation of a gatewayAPI and updates it. Returns the server's representation of the gatewayAPI, and an error, if there is any.
func (c *FakeGatewayAPIs) Update(ctx context.Context, gatewayAPI *v1alpha1.GatewayAPI, opts v1.UpdateOptions) (result *v1alpha1.GatewayAPI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(gatewayapisResource, c.ns, gatewayAPI), &v1alpha1.GatewayAPI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GatewayAPI), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGatewayAPIs) UpdateStatus(ctx context.Context, gatewayAPI *v1alpha1.GatewayAPI, opts v1.UpdateOptions) (*v1alpha1.GatewayAPI, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(gatewayapisResource, "status", c.ns, gatewayAPI), &v1alpha1.GatewayAPI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GatewayAPI), err
}

// Delete takes name of the gatewayAPI and deletes it. Returns an error if one occurs.
func (c *FakeGatewayAPIs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(gatewayapisResource, c.ns, name), &v1alpha1.GatewayAPI{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGatewayAPIs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(gatewayapisResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.GatewayAPIList{})
	return err
}

// Patch applies the patch and returns the patched gatewayAPI.
func (c *FakeGatewayAPIs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GatewayAPI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gatewayapisResource, c.ns, name, pt, data, subresources...), &v1alpha1.GatewayAPI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GatewayAPI), err
}
