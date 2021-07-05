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

// CloudAppRedisAssociationLister helps list CloudAppRedisAssociations.
// All objects returned here must be treated as read-only.
type CloudAppRedisAssociationLister interface {
	// List lists all CloudAppRedisAssociations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CloudAppRedisAssociation, err error)
	// CloudAppRedisAssociations returns an object that can list and get CloudAppRedisAssociations.
	CloudAppRedisAssociations(namespace string) CloudAppRedisAssociationNamespaceLister
	CloudAppRedisAssociationListerExpansion
}

// cloudAppRedisAssociationLister implements the CloudAppRedisAssociationLister interface.
type cloudAppRedisAssociationLister struct {
	indexer cache.Indexer
}

// NewCloudAppRedisAssociationLister returns a new CloudAppRedisAssociationLister.
func NewCloudAppRedisAssociationLister(indexer cache.Indexer) CloudAppRedisAssociationLister {
	return &cloudAppRedisAssociationLister{indexer: indexer}
}

// List lists all CloudAppRedisAssociations in the indexer.
func (s *cloudAppRedisAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.CloudAppRedisAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudAppRedisAssociation))
	})
	return ret, err
}

// CloudAppRedisAssociations returns an object that can list and get CloudAppRedisAssociations.
func (s *cloudAppRedisAssociationLister) CloudAppRedisAssociations(namespace string) CloudAppRedisAssociationNamespaceLister {
	return cloudAppRedisAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CloudAppRedisAssociationNamespaceLister helps list and get CloudAppRedisAssociations.
// All objects returned here must be treated as read-only.
type CloudAppRedisAssociationNamespaceLister interface {
	// List lists all CloudAppRedisAssociations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CloudAppRedisAssociation, err error)
	// Get retrieves the CloudAppRedisAssociation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CloudAppRedisAssociation, error)
	CloudAppRedisAssociationNamespaceListerExpansion
}

// cloudAppRedisAssociationNamespaceLister implements the CloudAppRedisAssociationNamespaceLister
// interface.
type cloudAppRedisAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CloudAppRedisAssociations in the indexer for a given namespace.
func (s cloudAppRedisAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CloudAppRedisAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudAppRedisAssociation))
	})
	return ret, err
}

// Get retrieves the CloudAppRedisAssociation from the indexer for a given namespace and name.
func (s cloudAppRedisAssociationNamespaceLister) Get(name string) (*v1alpha1.CloudAppRedisAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cloudappredisassociation"), name)
	}
	return obj.(*v1alpha1.CloudAppRedisAssociation), nil
}