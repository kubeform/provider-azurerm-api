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

// MachineScaleSetLister helps list MachineScaleSets.
// All objects returned here must be treated as read-only.
type MachineScaleSetLister interface {
	// List lists all MachineScaleSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MachineScaleSet, err error)
	// MachineScaleSets returns an object that can list and get MachineScaleSets.
	MachineScaleSets(namespace string) MachineScaleSetNamespaceLister
	MachineScaleSetListerExpansion
}

// machineScaleSetLister implements the MachineScaleSetLister interface.
type machineScaleSetLister struct {
	indexer cache.Indexer
}

// NewMachineScaleSetLister returns a new MachineScaleSetLister.
func NewMachineScaleSetLister(indexer cache.Indexer) MachineScaleSetLister {
	return &machineScaleSetLister{indexer: indexer}
}

// List lists all MachineScaleSets in the indexer.
func (s *machineScaleSetLister) List(selector labels.Selector) (ret []*v1alpha1.MachineScaleSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MachineScaleSet))
	})
	return ret, err
}

// MachineScaleSets returns an object that can list and get MachineScaleSets.
func (s *machineScaleSetLister) MachineScaleSets(namespace string) MachineScaleSetNamespaceLister {
	return machineScaleSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MachineScaleSetNamespaceLister helps list and get MachineScaleSets.
// All objects returned here must be treated as read-only.
type MachineScaleSetNamespaceLister interface {
	// List lists all MachineScaleSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MachineScaleSet, err error)
	// Get retrieves the MachineScaleSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.MachineScaleSet, error)
	MachineScaleSetNamespaceListerExpansion
}

// machineScaleSetNamespaceLister implements the MachineScaleSetNamespaceLister
// interface.
type machineScaleSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MachineScaleSets in the indexer for a given namespace.
func (s machineScaleSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MachineScaleSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MachineScaleSet))
	})
	return ret, err
}

// Get retrieves the MachineScaleSet from the indexer for a given namespace and name.
func (s machineScaleSetNamespaceLister) Get(name string) (*v1alpha1.MachineScaleSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("machinescaleset"), name)
	}
	return obj.(*v1alpha1.MachineScaleSet), nil
}
