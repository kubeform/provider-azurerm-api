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

// FakeFactoryTriggerTumblingWindows implements FactoryTriggerTumblingWindowInterface
type FakeFactoryTriggerTumblingWindows struct {
	Fake *FakeDataV1alpha1
	ns   string
}

var factorytriggertumblingwindowsResource = schema.GroupVersionResource{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Resource: "factorytriggertumblingwindows"}

var factorytriggertumblingwindowsKind = schema.GroupVersionKind{Group: "data.azurerm.kubeform.com", Version: "v1alpha1", Kind: "FactoryTriggerTumblingWindow"}

// Get takes name of the factoryTriggerTumblingWindow, and returns the corresponding factoryTriggerTumblingWindow object, and an error if there is any.
func (c *FakeFactoryTriggerTumblingWindows) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FactoryTriggerTumblingWindow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(factorytriggertumblingwindowsResource, c.ns, name), &v1alpha1.FactoryTriggerTumblingWindow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryTriggerTumblingWindow), err
}

// List takes label and field selectors, and returns the list of FactoryTriggerTumblingWindows that match those selectors.
func (c *FakeFactoryTriggerTumblingWindows) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FactoryTriggerTumblingWindowList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(factorytriggertumblingwindowsResource, factorytriggertumblingwindowsKind, c.ns, opts), &v1alpha1.FactoryTriggerTumblingWindowList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FactoryTriggerTumblingWindowList{ListMeta: obj.(*v1alpha1.FactoryTriggerTumblingWindowList).ListMeta}
	for _, item := range obj.(*v1alpha1.FactoryTriggerTumblingWindowList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested factoryTriggerTumblingWindows.
func (c *FakeFactoryTriggerTumblingWindows) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(factorytriggertumblingwindowsResource, c.ns, opts))

}

// Create takes the representation of a factoryTriggerTumblingWindow and creates it.  Returns the server's representation of the factoryTriggerTumblingWindow, and an error, if there is any.
func (c *FakeFactoryTriggerTumblingWindows) Create(ctx context.Context, factoryTriggerTumblingWindow *v1alpha1.FactoryTriggerTumblingWindow, opts v1.CreateOptions) (result *v1alpha1.FactoryTriggerTumblingWindow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(factorytriggertumblingwindowsResource, c.ns, factoryTriggerTumblingWindow), &v1alpha1.FactoryTriggerTumblingWindow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryTriggerTumblingWindow), err
}

// Update takes the representation of a factoryTriggerTumblingWindow and updates it. Returns the server's representation of the factoryTriggerTumblingWindow, and an error, if there is any.
func (c *FakeFactoryTriggerTumblingWindows) Update(ctx context.Context, factoryTriggerTumblingWindow *v1alpha1.FactoryTriggerTumblingWindow, opts v1.UpdateOptions) (result *v1alpha1.FactoryTriggerTumblingWindow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(factorytriggertumblingwindowsResource, c.ns, factoryTriggerTumblingWindow), &v1alpha1.FactoryTriggerTumblingWindow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryTriggerTumblingWindow), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFactoryTriggerTumblingWindows) UpdateStatus(ctx context.Context, factoryTriggerTumblingWindow *v1alpha1.FactoryTriggerTumblingWindow, opts v1.UpdateOptions) (*v1alpha1.FactoryTriggerTumblingWindow, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(factorytriggertumblingwindowsResource, "status", c.ns, factoryTriggerTumblingWindow), &v1alpha1.FactoryTriggerTumblingWindow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryTriggerTumblingWindow), err
}

// Delete takes name of the factoryTriggerTumblingWindow and deletes it. Returns an error if one occurs.
func (c *FakeFactoryTriggerTumblingWindows) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(factorytriggertumblingwindowsResource, c.ns, name), &v1alpha1.FactoryTriggerTumblingWindow{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFactoryTriggerTumblingWindows) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(factorytriggertumblingwindowsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FactoryTriggerTumblingWindowList{})
	return err
}

// Patch applies the patch and returns the patched factoryTriggerTumblingWindow.
func (c *FakeFactoryTriggerTumblingWindows) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FactoryTriggerTumblingWindow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(factorytriggertumblingwindowsResource, c.ns, name, pt, data, subresources...), &v1alpha1.FactoryTriggerTumblingWindow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FactoryTriggerTumblingWindow), err
}
