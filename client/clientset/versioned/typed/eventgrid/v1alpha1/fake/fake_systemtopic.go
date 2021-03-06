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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/eventgrid/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSystemTopics implements SystemTopicInterface
type FakeSystemTopics struct {
	Fake *FakeEventgridV1alpha1
	ns   string
}

var systemtopicsResource = schema.GroupVersionResource{Group: "eventgrid.azurerm.kubeform.com", Version: "v1alpha1", Resource: "systemtopics"}

var systemtopicsKind = schema.GroupVersionKind{Group: "eventgrid.azurerm.kubeform.com", Version: "v1alpha1", Kind: "SystemTopic"}

// Get takes name of the systemTopic, and returns the corresponding systemTopic object, and an error if there is any.
func (c *FakeSystemTopics) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SystemTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(systemtopicsResource, c.ns, name), &v1alpha1.SystemTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SystemTopic), err
}

// List takes label and field selectors, and returns the list of SystemTopics that match those selectors.
func (c *FakeSystemTopics) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SystemTopicList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(systemtopicsResource, systemtopicsKind, c.ns, opts), &v1alpha1.SystemTopicList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SystemTopicList{ListMeta: obj.(*v1alpha1.SystemTopicList).ListMeta}
	for _, item := range obj.(*v1alpha1.SystemTopicList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested systemTopics.
func (c *FakeSystemTopics) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(systemtopicsResource, c.ns, opts))

}

// Create takes the representation of a systemTopic and creates it.  Returns the server's representation of the systemTopic, and an error, if there is any.
func (c *FakeSystemTopics) Create(ctx context.Context, systemTopic *v1alpha1.SystemTopic, opts v1.CreateOptions) (result *v1alpha1.SystemTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(systemtopicsResource, c.ns, systemTopic), &v1alpha1.SystemTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SystemTopic), err
}

// Update takes the representation of a systemTopic and updates it. Returns the server's representation of the systemTopic, and an error, if there is any.
func (c *FakeSystemTopics) Update(ctx context.Context, systemTopic *v1alpha1.SystemTopic, opts v1.UpdateOptions) (result *v1alpha1.SystemTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(systemtopicsResource, c.ns, systemTopic), &v1alpha1.SystemTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SystemTopic), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSystemTopics) UpdateStatus(ctx context.Context, systemTopic *v1alpha1.SystemTopic, opts v1.UpdateOptions) (*v1alpha1.SystemTopic, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(systemtopicsResource, "status", c.ns, systemTopic), &v1alpha1.SystemTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SystemTopic), err
}

// Delete takes name of the systemTopic and deletes it. Returns an error if one occurs.
func (c *FakeSystemTopics) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(systemtopicsResource, c.ns, name), &v1alpha1.SystemTopic{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSystemTopics) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(systemtopicsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SystemTopicList{})
	return err
}

// Patch applies the patch and returns the patched systemTopic.
func (c *FakeSystemTopics) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SystemTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(systemtopicsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SystemTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SystemTopic), err
}
