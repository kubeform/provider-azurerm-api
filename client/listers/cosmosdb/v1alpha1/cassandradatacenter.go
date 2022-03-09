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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/cosmosdb/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CassandraDatacenterLister helps list CassandraDatacenters.
// All objects returned here must be treated as read-only.
type CassandraDatacenterLister interface {
	// List lists all CassandraDatacenters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CassandraDatacenter, err error)
	// CassandraDatacenters returns an object that can list and get CassandraDatacenters.
	CassandraDatacenters(namespace string) CassandraDatacenterNamespaceLister
	CassandraDatacenterListerExpansion
}

// cassandraDatacenterLister implements the CassandraDatacenterLister interface.
type cassandraDatacenterLister struct {
	indexer cache.Indexer
}

// NewCassandraDatacenterLister returns a new CassandraDatacenterLister.
func NewCassandraDatacenterLister(indexer cache.Indexer) CassandraDatacenterLister {
	return &cassandraDatacenterLister{indexer: indexer}
}

// List lists all CassandraDatacenters in the indexer.
func (s *cassandraDatacenterLister) List(selector labels.Selector) (ret []*v1alpha1.CassandraDatacenter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CassandraDatacenter))
	})
	return ret, err
}

// CassandraDatacenters returns an object that can list and get CassandraDatacenters.
func (s *cassandraDatacenterLister) CassandraDatacenters(namespace string) CassandraDatacenterNamespaceLister {
	return cassandraDatacenterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CassandraDatacenterNamespaceLister helps list and get CassandraDatacenters.
// All objects returned here must be treated as read-only.
type CassandraDatacenterNamespaceLister interface {
	// List lists all CassandraDatacenters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CassandraDatacenter, err error)
	// Get retrieves the CassandraDatacenter from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CassandraDatacenter, error)
	CassandraDatacenterNamespaceListerExpansion
}

// cassandraDatacenterNamespaceLister implements the CassandraDatacenterNamespaceLister
// interface.
type cassandraDatacenterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CassandraDatacenters in the indexer for a given namespace.
func (s cassandraDatacenterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CassandraDatacenter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CassandraDatacenter))
	})
	return ret, err
}

// Get retrieves the CassandraDatacenter from the indexer for a given namespace and name.
func (s cassandraDatacenterNamespaceLister) Get(name string) (*v1alpha1.CassandraDatacenter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cassandradatacenter"), name)
	}
	return obj.(*v1alpha1.CassandraDatacenter), nil
}
