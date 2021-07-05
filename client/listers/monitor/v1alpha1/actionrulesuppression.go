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

// ActionRuleSuppressionLister helps list ActionRuleSuppressions.
// All objects returned here must be treated as read-only.
type ActionRuleSuppressionLister interface {
	// List lists all ActionRuleSuppressions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ActionRuleSuppression, err error)
	// ActionRuleSuppressions returns an object that can list and get ActionRuleSuppressions.
	ActionRuleSuppressions(namespace string) ActionRuleSuppressionNamespaceLister
	ActionRuleSuppressionListerExpansion
}

// actionRuleSuppressionLister implements the ActionRuleSuppressionLister interface.
type actionRuleSuppressionLister struct {
	indexer cache.Indexer
}

// NewActionRuleSuppressionLister returns a new ActionRuleSuppressionLister.
func NewActionRuleSuppressionLister(indexer cache.Indexer) ActionRuleSuppressionLister {
	return &actionRuleSuppressionLister{indexer: indexer}
}

// List lists all ActionRuleSuppressions in the indexer.
func (s *actionRuleSuppressionLister) List(selector labels.Selector) (ret []*v1alpha1.ActionRuleSuppression, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ActionRuleSuppression))
	})
	return ret, err
}

// ActionRuleSuppressions returns an object that can list and get ActionRuleSuppressions.
func (s *actionRuleSuppressionLister) ActionRuleSuppressions(namespace string) ActionRuleSuppressionNamespaceLister {
	return actionRuleSuppressionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ActionRuleSuppressionNamespaceLister helps list and get ActionRuleSuppressions.
// All objects returned here must be treated as read-only.
type ActionRuleSuppressionNamespaceLister interface {
	// List lists all ActionRuleSuppressions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ActionRuleSuppression, err error)
	// Get retrieves the ActionRuleSuppression from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ActionRuleSuppression, error)
	ActionRuleSuppressionNamespaceListerExpansion
}

// actionRuleSuppressionNamespaceLister implements the ActionRuleSuppressionNamespaceLister
// interface.
type actionRuleSuppressionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ActionRuleSuppressions in the indexer for a given namespace.
func (s actionRuleSuppressionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ActionRuleSuppression, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ActionRuleSuppression))
	})
	return ret, err
}

// Get retrieves the ActionRuleSuppression from the indexer for a given namespace and name.
func (s actionRuleSuppressionNamespaceLister) Get(name string) (*v1alpha1.ActionRuleSuppression, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("actionrulesuppression"), name)
	}
	return obj.(*v1alpha1.ActionRuleSuppression), nil
}