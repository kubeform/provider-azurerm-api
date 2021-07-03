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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/user/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AssignedIdentityLister helps list AssignedIdentities.
// All objects returned here must be treated as read-only.
type AssignedIdentityLister interface {
	// List lists all AssignedIdentities in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AssignedIdentity, err error)
	// AssignedIdentities returns an object that can list and get AssignedIdentities.
	AssignedIdentities(namespace string) AssignedIdentityNamespaceLister
	AssignedIdentityListerExpansion
}

// assignedIdentityLister implements the AssignedIdentityLister interface.
type assignedIdentityLister struct {
	indexer cache.Indexer
}

// NewAssignedIdentityLister returns a new AssignedIdentityLister.
func NewAssignedIdentityLister(indexer cache.Indexer) AssignedIdentityLister {
	return &assignedIdentityLister{indexer: indexer}
}

// List lists all AssignedIdentities in the indexer.
func (s *assignedIdentityLister) List(selector labels.Selector) (ret []*v1alpha1.AssignedIdentity, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AssignedIdentity))
	})
	return ret, err
}

// AssignedIdentities returns an object that can list and get AssignedIdentities.
func (s *assignedIdentityLister) AssignedIdentities(namespace string) AssignedIdentityNamespaceLister {
	return assignedIdentityNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AssignedIdentityNamespaceLister helps list and get AssignedIdentities.
// All objects returned here must be treated as read-only.
type AssignedIdentityNamespaceLister interface {
	// List lists all AssignedIdentities in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AssignedIdentity, err error)
	// Get retrieves the AssignedIdentity from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AssignedIdentity, error)
	AssignedIdentityNamespaceListerExpansion
}

// assignedIdentityNamespaceLister implements the AssignedIdentityNamespaceLister
// interface.
type assignedIdentityNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AssignedIdentities in the indexer for a given namespace.
func (s assignedIdentityNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AssignedIdentity, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AssignedIdentity))
	})
	return ret, err
}

// Get retrieves the AssignedIdentity from the indexer for a given namespace and name.
func (s assignedIdentityNamespaceLister) Get(name string) (*v1alpha1.AssignedIdentity, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("assignedidentity"), name)
	}
	return obj.(*v1alpha1.AssignedIdentity), nil
}
