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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/virtual/v1alpha1"
	scheme "kubeform.dev/provider-azurerm-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HubIPsGetter has a method to return a HubIPInterface.
// A group's client should implement this interface.
type HubIPsGetter interface {
	HubIPs(namespace string) HubIPInterface
}

// HubIPInterface has methods to work with HubIP resources.
type HubIPInterface interface {
	Create(ctx context.Context, hubIP *v1alpha1.HubIP, opts v1.CreateOptions) (*v1alpha1.HubIP, error)
	Update(ctx context.Context, hubIP *v1alpha1.HubIP, opts v1.UpdateOptions) (*v1alpha1.HubIP, error)
	UpdateStatus(ctx context.Context, hubIP *v1alpha1.HubIP, opts v1.UpdateOptions) (*v1alpha1.HubIP, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.HubIP, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.HubIPList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HubIP, err error)
	HubIPExpansion
}

// hubIPs implements HubIPInterface
type hubIPs struct {
	client rest.Interface
	ns     string
}

// newHubIPs returns a HubIPs
func newHubIPs(c *VirtualV1alpha1Client, namespace string) *hubIPs {
	return &hubIPs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the hubIP, and returns the corresponding hubIP object, and an error if there is any.
func (c *hubIPs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HubIP, err error) {
	result = &v1alpha1.HubIP{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hubips").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HubIPs that match those selectors.
func (c *hubIPs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HubIPList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.HubIPList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hubips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hubIPs.
func (c *hubIPs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("hubips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a hubIP and creates it.  Returns the server's representation of the hubIP, and an error, if there is any.
func (c *hubIPs) Create(ctx context.Context, hubIP *v1alpha1.HubIP, opts v1.CreateOptions) (result *v1alpha1.HubIP, err error) {
	result = &v1alpha1.HubIP{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("hubips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hubIP).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a hubIP and updates it. Returns the server's representation of the hubIP, and an error, if there is any.
func (c *hubIPs) Update(ctx context.Context, hubIP *v1alpha1.HubIP, opts v1.UpdateOptions) (result *v1alpha1.HubIP, err error) {
	result = &v1alpha1.HubIP{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hubips").
		Name(hubIP.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hubIP).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *hubIPs) UpdateStatus(ctx context.Context, hubIP *v1alpha1.HubIP, opts v1.UpdateOptions) (result *v1alpha1.HubIP, err error) {
	result = &v1alpha1.HubIP{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hubips").
		Name(hubIP.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hubIP).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the hubIP and deletes it. Returns an error if one occurs.
func (c *hubIPs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hubips").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hubIPs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hubips").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched hubIP.
func (c *hubIPs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HubIP, err error) {
	result = &v1alpha1.HubIP{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("hubips").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
