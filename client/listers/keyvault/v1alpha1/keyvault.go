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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/keyvault/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KeyVaultLister helps list KeyVaults.
// All objects returned here must be treated as read-only.
type KeyVaultLister interface {
	// List lists all KeyVaults in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.KeyVault, err error)
	// KeyVaults returns an object that can list and get KeyVaults.
	KeyVaults(namespace string) KeyVaultNamespaceLister
	KeyVaultListerExpansion
}

// keyVaultLister implements the KeyVaultLister interface.
type keyVaultLister struct {
	indexer cache.Indexer
}

// NewKeyVaultLister returns a new KeyVaultLister.
func NewKeyVaultLister(indexer cache.Indexer) KeyVaultLister {
	return &keyVaultLister{indexer: indexer}
}

// List lists all KeyVaults in the indexer.
func (s *keyVaultLister) List(selector labels.Selector) (ret []*v1alpha1.KeyVault, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KeyVault))
	})
	return ret, err
}

// KeyVaults returns an object that can list and get KeyVaults.
func (s *keyVaultLister) KeyVaults(namespace string) KeyVaultNamespaceLister {
	return keyVaultNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KeyVaultNamespaceLister helps list and get KeyVaults.
// All objects returned here must be treated as read-only.
type KeyVaultNamespaceLister interface {
	// List lists all KeyVaults in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.KeyVault, err error)
	// Get retrieves the KeyVault from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.KeyVault, error)
	KeyVaultNamespaceListerExpansion
}

// keyVaultNamespaceLister implements the KeyVaultNamespaceLister
// interface.
type keyVaultNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KeyVaults in the indexer for a given namespace.
func (s keyVaultNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.KeyVault, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KeyVault))
	})
	return ret, err
}

// Get retrieves the KeyVault from the indexer for a given namespace and name.
func (s keyVaultNamespaceLister) Get(name string) (*v1alpha1.KeyVault, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("keyvault"), name)
	}
	return obj.(*v1alpha1.KeyVault), nil
}
