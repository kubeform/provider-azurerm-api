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

// FactoryDatasetAzureBlobLister helps list FactoryDatasetAzureBlobs.
// All objects returned here must be treated as read-only.
type FactoryDatasetAzureBlobLister interface {
	// List lists all FactoryDatasetAzureBlobs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FactoryDatasetAzureBlob, err error)
	// FactoryDatasetAzureBlobs returns an object that can list and get FactoryDatasetAzureBlobs.
	FactoryDatasetAzureBlobs(namespace string) FactoryDatasetAzureBlobNamespaceLister
	FactoryDatasetAzureBlobListerExpansion
}

// factoryDatasetAzureBlobLister implements the FactoryDatasetAzureBlobLister interface.
type factoryDatasetAzureBlobLister struct {
	indexer cache.Indexer
}

// NewFactoryDatasetAzureBlobLister returns a new FactoryDatasetAzureBlobLister.
func NewFactoryDatasetAzureBlobLister(indexer cache.Indexer) FactoryDatasetAzureBlobLister {
	return &factoryDatasetAzureBlobLister{indexer: indexer}
}

// List lists all FactoryDatasetAzureBlobs in the indexer.
func (s *factoryDatasetAzureBlobLister) List(selector labels.Selector) (ret []*v1alpha1.FactoryDatasetAzureBlob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FactoryDatasetAzureBlob))
	})
	return ret, err
}

// FactoryDatasetAzureBlobs returns an object that can list and get FactoryDatasetAzureBlobs.
func (s *factoryDatasetAzureBlobLister) FactoryDatasetAzureBlobs(namespace string) FactoryDatasetAzureBlobNamespaceLister {
	return factoryDatasetAzureBlobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FactoryDatasetAzureBlobNamespaceLister helps list and get FactoryDatasetAzureBlobs.
// All objects returned here must be treated as read-only.
type FactoryDatasetAzureBlobNamespaceLister interface {
	// List lists all FactoryDatasetAzureBlobs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FactoryDatasetAzureBlob, err error)
	// Get retrieves the FactoryDatasetAzureBlob from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FactoryDatasetAzureBlob, error)
	FactoryDatasetAzureBlobNamespaceListerExpansion
}

// factoryDatasetAzureBlobNamespaceLister implements the FactoryDatasetAzureBlobNamespaceLister
// interface.
type factoryDatasetAzureBlobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FactoryDatasetAzureBlobs in the indexer for a given namespace.
func (s factoryDatasetAzureBlobNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FactoryDatasetAzureBlob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FactoryDatasetAzureBlob))
	})
	return ret, err
}

// Get retrieves the FactoryDatasetAzureBlob from the indexer for a given namespace and name.
func (s factoryDatasetAzureBlobNamespaceLister) Get(name string) (*v1alpha1.FactoryDatasetAzureBlob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("factorydatasetazureblob"), name)
	}
	return obj.(*v1alpha1.FactoryDatasetAzureBlob), nil
}
