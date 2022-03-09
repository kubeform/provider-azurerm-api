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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/resource/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGroupCostManagementExports implements GroupCostManagementExportInterface
type FakeGroupCostManagementExports struct {
	Fake *FakeResourceV1alpha1
	ns   string
}

var groupcostmanagementexportsResource = schema.GroupVersionResource{Group: "resource.azurerm.kubeform.com", Version: "v1alpha1", Resource: "groupcostmanagementexports"}

var groupcostmanagementexportsKind = schema.GroupVersionKind{Group: "resource.azurerm.kubeform.com", Version: "v1alpha1", Kind: "GroupCostManagementExport"}

// Get takes name of the groupCostManagementExport, and returns the corresponding groupCostManagementExport object, and an error if there is any.
func (c *FakeGroupCostManagementExports) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GroupCostManagementExport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(groupcostmanagementexportsResource, c.ns, name), &v1alpha1.GroupCostManagementExport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupCostManagementExport), err
}

// List takes label and field selectors, and returns the list of GroupCostManagementExports that match those selectors.
func (c *FakeGroupCostManagementExports) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GroupCostManagementExportList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(groupcostmanagementexportsResource, groupcostmanagementexportsKind, c.ns, opts), &v1alpha1.GroupCostManagementExportList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GroupCostManagementExportList{ListMeta: obj.(*v1alpha1.GroupCostManagementExportList).ListMeta}
	for _, item := range obj.(*v1alpha1.GroupCostManagementExportList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested groupCostManagementExports.
func (c *FakeGroupCostManagementExports) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(groupcostmanagementexportsResource, c.ns, opts))

}

// Create takes the representation of a groupCostManagementExport and creates it.  Returns the server's representation of the groupCostManagementExport, and an error, if there is any.
func (c *FakeGroupCostManagementExports) Create(ctx context.Context, groupCostManagementExport *v1alpha1.GroupCostManagementExport, opts v1.CreateOptions) (result *v1alpha1.GroupCostManagementExport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(groupcostmanagementexportsResource, c.ns, groupCostManagementExport), &v1alpha1.GroupCostManagementExport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupCostManagementExport), err
}

// Update takes the representation of a groupCostManagementExport and updates it. Returns the server's representation of the groupCostManagementExport, and an error, if there is any.
func (c *FakeGroupCostManagementExports) Update(ctx context.Context, groupCostManagementExport *v1alpha1.GroupCostManagementExport, opts v1.UpdateOptions) (result *v1alpha1.GroupCostManagementExport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(groupcostmanagementexportsResource, c.ns, groupCostManagementExport), &v1alpha1.GroupCostManagementExport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupCostManagementExport), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGroupCostManagementExports) UpdateStatus(ctx context.Context, groupCostManagementExport *v1alpha1.GroupCostManagementExport, opts v1.UpdateOptions) (*v1alpha1.GroupCostManagementExport, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(groupcostmanagementexportsResource, "status", c.ns, groupCostManagementExport), &v1alpha1.GroupCostManagementExport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupCostManagementExport), err
}

// Delete takes name of the groupCostManagementExport and deletes it. Returns an error if one occurs.
func (c *FakeGroupCostManagementExports) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(groupcostmanagementexportsResource, c.ns, name), &v1alpha1.GroupCostManagementExport{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGroupCostManagementExports) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(groupcostmanagementexportsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.GroupCostManagementExportList{})
	return err
}

// Patch applies the patch and returns the patched groupCostManagementExport.
func (c *FakeGroupCostManagementExports) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GroupCostManagementExport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(groupcostmanagementexportsResource, c.ns, name, pt, data, subresources...), &v1alpha1.GroupCostManagementExport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupCostManagementExport), err
}
