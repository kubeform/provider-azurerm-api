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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/postgresql/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServerKeyLister helps list ServerKeys.
// All objects returned here must be treated as read-only.
type ServerKeyLister interface {
	// List lists all ServerKeys in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServerKey, err error)
	// ServerKeys returns an object that can list and get ServerKeys.
	ServerKeys(namespace string) ServerKeyNamespaceLister
	ServerKeyListerExpansion
}

// serverKeyLister implements the ServerKeyLister interface.
type serverKeyLister struct {
	indexer cache.Indexer
}

// NewServerKeyLister returns a new ServerKeyLister.
func NewServerKeyLister(indexer cache.Indexer) ServerKeyLister {
	return &serverKeyLister{indexer: indexer}
}

// List lists all ServerKeys in the indexer.
func (s *serverKeyLister) List(selector labels.Selector) (ret []*v1alpha1.ServerKey, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServerKey))
	})
	return ret, err
}

// ServerKeys returns an object that can list and get ServerKeys.
func (s *serverKeyLister) ServerKeys(namespace string) ServerKeyNamespaceLister {
	return serverKeyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServerKeyNamespaceLister helps list and get ServerKeys.
// All objects returned here must be treated as read-only.
type ServerKeyNamespaceLister interface {
	// List lists all ServerKeys in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServerKey, err error)
	// Get retrieves the ServerKey from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ServerKey, error)
	ServerKeyNamespaceListerExpansion
}

// serverKeyNamespaceLister implements the ServerKeyNamespaceLister
// interface.
type serverKeyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServerKeys in the indexer for a given namespace.
func (s serverKeyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServerKey, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServerKey))
	})
	return ret, err
}

// Get retrieves the ServerKey from the indexer for a given namespace and name.
func (s serverKeyNamespaceLister) Get(name string) (*v1alpha1.ServerKey, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("serverkey"), name)
	}
	return obj.(*v1alpha1.ServerKey), nil
}
