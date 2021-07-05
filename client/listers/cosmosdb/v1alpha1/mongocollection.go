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

// MongoCollectionLister helps list MongoCollections.
// All objects returned here must be treated as read-only.
type MongoCollectionLister interface {
	// List lists all MongoCollections in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MongoCollection, err error)
	// MongoCollections returns an object that can list and get MongoCollections.
	MongoCollections(namespace string) MongoCollectionNamespaceLister
	MongoCollectionListerExpansion
}

// mongoCollectionLister implements the MongoCollectionLister interface.
type mongoCollectionLister struct {
	indexer cache.Indexer
}

// NewMongoCollectionLister returns a new MongoCollectionLister.
func NewMongoCollectionLister(indexer cache.Indexer) MongoCollectionLister {
	return &mongoCollectionLister{indexer: indexer}
}

// List lists all MongoCollections in the indexer.
func (s *mongoCollectionLister) List(selector labels.Selector) (ret []*v1alpha1.MongoCollection, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MongoCollection))
	})
	return ret, err
}

// MongoCollections returns an object that can list and get MongoCollections.
func (s *mongoCollectionLister) MongoCollections(namespace string) MongoCollectionNamespaceLister {
	return mongoCollectionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MongoCollectionNamespaceLister helps list and get MongoCollections.
// All objects returned here must be treated as read-only.
type MongoCollectionNamespaceLister interface {
	// List lists all MongoCollections in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MongoCollection, err error)
	// Get retrieves the MongoCollection from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.MongoCollection, error)
	MongoCollectionNamespaceListerExpansion
}

// mongoCollectionNamespaceLister implements the MongoCollectionNamespaceLister
// interface.
type mongoCollectionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MongoCollections in the indexer for a given namespace.
func (s mongoCollectionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MongoCollection, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MongoCollection))
	})
	return ret, err
}

// Get retrieves the MongoCollection from the indexer for a given namespace and name.
func (s mongoCollectionNamespaceLister) Get(name string) (*v1alpha1.MongoCollection, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("mongocollection"), name)
	}
	return obj.(*v1alpha1.MongoCollection), nil
}