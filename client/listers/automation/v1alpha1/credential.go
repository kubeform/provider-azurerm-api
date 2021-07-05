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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/automation/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CredentialLister helps list Credentials.
// All objects returned here must be treated as read-only.
type CredentialLister interface {
	// List lists all Credentials in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Credential, err error)
	// Credentials returns an object that can list and get Credentials.
	Credentials(namespace string) CredentialNamespaceLister
	CredentialListerExpansion
}

// credentialLister implements the CredentialLister interface.
type credentialLister struct {
	indexer cache.Indexer
}

// NewCredentialLister returns a new CredentialLister.
func NewCredentialLister(indexer cache.Indexer) CredentialLister {
	return &credentialLister{indexer: indexer}
}

// List lists all Credentials in the indexer.
func (s *credentialLister) List(selector labels.Selector) (ret []*v1alpha1.Credential, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Credential))
	})
	return ret, err
}

// Credentials returns an object that can list and get Credentials.
func (s *credentialLister) Credentials(namespace string) CredentialNamespaceLister {
	return credentialNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CredentialNamespaceLister helps list and get Credentials.
// All objects returned here must be treated as read-only.
type CredentialNamespaceLister interface {
	// List lists all Credentials in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Credential, err error)
	// Get retrieves the Credential from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Credential, error)
	CredentialNamespaceListerExpansion
}

// credentialNamespaceLister implements the CredentialNamespaceLister
// interface.
type credentialNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Credentials in the indexer for a given namespace.
func (s credentialNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Credential, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Credential))
	})
	return ret, err
}

// Get retrieves the Credential from the indexer for a given namespace and name.
func (s credentialNamespaceLister) Get(name string) (*v1alpha1.Credential, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("credential"), name)
	}
	return obj.(*v1alpha1.Credential), nil
}