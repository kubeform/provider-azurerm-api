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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/service/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FabricManagedClusterLister helps list FabricManagedClusters.
// All objects returned here must be treated as read-only.
type FabricManagedClusterLister interface {
	// List lists all FabricManagedClusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FabricManagedCluster, err error)
	// FabricManagedClusters returns an object that can list and get FabricManagedClusters.
	FabricManagedClusters(namespace string) FabricManagedClusterNamespaceLister
	FabricManagedClusterListerExpansion
}

// fabricManagedClusterLister implements the FabricManagedClusterLister interface.
type fabricManagedClusterLister struct {
	indexer cache.Indexer
}

// NewFabricManagedClusterLister returns a new FabricManagedClusterLister.
func NewFabricManagedClusterLister(indexer cache.Indexer) FabricManagedClusterLister {
	return &fabricManagedClusterLister{indexer: indexer}
}

// List lists all FabricManagedClusters in the indexer.
func (s *fabricManagedClusterLister) List(selector labels.Selector) (ret []*v1alpha1.FabricManagedCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FabricManagedCluster))
	})
	return ret, err
}

// FabricManagedClusters returns an object that can list and get FabricManagedClusters.
func (s *fabricManagedClusterLister) FabricManagedClusters(namespace string) FabricManagedClusterNamespaceLister {
	return fabricManagedClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FabricManagedClusterNamespaceLister helps list and get FabricManagedClusters.
// All objects returned here must be treated as read-only.
type FabricManagedClusterNamespaceLister interface {
	// List lists all FabricManagedClusters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FabricManagedCluster, err error)
	// Get retrieves the FabricManagedCluster from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FabricManagedCluster, error)
	FabricManagedClusterNamespaceListerExpansion
}

// fabricManagedClusterNamespaceLister implements the FabricManagedClusterNamespaceLister
// interface.
type fabricManagedClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FabricManagedClusters in the indexer for a given namespace.
func (s fabricManagedClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FabricManagedCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FabricManagedCluster))
	})
	return ret, err
}

// Get retrieves the FabricManagedCluster from the indexer for a given namespace and name.
func (s fabricManagedClusterNamespaceLister) Get(name string) (*v1alpha1.FabricManagedCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("fabricmanagedcluster"), name)
	}
	return obj.(*v1alpha1.FabricManagedCluster), nil
}
