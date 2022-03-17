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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/mssql/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ManagedInstanceLister helps list ManagedInstances.
// All objects returned here must be treated as read-only.
type ManagedInstanceLister interface {
	// List lists all ManagedInstances in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagedInstance, err error)
	// ManagedInstances returns an object that can list and get ManagedInstances.
	ManagedInstances(namespace string) ManagedInstanceNamespaceLister
	ManagedInstanceListerExpansion
}

// managedInstanceLister implements the ManagedInstanceLister interface.
type managedInstanceLister struct {
	indexer cache.Indexer
}

// NewManagedInstanceLister returns a new ManagedInstanceLister.
func NewManagedInstanceLister(indexer cache.Indexer) ManagedInstanceLister {
	return &managedInstanceLister{indexer: indexer}
}

// List lists all ManagedInstances in the indexer.
func (s *managedInstanceLister) List(selector labels.Selector) (ret []*v1alpha1.ManagedInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagedInstance))
	})
	return ret, err
}

// ManagedInstances returns an object that can list and get ManagedInstances.
func (s *managedInstanceLister) ManagedInstances(namespace string) ManagedInstanceNamespaceLister {
	return managedInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ManagedInstanceNamespaceLister helps list and get ManagedInstances.
// All objects returned here must be treated as read-only.
type ManagedInstanceNamespaceLister interface {
	// List lists all ManagedInstances in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagedInstance, err error)
	// Get retrieves the ManagedInstance from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ManagedInstance, error)
	ManagedInstanceNamespaceListerExpansion
}

// managedInstanceNamespaceLister implements the ManagedInstanceNamespaceLister
// interface.
type managedInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ManagedInstances in the indexer for a given namespace.
func (s managedInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ManagedInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagedInstance))
	})
	return ret, err
}

// Get retrieves the ManagedInstance from the indexer for a given namespace and name.
func (s managedInstanceNamespaceLister) Get(name string) (*v1alpha1.ManagedInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("managedinstance"), name)
	}
	return obj.(*v1alpha1.ManagedInstance), nil
}