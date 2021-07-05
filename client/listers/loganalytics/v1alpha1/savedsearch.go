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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/loganalytics/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SavedSearchLister helps list SavedSearches.
// All objects returned here must be treated as read-only.
type SavedSearchLister interface {
	// List lists all SavedSearches in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SavedSearch, err error)
	// SavedSearches returns an object that can list and get SavedSearches.
	SavedSearches(namespace string) SavedSearchNamespaceLister
	SavedSearchListerExpansion
}

// savedSearchLister implements the SavedSearchLister interface.
type savedSearchLister struct {
	indexer cache.Indexer
}

// NewSavedSearchLister returns a new SavedSearchLister.
func NewSavedSearchLister(indexer cache.Indexer) SavedSearchLister {
	return &savedSearchLister{indexer: indexer}
}

// List lists all SavedSearches in the indexer.
func (s *savedSearchLister) List(selector labels.Selector) (ret []*v1alpha1.SavedSearch, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SavedSearch))
	})
	return ret, err
}

// SavedSearches returns an object that can list and get SavedSearches.
func (s *savedSearchLister) SavedSearches(namespace string) SavedSearchNamespaceLister {
	return savedSearchNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SavedSearchNamespaceLister helps list and get SavedSearches.
// All objects returned here must be treated as read-only.
type SavedSearchNamespaceLister interface {
	// List lists all SavedSearches in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SavedSearch, err error)
	// Get retrieves the SavedSearch from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SavedSearch, error)
	SavedSearchNamespaceListerExpansion
}

// savedSearchNamespaceLister implements the SavedSearchNamespaceLister
// interface.
type savedSearchNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SavedSearches in the indexer for a given namespace.
func (s savedSearchNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SavedSearch, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SavedSearch))
	})
	return ret, err
}

// Get retrieves the SavedSearch from the indexer for a given namespace and name.
func (s savedSearchNamespaceLister) Get(name string) (*v1alpha1.SavedSearch, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("savedsearch"), name)
	}
	return obj.(*v1alpha1.SavedSearch), nil
}