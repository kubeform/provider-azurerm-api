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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/kusto/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterPrincipalAssignmentLister helps list ClusterPrincipalAssignments.
// All objects returned here must be treated as read-only.
type ClusterPrincipalAssignmentLister interface {
	// List lists all ClusterPrincipalAssignments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterPrincipalAssignment, err error)
	// ClusterPrincipalAssignments returns an object that can list and get ClusterPrincipalAssignments.
	ClusterPrincipalAssignments(namespace string) ClusterPrincipalAssignmentNamespaceLister
	ClusterPrincipalAssignmentListerExpansion
}

// clusterPrincipalAssignmentLister implements the ClusterPrincipalAssignmentLister interface.
type clusterPrincipalAssignmentLister struct {
	indexer cache.Indexer
}

// NewClusterPrincipalAssignmentLister returns a new ClusterPrincipalAssignmentLister.
func NewClusterPrincipalAssignmentLister(indexer cache.Indexer) ClusterPrincipalAssignmentLister {
	return &clusterPrincipalAssignmentLister{indexer: indexer}
}

// List lists all ClusterPrincipalAssignments in the indexer.
func (s *clusterPrincipalAssignmentLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterPrincipalAssignment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterPrincipalAssignment))
	})
	return ret, err
}

// ClusterPrincipalAssignments returns an object that can list and get ClusterPrincipalAssignments.
func (s *clusterPrincipalAssignmentLister) ClusterPrincipalAssignments(namespace string) ClusterPrincipalAssignmentNamespaceLister {
	return clusterPrincipalAssignmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterPrincipalAssignmentNamespaceLister helps list and get ClusterPrincipalAssignments.
// All objects returned here must be treated as read-only.
type ClusterPrincipalAssignmentNamespaceLister interface {
	// List lists all ClusterPrincipalAssignments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterPrincipalAssignment, err error)
	// Get retrieves the ClusterPrincipalAssignment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterPrincipalAssignment, error)
	ClusterPrincipalAssignmentNamespaceListerExpansion
}

// clusterPrincipalAssignmentNamespaceLister implements the ClusterPrincipalAssignmentNamespaceLister
// interface.
type clusterPrincipalAssignmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterPrincipalAssignments in the indexer for a given namespace.
func (s clusterPrincipalAssignmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterPrincipalAssignment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterPrincipalAssignment))
	})
	return ret, err
}

// Get retrieves the ClusterPrincipalAssignment from the indexer for a given namespace and name.
func (s clusterPrincipalAssignmentNamespaceLister) Get(name string) (*v1alpha1.ClusterPrincipalAssignment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusterprincipalassignment"), name)
	}
	return obj.(*v1alpha1.ClusterPrincipalAssignment), nil
}
