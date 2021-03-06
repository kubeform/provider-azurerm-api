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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/keyvault/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ManagedStorageAccountSasTokenDefinitionLister helps list ManagedStorageAccountSasTokenDefinitions.
// All objects returned here must be treated as read-only.
type ManagedStorageAccountSasTokenDefinitionLister interface {
	// List lists all ManagedStorageAccountSasTokenDefinitions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagedStorageAccountSasTokenDefinition, err error)
	// ManagedStorageAccountSasTokenDefinitions returns an object that can list and get ManagedStorageAccountSasTokenDefinitions.
	ManagedStorageAccountSasTokenDefinitions(namespace string) ManagedStorageAccountSasTokenDefinitionNamespaceLister
	ManagedStorageAccountSasTokenDefinitionListerExpansion
}

// managedStorageAccountSasTokenDefinitionLister implements the ManagedStorageAccountSasTokenDefinitionLister interface.
type managedStorageAccountSasTokenDefinitionLister struct {
	indexer cache.Indexer
}

// NewManagedStorageAccountSasTokenDefinitionLister returns a new ManagedStorageAccountSasTokenDefinitionLister.
func NewManagedStorageAccountSasTokenDefinitionLister(indexer cache.Indexer) ManagedStorageAccountSasTokenDefinitionLister {
	return &managedStorageAccountSasTokenDefinitionLister{indexer: indexer}
}

// List lists all ManagedStorageAccountSasTokenDefinitions in the indexer.
func (s *managedStorageAccountSasTokenDefinitionLister) List(selector labels.Selector) (ret []*v1alpha1.ManagedStorageAccountSasTokenDefinition, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagedStorageAccountSasTokenDefinition))
	})
	return ret, err
}

// ManagedStorageAccountSasTokenDefinitions returns an object that can list and get ManagedStorageAccountSasTokenDefinitions.
func (s *managedStorageAccountSasTokenDefinitionLister) ManagedStorageAccountSasTokenDefinitions(namespace string) ManagedStorageAccountSasTokenDefinitionNamespaceLister {
	return managedStorageAccountSasTokenDefinitionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ManagedStorageAccountSasTokenDefinitionNamespaceLister helps list and get ManagedStorageAccountSasTokenDefinitions.
// All objects returned here must be treated as read-only.
type ManagedStorageAccountSasTokenDefinitionNamespaceLister interface {
	// List lists all ManagedStorageAccountSasTokenDefinitions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagedStorageAccountSasTokenDefinition, err error)
	// Get retrieves the ManagedStorageAccountSasTokenDefinition from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ManagedStorageAccountSasTokenDefinition, error)
	ManagedStorageAccountSasTokenDefinitionNamespaceListerExpansion
}

// managedStorageAccountSasTokenDefinitionNamespaceLister implements the ManagedStorageAccountSasTokenDefinitionNamespaceLister
// interface.
type managedStorageAccountSasTokenDefinitionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ManagedStorageAccountSasTokenDefinitions in the indexer for a given namespace.
func (s managedStorageAccountSasTokenDefinitionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ManagedStorageAccountSasTokenDefinition, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagedStorageAccountSasTokenDefinition))
	})
	return ret, err
}

// Get retrieves the ManagedStorageAccountSasTokenDefinition from the indexer for a given namespace and name.
func (s managedStorageAccountSasTokenDefinitionNamespaceLister) Get(name string) (*v1alpha1.ManagedStorageAccountSasTokenDefinition, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("managedstorageaccountsastokendefinition"), name)
	}
	return obj.(*v1alpha1.ManagedStorageAccountSasTokenDefinition), nil
}
