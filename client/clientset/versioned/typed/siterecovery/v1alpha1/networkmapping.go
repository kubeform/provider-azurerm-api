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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/siterecovery/v1alpha1"
	scheme "kubeform.dev/provider-azurerm-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NetworkMappingsGetter has a method to return a NetworkMappingInterface.
// A group's client should implement this interface.
type NetworkMappingsGetter interface {
	NetworkMappings(namespace string) NetworkMappingInterface
}

// NetworkMappingInterface has methods to work with NetworkMapping resources.
type NetworkMappingInterface interface {
	Create(ctx context.Context, networkMapping *v1alpha1.NetworkMapping, opts v1.CreateOptions) (*v1alpha1.NetworkMapping, error)
	Update(ctx context.Context, networkMapping *v1alpha1.NetworkMapping, opts v1.UpdateOptions) (*v1alpha1.NetworkMapping, error)
	UpdateStatus(ctx context.Context, networkMapping *v1alpha1.NetworkMapping, opts v1.UpdateOptions) (*v1alpha1.NetworkMapping, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.NetworkMapping, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.NetworkMappingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkMapping, err error)
	NetworkMappingExpansion
}

// networkMappings implements NetworkMappingInterface
type networkMappings struct {
	client rest.Interface
	ns     string
}

// newNetworkMappings returns a NetworkMappings
func newNetworkMappings(c *SiterecoveryV1alpha1Client, namespace string) *networkMappings {
	return &networkMappings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the networkMapping, and returns the corresponding networkMapping object, and an error if there is any.
func (c *networkMappings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkMapping, err error) {
	result = &v1alpha1.NetworkMapping{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkmappings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkMappings that match those selectors.
func (c *networkMappings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetworkMappingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NetworkMappingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkmappings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkMappings.
func (c *networkMappings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("networkmappings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a networkMapping and creates it.  Returns the server's representation of the networkMapping, and an error, if there is any.
func (c *networkMappings) Create(ctx context.Context, networkMapping *v1alpha1.NetworkMapping, opts v1.CreateOptions) (result *v1alpha1.NetworkMapping, err error) {
	result = &v1alpha1.NetworkMapping{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("networkmappings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkMapping).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a networkMapping and updates it. Returns the server's representation of the networkMapping, and an error, if there is any.
func (c *networkMappings) Update(ctx context.Context, networkMapping *v1alpha1.NetworkMapping, opts v1.UpdateOptions) (result *v1alpha1.NetworkMapping, err error) {
	result = &v1alpha1.NetworkMapping{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkmappings").
		Name(networkMapping.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkMapping).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *networkMappings) UpdateStatus(ctx context.Context, networkMapping *v1alpha1.NetworkMapping, opts v1.UpdateOptions) (result *v1alpha1.NetworkMapping, err error) {
	result = &v1alpha1.NetworkMapping{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkmappings").
		Name(networkMapping.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkMapping).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the networkMapping and deletes it. Returns an error if one occurs.
func (c *networkMappings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkmappings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkMappings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkmappings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched networkMapping.
func (c *networkMappings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkMapping, err error) {
	result = &v1alpha1.NetworkMapping{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("networkmappings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
