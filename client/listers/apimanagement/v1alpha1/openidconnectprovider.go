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

// OpenidConnectProviderLister helps list OpenidConnectProviders.
// All objects returned here must be treated as read-only.
type OpenidConnectProviderLister interface {
	// List lists all OpenidConnectProviders in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OpenidConnectProvider, err error)
	// OpenidConnectProviders returns an object that can list and get OpenidConnectProviders.
	OpenidConnectProviders(namespace string) OpenidConnectProviderNamespaceLister
	OpenidConnectProviderListerExpansion
}

// openidConnectProviderLister implements the OpenidConnectProviderLister interface.
type openidConnectProviderLister struct {
	indexer cache.Indexer
}

// NewOpenidConnectProviderLister returns a new OpenidConnectProviderLister.
func NewOpenidConnectProviderLister(indexer cache.Indexer) OpenidConnectProviderLister {
	return &openidConnectProviderLister{indexer: indexer}
}

// List lists all OpenidConnectProviders in the indexer.
func (s *openidConnectProviderLister) List(selector labels.Selector) (ret []*v1alpha1.OpenidConnectProvider, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OpenidConnectProvider))
	})
	return ret, err
}

// OpenidConnectProviders returns an object that can list and get OpenidConnectProviders.
func (s *openidConnectProviderLister) OpenidConnectProviders(namespace string) OpenidConnectProviderNamespaceLister {
	return openidConnectProviderNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OpenidConnectProviderNamespaceLister helps list and get OpenidConnectProviders.
// All objects returned here must be treated as read-only.
type OpenidConnectProviderNamespaceLister interface {
	// List lists all OpenidConnectProviders in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OpenidConnectProvider, err error)
	// Get retrieves the OpenidConnectProvider from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.OpenidConnectProvider, error)
	OpenidConnectProviderNamespaceListerExpansion
}

// openidConnectProviderNamespaceLister implements the OpenidConnectProviderNamespaceLister
// interface.
type openidConnectProviderNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OpenidConnectProviders in the indexer for a given namespace.
func (s openidConnectProviderNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.OpenidConnectProvider, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OpenidConnectProvider))
	})
	return ret, err
}

// Get retrieves the OpenidConnectProvider from the indexer for a given namespace and name.
func (s openidConnectProviderNamespaceLister) Get(name string) (*v1alpha1.OpenidConnectProvider, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("openidconnectprovider"), name)
	}
	return obj.(*v1alpha1.OpenidConnectProvider), nil
}
