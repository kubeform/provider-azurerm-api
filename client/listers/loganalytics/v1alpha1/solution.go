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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/loganalytics/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SolutionLister helps list Solutions.
// All objects returned here must be treated as read-only.
type SolutionLister interface {
	// List lists all Solutions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Solution, err error)
	// Solutions returns an object that can list and get Solutions.
	Solutions(namespace string) SolutionNamespaceLister
	SolutionListerExpansion
}

// solutionLister implements the SolutionLister interface.
type solutionLister struct {
	indexer cache.Indexer
}

// NewSolutionLister returns a new SolutionLister.
func NewSolutionLister(indexer cache.Indexer) SolutionLister {
	return &solutionLister{indexer: indexer}
}

// List lists all Solutions in the indexer.
func (s *solutionLister) List(selector labels.Selector) (ret []*v1alpha1.Solution, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Solution))
	})
	return ret, err
}

// Solutions returns an object that can list and get Solutions.
func (s *solutionLister) Solutions(namespace string) SolutionNamespaceLister {
	return solutionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SolutionNamespaceLister helps list and get Solutions.
// All objects returned here must be treated as read-only.
type SolutionNamespaceLister interface {
	// List lists all Solutions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Solution, err error)
	// Get retrieves the Solution from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Solution, error)
	SolutionNamespaceListerExpansion
}

// solutionNamespaceLister implements the SolutionNamespaceLister
// interface.
type solutionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Solutions in the indexer for a given namespace.
func (s solutionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Solution, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Solution))
	})
	return ret, err
}

// Get retrieves the Solution from the indexer for a given namespace and name.
func (s solutionNamespaceLister) Get(name string) (*v1alpha1.Solution, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("solution"), name)
	}
	return obj.(*v1alpha1.Solution), nil
}
