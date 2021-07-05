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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/security/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CenterSubscriptionPricingLister helps list CenterSubscriptionPricings.
// All objects returned here must be treated as read-only.
type CenterSubscriptionPricingLister interface {
	// List lists all CenterSubscriptionPricings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CenterSubscriptionPricing, err error)
	// CenterSubscriptionPricings returns an object that can list and get CenterSubscriptionPricings.
	CenterSubscriptionPricings(namespace string) CenterSubscriptionPricingNamespaceLister
	CenterSubscriptionPricingListerExpansion
}

// centerSubscriptionPricingLister implements the CenterSubscriptionPricingLister interface.
type centerSubscriptionPricingLister struct {
	indexer cache.Indexer
}

// NewCenterSubscriptionPricingLister returns a new CenterSubscriptionPricingLister.
func NewCenterSubscriptionPricingLister(indexer cache.Indexer) CenterSubscriptionPricingLister {
	return &centerSubscriptionPricingLister{indexer: indexer}
}

// List lists all CenterSubscriptionPricings in the indexer.
func (s *centerSubscriptionPricingLister) List(selector labels.Selector) (ret []*v1alpha1.CenterSubscriptionPricing, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CenterSubscriptionPricing))
	})
	return ret, err
}

// CenterSubscriptionPricings returns an object that can list and get CenterSubscriptionPricings.
func (s *centerSubscriptionPricingLister) CenterSubscriptionPricings(namespace string) CenterSubscriptionPricingNamespaceLister {
	return centerSubscriptionPricingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CenterSubscriptionPricingNamespaceLister helps list and get CenterSubscriptionPricings.
// All objects returned here must be treated as read-only.
type CenterSubscriptionPricingNamespaceLister interface {
	// List lists all CenterSubscriptionPricings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CenterSubscriptionPricing, err error)
	// Get retrieves the CenterSubscriptionPricing from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CenterSubscriptionPricing, error)
	CenterSubscriptionPricingNamespaceListerExpansion
}

// centerSubscriptionPricingNamespaceLister implements the CenterSubscriptionPricingNamespaceLister
// interface.
type centerSubscriptionPricingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CenterSubscriptionPricings in the indexer for a given namespace.
func (s centerSubscriptionPricingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CenterSubscriptionPricing, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CenterSubscriptionPricing))
	})
	return ret, err
}

// Get retrieves the CenterSubscriptionPricing from the indexer for a given namespace and name.
func (s centerSubscriptionPricingNamespaceLister) Get(name string) (*v1alpha1.CenterSubscriptionPricing, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("centersubscriptionpricing"), name)
	}
	return obj.(*v1alpha1.CenterSubscriptionPricing), nil
}