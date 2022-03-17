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

// FakeFactoryLinkedServiceOdbcs implements FactoryLinkedServiceOdbcInterface
type FakeFactoryLinkedServiceOdbcs struct {
	Fake *FakeDataV1alpha1
	ns   string
}

var factorylinkedserviceodbcsResource = schema.GroupVersionResource{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Resource: "factorylinkedserviceodbcs"}

var factorylinkedserviceodbcsKind = schema.GroupVersionKind{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Kind: "FactoryLinkedServiceOdbc"}

// Get takes name of the factoryLinkedServiceOdbc, and returns the corresponding factoryLinkedServiceOdbc object, and an error if there is any.
func (c *FakeFactoryLinkedServiceOdbcs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FactoryLinkedServiceOdbc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(factorylinkedserviceodbcsResource, c.ns, name), &v1alpha1.FactoryLinkedServiceOdbc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceOdbc), err
}

// List takes label and field selectors, and returns the list of FactoryLinkedServiceOdbcs that match those selectors.
func (c *FakeFactoryLinkedServiceOdbcs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FactoryLinkedServiceOdbcList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(factorylinkedserviceodbcsResource, factorylinkedserviceodbcsKind, c.ns, opts), &v1alpha1.FactoryLinkedServiceOdbcList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FactoryLinkedServiceOdbcList{ListMeta: obj.(*v1alpha1.FactoryLinkedServiceOdbcList).ListMeta}
	for _, item := range obj.(*v1alpha1.FactoryLinkedServiceOdbcList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested factoryLinkedServiceOdbcs.
func (c *FakeFactoryLinkedServiceOdbcs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(factorylinkedserviceodbcsResource, c.ns, opts))

}

// Create takes the representation of a factoryLinkedServiceOdbc and creates it.  Returns the server's representation of the factoryLinkedServiceOdbc, and an error, if there is any.
func (c *FakeFactoryLinkedServiceOdbcs) Create(ctx context.Context, factoryLinkedServiceOdbc *v1alpha1.FactoryLinkedServiceOdbc, opts v1.CreateOptions) (result *v1alpha1.FactoryLinkedServiceOdbc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(factorylinkedserviceodbcsResource, c.ns, factoryLinkedServiceOdbc), &v1alpha1.FactoryLinkedServiceOdbc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceOdbc), err
}

// Update takes the representation of a factoryLinkedServiceOdbc and updates it. Returns the server's representation of the factoryLinkedServiceOdbc, and an error, if there is any.
func (c *FakeFactoryLinkedServiceOdbcs) Update(ctx context.Context, factoryLinkedServiceOdbc *v1alpha1.FactoryLinkedServiceOdbc, opts v1.UpdateOptions) (result *v1alpha1.FactoryLinkedServiceOdbc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(factorylinkedserviceodbcsResource, c.ns, factoryLinkedServiceOdbc), &v1alpha1.FactoryLinkedServiceOdbc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceOdbc), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFactoryLinkedServiceOdbcs) UpdateStatus(ctx context.Context, factoryLinkedServiceOdbc *v1alpha1.FactoryLinkedServiceOdbc, opts v1.UpdateOptions) (*v1alpha1.FactoryLinkedServiceOdbc, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(factorylinkedserviceodbcsResource, "status", c.ns, factoryLinkedServiceOdbc), &v1alpha1.FactoryLinkedServiceOdbc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceOdbc), err
}

// Delete takes name of the factoryLinkedServiceOdbc and deletes it. Returns an error if one occurs.
func (c *FakeFactoryLinkedServiceOdbcs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(factorylinkedserviceodbcsResource, c.ns, name), &v1alpha1.FactoryLinkedServiceOdbc{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFactoryLinkedServiceOdbcs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(factorylinkedserviceodbcsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FactoryLinkedServiceOdbcList{})
	return err
}

// Patch applies the patch and returns the patched factoryLinkedServiceOdbc.
func (c *FakeFactoryLinkedServiceOdbcs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FactoryLinkedServiceOdbc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(factorylinkedserviceodbcsResource, c.ns, name, pt, data, subresources...), &v1alpha1.FactoryLinkedServiceOdbc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceOdbc), err
}