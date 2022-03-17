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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/sentinel/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WatchlistItemLister helps list WatchlistItems.
// All objects returned here must be treated as read-only.
type WatchlistItemLister interface {
	// List lists all WatchlistItems in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WatchlistItem, err error)
	// WatchlistItems returns an object that can list and get WatchlistItems.
	WatchlistItems(namespace string) WatchlistItemNamespaceLister
	WatchlistItemListerExpansion
}

// watchlistItemLister implements the WatchlistItemLister interface.
type watchlistItemLister struct {
	indexer cache.Indexer
}

// NewWatchlistItemLister returns a new WatchlistItemLister.
func NewWatchlistItemLister(indexer cache.Indexer) WatchlistItemLister {
	return &watchlistItemLister{indexer: indexer}
}

// List lists all WatchlistItems in the indexer.
func (s *watchlistItemLister) List(selector labels.Selector) (ret []*v1alpha1.WatchlistItem, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WatchlistItem))
	})
	return ret, err
}

// WatchlistItems returns an object that can list and get WatchlistItems.
func (s *watchlistItemLister) WatchlistItems(namespace string) WatchlistItemNamespaceLister {
	return watchlistItemNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WatchlistItemNamespaceLister helps list and get WatchlistItems.
// All objects returned here must be treated as read-only.
type WatchlistItemNamespaceLister interface {
	// List lists all WatchlistItems in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WatchlistItem, err error)
	// Get retrieves the WatchlistItem from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.WatchlistItem, error)
	WatchlistItemNamespaceListerExpansion
}

// watchlistItemNamespaceLister implements the WatchlistItemNamespaceLister
// interface.
type watchlistItemNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WatchlistItems in the indexer for a given namespace.
func (s watchlistItemNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WatchlistItem, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WatchlistItem))
	})
	return ret, err
}

// Get retrieves the WatchlistItem from the indexer for a given namespace and name.
func (s watchlistItemNamespaceLister) Get(name string) (*v1alpha1.WatchlistItem, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("watchlistitem"), name)
	}
	return obj.(*v1alpha1.WatchlistItem), nil
}