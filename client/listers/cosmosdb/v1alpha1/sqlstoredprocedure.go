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

// SqlStoredProcedureLister helps list SqlStoredProcedures.
// All objects returned here must be treated as read-only.
type SqlStoredProcedureLister interface {
	// List lists all SqlStoredProcedures in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SqlStoredProcedure, err error)
	// SqlStoredProcedures returns an object that can list and get SqlStoredProcedures.
	SqlStoredProcedures(namespace string) SqlStoredProcedureNamespaceLister
	SqlStoredProcedureListerExpansion
}

// sqlStoredProcedureLister implements the SqlStoredProcedureLister interface.
type sqlStoredProcedureLister struct {
	indexer cache.Indexer
}

// NewSqlStoredProcedureLister returns a new SqlStoredProcedureLister.
func NewSqlStoredProcedureLister(indexer cache.Indexer) SqlStoredProcedureLister {
	return &sqlStoredProcedureLister{indexer: indexer}
}

// List lists all SqlStoredProcedures in the indexer.
func (s *sqlStoredProcedureLister) List(selector labels.Selector) (ret []*v1alpha1.SqlStoredProcedure, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SqlStoredProcedure))
	})
	return ret, err
}

// SqlStoredProcedures returns an object that can list and get SqlStoredProcedures.
func (s *sqlStoredProcedureLister) SqlStoredProcedures(namespace string) SqlStoredProcedureNamespaceLister {
	return sqlStoredProcedureNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SqlStoredProcedureNamespaceLister helps list and get SqlStoredProcedures.
// All objects returned here must be treated as read-only.
type SqlStoredProcedureNamespaceLister interface {
	// List lists all SqlStoredProcedures in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SqlStoredProcedure, err error)
	// Get retrieves the SqlStoredProcedure from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SqlStoredProcedure, error)
	SqlStoredProcedureNamespaceListerExpansion
}

// sqlStoredProcedureNamespaceLister implements the SqlStoredProcedureNamespaceLister
// interface.
type sqlStoredProcedureNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SqlStoredProcedures in the indexer for a given namespace.
func (s sqlStoredProcedureNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SqlStoredProcedure, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SqlStoredProcedure))
	})
	return ret, err
}

// Get retrieves the SqlStoredProcedure from the indexer for a given namespace and name.
func (s sqlStoredProcedureNamespaceLister) Get(name string) (*v1alpha1.SqlStoredProcedure, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sqlstoredprocedure"), name)
	}
	return obj.(*v1alpha1.SqlStoredProcedure), nil
}
