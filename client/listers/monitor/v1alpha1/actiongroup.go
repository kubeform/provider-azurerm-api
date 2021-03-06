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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/monitor/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ActionGroupLister helps list ActionGroups.
// All objects returned here must be treated as read-only.
type ActionGroupLister interface {
	// List lists all ActionGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ActionGroup, err error)
	// ActionGroups returns an object that can list and get ActionGroups.
	ActionGroups(namespace string) ActionGroupNamespaceLister
	ActionGroupListerExpansion
}

// actionGroupLister implements the ActionGroupLister interface.
type actionGroupLister struct {
	indexer cache.Indexer
}

// NewActionGroupLister returns a new ActionGroupLister.
func NewActionGroupLister(indexer cache.Indexer) ActionGroupLister {
	return &actionGroupLister{indexer: indexer}
}

// List lists all ActionGroups in the indexer.
func (s *actionGroupLister) List(selector labels.Selector) (ret []*v1alpha1.ActionGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ActionGroup))
	})
	return ret, err
}

// ActionGroups returns an object that can list and get ActionGroups.
func (s *actionGroupLister) ActionGroups(namespace string) ActionGroupNamespaceLister {
	return actionGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ActionGroupNamespaceLister helps list and get ActionGroups.
// All objects returned here must be treated as read-only.
type ActionGroupNamespaceLister interface {
	// List lists all ActionGroups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ActionGroup, err error)
	// Get retrieves the ActionGroup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ActionGroup, error)
	ActionGroupNamespaceListerExpansion
}

// actionGroupNamespaceLister implements the ActionGroupNamespaceLister
// interface.
type actionGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ActionGroups in the indexer for a given namespace.
func (s actionGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ActionGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ActionGroup))
	})
	return ret, err
}

// Get retrieves the ActionGroup from the indexer for a given namespace and name.
func (s actionGroupNamespaceLister) Get(name string) (*v1alpha1.ActionGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("actiongroup"), name)
	}
	return obj.(*v1alpha1.ActionGroup), nil
}
