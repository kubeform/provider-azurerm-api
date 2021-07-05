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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/application/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InsightsSmartDetectionRuleLister helps list InsightsSmartDetectionRules.
// All objects returned here must be treated as read-only.
type InsightsSmartDetectionRuleLister interface {
	// List lists all InsightsSmartDetectionRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InsightsSmartDetectionRule, err error)
	// InsightsSmartDetectionRules returns an object that can list and get InsightsSmartDetectionRules.
	InsightsSmartDetectionRules(namespace string) InsightsSmartDetectionRuleNamespaceLister
	InsightsSmartDetectionRuleListerExpansion
}

// insightsSmartDetectionRuleLister implements the InsightsSmartDetectionRuleLister interface.
type insightsSmartDetectionRuleLister struct {
	indexer cache.Indexer
}

// NewInsightsSmartDetectionRuleLister returns a new InsightsSmartDetectionRuleLister.
func NewInsightsSmartDetectionRuleLister(indexer cache.Indexer) InsightsSmartDetectionRuleLister {
	return &insightsSmartDetectionRuleLister{indexer: indexer}
}

// List lists all InsightsSmartDetectionRules in the indexer.
func (s *insightsSmartDetectionRuleLister) List(selector labels.Selector) (ret []*v1alpha1.InsightsSmartDetectionRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InsightsSmartDetectionRule))
	})
	return ret, err
}

// InsightsSmartDetectionRules returns an object that can list and get InsightsSmartDetectionRules.
func (s *insightsSmartDetectionRuleLister) InsightsSmartDetectionRules(namespace string) InsightsSmartDetectionRuleNamespaceLister {
	return insightsSmartDetectionRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InsightsSmartDetectionRuleNamespaceLister helps list and get InsightsSmartDetectionRules.
// All objects returned here must be treated as read-only.
type InsightsSmartDetectionRuleNamespaceLister interface {
	// List lists all InsightsSmartDetectionRules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InsightsSmartDetectionRule, err error)
	// Get retrieves the InsightsSmartDetectionRule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.InsightsSmartDetectionRule, error)
	InsightsSmartDetectionRuleNamespaceListerExpansion
}

// insightsSmartDetectionRuleNamespaceLister implements the InsightsSmartDetectionRuleNamespaceLister
// interface.
type insightsSmartDetectionRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all InsightsSmartDetectionRules in the indexer for a given namespace.
func (s insightsSmartDetectionRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.InsightsSmartDetectionRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InsightsSmartDetectionRule))
	})
	return ret, err
}

// Get retrieves the InsightsSmartDetectionRule from the indexer for a given namespace and name.
func (s insightsSmartDetectionRuleNamespaceLister) Get(name string) (*v1alpha1.InsightsSmartDetectionRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("insightssmartdetectionrule"), name)
	}
	return obj.(*v1alpha1.InsightsSmartDetectionRule), nil
}