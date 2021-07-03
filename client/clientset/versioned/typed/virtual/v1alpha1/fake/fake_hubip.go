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

// FakeHubIPs implements HubIPInterface
type FakeHubIPs struct {
	Fake *FakeVirtualV1alpha1
	ns   string
}

var hubipsResource = schema.GroupVersionResource{Group: "virtual.azurerm.kubeform.com", Version: "v1alpha1", Resource: "hubips"}

var hubipsKind = schema.GroupVersionKind{Group: "virtual.azurerm.kubeform.com", Version: "v1alpha1", Kind: "HubIP"}

// Get takes name of the hubIP, and returns the corresponding hubIP object, and an error if there is any.
func (c *FakeHubIPs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HubIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hubipsResource, c.ns, name), &v1alpha1.HubIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HubIP), err
}

// List takes label and field selectors, and returns the list of HubIPs that match those selectors.
func (c *FakeHubIPs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HubIPList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hubipsResource, hubipsKind, c.ns, opts), &v1alpha1.HubIPList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HubIPList{ListMeta: obj.(*v1alpha1.HubIPList).ListMeta}
	for _, item := range obj.(*v1alpha1.HubIPList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hubIPs.
func (c *FakeHubIPs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hubipsResource, c.ns, opts))

}

// Create takes the representation of a hubIP and creates it.  Returns the server's representation of the hubIP, and an error, if there is any.
func (c *FakeHubIPs) Create(ctx context.Context, hubIP *v1alpha1.HubIP, opts v1.CreateOptions) (result *v1alpha1.HubIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hubipsResource, c.ns, hubIP), &v1alpha1.HubIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HubIP), err
}

// Update takes the representation of a hubIP and updates it. Returns the server's representation of the hubIP, and an error, if there is any.
func (c *FakeHubIPs) Update(ctx context.Context, hubIP *v1alpha1.HubIP, opts v1.UpdateOptions) (result *v1alpha1.HubIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hubipsResource, c.ns, hubIP), &v1alpha1.HubIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HubIP), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHubIPs) UpdateStatus(ctx context.Context, hubIP *v1alpha1.HubIP, opts v1.UpdateOptions) (*v1alpha1.HubIP, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(hubipsResource, "status", c.ns, hubIP), &v1alpha1.HubIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HubIP), err
}

// Delete takes name of the hubIP and deletes it. Returns an error if one occurs.
func (c *FakeHubIPs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(hubipsResource, c.ns, name), &v1alpha1.HubIP{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHubIPs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hubipsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HubIPList{})
	return err
}

// Patch applies the patch and returns the patched hubIP.
func (c *FakeHubIPs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HubIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hubipsResource, c.ns, name, pt, data, subresources...), &v1alpha1.HubIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HubIP), err
}
