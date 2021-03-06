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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/logicapp/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IntegrationAccountLister helps list IntegrationAccounts.
// All objects returned here must be treated as read-only.
type IntegrationAccountLister interface {
	// List lists all IntegrationAccounts in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IntegrationAccount, err error)
	// IntegrationAccounts returns an object that can list and get IntegrationAccounts.
	IntegrationAccounts(namespace string) IntegrationAccountNamespaceLister
	IntegrationAccountListerExpansion
}

// integrationAccountLister implements the IntegrationAccountLister interface.
type integrationAccountLister struct {
	indexer cache.Indexer
}

// NewIntegrationAccountLister returns a new IntegrationAccountLister.
func NewIntegrationAccountLister(indexer cache.Indexer) IntegrationAccountLister {
	return &integrationAccountLister{indexer: indexer}
}

// List lists all IntegrationAccounts in the indexer.
func (s *integrationAccountLister) List(selector labels.Selector) (ret []*v1alpha1.IntegrationAccount, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IntegrationAccount))
	})
	return ret, err
}

// IntegrationAccounts returns an object that can list and get IntegrationAccounts.
func (s *integrationAccountLister) IntegrationAccounts(namespace string) IntegrationAccountNamespaceLister {
	return integrationAccountNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IntegrationAccountNamespaceLister helps list and get IntegrationAccounts.
// All objects returned here must be treated as read-only.
type IntegrationAccountNamespaceLister interface {
	// List lists all IntegrationAccounts in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IntegrationAccount, err error)
	// Get retrieves the IntegrationAccount from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.IntegrationAccount, error)
	IntegrationAccountNamespaceListerExpansion
}

// integrationAccountNamespaceLister implements the IntegrationAccountNamespaceLister
// interface.
type integrationAccountNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IntegrationAccounts in the indexer for a given namespace.
func (s integrationAccountNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IntegrationAccount, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IntegrationAccount))
	})
	return ret, err
}

// Get retrieves the IntegrationAccount from the indexer for a given namespace and name.
func (s integrationAccountNamespaceLister) Get(name string) (*v1alpha1.IntegrationAccount, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("integrationaccount"), name)
	}
	return obj.(*v1alpha1.IntegrationAccount), nil
}
