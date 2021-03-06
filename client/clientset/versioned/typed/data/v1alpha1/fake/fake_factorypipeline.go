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

// FakeFactoryPipelines implements FactoryPipelineInterface
type FakeFactoryPipelines struct {
	Fake *FakeDataV1alpha1
	ns   string
}

var factorypipelinesResource = schema.GroupVersionResource{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Resource: "factorypipelines"}

var factorypipelinesKind = schema.GroupVersionKind{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Kind: "FactoryPipeline"}

// Get takes name of the factoryPipeline, and returns the corresponding factoryPipeline object, and an error if there is any.
func (c *FakeFactoryPipelines) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FactoryPipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(factorypipelinesResource, c.ns, name), &v1alpha1.FactoryPipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryPipeline), err
}

// List takes label and field selectors, and returns the list of FactoryPipelines that match those selectors.
func (c *FakeFactoryPipelines) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FactoryPipelineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(factorypipelinesResource, factorypipelinesKind, c.ns, opts), &v1alpha1.FactoryPipelineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FactoryPipelineList{ListMeta: obj.(*v1alpha1.FactoryPipelineList).ListMeta}
	for _, item := range obj.(*v1alpha1.FactoryPipelineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested factoryPipelines.
func (c *FakeFactoryPipelines) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(factorypipelinesResource, c.ns, opts))

}

// Create takes the representation of a factoryPipeline and creates it.  Returns the server's representation of the factoryPipeline, and an error, if there is any.
func (c *FakeFactoryPipelines) Create(ctx context.Context, factoryPipeline *v1alpha1.FactoryPipeline, opts v1.CreateOptions) (result *v1alpha1.FactoryPipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(factorypipelinesResource, c.ns, factoryPipeline), &v1alpha1.FactoryPipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryPipeline), err
}

// Update takes the representation of a factoryPipeline and updates it. Returns the server's representation of the factoryPipeline, and an error, if there is any.
func (c *FakeFactoryPipelines) Update(ctx context.Context, factoryPipeline *v1alpha1.FactoryPipeline, opts v1.UpdateOptions) (result *v1alpha1.FactoryPipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(factorypipelinesResource, c.ns, factoryPipeline), &v1alpha1.FactoryPipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryPipeline), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFactoryPipelines) UpdateStatus(ctx context.Context, factoryPipeline *v1alpha1.FactoryPipeline, opts v1.UpdateOptions) (*v1alpha1.FactoryPipeline, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(factorypipelinesResource, "status", c.ns, factoryPipeline), &v1alpha1.FactoryPipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryPipeline), err
}

// Delete takes name of the factoryPipeline and deletes it. Returns an error if one occurs.
func (c *FakeFactoryPipelines) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(factorypipelinesResource, c.ns, name), &v1alpha1.FactoryPipeline{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFactoryPipelines) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(factorypipelinesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FactoryPipelineList{})
	return err
}

// Patch applies the patch and returns the patched factoryPipeline.
func (c *FakeFactoryPipelines) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FactoryPipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(factorypipelinesResource, c.ns, name, pt, data, subresources...), &v1alpha1.FactoryPipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryPipeline), err
}
