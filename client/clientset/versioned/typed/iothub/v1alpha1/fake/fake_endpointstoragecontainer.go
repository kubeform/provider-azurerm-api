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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/iothub/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEndpointStorageContainers implements EndpointStorageContainerInterface
type FakeEndpointStorageContainers struct {
	Fake *FakeIothubV1alpha1
	ns   string
}

var endpointstoragecontainersResource = schema.GroupVersionResource{Group: "iothub.azurerm.kubeform.com", Version: "v1alpha1", Resource: "endpointstoragecontainers"}

var endpointstoragecontainersKind = schema.GroupVersionKind{Group: "iothub.azurerm.kubeform.com", Version: "v1alpha1", Kind: "EndpointStorageContainer"}

// Get takes name of the endpointStorageContainer, and returns the corresponding endpointStorageContainer object, and an error if there is any.
func (c *FakeEndpointStorageContainers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EndpointStorageContainer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(endpointstoragecontainersResource, c.ns, name), &v1alpha1.EndpointStorageContainer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EndpointStorageContainer), err
}

// List takes label and field selectors, and returns the list of EndpointStorageContainers that match those selectors.
func (c *FakeEndpointStorageContainers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EndpointStorageContainerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(endpointstoragecontainersResource, endpointstoragecontainersKind, c.ns, opts), &v1alpha1.EndpointStorageContainerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EndpointStorageContainerList{ListMeta: obj.(*v1alpha1.EndpointStorageContainerList).ListMeta}
	for _, item := range obj.(*v1alpha1.EndpointStorageContainerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested endpointStorageContainers.
func (c *FakeEndpointStorageContainers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(endpointstoragecontainersResource, c.ns, opts))

}

// Create takes the representation of a endpointStorageContainer and creates it.  Returns the server's representation of the endpointStorageContainer, and an error, if there is any.
func (c *FakeEndpointStorageContainers) Create(ctx context.Context, endpointStorageContainer *v1alpha1.EndpointStorageContainer, opts v1.CreateOptions) (result *v1alpha1.EndpointStorageContainer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(endpointstoragecontainersResource, c.ns, endpointStorageContainer), &v1alpha1.EndpointStorageContainer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EndpointStorageContainer), err
}

// Update takes the representation of a endpointStorageContainer and updates it. Returns the server's representation of the endpointStorageContainer, and an error, if there is any.
func (c *FakeEndpointStorageContainers) Update(ctx context.Context, endpointStorageContainer *v1alpha1.EndpointStorageContainer, opts v1.UpdateOptions) (result *v1alpha1.EndpointStorageContainer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(endpointstoragecontainersResource, c.ns, endpointStorageContainer), &v1alpha1.EndpointStorageContainer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EndpointStorageContainer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEndpointStorageContainers) UpdateStatus(ctx context.Context, endpointStorageContainer *v1alpha1.EndpointStorageContainer, opts v1.UpdateOptions) (*v1alpha1.EndpointStorageContainer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(endpointstoragecontainersResource, "status", c.ns, endpointStorageContainer), &v1alpha1.EndpointStorageContainer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EndpointStorageContainer), err
}

// Delete takes name of the endpointStorageContainer and deletes it. Returns an error if one occurs.
func (c *FakeEndpointStorageContainers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(endpointstoragecontainersResource, c.ns, name), &v1alpha1.EndpointStorageContainer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEndpointStorageContainers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(endpointstoragecontainersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EndpointStorageContainerList{})
	return err
}

// Patch applies the patch and returns the patched endpointStorageContainer.
func (c *FakeEndpointStorageContainers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EndpointStorageContainer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(endpointstoragecontainersResource, c.ns, name, pt, data, subresources...), &v1alpha1.EndpointStorageContainer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EndpointStorageContainer), err
}
