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

// GlobalVmShutdownScheduleLister helps list GlobalVmShutdownSchedules.
// All objects returned here must be treated as read-only.
type GlobalVmShutdownScheduleLister interface {
	// List lists all GlobalVmShutdownSchedules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GlobalVmShutdownSchedule, err error)
	// GlobalVmShutdownSchedules returns an object that can list and get GlobalVmShutdownSchedules.
	GlobalVmShutdownSchedules(namespace string) GlobalVmShutdownScheduleNamespaceLister
	GlobalVmShutdownScheduleListerExpansion
}

// globalVmShutdownScheduleLister implements the GlobalVmShutdownScheduleLister interface.
type globalVmShutdownScheduleLister struct {
	indexer cache.Indexer
}

// NewGlobalVmShutdownScheduleLister returns a new GlobalVmShutdownScheduleLister.
func NewGlobalVmShutdownScheduleLister(indexer cache.Indexer) GlobalVmShutdownScheduleLister {
	return &globalVmShutdownScheduleLister{indexer: indexer}
}

// List lists all GlobalVmShutdownSchedules in the indexer.
func (s *globalVmShutdownScheduleLister) List(selector labels.Selector) (ret []*v1alpha1.GlobalVmShutdownSchedule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GlobalVmShutdownSchedule))
	})
	return ret, err
}

// GlobalVmShutdownSchedules returns an object that can list and get GlobalVmShutdownSchedules.
func (s *globalVmShutdownScheduleLister) GlobalVmShutdownSchedules(namespace string) GlobalVmShutdownScheduleNamespaceLister {
	return globalVmShutdownScheduleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GlobalVmShutdownScheduleNamespaceLister helps list and get GlobalVmShutdownSchedules.
// All objects returned here must be treated as read-only.
type GlobalVmShutdownScheduleNamespaceLister interface {
	// List lists all GlobalVmShutdownSchedules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GlobalVmShutdownSchedule, err error)
	// Get retrieves the GlobalVmShutdownSchedule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.GlobalVmShutdownSchedule, error)
	GlobalVmShutdownScheduleNamespaceListerExpansion
}

// globalVmShutdownScheduleNamespaceLister implements the GlobalVmShutdownScheduleNamespaceLister
// interface.
type globalVmShutdownScheduleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GlobalVmShutdownSchedules in the indexer for a given namespace.
func (s globalVmShutdownScheduleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GlobalVmShutdownSchedule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GlobalVmShutdownSchedule))
	})
	return ret, err
}

// Get retrieves the GlobalVmShutdownSchedule from the indexer for a given namespace and name.
func (s globalVmShutdownScheduleNamespaceLister) Get(name string) (*v1alpha1.GlobalVmShutdownSchedule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("globalvmshutdownschedule"), name)
	}
	return obj.(*v1alpha1.GlobalVmShutdownSchedule), nil
}
