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

// FakeMongoDatabases implements MongoDatabaseInterface
type FakeMongoDatabases struct {
	Fake *FakeCosmosdbV1alpha1
	ns   string
}

var mongodatabasesResource = schema.GroupVersionResource{Group: "cosmosdb.azurerm.kubeform.com", Version: "v1alpha1", Resource: "mongodatabases"}

var mongodatabasesKind = schema.GroupVersionKind{Group: "cosmosdb.azurerm.kubeform.com", Version: "v1alpha1", Kind: "MongoDatabase"}

// Get takes name of the mongoDatabase, and returns the corresponding mongoDatabase object, and an error if there is any.
func (c *FakeMongoDatabases) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MongoDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mongodatabasesResource, c.ns, name), &v1alpha1.MongoDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MongoDatabase), err
}

// List takes label and field selectors, and returns the list of MongoDatabases that match those selectors.
func (c *FakeMongoDatabases) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MongoDatabaseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mongodatabasesResource, mongodatabasesKind, c.ns, opts), &v1alpha1.MongoDatabaseList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MongoDatabaseList{ListMeta: obj.(*v1alpha1.MongoDatabaseList).ListMeta}
	for _, item := range obj.(*v1alpha1.MongoDatabaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mongoDatabases.
func (c *FakeMongoDatabases) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mongodatabasesResource, c.ns, opts))

}

// Create takes the representation of a mongoDatabase and creates it.  Returns the server's representation of the mongoDatabase, and an error, if there is any.
func (c *FakeMongoDatabases) Create(ctx context.Context, mongoDatabase *v1alpha1.MongoDatabase, opts v1.CreateOptions) (result *v1alpha1.MongoDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mongodatabasesResource, c.ns, mongoDatabase), &v1alpha1.MongoDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MongoDatabase), err
}

// Update takes the representation of a mongoDatabase and updates it. Returns the server's representation of the mongoDatabase, and an error, if there is any.
func (c *FakeMongoDatabases) Update(ctx context.Context, mongoDatabase *v1alpha1.MongoDatabase, opts v1.UpdateOptions) (result *v1alpha1.MongoDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mongodatabasesResource, c.ns, mongoDatabase), &v1alpha1.MongoDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MongoDatabase), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMongoDatabases) UpdateStatus(ctx context.Context, mongoDatabase *v1alpha1.MongoDatabase, opts v1.UpdateOptions) (*v1alpha1.MongoDatabase, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mongodatabasesResource, "status", c.ns, mongoDatabase), &v1alpha1.MongoDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MongoDatabase), err
}

// Delete takes name of the mongoDatabase and deletes it. Returns an error if one occurs.
func (c *FakeMongoDatabases) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mongodatabasesResource, c.ns, name), &v1alpha1.MongoDatabase{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMongoDatabases) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mongodatabasesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MongoDatabaseList{})
	return err
}

// Patch applies the patch and returns the patched mongoDatabase.
func (c *FakeMongoDatabases) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MongoDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mongodatabasesResource, c.ns, name, pt, data, subresources...), &v1alpha1.MongoDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MongoDatabase), err
}