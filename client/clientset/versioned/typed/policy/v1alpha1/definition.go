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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/policy/v1alpha1"
	scheme "kubeform.dev/provider-azurerm-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DefinitionsGetter has a method to return a DefinitionInterface.
// A group's client should implement this interface.
type DefinitionsGetter interface {
	Definitions(namespace string) DefinitionInterface
}

// DefinitionInterface has methods to work with Definition resources.
type DefinitionInterface interface {
	Create(ctx context.Context, definition *v1alpha1.Definition, opts v1.CreateOptions) (*v1alpha1.Definition, error)
	Update(ctx context.Context, definition *v1alpha1.Definition, opts v1.UpdateOptions) (*v1alpha1.Definition, error)
	UpdateStatus(ctx context.Context, definition *v1alpha1.Definition, opts v1.UpdateOptions) (*v1alpha1.Definition, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Definition, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DefinitionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Definition, err error)
	DefinitionExpansion
}

// definitions implements DefinitionInterface
type definitions struct {
	client rest.Interface
	ns     string
}

// newDefinitions returns a Definitions
func newDefinitions(c *PolicyV1alpha1Client, namespace string) *definitions {
	return &definitions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the definition, and returns the corresponding definition object, and an error if there is any.
func (c *definitions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Definition, err error) {
	result = &v1alpha1.Definition{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("definitions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Definitions that match those selectors.
func (c *definitions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DefinitionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DefinitionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("definitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested definitions.
func (c *definitions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("definitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a definition and creates it.  Returns the server's representation of the definition, and an error, if there is any.
func (c *definitions) Create(ctx context.Context, definition *v1alpha1.Definition, opts v1.CreateOptions) (result *v1alpha1.Definition, err error) {
	result = &v1alpha1.Definition{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("definitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(definition).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a definition and updates it. Returns the server's representation of the definition, and an error, if there is any.
func (c *definitions) Update(ctx context.Context, definition *v1alpha1.Definition, opts v1.UpdateOptions) (result *v1alpha1.Definition, err error) {
	result = &v1alpha1.Definition{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("definitions").
		Name(definition.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(definition).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *definitions) UpdateStatus(ctx context.Context, definition *v1alpha1.Definition, opts v1.UpdateOptions) (result *v1alpha1.Definition, err error) {
	result = &v1alpha1.Definition{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("definitions").
		Name(definition.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(definition).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the definition and deletes it. Returns an error if one occurs.
func (c *definitions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("definitions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *definitions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("definitions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched definition.
func (c *definitions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Definition, err error) {
	result = &v1alpha1.Definition{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("definitions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
