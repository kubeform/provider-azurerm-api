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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/bot/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeChannelLines implements ChannelLineInterface
type FakeChannelLines struct {
	Fake *FakeBotV1alpha1
	ns   string
}

var channellinesResource = schema.GroupVersionResource{Group: "bot.azurerm.kubeform.com", Version: "v1alpha1", Resource: "channellines"}

var channellinesKind = schema.GroupVersionKind{Group: "bot.azurerm.kubeform.com", Version: "v1alpha1", Kind: "ChannelLine"}

// Get takes name of the channelLine, and returns the corresponding channelLine object, and an error if there is any.
func (c *FakeChannelLines) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ChannelLine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(channellinesResource, c.ns, name), &v1alpha1.ChannelLine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChannelLine), err
}

// List takes label and field selectors, and returns the list of ChannelLines that match those selectors.
func (c *FakeChannelLines) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ChannelLineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(channellinesResource, channellinesKind, c.ns, opts), &v1alpha1.ChannelLineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ChannelLineList{ListMeta: obj.(*v1alpha1.ChannelLineList).ListMeta}
	for _, item := range obj.(*v1alpha1.ChannelLineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested channelLines.
func (c *FakeChannelLines) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(channellinesResource, c.ns, opts))

}

// Create takes the representation of a channelLine and creates it.  Returns the server's representation of the channelLine, and an error, if there is any.
func (c *FakeChannelLines) Create(ctx context.Context, channelLine *v1alpha1.ChannelLine, opts v1.CreateOptions) (result *v1alpha1.ChannelLine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(channellinesResource, c.ns, channelLine), &v1alpha1.ChannelLine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChannelLine), err
}

// Update takes the representation of a channelLine and updates it. Returns the server's representation of the channelLine, and an error, if there is any.
func (c *FakeChannelLines) Update(ctx context.Context, channelLine *v1alpha1.ChannelLine, opts v1.UpdateOptions) (result *v1alpha1.ChannelLine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(channellinesResource, c.ns, channelLine), &v1alpha1.ChannelLine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChannelLine), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeChannelLines) UpdateStatus(ctx context.Context, channelLine *v1alpha1.ChannelLine, opts v1.UpdateOptions) (*v1alpha1.ChannelLine, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(channellinesResource, "status", c.ns, channelLine), &v1alpha1.ChannelLine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChannelLine), err
}

// Delete takes name of the channelLine and deletes it. Returns an error if one occurs.
func (c *FakeChannelLines) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(channellinesResource, c.ns, name), &v1alpha1.ChannelLine{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeChannelLines) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(channellinesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ChannelLineList{})
	return err
}

// Patch applies the patch and returns the patched channelLine.
func (c *FakeChannelLines) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ChannelLine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(channellinesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ChannelLine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChannelLine), err
}
