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

// FactoryLinkedServiceSnowflakeLister helps list FactoryLinkedServiceSnowflakes.
// All objects returned here must be treated as read-only.
type FactoryLinkedServiceSnowflakeLister interface {
	// List lists all FactoryLinkedServiceSnowflakes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FactoryLinkedServiceSnowflake, err error)
	// FactoryLinkedServiceSnowflakes returns an object that can list and get FactoryLinkedServiceSnowflakes.
	FactoryLinkedServiceSnowflakes(namespace string) FactoryLinkedServiceSnowflakeNamespaceLister
	FactoryLinkedServiceSnowflakeListerExpansion
}

// factoryLinkedServiceSnowflakeLister implements the FactoryLinkedServiceSnowflakeLister interface.
type factoryLinkedServiceSnowflakeLister struct {
	indexer cache.Indexer
}

// NewFactoryLinkedServiceSnowflakeLister returns a new FactoryLinkedServiceSnowflakeLister.
func NewFactoryLinkedServiceSnowflakeLister(indexer cache.Indexer) FactoryLinkedServiceSnowflakeLister {
	return &factoryLinkedServiceSnowflakeLister{indexer: indexer}
}

// List lists all FactoryLinkedServiceSnowflakes in the indexer.
func (s *factoryLinkedServiceSnowflakeLister) List(selector labels.Selector) (ret []*v1alpha1.FactoryLinkedServiceSnowflake, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FactoryLinkedServiceSnowflake))
	})
	return ret, err
}

// FactoryLinkedServiceSnowflakes returns an object that can list and get FactoryLinkedServiceSnowflakes.
func (s *factoryLinkedServiceSnowflakeLister) FactoryLinkedServiceSnowflakes(namespace string) FactoryLinkedServiceSnowflakeNamespaceLister {
	return factoryLinkedServiceSnowflakeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FactoryLinkedServiceSnowflakeNamespaceLister helps list and get FactoryLinkedServiceSnowflakes.
// All objects returned here must be treated as read-only.
type FactoryLinkedServiceSnowflakeNamespaceLister interface {
	// List lists all FactoryLinkedServiceSnowflakes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FactoryLinkedServiceSnowflake, err error)
	// Get retrieves the FactoryLinkedServiceSnowflake from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FactoryLinkedServiceSnowflake, error)
	FactoryLinkedServiceSnowflakeNamespaceListerExpansion
}

// factoryLinkedServiceSnowflakeNamespaceLister implements the FactoryLinkedServiceSnowflakeNamespaceLister
// interface.
type factoryLinkedServiceSnowflakeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FactoryLinkedServiceSnowflakes in the indexer for a given namespace.
func (s factoryLinkedServiceSnowflakeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FactoryLinkedServiceSnowflake, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FactoryLinkedServiceSnowflake))
	})
	return ret, err
}

// Get retrieves the FactoryLinkedServiceSnowflake from the indexer for a given namespace and name.
func (s factoryLinkedServiceSnowflakeNamespaceLister) Get(name string) (*v1alpha1.FactoryLinkedServiceSnowflake, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("factorylinkedservicesnowflake"), name)
	}
	return obj.(*v1alpha1.FactoryLinkedServiceSnowflake), nil
}