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

// VirtualNetworkRuleLister helps list VirtualNetworkRules.
// All objects returned here must be treated as read-only.
type VirtualNetworkRuleLister interface {
	// List lists all VirtualNetworkRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VirtualNetworkRule, err error)
	// VirtualNetworkRules returns an object that can list and get VirtualNetworkRules.
	VirtualNetworkRules(namespace string) VirtualNetworkRuleNamespaceLister
	VirtualNetworkRuleListerExpansion
}

// virtualNetworkRuleLister implements the VirtualNetworkRuleLister interface.
type virtualNetworkRuleLister struct {
	indexer cache.Indexer
}

// NewVirtualNetworkRuleLister returns a new VirtualNetworkRuleLister.
func NewVirtualNetworkRuleLister(indexer cache.Indexer) VirtualNetworkRuleLister {
	return &virtualNetworkRuleLister{indexer: indexer}
}

// List lists all VirtualNetworkRules in the indexer.
func (s *virtualNetworkRuleLister) List(selector labels.Selector) (ret []*v1alpha1.VirtualNetworkRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VirtualNetworkRule))
	})
	return ret, err
}

// VirtualNetworkRules returns an object that can list and get VirtualNetworkRules.
func (s *virtualNetworkRuleLister) VirtualNetworkRules(namespace string) VirtualNetworkRuleNamespaceLister {
	return virtualNetworkRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VirtualNetworkRuleNamespaceLister helps list and get VirtualNetworkRules.
// All objects returned here must be treated as read-only.
type VirtualNetworkRuleNamespaceLister interface {
	// List lists all VirtualNetworkRules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VirtualNetworkRule, err error)
	// Get retrieves the VirtualNetworkRule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.VirtualNetworkRule, error)
	VirtualNetworkRuleNamespaceListerExpansion
}

// virtualNetworkRuleNamespaceLister implements the VirtualNetworkRuleNamespaceLister
// interface.
type virtualNetworkRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VirtualNetworkRules in the indexer for a given namespace.
func (s virtualNetworkRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VirtualNetworkRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VirtualNetworkRule))
	})
	return ret, err
}

// Get retrieves the VirtualNetworkRule from the indexer for a given namespace and name.
func (s virtualNetworkRuleNamespaceLister) Get(name string) (*v1alpha1.VirtualNetworkRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("virtualnetworkrule"), name)
	}
	return obj.(*v1alpha1.VirtualNetworkRule), nil
}
