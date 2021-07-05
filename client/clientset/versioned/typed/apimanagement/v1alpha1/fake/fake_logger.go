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

// FakeLoggers implements LoggerInterface
type FakeLoggers struct {
	Fake *FakeApimanagementV1alpha1
	ns   string
}

var loggersResource = schema.GroupVersionResource{Group: "apimanagement.azurerm.kubeform.com", Version: "v1alpha1", Resource: "loggers"}

var loggersKind = schema.GroupVersionKind{Group: "apimanagement.azurerm.kubeform.com", Version: "v1alpha1", Kind: "Logger"}

// Get takes name of the logger, and returns the corresponding logger object, and an error if there is any.
func (c *FakeLoggers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Logger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(loggersResource, c.ns, name), &v1alpha1.Logger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Logger), err
}

// List takes label and field selectors, and returns the list of Loggers that match those selectors.
func (c *FakeLoggers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LoggerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(loggersResource, loggersKind, c.ns, opts), &v1alpha1.LoggerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LoggerList{ListMeta: obj.(*v1alpha1.LoggerList).ListMeta}
	for _, item := range obj.(*v1alpha1.LoggerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested loggers.
func (c *FakeLoggers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(loggersResource, c.ns, opts))

}

// Create takes the representation of a logger and creates it.  Returns the server's representation of the logger, and an error, if there is any.
func (c *FakeLoggers) Create(ctx context.Context, logger *v1alpha1.Logger, opts v1.CreateOptions) (result *v1alpha1.Logger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(loggersResource, c.ns, logger), &v1alpha1.Logger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Logger), err
}

// Update takes the representation of a logger and updates it. Returns the server's representation of the logger, and an error, if there is any.
func (c *FakeLoggers) Update(ctx context.Context, logger *v1alpha1.Logger, opts v1.UpdateOptions) (result *v1alpha1.Logger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(loggersResource, c.ns, logger), &v1alpha1.Logger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Logger), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLoggers) UpdateStatus(ctx context.Context, logger *v1alpha1.Logger, opts v1.UpdateOptions) (*v1alpha1.Logger, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(loggersResource, "status", c.ns, logger), &v1alpha1.Logger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Logger), err
}

// Delete takes name of the logger and deletes it. Returns an error if one occurs.
func (c *FakeLoggers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(loggersResource, c.ns, name), &v1alpha1.Logger{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLoggers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(loggersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LoggerList{})
	return err
}

// Patch applies the patch and returns the patched logger.
func (c *FakeLoggers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Logger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(loggersResource, c.ns, name, pt, data, subresources...), &v1alpha1.Logger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Logger), err
}