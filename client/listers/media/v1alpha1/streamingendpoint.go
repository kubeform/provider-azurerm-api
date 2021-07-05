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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/media/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StreamingEndpointLister helps list StreamingEndpoints.
// All objects returned here must be treated as read-only.
type StreamingEndpointLister interface {
	// List lists all StreamingEndpoints in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StreamingEndpoint, err error)
	// StreamingEndpoints returns an object that can list and get StreamingEndpoints.
	StreamingEndpoints(namespace string) StreamingEndpointNamespaceLister
	StreamingEndpointListerExpansion
}

// streamingEndpointLister implements the StreamingEndpointLister interface.
type streamingEndpointLister struct {
	indexer cache.Indexer
}

// NewStreamingEndpointLister returns a new StreamingEndpointLister.
func NewStreamingEndpointLister(indexer cache.Indexer) StreamingEndpointLister {
	return &streamingEndpointLister{indexer: indexer}
}

// List lists all StreamingEndpoints in the indexer.
func (s *streamingEndpointLister) List(selector labels.Selector) (ret []*v1alpha1.StreamingEndpoint, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StreamingEndpoint))
	})
	return ret, err
}

// StreamingEndpoints returns an object that can list and get StreamingEndpoints.
func (s *streamingEndpointLister) StreamingEndpoints(namespace string) StreamingEndpointNamespaceLister {
	return streamingEndpointNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StreamingEndpointNamespaceLister helps list and get StreamingEndpoints.
// All objects returned here must be treated as read-only.
type StreamingEndpointNamespaceLister interface {
	// List lists all StreamingEndpoints in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StreamingEndpoint, err error)
	// Get retrieves the StreamingEndpoint from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.StreamingEndpoint, error)
	StreamingEndpointNamespaceListerExpansion
}

// streamingEndpointNamespaceLister implements the StreamingEndpointNamespaceLister
// interface.
type streamingEndpointNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StreamingEndpoints in the indexer for a given namespace.
func (s streamingEndpointNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StreamingEndpoint, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StreamingEndpoint))
	})
	return ret, err
}

// Get retrieves the StreamingEndpoint from the indexer for a given namespace and name.
func (s streamingEndpointNamespaceLister) Get(name string) (*v1alpha1.StreamingEndpoint, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("streamingendpoint"), name)
	}
	return obj.(*v1alpha1.StreamingEndpoint), nil
}