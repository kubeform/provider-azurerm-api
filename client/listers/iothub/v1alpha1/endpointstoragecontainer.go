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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/iothub/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EndpointStorageContainerLister helps list EndpointStorageContainers.
// All objects returned here must be treated as read-only.
type EndpointStorageContainerLister interface {
	// List lists all EndpointStorageContainers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EndpointStorageContainer, err error)
	// EndpointStorageContainers returns an object that can list and get EndpointStorageContainers.
	EndpointStorageContainers(namespace string) EndpointStorageContainerNamespaceLister
	EndpointStorageContainerListerExpansion
}

// endpointStorageContainerLister implements the EndpointStorageContainerLister interface.
type endpointStorageContainerLister struct {
	indexer cache.Indexer
}

// NewEndpointStorageContainerLister returns a new EndpointStorageContainerLister.
func NewEndpointStorageContainerLister(indexer cache.Indexer) EndpointStorageContainerLister {
	return &endpointStorageContainerLister{indexer: indexer}
}

// List lists all EndpointStorageContainers in the indexer.
func (s *endpointStorageContainerLister) List(selector labels.Selector) (ret []*v1alpha1.EndpointStorageContainer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EndpointStorageContainer))
	})
	return ret, err
}

// EndpointStorageContainers returns an object that can list and get EndpointStorageContainers.
func (s *endpointStorageContainerLister) EndpointStorageContainers(namespace string) EndpointStorageContainerNamespaceLister {
	return endpointStorageContainerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EndpointStorageContainerNamespaceLister helps list and get EndpointStorageContainers.
// All objects returned here must be treated as read-only.
type EndpointStorageContainerNamespaceLister interface {
	// List lists all EndpointStorageContainers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EndpointStorageContainer, err error)
	// Get retrieves the EndpointStorageContainer from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EndpointStorageContainer, error)
	EndpointStorageContainerNamespaceListerExpansion
}

// endpointStorageContainerNamespaceLister implements the EndpointStorageContainerNamespaceLister
// interface.
type endpointStorageContainerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EndpointStorageContainers in the indexer for a given namespace.
func (s endpointStorageContainerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EndpointStorageContainer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EndpointStorageContainer))
	})
	return ret, err
}

// Get retrieves the EndpointStorageContainer from the indexer for a given namespace and name.
func (s endpointStorageContainerNamespaceLister) Get(name string) (*v1alpha1.EndpointStorageContainer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("endpointstoragecontainer"), name)
	}
	return obj.(*v1alpha1.EndpointStorageContainer), nil
}
