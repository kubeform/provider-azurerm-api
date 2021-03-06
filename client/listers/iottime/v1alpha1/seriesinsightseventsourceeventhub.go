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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/iottime/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SeriesInsightsEventSourceEventhubLister helps list SeriesInsightsEventSourceEventhubs.
// All objects returned here must be treated as read-only.
type SeriesInsightsEventSourceEventhubLister interface {
	// List lists all SeriesInsightsEventSourceEventhubs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SeriesInsightsEventSourceEventhub, err error)
	// SeriesInsightsEventSourceEventhubs returns an object that can list and get SeriesInsightsEventSourceEventhubs.
	SeriesInsightsEventSourceEventhubs(namespace string) SeriesInsightsEventSourceEventhubNamespaceLister
	SeriesInsightsEventSourceEventhubListerExpansion
}

// seriesInsightsEventSourceEventhubLister implements the SeriesInsightsEventSourceEventhubLister interface.
type seriesInsightsEventSourceEventhubLister struct {
	indexer cache.Indexer
}

// NewSeriesInsightsEventSourceEventhubLister returns a new SeriesInsightsEventSourceEventhubLister.
func NewSeriesInsightsEventSourceEventhubLister(indexer cache.Indexer) SeriesInsightsEventSourceEventhubLister {
	return &seriesInsightsEventSourceEventhubLister{indexer: indexer}
}

// List lists all SeriesInsightsEventSourceEventhubs in the indexer.
func (s *seriesInsightsEventSourceEventhubLister) List(selector labels.Selector) (ret []*v1alpha1.SeriesInsightsEventSourceEventhub, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SeriesInsightsEventSourceEventhub))
	})
	return ret, err
}

// SeriesInsightsEventSourceEventhubs returns an object that can list and get SeriesInsightsEventSourceEventhubs.
func (s *seriesInsightsEventSourceEventhubLister) SeriesInsightsEventSourceEventhubs(namespace string) SeriesInsightsEventSourceEventhubNamespaceLister {
	return seriesInsightsEventSourceEventhubNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SeriesInsightsEventSourceEventhubNamespaceLister helps list and get SeriesInsightsEventSourceEventhubs.
// All objects returned here must be treated as read-only.
type SeriesInsightsEventSourceEventhubNamespaceLister interface {
	// List lists all SeriesInsightsEventSourceEventhubs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SeriesInsightsEventSourceEventhub, err error)
	// Get retrieves the SeriesInsightsEventSourceEventhub from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SeriesInsightsEventSourceEventhub, error)
	SeriesInsightsEventSourceEventhubNamespaceListerExpansion
}

// seriesInsightsEventSourceEventhubNamespaceLister implements the SeriesInsightsEventSourceEventhubNamespaceLister
// interface.
type seriesInsightsEventSourceEventhubNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SeriesInsightsEventSourceEventhubs in the indexer for a given namespace.
func (s seriesInsightsEventSourceEventhubNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SeriesInsightsEventSourceEventhub, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SeriesInsightsEventSourceEventhub))
	})
	return ret, err
}

// Get retrieves the SeriesInsightsEventSourceEventhub from the indexer for a given namespace and name.
func (s seriesInsightsEventSourceEventhubNamespaceLister) Get(name string) (*v1alpha1.SeriesInsightsEventSourceEventhub, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("seriesinsightseventsourceeventhub"), name)
	}
	return obj.(*v1alpha1.SeriesInsightsEventSourceEventhub), nil
}
