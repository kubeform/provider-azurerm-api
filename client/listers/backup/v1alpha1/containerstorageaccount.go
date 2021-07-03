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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/backup/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ContainerStorageAccountLister helps list ContainerStorageAccounts.
// All objects returned here must be treated as read-only.
type ContainerStorageAccountLister interface {
	// List lists all ContainerStorageAccounts in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ContainerStorageAccount, err error)
	// ContainerStorageAccounts returns an object that can list and get ContainerStorageAccounts.
	ContainerStorageAccounts(namespace string) ContainerStorageAccountNamespaceLister
	ContainerStorageAccountListerExpansion
}

// containerStorageAccountLister implements the ContainerStorageAccountLister interface.
type containerStorageAccountLister struct {
	indexer cache.Indexer
}

// NewContainerStorageAccountLister returns a new ContainerStorageAccountLister.
func NewContainerStorageAccountLister(indexer cache.Indexer) ContainerStorageAccountLister {
	return &containerStorageAccountLister{indexer: indexer}
}

// List lists all ContainerStorageAccounts in the indexer.
func (s *containerStorageAccountLister) List(selector labels.Selector) (ret []*v1alpha1.ContainerStorageAccount, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ContainerStorageAccount))
	})
	return ret, err
}

// ContainerStorageAccounts returns an object that can list and get ContainerStorageAccounts.
func (s *containerStorageAccountLister) ContainerStorageAccounts(namespace string) ContainerStorageAccountNamespaceLister {
	return containerStorageAccountNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ContainerStorageAccountNamespaceLister helps list and get ContainerStorageAccounts.
// All objects returned here must be treated as read-only.
type ContainerStorageAccountNamespaceLister interface {
	// List lists all ContainerStorageAccounts in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ContainerStorageAccount, err error)
	// Get retrieves the ContainerStorageAccount from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ContainerStorageAccount, error)
	ContainerStorageAccountNamespaceListerExpansion
}

// containerStorageAccountNamespaceLister implements the ContainerStorageAccountNamespaceLister
// interface.
type containerStorageAccountNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ContainerStorageAccounts in the indexer for a given namespace.
func (s containerStorageAccountNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ContainerStorageAccount, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ContainerStorageAccount))
	})
	return ret, err
}

// Get retrieves the ContainerStorageAccount from the indexer for a given namespace and name.
func (s containerStorageAccountNamespaceLister) Get(name string) (*v1alpha1.ContainerStorageAccount, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("containerstorageaccount"), name)
	}
	return obj.(*v1alpha1.ContainerStorageAccount), nil
}
