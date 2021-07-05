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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/keyvault/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCertificateIssuers implements CertificateIssuerInterface
type FakeCertificateIssuers struct {
	Fake *FakeKeyvaultV1alpha1
	ns   string
}

var certificateissuersResource = schema.GroupVersionResource{Group: "keyvault.azurerm.kubeform.com", Version: "v1alpha1", Resource: "certificateissuers"}

var certificateissuersKind = schema.GroupVersionKind{Group: "keyvault.azurerm.kubeform.com", Version: "v1alpha1", Kind: "CertificateIssuer"}

// Get takes name of the certificateIssuer, and returns the corresponding certificateIssuer object, and an error if there is any.
func (c *FakeCertificateIssuers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CertificateIssuer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(certificateissuersResource, c.ns, name), &v1alpha1.CertificateIssuer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateIssuer), err
}

// List takes label and field selectors, and returns the list of CertificateIssuers that match those selectors.
func (c *FakeCertificateIssuers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CertificateIssuerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(certificateissuersResource, certificateissuersKind, c.ns, opts), &v1alpha1.CertificateIssuerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CertificateIssuerList{ListMeta: obj.(*v1alpha1.CertificateIssuerList).ListMeta}
	for _, item := range obj.(*v1alpha1.CertificateIssuerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested certificateIssuers.
func (c *FakeCertificateIssuers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(certificateissuersResource, c.ns, opts))

}

// Create takes the representation of a certificateIssuer and creates it.  Returns the server's representation of the certificateIssuer, and an error, if there is any.
func (c *FakeCertificateIssuers) Create(ctx context.Context, certificateIssuer *v1alpha1.CertificateIssuer, opts v1.CreateOptions) (result *v1alpha1.CertificateIssuer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(certificateissuersResource, c.ns, certificateIssuer), &v1alpha1.CertificateIssuer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateIssuer), err
}

// Update takes the representation of a certificateIssuer and updates it. Returns the server's representation of the certificateIssuer, and an error, if there is any.
func (c *FakeCertificateIssuers) Update(ctx context.Context, certificateIssuer *v1alpha1.CertificateIssuer, opts v1.UpdateOptions) (result *v1alpha1.CertificateIssuer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(certificateissuersResource, c.ns, certificateIssuer), &v1alpha1.CertificateIssuer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateIssuer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCertificateIssuers) UpdateStatus(ctx context.Context, certificateIssuer *v1alpha1.CertificateIssuer, opts v1.UpdateOptions) (*v1alpha1.CertificateIssuer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(certificateissuersResource, "status", c.ns, certificateIssuer), &v1alpha1.CertificateIssuer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateIssuer), err
}

// Delete takes name of the certificateIssuer and deletes it. Returns an error if one occurs.
func (c *FakeCertificateIssuers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(certificateissuersResource, c.ns, name), &v1alpha1.CertificateIssuer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCertificateIssuers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(certificateissuersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CertificateIssuerList{})
	return err
}

// Patch applies the patch and returns the patched certificateIssuer.
func (c *FakeCertificateIssuers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CertificateIssuer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(certificateissuersResource, c.ns, name, pt, data, subresources...), &v1alpha1.CertificateIssuer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateIssuer), err
}