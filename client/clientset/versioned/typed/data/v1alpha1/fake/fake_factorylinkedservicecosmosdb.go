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

// FakeFactoryLinkedServiceCosmosdbs implements FactoryLinkedServiceCosmosdbInterface
type FakeFactoryLinkedServiceCosmosdbs struct {
	Fake *FakeDataV1alpha1
	ns   string
}

var factorylinkedservicecosmosdbsResource = schema.GroupVersionResource{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Resource: "factorylinkedservicecosmosdbs"}

var factorylinkedservicecosmosdbsKind = schema.GroupVersionKind{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Kind: "FactoryLinkedServiceCosmosdb"}

// Get takes name of the factoryLinkedServiceCosmosdb, and returns the corresponding factoryLinkedServiceCosmosdb object, and an error if there is any.
func (c *FakeFactoryLinkedServiceCosmosdbs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FactoryLinkedServiceCosmosdb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(factorylinkedservicecosmosdbsResource, c.ns, name), &v1alpha1.FactoryLinkedServiceCosmosdb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceCosmosdb), err
}

// List takes label and field selectors, and returns the list of FactoryLinkedServiceCosmosdbs that match those selectors.
func (c *FakeFactoryLinkedServiceCosmosdbs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FactoryLinkedServiceCosmosdbList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(factorylinkedservicecosmosdbsResource, factorylinkedservicecosmosdbsKind, c.ns, opts), &v1alpha1.FactoryLinkedServiceCosmosdbList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FactoryLinkedServiceCosmosdbList{ListMeta: obj.(*v1alpha1.FactoryLinkedServiceCosmosdbList).ListMeta}
	for _, item := range obj.(*v1alpha1.FactoryLinkedServiceCosmosdbList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested factoryLinkedServiceCosmosdbs.
func (c *FakeFactoryLinkedServiceCosmosdbs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(factorylinkedservicecosmosdbsResource, c.ns, opts))

}

// Create takes the representation of a factoryLinkedServiceCosmosdb and creates it.  Returns the server's representation of the factoryLinkedServiceCosmosdb, and an error, if there is any.
func (c *FakeFactoryLinkedServiceCosmosdbs) Create(ctx context.Context, factoryLinkedServiceCosmosdb *v1alpha1.FactoryLinkedServiceCosmosdb, opts v1.CreateOptions) (result *v1alpha1.FactoryLinkedServiceCosmosdb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(factorylinkedservicecosmosdbsResource, c.ns, factoryLinkedServiceCosmosdb), &v1alpha1.FactoryLinkedServiceCosmosdb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceCosmosdb), err
}

// Update takes the representation of a factoryLinkedServiceCosmosdb and updates it. Returns the server's representation of the factoryLinkedServiceCosmosdb, and an error, if there is any.
func (c *FakeFactoryLinkedServiceCosmosdbs) Update(ctx context.Context, factoryLinkedServiceCosmosdb *v1alpha1.FactoryLinkedServiceCosmosdb, opts v1.UpdateOptions) (result *v1alpha1.FactoryLinkedServiceCosmosdb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(factorylinkedservicecosmosdbsResource, c.ns, factoryLinkedServiceCosmosdb), &v1alpha1.FactoryLinkedServiceCosmosdb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceCosmosdb), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFactoryLinkedServiceCosmosdbs) UpdateStatus(ctx context.Context, factoryLinkedServiceCosmosdb *v1alpha1.FactoryLinkedServiceCosmosdb, opts v1.UpdateOptions) (*v1alpha1.FactoryLinkedServiceCosmosdb, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(factorylinkedservicecosmosdbsResource, "status", c.ns, factoryLinkedServiceCosmosdb), &v1alpha1.FactoryLinkedServiceCosmosdb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceCosmosdb), err
}

// Delete takes name of the factoryLinkedServiceCosmosdb and deletes it. Returns an error if one occurs.
func (c *FakeFactoryLinkedServiceCosmosdbs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(factorylinkedservicecosmosdbsResource, c.ns, name), &v1alpha1.FactoryLinkedServiceCosmosdb{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFactoryLinkedServiceCosmosdbs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(factorylinkedservicecosmosdbsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FactoryLinkedServiceCosmosdbList{})
	return err
}

// Patch applies the patch and returns the patched factoryLinkedServiceCosmosdb.
func (c *FakeFactoryLinkedServiceCosmosdbs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FactoryLinkedServiceCosmosdb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(factorylinkedservicecosmosdbsResource, c.ns, name, pt, data, subresources...), &v1alpha1.FactoryLinkedServiceCosmosdb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceCosmosdb), err
}
