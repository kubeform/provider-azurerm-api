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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/media/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ContentKeyPolicyLister helps list ContentKeyPolicies.
// All objects returned here must be treated as read-only.
type ContentKeyPolicyLister interface {
	// List lists all ContentKeyPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ContentKeyPolicy, err error)
	// ContentKeyPolicies returns an object that can list and get ContentKeyPolicies.
	ContentKeyPolicies(namespace string) ContentKeyPolicyNamespaceLister
	ContentKeyPolicyListerExpansion
}

// contentKeyPolicyLister implements the ContentKeyPolicyLister interface.
type contentKeyPolicyLister struct {
	indexer cache.Indexer
}

// NewContentKeyPolicyLister returns a new ContentKeyPolicyLister.
func NewContentKeyPolicyLister(indexer cache.Indexer) ContentKeyPolicyLister {
	return &contentKeyPolicyLister{indexer: indexer}
}

// List lists all ContentKeyPolicies in the indexer.
func (s *contentKeyPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.ContentKeyPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ContentKeyPolicy))
	})
	return ret, err
}

// ContentKeyPolicies returns an object that can list and get ContentKeyPolicies.
func (s *contentKeyPolicyLister) ContentKeyPolicies(namespace string) ContentKeyPolicyNamespaceLister {
	return contentKeyPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ContentKeyPolicyNamespaceLister helps list and get ContentKeyPolicies.
// All objects returned here must be treated as read-only.
type ContentKeyPolicyNamespaceLister interface {
	// List lists all ContentKeyPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ContentKeyPolicy, err error)
	// Get retrieves the ContentKeyPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ContentKeyPolicy, error)
	ContentKeyPolicyNamespaceListerExpansion
}

// contentKeyPolicyNamespaceLister implements the ContentKeyPolicyNamespaceLister
// interface.
type contentKeyPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ContentKeyPolicies in the indexer for a given namespace.
func (s contentKeyPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ContentKeyPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ContentKeyPolicy))
	})
	return ret, err
}

// Get retrieves the ContentKeyPolicy from the indexer for a given namespace and name.
func (s contentKeyPolicyNamespaceLister) Get(name string) (*v1alpha1.ContentKeyPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("contentkeypolicy"), name)
	}
	return obj.(*v1alpha1.ContentKeyPolicy), nil
}