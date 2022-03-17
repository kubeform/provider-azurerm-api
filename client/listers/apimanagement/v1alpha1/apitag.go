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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/apimanagement/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ApiTagLister helps list ApiTags.
// All objects returned here must be treated as read-only.
type ApiTagLister interface {
	// List lists all ApiTags in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ApiTag, err error)
	// ApiTags returns an object that can list and get ApiTags.
	ApiTags(namespace string) ApiTagNamespaceLister
	ApiTagListerExpansion
}

// apiTagLister implements the ApiTagLister interface.
type apiTagLister struct {
	indexer cache.Indexer
}

// NewApiTagLister returns a new ApiTagLister.
func NewApiTagLister(indexer cache.Indexer) ApiTagLister {
	return &apiTagLister{indexer: indexer}
}

// List lists all ApiTags in the indexer.
func (s *apiTagLister) List(selector labels.Selector) (ret []*v1alpha1.ApiTag, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiTag))
	})
	return ret, err
}

// ApiTags returns an object that can list and get ApiTags.
func (s *apiTagLister) ApiTags(namespace string) ApiTagNamespaceLister {
	return apiTagNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApiTagNamespaceLister helps list and get ApiTags.
// All objects returned here must be treated as read-only.
type ApiTagNamespaceLister interface {
	// List lists all ApiTags in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ApiTag, err error)
	// Get retrieves the ApiTag from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ApiTag, error)
	ApiTagNamespaceListerExpansion
}

// apiTagNamespaceLister implements the ApiTagNamespaceLister
// interface.
type apiTagNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApiTags in the indexer for a given namespace.
func (s apiTagNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApiTag, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiTag))
	})
	return ret, err
}

// Get retrieves the ApiTag from the indexer for a given namespace and name.
func (s apiTagNamespaceLister) Get(name string) (*v1alpha1.ApiTag, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apitag"), name)
	}
	return obj.(*v1alpha1.ApiTag), nil
}