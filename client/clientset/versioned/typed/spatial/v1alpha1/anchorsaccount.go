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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/spatial/v1alpha1"
	scheme "kubeform.dev/provider-azurerm-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AnchorsAccountsGetter has a method to return a AnchorsAccountInterface.
// A group's client should implement this interface.
type AnchorsAccountsGetter interface {
	AnchorsAccounts(namespace string) AnchorsAccountInterface
}

// AnchorsAccountInterface has methods to work with AnchorsAccount resources.
type AnchorsAccountInterface interface {
	Create(ctx context.Context, anchorsAccount *v1alpha1.AnchorsAccount, opts v1.CreateOptions) (*v1alpha1.AnchorsAccount, error)
	Update(ctx context.Context, anchorsAccount *v1alpha1.AnchorsAccount, opts v1.UpdateOptions) (*v1alpha1.AnchorsAccount, error)
	UpdateStatus(ctx context.Context, anchorsAccount *v1alpha1.AnchorsAccount, opts v1.UpdateOptions) (*v1alpha1.AnchorsAccount, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AnchorsAccount, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AnchorsAccountList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AnchorsAccount, err error)
	AnchorsAccountExpansion
}

// anchorsAccounts implements AnchorsAccountInterface
type anchorsAccounts struct {
	client rest.Interface
	ns     string
}

// newAnchorsAccounts returns a AnchorsAccounts
func newAnchorsAccounts(c *SpatialV1alpha1Client, namespace string) *anchorsAccounts {
	return &anchorsAccounts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the anchorsAccount, and returns the corresponding anchorsAccount object, and an error if there is any.
func (c *anchorsAccounts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AnchorsAccount, err error) {
	result = &v1alpha1.AnchorsAccount{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("anchorsaccounts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AnchorsAccounts that match those selectors.
func (c *anchorsAccounts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AnchorsAccountList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AnchorsAccountList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("anchorsaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested anchorsAccounts.
func (c *anchorsAccounts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("anchorsaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a anchorsAccount and creates it.  Returns the server's representation of the anchorsAccount, and an error, if there is any.
func (c *anchorsAccounts) Create(ctx context.Context, anchorsAccount *v1alpha1.AnchorsAccount, opts v1.CreateOptions) (result *v1alpha1.AnchorsAccount, err error) {
	result = &v1alpha1.AnchorsAccount{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("anchorsaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(anchorsAccount).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a anchorsAccount and updates it. Returns the server's representation of the anchorsAccount, and an error, if there is any.
func (c *anchorsAccounts) Update(ctx context.Context, anchorsAccount *v1alpha1.AnchorsAccount, opts v1.UpdateOptions) (result *v1alpha1.AnchorsAccount, err error) {
	result = &v1alpha1.AnchorsAccount{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("anchorsaccounts").
		Name(anchorsAccount.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(anchorsAccount).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *anchorsAccounts) UpdateStatus(ctx context.Context, anchorsAccount *v1alpha1.AnchorsAccount, opts v1.UpdateOptions) (result *v1alpha1.AnchorsAccount, err error) {
	result = &v1alpha1.AnchorsAccount{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("anchorsaccounts").
		Name(anchorsAccount.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(anchorsAccount).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the anchorsAccount and deletes it. Returns an error if one occurs.
func (c *anchorsAccounts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("anchorsaccounts").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *anchorsAccounts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("anchorsaccounts").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched anchorsAccount.
func (c *anchorsAccounts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AnchorsAccount, err error) {
	result = &v1alpha1.AnchorsAccount{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("anchorsaccounts").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
