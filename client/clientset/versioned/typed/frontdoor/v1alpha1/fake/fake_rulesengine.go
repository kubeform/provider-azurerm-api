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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/frontdoor/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRulesEngines implements RulesEngineInterface
type FakeRulesEngines struct {
	Fake *FakeFrontdoorV1alpha1
	ns   string
}

var rulesenginesResource = schema.GroupVersionResource{Group: "frontdoor.azurerm.kubeform.com", Version: "v1alpha1", Resource: "rulesengines"}

var rulesenginesKind = schema.GroupVersionKind{Group: "frontdoor.azurerm.kubeform.com", Version: "v1alpha1", Kind: "RulesEngine"}

// Get takes name of the rulesEngine, and returns the corresponding rulesEngine object, and an error if there is any.
func (c *FakeRulesEngines) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RulesEngine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rulesenginesResource, c.ns, name), &v1alpha1.RulesEngine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RulesEngine), err
}

// List takes label and field selectors, and returns the list of RulesEngines that match those selectors.
func (c *FakeRulesEngines) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RulesEngineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rulesenginesResource, rulesenginesKind, c.ns, opts), &v1alpha1.RulesEngineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RulesEngineList{ListMeta: obj.(*v1alpha1.RulesEngineList).ListMeta}
	for _, item := range obj.(*v1alpha1.RulesEngineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rulesEngines.
func (c *FakeRulesEngines) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rulesenginesResource, c.ns, opts))

}

// Create takes the representation of a rulesEngine and creates it.  Returns the server's representation of the rulesEngine, and an error, if there is any.
func (c *FakeRulesEngines) Create(ctx context.Context, rulesEngine *v1alpha1.RulesEngine, opts v1.CreateOptions) (result *v1alpha1.RulesEngine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rulesenginesResource, c.ns, rulesEngine), &v1alpha1.RulesEngine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RulesEngine), err
}

// Update takes the representation of a rulesEngine and updates it. Returns the server's representation of the rulesEngine, and an error, if there is any.
func (c *FakeRulesEngines) Update(ctx context.Context, rulesEngine *v1alpha1.RulesEngine, opts v1.UpdateOptions) (result *v1alpha1.RulesEngine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rulesenginesResource, c.ns, rulesEngine), &v1alpha1.RulesEngine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RulesEngine), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRulesEngines) UpdateStatus(ctx context.Context, rulesEngine *v1alpha1.RulesEngine, opts v1.UpdateOptions) (*v1alpha1.RulesEngine, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(rulesenginesResource, "status", c.ns, rulesEngine), &v1alpha1.RulesEngine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RulesEngine), err
}

// Delete takes name of the rulesEngine and deletes it. Returns an error if one occurs.
func (c *FakeRulesEngines) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(rulesenginesResource, c.ns, name), &v1alpha1.RulesEngine{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRulesEngines) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rulesenginesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RulesEngineList{})
	return err
}

// Patch applies the patch and returns the patched rulesEngine.
func (c *FakeRulesEngines) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RulesEngine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rulesenginesResource, c.ns, name, pt, data, subresources...), &v1alpha1.RulesEngine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RulesEngine), err
}