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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/hdinsight/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InteractiveQueryClusterLister helps list InteractiveQueryClusters.
// All objects returned here must be treated as read-only.
type InteractiveQueryClusterLister interface {
	// List lists all InteractiveQueryClusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InteractiveQueryCluster, err error)
	// InteractiveQueryClusters returns an object that can list and get InteractiveQueryClusters.
	InteractiveQueryClusters(namespace string) InteractiveQueryClusterNamespaceLister
	InteractiveQueryClusterListerExpansion
}

// interactiveQueryClusterLister implements the InteractiveQueryClusterLister interface.
type interactiveQueryClusterLister struct {
	indexer cache.Indexer
}

// NewInteractiveQueryClusterLister returns a new InteractiveQueryClusterLister.
func NewInteractiveQueryClusterLister(indexer cache.Indexer) InteractiveQueryClusterLister {
	return &interactiveQueryClusterLister{indexer: indexer}
}

// List lists all InteractiveQueryClusters in the indexer.
func (s *interactiveQueryClusterLister) List(selector labels.Selector) (ret []*v1alpha1.InteractiveQueryCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InteractiveQueryCluster))
	})
	return ret, err
}

// InteractiveQueryClusters returns an object that can list and get InteractiveQueryClusters.
func (s *interactiveQueryClusterLister) InteractiveQueryClusters(namespace string) InteractiveQueryClusterNamespaceLister {
	return interactiveQueryClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InteractiveQueryClusterNamespaceLister helps list and get InteractiveQueryClusters.
// All objects returned here must be treated as read-only.
type InteractiveQueryClusterNamespaceLister interface {
	// List lists all InteractiveQueryClusters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InteractiveQueryCluster, err error)
	// Get retrieves the InteractiveQueryCluster from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.InteractiveQueryCluster, error)
	InteractiveQueryClusterNamespaceListerExpansion
}

// interactiveQueryClusterNamespaceLister implements the InteractiveQueryClusterNamespaceLister
// interface.
type interactiveQueryClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all InteractiveQueryClusters in the indexer for a given namespace.
func (s interactiveQueryClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.InteractiveQueryCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InteractiveQueryCluster))
	})
	return ret, err
}

// Get retrieves the InteractiveQueryCluster from the indexer for a given namespace and name.
func (s interactiveQueryClusterNamespaceLister) Get(name string) (*v1alpha1.InteractiveQueryCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("interactivequerycluster"), name)
	}
	return obj.(*v1alpha1.InteractiveQueryCluster), nil
}
