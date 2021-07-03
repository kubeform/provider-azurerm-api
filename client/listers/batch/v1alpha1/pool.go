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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/batch/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PoolLister helps list Pools.
// All objects returned here must be treated as read-only.
type PoolLister interface {
	// List lists all Pools in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Pool, err error)
	// Pools returns an object that can list and get Pools.
	Pools(namespace string) PoolNamespaceLister
	PoolListerExpansion
}

// poolLister implements the PoolLister interface.
type poolLister struct {
	indexer cache.Indexer
}

// NewPoolLister returns a new PoolLister.
func NewPoolLister(indexer cache.Indexer) PoolLister {
	return &poolLister{indexer: indexer}
}

// List lists all Pools in the indexer.
func (s *poolLister) List(selector labels.Selector) (ret []*v1alpha1.Pool, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Pool))
	})
	return ret, err
}

// Pools returns an object that can list and get Pools.
func (s *poolLister) Pools(namespace string) PoolNamespaceLister {
	return poolNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PoolNamespaceLister helps list and get Pools.
// All objects returned here must be treated as read-only.
type PoolNamespaceLister interface {
	// List lists all Pools in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Pool, err error)
	// Get retrieves the Pool from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Pool, error)
	PoolNamespaceListerExpansion
}

// poolNamespaceLister implements the PoolNamespaceLister
// interface.
type poolNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Pools in the indexer for a given namespace.
func (s poolNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Pool, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Pool))
	})
	return ret, err
}

// Get retrieves the Pool from the indexer for a given namespace and name.
func (s poolNamespaceLister) Get(name string) (*v1alpha1.Pool, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("pool"), name)
	}
	return obj.(*v1alpha1.Pool), nil
}
