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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/security/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCenterSettings implements CenterSettingInterface
type FakeCenterSettings struct {
	Fake *FakeSecurityV1alpha1
	ns   string
}

var centersettingsResource = schema.GroupVersionResource{Group: "security.azurerm.kubeform.com", Version: "v1alpha1", Resource: "centersettings"}

var centersettingsKind = schema.GroupVersionKind{Group: "security.azurerm.kubeform.com", Version: "v1alpha1", Kind: "CenterSetting"}

// Get takes name of the centerSetting, and returns the corresponding centerSetting object, and an error if there is any.
func (c *FakeCenterSettings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CenterSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(centersettingsResource, c.ns, name), &v1alpha1.CenterSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CenterSetting), err
}

// List takes label and field selectors, and returns the list of CenterSettings that match those selectors.
func (c *FakeCenterSettings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CenterSettingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(centersettingsResource, centersettingsKind, c.ns, opts), &v1alpha1.CenterSettingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CenterSettingList{ListMeta: obj.(*v1alpha1.CenterSettingList).ListMeta}
	for _, item := range obj.(*v1alpha1.CenterSettingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested centerSettings.
func (c *FakeCenterSettings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(centersettingsResource, c.ns, opts))

}

// Create takes the representation of a centerSetting and creates it.  Returns the server's representation of the centerSetting, and an error, if there is any.
func (c *FakeCenterSettings) Create(ctx context.Context, centerSetting *v1alpha1.CenterSetting, opts v1.CreateOptions) (result *v1alpha1.CenterSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(centersettingsResource, c.ns, centerSetting), &v1alpha1.CenterSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CenterSetting), err
}

// Update takes the representation of a centerSetting and updates it. Returns the server's representation of the centerSetting, and an error, if there is any.
func (c *FakeCenterSettings) Update(ctx context.Context, centerSetting *v1alpha1.CenterSetting, opts v1.UpdateOptions) (result *v1alpha1.CenterSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(centersettingsResource, c.ns, centerSetting), &v1alpha1.CenterSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CenterSetting), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCenterSettings) UpdateStatus(ctx context.Context, centerSetting *v1alpha1.CenterSetting, opts v1.UpdateOptions) (*v1alpha1.CenterSetting, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(centersettingsResource, "status", c.ns, centerSetting), &v1alpha1.CenterSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CenterSetting), err
}

// Delete takes name of the centerSetting and deletes it. Returns an error if one occurs.
func (c *FakeCenterSettings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(centersettingsResource, c.ns, name), &v1alpha1.CenterSetting{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCenterSettings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(centersettingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CenterSettingList{})
	return err
}

// Patch applies the patch and returns the patched centerSetting.
func (c *FakeCenterSettings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CenterSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(centersettingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CenterSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CenterSetting), err
}