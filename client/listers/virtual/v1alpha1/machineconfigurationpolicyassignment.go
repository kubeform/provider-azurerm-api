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

// MachineConfigurationPolicyAssignmentLister helps list MachineConfigurationPolicyAssignments.
// All objects returned here must be treated as read-only.
type MachineConfigurationPolicyAssignmentLister interface {
	// List lists all MachineConfigurationPolicyAssignments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MachineConfigurationPolicyAssignment, err error)
	// MachineConfigurationPolicyAssignments returns an object that can list and get MachineConfigurationPolicyAssignments.
	MachineConfigurationPolicyAssignments(namespace string) MachineConfigurationPolicyAssignmentNamespaceLister
	MachineConfigurationPolicyAssignmentListerExpansion
}

// machineConfigurationPolicyAssignmentLister implements the MachineConfigurationPolicyAssignmentLister interface.
type machineConfigurationPolicyAssignmentLister struct {
	indexer cache.Indexer
}

// NewMachineConfigurationPolicyAssignmentLister returns a new MachineConfigurationPolicyAssignmentLister.
func NewMachineConfigurationPolicyAssignmentLister(indexer cache.Indexer) MachineConfigurationPolicyAssignmentLister {
	return &machineConfigurationPolicyAssignmentLister{indexer: indexer}
}

// List lists all MachineConfigurationPolicyAssignments in the indexer.
func (s *machineConfigurationPolicyAssignmentLister) List(selector labels.Selector) (ret []*v1alpha1.MachineConfigurationPolicyAssignment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MachineConfigurationPolicyAssignment))
	})
	return ret, err
}

// MachineConfigurationPolicyAssignments returns an object that can list and get MachineConfigurationPolicyAssignments.
func (s *machineConfigurationPolicyAssignmentLister) MachineConfigurationPolicyAssignments(namespace string) MachineConfigurationPolicyAssignmentNamespaceLister {
	return machineConfigurationPolicyAssignmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MachineConfigurationPolicyAssignmentNamespaceLister helps list and get MachineConfigurationPolicyAssignments.
// All objects returned here must be treated as read-only.
type MachineConfigurationPolicyAssignmentNamespaceLister interface {
	// List lists all MachineConfigurationPolicyAssignments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MachineConfigurationPolicyAssignment, err error)
	// Get retrieves the MachineConfigurationPolicyAssignment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.MachineConfigurationPolicyAssignment, error)
	MachineConfigurationPolicyAssignmentNamespaceListerExpansion
}

// machineConfigurationPolicyAssignmentNamespaceLister implements the MachineConfigurationPolicyAssignmentNamespaceLister
// interface.
type machineConfigurationPolicyAssignmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MachineConfigurationPolicyAssignments in the indexer for a given namespace.
func (s machineConfigurationPolicyAssignmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MachineConfigurationPolicyAssignment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MachineConfigurationPolicyAssignment))
	})
	return ret, err
}

// Get retrieves the MachineConfigurationPolicyAssignment from the indexer for a given namespace and name.
func (s machineConfigurationPolicyAssignmentNamespaceLister) Get(name string) (*v1alpha1.MachineConfigurationPolicyAssignment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("machineconfigurationpolicyassignment"), name)
	}
	return obj.(*v1alpha1.MachineConfigurationPolicyAssignment), nil
}
