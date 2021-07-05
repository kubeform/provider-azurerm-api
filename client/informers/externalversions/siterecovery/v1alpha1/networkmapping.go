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

	siterecoveryv1alpha1 "kubeform.dev/provider-azurerm-api/apis/siterecovery/v1alpha1"
	versioned "kubeform.dev/provider-azurerm-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-azurerm-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/listers/siterecovery/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NetworkMappingInformer provides access to a shared informer and lister for
// NetworkMappings.
type NetworkMappingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.NetworkMappingLister
}

type networkMappingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewNetworkMappingInformer constructs a new informer for NetworkMapping type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNetworkMappingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNetworkMappingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredNetworkMappingInformer constructs a new informer for NetworkMapping type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNetworkMappingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SiterecoveryV1alpha1().NetworkMappings(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SiterecoveryV1alpha1().NetworkMappings(namespace).Watch(context.TODO(), options)
			},
		},
		&siterecoveryv1alpha1.NetworkMapping{},
		resyncPeriod,
		indexers,
	)
}

func (f *networkMappingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNetworkMappingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *networkMappingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&siterecoveryv1alpha1.NetworkMapping{}, f.defaultInformer)
}

func (f *networkMappingInformer) Lister() v1alpha1.NetworkMappingLister {
	return v1alpha1.NewNetworkMappingLister(f.Informer().GetIndexer())
}