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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/app/v1alpha1"
	scheme "kubeform.dev/provider-azurerm-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ServiceCertificateBindingsGetter has a method to return a ServiceCertificateBindingInterface.
// A group's client should implement this interface.
type ServiceCertificateBindingsGetter interface {
	ServiceCertificateBindings(namespace string) ServiceCertificateBindingInterface
}

// ServiceCertificateBindingInterface has methods to work with ServiceCertificateBinding resources.
type ServiceCertificateBindingInterface interface {
	Create(ctx context.Context, serviceCertificateBinding *v1alpha1.ServiceCertificateBinding, opts v1.CreateOptions) (*v1alpha1.ServiceCertificateBinding, error)
	Update(ctx context.Context, serviceCertificateBinding *v1alpha1.ServiceCertificateBinding, opts v1.UpdateOptions) (*v1alpha1.ServiceCertificateBinding, error)
	UpdateStatus(ctx context.Context, serviceCertificateBinding *v1alpha1.ServiceCertificateBinding, opts v1.UpdateOptions) (*v1alpha1.ServiceCertificateBinding, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ServiceCertificateBinding, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ServiceCertificateBindingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceCertificateBinding, err error)
	ServiceCertificateBindingExpansion
}

// serviceCertificateBindings implements ServiceCertificateBindingInterface
type serviceCertificateBindings struct {
	client rest.Interface
	ns     string
}

// newServiceCertificateBindings returns a ServiceCertificateBindings
func newServiceCertificateBindings(c *AppV1alpha1Client, namespace string) *serviceCertificateBindings {
	return &serviceCertificateBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serviceCertificateBinding, and returns the corresponding serviceCertificateBinding object, and an error if there is any.
func (c *serviceCertificateBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServiceCertificateBinding, err error) {
	result = &v1alpha1.ServiceCertificateBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicecertificatebindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceCertificateBindings that match those selectors.
func (c *serviceCertificateBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServiceCertificateBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ServiceCertificateBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicecertificatebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceCertificateBindings.
func (c *serviceCertificateBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("servicecertificatebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a serviceCertificateBinding and creates it.  Returns the server's representation of the serviceCertificateBinding, and an error, if there is any.
func (c *serviceCertificateBindings) Create(ctx context.Context, serviceCertificateBinding *v1alpha1.ServiceCertificateBinding, opts v1.CreateOptions) (result *v1alpha1.ServiceCertificateBinding, err error) {
	result = &v1alpha1.ServiceCertificateBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("servicecertificatebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceCertificateBinding).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a serviceCertificateBinding and updates it. Returns the server's representation of the serviceCertificateBinding, and an error, if there is any.
func (c *serviceCertificateBindings) Update(ctx context.Context, serviceCertificateBinding *v1alpha1.ServiceCertificateBinding, opts v1.UpdateOptions) (result *v1alpha1.ServiceCertificateBinding, err error) {
	result = &v1alpha1.ServiceCertificateBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicecertificatebindings").
		Name(serviceCertificateBinding.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceCertificateBinding).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *serviceCertificateBindings) UpdateStatus(ctx context.Context, serviceCertificateBinding *v1alpha1.ServiceCertificateBinding, opts v1.UpdateOptions) (result *v1alpha1.ServiceCertificateBinding, err error) {
	result = &v1alpha1.ServiceCertificateBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicecertificatebindings").
		Name(serviceCertificateBinding.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceCertificateBinding).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the serviceCertificateBinding and deletes it. Returns an error if one occurs.
func (c *serviceCertificateBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicecertificatebindings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serviceCertificateBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicecertificatebindings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched serviceCertificateBinding.
func (c *serviceCertificateBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceCertificateBinding, err error) {
	result = &v1alpha1.ServiceCertificateBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("servicecertificatebindings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
