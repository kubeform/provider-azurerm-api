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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/apimanagement/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApiDiagnostics implements ApiDiagnosticInterface
type FakeApiDiagnostics struct {
	Fake *FakeApimanagementV1alpha1
	ns   string
}

var apidiagnosticsResource = schema.GroupVersionResource{Group: "apimanagement.azurerm.kubeform.com", Version: "v1alpha1", Resource: "apidiagnostics"}

var apidiagnosticsKind = schema.GroupVersionKind{Group: "apimanagement.azurerm.kubeform.com", Version: "v1alpha1", Kind: "ApiDiagnostic"}

// Get takes name of the apiDiagnostic, and returns the corresponding apiDiagnostic object, and an error if there is any.
func (c *FakeApiDiagnostics) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApiDiagnostic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apidiagnosticsResource, c.ns, name), &v1alpha1.ApiDiagnostic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiDiagnostic), err
}

// List takes label and field selectors, and returns the list of ApiDiagnostics that match those selectors.
func (c *FakeApiDiagnostics) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApiDiagnosticList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apidiagnosticsResource, apidiagnosticsKind, c.ns, opts), &v1alpha1.ApiDiagnosticList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiDiagnosticList{ListMeta: obj.(*v1alpha1.ApiDiagnosticList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiDiagnosticList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiDiagnostics.
func (c *FakeApiDiagnostics) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apidiagnosticsResource, c.ns, opts))

}

// Create takes the representation of a apiDiagnostic and creates it.  Returns the server's representation of the apiDiagnostic, and an error, if there is any.
func (c *FakeApiDiagnostics) Create(ctx context.Context, apiDiagnostic *v1alpha1.ApiDiagnostic, opts v1.CreateOptions) (result *v1alpha1.ApiDiagnostic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apidiagnosticsResource, c.ns, apiDiagnostic), &v1alpha1.ApiDiagnostic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiDiagnostic), err
}

// Update takes the representation of a apiDiagnostic and updates it. Returns the server's representation of the apiDiagnostic, and an error, if there is any.
func (c *FakeApiDiagnostics) Update(ctx context.Context, apiDiagnostic *v1alpha1.ApiDiagnostic, opts v1.UpdateOptions) (result *v1alpha1.ApiDiagnostic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apidiagnosticsResource, c.ns, apiDiagnostic), &v1alpha1.ApiDiagnostic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiDiagnostic), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiDiagnostics) UpdateStatus(ctx context.Context, apiDiagnostic *v1alpha1.ApiDiagnostic, opts v1.UpdateOptions) (*v1alpha1.ApiDiagnostic, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apidiagnosticsResource, "status", c.ns, apiDiagnostic), &v1alpha1.ApiDiagnostic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiDiagnostic), err
}

// Delete takes name of the apiDiagnostic and deletes it. Returns an error if one occurs.
func (c *FakeApiDiagnostics) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apidiagnosticsResource, c.ns, name), &v1alpha1.ApiDiagnostic{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiDiagnostics) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apidiagnosticsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiDiagnosticList{})
	return err
}

// Patch applies the patch and returns the patched apiDiagnostic.
func (c *FakeApiDiagnostics) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiDiagnostic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apidiagnosticsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiDiagnostic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiDiagnostic), err
}
