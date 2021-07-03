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

// StreamingLocatorLister helps list StreamingLocators.
// All objects returned here must be treated as read-only.
type StreamingLocatorLister interface {
	// List lists all StreamingLocators in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StreamingLocator, err error)
	// StreamingLocators returns an object that can list and get StreamingLocators.
	StreamingLocators(namespace string) StreamingLocatorNamespaceLister
	StreamingLocatorListerExpansion
}

// streamingLocatorLister implements the StreamingLocatorLister interface.
type streamingLocatorLister struct {
	indexer cache.Indexer
}

// NewStreamingLocatorLister returns a new StreamingLocatorLister.
func NewStreamingLocatorLister(indexer cache.Indexer) StreamingLocatorLister {
	return &streamingLocatorLister{indexer: indexer}
}

// List lists all StreamingLocators in the indexer.
func (s *streamingLocatorLister) List(selector labels.Selector) (ret []*v1alpha1.StreamingLocator, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StreamingLocator))
	})
	return ret, err
}

// StreamingLocators returns an object that can list and get StreamingLocators.
func (s *streamingLocatorLister) StreamingLocators(namespace string) StreamingLocatorNamespaceLister {
	return streamingLocatorNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StreamingLocatorNamespaceLister helps list and get StreamingLocators.
// All objects returned here must be treated as read-only.
type StreamingLocatorNamespaceLister interface {
	// List lists all StreamingLocators in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StreamingLocator, err error)
	// Get retrieves the StreamingLocator from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.StreamingLocator, error)
	StreamingLocatorNamespaceListerExpansion
}

// streamingLocatorNamespaceLister implements the StreamingLocatorNamespaceLister
// interface.
type streamingLocatorNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StreamingLocators in the indexer for a given namespace.
func (s streamingLocatorNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StreamingLocator, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StreamingLocator))
	})
	return ret, err
}

// Get retrieves the StreamingLocator from the indexer for a given namespace and name.
func (s streamingLocatorNamespaceLister) Get(name string) (*v1alpha1.StreamingLocator, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("streaminglocator"), name)
	}
	return obj.(*v1alpha1.StreamingLocator), nil
}
