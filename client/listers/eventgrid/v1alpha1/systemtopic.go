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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/eventgrid/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SystemTopicLister helps list SystemTopics.
// All objects returned here must be treated as read-only.
type SystemTopicLister interface {
	// List lists all SystemTopics in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SystemTopic, err error)
	// SystemTopics returns an object that can list and get SystemTopics.
	SystemTopics(namespace string) SystemTopicNamespaceLister
	SystemTopicListerExpansion
}

// systemTopicLister implements the SystemTopicLister interface.
type systemTopicLister struct {
	indexer cache.Indexer
}

// NewSystemTopicLister returns a new SystemTopicLister.
func NewSystemTopicLister(indexer cache.Indexer) SystemTopicLister {
	return &systemTopicLister{indexer: indexer}
}

// List lists all SystemTopics in the indexer.
func (s *systemTopicLister) List(selector labels.Selector) (ret []*v1alpha1.SystemTopic, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SystemTopic))
	})
	return ret, err
}

// SystemTopics returns an object that can list and get SystemTopics.
func (s *systemTopicLister) SystemTopics(namespace string) SystemTopicNamespaceLister {
	return systemTopicNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SystemTopicNamespaceLister helps list and get SystemTopics.
// All objects returned here must be treated as read-only.
type SystemTopicNamespaceLister interface {
	// List lists all SystemTopics in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SystemTopic, err error)
	// Get retrieves the SystemTopic from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SystemTopic, error)
	SystemTopicNamespaceListerExpansion
}

// systemTopicNamespaceLister implements the SystemTopicNamespaceLister
// interface.
type systemTopicNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SystemTopics in the indexer for a given namespace.
func (s systemTopicNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SystemTopic, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SystemTopic))
	})
	return ret, err
}

// Get retrieves the SystemTopic from the indexer for a given namespace and name.
func (s systemTopicNamespaceLister) Get(name string) (*v1alpha1.SystemTopic, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("systemtopic"), name)
	}
	return obj.(*v1alpha1.SystemTopic), nil
}
