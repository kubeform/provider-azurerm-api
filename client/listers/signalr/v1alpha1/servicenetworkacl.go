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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/signalr/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServiceNetworkACLLister helps list ServiceNetworkACLs.
// All objects returned here must be treated as read-only.
type ServiceNetworkACLLister interface {
	// List lists all ServiceNetworkACLs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceNetworkACL, err error)
	// ServiceNetworkACLs returns an object that can list and get ServiceNetworkACLs.
	ServiceNetworkACLs(namespace string) ServiceNetworkACLNamespaceLister
	ServiceNetworkACLListerExpansion
}

// serviceNetworkACLLister implements the ServiceNetworkACLLister interface.
type serviceNetworkACLLister struct {
	indexer cache.Indexer
}

// NewServiceNetworkACLLister returns a new ServiceNetworkACLLister.
func NewServiceNetworkACLLister(indexer cache.Indexer) ServiceNetworkACLLister {
	return &serviceNetworkACLLister{indexer: indexer}
}

// List lists all ServiceNetworkACLs in the indexer.
func (s *serviceNetworkACLLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceNetworkACL, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceNetworkACL))
	})
	return ret, err
}

// ServiceNetworkACLs returns an object that can list and get ServiceNetworkACLs.
func (s *serviceNetworkACLLister) ServiceNetworkACLs(namespace string) ServiceNetworkACLNamespaceLister {
	return serviceNetworkACLNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServiceNetworkACLNamespaceLister helps list and get ServiceNetworkACLs.
// All objects returned here must be treated as read-only.
type ServiceNetworkACLNamespaceLister interface {
	// List lists all ServiceNetworkACLs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceNetworkACL, err error)
	// Get retrieves the ServiceNetworkACL from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ServiceNetworkACL, error)
	ServiceNetworkACLNamespaceListerExpansion
}

// serviceNetworkACLNamespaceLister implements the ServiceNetworkACLNamespaceLister
// interface.
type serviceNetworkACLNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServiceNetworkACLs in the indexer for a given namespace.
func (s serviceNetworkACLNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceNetworkACL, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceNetworkACL))
	})
	return ret, err
}

// Get retrieves the ServiceNetworkACL from the indexer for a given namespace and name.
func (s serviceNetworkACLNamespaceLister) Get(name string) (*v1alpha1.ServiceNetworkACL, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("servicenetworkacl"), name)
	}
	return obj.(*v1alpha1.ServiceNetworkACL), nil
}
