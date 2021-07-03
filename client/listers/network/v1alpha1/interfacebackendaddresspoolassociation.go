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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/network/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InterfaceBackendAddressPoolAssociationLister helps list InterfaceBackendAddressPoolAssociations.
// All objects returned here must be treated as read-only.
type InterfaceBackendAddressPoolAssociationLister interface {
	// List lists all InterfaceBackendAddressPoolAssociations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InterfaceBackendAddressPoolAssociation, err error)
	// InterfaceBackendAddressPoolAssociations returns an object that can list and get InterfaceBackendAddressPoolAssociations.
	InterfaceBackendAddressPoolAssociations(namespace string) InterfaceBackendAddressPoolAssociationNamespaceLister
	InterfaceBackendAddressPoolAssociationListerExpansion
}

// interfaceBackendAddressPoolAssociationLister implements the InterfaceBackendAddressPoolAssociationLister interface.
type interfaceBackendAddressPoolAssociationLister struct {
	indexer cache.Indexer
}

// NewInterfaceBackendAddressPoolAssociationLister returns a new InterfaceBackendAddressPoolAssociationLister.
func NewInterfaceBackendAddressPoolAssociationLister(indexer cache.Indexer) InterfaceBackendAddressPoolAssociationLister {
	return &interfaceBackendAddressPoolAssociationLister{indexer: indexer}
}

// List lists all InterfaceBackendAddressPoolAssociations in the indexer.
func (s *interfaceBackendAddressPoolAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.InterfaceBackendAddressPoolAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InterfaceBackendAddressPoolAssociation))
	})
	return ret, err
}

// InterfaceBackendAddressPoolAssociations returns an object that can list and get InterfaceBackendAddressPoolAssociations.
func (s *interfaceBackendAddressPoolAssociationLister) InterfaceBackendAddressPoolAssociations(namespace string) InterfaceBackendAddressPoolAssociationNamespaceLister {
	return interfaceBackendAddressPoolAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InterfaceBackendAddressPoolAssociationNamespaceLister helps list and get InterfaceBackendAddressPoolAssociations.
// All objects returned here must be treated as read-only.
type InterfaceBackendAddressPoolAssociationNamespaceLister interface {
	// List lists all InterfaceBackendAddressPoolAssociations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InterfaceBackendAddressPoolAssociation, err error)
	// Get retrieves the InterfaceBackendAddressPoolAssociation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.InterfaceBackendAddressPoolAssociation, error)
	InterfaceBackendAddressPoolAssociationNamespaceListerExpansion
}

// interfaceBackendAddressPoolAssociationNamespaceLister implements the InterfaceBackendAddressPoolAssociationNamespaceLister
// interface.
type interfaceBackendAddressPoolAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all InterfaceBackendAddressPoolAssociations in the indexer for a given namespace.
func (s interfaceBackendAddressPoolAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.InterfaceBackendAddressPoolAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InterfaceBackendAddressPoolAssociation))
	})
	return ret, err
}

// Get retrieves the InterfaceBackendAddressPoolAssociation from the indexer for a given namespace and name.
func (s interfaceBackendAddressPoolAssociationNamespaceLister) Get(name string) (*v1alpha1.InterfaceBackendAddressPoolAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("interfacebackendaddresspoolassociation"), name)
	}
	return obj.(*v1alpha1.InterfaceBackendAddressPoolAssociation), nil
}
