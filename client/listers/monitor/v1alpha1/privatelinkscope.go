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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/monitor/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PrivateLinkScopeLister helps list PrivateLinkScopes.
// All objects returned here must be treated as read-only.
type PrivateLinkScopeLister interface {
	// List lists all PrivateLinkScopes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PrivateLinkScope, err error)
	// PrivateLinkScopes returns an object that can list and get PrivateLinkScopes.
	PrivateLinkScopes(namespace string) PrivateLinkScopeNamespaceLister
	PrivateLinkScopeListerExpansion
}

// privateLinkScopeLister implements the PrivateLinkScopeLister interface.
type privateLinkScopeLister struct {
	indexer cache.Indexer
}

// NewPrivateLinkScopeLister returns a new PrivateLinkScopeLister.
func NewPrivateLinkScopeLister(indexer cache.Indexer) PrivateLinkScopeLister {
	return &privateLinkScopeLister{indexer: indexer}
}

// List lists all PrivateLinkScopes in the indexer.
func (s *privateLinkScopeLister) List(selector labels.Selector) (ret []*v1alpha1.PrivateLinkScope, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PrivateLinkScope))
	})
	return ret, err
}

// PrivateLinkScopes returns an object that can list and get PrivateLinkScopes.
func (s *privateLinkScopeLister) PrivateLinkScopes(namespace string) PrivateLinkScopeNamespaceLister {
	return privateLinkScopeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PrivateLinkScopeNamespaceLister helps list and get PrivateLinkScopes.
// All objects returned here must be treated as read-only.
type PrivateLinkScopeNamespaceLister interface {
	// List lists all PrivateLinkScopes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PrivateLinkScope, err error)
	// Get retrieves the PrivateLinkScope from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PrivateLinkScope, error)
	PrivateLinkScopeNamespaceListerExpansion
}

// privateLinkScopeNamespaceLister implements the PrivateLinkScopeNamespaceLister
// interface.
type privateLinkScopeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PrivateLinkScopes in the indexer for a given namespace.
func (s privateLinkScopeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PrivateLinkScope, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PrivateLinkScope))
	})
	return ret, err
}

// Get retrieves the PrivateLinkScope from the indexer for a given namespace and name.
func (s privateLinkScopeNamespaceLister) Get(name string) (*v1alpha1.PrivateLinkScope, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("privatelinkscope"), name)
	}
	return obj.(*v1alpha1.PrivateLinkScope), nil
}