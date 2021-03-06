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

// FakeEmailTemplates implements EmailTemplateInterface
type FakeEmailTemplates struct {
	Fake *FakeApimanagementV1alpha1
	ns   string
}

var emailtemplatesResource = schema.GroupVersionResource{Group: "apimanagement.azurerm.kubeform.com", Version: "v1alpha1", Resource: "emailtemplates"}

var emailtemplatesKind = schema.GroupVersionKind{Group: "apimanagement.azurerm.kubeform.com", Version: "v1alpha1", Kind: "EmailTemplate"}

// Get takes name of the emailTemplate, and returns the corresponding emailTemplate object, and an error if there is any.
func (c *FakeEmailTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EmailTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(emailtemplatesResource, c.ns, name), &v1alpha1.EmailTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EmailTemplate), err
}

// List takes label and field selectors, and returns the list of EmailTemplates that match those selectors.
func (c *FakeEmailTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EmailTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(emailtemplatesResource, emailtemplatesKind, c.ns, opts), &v1alpha1.EmailTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EmailTemplateList{ListMeta: obj.(*v1alpha1.EmailTemplateList).ListMeta}
	for _, item := range obj.(*v1alpha1.EmailTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested emailTemplates.
func (c *FakeEmailTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(emailtemplatesResource, c.ns, opts))

}

// Create takes the representation of a emailTemplate and creates it.  Returns the server's representation of the emailTemplate, and an error, if there is any.
func (c *FakeEmailTemplates) Create(ctx context.Context, emailTemplate *v1alpha1.EmailTemplate, opts v1.CreateOptions) (result *v1alpha1.EmailTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(emailtemplatesResource, c.ns, emailTemplate), &v1alpha1.EmailTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EmailTemplate), err
}

// Update takes the representation of a emailTemplate and updates it. Returns the server's representation of the emailTemplate, and an error, if there is any.
func (c *FakeEmailTemplates) Update(ctx context.Context, emailTemplate *v1alpha1.EmailTemplate, opts v1.UpdateOptions) (result *v1alpha1.EmailTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(emailtemplatesResource, c.ns, emailTemplate), &v1alpha1.EmailTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EmailTemplate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEmailTemplates) UpdateStatus(ctx context.Context, emailTemplate *v1alpha1.EmailTemplate, opts v1.UpdateOptions) (*v1alpha1.EmailTemplate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(emailtemplatesResource, "status", c.ns, emailTemplate), &v1alpha1.EmailTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EmailTemplate), err
}

// Delete takes name of the emailTemplate and deletes it. Returns an error if one occurs.
func (c *FakeEmailTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(emailtemplatesResource, c.ns, name), &v1alpha1.EmailTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEmailTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(emailtemplatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EmailTemplateList{})
	return err
}

// Patch applies the patch and returns the patched emailTemplate.
func (c *FakeEmailTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EmailTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(emailtemplatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.EmailTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EmailTemplate), err
}
