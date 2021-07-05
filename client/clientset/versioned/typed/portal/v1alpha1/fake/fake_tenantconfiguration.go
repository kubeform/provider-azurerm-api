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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/portal/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTenantConfigurations implements TenantConfigurationInterface
type FakeTenantConfigurations struct {
	Fake *FakePortalV1alpha1
	ns   string
}

var tenantconfigurationsResource = schema.GroupVersionResource{Group: "portal.azurerm.kubeform.com", Version: "v1alpha1", Resource: "tenantconfigurations"}

var tenantconfigurationsKind = schema.GroupVersionKind{Group: "portal.azurerm.kubeform.com", Version: "v1alpha1", Kind: "TenantConfiguration"}

// Get takes name of the tenantConfiguration, and returns the corresponding tenantConfiguration object, and an error if there is any.
func (c *FakeTenantConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TenantConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tenantconfigurationsResource, c.ns, name), &v1alpha1.TenantConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TenantConfiguration), err
}

// List takes label and field selectors, and returns the list of TenantConfigurations that match those selectors.
func (c *FakeTenantConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TenantConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tenantconfigurationsResource, tenantconfigurationsKind, c.ns, opts), &v1alpha1.TenantConfigurationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TenantConfigurationList{ListMeta: obj.(*v1alpha1.TenantConfigurationList).ListMeta}
	for _, item := range obj.(*v1alpha1.TenantConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tenantConfigurations.
func (c *FakeTenantConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tenantconfigurationsResource, c.ns, opts))

}

// Create takes the representation of a tenantConfiguration and creates it.  Returns the server's representation of the tenantConfiguration, and an error, if there is any.
func (c *FakeTenantConfigurations) Create(ctx context.Context, tenantConfiguration *v1alpha1.TenantConfiguration, opts v1.CreateOptions) (result *v1alpha1.TenantConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tenantconfigurationsResource, c.ns, tenantConfiguration), &v1alpha1.TenantConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TenantConfiguration), err
}

// Update takes the representation of a tenantConfiguration and updates it. Returns the server's representation of the tenantConfiguration, and an error, if there is any.
func (c *FakeTenantConfigurations) Update(ctx context.Context, tenantConfiguration *v1alpha1.TenantConfiguration, opts v1.UpdateOptions) (result *v1alpha1.TenantConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tenantconfigurationsResource, c.ns, tenantConfiguration), &v1alpha1.TenantConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TenantConfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTenantConfigurations) UpdateStatus(ctx context.Context, tenantConfiguration *v1alpha1.TenantConfiguration, opts v1.UpdateOptions) (*v1alpha1.TenantConfiguration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(tenantconfigurationsResource, "status", c.ns, tenantConfiguration), &v1alpha1.TenantConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TenantConfiguration), err
}

// Delete takes name of the tenantConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeTenantConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tenantconfigurationsResource, c.ns, name), &v1alpha1.TenantConfiguration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTenantConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tenantconfigurationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TenantConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched tenantConfiguration.
func (c *FakeTenantConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TenantConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tenantconfigurationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.TenantConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TenantConfiguration), err
}