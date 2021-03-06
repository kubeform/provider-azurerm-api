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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/siterecovery/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ProtectionContainerMappingLister helps list ProtectionContainerMappings.
// All objects returned here must be treated as read-only.
type ProtectionContainerMappingLister interface {
	// List lists all ProtectionContainerMappings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProtectionContainerMapping, err error)
	// ProtectionContainerMappings returns an object that can list and get ProtectionContainerMappings.
	ProtectionContainerMappings(namespace string) ProtectionContainerMappingNamespaceLister
	ProtectionContainerMappingListerExpansion
}

// protectionContainerMappingLister implements the ProtectionContainerMappingLister interface.
type protectionContainerMappingLister struct {
	indexer cache.Indexer
}

// NewProtectionContainerMappingLister returns a new ProtectionContainerMappingLister.
func NewProtectionContainerMappingLister(indexer cache.Indexer) ProtectionContainerMappingLister {
	return &protectionContainerMappingLister{indexer: indexer}
}

// List lists all ProtectionContainerMappings in the indexer.
func (s *protectionContainerMappingLister) List(selector labels.Selector) (ret []*v1alpha1.ProtectionContainerMapping, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProtectionContainerMapping))
	})
	return ret, err
}

// ProtectionContainerMappings returns an object that can list and get ProtectionContainerMappings.
func (s *protectionContainerMappingLister) ProtectionContainerMappings(namespace string) ProtectionContainerMappingNamespaceLister {
	return protectionContainerMappingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProtectionContainerMappingNamespaceLister helps list and get ProtectionContainerMappings.
// All objects returned here must be treated as read-only.
type ProtectionContainerMappingNamespaceLister interface {
	// List lists all ProtectionContainerMappings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProtectionContainerMapping, err error)
	// Get retrieves the ProtectionContainerMapping from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ProtectionContainerMapping, error)
	ProtectionContainerMappingNamespaceListerExpansion
}

// protectionContainerMappingNamespaceLister implements the ProtectionContainerMappingNamespaceLister
// interface.
type protectionContainerMappingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ProtectionContainerMappings in the indexer for a given namespace.
func (s protectionContainerMappingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ProtectionContainerMapping, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProtectionContainerMapping))
	})
	return ret, err
}

// Get retrieves the ProtectionContainerMapping from the indexer for a given namespace and name.
func (s protectionContainerMappingNamespaceLister) Get(name string) (*v1alpha1.ProtectionContainerMapping, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("protectioncontainermapping"), name)
	}
	return obj.(*v1alpha1.ProtectionContainerMapping), nil
}
