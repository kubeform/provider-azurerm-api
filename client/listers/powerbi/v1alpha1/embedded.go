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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/powerbi/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EmbeddedLister helps list Embeddeds.
// All objects returned here must be treated as read-only.
type EmbeddedLister interface {
	// List lists all Embeddeds in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Embedded, err error)
	// Embeddeds returns an object that can list and get Embeddeds.
	Embeddeds(namespace string) EmbeddedNamespaceLister
	EmbeddedListerExpansion
}

// embeddedLister implements the EmbeddedLister interface.
type embeddedLister struct {
	indexer cache.Indexer
}

// NewEmbeddedLister returns a new EmbeddedLister.
func NewEmbeddedLister(indexer cache.Indexer) EmbeddedLister {
	return &embeddedLister{indexer: indexer}
}

// List lists all Embeddeds in the indexer.
func (s *embeddedLister) List(selector labels.Selector) (ret []*v1alpha1.Embedded, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Embedded))
	})
	return ret, err
}

// Embeddeds returns an object that can list and get Embeddeds.
func (s *embeddedLister) Embeddeds(namespace string) EmbeddedNamespaceLister {
	return embeddedNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EmbeddedNamespaceLister helps list and get Embeddeds.
// All objects returned here must be treated as read-only.
type EmbeddedNamespaceLister interface {
	// List lists all Embeddeds in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Embedded, err error)
	// Get retrieves the Embedded from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Embedded, error)
	EmbeddedNamespaceListerExpansion
}

// embeddedNamespaceLister implements the EmbeddedNamespaceLister
// interface.
type embeddedNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Embeddeds in the indexer for a given namespace.
func (s embeddedNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Embedded, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Embedded))
	})
	return ret, err
}

// Get retrieves the Embedded from the indexer for a given namespace and name.
func (s embeddedNamespaceLister) Get(name string) (*v1alpha1.Embedded, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("embedded"), name)
	}
	return obj.(*v1alpha1.Embedded), nil
}