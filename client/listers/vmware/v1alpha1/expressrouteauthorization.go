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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/vmware/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ExpressRouteAuthorizationLister helps list ExpressRouteAuthorizations.
// All objects returned here must be treated as read-only.
type ExpressRouteAuthorizationLister interface {
	// List lists all ExpressRouteAuthorizations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ExpressRouteAuthorization, err error)
	// ExpressRouteAuthorizations returns an object that can list and get ExpressRouteAuthorizations.
	ExpressRouteAuthorizations(namespace string) ExpressRouteAuthorizationNamespaceLister
	ExpressRouteAuthorizationListerExpansion
}

// expressRouteAuthorizationLister implements the ExpressRouteAuthorizationLister interface.
type expressRouteAuthorizationLister struct {
	indexer cache.Indexer
}

// NewExpressRouteAuthorizationLister returns a new ExpressRouteAuthorizationLister.
func NewExpressRouteAuthorizationLister(indexer cache.Indexer) ExpressRouteAuthorizationLister {
	return &expressRouteAuthorizationLister{indexer: indexer}
}

// List lists all ExpressRouteAuthorizations in the indexer.
func (s *expressRouteAuthorizationLister) List(selector labels.Selector) (ret []*v1alpha1.ExpressRouteAuthorization, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ExpressRouteAuthorization))
	})
	return ret, err
}

// ExpressRouteAuthorizations returns an object that can list and get ExpressRouteAuthorizations.
func (s *expressRouteAuthorizationLister) ExpressRouteAuthorizations(namespace string) ExpressRouteAuthorizationNamespaceLister {
	return expressRouteAuthorizationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ExpressRouteAuthorizationNamespaceLister helps list and get ExpressRouteAuthorizations.
// All objects returned here must be treated as read-only.
type ExpressRouteAuthorizationNamespaceLister interface {
	// List lists all ExpressRouteAuthorizations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ExpressRouteAuthorization, err error)
	// Get retrieves the ExpressRouteAuthorization from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ExpressRouteAuthorization, error)
	ExpressRouteAuthorizationNamespaceListerExpansion
}

// expressRouteAuthorizationNamespaceLister implements the ExpressRouteAuthorizationNamespaceLister
// interface.
type expressRouteAuthorizationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ExpressRouteAuthorizations in the indexer for a given namespace.
func (s expressRouteAuthorizationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ExpressRouteAuthorization, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ExpressRouteAuthorization))
	})
	return ret, err
}

// Get retrieves the ExpressRouteAuthorization from the indexer for a given namespace and name.
func (s expressRouteAuthorizationNamespaceLister) Get(name string) (*v1alpha1.ExpressRouteAuthorization, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("expressrouteauthorization"), name)
	}
	return obj.(*v1alpha1.ExpressRouteAuthorization), nil
}
