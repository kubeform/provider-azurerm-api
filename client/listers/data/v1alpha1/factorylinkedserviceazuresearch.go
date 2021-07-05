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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/data/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FactoryLinkedServiceAzureSearchLister helps list FactoryLinkedServiceAzureSearches.
// All objects returned here must be treated as read-only.
type FactoryLinkedServiceAzureSearchLister interface {
	// List lists all FactoryLinkedServiceAzureSearches in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FactoryLinkedServiceAzureSearch, err error)
	// FactoryLinkedServiceAzureSearches returns an object that can list and get FactoryLinkedServiceAzureSearches.
	FactoryLinkedServiceAzureSearches(namespace string) FactoryLinkedServiceAzureSearchNamespaceLister
	FactoryLinkedServiceAzureSearchListerExpansion
}

// factoryLinkedServiceAzureSearchLister implements the FactoryLinkedServiceAzureSearchLister interface.
type factoryLinkedServiceAzureSearchLister struct {
	indexer cache.Indexer
}

// NewFactoryLinkedServiceAzureSearchLister returns a new FactoryLinkedServiceAzureSearchLister.
func NewFactoryLinkedServiceAzureSearchLister(indexer cache.Indexer) FactoryLinkedServiceAzureSearchLister {
	return &factoryLinkedServiceAzureSearchLister{indexer: indexer}
}

// List lists all FactoryLinkedServiceAzureSearches in the indexer.
func (s *factoryLinkedServiceAzureSearchLister) List(selector labels.Selector) (ret []*v1alpha1.FactoryLinkedServiceAzureSearch, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FactoryLinkedServiceAzureSearch))
	})
	return ret, err
}

// FactoryLinkedServiceAzureSearches returns an object that can list and get FactoryLinkedServiceAzureSearches.
func (s *factoryLinkedServiceAzureSearchLister) FactoryLinkedServiceAzureSearches(namespace string) FactoryLinkedServiceAzureSearchNamespaceLister {
	return factoryLinkedServiceAzureSearchNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FactoryLinkedServiceAzureSearchNamespaceLister helps list and get FactoryLinkedServiceAzureSearches.
// All objects returned here must be treated as read-only.
type FactoryLinkedServiceAzureSearchNamespaceLister interface {
	// List lists all FactoryLinkedServiceAzureSearches in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FactoryLinkedServiceAzureSearch, err error)
	// Get retrieves the FactoryLinkedServiceAzureSearch from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FactoryLinkedServiceAzureSearch, error)
	FactoryLinkedServiceAzureSearchNamespaceListerExpansion
}

// factoryLinkedServiceAzureSearchNamespaceLister implements the FactoryLinkedServiceAzureSearchNamespaceLister
// interface.
type factoryLinkedServiceAzureSearchNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FactoryLinkedServiceAzureSearches in the indexer for a given namespace.
func (s factoryLinkedServiceAzureSearchNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FactoryLinkedServiceAzureSearch, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FactoryLinkedServiceAzureSearch))
	})
	return ret, err
}

// Get retrieves the FactoryLinkedServiceAzureSearch from the indexer for a given namespace and name.
func (s factoryLinkedServiceAzureSearchNamespaceLister) Get(name string) (*v1alpha1.FactoryLinkedServiceAzureSearch, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("factorylinkedserviceazuresearch"), name)
	}
	return obj.(*v1alpha1.FactoryLinkedServiceAzureSearch), nil
}