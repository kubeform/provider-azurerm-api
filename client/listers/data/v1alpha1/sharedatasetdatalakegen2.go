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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/data/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ShareDatasetDataLakeGen2Lister helps list ShareDatasetDataLakeGen2s.
// All objects returned here must be treated as read-only.
type ShareDatasetDataLakeGen2Lister interface {
	// List lists all ShareDatasetDataLakeGen2s in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ShareDatasetDataLakeGen2, err error)
	// ShareDatasetDataLakeGen2s returns an object that can list and get ShareDatasetDataLakeGen2s.
	ShareDatasetDataLakeGen2s(namespace string) ShareDatasetDataLakeGen2NamespaceLister
	ShareDatasetDataLakeGen2ListerExpansion
}

// shareDatasetDataLakeGen2Lister implements the ShareDatasetDataLakeGen2Lister interface.
type shareDatasetDataLakeGen2Lister struct {
	indexer cache.Indexer
}

// NewShareDatasetDataLakeGen2Lister returns a new ShareDatasetDataLakeGen2Lister.
func NewShareDatasetDataLakeGen2Lister(indexer cache.Indexer) ShareDatasetDataLakeGen2Lister {
	return &shareDatasetDataLakeGen2Lister{indexer: indexer}
}

// List lists all ShareDatasetDataLakeGen2s in the indexer.
func (s *shareDatasetDataLakeGen2Lister) List(selector labels.Selector) (ret []*v1alpha1.ShareDatasetDataLakeGen2, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ShareDatasetDataLakeGen2))
	})
	return ret, err
}

// ShareDatasetDataLakeGen2s returns an object that can list and get ShareDatasetDataLakeGen2s.
func (s *shareDatasetDataLakeGen2Lister) ShareDatasetDataLakeGen2s(namespace string) ShareDatasetDataLakeGen2NamespaceLister {
	return shareDatasetDataLakeGen2NamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ShareDatasetDataLakeGen2NamespaceLister helps list and get ShareDatasetDataLakeGen2s.
// All objects returned here must be treated as read-only.
type ShareDatasetDataLakeGen2NamespaceLister interface {
	// List lists all ShareDatasetDataLakeGen2s in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ShareDatasetDataLakeGen2, err error)
	// Get retrieves the ShareDatasetDataLakeGen2 from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ShareDatasetDataLakeGen2, error)
	ShareDatasetDataLakeGen2NamespaceListerExpansion
}

// shareDatasetDataLakeGen2NamespaceLister implements the ShareDatasetDataLakeGen2NamespaceLister
// interface.
type shareDatasetDataLakeGen2NamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ShareDatasetDataLakeGen2s in the indexer for a given namespace.
func (s shareDatasetDataLakeGen2NamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ShareDatasetDataLakeGen2, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ShareDatasetDataLakeGen2))
	})
	return ret, err
}

// Get retrieves the ShareDatasetDataLakeGen2 from the indexer for a given namespace and name.
func (s shareDatasetDataLakeGen2NamespaceLister) Get(name string) (*v1alpha1.ShareDatasetDataLakeGen2, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sharedatasetdatalakegen2"), name)
	}
	return obj.(*v1alpha1.ShareDatasetDataLakeGen2), nil
}