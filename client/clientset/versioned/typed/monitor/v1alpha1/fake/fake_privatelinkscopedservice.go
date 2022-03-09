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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/monitor/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePrivateLinkScopedServices implements PrivateLinkScopedServiceInterface
type FakePrivateLinkScopedServices struct {
	Fake *FakeMonitorV1alpha1
	ns   string
}

var privatelinkscopedservicesResource = schema.GroupVersionResource{Group: "monitor.azurerm.kubeform.com", Version: "v1alpha1", Resource: "privatelinkscopedservices"}

var privatelinkscopedservicesKind = schema.GroupVersionKind{Group: "monitor.azurerm.kubeform.com", Version: "v1alpha1", Kind: "PrivateLinkScopedService"}

// Get takes name of the privateLinkScopedService, and returns the corresponding privateLinkScopedService object, and an error if there is any.
func (c *FakePrivateLinkScopedServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PrivateLinkScopedService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(privatelinkscopedservicesResource, c.ns, name), &v1alpha1.PrivateLinkScopedService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateLinkScopedService), err
}

// List takes label and field selectors, and returns the list of PrivateLinkScopedServices that match those selectors.
func (c *FakePrivateLinkScopedServices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PrivateLinkScopedServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(privatelinkscopedservicesResource, privatelinkscopedservicesKind, c.ns, opts), &v1alpha1.PrivateLinkScopedServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PrivateLinkScopedServiceList{ListMeta: obj.(*v1alpha1.PrivateLinkScopedServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.PrivateLinkScopedServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested privateLinkScopedServices.
func (c *FakePrivateLinkScopedServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(privatelinkscopedservicesResource, c.ns, opts))

}

// Create takes the representation of a privateLinkScopedService and creates it.  Returns the server's representation of the privateLinkScopedService, and an error, if there is any.
func (c *FakePrivateLinkScopedServices) Create(ctx context.Context, privateLinkScopedService *v1alpha1.PrivateLinkScopedService, opts v1.CreateOptions) (result *v1alpha1.PrivateLinkScopedService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(privatelinkscopedservicesResource, c.ns, privateLinkScopedService), &v1alpha1.PrivateLinkScopedService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateLinkScopedService), err
}

// Update takes the representation of a privateLinkScopedService and updates it. Returns the server's representation of the privateLinkScopedService, and an error, if there is any.
func (c *FakePrivateLinkScopedServices) Update(ctx context.Context, privateLinkScopedService *v1alpha1.PrivateLinkScopedService, opts v1.UpdateOptions) (result *v1alpha1.PrivateLinkScopedService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(privatelinkscopedservicesResource, c.ns, privateLinkScopedService), &v1alpha1.PrivateLinkScopedService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateLinkScopedService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePrivateLinkScopedServices) UpdateStatus(ctx context.Context, privateLinkScopedService *v1alpha1.PrivateLinkScopedService, opts v1.UpdateOptions) (*v1alpha1.PrivateLinkScopedService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(privatelinkscopedservicesResource, "status", c.ns, privateLinkScopedService), &v1alpha1.PrivateLinkScopedService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateLinkScopedService), err
}

// Delete takes name of the privateLinkScopedService and deletes it. Returns an error if one occurs.
func (c *FakePrivateLinkScopedServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(privatelinkscopedservicesResource, c.ns, name), &v1alpha1.PrivateLinkScopedService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePrivateLinkScopedServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(privatelinkscopedservicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PrivateLinkScopedServiceList{})
	return err
}

// Patch applies the patch and returns the patched privateLinkScopedService.
func (c *FakePrivateLinkScopedServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PrivateLinkScopedService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(privatelinkscopedservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.PrivateLinkScopedService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateLinkScopedService), err
}
