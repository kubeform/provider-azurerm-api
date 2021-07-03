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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/netapp/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AccountLister helps list Accounts.
// All objects returned here must be treated as read-only.
type AccountLister interface {
	// List lists all Accounts in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Account, err error)
	// Accounts returns an object that can list and get Accounts.
	Accounts(namespace string) AccountNamespaceLister
	AccountListerExpansion
}

// accountLister implements the AccountLister interface.
type accountLister struct {
	indexer cache.Indexer
}

// NewAccountLister returns a new AccountLister.
func NewAccountLister(indexer cache.Indexer) AccountLister {
	return &accountLister{indexer: indexer}
}

// List lists all Accounts in the indexer.
func (s *accountLister) List(selector labels.Selector) (ret []*v1alpha1.Account, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Account))
	})
	return ret, err
}

// Accounts returns an object that can list and get Accounts.
func (s *accountLister) Accounts(namespace string) AccountNamespaceLister {
	return accountNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AccountNamespaceLister helps list and get Accounts.
// All objects returned here must be treated as read-only.
type AccountNamespaceLister interface {
	// List lists all Accounts in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Account, err error)
	// Get retrieves the Account from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Account, error)
	AccountNamespaceListerExpansion
}

// accountNamespaceLister implements the AccountNamespaceLister
// interface.
type accountNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Accounts in the indexer for a given namespace.
func (s accountNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Account, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Account))
	})
	return ret, err
}

// Get retrieves the Account from the indexer for a given namespace and name.
func (s accountNamespaceLister) Get(name string) (*v1alpha1.Account, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("account"), name)
	}
	return obj.(*v1alpha1.Account), nil
}
