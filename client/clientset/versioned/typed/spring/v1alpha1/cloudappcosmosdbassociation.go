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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/spring/v1alpha1"
	scheme "kubeform.dev/provider-azurerm-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CloudAppCosmosdbAssociationsGetter has a method to return a CloudAppCosmosdbAssociationInterface.
// A group's client should implement this interface.
type CloudAppCosmosdbAssociationsGetter interface {
	CloudAppCosmosdbAssociations(namespace string) CloudAppCosmosdbAssociationInterface
}

// CloudAppCosmosdbAssociationInterface has methods to work with CloudAppCosmosdbAssociation resources.
type CloudAppCosmosdbAssociationInterface interface {
	Create(ctx context.Context, cloudAppCosmosdbAssociation *v1alpha1.CloudAppCosmosdbAssociation, opts v1.CreateOptions) (*v1alpha1.CloudAppCosmosdbAssociation, error)
	Update(ctx context.Context, cloudAppCosmosdbAssociation *v1alpha1.CloudAppCosmosdbAssociation, opts v1.UpdateOptions) (*v1alpha1.CloudAppCosmosdbAssociation, error)
	UpdateStatus(ctx context.Context, cloudAppCosmosdbAssociation *v1alpha1.CloudAppCosmosdbAssociation, opts v1.UpdateOptions) (*v1alpha1.CloudAppCosmosdbAssociation, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.CloudAppCosmosdbAssociation, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CloudAppCosmosdbAssociationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CloudAppCosmosdbAssociation, err error)
	CloudAppCosmosdbAssociationExpansion
}

// cloudAppCosmosdbAssociations implements CloudAppCosmosdbAssociationInterface
type cloudAppCosmosdbAssociations struct {
	client rest.Interface
	ns     string
}

// newCloudAppCosmosdbAssociations returns a CloudAppCosmosdbAssociations
func newCloudAppCosmosdbAssociations(c *SpringV1alpha1Client, namespace string) *cloudAppCosmosdbAssociations {
	return &cloudAppCosmosdbAssociations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cloudAppCosmosdbAssociation, and returns the corresponding cloudAppCosmosdbAssociation object, and an error if there is any.
func (c *cloudAppCosmosdbAssociations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CloudAppCosmosdbAssociation, err error) {
	result = &v1alpha1.CloudAppCosmosdbAssociation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudappcosmosdbassociations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudAppCosmosdbAssociations that match those selectors.
func (c *cloudAppCosmosdbAssociations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CloudAppCosmosdbAssociationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CloudAppCosmosdbAssociationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudappcosmosdbassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudAppCosmosdbAssociations.
func (c *cloudAppCosmosdbAssociations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cloudappcosmosdbassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cloudAppCosmosdbAssociation and creates it.  Returns the server's representation of the cloudAppCosmosdbAssociation, and an error, if there is any.
func (c *cloudAppCosmosdbAssociations) Create(ctx context.Context, cloudAppCosmosdbAssociation *v1alpha1.CloudAppCosmosdbAssociation, opts v1.CreateOptions) (result *v1alpha1.CloudAppCosmosdbAssociation, err error) {
	result = &v1alpha1.CloudAppCosmosdbAssociation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cloudappcosmosdbassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloudAppCosmosdbAssociation).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cloudAppCosmosdbAssociation and updates it. Returns the server's representation of the cloudAppCosmosdbAssociation, and an error, if there is any.
func (c *cloudAppCosmosdbAssociations) Update(ctx context.Context, cloudAppCosmosdbAssociation *v1alpha1.CloudAppCosmosdbAssociation, opts v1.UpdateOptions) (result *v1alpha1.CloudAppCosmosdbAssociation, err error) {
	result = &v1alpha1.CloudAppCosmosdbAssociation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudappcosmosdbassociations").
		Name(cloudAppCosmosdbAssociation.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloudAppCosmosdbAssociation).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *cloudAppCosmosdbAssociations) UpdateStatus(ctx context.Context, cloudAppCosmosdbAssociation *v1alpha1.CloudAppCosmosdbAssociation, opts v1.UpdateOptions) (result *v1alpha1.CloudAppCosmosdbAssociation, err error) {
	result = &v1alpha1.CloudAppCosmosdbAssociation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudappcosmosdbassociations").
		Name(cloudAppCosmosdbAssociation.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloudAppCosmosdbAssociation).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cloudAppCosmosdbAssociation and deletes it. Returns an error if one occurs.
func (c *cloudAppCosmosdbAssociations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudappcosmosdbassociations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudAppCosmosdbAssociations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudappcosmosdbassociations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cloudAppCosmosdbAssociation.
func (c *cloudAppCosmosdbAssociations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CloudAppCosmosdbAssociation, err error) {
	result = &v1alpha1.CloudAppCosmosdbAssociation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cloudappcosmosdbassociations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}