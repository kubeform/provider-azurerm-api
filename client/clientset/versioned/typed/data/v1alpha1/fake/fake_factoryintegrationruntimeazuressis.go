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

// FakeFactoryIntegrationRuntimeAzureSsises implements FactoryIntegrationRuntimeAzureSsisInterface
type FakeFactoryIntegrationRuntimeAzureSsises struct {
	Fake *FakeDataV1alpha1
	ns   string
}

var factoryintegrationruntimeazuressisesResource = schema.GroupVersionResource{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Resource: "factoryintegrationruntimeazuressises"}

var factoryintegrationruntimeazuressisesKind = schema.GroupVersionKind{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Kind: "FactoryIntegrationRuntimeAzureSsis"}

// Get takes name of the factoryIntegrationRuntimeAzureSsis, and returns the corresponding factoryIntegrationRuntimeAzureSsis object, and an error if there is any.
func (c *FakeFactoryIntegrationRuntimeAzureSsises) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FactoryIntegrationRuntimeAzureSsis, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(factoryintegrationruntimeazuressisesResource, c.ns, name), &v1alpha1.FactoryIntegrationRuntimeAzureSsis{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryIntegrationRuntimeAzureSsis), err
}

// List takes label and field selectors, and returns the list of FactoryIntegrationRuntimeAzureSsises that match those selectors.
func (c *FakeFactoryIntegrationRuntimeAzureSsises) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FactoryIntegrationRuntimeAzureSsisList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(factoryintegrationruntimeazuressisesResource, factoryintegrationruntimeazuressisesKind, c.ns, opts), &v1alpha1.FactoryIntegrationRuntimeAzureSsisList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FactoryIntegrationRuntimeAzureSsisList{ListMeta: obj.(*v1alpha1.FactoryIntegrationRuntimeAzureSsisList).ListMeta}
	for _, item := range obj.(*v1alpha1.FactoryIntegrationRuntimeAzureSsisList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested factoryIntegrationRuntimeAzureSsises.
func (c *FakeFactoryIntegrationRuntimeAzureSsises) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(factoryintegrationruntimeazuressisesResource, c.ns, opts))

}

// Create takes the representation of a factoryIntegrationRuntimeAzureSsis and creates it.  Returns the server's representation of the factoryIntegrationRuntimeAzureSsis, and an error, if there is any.
func (c *FakeFactoryIntegrationRuntimeAzureSsises) Create(ctx context.Context, factoryIntegrationRuntimeAzureSsis *v1alpha1.FactoryIntegrationRuntimeAzureSsis, opts v1.CreateOptions) (result *v1alpha1.FactoryIntegrationRuntimeAzureSsis, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(factoryintegrationruntimeazuressisesResource, c.ns, factoryIntegrationRuntimeAzureSsis), &v1alpha1.FactoryIntegrationRuntimeAzureSsis{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryIntegrationRuntimeAzureSsis), err
}

// Update takes the representation of a factoryIntegrationRuntimeAzureSsis and updates it. Returns the server's representation of the factoryIntegrationRuntimeAzureSsis, and an error, if there is any.
func (c *FakeFactoryIntegrationRuntimeAzureSsises) Update(ctx context.Context, factoryIntegrationRuntimeAzureSsis *v1alpha1.FactoryIntegrationRuntimeAzureSsis, opts v1.UpdateOptions) (result *v1alpha1.FactoryIntegrationRuntimeAzureSsis, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(factoryintegrationruntimeazuressisesResource, c.ns, factoryIntegrationRuntimeAzureSsis), &v1alpha1.FactoryIntegrationRuntimeAzureSsis{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryIntegrationRuntimeAzureSsis), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFactoryIntegrationRuntimeAzureSsises) UpdateStatus(ctx context.Context, factoryIntegrationRuntimeAzureSsis *v1alpha1.FactoryIntegrationRuntimeAzureSsis, opts v1.UpdateOptions) (*v1alpha1.FactoryIntegrationRuntimeAzureSsis, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(factoryintegrationruntimeazuressisesResource, "status", c.ns, factoryIntegrationRuntimeAzureSsis), &v1alpha1.FactoryIntegrationRuntimeAzureSsis{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryIntegrationRuntimeAzureSsis), err
}

// Delete takes name of the factoryIntegrationRuntimeAzureSsis and deletes it. Returns an error if one occurs.
func (c *FakeFactoryIntegrationRuntimeAzureSsises) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(factoryintegrationruntimeazuressisesResource, c.ns, name), &v1alpha1.FactoryIntegrationRuntimeAzureSsis{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFactoryIntegrationRuntimeAzureSsises) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(factoryintegrationruntimeazuressisesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FactoryIntegrationRuntimeAzureSsisList{})
	return err
}

// Patch applies the patch and returns the patched factoryIntegrationRuntimeAzureSsis.
func (c *FakeFactoryIntegrationRuntimeAzureSsises) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FactoryIntegrationRuntimeAzureSsis, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(factoryintegrationruntimeazuressisesResource, c.ns, name, pt, data, subresources...), &v1alpha1.FactoryIntegrationRuntimeAzureSsis{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryIntegrationRuntimeAzureSsis), err
}
