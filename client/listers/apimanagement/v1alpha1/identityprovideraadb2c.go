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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/apimanagement/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IdentityProviderAadb2cLister helps list IdentityProviderAadb2cs.
// All objects returned here must be treated as read-only.
type IdentityProviderAadb2cLister interface {
	// List lists all IdentityProviderAadb2cs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IdentityProviderAadb2c, err error)
	// IdentityProviderAadb2cs returns an object that can list and get IdentityProviderAadb2cs.
	IdentityProviderAadb2cs(namespace string) IdentityProviderAadb2cNamespaceLister
	IdentityProviderAadb2cListerExpansion
}

// identityProviderAadb2cLister implements the IdentityProviderAadb2cLister interface.
type identityProviderAadb2cLister struct {
	indexer cache.Indexer
}

// NewIdentityProviderAadb2cLister returns a new IdentityProviderAadb2cLister.
func NewIdentityProviderAadb2cLister(indexer cache.Indexer) IdentityProviderAadb2cLister {
	return &identityProviderAadb2cLister{indexer: indexer}
}

// List lists all IdentityProviderAadb2cs in the indexer.
func (s *identityProviderAadb2cLister) List(selector labels.Selector) (ret []*v1alpha1.IdentityProviderAadb2c, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IdentityProviderAadb2c))
	})
	return ret, err
}

// IdentityProviderAadb2cs returns an object that can list and get IdentityProviderAadb2cs.
func (s *identityProviderAadb2cLister) IdentityProviderAadb2cs(namespace string) IdentityProviderAadb2cNamespaceLister {
	return identityProviderAadb2cNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IdentityProviderAadb2cNamespaceLister helps list and get IdentityProviderAadb2cs.
// All objects returned here must be treated as read-only.
type IdentityProviderAadb2cNamespaceLister interface {
	// List lists all IdentityProviderAadb2cs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IdentityProviderAadb2c, err error)
	// Get retrieves the IdentityProviderAadb2c from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.IdentityProviderAadb2c, error)
	IdentityProviderAadb2cNamespaceListerExpansion
}

// identityProviderAadb2cNamespaceLister implements the IdentityProviderAadb2cNamespaceLister
// interface.
type identityProviderAadb2cNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IdentityProviderAadb2cs in the indexer for a given namespace.
func (s identityProviderAadb2cNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IdentityProviderAadb2c, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IdentityProviderAadb2c))
	})
	return ret, err
}

// Get retrieves the IdentityProviderAadb2c from the indexer for a given namespace and name.
func (s identityProviderAadb2cNamespaceLister) Get(name string) (*v1alpha1.IdentityProviderAadb2c, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("identityprovideraadb2c"), name)
	}
	return obj.(*v1alpha1.IdentityProviderAadb2c), nil
}
