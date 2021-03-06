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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/loganalytics/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDatasourceWindowsEvents implements DatasourceWindowsEventInterface
type FakeDatasourceWindowsEvents struct {
	Fake *FakeLoganalyticsV1alpha1
	ns   string
}

var datasourcewindowseventsResource = schema.GroupVersionResource{Group: "loganalytics.azurerm.kubeform.com", Version: "v1alpha1", Resource: "datasourcewindowsevents"}

var datasourcewindowseventsKind = schema.GroupVersionKind{Group: "loganalytics.azurerm.kubeform.com", Version: "v1alpha1", Kind: "DatasourceWindowsEvent"}

// Get takes name of the datasourceWindowsEvent, and returns the corresponding datasourceWindowsEvent object, and an error if there is any.
func (c *FakeDatasourceWindowsEvents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DatasourceWindowsEvent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datasourcewindowseventsResource, c.ns, name), &v1alpha1.DatasourceWindowsEvent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasourceWindowsEvent), err
}

// List takes label and field selectors, and returns the list of DatasourceWindowsEvents that match those selectors.
func (c *FakeDatasourceWindowsEvents) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DatasourceWindowsEventList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datasourcewindowseventsResource, datasourcewindowseventsKind, c.ns, opts), &v1alpha1.DatasourceWindowsEventList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DatasourceWindowsEventList{ListMeta: obj.(*v1alpha1.DatasourceWindowsEventList).ListMeta}
	for _, item := range obj.(*v1alpha1.DatasourceWindowsEventList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested datasourceWindowsEvents.
func (c *FakeDatasourceWindowsEvents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datasourcewindowseventsResource, c.ns, opts))

}

// Create takes the representation of a datasourceWindowsEvent and creates it.  Returns the server's representation of the datasourceWindowsEvent, and an error, if there is any.
func (c *FakeDatasourceWindowsEvents) Create(ctx context.Context, datasourceWindowsEvent *v1alpha1.DatasourceWindowsEvent, opts v1.CreateOptions) (result *v1alpha1.DatasourceWindowsEvent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datasourcewindowseventsResource, c.ns, datasourceWindowsEvent), &v1alpha1.DatasourceWindowsEvent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasourceWindowsEvent), err
}

// Update takes the representation of a datasourceWindowsEvent and updates it. Returns the server's representation of the datasourceWindowsEvent, and an error, if there is any.
func (c *FakeDatasourceWindowsEvents) Update(ctx context.Context, datasourceWindowsEvent *v1alpha1.DatasourceWindowsEvent, opts v1.UpdateOptions) (result *v1alpha1.DatasourceWindowsEvent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datasourcewindowseventsResource, c.ns, datasourceWindowsEvent), &v1alpha1.DatasourceWindowsEvent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasourceWindowsEvent), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDatasourceWindowsEvents) UpdateStatus(ctx context.Context, datasourceWindowsEvent *v1alpha1.DatasourceWindowsEvent, opts v1.UpdateOptions) (*v1alpha1.DatasourceWindowsEvent, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datasourcewindowseventsResource, "status", c.ns, datasourceWindowsEvent), &v1alpha1.DatasourceWindowsEvent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasourceWindowsEvent), err
}

// Delete takes name of the datasourceWindowsEvent and deletes it. Returns an error if one occurs.
func (c *FakeDatasourceWindowsEvents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(datasourcewindowseventsResource, c.ns, name), &v1alpha1.DatasourceWindowsEvent{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDatasourceWindowsEvents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datasourcewindowseventsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DatasourceWindowsEventList{})
	return err
}

// Patch applies the patch and returns the patched datasourceWindowsEvent.
func (c *FakeDatasourceWindowsEvents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DatasourceWindowsEvent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datasourcewindowseventsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DatasourceWindowsEvent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasourceWindowsEvent), err
}
