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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/management/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGroupTemplateDeployments implements GroupTemplateDeploymentInterface
type FakeGroupTemplateDeployments struct {
	Fake *FakeManagementV1alpha1
	ns   string
}

var grouptemplatedeploymentsResource = schema.GroupVersionResource{Group: "management.azurerm.kubeform.com", Version: "v1alpha1", Resource: "grouptemplatedeployments"}

var grouptemplatedeploymentsKind = schema.GroupVersionKind{Group: "management.azurerm.kubeform.com", Version: "v1alpha1", Kind: "GroupTemplateDeployment"}

// Get takes name of the groupTemplateDeployment, and returns the corresponding groupTemplateDeployment object, and an error if there is any.
func (c *FakeGroupTemplateDeployments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GroupTemplateDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(grouptemplatedeploymentsResource, c.ns, name), &v1alpha1.GroupTemplateDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupTemplateDeployment), err
}

// List takes label and field selectors, and returns the list of GroupTemplateDeployments that match those selectors.
func (c *FakeGroupTemplateDeployments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GroupTemplateDeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(grouptemplatedeploymentsResource, grouptemplatedeploymentsKind, c.ns, opts), &v1alpha1.GroupTemplateDeploymentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GroupTemplateDeploymentList{ListMeta: obj.(*v1alpha1.GroupTemplateDeploymentList).ListMeta}
	for _, item := range obj.(*v1alpha1.GroupTemplateDeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested groupTemplateDeployments.
func (c *FakeGroupTemplateDeployments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(grouptemplatedeploymentsResource, c.ns, opts))

}

// Create takes the representation of a groupTemplateDeployment and creates it.  Returns the server's representation of the groupTemplateDeployment, and an error, if there is any.
func (c *FakeGroupTemplateDeployments) Create(ctx context.Context, groupTemplateDeployment *v1alpha1.GroupTemplateDeployment, opts v1.CreateOptions) (result *v1alpha1.GroupTemplateDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(grouptemplatedeploymentsResource, c.ns, groupTemplateDeployment), &v1alpha1.GroupTemplateDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupTemplateDeployment), err
}

// Update takes the representation of a groupTemplateDeployment and updates it. Returns the server's representation of the groupTemplateDeployment, and an error, if there is any.
func (c *FakeGroupTemplateDeployments) Update(ctx context.Context, groupTemplateDeployment *v1alpha1.GroupTemplateDeployment, opts v1.UpdateOptions) (result *v1alpha1.GroupTemplateDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(grouptemplatedeploymentsResource, c.ns, groupTemplateDeployment), &v1alpha1.GroupTemplateDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupTemplateDeployment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGroupTemplateDeployments) UpdateStatus(ctx context.Context, groupTemplateDeployment *v1alpha1.GroupTemplateDeployment, opts v1.UpdateOptions) (*v1alpha1.GroupTemplateDeployment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(grouptemplatedeploymentsResource, "status", c.ns, groupTemplateDeployment), &v1alpha1.GroupTemplateDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupTemplateDeployment), err
}

// Delete takes name of the groupTemplateDeployment and deletes it. Returns an error if one occurs.
func (c *FakeGroupTemplateDeployments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(grouptemplatedeploymentsResource, c.ns, name), &v1alpha1.GroupTemplateDeployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGroupTemplateDeployments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(grouptemplatedeploymentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.GroupTemplateDeploymentList{})
	return err
}

// Patch applies the patch and returns the patched groupTemplateDeployment.
func (c *FakeGroupTemplateDeployments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GroupTemplateDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(grouptemplatedeploymentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.GroupTemplateDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupTemplateDeployment), err
}
