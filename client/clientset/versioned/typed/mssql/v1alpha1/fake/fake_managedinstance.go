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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/mssql/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeManagedInstances implements ManagedInstanceInterface
type FakeManagedInstances struct {
	Fake *FakeMssqlV1alpha1
	ns   string
}

var managedinstancesResource = schema.GroupVersionResource{Group: "mssql.azurerm.kubeform.com", Version: "v1alpha1", Resource: "managedinstances"}

var managedinstancesKind = schema.GroupVersionKind{Group: "mssql.azurerm.kubeform.com", Version: "v1alpha1", Kind: "ManagedInstance"}

// Get takes name of the managedInstance, and returns the corresponding managedInstance object, and an error if there is any.
func (c *FakeManagedInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ManagedInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(managedinstancesResource, c.ns, name), &v1alpha1.ManagedInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedInstance), err
}

// List takes label and field selectors, and returns the list of ManagedInstances that match those selectors.
func (c *FakeManagedInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ManagedInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(managedinstancesResource, managedinstancesKind, c.ns, opts), &v1alpha1.ManagedInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ManagedInstanceList{ListMeta: obj.(*v1alpha1.ManagedInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.ManagedInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managedInstances.
func (c *FakeManagedInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(managedinstancesResource, c.ns, opts))

}

// Create takes the representation of a managedInstance and creates it.  Returns the server's representation of the managedInstance, and an error, if there is any.
func (c *FakeManagedInstances) Create(ctx context.Context, managedInstance *v1alpha1.ManagedInstance, opts v1.CreateOptions) (result *v1alpha1.ManagedInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(managedinstancesResource, c.ns, managedInstance), &v1alpha1.ManagedInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedInstance), err
}

// Update takes the representation of a managedInstance and updates it. Returns the server's representation of the managedInstance, and an error, if there is any.
func (c *FakeManagedInstances) Update(ctx context.Context, managedInstance *v1alpha1.ManagedInstance, opts v1.UpdateOptions) (result *v1alpha1.ManagedInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(managedinstancesResource, c.ns, managedInstance), &v1alpha1.ManagedInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagedInstances) UpdateStatus(ctx context.Context, managedInstance *v1alpha1.ManagedInstance, opts v1.UpdateOptions) (*v1alpha1.ManagedInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(managedinstancesResource, "status", c.ns, managedInstance), &v1alpha1.ManagedInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedInstance), err
}

// Delete takes name of the managedInstance and deletes it. Returns an error if one occurs.
func (c *FakeManagedInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(managedinstancesResource, c.ns, name), &v1alpha1.ManagedInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagedInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(managedinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ManagedInstanceList{})
	return err
}

// Patch applies the patch and returns the patched managedInstance.
func (c *FakeManagedInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagedInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(managedinstancesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ManagedInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedInstance), err
}
