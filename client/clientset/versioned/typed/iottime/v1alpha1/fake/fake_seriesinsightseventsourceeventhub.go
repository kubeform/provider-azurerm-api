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

// FakeSeriesInsightsEventSourceEventhubs implements SeriesInsightsEventSourceEventhubInterface
type FakeSeriesInsightsEventSourceEventhubs struct {
	Fake *FakeIottimeV1alpha1
	ns   string
}

var seriesinsightseventsourceeventhubsResource = schema.GroupVersionResource{Group: "iottime.azurerm.kubeform.com", Version: "v1alpha1", Resource: "seriesinsightseventsourceeventhubs"}

var seriesinsightseventsourceeventhubsKind = schema.GroupVersionKind{Group: "iottime.azurerm.kubeform.com", Version: "v1alpha1", Kind: "SeriesInsightsEventSourceEventhub"}

// Get takes name of the seriesInsightsEventSourceEventhub, and returns the corresponding seriesInsightsEventSourceEventhub object, and an error if there is any.
func (c *FakeSeriesInsightsEventSourceEventhubs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SeriesInsightsEventSourceEventhub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(seriesinsightseventsourceeventhubsResource, c.ns, name), &v1alpha1.SeriesInsightsEventSourceEventhub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SeriesInsightsEventSourceEventhub), err
}

// List takes label and field selectors, and returns the list of SeriesInsightsEventSourceEventhubs that match those selectors.
func (c *FakeSeriesInsightsEventSourceEventhubs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SeriesInsightsEventSourceEventhubList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(seriesinsightseventsourceeventhubsResource, seriesinsightseventsourceeventhubsKind, c.ns, opts), &v1alpha1.SeriesInsightsEventSourceEventhubList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SeriesInsightsEventSourceEventhubList{ListMeta: obj.(*v1alpha1.SeriesInsightsEventSourceEventhubList).ListMeta}
	for _, item := range obj.(*v1alpha1.SeriesInsightsEventSourceEventhubList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested seriesInsightsEventSourceEventhubs.
func (c *FakeSeriesInsightsEventSourceEventhubs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(seriesinsightseventsourceeventhubsResource, c.ns, opts))

}

// Create takes the representation of a seriesInsightsEventSourceEventhub and creates it.  Returns the server's representation of the seriesInsightsEventSourceEventhub, and an error, if there is any.
func (c *FakeSeriesInsightsEventSourceEventhubs) Create(ctx context.Context, seriesInsightsEventSourceEventhub *v1alpha1.SeriesInsightsEventSourceEventhub, opts v1.CreateOptions) (result *v1alpha1.SeriesInsightsEventSourceEventhub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(seriesinsightseventsourceeventhubsResource, c.ns, seriesInsightsEventSourceEventhub), &v1alpha1.SeriesInsightsEventSourceEventhub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SeriesInsightsEventSourceEventhub), err
}

// Update takes the representation of a seriesInsightsEventSourceEventhub and updates it. Returns the server's representation of the seriesInsightsEventSourceEventhub, and an error, if there is any.
func (c *FakeSeriesInsightsEventSourceEventhubs) Update(ctx context.Context, seriesInsightsEventSourceEventhub *v1alpha1.SeriesInsightsEventSourceEventhub, opts v1.UpdateOptions) (result *v1alpha1.SeriesInsightsEventSourceEventhub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(seriesinsightseventsourceeventhubsResource, c.ns, seriesInsightsEventSourceEventhub), &v1alpha1.SeriesInsightsEventSourceEventhub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SeriesInsightsEventSourceEventhub), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSeriesInsightsEventSourceEventhubs) UpdateStatus(ctx context.Context, seriesInsightsEventSourceEventhub *v1alpha1.SeriesInsightsEventSourceEventhub, opts v1.UpdateOptions) (*v1alpha1.SeriesInsightsEventSourceEventhub, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(seriesinsightseventsourceeventhubsResource, "status", c.ns, seriesInsightsEventSourceEventhub), &v1alpha1.SeriesInsightsEventSourceEventhub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SeriesInsightsEventSourceEventhub), err
}

// Delete takes name of the seriesInsightsEventSourceEventhub and deletes it. Returns an error if one occurs.
func (c *FakeSeriesInsightsEventSourceEventhubs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(seriesinsightseventsourceeventhubsResource, c.ns, name), &v1alpha1.SeriesInsightsEventSourceEventhub{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSeriesInsightsEventSourceEventhubs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(seriesinsightseventsourceeventhubsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SeriesInsightsEventSourceEventhubList{})
	return err
}

// Patch applies the patch and returns the patched seriesInsightsEventSourceEventhub.
func (c *FakeSeriesInsightsEventSourceEventhubs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SeriesInsightsEventSourceEventhub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(seriesinsightseventsourceeventhubsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SeriesInsightsEventSourceEventhub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SeriesInsightsEventSourceEventhub), err
}