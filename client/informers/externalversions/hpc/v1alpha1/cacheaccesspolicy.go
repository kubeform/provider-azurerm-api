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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	hpcv1alpha1 "kubeform.dev/provider-azurerm-api/apis/hpc/v1alpha1"
	versioned "kubeform.dev/provider-azurerm-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-azurerm-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/listers/hpc/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CacheAccessPolicyInformer provides access to a shared informer and lister for
// CacheAccessPolicies.
type CacheAccessPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CacheAccessPolicyLister
}

type cacheAccessPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCacheAccessPolicyInformer constructs a new informer for CacheAccessPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCacheAccessPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCacheAccessPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCacheAccessPolicyInformer constructs a new informer for CacheAccessPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCacheAccessPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HpcV1alpha1().CacheAccessPolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HpcV1alpha1().CacheAccessPolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&hpcv1alpha1.CacheAccessPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *cacheAccessPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCacheAccessPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cacheAccessPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&hpcv1alpha1.CacheAccessPolicy{}, f.defaultInformer)
}

func (f *cacheAccessPolicyInformer) Lister() v1alpha1.CacheAccessPolicyLister {
	return v1alpha1.NewCacheAccessPolicyLister(f.Informer().GetIndexer())
}
