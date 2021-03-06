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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/storage/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ManagementPolicyLister helps list ManagementPolicies.
// All objects returned here must be treated as read-only.
type ManagementPolicyLister interface {
	// List lists all ManagementPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagementPolicy, err error)
	// ManagementPolicies returns an object that can list and get ManagementPolicies.
	ManagementPolicies(namespace string) ManagementPolicyNamespaceLister
	ManagementPolicyListerExpansion
}

// managementPolicyLister implements the ManagementPolicyLister interface.
type managementPolicyLister struct {
	indexer cache.Indexer
}

// NewManagementPolicyLister returns a new ManagementPolicyLister.
func NewManagementPolicyLister(indexer cache.Indexer) ManagementPolicyLister {
	return &managementPolicyLister{indexer: indexer}
}

// List lists all ManagementPolicies in the indexer.
func (s *managementPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.ManagementPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagementPolicy))
	})
	return ret, err
}

// ManagementPolicies returns an object that can list and get ManagementPolicies.
func (s *managementPolicyLister) ManagementPolicies(namespace string) ManagementPolicyNamespaceLister {
	return managementPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ManagementPolicyNamespaceLister helps list and get ManagementPolicies.
// All objects returned here must be treated as read-only.
type ManagementPolicyNamespaceLister interface {
	// List lists all ManagementPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagementPolicy, err error)
	// Get retrieves the ManagementPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ManagementPolicy, error)
	ManagementPolicyNamespaceListerExpansion
}

// managementPolicyNamespaceLister implements the ManagementPolicyNamespaceLister
// interface.
type managementPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ManagementPolicies in the indexer for a given namespace.
func (s managementPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ManagementPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagementPolicy))
	})
	return ret, err
}

// Get retrieves the ManagementPolicy from the indexer for a given namespace and name.
func (s managementPolicyNamespaceLister) Get(name string) (*v1alpha1.ManagementPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("managementpolicy"), name)
	}
	return obj.(*v1alpha1.ManagementPolicy), nil
}
