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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/databricks/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WorkspaceCustomerManagedKeyLister helps list WorkspaceCustomerManagedKeys.
// All objects returned here must be treated as read-only.
type WorkspaceCustomerManagedKeyLister interface {
	// List lists all WorkspaceCustomerManagedKeys in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WorkspaceCustomerManagedKey, err error)
	// WorkspaceCustomerManagedKeys returns an object that can list and get WorkspaceCustomerManagedKeys.
	WorkspaceCustomerManagedKeys(namespace string) WorkspaceCustomerManagedKeyNamespaceLister
	WorkspaceCustomerManagedKeyListerExpansion
}

// workspaceCustomerManagedKeyLister implements the WorkspaceCustomerManagedKeyLister interface.
type workspaceCustomerManagedKeyLister struct {
	indexer cache.Indexer
}

// NewWorkspaceCustomerManagedKeyLister returns a new WorkspaceCustomerManagedKeyLister.
func NewWorkspaceCustomerManagedKeyLister(indexer cache.Indexer) WorkspaceCustomerManagedKeyLister {
	return &workspaceCustomerManagedKeyLister{indexer: indexer}
}

// List lists all WorkspaceCustomerManagedKeys in the indexer.
func (s *workspaceCustomerManagedKeyLister) List(selector labels.Selector) (ret []*v1alpha1.WorkspaceCustomerManagedKey, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WorkspaceCustomerManagedKey))
	})
	return ret, err
}

// WorkspaceCustomerManagedKeys returns an object that can list and get WorkspaceCustomerManagedKeys.
func (s *workspaceCustomerManagedKeyLister) WorkspaceCustomerManagedKeys(namespace string) WorkspaceCustomerManagedKeyNamespaceLister {
	return workspaceCustomerManagedKeyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WorkspaceCustomerManagedKeyNamespaceLister helps list and get WorkspaceCustomerManagedKeys.
// All objects returned here must be treated as read-only.
type WorkspaceCustomerManagedKeyNamespaceLister interface {
	// List lists all WorkspaceCustomerManagedKeys in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WorkspaceCustomerManagedKey, err error)
	// Get retrieves the WorkspaceCustomerManagedKey from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.WorkspaceCustomerManagedKey, error)
	WorkspaceCustomerManagedKeyNamespaceListerExpansion
}

// workspaceCustomerManagedKeyNamespaceLister implements the WorkspaceCustomerManagedKeyNamespaceLister
// interface.
type workspaceCustomerManagedKeyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WorkspaceCustomerManagedKeys in the indexer for a given namespace.
func (s workspaceCustomerManagedKeyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WorkspaceCustomerManagedKey, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WorkspaceCustomerManagedKey))
	})
	return ret, err
}

// Get retrieves the WorkspaceCustomerManagedKey from the indexer for a given namespace and name.
func (s workspaceCustomerManagedKeyNamespaceLister) Get(name string) (*v1alpha1.WorkspaceCustomerManagedKey, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("workspacecustomermanagedkey"), name)
	}
	return obj.(*v1alpha1.WorkspaceCustomerManagedKey), nil
}
