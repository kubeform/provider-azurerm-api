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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/synapse/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WorkspaceExtendedAuditingPolicyLister helps list WorkspaceExtendedAuditingPolicies.
// All objects returned here must be treated as read-only.
type WorkspaceExtendedAuditingPolicyLister interface {
	// List lists all WorkspaceExtendedAuditingPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WorkspaceExtendedAuditingPolicy, err error)
	// WorkspaceExtendedAuditingPolicies returns an object that can list and get WorkspaceExtendedAuditingPolicies.
	WorkspaceExtendedAuditingPolicies(namespace string) WorkspaceExtendedAuditingPolicyNamespaceLister
	WorkspaceExtendedAuditingPolicyListerExpansion
}

// workspaceExtendedAuditingPolicyLister implements the WorkspaceExtendedAuditingPolicyLister interface.
type workspaceExtendedAuditingPolicyLister struct {
	indexer cache.Indexer
}

// NewWorkspaceExtendedAuditingPolicyLister returns a new WorkspaceExtendedAuditingPolicyLister.
func NewWorkspaceExtendedAuditingPolicyLister(indexer cache.Indexer) WorkspaceExtendedAuditingPolicyLister {
	return &workspaceExtendedAuditingPolicyLister{indexer: indexer}
}

// List lists all WorkspaceExtendedAuditingPolicies in the indexer.
func (s *workspaceExtendedAuditingPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.WorkspaceExtendedAuditingPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WorkspaceExtendedAuditingPolicy))
	})
	return ret, err
}

// WorkspaceExtendedAuditingPolicies returns an object that can list and get WorkspaceExtendedAuditingPolicies.
func (s *workspaceExtendedAuditingPolicyLister) WorkspaceExtendedAuditingPolicies(namespace string) WorkspaceExtendedAuditingPolicyNamespaceLister {
	return workspaceExtendedAuditingPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WorkspaceExtendedAuditingPolicyNamespaceLister helps list and get WorkspaceExtendedAuditingPolicies.
// All objects returned here must be treated as read-only.
type WorkspaceExtendedAuditingPolicyNamespaceLister interface {
	// List lists all WorkspaceExtendedAuditingPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WorkspaceExtendedAuditingPolicy, err error)
	// Get retrieves the WorkspaceExtendedAuditingPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.WorkspaceExtendedAuditingPolicy, error)
	WorkspaceExtendedAuditingPolicyNamespaceListerExpansion
}

// workspaceExtendedAuditingPolicyNamespaceLister implements the WorkspaceExtendedAuditingPolicyNamespaceLister
// interface.
type workspaceExtendedAuditingPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WorkspaceExtendedAuditingPolicies in the indexer for a given namespace.
func (s workspaceExtendedAuditingPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WorkspaceExtendedAuditingPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WorkspaceExtendedAuditingPolicy))
	})
	return ret, err
}

// Get retrieves the WorkspaceExtendedAuditingPolicy from the indexer for a given namespace and name.
func (s workspaceExtendedAuditingPolicyNamespaceLister) Get(name string) (*v1alpha1.WorkspaceExtendedAuditingPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("workspaceextendedauditingpolicy"), name)
	}
	return obj.(*v1alpha1.WorkspaceExtendedAuditingPolicy), nil
}
