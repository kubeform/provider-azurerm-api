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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/synapse/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSparkPools implements SparkPoolInterface
type FakeSparkPools struct {
	Fake *FakeSynapseV1alpha1
	ns   string
}

var sparkpoolsResource = schema.GroupVersionResource{Group: "synapse.azurerm.kubeform.com", Version: "v1alpha1", Resource: "sparkpools"}

var sparkpoolsKind = schema.GroupVersionKind{Group: "synapse.azurerm.kubeform.com", Version: "v1alpha1", Kind: "SparkPool"}

// Get takes name of the sparkPool, and returns the corresponding sparkPool object, and an error if there is any.
func (c *FakeSparkPools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SparkPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sparkpoolsResource, c.ns, name), &v1alpha1.SparkPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SparkPool), err
}

// List takes label and field selectors, and returns the list of SparkPools that match those selectors.
func (c *FakeSparkPools) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SparkPoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sparkpoolsResource, sparkpoolsKind, c.ns, opts), &v1alpha1.SparkPoolList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SparkPoolList{ListMeta: obj.(*v1alpha1.SparkPoolList).ListMeta}
	for _, item := range obj.(*v1alpha1.SparkPoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sparkPools.
func (c *FakeSparkPools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sparkpoolsResource, c.ns, opts))

}

// Create takes the representation of a sparkPool and creates it.  Returns the server's representation of the sparkPool, and an error, if there is any.
func (c *FakeSparkPools) Create(ctx context.Context, sparkPool *v1alpha1.SparkPool, opts v1.CreateOptions) (result *v1alpha1.SparkPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sparkpoolsResource, c.ns, sparkPool), &v1alpha1.SparkPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SparkPool), err
}

// Update takes the representation of a sparkPool and updates it. Returns the server's representation of the sparkPool, and an error, if there is any.
func (c *FakeSparkPools) Update(ctx context.Context, sparkPool *v1alpha1.SparkPool, opts v1.UpdateOptions) (result *v1alpha1.SparkPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sparkpoolsResource, c.ns, sparkPool), &v1alpha1.SparkPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SparkPool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSparkPools) UpdateStatus(ctx context.Context, sparkPool *v1alpha1.SparkPool, opts v1.UpdateOptions) (*v1alpha1.SparkPool, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sparkpoolsResource, "status", c.ns, sparkPool), &v1alpha1.SparkPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SparkPool), err
}

// Delete takes name of the sparkPool and deletes it. Returns an error if one occurs.
func (c *FakeSparkPools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sparkpoolsResource, c.ns, name), &v1alpha1.SparkPool{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSparkPools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sparkpoolsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SparkPoolList{})
	return err
}

// Patch applies the patch and returns the patched sparkPool.
func (c *FakeSparkPools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SparkPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sparkpoolsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SparkPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SparkPool), err
}
