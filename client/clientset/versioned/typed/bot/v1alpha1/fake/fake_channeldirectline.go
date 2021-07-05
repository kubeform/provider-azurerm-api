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

// FakeChannelDirectlines implements ChannelDirectlineInterface
type FakeChannelDirectlines struct {
	Fake *FakeBotV1alpha1
	ns   string
}

var channeldirectlinesResource = schema.GroupVersionResource{Group: "bot.azurerm.kubeform.com", Version: "v1alpha1", Resource: "channeldirectlines"}

var channeldirectlinesKind = schema.GroupVersionKind{Group: "bot.azurerm.kubeform.com", Version: "v1alpha1", Kind: "ChannelDirectline"}

// Get takes name of the channelDirectline, and returns the corresponding channelDirectline object, and an error if there is any.
func (c *FakeChannelDirectlines) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ChannelDirectline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(channeldirectlinesResource, c.ns, name), &v1alpha1.ChannelDirectline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChannelDirectline), err
}

// List takes label and field selectors, and returns the list of ChannelDirectlines that match those selectors.
func (c *FakeChannelDirectlines) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ChannelDirectlineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(channeldirectlinesResource, channeldirectlinesKind, c.ns, opts), &v1alpha1.ChannelDirectlineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ChannelDirectlineList{ListMeta: obj.(*v1alpha1.ChannelDirectlineList).ListMeta}
	for _, item := range obj.(*v1alpha1.ChannelDirectlineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested channelDirectlines.
func (c *FakeChannelDirectlines) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(channeldirectlinesResource, c.ns, opts))

}

// Create takes the representation of a channelDirectline and creates it.  Returns the server's representation of the channelDirectline, and an error, if there is any.
func (c *FakeChannelDirectlines) Create(ctx context.Context, channelDirectline *v1alpha1.ChannelDirectline, opts v1.CreateOptions) (result *v1alpha1.ChannelDirectline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(channeldirectlinesResource, c.ns, channelDirectline), &v1alpha1.ChannelDirectline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChannelDirectline), err
}

// Update takes the representation of a channelDirectline and updates it. Returns the server's representation of the channelDirectline, and an error, if there is any.
func (c *FakeChannelDirectlines) Update(ctx context.Context, channelDirectline *v1alpha1.ChannelDirectline, opts v1.UpdateOptions) (result *v1alpha1.ChannelDirectline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(channeldirectlinesResource, c.ns, channelDirectline), &v1alpha1.ChannelDirectline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChannelDirectline), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeChannelDirectlines) UpdateStatus(ctx context.Context, channelDirectline *v1alpha1.ChannelDirectline, opts v1.UpdateOptions) (*v1alpha1.ChannelDirectline, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(channeldirectlinesResource, "status", c.ns, channelDirectline), &v1alpha1.ChannelDirectline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChannelDirectline), err
}

// Delete takes name of the channelDirectline and deletes it. Returns an error if one occurs.
func (c *FakeChannelDirectlines) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(channeldirectlinesResource, c.ns, name), &v1alpha1.ChannelDirectline{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeChannelDirectlines) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(channeldirectlinesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ChannelDirectlineList{})
	return err
}

// Patch applies the patch and returns the patched channelDirectline.
func (c *FakeChannelDirectlines) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ChannelDirectline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(channeldirectlinesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ChannelDirectline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChannelDirectline), err
}