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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/bot/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ChannelsRegistrationLister helps list ChannelsRegistrations.
// All objects returned here must be treated as read-only.
type ChannelsRegistrationLister interface {
	// List lists all ChannelsRegistrations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ChannelsRegistration, err error)
	// ChannelsRegistrations returns an object that can list and get ChannelsRegistrations.
	ChannelsRegistrations(namespace string) ChannelsRegistrationNamespaceLister
	ChannelsRegistrationListerExpansion
}

// channelsRegistrationLister implements the ChannelsRegistrationLister interface.
type channelsRegistrationLister struct {
	indexer cache.Indexer
}

// NewChannelsRegistrationLister returns a new ChannelsRegistrationLister.
func NewChannelsRegistrationLister(indexer cache.Indexer) ChannelsRegistrationLister {
	return &channelsRegistrationLister{indexer: indexer}
}

// List lists all ChannelsRegistrations in the indexer.
func (s *channelsRegistrationLister) List(selector labels.Selector) (ret []*v1alpha1.ChannelsRegistration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ChannelsRegistration))
	})
	return ret, err
}

// ChannelsRegistrations returns an object that can list and get ChannelsRegistrations.
func (s *channelsRegistrationLister) ChannelsRegistrations(namespace string) ChannelsRegistrationNamespaceLister {
	return channelsRegistrationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ChannelsRegistrationNamespaceLister helps list and get ChannelsRegistrations.
// All objects returned here must be treated as read-only.
type ChannelsRegistrationNamespaceLister interface {
	// List lists all ChannelsRegistrations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ChannelsRegistration, err error)
	// Get retrieves the ChannelsRegistration from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ChannelsRegistration, error)
	ChannelsRegistrationNamespaceListerExpansion
}

// channelsRegistrationNamespaceLister implements the ChannelsRegistrationNamespaceLister
// interface.
type channelsRegistrationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ChannelsRegistrations in the indexer for a given namespace.
func (s channelsRegistrationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ChannelsRegistration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ChannelsRegistration))
	})
	return ret, err
}

// Get retrieves the ChannelsRegistration from the indexer for a given namespace and name.
func (s channelsRegistrationNamespaceLister) Get(name string) (*v1alpha1.ChannelsRegistration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("channelsregistration"), name)
	}
	return obj.(*v1alpha1.ChannelsRegistration), nil
}
