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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/loganalytics/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LinkedStorageAccountLister helps list LinkedStorageAccounts.
// All objects returned here must be treated as read-only.
type LinkedStorageAccountLister interface {
	// List lists all LinkedStorageAccounts in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LinkedStorageAccount, err error)
	// LinkedStorageAccounts returns an object that can list and get LinkedStorageAccounts.
	LinkedStorageAccounts(namespace string) LinkedStorageAccountNamespaceLister
	LinkedStorageAccountListerExpansion
}

// linkedStorageAccountLister implements the LinkedStorageAccountLister interface.
type linkedStorageAccountLister struct {
	indexer cache.Indexer
}

// NewLinkedStorageAccountLister returns a new LinkedStorageAccountLister.
func NewLinkedStorageAccountLister(indexer cache.Indexer) LinkedStorageAccountLister {
	return &linkedStorageAccountLister{indexer: indexer}
}

// List lists all LinkedStorageAccounts in the indexer.
func (s *linkedStorageAccountLister) List(selector labels.Selector) (ret []*v1alpha1.LinkedStorageAccount, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LinkedStorageAccount))
	})
	return ret, err
}

// LinkedStorageAccounts returns an object that can list and get LinkedStorageAccounts.
func (s *linkedStorageAccountLister) LinkedStorageAccounts(namespace string) LinkedStorageAccountNamespaceLister {
	return linkedStorageAccountNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LinkedStorageAccountNamespaceLister helps list and get LinkedStorageAccounts.
// All objects returned here must be treated as read-only.
type LinkedStorageAccountNamespaceLister interface {
	// List lists all LinkedStorageAccounts in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LinkedStorageAccount, err error)
	// Get retrieves the LinkedStorageAccount from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.LinkedStorageAccount, error)
	LinkedStorageAccountNamespaceListerExpansion
}

// linkedStorageAccountNamespaceLister implements the LinkedStorageAccountNamespaceLister
// interface.
type linkedStorageAccountNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LinkedStorageAccounts in the indexer for a given namespace.
func (s linkedStorageAccountNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LinkedStorageAccount, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LinkedStorageAccount))
	})
	return ret, err
}

// Get retrieves the LinkedStorageAccount from the indexer for a given namespace and name.
func (s linkedStorageAccountNamespaceLister) Get(name string) (*v1alpha1.LinkedStorageAccount, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("linkedstorageaccount"), name)
	}
	return obj.(*v1alpha1.LinkedStorageAccount), nil
}
