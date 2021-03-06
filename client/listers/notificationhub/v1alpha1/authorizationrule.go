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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/notificationhub/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AuthorizationRuleLister helps list AuthorizationRules.
// All objects returned here must be treated as read-only.
type AuthorizationRuleLister interface {
	// List lists all AuthorizationRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AuthorizationRule, err error)
	// AuthorizationRules returns an object that can list and get AuthorizationRules.
	AuthorizationRules(namespace string) AuthorizationRuleNamespaceLister
	AuthorizationRuleListerExpansion
}

// authorizationRuleLister implements the AuthorizationRuleLister interface.
type authorizationRuleLister struct {
	indexer cache.Indexer
}

// NewAuthorizationRuleLister returns a new AuthorizationRuleLister.
func NewAuthorizationRuleLister(indexer cache.Indexer) AuthorizationRuleLister {
	return &authorizationRuleLister{indexer: indexer}
}

// List lists all AuthorizationRules in the indexer.
func (s *authorizationRuleLister) List(selector labels.Selector) (ret []*v1alpha1.AuthorizationRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AuthorizationRule))
	})
	return ret, err
}

// AuthorizationRules returns an object that can list and get AuthorizationRules.
func (s *authorizationRuleLister) AuthorizationRules(namespace string) AuthorizationRuleNamespaceLister {
	return authorizationRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AuthorizationRuleNamespaceLister helps list and get AuthorizationRules.
// All objects returned here must be treated as read-only.
type AuthorizationRuleNamespaceLister interface {
	// List lists all AuthorizationRules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AuthorizationRule, err error)
	// Get retrieves the AuthorizationRule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AuthorizationRule, error)
	AuthorizationRuleNamespaceListerExpansion
}

// authorizationRuleNamespaceLister implements the AuthorizationRuleNamespaceLister
// interface.
type authorizationRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AuthorizationRules in the indexer for a given namespace.
func (s authorizationRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AuthorizationRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AuthorizationRule))
	})
	return ret, err
}

// Get retrieves the AuthorizationRule from the indexer for a given namespace and name.
func (s authorizationRuleNamespaceLister) Get(name string) (*v1alpha1.AuthorizationRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("authorizationrule"), name)
	}
	return obj.(*v1alpha1.AuthorizationRule), nil
}
