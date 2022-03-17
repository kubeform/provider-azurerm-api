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

	datav1alpha1 "kubeform.dev/provider-azurerm-api/apis/data/v1alpha1"
	versioned "kubeform.dev/provider-azurerm-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-azurerm-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/listers/data/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FactoryTriggerTumblingWindowInformer provides access to a shared informer and lister for
// FactoryTriggerTumblingWindows.
type FactoryTriggerTumblingWindowInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FactoryTriggerTumblingWindowLister
}

type factoryTriggerTumblingWindowInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFactoryTriggerTumblingWindowInformer constructs a new informer for FactoryTriggerTumblingWindow type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFactoryTriggerTumblingWindowInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFactoryTriggerTumblingWindowInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFactoryTriggerTumblingWindowInformer constructs a new informer for FactoryTriggerTumblingWindow type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFactoryTriggerTumblingWindowInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DataV1alpha1().FactoryTriggerTumblingWindows(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DataV1alpha1().FactoryTriggerTumblingWindows(namespace).Watch(context.TODO(), options)
			},
		},
		&datav1alpha1.FactoryTriggerTumblingWindow{},
		resyncPeriod,
		indexers,
	)
}

func (f *factoryTriggerTumblingWindowInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFactoryTriggerTumblingWindowInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *factoryTriggerTumblingWindowInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&datav1alpha1.FactoryTriggerTumblingWindow{}, f.defaultInformer)
}

func (f *factoryTriggerTumblingWindowInformer) Lister() v1alpha1.FactoryTriggerTumblingWindowLister {
	return v1alpha1.NewFactoryTriggerTumblingWindowLister(f.Informer().GetIndexer())
}