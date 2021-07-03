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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/eventhub/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NamespaceCustomerManagedKeyLister helps list NamespaceCustomerManagedKeys.
// All objects returned here must be treated as read-only.
type NamespaceCustomerManagedKeyLister interface {
	// List lists all NamespaceCustomerManagedKeys in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NamespaceCustomerManagedKey, err error)
	// NamespaceCustomerManagedKeys returns an object that can list and get NamespaceCustomerManagedKeys.
	NamespaceCustomerManagedKeys(namespace string) NamespaceCustomerManagedKeyNamespaceLister
	NamespaceCustomerManagedKeyListerExpansion
}

// namespaceCustomerManagedKeyLister implements the NamespaceCustomerManagedKeyLister interface.
type namespaceCustomerManagedKeyLister struct {
	indexer cache.Indexer
}

// NewNamespaceCustomerManagedKeyLister returns a new NamespaceCustomerManagedKeyLister.
func NewNamespaceCustomerManagedKeyLister(indexer cache.Indexer) NamespaceCustomerManagedKeyLister {
	return &namespaceCustomerManagedKeyLister{indexer: indexer}
}

// List lists all NamespaceCustomerManagedKeys in the indexer.
func (s *namespaceCustomerManagedKeyLister) List(selector labels.Selector) (ret []*v1alpha1.NamespaceCustomerManagedKey, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NamespaceCustomerManagedKey))
	})
	return ret, err
}

// NamespaceCustomerManagedKeys returns an object that can list and get NamespaceCustomerManagedKeys.
func (s *namespaceCustomerManagedKeyLister) NamespaceCustomerManagedKeys(namespace string) NamespaceCustomerManagedKeyNamespaceLister {
	return namespaceCustomerManagedKeyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NamespaceCustomerManagedKeyNamespaceLister helps list and get NamespaceCustomerManagedKeys.
// All objects returned here must be treated as read-only.
type NamespaceCustomerManagedKeyNamespaceLister interface {
	// List lists all NamespaceCustomerManagedKeys in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NamespaceCustomerManagedKey, err error)
	// Get retrieves the NamespaceCustomerManagedKey from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.NamespaceCustomerManagedKey, error)
	NamespaceCustomerManagedKeyNamespaceListerExpansion
}

// namespaceCustomerManagedKeyNamespaceLister implements the NamespaceCustomerManagedKeyNamespaceLister
// interface.
type namespaceCustomerManagedKeyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NamespaceCustomerManagedKeys in the indexer for a given namespace.
func (s namespaceCustomerManagedKeyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NamespaceCustomerManagedKey, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NamespaceCustomerManagedKey))
	})
	return ret, err
}

// Get retrieves the NamespaceCustomerManagedKey from the indexer for a given namespace and name.
func (s namespaceCustomerManagedKeyNamespaceLister) Get(name string) (*v1alpha1.NamespaceCustomerManagedKey, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("namespacecustomermanagedkey"), name)
	}
	return obj.(*v1alpha1.NamespaceCustomerManagedKey), nil
}
