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

// MachineExtensionLister helps list MachineExtensions.
// All objects returned here must be treated as read-only.
type MachineExtensionLister interface {
	// List lists all MachineExtensions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MachineExtension, err error)
	// MachineExtensions returns an object that can list and get MachineExtensions.
	MachineExtensions(namespace string) MachineExtensionNamespaceLister
	MachineExtensionListerExpansion
}

// machineExtensionLister implements the MachineExtensionLister interface.
type machineExtensionLister struct {
	indexer cache.Indexer
}

// NewMachineExtensionLister returns a new MachineExtensionLister.
func NewMachineExtensionLister(indexer cache.Indexer) MachineExtensionLister {
	return &machineExtensionLister{indexer: indexer}
}

// List lists all MachineExtensions in the indexer.
func (s *machineExtensionLister) List(selector labels.Selector) (ret []*v1alpha1.MachineExtension, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MachineExtension))
	})
	return ret, err
}

// MachineExtensions returns an object that can list and get MachineExtensions.
func (s *machineExtensionLister) MachineExtensions(namespace string) MachineExtensionNamespaceLister {
	return machineExtensionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MachineExtensionNamespaceLister helps list and get MachineExtensions.
// All objects returned here must be treated as read-only.
type MachineExtensionNamespaceLister interface {
	// List lists all MachineExtensions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MachineExtension, err error)
	// Get retrieves the MachineExtension from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.MachineExtension, error)
	MachineExtensionNamespaceListerExpansion
}

// machineExtensionNamespaceLister implements the MachineExtensionNamespaceLister
// interface.
type machineExtensionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MachineExtensions in the indexer for a given namespace.
func (s machineExtensionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MachineExtension, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MachineExtension))
	})
	return ret, err
}

// Get retrieves the MachineExtension from the indexer for a given namespace and name.
func (s machineExtensionNamespaceLister) Get(name string) (*v1alpha1.MachineExtension, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("machineextension"), name)
	}
	return obj.(*v1alpha1.MachineExtension), nil
}