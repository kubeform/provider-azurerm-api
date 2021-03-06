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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/stream/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AnalyticsOutputFunctionLister helps list AnalyticsOutputFunctions.
// All objects returned here must be treated as read-only.
type AnalyticsOutputFunctionLister interface {
	// List lists all AnalyticsOutputFunctions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AnalyticsOutputFunction, err error)
	// AnalyticsOutputFunctions returns an object that can list and get AnalyticsOutputFunctions.
	AnalyticsOutputFunctions(namespace string) AnalyticsOutputFunctionNamespaceLister
	AnalyticsOutputFunctionListerExpansion
}

// analyticsOutputFunctionLister implements the AnalyticsOutputFunctionLister interface.
type analyticsOutputFunctionLister struct {
	indexer cache.Indexer
}

// NewAnalyticsOutputFunctionLister returns a new AnalyticsOutputFunctionLister.
func NewAnalyticsOutputFunctionLister(indexer cache.Indexer) AnalyticsOutputFunctionLister {
	return &analyticsOutputFunctionLister{indexer: indexer}
}

// List lists all AnalyticsOutputFunctions in the indexer.
func (s *analyticsOutputFunctionLister) List(selector labels.Selector) (ret []*v1alpha1.AnalyticsOutputFunction, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AnalyticsOutputFunction))
	})
	return ret, err
}

// AnalyticsOutputFunctions returns an object that can list and get AnalyticsOutputFunctions.
func (s *analyticsOutputFunctionLister) AnalyticsOutputFunctions(namespace string) AnalyticsOutputFunctionNamespaceLister {
	return analyticsOutputFunctionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AnalyticsOutputFunctionNamespaceLister helps list and get AnalyticsOutputFunctions.
// All objects returned here must be treated as read-only.
type AnalyticsOutputFunctionNamespaceLister interface {
	// List lists all AnalyticsOutputFunctions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AnalyticsOutputFunction, err error)
	// Get retrieves the AnalyticsOutputFunction from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AnalyticsOutputFunction, error)
	AnalyticsOutputFunctionNamespaceListerExpansion
}

// analyticsOutputFunctionNamespaceLister implements the AnalyticsOutputFunctionNamespaceLister
// interface.
type analyticsOutputFunctionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AnalyticsOutputFunctions in the indexer for a given namespace.
func (s analyticsOutputFunctionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AnalyticsOutputFunction, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AnalyticsOutputFunction))
	})
	return ret, err
}

// Get retrieves the AnalyticsOutputFunction from the indexer for a given namespace and name.
func (s analyticsOutputFunctionNamespaceLister) Get(name string) (*v1alpha1.AnalyticsOutputFunction, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("analyticsoutputfunction"), name)
	}
	return obj.(*v1alpha1.AnalyticsOutputFunction), nil
}
