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

// NamespaceAuthorizationRuleLister helps list NamespaceAuthorizationRules.
// All objects returned here must be treated as read-only.
type NamespaceAuthorizationRuleLister interface {
	// List lists all NamespaceAuthorizationRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NamespaceAuthorizationRule, err error)
	// NamespaceAuthorizationRules returns an object that can list and get NamespaceAuthorizationRules.
	NamespaceAuthorizationRules(namespace string) NamespaceAuthorizationRuleNamespaceLister
	NamespaceAuthorizationRuleListerExpansion
}

// namespaceAuthorizationRuleLister implements the NamespaceAuthorizationRuleLister interface.
type namespaceAuthorizationRuleLister struct {
	indexer cache.Indexer
}

// NewNamespaceAuthorizationRuleLister returns a new NamespaceAuthorizationRuleLister.
func NewNamespaceAuthorizationRuleLister(indexer cache.Indexer) NamespaceAuthorizationRuleLister {
	return &namespaceAuthorizationRuleLister{indexer: indexer}
}

// List lists all NamespaceAuthorizationRules in the indexer.
func (s *namespaceAuthorizationRuleLister) List(selector labels.Selector) (ret []*v1alpha1.NamespaceAuthorizationRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NamespaceAuthorizationRule))
	})
	return ret, err
}

// NamespaceAuthorizationRules returns an object that can list and get NamespaceAuthorizationRules.
func (s *namespaceAuthorizationRuleLister) NamespaceAuthorizationRules(namespace string) NamespaceAuthorizationRuleNamespaceLister {
	return namespaceAuthorizationRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NamespaceAuthorizationRuleNamespaceLister helps list and get NamespaceAuthorizationRules.
// All objects returned here must be treated as read-only.
type NamespaceAuthorizationRuleNamespaceLister interface {
	// List lists all NamespaceAuthorizationRules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NamespaceAuthorizationRule, err error)
	// Get retrieves the NamespaceAuthorizationRule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.NamespaceAuthorizationRule, error)
	NamespaceAuthorizationRuleNamespaceListerExpansion
}

// namespaceAuthorizationRuleNamespaceLister implements the NamespaceAuthorizationRuleNamespaceLister
// interface.
type namespaceAuthorizationRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NamespaceAuthorizationRules in the indexer for a given namespace.
func (s namespaceAuthorizationRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NamespaceAuthorizationRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NamespaceAuthorizationRule))
	})
	return ret, err
}

// Get retrieves the NamespaceAuthorizationRule from the indexer for a given namespace and name.
func (s namespaceAuthorizationRuleNamespaceLister) Get(name string) (*v1alpha1.NamespaceAuthorizationRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("namespaceauthorizationrule"), name)
	}
	return obj.(*v1alpha1.NamespaceAuthorizationRule), nil
}
