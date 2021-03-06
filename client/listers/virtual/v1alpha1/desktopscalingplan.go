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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/virtual/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DesktopScalingPlanLister helps list DesktopScalingPlans.
// All objects returned here must be treated as read-only.
type DesktopScalingPlanLister interface {
	// List lists all DesktopScalingPlans in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DesktopScalingPlan, err error)
	// DesktopScalingPlans returns an object that can list and get DesktopScalingPlans.
	DesktopScalingPlans(namespace string) DesktopScalingPlanNamespaceLister
	DesktopScalingPlanListerExpansion
}

// desktopScalingPlanLister implements the DesktopScalingPlanLister interface.
type desktopScalingPlanLister struct {
	indexer cache.Indexer
}

// NewDesktopScalingPlanLister returns a new DesktopScalingPlanLister.
func NewDesktopScalingPlanLister(indexer cache.Indexer) DesktopScalingPlanLister {
	return &desktopScalingPlanLister{indexer: indexer}
}

// List lists all DesktopScalingPlans in the indexer.
func (s *desktopScalingPlanLister) List(selector labels.Selector) (ret []*v1alpha1.DesktopScalingPlan, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DesktopScalingPlan))
	})
	return ret, err
}

// DesktopScalingPlans returns an object that can list and get DesktopScalingPlans.
func (s *desktopScalingPlanLister) DesktopScalingPlans(namespace string) DesktopScalingPlanNamespaceLister {
	return desktopScalingPlanNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DesktopScalingPlanNamespaceLister helps list and get DesktopScalingPlans.
// All objects returned here must be treated as read-only.
type DesktopScalingPlanNamespaceLister interface {
	// List lists all DesktopScalingPlans in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DesktopScalingPlan, err error)
	// Get retrieves the DesktopScalingPlan from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DesktopScalingPlan, error)
	DesktopScalingPlanNamespaceListerExpansion
}

// desktopScalingPlanNamespaceLister implements the DesktopScalingPlanNamespaceLister
// interface.
type desktopScalingPlanNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DesktopScalingPlans in the indexer for a given namespace.
func (s desktopScalingPlanNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DesktopScalingPlan, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DesktopScalingPlan))
	})
	return ret, err
}

// Get retrieves the DesktopScalingPlan from the indexer for a given namespace and name.
func (s desktopScalingPlanNamespaceLister) Get(name string) (*v1alpha1.DesktopScalingPlan, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("desktopscalingplan"), name)
	}
	return obj.(*v1alpha1.DesktopScalingPlan), nil
}
