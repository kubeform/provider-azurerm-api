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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/mssql/v1alpha1"
	scheme "kubeform.dev/provider-azurerm-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ServerTransparentDataEncryptionsGetter has a method to return a ServerTransparentDataEncryptionInterface.
// A group's client should implement this interface.
type ServerTransparentDataEncryptionsGetter interface {
	ServerTransparentDataEncryptions(namespace string) ServerTransparentDataEncryptionInterface
}

// ServerTransparentDataEncryptionInterface has methods to work with ServerTransparentDataEncryption resources.
type ServerTransparentDataEncryptionInterface interface {
	Create(ctx context.Context, serverTransparentDataEncryption *v1alpha1.ServerTransparentDataEncryption, opts v1.CreateOptions) (*v1alpha1.ServerTransparentDataEncryption, error)
	Update(ctx context.Context, serverTransparentDataEncryption *v1alpha1.ServerTransparentDataEncryption, opts v1.UpdateOptions) (*v1alpha1.ServerTransparentDataEncryption, error)
	UpdateStatus(ctx context.Context, serverTransparentDataEncryption *v1alpha1.ServerTransparentDataEncryption, opts v1.UpdateOptions) (*v1alpha1.ServerTransparentDataEncryption, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ServerTransparentDataEncryption, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ServerTransparentDataEncryptionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServerTransparentDataEncryption, err error)
	ServerTransparentDataEncryptionExpansion
}

// serverTransparentDataEncryptions implements ServerTransparentDataEncryptionInterface
type serverTransparentDataEncryptions struct {
	client rest.Interface
	ns     string
}

// newServerTransparentDataEncryptions returns a ServerTransparentDataEncryptions
func newServerTransparentDataEncryptions(c *MssqlV1alpha1Client, namespace string) *serverTransparentDataEncryptions {
	return &serverTransparentDataEncryptions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serverTransparentDataEncryption, and returns the corresponding serverTransparentDataEncryption object, and an error if there is any.
func (c *serverTransparentDataEncryptions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServerTransparentDataEncryption, err error) {
	result = &v1alpha1.ServerTransparentDataEncryption{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servertransparentdataencryptions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServerTransparentDataEncryptions that match those selectors.
func (c *serverTransparentDataEncryptions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServerTransparentDataEncryptionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ServerTransparentDataEncryptionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servertransparentdataencryptions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serverTransparentDataEncryptions.
func (c *serverTransparentDataEncryptions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("servertransparentdataencryptions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a serverTransparentDataEncryption and creates it.  Returns the server's representation of the serverTransparentDataEncryption, and an error, if there is any.
func (c *serverTransparentDataEncryptions) Create(ctx context.Context, serverTransparentDataEncryption *v1alpha1.ServerTransparentDataEncryption, opts v1.CreateOptions) (result *v1alpha1.ServerTransparentDataEncryption, err error) {
	result = &v1alpha1.ServerTransparentDataEncryption{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("servertransparentdataencryptions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serverTransparentDataEncryption).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a serverTransparentDataEncryption and updates it. Returns the server's representation of the serverTransparentDataEncryption, and an error, if there is any.
func (c *serverTransparentDataEncryptions) Update(ctx context.Context, serverTransparentDataEncryption *v1alpha1.ServerTransparentDataEncryption, opts v1.UpdateOptions) (result *v1alpha1.ServerTransparentDataEncryption, err error) {
	result = &v1alpha1.ServerTransparentDataEncryption{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servertransparentdataencryptions").
		Name(serverTransparentDataEncryption.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serverTransparentDataEncryption).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *serverTransparentDataEncryptions) UpdateStatus(ctx context.Context, serverTransparentDataEncryption *v1alpha1.ServerTransparentDataEncryption, opts v1.UpdateOptions) (result *v1alpha1.ServerTransparentDataEncryption, err error) {
	result = &v1alpha1.ServerTransparentDataEncryption{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servertransparentdataencryptions").
		Name(serverTransparentDataEncryption.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serverTransparentDataEncryption).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the serverTransparentDataEncryption and deletes it. Returns an error if one occurs.
func (c *serverTransparentDataEncryptions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servertransparentdataencryptions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serverTransparentDataEncryptions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servertransparentdataencryptions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched serverTransparentDataEncryption.
func (c *serverTransparentDataEncryptions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServerTransparentDataEncryption, err error) {
	result = &v1alpha1.ServerTransparentDataEncryption{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("servertransparentdataencryptions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
