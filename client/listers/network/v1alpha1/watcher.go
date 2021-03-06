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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/network/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WatcherLister helps list Watchers.
// All objects returned here must be treated as read-only.
type WatcherLister interface {
	// List lists all Watchers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Watcher, err error)
	// Watchers returns an object that can list and get Watchers.
	Watchers(namespace string) WatcherNamespaceLister
	WatcherListerExpansion
}

// watcherLister implements the WatcherLister interface.
type watcherLister struct {
	indexer cache.Indexer
}

// NewWatcherLister returns a new WatcherLister.
func NewWatcherLister(indexer cache.Indexer) WatcherLister {
	return &watcherLister{indexer: indexer}
}

// List lists all Watchers in the indexer.
func (s *watcherLister) List(selector labels.Selector) (ret []*v1alpha1.Watcher, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Watcher))
	})
	return ret, err
}

// Watchers returns an object that can list and get Watchers.
func (s *watcherLister) Watchers(namespace string) WatcherNamespaceLister {
	return watcherNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WatcherNamespaceLister helps list and get Watchers.
// All objects returned here must be treated as read-only.
type WatcherNamespaceLister interface {
	// List lists all Watchers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Watcher, err error)
	// Get retrieves the Watcher from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Watcher, error)
	WatcherNamespaceListerExpansion
}

// watcherNamespaceLister implements the WatcherNamespaceLister
// interface.
type watcherNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Watchers in the indexer for a given namespace.
func (s watcherNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Watcher, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Watcher))
	})
	return ret, err
}

// Get retrieves the Watcher from the indexer for a given namespace and name.
func (s watcherNamespaceLister) Get(name string) (*v1alpha1.Watcher, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("watcher"), name)
	}
	return obj.(*v1alpha1.Watcher), nil
}
