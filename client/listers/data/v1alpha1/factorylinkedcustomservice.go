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

// FactoryLinkedCustomServiceLister helps list FactoryLinkedCustomServices.
// All objects returned here must be treated as read-only.
type FactoryLinkedCustomServiceLister interface {
	// List lists all FactoryLinkedCustomServices in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FactoryLinkedCustomService, err error)
	// FactoryLinkedCustomServices returns an object that can list and get FactoryLinkedCustomServices.
	FactoryLinkedCustomServices(namespace string) FactoryLinkedCustomServiceNamespaceLister
	FactoryLinkedCustomServiceListerExpansion
}

// factoryLinkedCustomServiceLister implements the FactoryLinkedCustomServiceLister interface.
type factoryLinkedCustomServiceLister struct {
	indexer cache.Indexer
}

// NewFactoryLinkedCustomServiceLister returns a new FactoryLinkedCustomServiceLister.
func NewFactoryLinkedCustomServiceLister(indexer cache.Indexer) FactoryLinkedCustomServiceLister {
	return &factoryLinkedCustomServiceLister{indexer: indexer}
}

// List lists all FactoryLinkedCustomServices in the indexer.
func (s *factoryLinkedCustomServiceLister) List(selector labels.Selector) (ret []*v1alpha1.FactoryLinkedCustomService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FactoryLinkedCustomService))
	})
	return ret, err
}

// FactoryLinkedCustomServices returns an object that can list and get FactoryLinkedCustomServices.
func (s *factoryLinkedCustomServiceLister) FactoryLinkedCustomServices(namespace string) FactoryLinkedCustomServiceNamespaceLister {
	return factoryLinkedCustomServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FactoryLinkedCustomServiceNamespaceLister helps list and get FactoryLinkedCustomServices.
// All objects returned here must be treated as read-only.
type FactoryLinkedCustomServiceNamespaceLister interface {
	// List lists all FactoryLinkedCustomServices in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FactoryLinkedCustomService, err error)
	// Get retrieves the FactoryLinkedCustomService from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FactoryLinkedCustomService, error)
	FactoryLinkedCustomServiceNamespaceListerExpansion
}

// factoryLinkedCustomServiceNamespaceLister implements the FactoryLinkedCustomServiceNamespaceLister
// interface.
type factoryLinkedCustomServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FactoryLinkedCustomServices in the indexer for a given namespace.
func (s factoryLinkedCustomServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FactoryLinkedCustomService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FactoryLinkedCustomService))
	})
	return ret, err
}

// Get retrieves the FactoryLinkedCustomService from the indexer for a given namespace and name.
func (s factoryLinkedCustomServiceNamespaceLister) Get(name string) (*v1alpha1.FactoryLinkedCustomService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("factorylinkedcustomservice"), name)
	}
	return obj.(*v1alpha1.FactoryLinkedCustomService), nil
}