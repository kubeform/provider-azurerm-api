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

// SeriesInsightsEventSourceIothubLister helps list SeriesInsightsEventSourceIothubs.
// All objects returned here must be treated as read-only.
type SeriesInsightsEventSourceIothubLister interface {
	// List lists all SeriesInsightsEventSourceIothubs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SeriesInsightsEventSourceIothub, err error)
	// SeriesInsightsEventSourceIothubs returns an object that can list and get SeriesInsightsEventSourceIothubs.
	SeriesInsightsEventSourceIothubs(namespace string) SeriesInsightsEventSourceIothubNamespaceLister
	SeriesInsightsEventSourceIothubListerExpansion
}

// seriesInsightsEventSourceIothubLister implements the SeriesInsightsEventSourceIothubLister interface.
type seriesInsightsEventSourceIothubLister struct {
	indexer cache.Indexer
}

// NewSeriesInsightsEventSourceIothubLister returns a new SeriesInsightsEventSourceIothubLister.
func NewSeriesInsightsEventSourceIothubLister(indexer cache.Indexer) SeriesInsightsEventSourceIothubLister {
	return &seriesInsightsEventSourceIothubLister{indexer: indexer}
}

// List lists all SeriesInsightsEventSourceIothubs in the indexer.
func (s *seriesInsightsEventSourceIothubLister) List(selector labels.Selector) (ret []*v1alpha1.SeriesInsightsEventSourceIothub, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SeriesInsightsEventSourceIothub))
	})
	return ret, err
}

// SeriesInsightsEventSourceIothubs returns an object that can list and get SeriesInsightsEventSourceIothubs.
func (s *seriesInsightsEventSourceIothubLister) SeriesInsightsEventSourceIothubs(namespace string) SeriesInsightsEventSourceIothubNamespaceLister {
	return seriesInsightsEventSourceIothubNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SeriesInsightsEventSourceIothubNamespaceLister helps list and get SeriesInsightsEventSourceIothubs.
// All objects returned here must be treated as read-only.
type SeriesInsightsEventSourceIothubNamespaceLister interface {
	// List lists all SeriesInsightsEventSourceIothubs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SeriesInsightsEventSourceIothub, err error)
	// Get retrieves the SeriesInsightsEventSourceIothub from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SeriesInsightsEventSourceIothub, error)
	SeriesInsightsEventSourceIothubNamespaceListerExpansion
}

// seriesInsightsEventSourceIothubNamespaceLister implements the SeriesInsightsEventSourceIothubNamespaceLister
// interface.
type seriesInsightsEventSourceIothubNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SeriesInsightsEventSourceIothubs in the indexer for a given namespace.
func (s seriesInsightsEventSourceIothubNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SeriesInsightsEventSourceIothub, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SeriesInsightsEventSourceIothub))
	})
	return ret, err
}

// Get retrieves the SeriesInsightsEventSourceIothub from the indexer for a given namespace and name.
func (s seriesInsightsEventSourceIothubNamespaceLister) Get(name string) (*v1alpha1.SeriesInsightsEventSourceIothub, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("seriesinsightseventsourceiothub"), name)
	}
	return obj.(*v1alpha1.SeriesInsightsEventSourceIothub), nil
}
