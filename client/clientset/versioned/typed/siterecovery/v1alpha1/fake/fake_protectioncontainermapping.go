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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/siterecovery/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeProtectionContainerMappings implements ProtectionContainerMappingInterface
type FakeProtectionContainerMappings struct {
	Fake *FakeSiterecoveryV1alpha1
	ns   string
}

var protectioncontainermappingsResource = schema.GroupVersionResource{Group: "siterecovery.azurerm.kubeform.com", Version: "v1alpha1", Resource: "protectioncontainermappings"}

var protectioncontainermappingsKind = schema.GroupVersionKind{Group: "siterecovery.azurerm.kubeform.com", Version: "v1alpha1", Kind: "ProtectionContainerMapping"}

// Get takes name of the protectionContainerMapping, and returns the corresponding protectionContainerMapping object, and an error if there is any.
func (c *FakeProtectionContainerMappings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ProtectionContainerMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(protectioncontainermappingsResource, c.ns, name), &v1alpha1.ProtectionContainerMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProtectionContainerMapping), err
}

// List takes label and field selectors, and returns the list of ProtectionContainerMappings that match those selectors.
func (c *FakeProtectionContainerMappings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ProtectionContainerMappingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(protectioncontainermappingsResource, protectioncontainermappingsKind, c.ns, opts), &v1alpha1.ProtectionContainerMappingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ProtectionContainerMappingList{ListMeta: obj.(*v1alpha1.ProtectionContainerMappingList).ListMeta}
	for _, item := range obj.(*v1alpha1.ProtectionContainerMappingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested protectionContainerMappings.
func (c *FakeProtectionContainerMappings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(protectioncontainermappingsResource, c.ns, opts))

}

// Create takes the representation of a protectionContainerMapping and creates it.  Returns the server's representation of the protectionContainerMapping, and an error, if there is any.
func (c *FakeProtectionContainerMappings) Create(ctx context.Context, protectionContainerMapping *v1alpha1.ProtectionContainerMapping, opts v1.CreateOptions) (result *v1alpha1.ProtectionContainerMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(protectioncontainermappingsResource, c.ns, protectionContainerMapping), &v1alpha1.ProtectionContainerMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProtectionContainerMapping), err
}

// Update takes the representation of a protectionContainerMapping and updates it. Returns the server's representation of the protectionContainerMapping, and an error, if there is any.
func (c *FakeProtectionContainerMappings) Update(ctx context.Context, protectionContainerMapping *v1alpha1.ProtectionContainerMapping, opts v1.UpdateOptions) (result *v1alpha1.ProtectionContainerMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(protectioncontainermappingsResource, c.ns, protectionContainerMapping), &v1alpha1.ProtectionContainerMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProtectionContainerMapping), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProtectionContainerMappings) UpdateStatus(ctx context.Context, protectionContainerMapping *v1alpha1.ProtectionContainerMapping, opts v1.UpdateOptions) (*v1alpha1.ProtectionContainerMapping, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(protectioncontainermappingsResource, "status", c.ns, protectionContainerMapping), &v1alpha1.ProtectionContainerMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProtectionContainerMapping), err
}

// Delete takes name of the protectionContainerMapping and deletes it. Returns an error if one occurs.
func (c *FakeProtectionContainerMappings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(protectioncontainermappingsResource, c.ns, name), &v1alpha1.ProtectionContainerMapping{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProtectionContainerMappings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(protectioncontainermappingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ProtectionContainerMappingList{})
	return err
}

// Patch applies the patch and returns the patched protectionContainerMapping.
func (c *FakeProtectionContainerMappings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ProtectionContainerMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(protectioncontainermappingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ProtectionContainerMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProtectionContainerMapping), err
}