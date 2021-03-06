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

// ShareDatasetKustoDatabaseLister helps list ShareDatasetKustoDatabases.
// All objects returned here must be treated as read-only.
type ShareDatasetKustoDatabaseLister interface {
	// List lists all ShareDatasetKustoDatabases in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ShareDatasetKustoDatabase, err error)
	// ShareDatasetKustoDatabases returns an object that can list and get ShareDatasetKustoDatabases.
	ShareDatasetKustoDatabases(namespace string) ShareDatasetKustoDatabaseNamespaceLister
	ShareDatasetKustoDatabaseListerExpansion
}

// shareDatasetKustoDatabaseLister implements the ShareDatasetKustoDatabaseLister interface.
type shareDatasetKustoDatabaseLister struct {
	indexer cache.Indexer
}

// NewShareDatasetKustoDatabaseLister returns a new ShareDatasetKustoDatabaseLister.
func NewShareDatasetKustoDatabaseLister(indexer cache.Indexer) ShareDatasetKustoDatabaseLister {
	return &shareDatasetKustoDatabaseLister{indexer: indexer}
}

// List lists all ShareDatasetKustoDatabases in the indexer.
func (s *shareDatasetKustoDatabaseLister) List(selector labels.Selector) (ret []*v1alpha1.ShareDatasetKustoDatabase, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ShareDatasetKustoDatabase))
	})
	return ret, err
}

// ShareDatasetKustoDatabases returns an object that can list and get ShareDatasetKustoDatabases.
func (s *shareDatasetKustoDatabaseLister) ShareDatasetKustoDatabases(namespace string) ShareDatasetKustoDatabaseNamespaceLister {
	return shareDatasetKustoDatabaseNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ShareDatasetKustoDatabaseNamespaceLister helps list and get ShareDatasetKustoDatabases.
// All objects returned here must be treated as read-only.
type ShareDatasetKustoDatabaseNamespaceLister interface {
	// List lists all ShareDatasetKustoDatabases in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ShareDatasetKustoDatabase, err error)
	// Get retrieves the ShareDatasetKustoDatabase from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ShareDatasetKustoDatabase, error)
	ShareDatasetKustoDatabaseNamespaceListerExpansion
}

// shareDatasetKustoDatabaseNamespaceLister implements the ShareDatasetKustoDatabaseNamespaceLister
// interface.
type shareDatasetKustoDatabaseNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ShareDatasetKustoDatabases in the indexer for a given namespace.
func (s shareDatasetKustoDatabaseNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ShareDatasetKustoDatabase, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ShareDatasetKustoDatabase))
	})
	return ret, err
}

// Get retrieves the ShareDatasetKustoDatabase from the indexer for a given namespace and name.
func (s shareDatasetKustoDatabaseNamespaceLister) Get(name string) (*v1alpha1.ShareDatasetKustoDatabase, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sharedatasetkustodatabase"), name)
	}
	return obj.(*v1alpha1.ShareDatasetKustoDatabase), nil
}
