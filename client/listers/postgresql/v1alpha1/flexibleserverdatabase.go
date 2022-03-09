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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/postgresql/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FlexibleServerDatabaseLister helps list FlexibleServerDatabases.
// All objects returned here must be treated as read-only.
type FlexibleServerDatabaseLister interface {
	// List lists all FlexibleServerDatabases in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FlexibleServerDatabase, err error)
	// FlexibleServerDatabases returns an object that can list and get FlexibleServerDatabases.
	FlexibleServerDatabases(namespace string) FlexibleServerDatabaseNamespaceLister
	FlexibleServerDatabaseListerExpansion
}

// flexibleServerDatabaseLister implements the FlexibleServerDatabaseLister interface.
type flexibleServerDatabaseLister struct {
	indexer cache.Indexer
}

// NewFlexibleServerDatabaseLister returns a new FlexibleServerDatabaseLister.
func NewFlexibleServerDatabaseLister(indexer cache.Indexer) FlexibleServerDatabaseLister {
	return &flexibleServerDatabaseLister{indexer: indexer}
}

// List lists all FlexibleServerDatabases in the indexer.
func (s *flexibleServerDatabaseLister) List(selector labels.Selector) (ret []*v1alpha1.FlexibleServerDatabase, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FlexibleServerDatabase))
	})
	return ret, err
}

// FlexibleServerDatabases returns an object that can list and get FlexibleServerDatabases.
func (s *flexibleServerDatabaseLister) FlexibleServerDatabases(namespace string) FlexibleServerDatabaseNamespaceLister {
	return flexibleServerDatabaseNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FlexibleServerDatabaseNamespaceLister helps list and get FlexibleServerDatabases.
// All objects returned here must be treated as read-only.
type FlexibleServerDatabaseNamespaceLister interface {
	// List lists all FlexibleServerDatabases in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FlexibleServerDatabase, err error)
	// Get retrieves the FlexibleServerDatabase from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FlexibleServerDatabase, error)
	FlexibleServerDatabaseNamespaceListerExpansion
}

// flexibleServerDatabaseNamespaceLister implements the FlexibleServerDatabaseNamespaceLister
// interface.
type flexibleServerDatabaseNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FlexibleServerDatabases in the indexer for a given namespace.
func (s flexibleServerDatabaseNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FlexibleServerDatabase, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FlexibleServerDatabase))
	})
	return ret, err
}

// Get retrieves the FlexibleServerDatabase from the indexer for a given namespace and name.
func (s flexibleServerDatabaseNamespaceLister) Get(name string) (*v1alpha1.FlexibleServerDatabase, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("flexibleserverdatabase"), name)
	}
	return obj.(*v1alpha1.FlexibleServerDatabase), nil
}
