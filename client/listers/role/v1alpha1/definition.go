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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/role/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DefinitionLister helps list Definitions.
// All objects returned here must be treated as read-only.
type DefinitionLister interface {
	// List lists all Definitions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Definition, err error)
	// Definitions returns an object that can list and get Definitions.
	Definitions(namespace string) DefinitionNamespaceLister
	DefinitionListerExpansion
}

// definitionLister implements the DefinitionLister interface.
type definitionLister struct {
	indexer cache.Indexer
}

// NewDefinitionLister returns a new DefinitionLister.
func NewDefinitionLister(indexer cache.Indexer) DefinitionLister {
	return &definitionLister{indexer: indexer}
}

// List lists all Definitions in the indexer.
func (s *definitionLister) List(selector labels.Selector) (ret []*v1alpha1.Definition, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Definition))
	})
	return ret, err
}

// Definitions returns an object that can list and get Definitions.
func (s *definitionLister) Definitions(namespace string) DefinitionNamespaceLister {
	return definitionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DefinitionNamespaceLister helps list and get Definitions.
// All objects returned here must be treated as read-only.
type DefinitionNamespaceLister interface {
	// List lists all Definitions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Definition, err error)
	// Get retrieves the Definition from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Definition, error)
	DefinitionNamespaceListerExpansion
}

// definitionNamespaceLister implements the DefinitionNamespaceLister
// interface.
type definitionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Definitions in the indexer for a given namespace.
func (s definitionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Definition, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Definition))
	})
	return ret, err
}

// Get retrieves the Definition from the indexer for a given namespace and name.
func (s definitionNamespaceLister) Get(name string) (*v1alpha1.Definition, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("definition"), name)
	}
	return obj.(*v1alpha1.Definition), nil
}
