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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/synapse/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IntegrationRuntimeAzureLister helps list IntegrationRuntimeAzures.
// All objects returned here must be treated as read-only.
type IntegrationRuntimeAzureLister interface {
	// List lists all IntegrationRuntimeAzures in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IntegrationRuntimeAzure, err error)
	// IntegrationRuntimeAzures returns an object that can list and get IntegrationRuntimeAzures.
	IntegrationRuntimeAzures(namespace string) IntegrationRuntimeAzureNamespaceLister
	IntegrationRuntimeAzureListerExpansion
}

// integrationRuntimeAzureLister implements the IntegrationRuntimeAzureLister interface.
type integrationRuntimeAzureLister struct {
	indexer cache.Indexer
}

// NewIntegrationRuntimeAzureLister returns a new IntegrationRuntimeAzureLister.
func NewIntegrationRuntimeAzureLister(indexer cache.Indexer) IntegrationRuntimeAzureLister {
	return &integrationRuntimeAzureLister{indexer: indexer}
}

// List lists all IntegrationRuntimeAzures in the indexer.
func (s *integrationRuntimeAzureLister) List(selector labels.Selector) (ret []*v1alpha1.IntegrationRuntimeAzure, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IntegrationRuntimeAzure))
	})
	return ret, err
}

// IntegrationRuntimeAzures returns an object that can list and get IntegrationRuntimeAzures.
func (s *integrationRuntimeAzureLister) IntegrationRuntimeAzures(namespace string) IntegrationRuntimeAzureNamespaceLister {
	return integrationRuntimeAzureNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IntegrationRuntimeAzureNamespaceLister helps list and get IntegrationRuntimeAzures.
// All objects returned here must be treated as read-only.
type IntegrationRuntimeAzureNamespaceLister interface {
	// List lists all IntegrationRuntimeAzures in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IntegrationRuntimeAzure, err error)
	// Get retrieves the IntegrationRuntimeAzure from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.IntegrationRuntimeAzure, error)
	IntegrationRuntimeAzureNamespaceListerExpansion
}

// integrationRuntimeAzureNamespaceLister implements the IntegrationRuntimeAzureNamespaceLister
// interface.
type integrationRuntimeAzureNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IntegrationRuntimeAzures in the indexer for a given namespace.
func (s integrationRuntimeAzureNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IntegrationRuntimeAzure, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IntegrationRuntimeAzure))
	})
	return ret, err
}

// Get retrieves the IntegrationRuntimeAzure from the indexer for a given namespace and name.
func (s integrationRuntimeAzureNamespaceLister) Get(name string) (*v1alpha1.IntegrationRuntimeAzure, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("integrationruntimeazure"), name)
	}
	return obj.(*v1alpha1.IntegrationRuntimeAzure), nil
}
