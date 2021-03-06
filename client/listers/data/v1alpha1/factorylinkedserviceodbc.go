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

// FactoryLinkedServiceOdbcLister helps list FactoryLinkedServiceOdbcs.
// All objects returned here must be treated as read-only.
type FactoryLinkedServiceOdbcLister interface {
	// List lists all FactoryLinkedServiceOdbcs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FactoryLinkedServiceOdbc, err error)
	// FactoryLinkedServiceOdbcs returns an object that can list and get FactoryLinkedServiceOdbcs.
	FactoryLinkedServiceOdbcs(namespace string) FactoryLinkedServiceOdbcNamespaceLister
	FactoryLinkedServiceOdbcListerExpansion
}

// factoryLinkedServiceOdbcLister implements the FactoryLinkedServiceOdbcLister interface.
type factoryLinkedServiceOdbcLister struct {
	indexer cache.Indexer
}

// NewFactoryLinkedServiceOdbcLister returns a new FactoryLinkedServiceOdbcLister.
func NewFactoryLinkedServiceOdbcLister(indexer cache.Indexer) FactoryLinkedServiceOdbcLister {
	return &factoryLinkedServiceOdbcLister{indexer: indexer}
}

// List lists all FactoryLinkedServiceOdbcs in the indexer.
func (s *factoryLinkedServiceOdbcLister) List(selector labels.Selector) (ret []*v1alpha1.FactoryLinkedServiceOdbc, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FactoryLinkedServiceOdbc))
	})
	return ret, err
}

// FactoryLinkedServiceOdbcs returns an object that can list and get FactoryLinkedServiceOdbcs.
func (s *factoryLinkedServiceOdbcLister) FactoryLinkedServiceOdbcs(namespace string) FactoryLinkedServiceOdbcNamespaceLister {
	return factoryLinkedServiceOdbcNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FactoryLinkedServiceOdbcNamespaceLister helps list and get FactoryLinkedServiceOdbcs.
// All objects returned here must be treated as read-only.
type FactoryLinkedServiceOdbcNamespaceLister interface {
	// List lists all FactoryLinkedServiceOdbcs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FactoryLinkedServiceOdbc, err error)
	// Get retrieves the FactoryLinkedServiceOdbc from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FactoryLinkedServiceOdbc, error)
	FactoryLinkedServiceOdbcNamespaceListerExpansion
}

// factoryLinkedServiceOdbcNamespaceLister implements the FactoryLinkedServiceOdbcNamespaceLister
// interface.
type factoryLinkedServiceOdbcNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FactoryLinkedServiceOdbcs in the indexer for a given namespace.
func (s factoryLinkedServiceOdbcNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FactoryLinkedServiceOdbc, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FactoryLinkedServiceOdbc))
	})
	return ret, err
}

// Get retrieves the FactoryLinkedServiceOdbc from the indexer for a given namespace and name.
func (s factoryLinkedServiceOdbcNamespaceLister) Get(name string) (*v1alpha1.FactoryLinkedServiceOdbc, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("factorylinkedserviceodbc"), name)
	}
	return obj.(*v1alpha1.FactoryLinkedServiceOdbc), nil
}
