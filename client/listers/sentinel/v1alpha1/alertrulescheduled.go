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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/sentinel/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AlertRuleScheduledLister helps list AlertRuleScheduleds.
// All objects returned here must be treated as read-only.
type AlertRuleScheduledLister interface {
	// List lists all AlertRuleScheduleds in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AlertRuleScheduled, err error)
	// AlertRuleScheduleds returns an object that can list and get AlertRuleScheduleds.
	AlertRuleScheduleds(namespace string) AlertRuleScheduledNamespaceLister
	AlertRuleScheduledListerExpansion
}

// alertRuleScheduledLister implements the AlertRuleScheduledLister interface.
type alertRuleScheduledLister struct {
	indexer cache.Indexer
}

// NewAlertRuleScheduledLister returns a new AlertRuleScheduledLister.
func NewAlertRuleScheduledLister(indexer cache.Indexer) AlertRuleScheduledLister {
	return &alertRuleScheduledLister{indexer: indexer}
}

// List lists all AlertRuleScheduleds in the indexer.
func (s *alertRuleScheduledLister) List(selector labels.Selector) (ret []*v1alpha1.AlertRuleScheduled, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AlertRuleScheduled))
	})
	return ret, err
}

// AlertRuleScheduleds returns an object that can list and get AlertRuleScheduleds.
func (s *alertRuleScheduledLister) AlertRuleScheduleds(namespace string) AlertRuleScheduledNamespaceLister {
	return alertRuleScheduledNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AlertRuleScheduledNamespaceLister helps list and get AlertRuleScheduleds.
// All objects returned here must be treated as read-only.
type AlertRuleScheduledNamespaceLister interface {
	// List lists all AlertRuleScheduleds in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AlertRuleScheduled, err error)
	// Get retrieves the AlertRuleScheduled from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AlertRuleScheduled, error)
	AlertRuleScheduledNamespaceListerExpansion
}

// alertRuleScheduledNamespaceLister implements the AlertRuleScheduledNamespaceLister
// interface.
type alertRuleScheduledNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AlertRuleScheduleds in the indexer for a given namespace.
func (s alertRuleScheduledNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AlertRuleScheduled, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AlertRuleScheduled))
	})
	return ret, err
}

// Get retrieves the AlertRuleScheduled from the indexer for a given namespace and name.
func (s alertRuleScheduledNamespaceLister) Get(name string) (*v1alpha1.AlertRuleScheduled, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("alertrulescheduled"), name)
	}
	return obj.(*v1alpha1.AlertRuleScheduled), nil
}