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

// AuthorizationServerLister helps list AuthorizationServers.
// All objects returned here must be treated as read-only.
type AuthorizationServerLister interface {
	// List lists all AuthorizationServers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AuthorizationServer, err error)
	// AuthorizationServers returns an object that can list and get AuthorizationServers.
	AuthorizationServers(namespace string) AuthorizationServerNamespaceLister
	AuthorizationServerListerExpansion
}

// authorizationServerLister implements the AuthorizationServerLister interface.
type authorizationServerLister struct {
	indexer cache.Indexer
}

// NewAuthorizationServerLister returns a new AuthorizationServerLister.
func NewAuthorizationServerLister(indexer cache.Indexer) AuthorizationServerLister {
	return &authorizationServerLister{indexer: indexer}
}

// List lists all AuthorizationServers in the indexer.
func (s *authorizationServerLister) List(selector labels.Selector) (ret []*v1alpha1.AuthorizationServer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AuthorizationServer))
	})
	return ret, err
}

// AuthorizationServers returns an object that can list and get AuthorizationServers.
func (s *authorizationServerLister) AuthorizationServers(namespace string) AuthorizationServerNamespaceLister {
	return authorizationServerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AuthorizationServerNamespaceLister helps list and get AuthorizationServers.
// All objects returned here must be treated as read-only.
type AuthorizationServerNamespaceLister interface {
	// List lists all AuthorizationServers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AuthorizationServer, err error)
	// Get retrieves the AuthorizationServer from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AuthorizationServer, error)
	AuthorizationServerNamespaceListerExpansion
}

// authorizationServerNamespaceLister implements the AuthorizationServerNamespaceLister
// interface.
type authorizationServerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AuthorizationServers in the indexer for a given namespace.
func (s authorizationServerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AuthorizationServer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AuthorizationServer))
	})
	return ret, err
}

// Get retrieves the AuthorizationServer from the indexer for a given namespace and name.
func (s authorizationServerNamespaceLister) Get(name string) (*v1alpha1.AuthorizationServer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("authorizationserver"), name)
	}
	return obj.(*v1alpha1.AuthorizationServer), nil
}
