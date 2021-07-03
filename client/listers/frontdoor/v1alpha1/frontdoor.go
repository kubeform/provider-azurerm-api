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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/frontdoor/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FrontdoorLister helps list Frontdoors.
// All objects returned here must be treated as read-only.
type FrontdoorLister interface {
	// List lists all Frontdoors in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Frontdoor, err error)
	// Frontdoors returns an object that can list and get Frontdoors.
	Frontdoors(namespace string) FrontdoorNamespaceLister
	FrontdoorListerExpansion
}

// frontdoorLister implements the FrontdoorLister interface.
type frontdoorLister struct {
	indexer cache.Indexer
}

// NewFrontdoorLister returns a new FrontdoorLister.
func NewFrontdoorLister(indexer cache.Indexer) FrontdoorLister {
	return &frontdoorLister{indexer: indexer}
}

// List lists all Frontdoors in the indexer.
func (s *frontdoorLister) List(selector labels.Selector) (ret []*v1alpha1.Frontdoor, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Frontdoor))
	})
	return ret, err
}

// Frontdoors returns an object that can list and get Frontdoors.
func (s *frontdoorLister) Frontdoors(namespace string) FrontdoorNamespaceLister {
	return frontdoorNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FrontdoorNamespaceLister helps list and get Frontdoors.
// All objects returned here must be treated as read-only.
type FrontdoorNamespaceLister interface {
	// List lists all Frontdoors in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Frontdoor, err error)
	// Get retrieves the Frontdoor from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Frontdoor, error)
	FrontdoorNamespaceListerExpansion
}

// frontdoorNamespaceLister implements the FrontdoorNamespaceLister
// interface.
type frontdoorNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Frontdoors in the indexer for a given namespace.
func (s frontdoorNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Frontdoor, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Frontdoor))
	})
	return ret, err
}

// Get retrieves the Frontdoor from the indexer for a given namespace and name.
func (s frontdoorNamespaceLister) Get(name string) (*v1alpha1.Frontdoor, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("frontdoor"), name)
	}
	return obj.(*v1alpha1.Frontdoor), nil
}
