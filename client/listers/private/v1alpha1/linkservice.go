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

// LinkServiceLister helps list LinkServices.
// All objects returned here must be treated as read-only.
type LinkServiceLister interface {
	// List lists all LinkServices in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LinkService, err error)
	// LinkServices returns an object that can list and get LinkServices.
	LinkServices(namespace string) LinkServiceNamespaceLister
	LinkServiceListerExpansion
}

// linkServiceLister implements the LinkServiceLister interface.
type linkServiceLister struct {
	indexer cache.Indexer
}

// NewLinkServiceLister returns a new LinkServiceLister.
func NewLinkServiceLister(indexer cache.Indexer) LinkServiceLister {
	return &linkServiceLister{indexer: indexer}
}

// List lists all LinkServices in the indexer.
func (s *linkServiceLister) List(selector labels.Selector) (ret []*v1alpha1.LinkService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LinkService))
	})
	return ret, err
}

// LinkServices returns an object that can list and get LinkServices.
func (s *linkServiceLister) LinkServices(namespace string) LinkServiceNamespaceLister {
	return linkServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LinkServiceNamespaceLister helps list and get LinkServices.
// All objects returned here must be treated as read-only.
type LinkServiceNamespaceLister interface {
	// List lists all LinkServices in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LinkService, err error)
	// Get retrieves the LinkService from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.LinkService, error)
	LinkServiceNamespaceListerExpansion
}

// linkServiceNamespaceLister implements the LinkServiceNamespaceLister
// interface.
type linkServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LinkServices in the indexer for a given namespace.
func (s linkServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LinkService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LinkService))
	})
	return ret, err
}

// Get retrieves the LinkService from the indexer for a given namespace and name.
func (s linkServiceNamespaceLister) Get(name string) (*v1alpha1.LinkService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("linkservice"), name)
	}
	return obj.(*v1alpha1.LinkService), nil
}
