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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/private/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DnsZoneLister helps list DnsZones.
// All objects returned here must be treated as read-only.
type DnsZoneLister interface {
	// List lists all DnsZones in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DnsZone, err error)
	// DnsZones returns an object that can list and get DnsZones.
	DnsZones(namespace string) DnsZoneNamespaceLister
	DnsZoneListerExpansion
}

// dnsZoneLister implements the DnsZoneLister interface.
type dnsZoneLister struct {
	indexer cache.Indexer
}

// NewDnsZoneLister returns a new DnsZoneLister.
func NewDnsZoneLister(indexer cache.Indexer) DnsZoneLister {
	return &dnsZoneLister{indexer: indexer}
}

// List lists all DnsZones in the indexer.
func (s *dnsZoneLister) List(selector labels.Selector) (ret []*v1alpha1.DnsZone, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DnsZone))
	})
	return ret, err
}

// DnsZones returns an object that can list and get DnsZones.
func (s *dnsZoneLister) DnsZones(namespace string) DnsZoneNamespaceLister {
	return dnsZoneNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DnsZoneNamespaceLister helps list and get DnsZones.
// All objects returned here must be treated as read-only.
type DnsZoneNamespaceLister interface {
	// List lists all DnsZones in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DnsZone, err error)
	// Get retrieves the DnsZone from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DnsZone, error)
	DnsZoneNamespaceListerExpansion
}

// dnsZoneNamespaceLister implements the DnsZoneNamespaceLister
// interface.
type dnsZoneNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DnsZones in the indexer for a given namespace.
func (s dnsZoneNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DnsZone, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DnsZone))
	})
	return ret, err
}

// Get retrieves the DnsZone from the indexer for a given namespace and name.
func (s dnsZoneNamespaceLister) Get(name string) (*v1alpha1.DnsZone, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dnszone"), name)
	}
	return obj.(*v1alpha1.DnsZone), nil
}