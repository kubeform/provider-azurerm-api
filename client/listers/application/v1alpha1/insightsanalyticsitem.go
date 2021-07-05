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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/application/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InsightsAnalyticsItemLister helps list InsightsAnalyticsItems.
// All objects returned here must be treated as read-only.
type InsightsAnalyticsItemLister interface {
	// List lists all InsightsAnalyticsItems in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InsightsAnalyticsItem, err error)
	// InsightsAnalyticsItems returns an object that can list and get InsightsAnalyticsItems.
	InsightsAnalyticsItems(namespace string) InsightsAnalyticsItemNamespaceLister
	InsightsAnalyticsItemListerExpansion
}

// insightsAnalyticsItemLister implements the InsightsAnalyticsItemLister interface.
type insightsAnalyticsItemLister struct {
	indexer cache.Indexer
}

// NewInsightsAnalyticsItemLister returns a new InsightsAnalyticsItemLister.
func NewInsightsAnalyticsItemLister(indexer cache.Indexer) InsightsAnalyticsItemLister {
	return &insightsAnalyticsItemLister{indexer: indexer}
}

// List lists all InsightsAnalyticsItems in the indexer.
func (s *insightsAnalyticsItemLister) List(selector labels.Selector) (ret []*v1alpha1.InsightsAnalyticsItem, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InsightsAnalyticsItem))
	})
	return ret, err
}

// InsightsAnalyticsItems returns an object that can list and get InsightsAnalyticsItems.
func (s *insightsAnalyticsItemLister) InsightsAnalyticsItems(namespace string) InsightsAnalyticsItemNamespaceLister {
	return insightsAnalyticsItemNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InsightsAnalyticsItemNamespaceLister helps list and get InsightsAnalyticsItems.
// All objects returned here must be treated as read-only.
type InsightsAnalyticsItemNamespaceLister interface {
	// List lists all InsightsAnalyticsItems in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InsightsAnalyticsItem, err error)
	// Get retrieves the InsightsAnalyticsItem from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.InsightsAnalyticsItem, error)
	InsightsAnalyticsItemNamespaceListerExpansion
}

// insightsAnalyticsItemNamespaceLister implements the InsightsAnalyticsItemNamespaceLister
// interface.
type insightsAnalyticsItemNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all InsightsAnalyticsItems in the indexer for a given namespace.
func (s insightsAnalyticsItemNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.InsightsAnalyticsItem, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InsightsAnalyticsItem))
	})
	return ret, err
}

// Get retrieves the InsightsAnalyticsItem from the indexer for a given namespace and name.
func (s insightsAnalyticsItemNamespaceLister) Get(name string) (*v1alpha1.InsightsAnalyticsItem, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("insightsanalyticsitem"), name)
	}
	return obj.(*v1alpha1.InsightsAnalyticsItem), nil
}