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

// FakeManagedInstanceActiveDirectoryAdministrators implements ManagedInstanceActiveDirectoryAdministratorInterface
type FakeManagedInstanceActiveDirectoryAdministrators struct {
	Fake *FakeMssqlV1alpha1
	ns   string
}

var managedinstanceactivedirectoryadministratorsResource = schema.GroupVersionResource{Group: "mssql.azurerm.kubeform.com", Version: "v1alpha1", Resource: "managedinstanceactivedirectoryadministrators"}

var managedinstanceactivedirectoryadministratorsKind = schema.GroupVersionKind{Group: "mssql.azurerm.kubeform.com", Version: "v1alpha1", Kind: "ManagedInstanceActiveDirectoryAdministrator"}

// Get takes name of the managedInstanceActiveDirectoryAdministrator, and returns the corresponding managedInstanceActiveDirectoryAdministrator object, and an error if there is any.
func (c *FakeManagedInstanceActiveDirectoryAdministrators) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ManagedInstanceActiveDirectoryAdministrator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(managedinstanceactivedirectoryadministratorsResource, c.ns, name), &v1alpha1.ManagedInstanceActiveDirectoryAdministrator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedInstanceActiveDirectoryAdministrator), err
}

// List takes label and field selectors, and returns the list of ManagedInstanceActiveDirectoryAdministrators that match those selectors.
func (c *FakeManagedInstanceActiveDirectoryAdministrators) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ManagedInstanceActiveDirectoryAdministratorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(managedinstanceactivedirectoryadministratorsResource, managedinstanceactivedirectoryadministratorsKind, c.ns, opts), &v1alpha1.ManagedInstanceActiveDirectoryAdministratorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ManagedInstanceActiveDirectoryAdministratorList{ListMeta: obj.(*v1alpha1.ManagedInstanceActiveDirectoryAdministratorList).ListMeta}
	for _, item := range obj.(*v1alpha1.ManagedInstanceActiveDirectoryAdministratorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managedInstanceActiveDirectoryAdministrators.
func (c *FakeManagedInstanceActiveDirectoryAdministrators) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(managedinstanceactivedirectoryadministratorsResource, c.ns, opts))

}

// Create takes the representation of a managedInstanceActiveDirectoryAdministrator and creates it.  Returns the server's representation of the managedInstanceActiveDirectoryAdministrator, and an error, if there is any.
func (c *FakeManagedInstanceActiveDirectoryAdministrators) Create(ctx context.Context, managedInstanceActiveDirectoryAdministrator *v1alpha1.ManagedInstanceActiveDirectoryAdministrator, opts v1.CreateOptions) (result *v1alpha1.ManagedInstanceActiveDirectoryAdministrator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(managedinstanceactivedirectoryadministratorsResource, c.ns, managedInstanceActiveDirectoryAdministrator), &v1alpha1.ManagedInstanceActiveDirectoryAdministrator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedInstanceActiveDirectoryAdministrator), err
}

// Update takes the representation of a managedInstanceActiveDirectoryAdministrator and updates it. Returns the server's representation of the managedInstanceActiveDirectoryAdministrator, and an error, if there is any.
func (c *FakeManagedInstanceActiveDirectoryAdministrators) Update(ctx context.Context, managedInstanceActiveDirectoryAdministrator *v1alpha1.ManagedInstanceActiveDirectoryAdministrator, opts v1.UpdateOptions) (result *v1alpha1.ManagedInstanceActiveDirectoryAdministrator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(managedinstanceactivedirectoryadministratorsResource, c.ns, managedInstanceActiveDirectoryAdministrator), &v1alpha1.ManagedInstanceActiveDirectoryAdministrator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedInstanceActiveDirectoryAdministrator), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagedInstanceActiveDirectoryAdministrators) UpdateStatus(ctx context.Context, managedInstanceActiveDirectoryAdministrator *v1alpha1.ManagedInstanceActiveDirectoryAdministrator, opts v1.UpdateOptions) (*v1alpha1.ManagedInstanceActiveDirectoryAdministrator, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(managedinstanceactivedirectoryadministratorsResource, "status", c.ns, managedInstanceActiveDirectoryAdministrator), &v1alpha1.ManagedInstanceActiveDirectoryAdministrator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedInstanceActiveDirectoryAdministrator), err
}

// Delete takes name of the managedInstanceActiveDirectoryAdministrator and deletes it. Returns an error if one occurs.
func (c *FakeManagedInstanceActiveDirectoryAdministrators) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(managedinstanceactivedirectoryadministratorsResource, c.ns, name), &v1alpha1.ManagedInstanceActiveDirectoryAdministrator{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagedInstanceActiveDirectoryAdministrators) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(managedinstanceactivedirectoryadministratorsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ManagedInstanceActiveDirectoryAdministratorList{})
	return err
}

// Patch applies the patch and returns the patched managedInstanceActiveDirectoryAdministrator.
func (c *FakeManagedInstanceActiveDirectoryAdministrators) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagedInstanceActiveDirectoryAdministrator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(managedinstanceactivedirectoryadministratorsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ManagedInstanceActiveDirectoryAdministrator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedInstanceActiveDirectoryAdministrator), err
}
