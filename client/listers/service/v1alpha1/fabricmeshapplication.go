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

// FabricMeshApplicationLister helps list FabricMeshApplications.
// All objects returned here must be treated as read-only.
type FabricMeshApplicationLister interface {
	// List lists all FabricMeshApplications in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FabricMeshApplication, err error)
	// FabricMeshApplications returns an object that can list and get FabricMeshApplications.
	FabricMeshApplications(namespace string) FabricMeshApplicationNamespaceLister
	FabricMeshApplicationListerExpansion
}

// fabricMeshApplicationLister implements the FabricMeshApplicationLister interface.
type fabricMeshApplicationLister struct {
	indexer cache.Indexer
}

// NewFabricMeshApplicationLister returns a new FabricMeshApplicationLister.
func NewFabricMeshApplicationLister(indexer cache.Indexer) FabricMeshApplicationLister {
	return &fabricMeshApplicationLister{indexer: indexer}
}

// List lists all FabricMeshApplications in the indexer.
func (s *fabricMeshApplicationLister) List(selector labels.Selector) (ret []*v1alpha1.FabricMeshApplication, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FabricMeshApplication))
	})
	return ret, err
}

// FabricMeshApplications returns an object that can list and get FabricMeshApplications.
func (s *fabricMeshApplicationLister) FabricMeshApplications(namespace string) FabricMeshApplicationNamespaceLister {
	return fabricMeshApplicationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FabricMeshApplicationNamespaceLister helps list and get FabricMeshApplications.
// All objects returned here must be treated as read-only.
type FabricMeshApplicationNamespaceLister interface {
	// List lists all FabricMeshApplications in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FabricMeshApplication, err error)
	// Get retrieves the FabricMeshApplication from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FabricMeshApplication, error)
	FabricMeshApplicationNamespaceListerExpansion
}

// fabricMeshApplicationNamespaceLister implements the FabricMeshApplicationNamespaceLister
// interface.
type fabricMeshApplicationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FabricMeshApplications in the indexer for a given namespace.
func (s fabricMeshApplicationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FabricMeshApplication, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FabricMeshApplication))
	})
	return ret, err
}

// Get retrieves the FabricMeshApplication from the indexer for a given namespace and name.
func (s fabricMeshApplicationNamespaceLister) Get(name string) (*v1alpha1.FabricMeshApplication, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("fabricmeshapplication"), name)
	}
	return obj.(*v1alpha1.FabricMeshApplication), nil
}