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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/kusto/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AttachedDatabaseConfigurationLister helps list AttachedDatabaseConfigurations.
// All objects returned here must be treated as read-only.
type AttachedDatabaseConfigurationLister interface {
	// List lists all AttachedDatabaseConfigurations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AttachedDatabaseConfiguration, err error)
	// AttachedDatabaseConfigurations returns an object that can list and get AttachedDatabaseConfigurations.
	AttachedDatabaseConfigurations(namespace string) AttachedDatabaseConfigurationNamespaceLister
	AttachedDatabaseConfigurationListerExpansion
}

// attachedDatabaseConfigurationLister implements the AttachedDatabaseConfigurationLister interface.
type attachedDatabaseConfigurationLister struct {
	indexer cache.Indexer
}

// NewAttachedDatabaseConfigurationLister returns a new AttachedDatabaseConfigurationLister.
func NewAttachedDatabaseConfigurationLister(indexer cache.Indexer) AttachedDatabaseConfigurationLister {
	return &attachedDatabaseConfigurationLister{indexer: indexer}
}

// List lists all AttachedDatabaseConfigurations in the indexer.
func (s *attachedDatabaseConfigurationLister) List(selector labels.Selector) (ret []*v1alpha1.AttachedDatabaseConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AttachedDatabaseConfiguration))
	})
	return ret, err
}

// AttachedDatabaseConfigurations returns an object that can list and get AttachedDatabaseConfigurations.
func (s *attachedDatabaseConfigurationLister) AttachedDatabaseConfigurations(namespace string) AttachedDatabaseConfigurationNamespaceLister {
	return attachedDatabaseConfigurationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AttachedDatabaseConfigurationNamespaceLister helps list and get AttachedDatabaseConfigurations.
// All objects returned here must be treated as read-only.
type AttachedDatabaseConfigurationNamespaceLister interface {
	// List lists all AttachedDatabaseConfigurations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AttachedDatabaseConfiguration, err error)
	// Get retrieves the AttachedDatabaseConfiguration from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AttachedDatabaseConfiguration, error)
	AttachedDatabaseConfigurationNamespaceListerExpansion
}

// attachedDatabaseConfigurationNamespaceLister implements the AttachedDatabaseConfigurationNamespaceLister
// interface.
type attachedDatabaseConfigurationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AttachedDatabaseConfigurations in the indexer for a given namespace.
func (s attachedDatabaseConfigurationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AttachedDatabaseConfiguration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AttachedDatabaseConfiguration))
	})
	return ret, err
}

// Get retrieves the AttachedDatabaseConfiguration from the indexer for a given namespace and name.
func (s attachedDatabaseConfigurationNamespaceLister) Get(name string) (*v1alpha1.AttachedDatabaseConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("attacheddatabaseconfiguration"), name)
	}
	return obj.(*v1alpha1.AttachedDatabaseConfiguration), nil
}