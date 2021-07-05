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

// HubIPLister helps list HubIPs.
// All objects returned here must be treated as read-only.
type HubIPLister interface {
	// List lists all HubIPs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HubIP, err error)
	// HubIPs returns an object that can list and get HubIPs.
	HubIPs(namespace string) HubIPNamespaceLister
	HubIPListerExpansion
}

// hubIPLister implements the HubIPLister interface.
type hubIPLister struct {
	indexer cache.Indexer
}

// NewHubIPLister returns a new HubIPLister.
func NewHubIPLister(indexer cache.Indexer) HubIPLister {
	return &hubIPLister{indexer: indexer}
}

// List lists all HubIPs in the indexer.
func (s *hubIPLister) List(selector labels.Selector) (ret []*v1alpha1.HubIP, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HubIP))
	})
	return ret, err
}

// HubIPs returns an object that can list and get HubIPs.
func (s *hubIPLister) HubIPs(namespace string) HubIPNamespaceLister {
	return hubIPNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HubIPNamespaceLister helps list and get HubIPs.
// All objects returned here must be treated as read-only.
type HubIPNamespaceLister interface {
	// List lists all HubIPs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HubIP, err error)
	// Get retrieves the HubIP from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.HubIP, error)
	HubIPNamespaceListerExpansion
}

// hubIPNamespaceLister implements the HubIPNamespaceLister
// interface.
type hubIPNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HubIPs in the indexer for a given namespace.
func (s hubIPNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HubIP, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HubIP))
	})
	return ret, err
}

// Get retrieves the HubIP from the indexer for a given namespace and name.
func (s hubIPNamespaceLister) Get(name string) (*v1alpha1.HubIP, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("hubip"), name)
	}
	return obj.(*v1alpha1.HubIP), nil
}