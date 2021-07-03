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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/databox/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EdgeOrderLister helps list EdgeOrders.
// All objects returned here must be treated as read-only.
type EdgeOrderLister interface {
	// List lists all EdgeOrders in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EdgeOrder, err error)
	// EdgeOrders returns an object that can list and get EdgeOrders.
	EdgeOrders(namespace string) EdgeOrderNamespaceLister
	EdgeOrderListerExpansion
}

// edgeOrderLister implements the EdgeOrderLister interface.
type edgeOrderLister struct {
	indexer cache.Indexer
}

// NewEdgeOrderLister returns a new EdgeOrderLister.
func NewEdgeOrderLister(indexer cache.Indexer) EdgeOrderLister {
	return &edgeOrderLister{indexer: indexer}
}

// List lists all EdgeOrders in the indexer.
func (s *edgeOrderLister) List(selector labels.Selector) (ret []*v1alpha1.EdgeOrder, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EdgeOrder))
	})
	return ret, err
}

// EdgeOrders returns an object that can list and get EdgeOrders.
func (s *edgeOrderLister) EdgeOrders(namespace string) EdgeOrderNamespaceLister {
	return edgeOrderNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EdgeOrderNamespaceLister helps list and get EdgeOrders.
// All objects returned here must be treated as read-only.
type EdgeOrderNamespaceLister interface {
	// List lists all EdgeOrders in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EdgeOrder, err error)
	// Get retrieves the EdgeOrder from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EdgeOrder, error)
	EdgeOrderNamespaceListerExpansion
}

// edgeOrderNamespaceLister implements the EdgeOrderNamespaceLister
// interface.
type edgeOrderNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EdgeOrders in the indexer for a given namespace.
func (s edgeOrderNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EdgeOrder, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EdgeOrder))
	})
	return ret, err
}

// Get retrieves the EdgeOrder from the indexer for a given namespace and name.
func (s edgeOrderNamespaceLister) Get(name string) (*v1alpha1.EdgeOrder, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("edgeorder"), name)
	}
	return obj.(*v1alpha1.EdgeOrder), nil
}
