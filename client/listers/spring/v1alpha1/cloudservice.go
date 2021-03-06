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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/spring/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CloudServiceLister helps list CloudServices.
// All objects returned here must be treated as read-only.
type CloudServiceLister interface {
	// List lists all CloudServices in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CloudService, err error)
	// CloudServices returns an object that can list and get CloudServices.
	CloudServices(namespace string) CloudServiceNamespaceLister
	CloudServiceListerExpansion
}

// cloudServiceLister implements the CloudServiceLister interface.
type cloudServiceLister struct {
	indexer cache.Indexer
}

// NewCloudServiceLister returns a new CloudServiceLister.
func NewCloudServiceLister(indexer cache.Indexer) CloudServiceLister {
	return &cloudServiceLister{indexer: indexer}
}

// List lists all CloudServices in the indexer.
func (s *cloudServiceLister) List(selector labels.Selector) (ret []*v1alpha1.CloudService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudService))
	})
	return ret, err
}

// CloudServices returns an object that can list and get CloudServices.
func (s *cloudServiceLister) CloudServices(namespace string) CloudServiceNamespaceLister {
	return cloudServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CloudServiceNamespaceLister helps list and get CloudServices.
// All objects returned here must be treated as read-only.
type CloudServiceNamespaceLister interface {
	// List lists all CloudServices in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CloudService, err error)
	// Get retrieves the CloudService from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CloudService, error)
	CloudServiceNamespaceListerExpansion
}

// cloudServiceNamespaceLister implements the CloudServiceNamespaceLister
// interface.
type cloudServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CloudServices in the indexer for a given namespace.
func (s cloudServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CloudService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudService))
	})
	return ret, err
}

// Get retrieves the CloudService from the indexer for a given namespace and name.
func (s cloudServiceNamespaceLister) Get(name string) (*v1alpha1.CloudService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cloudservice"), name)
	}
	return obj.(*v1alpha1.CloudService), nil
}
