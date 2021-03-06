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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/postgresql/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFlexibleServerDatabases implements FlexibleServerDatabaseInterface
type FakeFlexibleServerDatabases struct {
	Fake *FakePostgresqlV1alpha1
	ns   string
}

var flexibleserverdatabasesResource = schema.GroupVersionResource{Group: "postgresql.azurerm.kubeform.com", Version: "v1alpha1", Resource: "flexibleserverdatabases"}

var flexibleserverdatabasesKind = schema.GroupVersionKind{Group: "postgresql.azurerm.kubeform.com", Version: "v1alpha1", Kind: "FlexibleServerDatabase"}

// Get takes name of the flexibleServerDatabase, and returns the corresponding flexibleServerDatabase object, and an error if there is any.
func (c *FakeFlexibleServerDatabases) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FlexibleServerDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(flexibleserverdatabasesResource, c.ns, name), &v1alpha1.FlexibleServerDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FlexibleServerDatabase), err
}

// List takes label and field selectors, and returns the list of FlexibleServerDatabases that match those selectors.
func (c *FakeFlexibleServerDatabases) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FlexibleServerDatabaseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(flexibleserverdatabasesResource, flexibleserverdatabasesKind, c.ns, opts), &v1alpha1.FlexibleServerDatabaseList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FlexibleServerDatabaseList{ListMeta: obj.(*v1alpha1.FlexibleServerDatabaseList).ListMeta}
	for _, item := range obj.(*v1alpha1.FlexibleServerDatabaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested flexibleServerDatabases.
func (c *FakeFlexibleServerDatabases) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(flexibleserverdatabasesResource, c.ns, opts))

}

// Create takes the representation of a flexibleServerDatabase and creates it.  Returns the server's representation of the flexibleServerDatabase, and an error, if there is any.
func (c *FakeFlexibleServerDatabases) Create(ctx context.Context, flexibleServerDatabase *v1alpha1.FlexibleServerDatabase, opts v1.CreateOptions) (result *v1alpha1.FlexibleServerDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(flexibleserverdatabasesResource, c.ns, flexibleServerDatabase), &v1alpha1.FlexibleServerDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FlexibleServerDatabase), err
}

// Update takes the representation of a flexibleServerDatabase and updates it. Returns the server's representation of the flexibleServerDatabase, and an error, if there is any.
func (c *FakeFlexibleServerDatabases) Update(ctx context.Context, flexibleServerDatabase *v1alpha1.FlexibleServerDatabase, opts v1.UpdateOptions) (result *v1alpha1.FlexibleServerDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(flexibleserverdatabasesResource, c.ns, flexibleServerDatabase), &v1alpha1.FlexibleServerDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FlexibleServerDatabase), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFlexibleServerDatabases) UpdateStatus(ctx context.Context, flexibleServerDatabase *v1alpha1.FlexibleServerDatabase, opts v1.UpdateOptions) (*v1alpha1.FlexibleServerDatabase, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(flexibleserverdatabasesResource, "status", c.ns, flexibleServerDatabase), &v1alpha1.FlexibleServerDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FlexibleServerDatabase), err
}

// Delete takes name of the flexibleServerDatabase and deletes it. Returns an error if one occurs.
func (c *FakeFlexibleServerDatabases) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(flexibleserverdatabasesResource, c.ns, name), &v1alpha1.FlexibleServerDatabase{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFlexibleServerDatabases) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(flexibleserverdatabasesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FlexibleServerDatabaseList{})
	return err
}

// Patch applies the patch and returns the patched flexibleServerDatabase.
func (c *FakeFlexibleServerDatabases) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FlexibleServerDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(flexibleserverdatabasesResource, c.ns, name, pt, data, subresources...), &v1alpha1.FlexibleServerDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FlexibleServerDatabase), err
}
