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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/packet/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CaptureLister helps list Captures.
// All objects returned here must be treated as read-only.
type CaptureLister interface {
	// List lists all Captures in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Capture, err error)
	// Captures returns an object that can list and get Captures.
	Captures(namespace string) CaptureNamespaceLister
	CaptureListerExpansion
}

// captureLister implements the CaptureLister interface.
type captureLister struct {
	indexer cache.Indexer
}

// NewCaptureLister returns a new CaptureLister.
func NewCaptureLister(indexer cache.Indexer) CaptureLister {
	return &captureLister{indexer: indexer}
}

// List lists all Captures in the indexer.
func (s *captureLister) List(selector labels.Selector) (ret []*v1alpha1.Capture, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Capture))
	})
	return ret, err
}

// Captures returns an object that can list and get Captures.
func (s *captureLister) Captures(namespace string) CaptureNamespaceLister {
	return captureNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CaptureNamespaceLister helps list and get Captures.
// All objects returned here must be treated as read-only.
type CaptureNamespaceLister interface {
	// List lists all Captures in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Capture, err error)
	// Get retrieves the Capture from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Capture, error)
	CaptureNamespaceListerExpansion
}

// captureNamespaceLister implements the CaptureNamespaceLister
// interface.
type captureNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Captures in the indexer for a given namespace.
func (s captureNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Capture, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Capture))
	})
	return ret, err
}

// Get retrieves the Capture from the indexer for a given namespace and name.
func (s captureNamespaceLister) Get(name string) (*v1alpha1.Capture, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("capture"), name)
	}
	return obj.(*v1alpha1.Capture), nil
}
