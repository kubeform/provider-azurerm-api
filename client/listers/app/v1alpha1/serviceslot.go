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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/app/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServiceSlotLister helps list ServiceSlots.
// All objects returned here must be treated as read-only.
type ServiceSlotLister interface {
	// List lists all ServiceSlots in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceSlot, err error)
	// ServiceSlots returns an object that can list and get ServiceSlots.
	ServiceSlots(namespace string) ServiceSlotNamespaceLister
	ServiceSlotListerExpansion
}

// serviceSlotLister implements the ServiceSlotLister interface.
type serviceSlotLister struct {
	indexer cache.Indexer
}

// NewServiceSlotLister returns a new ServiceSlotLister.
func NewServiceSlotLister(indexer cache.Indexer) ServiceSlotLister {
	return &serviceSlotLister{indexer: indexer}
}

// List lists all ServiceSlots in the indexer.
func (s *serviceSlotLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceSlot, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceSlot))
	})
	return ret, err
}

// ServiceSlots returns an object that can list and get ServiceSlots.
func (s *serviceSlotLister) ServiceSlots(namespace string) ServiceSlotNamespaceLister {
	return serviceSlotNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServiceSlotNamespaceLister helps list and get ServiceSlots.
// All objects returned here must be treated as read-only.
type ServiceSlotNamespaceLister interface {
	// List lists all ServiceSlots in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceSlot, err error)
	// Get retrieves the ServiceSlot from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ServiceSlot, error)
	ServiceSlotNamespaceListerExpansion
}

// serviceSlotNamespaceLister implements the ServiceSlotNamespaceLister
// interface.
type serviceSlotNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServiceSlots in the indexer for a given namespace.
func (s serviceSlotNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceSlot, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceSlot))
	})
	return ret, err
}

// Get retrieves the ServiceSlot from the indexer for a given namespace and name.
func (s serviceSlotNamespaceLister) Get(name string) (*v1alpha1.ServiceSlot, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("serviceslot"), name)
	}
	return obj.(*v1alpha1.ServiceSlot), nil
}
