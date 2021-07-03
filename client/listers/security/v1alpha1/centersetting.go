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

// CenterSettingLister helps list CenterSettings.
// All objects returned here must be treated as read-only.
type CenterSettingLister interface {
	// List lists all CenterSettings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CenterSetting, err error)
	// CenterSettings returns an object that can list and get CenterSettings.
	CenterSettings(namespace string) CenterSettingNamespaceLister
	CenterSettingListerExpansion
}

// centerSettingLister implements the CenterSettingLister interface.
type centerSettingLister struct {
	indexer cache.Indexer
}

// NewCenterSettingLister returns a new CenterSettingLister.
func NewCenterSettingLister(indexer cache.Indexer) CenterSettingLister {
	return &centerSettingLister{indexer: indexer}
}

// List lists all CenterSettings in the indexer.
func (s *centerSettingLister) List(selector labels.Selector) (ret []*v1alpha1.CenterSetting, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CenterSetting))
	})
	return ret, err
}

// CenterSettings returns an object that can list and get CenterSettings.
func (s *centerSettingLister) CenterSettings(namespace string) CenterSettingNamespaceLister {
	return centerSettingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CenterSettingNamespaceLister helps list and get CenterSettings.
// All objects returned here must be treated as read-only.
type CenterSettingNamespaceLister interface {
	// List lists all CenterSettings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CenterSetting, err error)
	// Get retrieves the CenterSetting from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CenterSetting, error)
	CenterSettingNamespaceListerExpansion
}

// centerSettingNamespaceLister implements the CenterSettingNamespaceLister
// interface.
type centerSettingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CenterSettings in the indexer for a given namespace.
func (s centerSettingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CenterSetting, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CenterSetting))
	})
	return ret, err
}

// Get retrieves the CenterSetting from the indexer for a given namespace and name.
func (s centerSettingNamespaceLister) Get(name string) (*v1alpha1.CenterSetting, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("centersetting"), name)
	}
	return obj.(*v1alpha1.CenterSetting), nil
}
