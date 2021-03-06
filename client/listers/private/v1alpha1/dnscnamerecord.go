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

// DnsCnameRecordLister helps list DnsCnameRecords.
// All objects returned here must be treated as read-only.
type DnsCnameRecordLister interface {
	// List lists all DnsCnameRecords in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DnsCnameRecord, err error)
	// DnsCnameRecords returns an object that can list and get DnsCnameRecords.
	DnsCnameRecords(namespace string) DnsCnameRecordNamespaceLister
	DnsCnameRecordListerExpansion
}

// dnsCnameRecordLister implements the DnsCnameRecordLister interface.
type dnsCnameRecordLister struct {
	indexer cache.Indexer
}

// NewDnsCnameRecordLister returns a new DnsCnameRecordLister.
func NewDnsCnameRecordLister(indexer cache.Indexer) DnsCnameRecordLister {
	return &dnsCnameRecordLister{indexer: indexer}
}

// List lists all DnsCnameRecords in the indexer.
func (s *dnsCnameRecordLister) List(selector labels.Selector) (ret []*v1alpha1.DnsCnameRecord, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DnsCnameRecord))
	})
	return ret, err
}

// DnsCnameRecords returns an object that can list and get DnsCnameRecords.
func (s *dnsCnameRecordLister) DnsCnameRecords(namespace string) DnsCnameRecordNamespaceLister {
	return dnsCnameRecordNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DnsCnameRecordNamespaceLister helps list and get DnsCnameRecords.
// All objects returned here must be treated as read-only.
type DnsCnameRecordNamespaceLister interface {
	// List lists all DnsCnameRecords in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DnsCnameRecord, err error)
	// Get retrieves the DnsCnameRecord from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DnsCnameRecord, error)
	DnsCnameRecordNamespaceListerExpansion
}

// dnsCnameRecordNamespaceLister implements the DnsCnameRecordNamespaceLister
// interface.
type dnsCnameRecordNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DnsCnameRecords in the indexer for a given namespace.
func (s dnsCnameRecordNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DnsCnameRecord, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DnsCnameRecord))
	})
	return ret, err
}

// Get retrieves the DnsCnameRecord from the indexer for a given namespace and name.
func (s dnsCnameRecordNamespaceLister) Get(name string) (*v1alpha1.DnsCnameRecord, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dnscnamerecord"), name)
	}
	return obj.(*v1alpha1.DnsCnameRecord), nil
}
