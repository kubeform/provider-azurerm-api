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
	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/bot/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ChannelDirectLineSpeechLister helps list ChannelDirectLineSpeeches.
// All objects returned here must be treated as read-only.
type ChannelDirectLineSpeechLister interface {
	// List lists all ChannelDirectLineSpeeches in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ChannelDirectLineSpeech, err error)
	// ChannelDirectLineSpeeches returns an object that can list and get ChannelDirectLineSpeeches.
	ChannelDirectLineSpeeches(namespace string) ChannelDirectLineSpeechNamespaceLister
	ChannelDirectLineSpeechListerExpansion
}

// channelDirectLineSpeechLister implements the ChannelDirectLineSpeechLister interface.
type channelDirectLineSpeechLister struct {
	indexer cache.Indexer
}

// NewChannelDirectLineSpeechLister returns a new ChannelDirectLineSpeechLister.
func NewChannelDirectLineSpeechLister(indexer cache.Indexer) ChannelDirectLineSpeechLister {
	return &channelDirectLineSpeechLister{indexer: indexer}
}

// List lists all ChannelDirectLineSpeeches in the indexer.
func (s *channelDirectLineSpeechLister) List(selector labels.Selector) (ret []*v1alpha1.ChannelDirectLineSpeech, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ChannelDirectLineSpeech))
	})
	return ret, err
}

// ChannelDirectLineSpeeches returns an object that can list and get ChannelDirectLineSpeeches.
func (s *channelDirectLineSpeechLister) ChannelDirectLineSpeeches(namespace string) ChannelDirectLineSpeechNamespaceLister {
	return channelDirectLineSpeechNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ChannelDirectLineSpeechNamespaceLister helps list and get ChannelDirectLineSpeeches.
// All objects returned here must be treated as read-only.
type ChannelDirectLineSpeechNamespaceLister interface {
	// List lists all ChannelDirectLineSpeeches in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ChannelDirectLineSpeech, err error)
	// Get retrieves the ChannelDirectLineSpeech from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ChannelDirectLineSpeech, error)
	ChannelDirectLineSpeechNamespaceListerExpansion
}

// channelDirectLineSpeechNamespaceLister implements the ChannelDirectLineSpeechNamespaceLister
// interface.
type channelDirectLineSpeechNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ChannelDirectLineSpeeches in the indexer for a given namespace.
func (s channelDirectLineSpeechNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ChannelDirectLineSpeech, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ChannelDirectLineSpeech))
	})
	return ret, err
}

// Get retrieves the ChannelDirectLineSpeech from the indexer for a given namespace and name.
func (s channelDirectLineSpeechNamespaceLister) Get(name string) (*v1alpha1.ChannelDirectLineSpeech, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("channeldirectlinespeech"), name)
	}
	return obj.(*v1alpha1.ChannelDirectLineSpeech), nil
}