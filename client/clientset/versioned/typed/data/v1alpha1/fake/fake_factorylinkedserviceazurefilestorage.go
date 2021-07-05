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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/data/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFactoryLinkedServiceAzureFileStorages implements FactoryLinkedServiceAzureFileStorageInterface
type FakeFactoryLinkedServiceAzureFileStorages struct {
	Fake *FakeDataV1alpha1
	ns   string
}

var factorylinkedserviceazurefilestoragesResource = schema.GroupVersionResource{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Resource: "factorylinkedserviceazurefilestorages"}

var factorylinkedserviceazurefilestoragesKind = schema.GroupVersionKind{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Kind: "FactoryLinkedServiceAzureFileStorage"}

// Get takes name of the factoryLinkedServiceAzureFileStorage, and returns the corresponding factoryLinkedServiceAzureFileStorage object, and an error if there is any.
func (c *FakeFactoryLinkedServiceAzureFileStorages) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FactoryLinkedServiceAzureFileStorage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(factorylinkedserviceazurefilestoragesResource, c.ns, name), &v1alpha1.FactoryLinkedServiceAzureFileStorage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceAzureFileStorage), err
}

// List takes label and field selectors, and returns the list of FactoryLinkedServiceAzureFileStorages that match those selectors.
func (c *FakeFactoryLinkedServiceAzureFileStorages) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FactoryLinkedServiceAzureFileStorageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(factorylinkedserviceazurefilestoragesResource, factorylinkedserviceazurefilestoragesKind, c.ns, opts), &v1alpha1.FactoryLinkedServiceAzureFileStorageList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FactoryLinkedServiceAzureFileStorageList{ListMeta: obj.(*v1alpha1.FactoryLinkedServiceAzureFileStorageList).ListMeta}
	for _, item := range obj.(*v1alpha1.FactoryLinkedServiceAzureFileStorageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested factoryLinkedServiceAzureFileStorages.
func (c *FakeFactoryLinkedServiceAzureFileStorages) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(factorylinkedserviceazurefilestoragesResource, c.ns, opts))

}

// Create takes the representation of a factoryLinkedServiceAzureFileStorage and creates it.  Returns the server's representation of the factoryLinkedServiceAzureFileStorage, and an error, if there is any.
func (c *FakeFactoryLinkedServiceAzureFileStorages) Create(ctx context.Context, factoryLinkedServiceAzureFileStorage *v1alpha1.FactoryLinkedServiceAzureFileStorage, opts v1.CreateOptions) (result *v1alpha1.FactoryLinkedServiceAzureFileStorage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(factorylinkedserviceazurefilestoragesResource, c.ns, factoryLinkedServiceAzureFileStorage), &v1alpha1.FactoryLinkedServiceAzureFileStorage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceAzureFileStorage), err
}

// Update takes the representation of a factoryLinkedServiceAzureFileStorage and updates it. Returns the server's representation of the factoryLinkedServiceAzureFileStorage, and an error, if there is any.
func (c *FakeFactoryLinkedServiceAzureFileStorages) Update(ctx context.Context, factoryLinkedServiceAzureFileStorage *v1alpha1.FactoryLinkedServiceAzureFileStorage, opts v1.UpdateOptions) (result *v1alpha1.FactoryLinkedServiceAzureFileStorage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(factorylinkedserviceazurefilestoragesResource, c.ns, factoryLinkedServiceAzureFileStorage), &v1alpha1.FactoryLinkedServiceAzureFileStorage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceAzureFileStorage), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFactoryLinkedServiceAzureFileStorages) UpdateStatus(ctx context.Context, factoryLinkedServiceAzureFileStorage *v1alpha1.FactoryLinkedServiceAzureFileStorage, opts v1.UpdateOptions) (*v1alpha1.FactoryLinkedServiceAzureFileStorage, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(factorylinkedserviceazurefilestoragesResource, "status", c.ns, factoryLinkedServiceAzureFileStorage), &v1alpha1.FactoryLinkedServiceAzureFileStorage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceAzureFileStorage), err
}

// Delete takes name of the factoryLinkedServiceAzureFileStorage and deletes it. Returns an error if one occurs.
func (c *FakeFactoryLinkedServiceAzureFileStorages) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(factorylinkedserviceazurefilestoragesResource, c.ns, name), &v1alpha1.FactoryLinkedServiceAzureFileStorage{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFactoryLinkedServiceAzureFileStorages) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(factorylinkedserviceazurefilestoragesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FactoryLinkedServiceAzureFileStorageList{})
	return err
}

// Patch applies the patch and returns the patched factoryLinkedServiceAzureFileStorage.
func (c *FakeFactoryLinkedServiceAzureFileStorages) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FactoryLinkedServiceAzureFileStorage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(factorylinkedserviceazurefilestoragesResource, c.ns, name, pt, data, subresources...), &v1alpha1.FactoryLinkedServiceAzureFileStorage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceAzureFileStorage), err
}