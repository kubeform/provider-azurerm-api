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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/hpc/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CacheBlobTargetLister helps list CacheBlobTargets.
// All objects returned here must be treated as read-only.
type CacheBlobTargetLister interface {
	// List lists all CacheBlobTargets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CacheBlobTarget, err error)
	// CacheBlobTargets returns an object that can list and get CacheBlobTargets.
	CacheBlobTargets(namespace string) CacheBlobTargetNamespaceLister
	CacheBlobTargetListerExpansion
}

// cacheBlobTargetLister implements the CacheBlobTargetLister interface.
type cacheBlobTargetLister struct {
	indexer cache.Indexer
}

// NewCacheBlobTargetLister returns a new CacheBlobTargetLister.
func NewCacheBlobTargetLister(indexer cache.Indexer) CacheBlobTargetLister {
	return &cacheBlobTargetLister{indexer: indexer}
}

// List lists all CacheBlobTargets in the indexer.
func (s *cacheBlobTargetLister) List(selector labels.Selector) (ret []*v1alpha1.CacheBlobTarget, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CacheBlobTarget))
	})
	return ret, err
}

// CacheBlobTargets returns an object that can list and get CacheBlobTargets.
func (s *cacheBlobTargetLister) CacheBlobTargets(namespace string) CacheBlobTargetNamespaceLister {
	return cacheBlobTargetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CacheBlobTargetNamespaceLister helps list and get CacheBlobTargets.
// All objects returned here must be treated as read-only.
type CacheBlobTargetNamespaceLister interface {
	// List lists all CacheBlobTargets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CacheBlobTarget, err error)
	// Get retrieves the CacheBlobTarget from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CacheBlobTarget, error)
	CacheBlobTargetNamespaceListerExpansion
}

// cacheBlobTargetNamespaceLister implements the CacheBlobTargetNamespaceLister
// interface.
type cacheBlobTargetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CacheBlobTargets in the indexer for a given namespace.
func (s cacheBlobTargetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CacheBlobTarget, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CacheBlobTarget))
	})
	return ret, err
}

// Get retrieves the CacheBlobTarget from the indexer for a given namespace and name.
func (s cacheBlobTargetNamespaceLister) Get(name string) (*v1alpha1.CacheBlobTarget, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cacheblobtarget"), name)
	}
	return obj.(*v1alpha1.CacheBlobTarget), nil
}