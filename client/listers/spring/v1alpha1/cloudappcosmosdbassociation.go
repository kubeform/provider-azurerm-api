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

// CloudAppCosmosdbAssociationLister helps list CloudAppCosmosdbAssociations.
// All objects returned here must be treated as read-only.
type CloudAppCosmosdbAssociationLister interface {
	// List lists all CloudAppCosmosdbAssociations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CloudAppCosmosdbAssociation, err error)
	// CloudAppCosmosdbAssociations returns an object that can list and get CloudAppCosmosdbAssociations.
	CloudAppCosmosdbAssociations(namespace string) CloudAppCosmosdbAssociationNamespaceLister
	CloudAppCosmosdbAssociationListerExpansion
}

// cloudAppCosmosdbAssociationLister implements the CloudAppCosmosdbAssociationLister interface.
type cloudAppCosmosdbAssociationLister struct {
	indexer cache.Indexer
}

// NewCloudAppCosmosdbAssociationLister returns a new CloudAppCosmosdbAssociationLister.
func NewCloudAppCosmosdbAssociationLister(indexer cache.Indexer) CloudAppCosmosdbAssociationLister {
	return &cloudAppCosmosdbAssociationLister{indexer: indexer}
}

// List lists all CloudAppCosmosdbAssociations in the indexer.
func (s *cloudAppCosmosdbAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.CloudAppCosmosdbAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudAppCosmosdbAssociation))
	})
	return ret, err
}

// CloudAppCosmosdbAssociations returns an object that can list and get CloudAppCosmosdbAssociations.
func (s *cloudAppCosmosdbAssociationLister) CloudAppCosmosdbAssociations(namespace string) CloudAppCosmosdbAssociationNamespaceLister {
	return cloudAppCosmosdbAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CloudAppCosmosdbAssociationNamespaceLister helps list and get CloudAppCosmosdbAssociations.
// All objects returned here must be treated as read-only.
type CloudAppCosmosdbAssociationNamespaceLister interface {
	// List lists all CloudAppCosmosdbAssociations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CloudAppCosmosdbAssociation, err error)
	// Get retrieves the CloudAppCosmosdbAssociation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CloudAppCosmosdbAssociation, error)
	CloudAppCosmosdbAssociationNamespaceListerExpansion
}

// cloudAppCosmosdbAssociationNamespaceLister implements the CloudAppCosmosdbAssociationNamespaceLister
// interface.
type cloudAppCosmosdbAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CloudAppCosmosdbAssociations in the indexer for a given namespace.
func (s cloudAppCosmosdbAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CloudAppCosmosdbAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudAppCosmosdbAssociation))
	})
	return ret, err
}

// Get retrieves the CloudAppCosmosdbAssociation from the indexer for a given namespace and name.
func (s cloudAppCosmosdbAssociationNamespaceLister) Get(name string) (*v1alpha1.CloudAppCosmosdbAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cloudappcosmosdbassociation"), name)
	}
	return obj.(*v1alpha1.CloudAppCosmosdbAssociation), nil
}