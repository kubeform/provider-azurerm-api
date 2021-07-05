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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/relay/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NamespaceLister helps list Namespaces.
// All objects returned here must be treated as read-only.
type NamespaceLister interface {
	// List lists all Namespaces in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Namespace, err error)
	// Namespaces returns an object that can list and get Namespaces.
	Namespaces(namespace string) NamespaceNamespaceLister
	NamespaceListerExpansion
}

// namespaceLister implements the NamespaceLister interface.
type namespaceLister struct {
	indexer cache.Indexer
}

// NewNamespaceLister returns a new NamespaceLister.
func NewNamespaceLister(indexer cache.Indexer) NamespaceLister {
	return &namespaceLister{indexer: indexer}
}

// List lists all Namespaces in the indexer.
func (s *namespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Namespace, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Namespace))
	})
	return ret, err
}

// Namespaces returns an object that can list and get Namespaces.
func (s *namespaceLister) Namespaces(namespace string) NamespaceNamespaceLister {
	return namespaceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NamespaceNamespaceLister helps list and get Namespaces.
// All objects returned here must be treated as read-only.
type NamespaceNamespaceLister interface {
	// List lists all Namespaces in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Namespace, err error)
	// Get retrieves the Namespace from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Namespace, error)
	NamespaceNamespaceListerExpansion
}

// namespaceNamespaceLister implements the NamespaceNamespaceLister
// interface.
type namespaceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Namespaces in the indexer for a given namespace.
func (s namespaceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Namespace, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Namespace))
	})
	return ret, err
}

// Get retrieves the Namespace from the indexer for a given namespace and name.
func (s namespaceNamespaceLister) Get(name string) (*v1alpha1.Namespace, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("namespace"), name)
	}
	return obj.(*v1alpha1.Namespace), nil
}
