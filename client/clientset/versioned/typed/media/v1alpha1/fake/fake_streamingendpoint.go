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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/media/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStreamingEndpoints implements StreamingEndpointInterface
type FakeStreamingEndpoints struct {
	Fake *FakeMediaV1alpha1
	ns   string
}

var streamingendpointsResource = schema.GroupVersionResource{Group: "media.azurerm.kubeform.com", Version: "v1alpha1", Resource: "streamingendpoints"}

var streamingendpointsKind = schema.GroupVersionKind{Group: "media.azurerm.kubeform.com", Version: "v1alpha1", Kind: "StreamingEndpoint"}

// Get takes name of the streamingEndpoint, and returns the corresponding streamingEndpoint object, and an error if there is any.
func (c *FakeStreamingEndpoints) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.StreamingEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(streamingendpointsResource, c.ns, name), &v1alpha1.StreamingEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamingEndpoint), err
}

// List takes label and field selectors, and returns the list of StreamingEndpoints that match those selectors.
func (c *FakeStreamingEndpoints) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.StreamingEndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(streamingendpointsResource, streamingendpointsKind, c.ns, opts), &v1alpha1.StreamingEndpointList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StreamingEndpointList{ListMeta: obj.(*v1alpha1.StreamingEndpointList).ListMeta}
	for _, item := range obj.(*v1alpha1.StreamingEndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested streamingEndpoints.
func (c *FakeStreamingEndpoints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(streamingendpointsResource, c.ns, opts))

}

// Create takes the representation of a streamingEndpoint and creates it.  Returns the server's representation of the streamingEndpoint, and an error, if there is any.
func (c *FakeStreamingEndpoints) Create(ctx context.Context, streamingEndpoint *v1alpha1.StreamingEndpoint, opts v1.CreateOptions) (result *v1alpha1.StreamingEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(streamingendpointsResource, c.ns, streamingEndpoint), &v1alpha1.StreamingEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamingEndpoint), err
}

// Update takes the representation of a streamingEndpoint and updates it. Returns the server's representation of the streamingEndpoint, and an error, if there is any.
func (c *FakeStreamingEndpoints) Update(ctx context.Context, streamingEndpoint *v1alpha1.StreamingEndpoint, opts v1.UpdateOptions) (result *v1alpha1.StreamingEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(streamingendpointsResource, c.ns, streamingEndpoint), &v1alpha1.StreamingEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamingEndpoint), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStreamingEndpoints) UpdateStatus(ctx context.Context, streamingEndpoint *v1alpha1.StreamingEndpoint, opts v1.UpdateOptions) (*v1alpha1.StreamingEndpoint, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(streamingendpointsResource, "status", c.ns, streamingEndpoint), &v1alpha1.StreamingEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamingEndpoint), err
}

// Delete takes name of the streamingEndpoint and deletes it. Returns an error if one occurs.
func (c *FakeStreamingEndpoints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(streamingendpointsResource, c.ns, name), &v1alpha1.StreamingEndpoint{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStreamingEndpoints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(streamingendpointsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.StreamingEndpointList{})
	return err
}

// Patch applies the patch and returns the patched streamingEndpoint.
func (c *FakeStreamingEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.StreamingEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(streamingendpointsResource, c.ns, name, pt, data, subresources...), &v1alpha1.StreamingEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamingEndpoint), err
}