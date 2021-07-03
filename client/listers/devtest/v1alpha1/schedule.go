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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/devtest/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ScheduleLister helps list Schedules.
// All objects returned here must be treated as read-only.
type ScheduleLister interface {
	// List lists all Schedules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Schedule, err error)
	// Schedules returns an object that can list and get Schedules.
	Schedules(namespace string) ScheduleNamespaceLister
	ScheduleListerExpansion
}

// scheduleLister implements the ScheduleLister interface.
type scheduleLister struct {
	indexer cache.Indexer
}

// NewScheduleLister returns a new ScheduleLister.
func NewScheduleLister(indexer cache.Indexer) ScheduleLister {
	return &scheduleLister{indexer: indexer}
}

// List lists all Schedules in the indexer.
func (s *scheduleLister) List(selector labels.Selector) (ret []*v1alpha1.Schedule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Schedule))
	})
	return ret, err
}

// Schedules returns an object that can list and get Schedules.
func (s *scheduleLister) Schedules(namespace string) ScheduleNamespaceLister {
	return scheduleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ScheduleNamespaceLister helps list and get Schedules.
// All objects returned here must be treated as read-only.
type ScheduleNamespaceLister interface {
	// List lists all Schedules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Schedule, err error)
	// Get retrieves the Schedule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Schedule, error)
	ScheduleNamespaceListerExpansion
}

// scheduleNamespaceLister implements the ScheduleNamespaceLister
// interface.
type scheduleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Schedules in the indexer for a given namespace.
func (s scheduleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Schedule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Schedule))
	})
	return ret, err
}

// Get retrieves the Schedule from the indexer for a given namespace and name.
func (s scheduleNamespaceLister) Get(name string) (*v1alpha1.Schedule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("schedule"), name)
	}
	return obj.(*v1alpha1.Schedule), nil
}
