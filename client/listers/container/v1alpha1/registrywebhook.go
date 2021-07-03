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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/container/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RegistryWebhookLister helps list RegistryWebhooks.
// All objects returned here must be treated as read-only.
type RegistryWebhookLister interface {
	// List lists all RegistryWebhooks in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RegistryWebhook, err error)
	// RegistryWebhooks returns an object that can list and get RegistryWebhooks.
	RegistryWebhooks(namespace string) RegistryWebhookNamespaceLister
	RegistryWebhookListerExpansion
}

// registryWebhookLister implements the RegistryWebhookLister interface.
type registryWebhookLister struct {
	indexer cache.Indexer
}

// NewRegistryWebhookLister returns a new RegistryWebhookLister.
func NewRegistryWebhookLister(indexer cache.Indexer) RegistryWebhookLister {
	return &registryWebhookLister{indexer: indexer}
}

// List lists all RegistryWebhooks in the indexer.
func (s *registryWebhookLister) List(selector labels.Selector) (ret []*v1alpha1.RegistryWebhook, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RegistryWebhook))
	})
	return ret, err
}

// RegistryWebhooks returns an object that can list and get RegistryWebhooks.
func (s *registryWebhookLister) RegistryWebhooks(namespace string) RegistryWebhookNamespaceLister {
	return registryWebhookNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RegistryWebhookNamespaceLister helps list and get RegistryWebhooks.
// All objects returned here must be treated as read-only.
type RegistryWebhookNamespaceLister interface {
	// List lists all RegistryWebhooks in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RegistryWebhook, err error)
	// Get retrieves the RegistryWebhook from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.RegistryWebhook, error)
	RegistryWebhookNamespaceListerExpansion
}

// registryWebhookNamespaceLister implements the RegistryWebhookNamespaceLister
// interface.
type registryWebhookNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RegistryWebhooks in the indexer for a given namespace.
func (s registryWebhookNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RegistryWebhook, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RegistryWebhook))
	})
	return ret, err
}

// Get retrieves the RegistryWebhook from the indexer for a given namespace and name.
func (s registryWebhookNamespaceLister) Get(name string) (*v1alpha1.RegistryWebhook, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("registrywebhook"), name)
	}
	return obj.(*v1alpha1.RegistryWebhook), nil
}
