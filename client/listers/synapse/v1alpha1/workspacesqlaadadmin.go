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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/synapse/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WorkspaceSQLAadAdminLister helps list WorkspaceSQLAadAdmins.
// All objects returned here must be treated as read-only.
type WorkspaceSQLAadAdminLister interface {
	// List lists all WorkspaceSQLAadAdmins in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WorkspaceSQLAadAdmin, err error)
	// WorkspaceSQLAadAdmins returns an object that can list and get WorkspaceSQLAadAdmins.
	WorkspaceSQLAadAdmins(namespace string) WorkspaceSQLAadAdminNamespaceLister
	WorkspaceSQLAadAdminListerExpansion
}

// workspaceSQLAadAdminLister implements the WorkspaceSQLAadAdminLister interface.
type workspaceSQLAadAdminLister struct {
	indexer cache.Indexer
}

// NewWorkspaceSQLAadAdminLister returns a new WorkspaceSQLAadAdminLister.
func NewWorkspaceSQLAadAdminLister(indexer cache.Indexer) WorkspaceSQLAadAdminLister {
	return &workspaceSQLAadAdminLister{indexer: indexer}
}

// List lists all WorkspaceSQLAadAdmins in the indexer.
func (s *workspaceSQLAadAdminLister) List(selector labels.Selector) (ret []*v1alpha1.WorkspaceSQLAadAdmin, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WorkspaceSQLAadAdmin))
	})
	return ret, err
}

// WorkspaceSQLAadAdmins returns an object that can list and get WorkspaceSQLAadAdmins.
func (s *workspaceSQLAadAdminLister) WorkspaceSQLAadAdmins(namespace string) WorkspaceSQLAadAdminNamespaceLister {
	return workspaceSQLAadAdminNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WorkspaceSQLAadAdminNamespaceLister helps list and get WorkspaceSQLAadAdmins.
// All objects returned here must be treated as read-only.
type WorkspaceSQLAadAdminNamespaceLister interface {
	// List lists all WorkspaceSQLAadAdmins in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WorkspaceSQLAadAdmin, err error)
	// Get retrieves the WorkspaceSQLAadAdmin from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.WorkspaceSQLAadAdmin, error)
	WorkspaceSQLAadAdminNamespaceListerExpansion
}

// workspaceSQLAadAdminNamespaceLister implements the WorkspaceSQLAadAdminNamespaceLister
// interface.
type workspaceSQLAadAdminNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WorkspaceSQLAadAdmins in the indexer for a given namespace.
func (s workspaceSQLAadAdminNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WorkspaceSQLAadAdmin, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WorkspaceSQLAadAdmin))
	})
	return ret, err
}

// Get retrieves the WorkspaceSQLAadAdmin from the indexer for a given namespace and name.
func (s workspaceSQLAadAdminNamespaceLister) Get(name string) (*v1alpha1.WorkspaceSQLAadAdmin, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("workspacesqlaadadmin"), name)
	}
	return obj.(*v1alpha1.WorkspaceSQLAadAdmin), nil
}
