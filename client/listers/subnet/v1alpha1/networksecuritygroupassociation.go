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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/subnet/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NetworkSecurityGroupAssociationLister helps list NetworkSecurityGroupAssociations.
// All objects returned here must be treated as read-only.
type NetworkSecurityGroupAssociationLister interface {
	// List lists all NetworkSecurityGroupAssociations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkSecurityGroupAssociation, err error)
	// NetworkSecurityGroupAssociations returns an object that can list and get NetworkSecurityGroupAssociations.
	NetworkSecurityGroupAssociations(namespace string) NetworkSecurityGroupAssociationNamespaceLister
	NetworkSecurityGroupAssociationListerExpansion
}

// networkSecurityGroupAssociationLister implements the NetworkSecurityGroupAssociationLister interface.
type networkSecurityGroupAssociationLister struct {
	indexer cache.Indexer
}

// NewNetworkSecurityGroupAssociationLister returns a new NetworkSecurityGroupAssociationLister.
func NewNetworkSecurityGroupAssociationLister(indexer cache.Indexer) NetworkSecurityGroupAssociationLister {
	return &networkSecurityGroupAssociationLister{indexer: indexer}
}

// List lists all NetworkSecurityGroupAssociations in the indexer.
func (s *networkSecurityGroupAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkSecurityGroupAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkSecurityGroupAssociation))
	})
	return ret, err
}

// NetworkSecurityGroupAssociations returns an object that can list and get NetworkSecurityGroupAssociations.
func (s *networkSecurityGroupAssociationLister) NetworkSecurityGroupAssociations(namespace string) NetworkSecurityGroupAssociationNamespaceLister {
	return networkSecurityGroupAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NetworkSecurityGroupAssociationNamespaceLister helps list and get NetworkSecurityGroupAssociations.
// All objects returned here must be treated as read-only.
type NetworkSecurityGroupAssociationNamespaceLister interface {
	// List lists all NetworkSecurityGroupAssociations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkSecurityGroupAssociation, err error)
	// Get retrieves the NetworkSecurityGroupAssociation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.NetworkSecurityGroupAssociation, error)
	NetworkSecurityGroupAssociationNamespaceListerExpansion
}

// networkSecurityGroupAssociationNamespaceLister implements the NetworkSecurityGroupAssociationNamespaceLister
// interface.
type networkSecurityGroupAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NetworkSecurityGroupAssociations in the indexer for a given namespace.
func (s networkSecurityGroupAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkSecurityGroupAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkSecurityGroupAssociation))
	})
	return ret, err
}

// Get retrieves the NetworkSecurityGroupAssociation from the indexer for a given namespace and name.
func (s networkSecurityGroupAssociationNamespaceLister) Get(name string) (*v1alpha1.NetworkSecurityGroupAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("networksecuritygroupassociation"), name)
	}
	return obj.(*v1alpha1.NetworkSecurityGroupAssociation), nil
}
