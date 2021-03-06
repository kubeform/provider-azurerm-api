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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/analysis/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServicesServerLister helps list ServicesServers.
// All objects returned here must be treated as read-only.
type ServicesServerLister interface {
	// List lists all ServicesServers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServicesServer, err error)
	// ServicesServers returns an object that can list and get ServicesServers.
	ServicesServers(namespace string) ServicesServerNamespaceLister
	ServicesServerListerExpansion
}

// servicesServerLister implements the ServicesServerLister interface.
type servicesServerLister struct {
	indexer cache.Indexer
}

// NewServicesServerLister returns a new ServicesServerLister.
func NewServicesServerLister(indexer cache.Indexer) ServicesServerLister {
	return &servicesServerLister{indexer: indexer}
}

// List lists all ServicesServers in the indexer.
func (s *servicesServerLister) List(selector labels.Selector) (ret []*v1alpha1.ServicesServer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServicesServer))
	})
	return ret, err
}

// ServicesServers returns an object that can list and get ServicesServers.
func (s *servicesServerLister) ServicesServers(namespace string) ServicesServerNamespaceLister {
	return servicesServerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServicesServerNamespaceLister helps list and get ServicesServers.
// All objects returned here must be treated as read-only.
type ServicesServerNamespaceLister interface {
	// List lists all ServicesServers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServicesServer, err error)
	// Get retrieves the ServicesServer from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ServicesServer, error)
	ServicesServerNamespaceListerExpansion
}

// servicesServerNamespaceLister implements the ServicesServerNamespaceLister
// interface.
type servicesServerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServicesServers in the indexer for a given namespace.
func (s servicesServerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServicesServer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServicesServer))
	})
	return ret, err
}

// Get retrieves the ServicesServer from the indexer for a given namespace and name.
func (s servicesServerNamespaceLister) Get(name string) (*v1alpha1.ServicesServer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("servicesserver"), name)
	}
	return obj.(*v1alpha1.ServicesServer), nil
}
