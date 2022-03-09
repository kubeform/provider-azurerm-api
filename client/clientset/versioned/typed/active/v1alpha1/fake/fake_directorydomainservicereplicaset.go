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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/active/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDirectoryDomainServiceReplicaSets implements DirectoryDomainServiceReplicaSetInterface
type FakeDirectoryDomainServiceReplicaSets struct {
	Fake *FakeActiveV1alpha1
	ns   string
}

var directorydomainservicereplicasetsResource = schema.GroupVersionResource{Group: "active.azurerm.kubeform.com", Version: "v1alpha1", Resource: "directorydomainservicereplicasets"}

var directorydomainservicereplicasetsKind = schema.GroupVersionKind{Group: "active.azurerm.kubeform.com", Version: "v1alpha1", Kind: "DirectoryDomainServiceReplicaSet"}

// Get takes name of the directoryDomainServiceReplicaSet, and returns the corresponding directoryDomainServiceReplicaSet object, and an error if there is any.
func (c *FakeDirectoryDomainServiceReplicaSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DirectoryDomainServiceReplicaSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(directorydomainservicereplicasetsResource, c.ns, name), &v1alpha1.DirectoryDomainServiceReplicaSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DirectoryDomainServiceReplicaSet), err
}

// List takes label and field selectors, and returns the list of DirectoryDomainServiceReplicaSets that match those selectors.
func (c *FakeDirectoryDomainServiceReplicaSets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DirectoryDomainServiceReplicaSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(directorydomainservicereplicasetsResource, directorydomainservicereplicasetsKind, c.ns, opts), &v1alpha1.DirectoryDomainServiceReplicaSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DirectoryDomainServiceReplicaSetList{ListMeta: obj.(*v1alpha1.DirectoryDomainServiceReplicaSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.DirectoryDomainServiceReplicaSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested directoryDomainServiceReplicaSets.
func (c *FakeDirectoryDomainServiceReplicaSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(directorydomainservicereplicasetsResource, c.ns, opts))

}

// Create takes the representation of a directoryDomainServiceReplicaSet and creates it.  Returns the server's representation of the directoryDomainServiceReplicaSet, and an error, if there is any.
func (c *FakeDirectoryDomainServiceReplicaSets) Create(ctx context.Context, directoryDomainServiceReplicaSet *v1alpha1.DirectoryDomainServiceReplicaSet, opts v1.CreateOptions) (result *v1alpha1.DirectoryDomainServiceReplicaSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(directorydomainservicereplicasetsResource, c.ns, directoryDomainServiceReplicaSet), &v1alpha1.DirectoryDomainServiceReplicaSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DirectoryDomainServiceReplicaSet), err
}

// Update takes the representation of a directoryDomainServiceReplicaSet and updates it. Returns the server's representation of the directoryDomainServiceReplicaSet, and an error, if there is any.
func (c *FakeDirectoryDomainServiceReplicaSets) Update(ctx context.Context, directoryDomainServiceReplicaSet *v1alpha1.DirectoryDomainServiceReplicaSet, opts v1.UpdateOptions) (result *v1alpha1.DirectoryDomainServiceReplicaSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(directorydomainservicereplicasetsResource, c.ns, directoryDomainServiceReplicaSet), &v1alpha1.DirectoryDomainServiceReplicaSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DirectoryDomainServiceReplicaSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDirectoryDomainServiceReplicaSets) UpdateStatus(ctx context.Context, directoryDomainServiceReplicaSet *v1alpha1.DirectoryDomainServiceReplicaSet, opts v1.UpdateOptions) (*v1alpha1.DirectoryDomainServiceReplicaSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(directorydomainservicereplicasetsResource, "status", c.ns, directoryDomainServiceReplicaSet), &v1alpha1.DirectoryDomainServiceReplicaSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DirectoryDomainServiceReplicaSet), err
}

// Delete takes name of the directoryDomainServiceReplicaSet and deletes it. Returns an error if one occurs.
func (c *FakeDirectoryDomainServiceReplicaSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(directorydomainservicereplicasetsResource, c.ns, name), &v1alpha1.DirectoryDomainServiceReplicaSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDirectoryDomainServiceReplicaSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(directorydomainservicereplicasetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DirectoryDomainServiceReplicaSetList{})
	return err
}

// Patch applies the patch and returns the patched directoryDomainServiceReplicaSet.
func (c *FakeDirectoryDomainServiceReplicaSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DirectoryDomainServiceReplicaSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(directorydomainservicereplicasetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DirectoryDomainServiceReplicaSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DirectoryDomainServiceReplicaSet), err
}
