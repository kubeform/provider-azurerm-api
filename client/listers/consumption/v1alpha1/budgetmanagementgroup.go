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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/consumption/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BudgetManagementGroupLister helps list BudgetManagementGroups.
// All objects returned here must be treated as read-only.
type BudgetManagementGroupLister interface {
	// List lists all BudgetManagementGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BudgetManagementGroup, err error)
	// BudgetManagementGroups returns an object that can list and get BudgetManagementGroups.
	BudgetManagementGroups(namespace string) BudgetManagementGroupNamespaceLister
	BudgetManagementGroupListerExpansion
}

// budgetManagementGroupLister implements the BudgetManagementGroupLister interface.
type budgetManagementGroupLister struct {
	indexer cache.Indexer
}

// NewBudgetManagementGroupLister returns a new BudgetManagementGroupLister.
func NewBudgetManagementGroupLister(indexer cache.Indexer) BudgetManagementGroupLister {
	return &budgetManagementGroupLister{indexer: indexer}
}

// List lists all BudgetManagementGroups in the indexer.
func (s *budgetManagementGroupLister) List(selector labels.Selector) (ret []*v1alpha1.BudgetManagementGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BudgetManagementGroup))
	})
	return ret, err
}

// BudgetManagementGroups returns an object that can list and get BudgetManagementGroups.
func (s *budgetManagementGroupLister) BudgetManagementGroups(namespace string) BudgetManagementGroupNamespaceLister {
	return budgetManagementGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BudgetManagementGroupNamespaceLister helps list and get BudgetManagementGroups.
// All objects returned here must be treated as read-only.
type BudgetManagementGroupNamespaceLister interface {
	// List lists all BudgetManagementGroups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BudgetManagementGroup, err error)
	// Get retrieves the BudgetManagementGroup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BudgetManagementGroup, error)
	BudgetManagementGroupNamespaceListerExpansion
}

// budgetManagementGroupNamespaceLister implements the BudgetManagementGroupNamespaceLister
// interface.
type budgetManagementGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BudgetManagementGroups in the indexer for a given namespace.
func (s budgetManagementGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BudgetManagementGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BudgetManagementGroup))
	})
	return ret, err
}

// Get retrieves the BudgetManagementGroup from the indexer for a given namespace and name.
func (s budgetManagementGroupNamespaceLister) Get(name string) (*v1alpha1.BudgetManagementGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("budgetmanagementgroup"), name)
	}
	return obj.(*v1alpha1.BudgetManagementGroup), nil
}