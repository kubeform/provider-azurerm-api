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

// FakeFactoryLinkedServiceAzureDatabrickses implements FactoryLinkedServiceAzureDatabricksInterface
type FakeFactoryLinkedServiceAzureDatabrickses struct {
	Fake *FakeDataV1alpha1
	ns   string
}

var factorylinkedserviceazuredatabricksesResource = schema.GroupVersionResource{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Resource: "factorylinkedserviceazuredatabrickses"}

var factorylinkedserviceazuredatabricksesKind = schema.GroupVersionKind{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Kind: "FactoryLinkedServiceAzureDatabricks"}

// Get takes name of the factoryLinkedServiceAzureDatabricks, and returns the corresponding factoryLinkedServiceAzureDatabricks object, and an error if there is any.
func (c *FakeFactoryLinkedServiceAzureDatabrickses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FactoryLinkedServiceAzureDatabricks, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(factorylinkedserviceazuredatabricksesResource, c.ns, name), &v1alpha1.FactoryLinkedServiceAzureDatabricks{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceAzureDatabricks), err
}

// List takes label and field selectors, and returns the list of FactoryLinkedServiceAzureDatabrickses that match those selectors.
func (c *FakeFactoryLinkedServiceAzureDatabrickses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FactoryLinkedServiceAzureDatabricksList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(factorylinkedserviceazuredatabricksesResource, factorylinkedserviceazuredatabricksesKind, c.ns, opts), &v1alpha1.FactoryLinkedServiceAzureDatabricksList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FactoryLinkedServiceAzureDatabricksList{ListMeta: obj.(*v1alpha1.FactoryLinkedServiceAzureDatabricksList).ListMeta}
	for _, item := range obj.(*v1alpha1.FactoryLinkedServiceAzureDatabricksList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested factoryLinkedServiceAzureDatabrickses.
func (c *FakeFactoryLinkedServiceAzureDatabrickses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(factorylinkedserviceazuredatabricksesResource, c.ns, opts))

}

// Create takes the representation of a factoryLinkedServiceAzureDatabricks and creates it.  Returns the server's representation of the factoryLinkedServiceAzureDatabricks, and an error, if there is any.
func (c *FakeFactoryLinkedServiceAzureDatabrickses) Create(ctx context.Context, factoryLinkedServiceAzureDatabricks *v1alpha1.FactoryLinkedServiceAzureDatabricks, opts v1.CreateOptions) (result *v1alpha1.FactoryLinkedServiceAzureDatabricks, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(factorylinkedserviceazuredatabricksesResource, c.ns, factoryLinkedServiceAzureDatabricks), &v1alpha1.FactoryLinkedServiceAzureDatabricks{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceAzureDatabricks), err
}

// Update takes the representation of a factoryLinkedServiceAzureDatabricks and updates it. Returns the server's representation of the factoryLinkedServiceAzureDatabricks, and an error, if there is any.
func (c *FakeFactoryLinkedServiceAzureDatabrickses) Update(ctx context.Context, factoryLinkedServiceAzureDatabricks *v1alpha1.FactoryLinkedServiceAzureDatabricks, opts v1.UpdateOptions) (result *v1alpha1.FactoryLinkedServiceAzureDatabricks, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(factorylinkedserviceazuredatabricksesResource, c.ns, factoryLinkedServiceAzureDatabricks), &v1alpha1.FactoryLinkedServiceAzureDatabricks{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceAzureDatabricks), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFactoryLinkedServiceAzureDatabrickses) UpdateStatus(ctx context.Context, factoryLinkedServiceAzureDatabricks *v1alpha1.FactoryLinkedServiceAzureDatabricks, opts v1.UpdateOptions) (*v1alpha1.FactoryLinkedServiceAzureDatabricks, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(factorylinkedserviceazuredatabricksesResource, "status", c.ns, factoryLinkedServiceAzureDatabricks), &v1alpha1.FactoryLinkedServiceAzureDatabricks{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceAzureDatabricks), err
}

// Delete takes name of the factoryLinkedServiceAzureDatabricks and deletes it. Returns an error if one occurs.
func (c *FakeFactoryLinkedServiceAzureDatabrickses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(factorylinkedserviceazuredatabricksesResource, c.ns, name), &v1alpha1.FactoryLinkedServiceAzureDatabricks{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFactoryLinkedServiceAzureDatabrickses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(factorylinkedserviceazuredatabricksesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FactoryLinkedServiceAzureDatabricksList{})
	return err
}

// Patch applies the patch and returns the patched factoryLinkedServiceAzureDatabricks.
func (c *FakeFactoryLinkedServiceAzureDatabrickses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FactoryLinkedServiceAzureDatabricks, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(factorylinkedserviceazuredatabricksesResource, c.ns, name, pt, data, subresources...), &v1alpha1.FactoryLinkedServiceAzureDatabricks{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryLinkedServiceAzureDatabricks), err
}
