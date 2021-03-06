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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/service/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFabricMeshApplications implements FabricMeshApplicationInterface
type FakeFabricMeshApplications struct {
	Fake *FakeServiceV1alpha1
	ns   string
}

var fabricmeshapplicationsResource = schema.GroupVersionResource{Group: "service.azurerm.kubeform.com", Version: "v1alpha1", Resource: "fabricmeshapplications"}

var fabricmeshapplicationsKind = schema.GroupVersionKind{Group: "service.azurerm.kubeform.com", Version: "v1alpha1", Kind: "FabricMeshApplication"}

// Get takes name of the fabricMeshApplication, and returns the corresponding fabricMeshApplication object, and an error if there is any.
func (c *FakeFabricMeshApplications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FabricMeshApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(fabricmeshapplicationsResource, c.ns, name), &v1alpha1.FabricMeshApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricMeshApplication), err
}

// List takes label and field selectors, and returns the list of FabricMeshApplications that match those selectors.
func (c *FakeFabricMeshApplications) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FabricMeshApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(fabricmeshapplicationsResource, fabricmeshapplicationsKind, c.ns, opts), &v1alpha1.FabricMeshApplicationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FabricMeshApplicationList{ListMeta: obj.(*v1alpha1.FabricMeshApplicationList).ListMeta}
	for _, item := range obj.(*v1alpha1.FabricMeshApplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fabricMeshApplications.
func (c *FakeFabricMeshApplications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(fabricmeshapplicationsResource, c.ns, opts))

}

// Create takes the representation of a fabricMeshApplication and creates it.  Returns the server's representation of the fabricMeshApplication, and an error, if there is any.
func (c *FakeFabricMeshApplications) Create(ctx context.Context, fabricMeshApplication *v1alpha1.FabricMeshApplication, opts v1.CreateOptions) (result *v1alpha1.FabricMeshApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(fabricmeshapplicationsResource, c.ns, fabricMeshApplication), &v1alpha1.FabricMeshApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricMeshApplication), err
}

// Update takes the representation of a fabricMeshApplication and updates it. Returns the server's representation of the fabricMeshApplication, and an error, if there is any.
func (c *FakeFabricMeshApplications) Update(ctx context.Context, fabricMeshApplication *v1alpha1.FabricMeshApplication, opts v1.UpdateOptions) (result *v1alpha1.FabricMeshApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(fabricmeshapplicationsResource, c.ns, fabricMeshApplication), &v1alpha1.FabricMeshApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricMeshApplication), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFabricMeshApplications) UpdateStatus(ctx context.Context, fabricMeshApplication *v1alpha1.FabricMeshApplication, opts v1.UpdateOptions) (*v1alpha1.FabricMeshApplication, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(fabricmeshapplicationsResource, "status", c.ns, fabricMeshApplication), &v1alpha1.FabricMeshApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricMeshApplication), err
}

// Delete takes name of the fabricMeshApplication and deletes it. Returns an error if one occurs.
func (c *FakeFabricMeshApplications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(fabricmeshapplicationsResource, c.ns, name), &v1alpha1.FabricMeshApplication{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFabricMeshApplications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(fabricmeshapplicationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FabricMeshApplicationList{})
	return err
}

// Patch applies the patch and returns the patched fabricMeshApplication.
func (c *FakeFabricMeshApplications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FabricMeshApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fabricmeshapplicationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.FabricMeshApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricMeshApplication), err
}
