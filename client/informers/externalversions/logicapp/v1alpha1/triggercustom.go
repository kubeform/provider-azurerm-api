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

	logicappv1alpha1 "kubeform.dev/provider-azurerm-api/apis/logicapp/v1alpha1"
	versioned "kubeform.dev/provider-azurerm-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-azurerm-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/listers/logicapp/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TriggerCustomInformer provides access to a shared informer and lister for
// TriggerCustoms.
type TriggerCustomInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TriggerCustomLister
}

type triggerCustomInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTriggerCustomInformer constructs a new informer for TriggerCustom type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTriggerCustomInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTriggerCustomInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTriggerCustomInformer constructs a new informer for TriggerCustom type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTriggerCustomInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LogicappV1alpha1().TriggerCustoms(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LogicappV1alpha1().TriggerCustoms(namespace).Watch(context.TODO(), options)
			},
		},
		&logicappv1alpha1.TriggerCustom{},
		resyncPeriod,
		indexers,
	)
}

func (f *triggerCustomInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTriggerCustomInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *triggerCustomInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&logicappv1alpha1.TriggerCustom{}, f.defaultInformer)
}

func (f *triggerCustomInformer) Lister() v1alpha1.TriggerCustomLister {
	return v1alpha1.NewTriggerCustomLister(f.Informer().GetIndexer())
}
