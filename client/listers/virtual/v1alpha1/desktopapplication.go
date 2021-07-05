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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/virtual/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DesktopApplicationLister helps list DesktopApplications.
// All objects returned here must be treated as read-only.
type DesktopApplicationLister interface {
	// List lists all DesktopApplications in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DesktopApplication, err error)
	// DesktopApplications returns an object that can list and get DesktopApplications.
	DesktopApplications(namespace string) DesktopApplicationNamespaceLister
	DesktopApplicationListerExpansion
}

// desktopApplicationLister implements the DesktopApplicationLister interface.
type desktopApplicationLister struct {
	indexer cache.Indexer
}

// NewDesktopApplicationLister returns a new DesktopApplicationLister.
func NewDesktopApplicationLister(indexer cache.Indexer) DesktopApplicationLister {
	return &desktopApplicationLister{indexer: indexer}
}

// List lists all DesktopApplications in the indexer.
func (s *desktopApplicationLister) List(selector labels.Selector) (ret []*v1alpha1.DesktopApplication, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DesktopApplication))
	})
	return ret, err
}

// DesktopApplications returns an object that can list and get DesktopApplications.
func (s *desktopApplicationLister) DesktopApplications(namespace string) DesktopApplicationNamespaceLister {
	return desktopApplicationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DesktopApplicationNamespaceLister helps list and get DesktopApplications.
// All objects returned here must be treated as read-only.
type DesktopApplicationNamespaceLister interface {
	// List lists all DesktopApplications in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DesktopApplication, err error)
	// Get retrieves the DesktopApplication from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DesktopApplication, error)
	DesktopApplicationNamespaceListerExpansion
}

// desktopApplicationNamespaceLister implements the DesktopApplicationNamespaceLister
// interface.
type desktopApplicationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DesktopApplications in the indexer for a given namespace.
func (s desktopApplicationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DesktopApplication, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DesktopApplication))
	})
	return ret, err
}

// Get retrieves the DesktopApplication from the indexer for a given namespace and name.
func (s desktopApplicationNamespaceLister) Get(name string) (*v1alpha1.DesktopApplication, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("desktopapplication"), name)
	}
	return obj.(*v1alpha1.DesktopApplication), nil
}