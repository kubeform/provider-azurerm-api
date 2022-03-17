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

// ProtectionBackupInstanceBlobStorageLister helps list ProtectionBackupInstanceBlobStorages.
// All objects returned here must be treated as read-only.
type ProtectionBackupInstanceBlobStorageLister interface {
	// List lists all ProtectionBackupInstanceBlobStorages in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProtectionBackupInstanceBlobStorage, err error)
	// ProtectionBackupInstanceBlobStorages returns an object that can list and get ProtectionBackupInstanceBlobStorages.
	ProtectionBackupInstanceBlobStorages(namespace string) ProtectionBackupInstanceBlobStorageNamespaceLister
	ProtectionBackupInstanceBlobStorageListerExpansion
}

// protectionBackupInstanceBlobStorageLister implements the ProtectionBackupInstanceBlobStorageLister interface.
type protectionBackupInstanceBlobStorageLister struct {
	indexer cache.Indexer
}

// NewProtectionBackupInstanceBlobStorageLister returns a new ProtectionBackupInstanceBlobStorageLister.
func NewProtectionBackupInstanceBlobStorageLister(indexer cache.Indexer) ProtectionBackupInstanceBlobStorageLister {
	return &protectionBackupInstanceBlobStorageLister{indexer: indexer}
}

// List lists all ProtectionBackupInstanceBlobStorages in the indexer.
func (s *protectionBackupInstanceBlobStorageLister) List(selector labels.Selector) (ret []*v1alpha1.ProtectionBackupInstanceBlobStorage, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProtectionBackupInstanceBlobStorage))
	})
	return ret, err
}

// ProtectionBackupInstanceBlobStorages returns an object that can list and get ProtectionBackupInstanceBlobStorages.
func (s *protectionBackupInstanceBlobStorageLister) ProtectionBackupInstanceBlobStorages(namespace string) ProtectionBackupInstanceBlobStorageNamespaceLister {
	return protectionBackupInstanceBlobStorageNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProtectionBackupInstanceBlobStorageNamespaceLister helps list and get ProtectionBackupInstanceBlobStorages.
// All objects returned here must be treated as read-only.
type ProtectionBackupInstanceBlobStorageNamespaceLister interface {
	// List lists all ProtectionBackupInstanceBlobStorages in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProtectionBackupInstanceBlobStorage, err error)
	// Get retrieves the ProtectionBackupInstanceBlobStorage from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ProtectionBackupInstanceBlobStorage, error)
	ProtectionBackupInstanceBlobStorageNamespaceListerExpansion
}

// protectionBackupInstanceBlobStorageNamespaceLister implements the ProtectionBackupInstanceBlobStorageNamespaceLister
// interface.
type protectionBackupInstanceBlobStorageNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ProtectionBackupInstanceBlobStorages in the indexer for a given namespace.
func (s protectionBackupInstanceBlobStorageNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ProtectionBackupInstanceBlobStorage, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProtectionBackupInstanceBlobStorage))
	})
	return ret, err
}

// Get retrieves the ProtectionBackupInstanceBlobStorage from the indexer for a given namespace and name.
func (s protectionBackupInstanceBlobStorageNamespaceLister) Get(name string) (*v1alpha1.ProtectionBackupInstanceBlobStorage, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("protectionbackupinstanceblobstorage"), name)
	}
	return obj.(*v1alpha1.ProtectionBackupInstanceBlobStorage), nil
}