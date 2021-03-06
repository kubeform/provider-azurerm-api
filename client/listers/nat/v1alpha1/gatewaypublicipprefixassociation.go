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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/nat/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GatewayPublicIPPrefixAssociationLister helps list GatewayPublicIPPrefixAssociations.
// All objects returned here must be treated as read-only.
type GatewayPublicIPPrefixAssociationLister interface {
	// List lists all GatewayPublicIPPrefixAssociations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GatewayPublicIPPrefixAssociation, err error)
	// GatewayPublicIPPrefixAssociations returns an object that can list and get GatewayPublicIPPrefixAssociations.
	GatewayPublicIPPrefixAssociations(namespace string) GatewayPublicIPPrefixAssociationNamespaceLister
	GatewayPublicIPPrefixAssociationListerExpansion
}

// gatewayPublicIPPrefixAssociationLister implements the GatewayPublicIPPrefixAssociationLister interface.
type gatewayPublicIPPrefixAssociationLister struct {
	indexer cache.Indexer
}

// NewGatewayPublicIPPrefixAssociationLister returns a new GatewayPublicIPPrefixAssociationLister.
func NewGatewayPublicIPPrefixAssociationLister(indexer cache.Indexer) GatewayPublicIPPrefixAssociationLister {
	return &gatewayPublicIPPrefixAssociationLister{indexer: indexer}
}

// List lists all GatewayPublicIPPrefixAssociations in the indexer.
func (s *gatewayPublicIPPrefixAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.GatewayPublicIPPrefixAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GatewayPublicIPPrefixAssociation))
	})
	return ret, err
}

// GatewayPublicIPPrefixAssociations returns an object that can list and get GatewayPublicIPPrefixAssociations.
func (s *gatewayPublicIPPrefixAssociationLister) GatewayPublicIPPrefixAssociations(namespace string) GatewayPublicIPPrefixAssociationNamespaceLister {
	return gatewayPublicIPPrefixAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GatewayPublicIPPrefixAssociationNamespaceLister helps list and get GatewayPublicIPPrefixAssociations.
// All objects returned here must be treated as read-only.
type GatewayPublicIPPrefixAssociationNamespaceLister interface {
	// List lists all GatewayPublicIPPrefixAssociations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GatewayPublicIPPrefixAssociation, err error)
	// Get retrieves the GatewayPublicIPPrefixAssociation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.GatewayPublicIPPrefixAssociation, error)
	GatewayPublicIPPrefixAssociationNamespaceListerExpansion
}

// gatewayPublicIPPrefixAssociationNamespaceLister implements the GatewayPublicIPPrefixAssociationNamespaceLister
// interface.
type gatewayPublicIPPrefixAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GatewayPublicIPPrefixAssociations in the indexer for a given namespace.
func (s gatewayPublicIPPrefixAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GatewayPublicIPPrefixAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GatewayPublicIPPrefixAssociation))
	})
	return ret, err
}

// Get retrieves the GatewayPublicIPPrefixAssociation from the indexer for a given namespace and name.
func (s gatewayPublicIPPrefixAssociationNamespaceLister) Get(name string) (*v1alpha1.GatewayPublicIPPrefixAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("gatewaypublicipprefixassociation"), name)
	}
	return obj.(*v1alpha1.GatewayPublicIPPrefixAssociation), nil
}
