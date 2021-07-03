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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/service/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FabricMeshSecretLister helps list FabricMeshSecrets.
// All objects returned here must be treated as read-only.
type FabricMeshSecretLister interface {
	// List lists all FabricMeshSecrets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FabricMeshSecret, err error)
	// FabricMeshSecrets returns an object that can list and get FabricMeshSecrets.
	FabricMeshSecrets(namespace string) FabricMeshSecretNamespaceLister
	FabricMeshSecretListerExpansion
}

// fabricMeshSecretLister implements the FabricMeshSecretLister interface.
type fabricMeshSecretLister struct {
	indexer cache.Indexer
}

// NewFabricMeshSecretLister returns a new FabricMeshSecretLister.
func NewFabricMeshSecretLister(indexer cache.Indexer) FabricMeshSecretLister {
	return &fabricMeshSecretLister{indexer: indexer}
}

// List lists all FabricMeshSecrets in the indexer.
func (s *fabricMeshSecretLister) List(selector labels.Selector) (ret []*v1alpha1.FabricMeshSecret, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FabricMeshSecret))
	})
	return ret, err
}

// FabricMeshSecrets returns an object that can list and get FabricMeshSecrets.
func (s *fabricMeshSecretLister) FabricMeshSecrets(namespace string) FabricMeshSecretNamespaceLister {
	return fabricMeshSecretNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FabricMeshSecretNamespaceLister helps list and get FabricMeshSecrets.
// All objects returned here must be treated as read-only.
type FabricMeshSecretNamespaceLister interface {
	// List lists all FabricMeshSecrets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FabricMeshSecret, err error)
	// Get retrieves the FabricMeshSecret from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FabricMeshSecret, error)
	FabricMeshSecretNamespaceListerExpansion
}

// fabricMeshSecretNamespaceLister implements the FabricMeshSecretNamespaceLister
// interface.
type fabricMeshSecretNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FabricMeshSecrets in the indexer for a given namespace.
func (s fabricMeshSecretNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FabricMeshSecret, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FabricMeshSecret))
	})
	return ret, err
}

// Get retrieves the FabricMeshSecret from the indexer for a given namespace and name.
func (s fabricMeshSecretNamespaceLister) Get(name string) (*v1alpha1.FabricMeshSecret, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("fabricmeshsecret"), name)
	}
	return obj.(*v1alpha1.FabricMeshSecret), nil
}
