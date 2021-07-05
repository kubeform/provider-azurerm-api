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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/stream/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAnalyticsJobs implements AnalyticsJobInterface
type FakeAnalyticsJobs struct {
	Fake *FakeStreamV1alpha1
	ns   string
}

var analyticsjobsResource = schema.GroupVersionResource{Group: "stream.azurerm.kubeform.com", Version: "v1alpha1", Resource: "analyticsjobs"}

var analyticsjobsKind = schema.GroupVersionKind{Group: "stream.azurerm.kubeform.com", Version: "v1alpha1", Kind: "AnalyticsJob"}

// Get takes name of the analyticsJob, and returns the corresponding analyticsJob object, and an error if there is any.
func (c *FakeAnalyticsJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AnalyticsJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(analyticsjobsResource, c.ns, name), &v1alpha1.AnalyticsJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnalyticsJob), err
}

// List takes label and field selectors, and returns the list of AnalyticsJobs that match those selectors.
func (c *FakeAnalyticsJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AnalyticsJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(analyticsjobsResource, analyticsjobsKind, c.ns, opts), &v1alpha1.AnalyticsJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AnalyticsJobList{ListMeta: obj.(*v1alpha1.AnalyticsJobList).ListMeta}
	for _, item := range obj.(*v1alpha1.AnalyticsJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested analyticsJobs.
func (c *FakeAnalyticsJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(analyticsjobsResource, c.ns, opts))

}

// Create takes the representation of a analyticsJob and creates it.  Returns the server's representation of the analyticsJob, and an error, if there is any.
func (c *FakeAnalyticsJobs) Create(ctx context.Context, analyticsJob *v1alpha1.AnalyticsJob, opts v1.CreateOptions) (result *v1alpha1.AnalyticsJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(analyticsjobsResource, c.ns, analyticsJob), &v1alpha1.AnalyticsJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnalyticsJob), err
}

// Update takes the representation of a analyticsJob and updates it. Returns the server's representation of the analyticsJob, and an error, if there is any.
func (c *FakeAnalyticsJobs) Update(ctx context.Context, analyticsJob *v1alpha1.AnalyticsJob, opts v1.UpdateOptions) (result *v1alpha1.AnalyticsJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(analyticsjobsResource, c.ns, analyticsJob), &v1alpha1.AnalyticsJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnalyticsJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAnalyticsJobs) UpdateStatus(ctx context.Context, analyticsJob *v1alpha1.AnalyticsJob, opts v1.UpdateOptions) (*v1alpha1.AnalyticsJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(analyticsjobsResource, "status", c.ns, analyticsJob), &v1alpha1.AnalyticsJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnalyticsJob), err
}

// Delete takes name of the analyticsJob and deletes it. Returns an error if one occurs.
func (c *FakeAnalyticsJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(analyticsjobsResource, c.ns, name), &v1alpha1.AnalyticsJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAnalyticsJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(analyticsjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AnalyticsJobList{})
	return err
}

// Patch applies the patch and returns the patched analyticsJob.
func (c *FakeAnalyticsJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AnalyticsJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(analyticsjobsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AnalyticsJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnalyticsJob), err
}