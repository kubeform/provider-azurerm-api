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

// LinkedServiceLister helps list LinkedServices.
// All objects returned here must be treated as read-only.
type LinkedServiceLister interface {
	// List lists all LinkedServices in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LinkedService, err error)
	// LinkedServices returns an object that can list and get LinkedServices.
	LinkedServices(namespace string) LinkedServiceNamespaceLister
	LinkedServiceListerExpansion
}

// linkedServiceLister implements the LinkedServiceLister interface.
type linkedServiceLister struct {
	indexer cache.Indexer
}

// NewLinkedServiceLister returns a new LinkedServiceLister.
func NewLinkedServiceLister(indexer cache.Indexer) LinkedServiceLister {
	return &linkedServiceLister{indexer: indexer}
}

// List lists all LinkedServices in the indexer.
func (s *linkedServiceLister) List(selector labels.Selector) (ret []*v1alpha1.LinkedService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LinkedService))
	})
	return ret, err
}

// LinkedServices returns an object that can list and get LinkedServices.
func (s *linkedServiceLister) LinkedServices(namespace string) LinkedServiceNamespaceLister {
	return linkedServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LinkedServiceNamespaceLister helps list and get LinkedServices.
// All objects returned here must be treated as read-only.
type LinkedServiceNamespaceLister interface {
	// List lists all LinkedServices in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LinkedService, err error)
	// Get retrieves the LinkedService from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.LinkedService, error)
	LinkedServiceNamespaceListerExpansion
}

// linkedServiceNamespaceLister implements the LinkedServiceNamespaceLister
// interface.
type linkedServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LinkedServices in the indexer for a given namespace.
func (s linkedServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LinkedService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LinkedService))
	})
	return ret, err
}

// Get retrieves the LinkedService from the indexer for a given namespace and name.
func (s linkedServiceNamespaceLister) Get(name string) (*v1alpha1.LinkedService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("linkedservice"), name)
	}
	return obj.(*v1alpha1.LinkedService), nil
}
