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

// DnsAaaaRecordLister helps list DnsAaaaRecords.
// All objects returned here must be treated as read-only.
type DnsAaaaRecordLister interface {
	// List lists all DnsAaaaRecords in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DnsAaaaRecord, err error)
	// DnsAaaaRecords returns an object that can list and get DnsAaaaRecords.
	DnsAaaaRecords(namespace string) DnsAaaaRecordNamespaceLister
	DnsAaaaRecordListerExpansion
}

// dnsAaaaRecordLister implements the DnsAaaaRecordLister interface.
type dnsAaaaRecordLister struct {
	indexer cache.Indexer
}

// NewDnsAaaaRecordLister returns a new DnsAaaaRecordLister.
func NewDnsAaaaRecordLister(indexer cache.Indexer) DnsAaaaRecordLister {
	return &dnsAaaaRecordLister{indexer: indexer}
}

// List lists all DnsAaaaRecords in the indexer.
func (s *dnsAaaaRecordLister) List(selector labels.Selector) (ret []*v1alpha1.DnsAaaaRecord, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DnsAaaaRecord))
	})
	return ret, err
}

// DnsAaaaRecords returns an object that can list and get DnsAaaaRecords.
func (s *dnsAaaaRecordLister) DnsAaaaRecords(namespace string) DnsAaaaRecordNamespaceLister {
	return dnsAaaaRecordNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DnsAaaaRecordNamespaceLister helps list and get DnsAaaaRecords.
// All objects returned here must be treated as read-only.
type DnsAaaaRecordNamespaceLister interface {
	// List lists all DnsAaaaRecords in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DnsAaaaRecord, err error)
	// Get retrieves the DnsAaaaRecord from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DnsAaaaRecord, error)
	DnsAaaaRecordNamespaceListerExpansion
}

// dnsAaaaRecordNamespaceLister implements the DnsAaaaRecordNamespaceLister
// interface.
type dnsAaaaRecordNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DnsAaaaRecords in the indexer for a given namespace.
func (s dnsAaaaRecordNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DnsAaaaRecord, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DnsAaaaRecord))
	})
	return ret, err
}

// Get retrieves the DnsAaaaRecord from the indexer for a given namespace and name.
func (s dnsAaaaRecordNamespaceLister) Get(name string) (*v1alpha1.DnsAaaaRecord, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dnsaaaarecord"), name)
	}
	return obj.(*v1alpha1.DnsAaaaRecord), nil
}
