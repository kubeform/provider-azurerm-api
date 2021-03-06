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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/monitor/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeScheduledQueryRulesLogs implements ScheduledQueryRulesLogInterface
type FakeScheduledQueryRulesLogs struct {
	Fake *FakeMonitorV1alpha1
	ns   string
}

var scheduledqueryruleslogsResource = schema.GroupVersionResource{Group: "monitor.azurerm.kubeform.com", Version: "v1alpha1", Resource: "scheduledqueryruleslogs"}

var scheduledqueryruleslogsKind = schema.GroupVersionKind{Group: "monitor.azurerm.kubeform.com", Version: "v1alpha1", Kind: "ScheduledQueryRulesLog"}

// Get takes name of the scheduledQueryRulesLog, and returns the corresponding scheduledQueryRulesLog object, and an error if there is any.
func (c *FakeScheduledQueryRulesLogs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ScheduledQueryRulesLog, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(scheduledqueryruleslogsResource, c.ns, name), &v1alpha1.ScheduledQueryRulesLog{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScheduledQueryRulesLog), err
}

// List takes label and field selectors, and returns the list of ScheduledQueryRulesLogs that match those selectors.
func (c *FakeScheduledQueryRulesLogs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ScheduledQueryRulesLogList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(scheduledqueryruleslogsResource, scheduledqueryruleslogsKind, c.ns, opts), &v1alpha1.ScheduledQueryRulesLogList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ScheduledQueryRulesLogList{ListMeta: obj.(*v1alpha1.ScheduledQueryRulesLogList).ListMeta}
	for _, item := range obj.(*v1alpha1.ScheduledQueryRulesLogList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested scheduledQueryRulesLogs.
func (c *FakeScheduledQueryRulesLogs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(scheduledqueryruleslogsResource, c.ns, opts))

}

// Create takes the representation of a scheduledQueryRulesLog and creates it.  Returns the server's representation of the scheduledQueryRulesLog, and an error, if there is any.
func (c *FakeScheduledQueryRulesLogs) Create(ctx context.Context, scheduledQueryRulesLog *v1alpha1.ScheduledQueryRulesLog, opts v1.CreateOptions) (result *v1alpha1.ScheduledQueryRulesLog, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(scheduledqueryruleslogsResource, c.ns, scheduledQueryRulesLog), &v1alpha1.ScheduledQueryRulesLog{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScheduledQueryRulesLog), err
}

// Update takes the representation of a scheduledQueryRulesLog and updates it. Returns the server's representation of the scheduledQueryRulesLog, and an error, if there is any.
func (c *FakeScheduledQueryRulesLogs) Update(ctx context.Context, scheduledQueryRulesLog *v1alpha1.ScheduledQueryRulesLog, opts v1.UpdateOptions) (result *v1alpha1.ScheduledQueryRulesLog, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(scheduledqueryruleslogsResource, c.ns, scheduledQueryRulesLog), &v1alpha1.ScheduledQueryRulesLog{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScheduledQueryRulesLog), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeScheduledQueryRulesLogs) UpdateStatus(ctx context.Context, scheduledQueryRulesLog *v1alpha1.ScheduledQueryRulesLog, opts v1.UpdateOptions) (*v1alpha1.ScheduledQueryRulesLog, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(scheduledqueryruleslogsResource, "status", c.ns, scheduledQueryRulesLog), &v1alpha1.ScheduledQueryRulesLog{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScheduledQueryRulesLog), err
}

// Delete takes name of the scheduledQueryRulesLog and deletes it. Returns an error if one occurs.
func (c *FakeScheduledQueryRulesLogs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(scheduledqueryruleslogsResource, c.ns, name), &v1alpha1.ScheduledQueryRulesLog{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeScheduledQueryRulesLogs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(scheduledqueryruleslogsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ScheduledQueryRulesLogList{})
	return err
}

// Patch applies the patch and returns the patched scheduledQueryRulesLog.
func (c *FakeScheduledQueryRulesLogs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ScheduledQueryRulesLog, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(scheduledqueryruleslogsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ScheduledQueryRulesLog{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScheduledQueryRulesLog), err
}
