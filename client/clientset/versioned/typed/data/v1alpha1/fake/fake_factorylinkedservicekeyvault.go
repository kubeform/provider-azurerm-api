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

// FakeFactoryLinkedServiceKeyVaults implements FactoryLinkedServiceKeyVaultInterface
type FakeFactoryLinkedServiceKeyVaults struct {
	Fake *FakeDataV1alpha1
	ns   string
}

var factorylinkedservicekeyvaultsResource = schema.GroupVersionResource{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Resource: "factorylinkedservicekeyvaults"}

var factorylinkedservicekeyvaultsKind = schema.GroupVersionKind{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Kind: "FactoryLinkedServiceKeyVault"}

// Get takes name of the factoryLinkedServiceKeyVault, and returns the corresponding factoryLinkedServiceKeyVault object, and an error if there is any.
func (c *FakeFactoryLinkedServiceKeyVaults) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FactoryLinkedServiceKeyVault, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(factorylinkedservicekeyvaultsResource, c.ns, name), &v1alpha1.FactoryLinkedServiceKeyVault{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceKeyVault), err
}

// List takes label and field selectors, and returns the list of FactoryLinkedServiceKeyVaults that match those selectors.
func (c *FakeFactoryLinkedServiceKeyVaults) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FactoryLinkedServiceKeyVaultList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(factorylinkedservicekeyvaultsResource, factorylinkedservicekeyvaultsKind, c.ns, opts), &v1alpha1.FactoryLinkedServiceKeyVaultList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FactoryLinkedServiceKeyVaultList{ListMeta: obj.(*v1alpha1.FactoryLinkedServiceKeyVaultList).ListMeta}
	for _, item := range obj.(*v1alpha1.FactoryLinkedServiceKeyVaultList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested factoryLinkedServiceKeyVaults.
func (c *FakeFactoryLinkedServiceKeyVaults) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(factorylinkedservicekeyvaultsResource, c.ns, opts))

}

// Create takes the representation of a factoryLinkedServiceKeyVault and creates it.  Returns the server's representation of the factoryLinkedServiceKeyVault, and an error, if there is any.
func (c *FakeFactoryLinkedServiceKeyVaults) Create(ctx context.Context, factoryLinkedServiceKeyVault *v1alpha1.FactoryLinkedServiceKeyVault, opts v1.CreateOptions) (result *v1alpha1.FactoryLinkedServiceKeyVault, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(factorylinkedservicekeyvaultsResource, c.ns, factoryLinkedServiceKeyVault), &v1alpha1.FactoryLinkedServiceKeyVault{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceKeyVault), err
}

// Update takes the representation of a factoryLinkedServiceKeyVault and updates it. Returns the server's representation of the factoryLinkedServiceKeyVault, and an error, if there is any.
func (c *FakeFactoryLinkedServiceKeyVaults) Update(ctx context.Context, factoryLinkedServiceKeyVault *v1alpha1.FactoryLinkedServiceKeyVault, opts v1.UpdateOptions) (result *v1alpha1.FactoryLinkedServiceKeyVault, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(factorylinkedservicekeyvaultsResource, c.ns, factoryLinkedServiceKeyVault), &v1alpha1.FactoryLinkedServiceKeyVault{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceKeyVault), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFactoryLinkedServiceKeyVaults) UpdateStatus(ctx context.Context, factoryLinkedServiceKeyVault *v1alpha1.FactoryLinkedServiceKeyVault, opts v1.UpdateOptions) (*v1alpha1.FactoryLinkedServiceKeyVault, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(factorylinkedservicekeyvaultsResource, "status", c.ns, factoryLinkedServiceKeyVault), &v1alpha1.FactoryLinkedServiceKeyVault{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceKeyVault), err
}

// Delete takes name of the factoryLinkedServiceKeyVault and deletes it. Returns an error if one occurs.
func (c *FakeFactoryLinkedServiceKeyVaults) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(factorylinkedservicekeyvaultsResource, c.ns, name), &v1alpha1.FactoryLinkedServiceKeyVault{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFactoryLinkedServiceKeyVaults) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(factorylinkedservicekeyvaultsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FactoryLinkedServiceKeyVaultList{})
	return err
}

// Patch applies the patch and returns the patched factoryLinkedServiceKeyVault.
func (c *FakeFactoryLinkedServiceKeyVaults) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FactoryLinkedServiceKeyVault, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(factorylinkedservicekeyvaultsResource, c.ns, name, pt, data, subresources...), &v1alpha1.FactoryLinkedServiceKeyVault{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceKeyVault), err
}
