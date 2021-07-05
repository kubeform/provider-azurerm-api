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

// StorageInsightsLister helps list StorageInsightses.
// All objects returned here must be treated as read-only.
type StorageInsightsLister interface {
	// List lists all StorageInsightses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StorageInsights, err error)
	// StorageInsightses returns an object that can list and get StorageInsightses.
	StorageInsightses(namespace string) StorageInsightsNamespaceLister
	StorageInsightsListerExpansion
}

// storageInsightsLister implements the StorageInsightsLister interface.
type storageInsightsLister struct {
	indexer cache.Indexer
}

// NewStorageInsightsLister returns a new StorageInsightsLister.
func NewStorageInsightsLister(indexer cache.Indexer) StorageInsightsLister {
	return &storageInsightsLister{indexer: indexer}
}

// List lists all StorageInsightses in the indexer.
func (s *storageInsightsLister) List(selector labels.Selector) (ret []*v1alpha1.StorageInsights, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageInsights))
	})
	return ret, err
}

// StorageInsightses returns an object that can list and get StorageInsightses.
func (s *storageInsightsLister) StorageInsightses(namespace string) StorageInsightsNamespaceLister {
	return storageInsightsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StorageInsightsNamespaceLister helps list and get StorageInsightses.
// All objects returned here must be treated as read-only.
type StorageInsightsNamespaceLister interface {
	// List lists all StorageInsightses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StorageInsights, err error)
	// Get retrieves the StorageInsights from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.StorageInsights, error)
	StorageInsightsNamespaceListerExpansion
}

// storageInsightsNamespaceLister implements the StorageInsightsNamespaceLister
// interface.
type storageInsightsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StorageInsightses in the indexer for a given namespace.
func (s storageInsightsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StorageInsights, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageInsights))
	})
	return ret, err
}

// Get retrieves the StorageInsights from the indexer for a given namespace and name.
func (s storageInsightsNamespaceLister) Get(name string) (*v1alpha1.StorageInsights, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("storageinsights"), name)
	}
	return obj.(*v1alpha1.StorageInsights), nil
}