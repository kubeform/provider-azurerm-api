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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/virtual/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HubRouteTableLister helps list HubRouteTables.
// All objects returned here must be treated as read-only.
type HubRouteTableLister interface {
	// List lists all HubRouteTables in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HubRouteTable, err error)
	// HubRouteTables returns an object that can list and get HubRouteTables.
	HubRouteTables(namespace string) HubRouteTableNamespaceLister
	HubRouteTableListerExpansion
}

// hubRouteTableLister implements the HubRouteTableLister interface.
type hubRouteTableLister struct {
	indexer cache.Indexer
}

// NewHubRouteTableLister returns a new HubRouteTableLister.
func NewHubRouteTableLister(indexer cache.Indexer) HubRouteTableLister {
	return &hubRouteTableLister{indexer: indexer}
}

// List lists all HubRouteTables in the indexer.
func (s *hubRouteTableLister) List(selector labels.Selector) (ret []*v1alpha1.HubRouteTable, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HubRouteTable))
	})
	return ret, err
}

// HubRouteTables returns an object that can list and get HubRouteTables.
func (s *hubRouteTableLister) HubRouteTables(namespace string) HubRouteTableNamespaceLister {
	return hubRouteTableNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HubRouteTableNamespaceLister helps list and get HubRouteTables.
// All objects returned here must be treated as read-only.
type HubRouteTableNamespaceLister interface {
	// List lists all HubRouteTables in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HubRouteTable, err error)
	// Get retrieves the HubRouteTable from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.HubRouteTable, error)
	HubRouteTableNamespaceListerExpansion
}

// hubRouteTableNamespaceLister implements the HubRouteTableNamespaceLister
// interface.
type hubRouteTableNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HubRouteTables in the indexer for a given namespace.
func (s hubRouteTableNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HubRouteTable, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HubRouteTable))
	})
	return ret, err
}

// Get retrieves the HubRouteTable from the indexer for a given namespace and name.
func (s hubRouteTableNamespaceLister) Get(name string) (*v1alpha1.HubRouteTable, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("hubroutetable"), name)
	}
	return obj.(*v1alpha1.HubRouteTable), nil
}
