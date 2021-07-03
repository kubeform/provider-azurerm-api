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

// FlexibleServerFirewallRuleLister helps list FlexibleServerFirewallRules.
// All objects returned here must be treated as read-only.
type FlexibleServerFirewallRuleLister interface {
	// List lists all FlexibleServerFirewallRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FlexibleServerFirewallRule, err error)
	// FlexibleServerFirewallRules returns an object that can list and get FlexibleServerFirewallRules.
	FlexibleServerFirewallRules(namespace string) FlexibleServerFirewallRuleNamespaceLister
	FlexibleServerFirewallRuleListerExpansion
}

// flexibleServerFirewallRuleLister implements the FlexibleServerFirewallRuleLister interface.
type flexibleServerFirewallRuleLister struct {
	indexer cache.Indexer
}

// NewFlexibleServerFirewallRuleLister returns a new FlexibleServerFirewallRuleLister.
func NewFlexibleServerFirewallRuleLister(indexer cache.Indexer) FlexibleServerFirewallRuleLister {
	return &flexibleServerFirewallRuleLister{indexer: indexer}
}

// List lists all FlexibleServerFirewallRules in the indexer.
func (s *flexibleServerFirewallRuleLister) List(selector labels.Selector) (ret []*v1alpha1.FlexibleServerFirewallRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FlexibleServerFirewallRule))
	})
	return ret, err
}

// FlexibleServerFirewallRules returns an object that can list and get FlexibleServerFirewallRules.
func (s *flexibleServerFirewallRuleLister) FlexibleServerFirewallRules(namespace string) FlexibleServerFirewallRuleNamespaceLister {
	return flexibleServerFirewallRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FlexibleServerFirewallRuleNamespaceLister helps list and get FlexibleServerFirewallRules.
// All objects returned here must be treated as read-only.
type FlexibleServerFirewallRuleNamespaceLister interface {
	// List lists all FlexibleServerFirewallRules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FlexibleServerFirewallRule, err error)
	// Get retrieves the FlexibleServerFirewallRule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FlexibleServerFirewallRule, error)
	FlexibleServerFirewallRuleNamespaceListerExpansion
}

// flexibleServerFirewallRuleNamespaceLister implements the FlexibleServerFirewallRuleNamespaceLister
// interface.
type flexibleServerFirewallRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FlexibleServerFirewallRules in the indexer for a given namespace.
func (s flexibleServerFirewallRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FlexibleServerFirewallRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FlexibleServerFirewallRule))
	})
	return ret, err
}

// Get retrieves the FlexibleServerFirewallRule from the indexer for a given namespace and name.
func (s flexibleServerFirewallRuleNamespaceLister) Get(name string) (*v1alpha1.FlexibleServerFirewallRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("flexibleserverfirewallrule"), name)
	}
	return obj.(*v1alpha1.FlexibleServerFirewallRule), nil
}
