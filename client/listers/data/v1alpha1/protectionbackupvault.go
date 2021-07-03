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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/data/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ProtectionBackupVaultLister helps list ProtectionBackupVaults.
// All objects returned here must be treated as read-only.
type ProtectionBackupVaultLister interface {
	// List lists all ProtectionBackupVaults in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProtectionBackupVault, err error)
	// ProtectionBackupVaults returns an object that can list and get ProtectionBackupVaults.
	ProtectionBackupVaults(namespace string) ProtectionBackupVaultNamespaceLister
	ProtectionBackupVaultListerExpansion
}

// protectionBackupVaultLister implements the ProtectionBackupVaultLister interface.
type protectionBackupVaultLister struct {
	indexer cache.Indexer
}

// NewProtectionBackupVaultLister returns a new ProtectionBackupVaultLister.
func NewProtectionBackupVaultLister(indexer cache.Indexer) ProtectionBackupVaultLister {
	return &protectionBackupVaultLister{indexer: indexer}
}

// List lists all ProtectionBackupVaults in the indexer.
func (s *protectionBackupVaultLister) List(selector labels.Selector) (ret []*v1alpha1.ProtectionBackupVault, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProtectionBackupVault))
	})
	return ret, err
}

// ProtectionBackupVaults returns an object that can list and get ProtectionBackupVaults.
func (s *protectionBackupVaultLister) ProtectionBackupVaults(namespace string) ProtectionBackupVaultNamespaceLister {
	return protectionBackupVaultNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProtectionBackupVaultNamespaceLister helps list and get ProtectionBackupVaults.
// All objects returned here must be treated as read-only.
type ProtectionBackupVaultNamespaceLister interface {
	// List lists all ProtectionBackupVaults in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProtectionBackupVault, err error)
	// Get retrieves the ProtectionBackupVault from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ProtectionBackupVault, error)
	ProtectionBackupVaultNamespaceListerExpansion
}

// protectionBackupVaultNamespaceLister implements the ProtectionBackupVaultNamespaceLister
// interface.
type protectionBackupVaultNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ProtectionBackupVaults in the indexer for a given namespace.
func (s protectionBackupVaultNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ProtectionBackupVault, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProtectionBackupVault))
	})
	return ret, err
}

// Get retrieves the ProtectionBackupVault from the indexer for a given namespace and name.
func (s protectionBackupVaultNamespaceLister) Get(name string) (*v1alpha1.ProtectionBackupVault, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("protectionbackupvault"), name)
	}
	return obj.(*v1alpha1.ProtectionBackupVault), nil
}
