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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/spring/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCloudJavaDeployments implements CloudJavaDeploymentInterface
type FakeCloudJavaDeployments struct {
	Fake *FakeSpringV1alpha1
	ns   string
}

var cloudjavadeploymentsResource = schema.GroupVersionResource{Group: "spring.azurerm.kubeform.com", Version: "v1alpha1", Resource: "cloudjavadeployments"}

var cloudjavadeploymentsKind = schema.GroupVersionKind{Group: "spring.azurerm.kubeform.com", Version: "v1alpha1", Kind: "CloudJavaDeployment"}

// Get takes name of the cloudJavaDeployment, and returns the corresponding cloudJavaDeployment object, and an error if there is any.
func (c *FakeCloudJavaDeployments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CloudJavaDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudjavadeploymentsResource, c.ns, name), &v1alpha1.CloudJavaDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudJavaDeployment), err
}

// List takes label and field selectors, and returns the list of CloudJavaDeployments that match those selectors.
func (c *FakeCloudJavaDeployments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CloudJavaDeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudjavadeploymentsResource, cloudjavadeploymentsKind, c.ns, opts), &v1alpha1.CloudJavaDeploymentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloudJavaDeploymentList{ListMeta: obj.(*v1alpha1.CloudJavaDeploymentList).ListMeta}
	for _, item := range obj.(*v1alpha1.CloudJavaDeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudJavaDeployments.
func (c *FakeCloudJavaDeployments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudjavadeploymentsResource, c.ns, opts))

}

// Create takes the representation of a cloudJavaDeployment and creates it.  Returns the server's representation of the cloudJavaDeployment, and an error, if there is any.
func (c *FakeCloudJavaDeployments) Create(ctx context.Context, cloudJavaDeployment *v1alpha1.CloudJavaDeployment, opts v1.CreateOptions) (result *v1alpha1.CloudJavaDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudjavadeploymentsResource, c.ns, cloudJavaDeployment), &v1alpha1.CloudJavaDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudJavaDeployment), err
}

// Update takes the representation of a cloudJavaDeployment and updates it. Returns the server's representation of the cloudJavaDeployment, and an error, if there is any.
func (c *FakeCloudJavaDeployments) Update(ctx context.Context, cloudJavaDeployment *v1alpha1.CloudJavaDeployment, opts v1.UpdateOptions) (result *v1alpha1.CloudJavaDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudjavadeploymentsResource, c.ns, cloudJavaDeployment), &v1alpha1.CloudJavaDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudJavaDeployment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudJavaDeployments) UpdateStatus(ctx context.Context, cloudJavaDeployment *v1alpha1.CloudJavaDeployment, opts v1.UpdateOptions) (*v1alpha1.CloudJavaDeployment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudjavadeploymentsResource, "status", c.ns, cloudJavaDeployment), &v1alpha1.CloudJavaDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudJavaDeployment), err
}

// Delete takes name of the cloudJavaDeployment and deletes it. Returns an error if one occurs.
func (c *FakeCloudJavaDeployments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cloudjavadeploymentsResource, c.ns, name), &v1alpha1.CloudJavaDeployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudJavaDeployments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudjavadeploymentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloudJavaDeploymentList{})
	return err
}

// Patch applies the patch and returns the patched cloudJavaDeployment.
func (c *FakeCloudJavaDeployments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CloudJavaDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudjavadeploymentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CloudJavaDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudJavaDeployment), err
}
