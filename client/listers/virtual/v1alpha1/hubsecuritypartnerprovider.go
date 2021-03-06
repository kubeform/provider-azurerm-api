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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/virtual/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HubSecurityPartnerProviderLister helps list HubSecurityPartnerProviders.
// All objects returned here must be treated as read-only.
type HubSecurityPartnerProviderLister interface {
	// List lists all HubSecurityPartnerProviders in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HubSecurityPartnerProvider, err error)
	// HubSecurityPartnerProviders returns an object that can list and get HubSecurityPartnerProviders.
	HubSecurityPartnerProviders(namespace string) HubSecurityPartnerProviderNamespaceLister
	HubSecurityPartnerProviderListerExpansion
}

// hubSecurityPartnerProviderLister implements the HubSecurityPartnerProviderLister interface.
type hubSecurityPartnerProviderLister struct {
	indexer cache.Indexer
}

// NewHubSecurityPartnerProviderLister returns a new HubSecurityPartnerProviderLister.
func NewHubSecurityPartnerProviderLister(indexer cache.Indexer) HubSecurityPartnerProviderLister {
	return &hubSecurityPartnerProviderLister{indexer: indexer}
}

// List lists all HubSecurityPartnerProviders in the indexer.
func (s *hubSecurityPartnerProviderLister) List(selector labels.Selector) (ret []*v1alpha1.HubSecurityPartnerProvider, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HubSecurityPartnerProvider))
	})
	return ret, err
}

// HubSecurityPartnerProviders returns an object that can list and get HubSecurityPartnerProviders.
func (s *hubSecurityPartnerProviderLister) HubSecurityPartnerProviders(namespace string) HubSecurityPartnerProviderNamespaceLister {
	return hubSecurityPartnerProviderNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HubSecurityPartnerProviderNamespaceLister helps list and get HubSecurityPartnerProviders.
// All objects returned here must be treated as read-only.
type HubSecurityPartnerProviderNamespaceLister interface {
	// List lists all HubSecurityPartnerProviders in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HubSecurityPartnerProvider, err error)
	// Get retrieves the HubSecurityPartnerProvider from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.HubSecurityPartnerProvider, error)
	HubSecurityPartnerProviderNamespaceListerExpansion
}

// hubSecurityPartnerProviderNamespaceLister implements the HubSecurityPartnerProviderNamespaceLister
// interface.
type hubSecurityPartnerProviderNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HubSecurityPartnerProviders in the indexer for a given namespace.
func (s hubSecurityPartnerProviderNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HubSecurityPartnerProvider, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HubSecurityPartnerProvider))
	})
	return ret, err
}

// Get retrieves the HubSecurityPartnerProvider from the indexer for a given namespace and name.
func (s hubSecurityPartnerProviderNamespaceLister) Get(name string) (*v1alpha1.HubSecurityPartnerProvider, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("hubsecuritypartnerprovider"), name)
	}
	return obj.(*v1alpha1.HubSecurityPartnerProvider), nil
}
