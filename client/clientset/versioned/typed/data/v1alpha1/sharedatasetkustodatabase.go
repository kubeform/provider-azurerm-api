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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/data/v1alpha1"
	scheme "kubeform.dev/provider-azurerm-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ShareDatasetKustoDatabasesGetter has a method to return a ShareDatasetKustoDatabaseInterface.
// A group's client should implement this interface.
type ShareDatasetKustoDatabasesGetter interface {
	ShareDatasetKustoDatabases(namespace string) ShareDatasetKustoDatabaseInterface
}

// ShareDatasetKustoDatabaseInterface has methods to work with ShareDatasetKustoDatabase resources.
type ShareDatasetKustoDatabaseInterface interface {
	Create(ctx context.Context, shareDatasetKustoDatabase *v1alpha1.ShareDatasetKustoDatabase, opts v1.CreateOptions) (*v1alpha1.ShareDatasetKustoDatabase, error)
	Update(ctx context.Context, shareDatasetKustoDatabase *v1alpha1.ShareDatasetKustoDatabase, opts v1.UpdateOptions) (*v1alpha1.ShareDatasetKustoDatabase, error)
	UpdateStatus(ctx context.Context, shareDatasetKustoDatabase *v1alpha1.ShareDatasetKustoDatabase, opts v1.UpdateOptions) (*v1alpha1.ShareDatasetKustoDatabase, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ShareDatasetKustoDatabase, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ShareDatasetKustoDatabaseList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ShareDatasetKustoDatabase, err error)
	ShareDatasetKustoDatabaseExpansion
}

// shareDatasetKustoDatabases implements ShareDatasetKustoDatabaseInterface
type shareDatasetKustoDatabases struct {
	client rest.Interface
	ns     string
}

// newShareDatasetKustoDatabases returns a ShareDatasetKustoDatabases
func newShareDatasetKustoDatabases(c *DataV1alpha1Client, namespace string) *shareDatasetKustoDatabases {
	return &shareDatasetKustoDatabases{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the shareDatasetKustoDatabase, and returns the corresponding shareDatasetKustoDatabase object, and an error if there is any.
func (c *shareDatasetKustoDatabases) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ShareDatasetKustoDatabase, err error) {
	result = &v1alpha1.ShareDatasetKustoDatabase{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sharedatasetkustodatabases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ShareDatasetKustoDatabases that match those selectors.
func (c *shareDatasetKustoDatabases) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ShareDatasetKustoDatabaseList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ShareDatasetKustoDatabaseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sharedatasetkustodatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested shareDatasetKustoDatabases.
func (c *shareDatasetKustoDatabases) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sharedatasetkustodatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a shareDatasetKustoDatabase and creates it.  Returns the server's representation of the shareDatasetKustoDatabase, and an error, if there is any.
func (c *shareDatasetKustoDatabases) Create(ctx context.Context, shareDatasetKustoDatabase *v1alpha1.ShareDatasetKustoDatabase, opts v1.CreateOptions) (result *v1alpha1.ShareDatasetKustoDatabase, err error) {
	result = &v1alpha1.ShareDatasetKustoDatabase{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sharedatasetkustodatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(shareDatasetKustoDatabase).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a shareDatasetKustoDatabase and updates it. Returns the server's representation of the shareDatasetKustoDatabase, and an error, if there is any.
func (c *shareDatasetKustoDatabases) Update(ctx context.Context, shareDatasetKustoDatabase *v1alpha1.ShareDatasetKustoDatabase, opts v1.UpdateOptions) (result *v1alpha1.ShareDatasetKustoDatabase, err error) {
	result = &v1alpha1.ShareDatasetKustoDatabase{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sharedatasetkustodatabases").
		Name(shareDatasetKustoDatabase.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(shareDatasetKustoDatabase).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *shareDatasetKustoDatabases) UpdateStatus(ctx context.Context, shareDatasetKustoDatabase *v1alpha1.ShareDatasetKustoDatabase, opts v1.UpdateOptions) (result *v1alpha1.ShareDatasetKustoDatabase, err error) {
	result = &v1alpha1.ShareDatasetKustoDatabase{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sharedatasetkustodatabases").
		Name(shareDatasetKustoDatabase.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(shareDatasetKustoDatabase).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the shareDatasetKustoDatabase and deletes it. Returns an error if one occurs.
func (c *shareDatasetKustoDatabases) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sharedatasetkustodatabases").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *shareDatasetKustoDatabases) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sharedatasetkustodatabases").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched shareDatasetKustoDatabase.
func (c *shareDatasetKustoDatabases) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ShareDatasetKustoDatabase, err error) {
	result = &v1alpha1.ShareDatasetKustoDatabase{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sharedatasetkustodatabases").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}