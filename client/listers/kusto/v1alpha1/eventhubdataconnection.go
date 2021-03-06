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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/kusto/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EventhubDataConnectionLister helps list EventhubDataConnections.
// All objects returned here must be treated as read-only.
type EventhubDataConnectionLister interface {
	// List lists all EventhubDataConnections in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EventhubDataConnection, err error)
	// EventhubDataConnections returns an object that can list and get EventhubDataConnections.
	EventhubDataConnections(namespace string) EventhubDataConnectionNamespaceLister
	EventhubDataConnectionListerExpansion
}

// eventhubDataConnectionLister implements the EventhubDataConnectionLister interface.
type eventhubDataConnectionLister struct {
	indexer cache.Indexer
}

// NewEventhubDataConnectionLister returns a new EventhubDataConnectionLister.
func NewEventhubDataConnectionLister(indexer cache.Indexer) EventhubDataConnectionLister {
	return &eventhubDataConnectionLister{indexer: indexer}
}

// List lists all EventhubDataConnections in the indexer.
func (s *eventhubDataConnectionLister) List(selector labels.Selector) (ret []*v1alpha1.EventhubDataConnection, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EventhubDataConnection))
	})
	return ret, err
}

// EventhubDataConnections returns an object that can list and get EventhubDataConnections.
func (s *eventhubDataConnectionLister) EventhubDataConnections(namespace string) EventhubDataConnectionNamespaceLister {
	return eventhubDataConnectionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EventhubDataConnectionNamespaceLister helps list and get EventhubDataConnections.
// All objects returned here must be treated as read-only.
type EventhubDataConnectionNamespaceLister interface {
	// List lists all EventhubDataConnections in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EventhubDataConnection, err error)
	// Get retrieves the EventhubDataConnection from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EventhubDataConnection, error)
	EventhubDataConnectionNamespaceListerExpansion
}

// eventhubDataConnectionNamespaceLister implements the EventhubDataConnectionNamespaceLister
// interface.
type eventhubDataConnectionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EventhubDataConnections in the indexer for a given namespace.
func (s eventhubDataConnectionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EventhubDataConnection, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EventhubDataConnection))
	})
	return ret, err
}

// Get retrieves the EventhubDataConnection from the indexer for a given namespace and name.
func (s eventhubDataConnectionNamespaceLister) Get(name string) (*v1alpha1.EventhubDataConnection, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("eventhubdataconnection"), name)
	}
	return obj.(*v1alpha1.EventhubDataConnection), nil
}
