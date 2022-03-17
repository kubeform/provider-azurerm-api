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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/machine/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLearningComputeInstances implements LearningComputeInstanceInterface
type FakeLearningComputeInstances struct {
	Fake *FakeMachineV1alpha1
	ns   string
}

var learningcomputeinstancesResource = schema.GroupVersionResource{Group: "machine.azurerm.kubeform.com", Version: "v1alpha1", Resource: "learningcomputeinstances"}

var learningcomputeinstancesKind = schema.GroupVersionKind{Group: "machine.azurerm.kubeform.com", Version: "v1alpha1", Kind: "LearningComputeInstance"}

// Get takes name of the learningComputeInstance, and returns the corresponding learningComputeInstance object, and an error if there is any.
func (c *FakeLearningComputeInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LearningComputeInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(learningcomputeinstancesResource, c.ns, name), &v1alpha1.LearningComputeInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LearningComputeInstance), err
}

// List takes label and field selectors, and returns the list of LearningComputeInstances that match those selectors.
func (c *FakeLearningComputeInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LearningComputeInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(learningcomputeinstancesResource, learningcomputeinstancesKind, c.ns, opts), &v1alpha1.LearningComputeInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LearningComputeInstanceList{ListMeta: obj.(*v1alpha1.LearningComputeInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.LearningComputeInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested learningComputeInstances.
func (c *FakeLearningComputeInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(learningcomputeinstancesResource, c.ns, opts))

}

// Create takes the representation of a learningComputeInstance and creates it.  Returns the server's representation of the learningComputeInstance, and an error, if there is any.
func (c *FakeLearningComputeInstances) Create(ctx context.Context, learningComputeInstance *v1alpha1.LearningComputeInstance, opts v1.CreateOptions) (result *v1alpha1.LearningComputeInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(learningcomputeinstancesResource, c.ns, learningComputeInstance), &v1alpha1.LearningComputeInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LearningComputeInstance), err
}

// Update takes the representation of a learningComputeInstance and updates it. Returns the server's representation of the learningComputeInstance, and an error, if there is any.
func (c *FakeLearningComputeInstances) Update(ctx context.Context, learningComputeInstance *v1alpha1.LearningComputeInstance, opts v1.UpdateOptions) (result *v1alpha1.LearningComputeInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(learningcomputeinstancesResource, c.ns, learningComputeInstance), &v1alpha1.LearningComputeInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LearningComputeInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLearningComputeInstances) UpdateStatus(ctx context.Context, learningComputeInstance *v1alpha1.LearningComputeInstance, opts v1.UpdateOptions) (*v1alpha1.LearningComputeInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(learningcomputeinstancesResource, "status", c.ns, learningComputeInstance), &v1alpha1.LearningComputeInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LearningComputeInstance), err
}

// Delete takes name of the learningComputeInstance and deletes it. Returns an error if one occurs.
func (c *FakeLearningComputeInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(learningcomputeinstancesResource, c.ns, name), &v1alpha1.LearningComputeInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLearningComputeInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(learningcomputeinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LearningComputeInstanceList{})
	return err
}

// Patch applies the patch and returns the patched learningComputeInstance.
func (c *FakeLearningComputeInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LearningComputeInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(learningcomputeinstancesResource, c.ns, name, pt, data, subresources...), &v1alpha1.LearningComputeInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LearningComputeInstance), err
}