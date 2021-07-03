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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/windows/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VirtualMachineScaleSetLister helps list VirtualMachineScaleSets.
// All objects returned here must be treated as read-only.
type VirtualMachineScaleSetLister interface {
	// List lists all VirtualMachineScaleSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VirtualMachineScaleSet, err error)
	// VirtualMachineScaleSets returns an object that can list and get VirtualMachineScaleSets.
	VirtualMachineScaleSets(namespace string) VirtualMachineScaleSetNamespaceLister
	VirtualMachineScaleSetListerExpansion
}

// virtualMachineScaleSetLister implements the VirtualMachineScaleSetLister interface.
type virtualMachineScaleSetLister struct {
	indexer cache.Indexer
}

// NewVirtualMachineScaleSetLister returns a new VirtualMachineScaleSetLister.
func NewVirtualMachineScaleSetLister(indexer cache.Indexer) VirtualMachineScaleSetLister {
	return &virtualMachineScaleSetLister{indexer: indexer}
}

// List lists all VirtualMachineScaleSets in the indexer.
func (s *virtualMachineScaleSetLister) List(selector labels.Selector) (ret []*v1alpha1.VirtualMachineScaleSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VirtualMachineScaleSet))
	})
	return ret, err
}

// VirtualMachineScaleSets returns an object that can list and get VirtualMachineScaleSets.
func (s *virtualMachineScaleSetLister) VirtualMachineScaleSets(namespace string) VirtualMachineScaleSetNamespaceLister {
	return virtualMachineScaleSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VirtualMachineScaleSetNamespaceLister helps list and get VirtualMachineScaleSets.
// All objects returned here must be treated as read-only.
type VirtualMachineScaleSetNamespaceLister interface {
	// List lists all VirtualMachineScaleSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VirtualMachineScaleSet, err error)
	// Get retrieves the VirtualMachineScaleSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.VirtualMachineScaleSet, error)
	VirtualMachineScaleSetNamespaceListerExpansion
}

// virtualMachineScaleSetNamespaceLister implements the VirtualMachineScaleSetNamespaceLister
// interface.
type virtualMachineScaleSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VirtualMachineScaleSets in the indexer for a given namespace.
func (s virtualMachineScaleSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VirtualMachineScaleSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VirtualMachineScaleSet))
	})
	return ret, err
}

// Get retrieves the VirtualMachineScaleSet from the indexer for a given namespace and name.
func (s virtualMachineScaleSetNamespaceLister) Get(name string) (*v1alpha1.VirtualMachineScaleSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("virtualmachinescaleset"), name)
	}
	return obj.(*v1alpha1.VirtualMachineScaleSet), nil
}
