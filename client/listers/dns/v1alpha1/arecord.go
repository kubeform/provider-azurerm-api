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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/dns/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ARecordLister helps list ARecords.
// All objects returned here must be treated as read-only.
type ARecordLister interface {
	// List lists all ARecords in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ARecord, err error)
	// ARecords returns an object that can list and get ARecords.
	ARecords(namespace string) ARecordNamespaceLister
	ARecordListerExpansion
}

// aRecordLister implements the ARecordLister interface.
type aRecordLister struct {
	indexer cache.Indexer
}

// NewARecordLister returns a new ARecordLister.
func NewARecordLister(indexer cache.Indexer) ARecordLister {
	return &aRecordLister{indexer: indexer}
}

// List lists all ARecords in the indexer.
func (s *aRecordLister) List(selector labels.Selector) (ret []*v1alpha1.ARecord, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ARecord))
	})
	return ret, err
}

// ARecords returns an object that can list and get ARecords.
func (s *aRecordLister) ARecords(namespace string) ARecordNamespaceLister {
	return aRecordNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ARecordNamespaceLister helps list and get ARecords.
// All objects returned here must be treated as read-only.
type ARecordNamespaceLister interface {
	// List lists all ARecords in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ARecord, err error)
	// Get retrieves the ARecord from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ARecord, error)
	ARecordNamespaceListerExpansion
}

// aRecordNamespaceLister implements the ARecordNamespaceLister
// interface.
type aRecordNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ARecords in the indexer for a given namespace.
func (s aRecordNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ARecord, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ARecord))
	})
	return ret, err
}

// Get retrieves the ARecord from the indexer for a given namespace and name.
func (s aRecordNamespaceLister) Get(name string) (*v1alpha1.ARecord, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("arecord"), name)
	}
	return obj.(*v1alpha1.ARecord), nil
}
