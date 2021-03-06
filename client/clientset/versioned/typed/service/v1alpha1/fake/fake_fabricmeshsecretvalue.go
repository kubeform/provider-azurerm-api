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

// FakeFabricMeshSecretValues implements FabricMeshSecretValueInterface
type FakeFabricMeshSecretValues struct {
	Fake *FakeServiceV1alpha1
	ns   string
}

var fabricmeshsecretvaluesResource = schema.GroupVersionResource{Group: "service.azurerm.kubeform.com", Version: "v1alpha1", Resource: "fabricmeshsecretvalues"}

var fabricmeshsecretvaluesKind = schema.GroupVersionKind{Group: "service.azurerm.kubeform.com", Version: "v1alpha1", Kind: "FabricMeshSecretValue"}

// Get takes name of the fabricMeshSecretValue, and returns the corresponding fabricMeshSecretValue object, and an error if there is any.
func (c *FakeFabricMeshSecretValues) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FabricMeshSecretValue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(fabricmeshsecretvaluesResource, c.ns, name), &v1alpha1.FabricMeshSecretValue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricMeshSecretValue), err
}

// List takes label and field selectors, and returns the list of FabricMeshSecretValues that match those selectors.
func (c *FakeFabricMeshSecretValues) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FabricMeshSecretValueList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(fabricmeshsecretvaluesResource, fabricmeshsecretvaluesKind, c.ns, opts), &v1alpha1.FabricMeshSecretValueList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FabricMeshSecretValueList{ListMeta: obj.(*v1alpha1.FabricMeshSecretValueList).ListMeta}
	for _, item := range obj.(*v1alpha1.FabricMeshSecretValueList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fabricMeshSecretValues.
func (c *FakeFabricMeshSecretValues) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(fabricmeshsecretvaluesResource, c.ns, opts))

}

// Create takes the representation of a fabricMeshSecretValue and creates it.  Returns the server's representation of the fabricMeshSecretValue, and an error, if there is any.
func (c *FakeFabricMeshSecretValues) Create(ctx context.Context, fabricMeshSecretValue *v1alpha1.FabricMeshSecretValue, opts v1.CreateOptions) (result *v1alpha1.FabricMeshSecretValue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(fabricmeshsecretvaluesResource, c.ns, fabricMeshSecretValue), &v1alpha1.FabricMeshSecretValue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricMeshSecretValue), err
}

// Update takes the representation of a fabricMeshSecretValue and updates it. Returns the server's representation of the fabricMeshSecretValue, and an error, if there is any.
func (c *FakeFabricMeshSecretValues) Update(ctx context.Context, fabricMeshSecretValue *v1alpha1.FabricMeshSecretValue, opts v1.UpdateOptions) (result *v1alpha1.FabricMeshSecretValue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(fabricmeshsecretvaluesResource, c.ns, fabricMeshSecretValue), &v1alpha1.FabricMeshSecretValue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricMeshSecretValue), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFabricMeshSecretValues) UpdateStatus(ctx context.Context, fabricMeshSecretValue *v1alpha1.FabricMeshSecretValue, opts v1.UpdateOptions) (*v1alpha1.FabricMeshSecretValue, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(fabricmeshsecretvaluesResource, "status", c.ns, fabricMeshSecretValue), &v1alpha1.FabricMeshSecretValue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricMeshSecretValue), err
}

// Delete takes name of the fabricMeshSecretValue and deletes it. Returns an error if one occurs.
func (c *FakeFabricMeshSecretValues) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(fabricmeshsecretvaluesResource, c.ns, name), &v1alpha1.FabricMeshSecretValue{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFabricMeshSecretValues) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(fabricmeshsecretvaluesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FabricMeshSecretValueList{})
	return err
}

// Patch applies the patch and returns the patched fabricMeshSecretValue.
func (c *FakeFabricMeshSecretValues) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FabricMeshSecretValue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fabricmeshsecretvaluesResource, c.ns, name, pt, data, subresources...), &v1alpha1.FabricMeshSecretValue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricMeshSecretValue), err
}
