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

// DnsTxtRecordLister helps list DnsTxtRecords.
// All objects returned here must be treated as read-only.
type DnsTxtRecordLister interface {
	// List lists all DnsTxtRecords in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DnsTxtRecord, err error)
	// DnsTxtRecords returns an object that can list and get DnsTxtRecords.
	DnsTxtRecords(namespace string) DnsTxtRecordNamespaceLister
	DnsTxtRecordListerExpansion
}

// dnsTxtRecordLister implements the DnsTxtRecordLister interface.
type dnsTxtRecordLister struct {
	indexer cache.Indexer
}

// NewDnsTxtRecordLister returns a new DnsTxtRecordLister.
func NewDnsTxtRecordLister(indexer cache.Indexer) DnsTxtRecordLister {
	return &dnsTxtRecordLister{indexer: indexer}
}

// List lists all DnsTxtRecords in the indexer.
func (s *dnsTxtRecordLister) List(selector labels.Selector) (ret []*v1alpha1.DnsTxtRecord, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DnsTxtRecord))
	})
	return ret, err
}

// DnsTxtRecords returns an object that can list and get DnsTxtRecords.
func (s *dnsTxtRecordLister) DnsTxtRecords(namespace string) DnsTxtRecordNamespaceLister {
	return dnsTxtRecordNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DnsTxtRecordNamespaceLister helps list and get DnsTxtRecords.
// All objects returned here must be treated as read-only.
type DnsTxtRecordNamespaceLister interface {
	// List lists all DnsTxtRecords in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DnsTxtRecord, err error)
	// Get retrieves the DnsTxtRecord from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DnsTxtRecord, error)
	DnsTxtRecordNamespaceListerExpansion
}

// dnsTxtRecordNamespaceLister implements the DnsTxtRecordNamespaceLister
// interface.
type dnsTxtRecordNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DnsTxtRecords in the indexer for a given namespace.
func (s dnsTxtRecordNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DnsTxtRecord, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DnsTxtRecord))
	})
	return ret, err
}

// Get retrieves the DnsTxtRecord from the indexer for a given namespace and name.
func (s dnsTxtRecordNamespaceLister) Get(name string) (*v1alpha1.DnsTxtRecord, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dnstxtrecord"), name)
	}
	return obj.(*v1alpha1.DnsTxtRecord), nil
}
