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

// EndpointServicebusQueueLister helps list EndpointServicebusQueues.
// All objects returned here must be treated as read-only.
type EndpointServicebusQueueLister interface {
	// List lists all EndpointServicebusQueues in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EndpointServicebusQueue, err error)
	// EndpointServicebusQueues returns an object that can list and get EndpointServicebusQueues.
	EndpointServicebusQueues(namespace string) EndpointServicebusQueueNamespaceLister
	EndpointServicebusQueueListerExpansion
}

// endpointServicebusQueueLister implements the EndpointServicebusQueueLister interface.
type endpointServicebusQueueLister struct {
	indexer cache.Indexer
}

// NewEndpointServicebusQueueLister returns a new EndpointServicebusQueueLister.
func NewEndpointServicebusQueueLister(indexer cache.Indexer) EndpointServicebusQueueLister {
	return &endpointServicebusQueueLister{indexer: indexer}
}

// List lists all EndpointServicebusQueues in the indexer.
func (s *endpointServicebusQueueLister) List(selector labels.Selector) (ret []*v1alpha1.EndpointServicebusQueue, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EndpointServicebusQueue))
	})
	return ret, err
}

// EndpointServicebusQueues returns an object that can list and get EndpointServicebusQueues.
func (s *endpointServicebusQueueLister) EndpointServicebusQueues(namespace string) EndpointServicebusQueueNamespaceLister {
	return endpointServicebusQueueNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EndpointServicebusQueueNamespaceLister helps list and get EndpointServicebusQueues.
// All objects returned here must be treated as read-only.
type EndpointServicebusQueueNamespaceLister interface {
	// List lists all EndpointServicebusQueues in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EndpointServicebusQueue, err error)
	// Get retrieves the EndpointServicebusQueue from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EndpointServicebusQueue, error)
	EndpointServicebusQueueNamespaceListerExpansion
}

// endpointServicebusQueueNamespaceLister implements the EndpointServicebusQueueNamespaceLister
// interface.
type endpointServicebusQueueNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EndpointServicebusQueues in the indexer for a given namespace.
func (s endpointServicebusQueueNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EndpointServicebusQueue, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EndpointServicebusQueue))
	})
	return ret, err
}

// Get retrieves the EndpointServicebusQueue from the indexer for a given namespace and name.
func (s endpointServicebusQueueNamespaceLister) Get(name string) (*v1alpha1.EndpointServicebusQueue, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("endpointservicebusqueue"), name)
	}
	return obj.(*v1alpha1.EndpointServicebusQueue), nil
}
