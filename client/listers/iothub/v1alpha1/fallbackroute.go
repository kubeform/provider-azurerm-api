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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/iothub/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FallbackRouteLister helps list FallbackRoutes.
// All objects returned here must be treated as read-only.
type FallbackRouteLister interface {
	// List lists all FallbackRoutes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FallbackRoute, err error)
	// FallbackRoutes returns an object that can list and get FallbackRoutes.
	FallbackRoutes(namespace string) FallbackRouteNamespaceLister
	FallbackRouteListerExpansion
}

// fallbackRouteLister implements the FallbackRouteLister interface.
type fallbackRouteLister struct {
	indexer cache.Indexer
}

// NewFallbackRouteLister returns a new FallbackRouteLister.
func NewFallbackRouteLister(indexer cache.Indexer) FallbackRouteLister {
	return &fallbackRouteLister{indexer: indexer}
}

// List lists all FallbackRoutes in the indexer.
func (s *fallbackRouteLister) List(selector labels.Selector) (ret []*v1alpha1.FallbackRoute, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FallbackRoute))
	})
	return ret, err
}

// FallbackRoutes returns an object that can list and get FallbackRoutes.
func (s *fallbackRouteLister) FallbackRoutes(namespace string) FallbackRouteNamespaceLister {
	return fallbackRouteNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FallbackRouteNamespaceLister helps list and get FallbackRoutes.
// All objects returned here must be treated as read-only.
type FallbackRouteNamespaceLister interface {
	// List lists all FallbackRoutes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FallbackRoute, err error)
	// Get retrieves the FallbackRoute from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FallbackRoute, error)
	FallbackRouteNamespaceListerExpansion
}

// fallbackRouteNamespaceLister implements the FallbackRouteNamespaceLister
// interface.
type fallbackRouteNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FallbackRoutes in the indexer for a given namespace.
func (s fallbackRouteNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FallbackRoute, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FallbackRoute))
	})
	return ret, err
}

// Get retrieves the FallbackRoute from the indexer for a given namespace and name.
func (s fallbackRouteNamespaceLister) Get(name string) (*v1alpha1.FallbackRoute, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("fallbackroute"), name)
	}
	return obj.(*v1alpha1.FallbackRoute), nil
}