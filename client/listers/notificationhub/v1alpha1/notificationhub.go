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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/notificationhub/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NotificationHubLister helps list NotificationHubs.
// All objects returned here must be treated as read-only.
type NotificationHubLister interface {
	// List lists all NotificationHubs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NotificationHub, err error)
	// NotificationHubs returns an object that can list and get NotificationHubs.
	NotificationHubs(namespace string) NotificationHubNamespaceLister
	NotificationHubListerExpansion
}

// notificationHubLister implements the NotificationHubLister interface.
type notificationHubLister struct {
	indexer cache.Indexer
}

// NewNotificationHubLister returns a new NotificationHubLister.
func NewNotificationHubLister(indexer cache.Indexer) NotificationHubLister {
	return &notificationHubLister{indexer: indexer}
}

// List lists all NotificationHubs in the indexer.
func (s *notificationHubLister) List(selector labels.Selector) (ret []*v1alpha1.NotificationHub, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NotificationHub))
	})
	return ret, err
}

// NotificationHubs returns an object that can list and get NotificationHubs.
func (s *notificationHubLister) NotificationHubs(namespace string) NotificationHubNamespaceLister {
	return notificationHubNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NotificationHubNamespaceLister helps list and get NotificationHubs.
// All objects returned here must be treated as read-only.
type NotificationHubNamespaceLister interface {
	// List lists all NotificationHubs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NotificationHub, err error)
	// Get retrieves the NotificationHub from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.NotificationHub, error)
	NotificationHubNamespaceListerExpansion
}

// notificationHubNamespaceLister implements the NotificationHubNamespaceLister
// interface.
type notificationHubNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NotificationHubs in the indexer for a given namespace.
func (s notificationHubNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NotificationHub, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NotificationHub))
	})
	return ret, err
}

// Get retrieves the NotificationHub from the indexer for a given namespace and name.
func (s notificationHubNamespaceLister) Get(name string) (*v1alpha1.NotificationHub, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("notificationhub"), name)
	}
	return obj.(*v1alpha1.NotificationHub), nil
}
