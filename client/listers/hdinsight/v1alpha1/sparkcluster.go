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

// SparkClusterLister helps list SparkClusters.
// All objects returned here must be treated as read-only.
type SparkClusterLister interface {
	// List lists all SparkClusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SparkCluster, err error)
	// SparkClusters returns an object that can list and get SparkClusters.
	SparkClusters(namespace string) SparkClusterNamespaceLister
	SparkClusterListerExpansion
}

// sparkClusterLister implements the SparkClusterLister interface.
type sparkClusterLister struct {
	indexer cache.Indexer
}

// NewSparkClusterLister returns a new SparkClusterLister.
func NewSparkClusterLister(indexer cache.Indexer) SparkClusterLister {
	return &sparkClusterLister{indexer: indexer}
}

// List lists all SparkClusters in the indexer.
func (s *sparkClusterLister) List(selector labels.Selector) (ret []*v1alpha1.SparkCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SparkCluster))
	})
	return ret, err
}

// SparkClusters returns an object that can list and get SparkClusters.
func (s *sparkClusterLister) SparkClusters(namespace string) SparkClusterNamespaceLister {
	return sparkClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SparkClusterNamespaceLister helps list and get SparkClusters.
// All objects returned here must be treated as read-only.
type SparkClusterNamespaceLister interface {
	// List lists all SparkClusters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SparkCluster, err error)
	// Get retrieves the SparkCluster from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SparkCluster, error)
	SparkClusterNamespaceListerExpansion
}

// sparkClusterNamespaceLister implements the SparkClusterNamespaceLister
// interface.
type sparkClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SparkClusters in the indexer for a given namespace.
func (s sparkClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SparkCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SparkCluster))
	})
	return ret, err
}

// Get retrieves the SparkCluster from the indexer for a given namespace and name.
func (s sparkClusterNamespaceLister) Get(name string) (*v1alpha1.SparkCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sparkcluster"), name)
	}
	return obj.(*v1alpha1.SparkCluster), nil
}
