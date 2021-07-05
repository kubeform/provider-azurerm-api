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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/iottime/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSeriesInsightsGen2Environments implements SeriesInsightsGen2EnvironmentInterface
type FakeSeriesInsightsGen2Environments struct {
	Fake *FakeIottimeV1alpha1
	ns   string
}

var seriesinsightsgen2environmentsResource = schema.GroupVersionResource{Group: "iottime.azurerm.kubeform.com", Version: "v1alpha1", Resource: "seriesinsightsgen2environments"}

var seriesinsightsgen2environmentsKind = schema.GroupVersionKind{Group: "iottime.azurerm.kubeform.com", Version: "v1alpha1", Kind: "SeriesInsightsGen2Environment"}

// Get takes name of the seriesInsightsGen2Environment, and returns the corresponding seriesInsightsGen2Environment object, and an error if there is any.
func (c *FakeSeriesInsightsGen2Environments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SeriesInsightsGen2Environment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(seriesinsightsgen2environmentsResource, c.ns, name), &v1alpha1.SeriesInsightsGen2Environment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SeriesInsightsGen2Environment), err
}

// List takes label and field selectors, and returns the list of SeriesInsightsGen2Environments that match those selectors.
func (c *FakeSeriesInsightsGen2Environments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SeriesInsightsGen2EnvironmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(seriesinsightsgen2environmentsResource, seriesinsightsgen2environmentsKind, c.ns, opts), &v1alpha1.SeriesInsightsGen2EnvironmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SeriesInsightsGen2EnvironmentList{ListMeta: obj.(*v1alpha1.SeriesInsightsGen2EnvironmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.SeriesInsightsGen2EnvironmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested seriesInsightsGen2Environments.
func (c *FakeSeriesInsightsGen2Environments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(seriesinsightsgen2environmentsResource, c.ns, opts))

}

// Create takes the representation of a seriesInsightsGen2Environment and creates it.  Returns the server's representation of the seriesInsightsGen2Environment, and an error, if there is any.
func (c *FakeSeriesInsightsGen2Environments) Create(ctx context.Context, seriesInsightsGen2Environment *v1alpha1.SeriesInsightsGen2Environment, opts v1.CreateOptions) (result *v1alpha1.SeriesInsightsGen2Environment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(seriesinsightsgen2environmentsResource, c.ns, seriesInsightsGen2Environment), &v1alpha1.SeriesInsightsGen2Environment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SeriesInsightsGen2Environment), err
}

// Update takes the representation of a seriesInsightsGen2Environment and updates it. Returns the server's representation of the seriesInsightsGen2Environment, and an error, if there is any.
func (c *FakeSeriesInsightsGen2Environments) Update(ctx context.Context, seriesInsightsGen2Environment *v1alpha1.SeriesInsightsGen2Environment, opts v1.UpdateOptions) (result *v1alpha1.SeriesInsightsGen2Environment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(seriesinsightsgen2environmentsResource, c.ns, seriesInsightsGen2Environment), &v1alpha1.SeriesInsightsGen2Environment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SeriesInsightsGen2Environment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSeriesInsightsGen2Environments) UpdateStatus(ctx context.Context, seriesInsightsGen2Environment *v1alpha1.SeriesInsightsGen2Environment, opts v1.UpdateOptions) (*v1alpha1.SeriesInsightsGen2Environment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(seriesinsightsgen2environmentsResource, "status", c.ns, seriesInsightsGen2Environment), &v1alpha1.SeriesInsightsGen2Environment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SeriesInsightsGen2Environment), err
}

// Delete takes name of the seriesInsightsGen2Environment and deletes it. Returns an error if one occurs.
func (c *FakeSeriesInsightsGen2Environments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(seriesinsightsgen2environmentsResource, c.ns, name), &v1alpha1.SeriesInsightsGen2Environment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSeriesInsightsGen2Environments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(seriesinsightsgen2environmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SeriesInsightsGen2EnvironmentList{})
	return err
}

// Patch applies the patch and returns the patched seriesInsightsGen2Environment.
func (c *FakeSeriesInsightsGen2Environments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SeriesInsightsGen2Environment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(seriesinsightsgen2environmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SeriesInsightsGen2Environment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SeriesInsightsGen2Environment), err
}