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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/expressroute/v1alpha1"
	scheme "kubeform.dev/provider-azurerm-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CircuitPeeringsGetter has a method to return a CircuitPeeringInterface.
// A group's client should implement this interface.
type CircuitPeeringsGetter interface {
	CircuitPeerings(namespace string) CircuitPeeringInterface
}

// CircuitPeeringInterface has methods to work with CircuitPeering resources.
type CircuitPeeringInterface interface {
	Create(ctx context.Context, circuitPeering *v1alpha1.CircuitPeering, opts v1.CreateOptions) (*v1alpha1.CircuitPeering, error)
	Update(ctx context.Context, circuitPeering *v1alpha1.CircuitPeering, opts v1.UpdateOptions) (*v1alpha1.CircuitPeering, error)
	UpdateStatus(ctx context.Context, circuitPeering *v1alpha1.CircuitPeering, opts v1.UpdateOptions) (*v1alpha1.CircuitPeering, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.CircuitPeering, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CircuitPeeringList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CircuitPeering, err error)
	CircuitPeeringExpansion
}

// circuitPeerings implements CircuitPeeringInterface
type circuitPeerings struct {
	client rest.Interface
	ns     string
}

// newCircuitPeerings returns a CircuitPeerings
func newCircuitPeerings(c *ExpressrouteV1alpha1Client, namespace string) *circuitPeerings {
	return &circuitPeerings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the circuitPeering, and returns the corresponding circuitPeering object, and an error if there is any.
func (c *circuitPeerings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CircuitPeering, err error) {
	result = &v1alpha1.CircuitPeering{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("circuitpeerings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CircuitPeerings that match those selectors.
func (c *circuitPeerings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CircuitPeeringList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CircuitPeeringList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("circuitpeerings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested circuitPeerings.
func (c *circuitPeerings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("circuitpeerings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a circuitPeering and creates it.  Returns the server's representation of the circuitPeering, and an error, if there is any.
func (c *circuitPeerings) Create(ctx context.Context, circuitPeering *v1alpha1.CircuitPeering, opts v1.CreateOptions) (result *v1alpha1.CircuitPeering, err error) {
	result = &v1alpha1.CircuitPeering{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("circuitpeerings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(circuitPeering).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a circuitPeering and updates it. Returns the server's representation of the circuitPeering, and an error, if there is any.
func (c *circuitPeerings) Update(ctx context.Context, circuitPeering *v1alpha1.CircuitPeering, opts v1.UpdateOptions) (result *v1alpha1.CircuitPeering, err error) {
	result = &v1alpha1.CircuitPeering{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("circuitpeerings").
		Name(circuitPeering.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(circuitPeering).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *circuitPeerings) UpdateStatus(ctx context.Context, circuitPeering *v1alpha1.CircuitPeering, opts v1.UpdateOptions) (result *v1alpha1.CircuitPeering, err error) {
	result = &v1alpha1.CircuitPeering{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("circuitpeerings").
		Name(circuitPeering.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(circuitPeering).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the circuitPeering and deletes it. Returns an error if one occurs.
func (c *circuitPeerings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("circuitpeerings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *circuitPeerings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("circuitpeerings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched circuitPeering.
func (c *circuitPeerings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CircuitPeering, err error) {
	result = &v1alpha1.CircuitPeering{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("circuitpeerings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
