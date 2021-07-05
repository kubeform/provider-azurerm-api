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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/redis/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLinkedServers implements LinkedServerInterface
type FakeLinkedServers struct {
	Fake *FakeRedisV1alpha1
	ns   string
}

var linkedserversResource = schema.GroupVersionResource{Group: "redis.azurerm.kubeform.com", Version: "v1alpha1", Resource: "linkedservers"}

var linkedserversKind = schema.GroupVersionKind{Group: "redis.azurerm.kubeform.com", Version: "v1alpha1", Kind: "LinkedServer"}

// Get takes name of the linkedServer, and returns the corresponding linkedServer object, and an error if there is any.
func (c *FakeLinkedServers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LinkedServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(linkedserversResource, c.ns, name), &v1alpha1.LinkedServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinkedServer), err
}

// List takes label and field selectors, and returns the list of LinkedServers that match those selectors.
func (c *FakeLinkedServers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LinkedServerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(linkedserversResource, linkedserversKind, c.ns, opts), &v1alpha1.LinkedServerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LinkedServerList{ListMeta: obj.(*v1alpha1.LinkedServerList).ListMeta}
	for _, item := range obj.(*v1alpha1.LinkedServerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested linkedServers.
func (c *FakeLinkedServers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(linkedserversResource, c.ns, opts))

}

// Create takes the representation of a linkedServer and creates it.  Returns the server's representation of the linkedServer, and an error, if there is any.
func (c *FakeLinkedServers) Create(ctx context.Context, linkedServer *v1alpha1.LinkedServer, opts v1.CreateOptions) (result *v1alpha1.LinkedServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(linkedserversResource, c.ns, linkedServer), &v1alpha1.LinkedServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinkedServer), err
}

// Update takes the representation of a linkedServer and updates it. Returns the server's representation of the linkedServer, and an error, if there is any.
func (c *FakeLinkedServers) Update(ctx context.Context, linkedServer *v1alpha1.LinkedServer, opts v1.UpdateOptions) (result *v1alpha1.LinkedServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(linkedserversResource, c.ns, linkedServer), &v1alpha1.LinkedServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinkedServer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLinkedServers) UpdateStatus(ctx context.Context, linkedServer *v1alpha1.LinkedServer, opts v1.UpdateOptions) (*v1alpha1.LinkedServer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(linkedserversResource, "status", c.ns, linkedServer), &v1alpha1.LinkedServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinkedServer), err
}

// Delete takes name of the linkedServer and deletes it. Returns an error if one occurs.
func (c *FakeLinkedServers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(linkedserversResource, c.ns, name), &v1alpha1.LinkedServer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLinkedServers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(linkedserversResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LinkedServerList{})
	return err
}

// Patch applies the patch and returns the patched linkedServer.
func (c *FakeLinkedServers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LinkedServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(linkedserversResource, c.ns, name, pt, data, subresources...), &v1alpha1.LinkedServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinkedServer), err
}