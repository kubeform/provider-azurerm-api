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

// FakeAnalyticsFunctionJavascriptUdves implements AnalyticsFunctionJavascriptUdfInterface
type FakeAnalyticsFunctionJavascriptUdves struct {
	Fake *FakeStreamV1alpha1
	ns   string
}

var analyticsfunctionjavascriptudvesResource = schema.GroupVersionResource{Group: "stream.azurerm.kubeform.com", Version: "v1alpha1", Resource: "analyticsfunctionjavascriptudves"}

var analyticsfunctionjavascriptudvesKind = schema.GroupVersionKind{Group: "stream.azurerm.kubeform.com", Version: "v1alpha1", Kind: "AnalyticsFunctionJavascriptUdf"}

// Get takes name of the analyticsFunctionJavascriptUdf, and returns the corresponding analyticsFunctionJavascriptUdf object, and an error if there is any.
func (c *FakeAnalyticsFunctionJavascriptUdves) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AnalyticsFunctionJavascriptUdf, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(analyticsfunctionjavascriptudvesResource, c.ns, name), &v1alpha1.AnalyticsFunctionJavascriptUdf{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnalyticsFunctionJavascriptUdf), err
}

// List takes label and field selectors, and returns the list of AnalyticsFunctionJavascriptUdves that match those selectors.
func (c *FakeAnalyticsFunctionJavascriptUdves) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AnalyticsFunctionJavascriptUdfList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(analyticsfunctionjavascriptudvesResource, analyticsfunctionjavascriptudvesKind, c.ns, opts), &v1alpha1.AnalyticsFunctionJavascriptUdfList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AnalyticsFunctionJavascriptUdfList{ListMeta: obj.(*v1alpha1.AnalyticsFunctionJavascriptUdfList).ListMeta}
	for _, item := range obj.(*v1alpha1.AnalyticsFunctionJavascriptUdfList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested analyticsFunctionJavascriptUdves.
func (c *FakeAnalyticsFunctionJavascriptUdves) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(analyticsfunctionjavascriptudvesResource, c.ns, opts))

}

// Create takes the representation of a analyticsFunctionJavascriptUdf and creates it.  Returns the server's representation of the analyticsFunctionJavascriptUdf, and an error, if there is any.
func (c *FakeAnalyticsFunctionJavascriptUdves) Create(ctx context.Context, analyticsFunctionJavascriptUdf *v1alpha1.AnalyticsFunctionJavascriptUdf, opts v1.CreateOptions) (result *v1alpha1.AnalyticsFunctionJavascriptUdf, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(analyticsfunctionjavascriptudvesResource, c.ns, analyticsFunctionJavascriptUdf), &v1alpha1.AnalyticsFunctionJavascriptUdf{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnalyticsFunctionJavascriptUdf), err
}

// Update takes the representation of a analyticsFunctionJavascriptUdf and updates it. Returns the server's representation of the analyticsFunctionJavascriptUdf, and an error, if there is any.
func (c *FakeAnalyticsFunctionJavascriptUdves) Update(ctx context.Context, analyticsFunctionJavascriptUdf *v1alpha1.AnalyticsFunctionJavascriptUdf, opts v1.UpdateOptions) (result *v1alpha1.AnalyticsFunctionJavascriptUdf, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(analyticsfunctionjavascriptudvesResource, c.ns, analyticsFunctionJavascriptUdf), &v1alpha1.AnalyticsFunctionJavascriptUdf{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnalyticsFunctionJavascriptUdf), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAnalyticsFunctionJavascriptUdves) UpdateStatus(ctx context.Context, analyticsFunctionJavascriptUdf *v1alpha1.AnalyticsFunctionJavascriptUdf, opts v1.UpdateOptions) (*v1alpha1.AnalyticsFunctionJavascriptUdf, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(analyticsfunctionjavascriptudvesResource, "status", c.ns, analyticsFunctionJavascriptUdf), &v1alpha1.AnalyticsFunctionJavascriptUdf{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnalyticsFunctionJavascriptUdf), err
}

// Delete takes name of the analyticsFunctionJavascriptUdf and deletes it. Returns an error if one occurs.
func (c *FakeAnalyticsFunctionJavascriptUdves) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(analyticsfunctionjavascriptudvesResource, c.ns, name), &v1alpha1.AnalyticsFunctionJavascriptUdf{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAnalyticsFunctionJavascriptUdves) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(analyticsfunctionjavascriptudvesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AnalyticsFunctionJavascriptUdfList{})
	return err
}

// Patch applies the patch and returns the patched analyticsFunctionJavascriptUdf.
func (c *FakeAnalyticsFunctionJavascriptUdves) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AnalyticsFunctionJavascriptUdf, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(analyticsfunctionjavascriptudvesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AnalyticsFunctionJavascriptUdf{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AnalyticsFunctionJavascriptUdf), err
}
