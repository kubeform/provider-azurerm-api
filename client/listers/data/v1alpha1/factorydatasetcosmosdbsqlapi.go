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

// FactoryDatasetCosmosdbSqlapiLister helps list FactoryDatasetCosmosdbSqlapis.
// All objects returned here must be treated as read-only.
type FactoryDatasetCosmosdbSqlapiLister interface {
	// List lists all FactoryDatasetCosmosdbSqlapis in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FactoryDatasetCosmosdbSqlapi, err error)
	// FactoryDatasetCosmosdbSqlapis returns an object that can list and get FactoryDatasetCosmosdbSqlapis.
	FactoryDatasetCosmosdbSqlapis(namespace string) FactoryDatasetCosmosdbSqlapiNamespaceLister
	FactoryDatasetCosmosdbSqlapiListerExpansion
}

// factoryDatasetCosmosdbSqlapiLister implements the FactoryDatasetCosmosdbSqlapiLister interface.
type factoryDatasetCosmosdbSqlapiLister struct {
	indexer cache.Indexer
}

// NewFactoryDatasetCosmosdbSqlapiLister returns a new FactoryDatasetCosmosdbSqlapiLister.
func NewFactoryDatasetCosmosdbSqlapiLister(indexer cache.Indexer) FactoryDatasetCosmosdbSqlapiLister {
	return &factoryDatasetCosmosdbSqlapiLister{indexer: indexer}
}

// List lists all FactoryDatasetCosmosdbSqlapis in the indexer.
func (s *factoryDatasetCosmosdbSqlapiLister) List(selector labels.Selector) (ret []*v1alpha1.FactoryDatasetCosmosdbSqlapi, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FactoryDatasetCosmosdbSqlapi))
	})
	return ret, err
}

// FactoryDatasetCosmosdbSqlapis returns an object that can list and get FactoryDatasetCosmosdbSqlapis.
func (s *factoryDatasetCosmosdbSqlapiLister) FactoryDatasetCosmosdbSqlapis(namespace string) FactoryDatasetCosmosdbSqlapiNamespaceLister {
	return factoryDatasetCosmosdbSqlapiNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FactoryDatasetCosmosdbSqlapiNamespaceLister helps list and get FactoryDatasetCosmosdbSqlapis.
// All objects returned here must be treated as read-only.
type FactoryDatasetCosmosdbSqlapiNamespaceLister interface {
	// List lists all FactoryDatasetCosmosdbSqlapis in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FactoryDatasetCosmosdbSqlapi, err error)
	// Get retrieves the FactoryDatasetCosmosdbSqlapi from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FactoryDatasetCosmosdbSqlapi, error)
	FactoryDatasetCosmosdbSqlapiNamespaceListerExpansion
}

// factoryDatasetCosmosdbSqlapiNamespaceLister implements the FactoryDatasetCosmosdbSqlapiNamespaceLister
// interface.
type factoryDatasetCosmosdbSqlapiNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FactoryDatasetCosmosdbSqlapis in the indexer for a given namespace.
func (s factoryDatasetCosmosdbSqlapiNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FactoryDatasetCosmosdbSqlapi, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FactoryDatasetCosmosdbSqlapi))
	})
	return ret, err
}

// Get retrieves the FactoryDatasetCosmosdbSqlapi from the indexer for a given namespace and name.
func (s factoryDatasetCosmosdbSqlapiNamespaceLister) Get(name string) (*v1alpha1.FactoryDatasetCosmosdbSqlapi, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("factorydatasetcosmosdbsqlapi"), name)
	}
	return obj.(*v1alpha1.FactoryDatasetCosmosdbSqlapi), nil
}