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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/disk/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEncryptionSets implements EncryptionSetInterface
type FakeEncryptionSets struct {
	Fake *FakeDiskV1alpha1
	ns   string
}

var encryptionsetsResource = schema.GroupVersionResource{Group: "disk.azurerm.kubeform.com", Version: "v1alpha1", Resource: "encryptionsets"}

var encryptionsetsKind = schema.GroupVersionKind{Group: "disk.azurerm.kubeform.com", Version: "v1alpha1", Kind: "EncryptionSet"}

// Get takes name of the encryptionSet, and returns the corresponding encryptionSet object, and an error if there is any.
func (c *FakeEncryptionSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EncryptionSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(encryptionsetsResource, c.ns, name), &v1alpha1.EncryptionSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EncryptionSet), err
}

// List takes label and field selectors, and returns the list of EncryptionSets that match those selectors.
func (c *FakeEncryptionSets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EncryptionSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(encryptionsetsResource, encryptionsetsKind, c.ns, opts), &v1alpha1.EncryptionSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EncryptionSetList{ListMeta: obj.(*v1alpha1.EncryptionSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.EncryptionSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested encryptionSets.
func (c *FakeEncryptionSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(encryptionsetsResource, c.ns, opts))

}

// Create takes the representation of a encryptionSet and creates it.  Returns the server's representation of the encryptionSet, and an error, if there is any.
func (c *FakeEncryptionSets) Create(ctx context.Context, encryptionSet *v1alpha1.EncryptionSet, opts v1.CreateOptions) (result *v1alpha1.EncryptionSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(encryptionsetsResource, c.ns, encryptionSet), &v1alpha1.EncryptionSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EncryptionSet), err
}

// Update takes the representation of a encryptionSet and updates it. Returns the server's representation of the encryptionSet, and an error, if there is any.
func (c *FakeEncryptionSets) Update(ctx context.Context, encryptionSet *v1alpha1.EncryptionSet, opts v1.UpdateOptions) (result *v1alpha1.EncryptionSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(encryptionsetsResource, c.ns, encryptionSet), &v1alpha1.EncryptionSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EncryptionSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEncryptionSets) UpdateStatus(ctx context.Context, encryptionSet *v1alpha1.EncryptionSet, opts v1.UpdateOptions) (*v1alpha1.EncryptionSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(encryptionsetsResource, "status", c.ns, encryptionSet), &v1alpha1.EncryptionSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EncryptionSet), err
}

// Delete takes name of the encryptionSet and deletes it. Returns an error if one occurs.
func (c *FakeEncryptionSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(encryptionsetsResource, c.ns, name), &v1alpha1.EncryptionSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEncryptionSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(encryptionsetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EncryptionSetList{})
	return err
}

// Patch applies the patch and returns the patched encryptionSet.
func (c *FakeEncryptionSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EncryptionSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(encryptionsetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.EncryptionSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EncryptionSet), err
}
