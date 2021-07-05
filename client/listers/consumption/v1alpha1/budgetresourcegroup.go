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

// BudgetResourceGroupLister helps list BudgetResourceGroups.
// All objects returned here must be treated as read-only.
type BudgetResourceGroupLister interface {
	// List lists all BudgetResourceGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BudgetResourceGroup, err error)
	// BudgetResourceGroups returns an object that can list and get BudgetResourceGroups.
	BudgetResourceGroups(namespace string) BudgetResourceGroupNamespaceLister
	BudgetResourceGroupListerExpansion
}

// budgetResourceGroupLister implements the BudgetResourceGroupLister interface.
type budgetResourceGroupLister struct {
	indexer cache.Indexer
}

// NewBudgetResourceGroupLister returns a new BudgetResourceGroupLister.
func NewBudgetResourceGroupLister(indexer cache.Indexer) BudgetResourceGroupLister {
	return &budgetResourceGroupLister{indexer: indexer}
}

// List lists all BudgetResourceGroups in the indexer.
func (s *budgetResourceGroupLister) List(selector labels.Selector) (ret []*v1alpha1.BudgetResourceGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BudgetResourceGroup))
	})
	return ret, err
}

// BudgetResourceGroups returns an object that can list and get BudgetResourceGroups.
func (s *budgetResourceGroupLister) BudgetResourceGroups(namespace string) BudgetResourceGroupNamespaceLister {
	return budgetResourceGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BudgetResourceGroupNamespaceLister helps list and get BudgetResourceGroups.
// All objects returned here must be treated as read-only.
type BudgetResourceGroupNamespaceLister interface {
	// List lists all BudgetResourceGroups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BudgetResourceGroup, err error)
	// Get retrieves the BudgetResourceGroup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BudgetResourceGroup, error)
	BudgetResourceGroupNamespaceListerExpansion
}

// budgetResourceGroupNamespaceLister implements the BudgetResourceGroupNamespaceLister
// interface.
type budgetResourceGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BudgetResourceGroups in the indexer for a given namespace.
func (s budgetResourceGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BudgetResourceGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BudgetResourceGroup))
	})
	return ret, err
}

// Get retrieves the BudgetResourceGroup from the indexer for a given namespace and name.
func (s budgetResourceGroupNamespaceLister) Get(name string) (*v1alpha1.BudgetResourceGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("budgetresourcegroup"), name)
	}
	return obj.(*v1alpha1.BudgetResourceGroup), nil
}