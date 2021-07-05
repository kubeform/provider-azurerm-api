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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/availability/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SetLister helps list Sets.
// All objects returned here must be treated as read-only.
type SetLister interface {
	// List lists all Sets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Set, err error)
	// Sets returns an object that can list and get Sets.
	Sets(namespace string) SetNamespaceLister
	SetListerExpansion
}

// setLister implements the SetLister interface.
type setLister struct {
	indexer cache.Indexer
}

// NewSetLister returns a new SetLister.
func NewSetLister(indexer cache.Indexer) SetLister {
	return &setLister{indexer: indexer}
}

// List lists all Sets in the indexer.
func (s *setLister) List(selector labels.Selector) (ret []*v1alpha1.Set, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Set))
	})
	return ret, err
}

// Sets returns an object that can list and get Sets.
func (s *setLister) Sets(namespace string) SetNamespaceLister {
	return setNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SetNamespaceLister helps list and get Sets.
// All objects returned here must be treated as read-only.
type SetNamespaceLister interface {
	// List lists all Sets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Set, err error)
	// Get retrieves the Set from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Set, error)
	SetNamespaceListerExpansion
}

// setNamespaceLister implements the SetNamespaceLister
// interface.
type setNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Sets in the indexer for a given namespace.
func (s setNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Set, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Set))
	})
	return ret, err
}

// Get retrieves the Set from the indexer for a given namespace and name.
func (s setNamespaceLister) Get(name string) (*v1alpha1.Set, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("set"), name)
	}
	return obj.(*v1alpha1.Set), nil
}