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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/app/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServiceCertificateLister helps list ServiceCertificates.
// All objects returned here must be treated as read-only.
type ServiceCertificateLister interface {
	// List lists all ServiceCertificates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceCertificate, err error)
	// ServiceCertificates returns an object that can list and get ServiceCertificates.
	ServiceCertificates(namespace string) ServiceCertificateNamespaceLister
	ServiceCertificateListerExpansion
}

// serviceCertificateLister implements the ServiceCertificateLister interface.
type serviceCertificateLister struct {
	indexer cache.Indexer
}

// NewServiceCertificateLister returns a new ServiceCertificateLister.
func NewServiceCertificateLister(indexer cache.Indexer) ServiceCertificateLister {
	return &serviceCertificateLister{indexer: indexer}
}

// List lists all ServiceCertificates in the indexer.
func (s *serviceCertificateLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceCertificate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceCertificate))
	})
	return ret, err
}

// ServiceCertificates returns an object that can list and get ServiceCertificates.
func (s *serviceCertificateLister) ServiceCertificates(namespace string) ServiceCertificateNamespaceLister {
	return serviceCertificateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServiceCertificateNamespaceLister helps list and get ServiceCertificates.
// All objects returned here must be treated as read-only.
type ServiceCertificateNamespaceLister interface {
	// List lists all ServiceCertificates in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceCertificate, err error)
	// Get retrieves the ServiceCertificate from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ServiceCertificate, error)
	ServiceCertificateNamespaceListerExpansion
}

// serviceCertificateNamespaceLister implements the ServiceCertificateNamespaceLister
// interface.
type serviceCertificateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServiceCertificates in the indexer for a given namespace.
func (s serviceCertificateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceCertificate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceCertificate))
	})
	return ret, err
}

// Get retrieves the ServiceCertificate from the indexer for a given namespace and name.
func (s serviceCertificateNamespaceLister) Get(name string) (*v1alpha1.ServiceCertificate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("servicecertificate"), name)
	}
	return obj.(*v1alpha1.ServiceCertificate), nil
}
