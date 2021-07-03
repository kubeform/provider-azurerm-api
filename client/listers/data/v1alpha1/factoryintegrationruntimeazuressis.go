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

// FactoryIntegrationRuntimeAzureSsisLister helps list FactoryIntegrationRuntimeAzureSsises.
// All objects returned here must be treated as read-only.
type FactoryIntegrationRuntimeAzureSsisLister interface {
	// List lists all FactoryIntegrationRuntimeAzureSsises in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FactoryIntegrationRuntimeAzureSsis, err error)
	// FactoryIntegrationRuntimeAzureSsises returns an object that can list and get FactoryIntegrationRuntimeAzureSsises.
	FactoryIntegrationRuntimeAzureSsises(namespace string) FactoryIntegrationRuntimeAzureSsisNamespaceLister
	FactoryIntegrationRuntimeAzureSsisListerExpansion
}

// factoryIntegrationRuntimeAzureSsisLister implements the FactoryIntegrationRuntimeAzureSsisLister interface.
type factoryIntegrationRuntimeAzureSsisLister struct {
	indexer cache.Indexer
}

// NewFactoryIntegrationRuntimeAzureSsisLister returns a new FactoryIntegrationRuntimeAzureSsisLister.
func NewFactoryIntegrationRuntimeAzureSsisLister(indexer cache.Indexer) FactoryIntegrationRuntimeAzureSsisLister {
	return &factoryIntegrationRuntimeAzureSsisLister{indexer: indexer}
}

// List lists all FactoryIntegrationRuntimeAzureSsises in the indexer.
func (s *factoryIntegrationRuntimeAzureSsisLister) List(selector labels.Selector) (ret []*v1alpha1.FactoryIntegrationRuntimeAzureSsis, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FactoryIntegrationRuntimeAzureSsis))
	})
	return ret, err
}

// FactoryIntegrationRuntimeAzureSsises returns an object that can list and get FactoryIntegrationRuntimeAzureSsises.
func (s *factoryIntegrationRuntimeAzureSsisLister) FactoryIntegrationRuntimeAzureSsises(namespace string) FactoryIntegrationRuntimeAzureSsisNamespaceLister {
	return factoryIntegrationRuntimeAzureSsisNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FactoryIntegrationRuntimeAzureSsisNamespaceLister helps list and get FactoryIntegrationRuntimeAzureSsises.
// All objects returned here must be treated as read-only.
type FactoryIntegrationRuntimeAzureSsisNamespaceLister interface {
	// List lists all FactoryIntegrationRuntimeAzureSsises in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FactoryIntegrationRuntimeAzureSsis, err error)
	// Get retrieves the FactoryIntegrationRuntimeAzureSsis from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FactoryIntegrationRuntimeAzureSsis, error)
	FactoryIntegrationRuntimeAzureSsisNamespaceListerExpansion
}

// factoryIntegrationRuntimeAzureSsisNamespaceLister implements the FactoryIntegrationRuntimeAzureSsisNamespaceLister
// interface.
type factoryIntegrationRuntimeAzureSsisNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FactoryIntegrationRuntimeAzureSsises in the indexer for a given namespace.
func (s factoryIntegrationRuntimeAzureSsisNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FactoryIntegrationRuntimeAzureSsis, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FactoryIntegrationRuntimeAzureSsis))
	})
	return ret, err
}

// Get retrieves the FactoryIntegrationRuntimeAzureSsis from the indexer for a given namespace and name.
func (s factoryIntegrationRuntimeAzureSsisNamespaceLister) Get(name string) (*v1alpha1.FactoryIntegrationRuntimeAzureSsis, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("factoryintegrationruntimeazuressis"), name)
	}
	return obj.(*v1alpha1.FactoryIntegrationRuntimeAzureSsis), nil
}
