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

// FakeChannelAlexas implements ChannelAlexaInterface
type FakeChannelAlexas struct {
	Fake *FakeBotV1alpha1
	ns   string
}

var channelalexasResource = schema.GroupVersionResource{Group: "bot.azurerm.kubeform.com", Version: "v1alpha1", Resource: "channelalexas"}

var channelalexasKind = schema.GroupVersionKind{Group: "bot.azurerm.kubeform.com", Version: "v1alpha1", Kind: "ChannelAlexa"}

// Get takes name of the channelAlexa, and returns the corresponding channelAlexa object, and an error if there is any.
func (c *FakeChannelAlexas) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ChannelAlexa, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(channelalexasResource, c.ns, name), &v1alpha1.ChannelAlexa{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChannelAlexa), err
}

// List takes label and field selectors, and returns the list of ChannelAlexas that match those selectors.
func (c *FakeChannelAlexas) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ChannelAlexaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(channelalexasResource, channelalexasKind, c.ns, opts), &v1alpha1.ChannelAlexaList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ChannelAlexaList{ListMeta: obj.(*v1alpha1.ChannelAlexaList).ListMeta}
	for _, item := range obj.(*v1alpha1.ChannelAlexaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested channelAlexas.
func (c *FakeChannelAlexas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(channelalexasResource, c.ns, opts))

}

// Create takes the representation of a channelAlexa and creates it.  Returns the server's representation of the channelAlexa, and an error, if there is any.
func (c *FakeChannelAlexas) Create(ctx context.Context, channelAlexa *v1alpha1.ChannelAlexa, opts v1.CreateOptions) (result *v1alpha1.ChannelAlexa, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(channelalexasResource, c.ns, channelAlexa), &v1alpha1.ChannelAlexa{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChannelAlexa), err
}

// Update takes the representation of a channelAlexa and updates it. Returns the server's representation of the channelAlexa, and an error, if there is any.
func (c *FakeChannelAlexas) Update(ctx context.Context, channelAlexa *v1alpha1.ChannelAlexa, opts v1.UpdateOptions) (result *v1alpha1.ChannelAlexa, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(channelalexasResource, c.ns, channelAlexa), &v1alpha1.ChannelAlexa{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChannelAlexa), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeChannelAlexas) UpdateStatus(ctx context.Context, channelAlexa *v1alpha1.ChannelAlexa, opts v1.UpdateOptions) (*v1alpha1.ChannelAlexa, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(channelalexasResource, "status", c.ns, channelAlexa), &v1alpha1.ChannelAlexa{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChannelAlexa), err
}

// Delete takes name of the channelAlexa and deletes it. Returns an error if one occurs.
func (c *FakeChannelAlexas) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(channelalexasResource, c.ns, name), &v1alpha1.ChannelAlexa{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeChannelAlexas) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(channelalexasResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ChannelAlexaList{})
	return err
}

// Patch applies the patch and returns the patched channelAlexa.
func (c *FakeChannelAlexas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ChannelAlexa, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(channelalexasResource, c.ns, name, pt, data, subresources...), &v1alpha1.ChannelAlexa{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChannelAlexa), err
}
