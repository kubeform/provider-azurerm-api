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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/trafficmanager/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AzureEndpointLister helps list AzureEndpoints.
// All objects returned here must be treated as read-only.
type AzureEndpointLister interface {
	// List lists all AzureEndpoints in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AzureEndpoint, err error)
	// AzureEndpoints returns an object that can list and get AzureEndpoints.
	AzureEndpoints(namespace string) AzureEndpointNamespaceLister
	AzureEndpointListerExpansion
}

// azureEndpointLister implements the AzureEndpointLister interface.
type azureEndpointLister struct {
	indexer cache.Indexer
}

// NewAzureEndpointLister returns a new AzureEndpointLister.
func NewAzureEndpointLister(indexer cache.Indexer) AzureEndpointLister {
	return &azureEndpointLister{indexer: indexer}
}

// List lists all AzureEndpoints in the indexer.
func (s *azureEndpointLister) List(selector labels.Selector) (ret []*v1alpha1.AzureEndpoint, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AzureEndpoint))
	})
	return ret, err
}

// AzureEndpoints returns an object that can list and get AzureEndpoints.
func (s *azureEndpointLister) AzureEndpoints(namespace string) AzureEndpointNamespaceLister {
	return azureEndpointNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AzureEndpointNamespaceLister helps list and get AzureEndpoints.
// All objects returned here must be treated as read-only.
type AzureEndpointNamespaceLister interface {
	// List lists all AzureEndpoints in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AzureEndpoint, err error)
	// Get retrieves the AzureEndpoint from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AzureEndpoint, error)
	AzureEndpointNamespaceListerExpansion
}

// azureEndpointNamespaceLister implements the AzureEndpointNamespaceLister
// interface.
type azureEndpointNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AzureEndpoints in the indexer for a given namespace.
func (s azureEndpointNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AzureEndpoint, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AzureEndpoint))
	})
	return ret, err
}

// Get retrieves the AzureEndpoint from the indexer for a given namespace and name.
func (s azureEndpointNamespaceLister) Get(name string) (*v1alpha1.AzureEndpoint, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("azureendpoint"), name)
	}
	return obj.(*v1alpha1.AzureEndpoint), nil
}