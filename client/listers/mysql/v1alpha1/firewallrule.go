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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/mysql/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FirewallRuleLister helps list FirewallRules.
// All objects returned here must be treated as read-only.
type FirewallRuleLister interface {
	// List lists all FirewallRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FirewallRule, err error)
	// FirewallRules returns an object that can list and get FirewallRules.
	FirewallRules(namespace string) FirewallRuleNamespaceLister
	FirewallRuleListerExpansion
}

// firewallRuleLister implements the FirewallRuleLister interface.
type firewallRuleLister struct {
	indexer cache.Indexer
}

// NewFirewallRuleLister returns a new FirewallRuleLister.
func NewFirewallRuleLister(indexer cache.Indexer) FirewallRuleLister {
	return &firewallRuleLister{indexer: indexer}
}

// List lists all FirewallRules in the indexer.
func (s *firewallRuleLister) List(selector labels.Selector) (ret []*v1alpha1.FirewallRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FirewallRule))
	})
	return ret, err
}

// FirewallRules returns an object that can list and get FirewallRules.
func (s *firewallRuleLister) FirewallRules(namespace string) FirewallRuleNamespaceLister {
	return firewallRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FirewallRuleNamespaceLister helps list and get FirewallRules.
// All objects returned here must be treated as read-only.
type FirewallRuleNamespaceLister interface {
	// List lists all FirewallRules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FirewallRule, err error)
	// Get retrieves the FirewallRule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FirewallRule, error)
	FirewallRuleNamespaceListerExpansion
}

// firewallRuleNamespaceLister implements the FirewallRuleNamespaceLister
// interface.
type firewallRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FirewallRules in the indexer for a given namespace.
func (s firewallRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FirewallRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FirewallRule))
	})
	return ret, err
}

// Get retrieves the FirewallRule from the indexer for a given namespace and name.
func (s firewallRuleNamespaceLister) Get(name string) (*v1alpha1.FirewallRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("firewallrule"), name)
	}
	return obj.(*v1alpha1.FirewallRule), nil
}
