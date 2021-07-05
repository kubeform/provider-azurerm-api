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

	privatev1alpha1 "kubeform.dev/provider-azurerm-api/apis/private/v1alpha1"
	versioned "kubeform.dev/provider-azurerm-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-azurerm-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/listers/private/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DnsMxRecordInformer provides access to a shared informer and lister for
// DnsMxRecords.
type DnsMxRecordInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.DnsMxRecordLister
}

type dnsMxRecordInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDnsMxRecordInformer constructs a new informer for DnsMxRecord type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDnsMxRecordInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDnsMxRecordInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDnsMxRecordInformer constructs a new informer for DnsMxRecord type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDnsMxRecordInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PrivateV1alpha1().DnsMxRecords(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PrivateV1alpha1().DnsMxRecords(namespace).Watch(context.TODO(), options)
			},
		},
		&privatev1alpha1.DnsMxRecord{},
		resyncPeriod,
		indexers,
	)
}

func (f *dnsMxRecordInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDnsMxRecordInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *dnsMxRecordInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&privatev1alpha1.DnsMxRecord{}, f.defaultInformer)
}

func (f *dnsMxRecordInformer) Lister() v1alpha1.DnsMxRecordLister {
	return v1alpha1.NewDnsMxRecordLister(f.Informer().GetIndexer())
}