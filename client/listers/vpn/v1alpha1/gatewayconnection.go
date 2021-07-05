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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/vpn/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GatewayConnectionLister helps list GatewayConnections.
// All objects returned here must be treated as read-only.
type GatewayConnectionLister interface {
	// List lists all GatewayConnections in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GatewayConnection, err error)
	// GatewayConnections returns an object that can list and get GatewayConnections.
	GatewayConnections(namespace string) GatewayConnectionNamespaceLister
	GatewayConnectionListerExpansion
}

// gatewayConnectionLister implements the GatewayConnectionLister interface.
type gatewayConnectionLister struct {
	indexer cache.Indexer
}

// NewGatewayConnectionLister returns a new GatewayConnectionLister.
func NewGatewayConnectionLister(indexer cache.Indexer) GatewayConnectionLister {
	return &gatewayConnectionLister{indexer: indexer}
}

// List lists all GatewayConnections in the indexer.
func (s *gatewayConnectionLister) List(selector labels.Selector) (ret []*v1alpha1.GatewayConnection, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GatewayConnection))
	})
	return ret, err
}

// GatewayConnections returns an object that can list and get GatewayConnections.
func (s *gatewayConnectionLister) GatewayConnections(namespace string) GatewayConnectionNamespaceLister {
	return gatewayConnectionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GatewayConnectionNamespaceLister helps list and get GatewayConnections.
// All objects returned here must be treated as read-only.
type GatewayConnectionNamespaceLister interface {
	// List lists all GatewayConnections in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GatewayConnection, err error)
	// Get retrieves the GatewayConnection from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.GatewayConnection, error)
	GatewayConnectionNamespaceListerExpansion
}

// gatewayConnectionNamespaceLister implements the GatewayConnectionNamespaceLister
// interface.
type gatewayConnectionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GatewayConnections in the indexer for a given namespace.
func (s gatewayConnectionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GatewayConnection, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GatewayConnection))
	})
	return ret, err
}

// Get retrieves the GatewayConnection from the indexer for a given namespace and name.
func (s gatewayConnectionNamespaceLister) Get(name string) (*v1alpha1.GatewayConnection, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("gatewayconnection"), name)
	}
	return obj.(*v1alpha1.GatewayConnection), nil
}