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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/backup/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ProtectedFileShareLister helps list ProtectedFileShares.
// All objects returned here must be treated as read-only.
type ProtectedFileShareLister interface {
	// List lists all ProtectedFileShares in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProtectedFileShare, err error)
	// ProtectedFileShares returns an object that can list and get ProtectedFileShares.
	ProtectedFileShares(namespace string) ProtectedFileShareNamespaceLister
	ProtectedFileShareListerExpansion
}

// protectedFileShareLister implements the ProtectedFileShareLister interface.
type protectedFileShareLister struct {
	indexer cache.Indexer
}

// NewProtectedFileShareLister returns a new ProtectedFileShareLister.
func NewProtectedFileShareLister(indexer cache.Indexer) ProtectedFileShareLister {
	return &protectedFileShareLister{indexer: indexer}
}

// List lists all ProtectedFileShares in the indexer.
func (s *protectedFileShareLister) List(selector labels.Selector) (ret []*v1alpha1.ProtectedFileShare, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProtectedFileShare))
	})
	return ret, err
}

// ProtectedFileShares returns an object that can list and get ProtectedFileShares.
func (s *protectedFileShareLister) ProtectedFileShares(namespace string) ProtectedFileShareNamespaceLister {
	return protectedFileShareNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProtectedFileShareNamespaceLister helps list and get ProtectedFileShares.
// All objects returned here must be treated as read-only.
type ProtectedFileShareNamespaceLister interface {
	// List lists all ProtectedFileShares in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProtectedFileShare, err error)
	// Get retrieves the ProtectedFileShare from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ProtectedFileShare, error)
	ProtectedFileShareNamespaceListerExpansion
}

// protectedFileShareNamespaceLister implements the ProtectedFileShareNamespaceLister
// interface.
type protectedFileShareNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ProtectedFileShares in the indexer for a given namespace.
func (s protectedFileShareNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ProtectedFileShare, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProtectedFileShare))
	})
	return ret, err
}

// Get retrieves the ProtectedFileShare from the indexer for a given namespace and name.
func (s protectedFileShareNamespaceLister) Get(name string) (*v1alpha1.ProtectedFileShare, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("protectedfileshare"), name)
	}
	return obj.(*v1alpha1.ProtectedFileShare), nil
}
