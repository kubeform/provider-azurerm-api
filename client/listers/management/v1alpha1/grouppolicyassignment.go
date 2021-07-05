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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/management/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GroupPolicyAssignmentLister helps list GroupPolicyAssignments.
// All objects returned here must be treated as read-only.
type GroupPolicyAssignmentLister interface {
	// List lists all GroupPolicyAssignments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GroupPolicyAssignment, err error)
	// GroupPolicyAssignments returns an object that can list and get GroupPolicyAssignments.
	GroupPolicyAssignments(namespace string) GroupPolicyAssignmentNamespaceLister
	GroupPolicyAssignmentListerExpansion
}

// groupPolicyAssignmentLister implements the GroupPolicyAssignmentLister interface.
type groupPolicyAssignmentLister struct {
	indexer cache.Indexer
}

// NewGroupPolicyAssignmentLister returns a new GroupPolicyAssignmentLister.
func NewGroupPolicyAssignmentLister(indexer cache.Indexer) GroupPolicyAssignmentLister {
	return &groupPolicyAssignmentLister{indexer: indexer}
}

// List lists all GroupPolicyAssignments in the indexer.
func (s *groupPolicyAssignmentLister) List(selector labels.Selector) (ret []*v1alpha1.GroupPolicyAssignment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GroupPolicyAssignment))
	})
	return ret, err
}

// GroupPolicyAssignments returns an object that can list and get GroupPolicyAssignments.
func (s *groupPolicyAssignmentLister) GroupPolicyAssignments(namespace string) GroupPolicyAssignmentNamespaceLister {
	return groupPolicyAssignmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GroupPolicyAssignmentNamespaceLister helps list and get GroupPolicyAssignments.
// All objects returned here must be treated as read-only.
type GroupPolicyAssignmentNamespaceLister interface {
	// List lists all GroupPolicyAssignments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GroupPolicyAssignment, err error)
	// Get retrieves the GroupPolicyAssignment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.GroupPolicyAssignment, error)
	GroupPolicyAssignmentNamespaceListerExpansion
}

// groupPolicyAssignmentNamespaceLister implements the GroupPolicyAssignmentNamespaceLister
// interface.
type groupPolicyAssignmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GroupPolicyAssignments in the indexer for a given namespace.
func (s groupPolicyAssignmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GroupPolicyAssignment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GroupPolicyAssignment))
	})
	return ret, err
}

// Get retrieves the GroupPolicyAssignment from the indexer for a given namespace and name.
func (s groupPolicyAssignmentNamespaceLister) Get(name string) (*v1alpha1.GroupPolicyAssignment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("grouppolicyassignment"), name)
	}
	return obj.(*v1alpha1.GroupPolicyAssignment), nil
}