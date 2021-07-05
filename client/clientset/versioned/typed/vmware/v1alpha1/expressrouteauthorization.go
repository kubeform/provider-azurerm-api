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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/vmware/v1alpha1"
	scheme "kubeform.dev/provider-azurerm-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ExpressRouteAuthorizationsGetter has a method to return a ExpressRouteAuthorizationInterface.
// A group's client should implement this interface.
type ExpressRouteAuthorizationsGetter interface {
	ExpressRouteAuthorizations(namespace string) ExpressRouteAuthorizationInterface
}

// ExpressRouteAuthorizationInterface has methods to work with ExpressRouteAuthorization resources.
type ExpressRouteAuthorizationInterface interface {
	Create(ctx context.Context, expressRouteAuthorization *v1alpha1.ExpressRouteAuthorization, opts v1.CreateOptions) (*v1alpha1.ExpressRouteAuthorization, error)
	Update(ctx context.Context, expressRouteAuthorization *v1alpha1.ExpressRouteAuthorization, opts v1.UpdateOptions) (*v1alpha1.ExpressRouteAuthorization, error)
	UpdateStatus(ctx context.Context, expressRouteAuthorization *v1alpha1.ExpressRouteAuthorization, opts v1.UpdateOptions) (*v1alpha1.ExpressRouteAuthorization, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ExpressRouteAuthorization, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ExpressRouteAuthorizationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ExpressRouteAuthorization, err error)
	ExpressRouteAuthorizationExpansion
}

// expressRouteAuthorizations implements ExpressRouteAuthorizationInterface
type expressRouteAuthorizations struct {
	client rest.Interface
	ns     string
}

// newExpressRouteAuthorizations returns a ExpressRouteAuthorizations
func newExpressRouteAuthorizations(c *VmwareV1alpha1Client, namespace string) *expressRouteAuthorizations {
	return &expressRouteAuthorizations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the expressRouteAuthorization, and returns the corresponding expressRouteAuthorization object, and an error if there is any.
func (c *expressRouteAuthorizations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ExpressRouteAuthorization, err error) {
	result = &v1alpha1.ExpressRouteAuthorization{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("expressrouteauthorizations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ExpressRouteAuthorizations that match those selectors.
func (c *expressRouteAuthorizations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ExpressRouteAuthorizationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ExpressRouteAuthorizationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("expressrouteauthorizations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested expressRouteAuthorizations.
func (c *expressRouteAuthorizations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("expressrouteauthorizations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a expressRouteAuthorization and creates it.  Returns the server's representation of the expressRouteAuthorization, and an error, if there is any.
func (c *expressRouteAuthorizations) Create(ctx context.Context, expressRouteAuthorization *v1alpha1.ExpressRouteAuthorization, opts v1.CreateOptions) (result *v1alpha1.ExpressRouteAuthorization, err error) {
	result = &v1alpha1.ExpressRouteAuthorization{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("expressrouteauthorizations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(expressRouteAuthorization).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a expressRouteAuthorization and updates it. Returns the server's representation of the expressRouteAuthorization, and an error, if there is any.
func (c *expressRouteAuthorizations) Update(ctx context.Context, expressRouteAuthorization *v1alpha1.ExpressRouteAuthorization, opts v1.UpdateOptions) (result *v1alpha1.ExpressRouteAuthorization, err error) {
	result = &v1alpha1.ExpressRouteAuthorization{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("expressrouteauthorizations").
		Name(expressRouteAuthorization.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(expressRouteAuthorization).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *expressRouteAuthorizations) UpdateStatus(ctx context.Context, expressRouteAuthorization *v1alpha1.ExpressRouteAuthorization, opts v1.UpdateOptions) (result *v1alpha1.ExpressRouteAuthorization, err error) {
	result = &v1alpha1.ExpressRouteAuthorization{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("expressrouteauthorizations").
		Name(expressRouteAuthorization.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(expressRouteAuthorization).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the expressRouteAuthorization and deletes it. Returns an error if one occurs.
func (c *expressRouteAuthorizations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("expressrouteauthorizations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *expressRouteAuthorizations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("expressrouteauthorizations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched expressRouteAuthorization.
func (c *expressRouteAuthorizations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ExpressRouteAuthorization, err error) {
	result = &v1alpha1.ExpressRouteAuthorization{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("expressrouteauthorizations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}