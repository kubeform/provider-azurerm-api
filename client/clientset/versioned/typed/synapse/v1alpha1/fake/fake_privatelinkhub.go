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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/synapse/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePrivateLinkHubs implements PrivateLinkHubInterface
type FakePrivateLinkHubs struct {
	Fake *FakeSynapseV1alpha1
	ns   string
}

var privatelinkhubsResource = schema.GroupVersionResource{Group: "synapse.azurerm.kubeform.com", Version: "v1alpha1", Resource: "privatelinkhubs"}

var privatelinkhubsKind = schema.GroupVersionKind{Group: "synapse.azurerm.kubeform.com", Version: "v1alpha1", Kind: "PrivateLinkHub"}

// Get takes name of the privateLinkHub, and returns the corresponding privateLinkHub object, and an error if there is any.
func (c *FakePrivateLinkHubs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PrivateLinkHub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(privatelinkhubsResource, c.ns, name), &v1alpha1.PrivateLinkHub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateLinkHub), err
}

// List takes label and field selectors, and returns the list of PrivateLinkHubs that match those selectors.
func (c *FakePrivateLinkHubs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PrivateLinkHubList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(privatelinkhubsResource, privatelinkhubsKind, c.ns, opts), &v1alpha1.PrivateLinkHubList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PrivateLinkHubList{ListMeta: obj.(*v1alpha1.PrivateLinkHubList).ListMeta}
	for _, item := range obj.(*v1alpha1.PrivateLinkHubList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested privateLinkHubs.
func (c *FakePrivateLinkHubs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(privatelinkhubsResource, c.ns, opts))

}

// Create takes the representation of a privateLinkHub and creates it.  Returns the server's representation of the privateLinkHub, and an error, if there is any.
func (c *FakePrivateLinkHubs) Create(ctx context.Context, privateLinkHub *v1alpha1.PrivateLinkHub, opts v1.CreateOptions) (result *v1alpha1.PrivateLinkHub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(privatelinkhubsResource, c.ns, privateLinkHub), &v1alpha1.PrivateLinkHub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateLinkHub), err
}

// Update takes the representation of a privateLinkHub and updates it. Returns the server's representation of the privateLinkHub, and an error, if there is any.
func (c *FakePrivateLinkHubs) Update(ctx context.Context, privateLinkHub *v1alpha1.PrivateLinkHub, opts v1.UpdateOptions) (result *v1alpha1.PrivateLinkHub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(privatelinkhubsResource, c.ns, privateLinkHub), &v1alpha1.PrivateLinkHub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateLinkHub), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePrivateLinkHubs) UpdateStatus(ctx context.Context, privateLinkHub *v1alpha1.PrivateLinkHub, opts v1.UpdateOptions) (*v1alpha1.PrivateLinkHub, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(privatelinkhubsResource, "status", c.ns, privateLinkHub), &v1alpha1.PrivateLinkHub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateLinkHub), err
}

// Delete takes name of the privateLinkHub and deletes it. Returns an error if one occurs.
func (c *FakePrivateLinkHubs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(privatelinkhubsResource, c.ns, name), &v1alpha1.PrivateLinkHub{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePrivateLinkHubs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(privatelinkhubsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PrivateLinkHubList{})
	return err
}

// Patch applies the patch and returns the patched privateLinkHub.
func (c *FakePrivateLinkHubs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PrivateLinkHub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(privatelinkhubsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PrivateLinkHub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateLinkHub), err
}
