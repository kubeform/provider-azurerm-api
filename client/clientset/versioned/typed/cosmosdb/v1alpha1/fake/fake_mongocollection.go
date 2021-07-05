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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/cosmosdb/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMongoCollections implements MongoCollectionInterface
type FakeMongoCollections struct {
	Fake *FakeCosmosdbV1alpha1
	ns   string
}

var mongocollectionsResource = schema.GroupVersionResource{Group: "cosmosdb.azurerm.kubeform.com", Version: "v1alpha1", Resource: "mongocollections"}

var mongocollectionsKind = schema.GroupVersionKind{Group: "cosmosdb.azurerm.kubeform.com", Version: "v1alpha1", Kind: "MongoCollection"}

// Get takes name of the mongoCollection, and returns the corresponding mongoCollection object, and an error if there is any.
func (c *FakeMongoCollections) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MongoCollection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mongocollectionsResource, c.ns, name), &v1alpha1.MongoCollection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MongoCollection), err
}

// List takes label and field selectors, and returns the list of MongoCollections that match those selectors.
func (c *FakeMongoCollections) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MongoCollectionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mongocollectionsResource, mongocollectionsKind, c.ns, opts), &v1alpha1.MongoCollectionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MongoCollectionList{ListMeta: obj.(*v1alpha1.MongoCollectionList).ListMeta}
	for _, item := range obj.(*v1alpha1.MongoCollectionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mongoCollections.
func (c *FakeMongoCollections) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mongocollectionsResource, c.ns, opts))

}

// Create takes the representation of a mongoCollection and creates it.  Returns the server's representation of the mongoCollection, and an error, if there is any.
func (c *FakeMongoCollections) Create(ctx context.Context, mongoCollection *v1alpha1.MongoCollection, opts v1.CreateOptions) (result *v1alpha1.MongoCollection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mongocollectionsResource, c.ns, mongoCollection), &v1alpha1.MongoCollection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MongoCollection), err
}

// Update takes the representation of a mongoCollection and updates it. Returns the server's representation of the mongoCollection, and an error, if there is any.
func (c *FakeMongoCollections) Update(ctx context.Context, mongoCollection *v1alpha1.MongoCollection, opts v1.UpdateOptions) (result *v1alpha1.MongoCollection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mongocollectionsResource, c.ns, mongoCollection), &v1alpha1.MongoCollection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MongoCollection), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMongoCollections) UpdateStatus(ctx context.Context, mongoCollection *v1alpha1.MongoCollection, opts v1.UpdateOptions) (*v1alpha1.MongoCollection, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mongocollectionsResource, "status", c.ns, mongoCollection), &v1alpha1.MongoCollection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MongoCollection), err
}

// Delete takes name of the mongoCollection and deletes it. Returns an error if one occurs.
func (c *FakeMongoCollections) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mongocollectionsResource, c.ns, name), &v1alpha1.MongoCollection{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMongoCollections) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mongocollectionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MongoCollectionList{})
	return err
}

// Patch applies the patch and returns the patched mongoCollection.
func (c *FakeMongoCollections) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MongoCollection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mongocollectionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MongoCollection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MongoCollection), err
}