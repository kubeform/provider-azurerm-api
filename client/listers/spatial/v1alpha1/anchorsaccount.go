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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/spatial/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AnchorsAccountLister helps list AnchorsAccounts.
// All objects returned here must be treated as read-only.
type AnchorsAccountLister interface {
	// List lists all AnchorsAccounts in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AnchorsAccount, err error)
	// AnchorsAccounts returns an object that can list and get AnchorsAccounts.
	AnchorsAccounts(namespace string) AnchorsAccountNamespaceLister
	AnchorsAccountListerExpansion
}

// anchorsAccountLister implements the AnchorsAccountLister interface.
type anchorsAccountLister struct {
	indexer cache.Indexer
}

// NewAnchorsAccountLister returns a new AnchorsAccountLister.
func NewAnchorsAccountLister(indexer cache.Indexer) AnchorsAccountLister {
	return &anchorsAccountLister{indexer: indexer}
}

// List lists all AnchorsAccounts in the indexer.
func (s *anchorsAccountLister) List(selector labels.Selector) (ret []*v1alpha1.AnchorsAccount, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AnchorsAccount))
	})
	return ret, err
}

// AnchorsAccounts returns an object that can list and get AnchorsAccounts.
func (s *anchorsAccountLister) AnchorsAccounts(namespace string) AnchorsAccountNamespaceLister {
	return anchorsAccountNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AnchorsAccountNamespaceLister helps list and get AnchorsAccounts.
// All objects returned here must be treated as read-only.
type AnchorsAccountNamespaceLister interface {
	// List lists all AnchorsAccounts in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AnchorsAccount, err error)
	// Get retrieves the AnchorsAccount from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AnchorsAccount, error)
	AnchorsAccountNamespaceListerExpansion
}

// anchorsAccountNamespaceLister implements the AnchorsAccountNamespaceLister
// interface.
type anchorsAccountNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AnchorsAccounts in the indexer for a given namespace.
func (s anchorsAccountNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AnchorsAccount, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AnchorsAccount))
	})
	return ret, err
}

// Get retrieves the AnchorsAccount from the indexer for a given namespace and name.
func (s anchorsAccountNamespaceLister) Get(name string) (*v1alpha1.AnchorsAccount, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("anchorsaccount"), name)
	}
	return obj.(*v1alpha1.AnchorsAccount), nil
}