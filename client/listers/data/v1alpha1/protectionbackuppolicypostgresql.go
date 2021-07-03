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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/data/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ProtectionBackupPolicyPostgresqlLister helps list ProtectionBackupPolicyPostgresqls.
// All objects returned here must be treated as read-only.
type ProtectionBackupPolicyPostgresqlLister interface {
	// List lists all ProtectionBackupPolicyPostgresqls in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProtectionBackupPolicyPostgresql, err error)
	// ProtectionBackupPolicyPostgresqls returns an object that can list and get ProtectionBackupPolicyPostgresqls.
	ProtectionBackupPolicyPostgresqls(namespace string) ProtectionBackupPolicyPostgresqlNamespaceLister
	ProtectionBackupPolicyPostgresqlListerExpansion
}

// protectionBackupPolicyPostgresqlLister implements the ProtectionBackupPolicyPostgresqlLister interface.
type protectionBackupPolicyPostgresqlLister struct {
	indexer cache.Indexer
}

// NewProtectionBackupPolicyPostgresqlLister returns a new ProtectionBackupPolicyPostgresqlLister.
func NewProtectionBackupPolicyPostgresqlLister(indexer cache.Indexer) ProtectionBackupPolicyPostgresqlLister {
	return &protectionBackupPolicyPostgresqlLister{indexer: indexer}
}

// List lists all ProtectionBackupPolicyPostgresqls in the indexer.
func (s *protectionBackupPolicyPostgresqlLister) List(selector labels.Selector) (ret []*v1alpha1.ProtectionBackupPolicyPostgresql, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProtectionBackupPolicyPostgresql))
	})
	return ret, err
}

// ProtectionBackupPolicyPostgresqls returns an object that can list and get ProtectionBackupPolicyPostgresqls.
func (s *protectionBackupPolicyPostgresqlLister) ProtectionBackupPolicyPostgresqls(namespace string) ProtectionBackupPolicyPostgresqlNamespaceLister {
	return protectionBackupPolicyPostgresqlNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProtectionBackupPolicyPostgresqlNamespaceLister helps list and get ProtectionBackupPolicyPostgresqls.
// All objects returned here must be treated as read-only.
type ProtectionBackupPolicyPostgresqlNamespaceLister interface {
	// List lists all ProtectionBackupPolicyPostgresqls in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProtectionBackupPolicyPostgresql, err error)
	// Get retrieves the ProtectionBackupPolicyPostgresql from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ProtectionBackupPolicyPostgresql, error)
	ProtectionBackupPolicyPostgresqlNamespaceListerExpansion
}

// protectionBackupPolicyPostgresqlNamespaceLister implements the ProtectionBackupPolicyPostgresqlNamespaceLister
// interface.
type protectionBackupPolicyPostgresqlNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ProtectionBackupPolicyPostgresqls in the indexer for a given namespace.
func (s protectionBackupPolicyPostgresqlNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ProtectionBackupPolicyPostgresql, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProtectionBackupPolicyPostgresql))
	})
	return ret, err
}

// Get retrieves the ProtectionBackupPolicyPostgresql from the indexer for a given namespace and name.
func (s protectionBackupPolicyPostgresqlNamespaceLister) Get(name string) (*v1alpha1.ProtectionBackupPolicyPostgresql, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("protectionbackuppolicypostgresql"), name)
	}
	return obj.(*v1alpha1.ProtectionBackupPolicyPostgresql), nil
}
