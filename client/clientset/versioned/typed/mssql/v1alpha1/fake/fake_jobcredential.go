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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/mssql/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeJobCredentials implements JobCredentialInterface
type FakeJobCredentials struct {
	Fake *FakeMssqlV1alpha1
	ns   string
}

var jobcredentialsResource = schema.GroupVersionResource{Group: "mssql.azurerm.kubeform.com", Version: "v1alpha1", Resource: "jobcredentials"}

var jobcredentialsKind = schema.GroupVersionKind{Group: "mssql.azurerm.kubeform.com", Version: "v1alpha1", Kind: "JobCredential"}

// Get takes name of the jobCredential, and returns the corresponding jobCredential object, and an error if there is any.
func (c *FakeJobCredentials) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.JobCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(jobcredentialsResource, c.ns, name), &v1alpha1.JobCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JobCredential), err
}

// List takes label and field selectors, and returns the list of JobCredentials that match those selectors.
func (c *FakeJobCredentials) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.JobCredentialList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(jobcredentialsResource, jobcredentialsKind, c.ns, opts), &v1alpha1.JobCredentialList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.JobCredentialList{ListMeta: obj.(*v1alpha1.JobCredentialList).ListMeta}
	for _, item := range obj.(*v1alpha1.JobCredentialList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested jobCredentials.
func (c *FakeJobCredentials) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(jobcredentialsResource, c.ns, opts))

}

// Create takes the representation of a jobCredential and creates it.  Returns the server's representation of the jobCredential, and an error, if there is any.
func (c *FakeJobCredentials) Create(ctx context.Context, jobCredential *v1alpha1.JobCredential, opts v1.CreateOptions) (result *v1alpha1.JobCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(jobcredentialsResource, c.ns, jobCredential), &v1alpha1.JobCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JobCredential), err
}

// Update takes the representation of a jobCredential and updates it. Returns the server's representation of the jobCredential, and an error, if there is any.
func (c *FakeJobCredentials) Update(ctx context.Context, jobCredential *v1alpha1.JobCredential, opts v1.UpdateOptions) (result *v1alpha1.JobCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(jobcredentialsResource, c.ns, jobCredential), &v1alpha1.JobCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JobCredential), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeJobCredentials) UpdateStatus(ctx context.Context, jobCredential *v1alpha1.JobCredential, opts v1.UpdateOptions) (*v1alpha1.JobCredential, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(jobcredentialsResource, "status", c.ns, jobCredential), &v1alpha1.JobCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JobCredential), err
}

// Delete takes name of the jobCredential and deletes it. Returns an error if one occurs.
func (c *FakeJobCredentials) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(jobcredentialsResource, c.ns, name), &v1alpha1.JobCredential{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeJobCredentials) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(jobcredentialsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.JobCredentialList{})
	return err
}

// Patch applies the patch and returns the patched jobCredential.
func (c *FakeJobCredentials) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.JobCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(jobcredentialsResource, c.ns, name, pt, data, subresources...), &v1alpha1.JobCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JobCredential), err
}
