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

	cognitivev1alpha1 "kubeform.dev/provider-azurerm-api/apis/cognitive/v1alpha1"
	versioned "kubeform.dev/provider-azurerm-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-azurerm-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/listers/cognitive/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AccountCustomerManagedKeyInformer provides access to a shared informer and lister for
// AccountCustomerManagedKeys.
type AccountCustomerManagedKeyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AccountCustomerManagedKeyLister
}

type accountCustomerManagedKeyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAccountCustomerManagedKeyInformer constructs a new informer for AccountCustomerManagedKey type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAccountCustomerManagedKeyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAccountCustomerManagedKeyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAccountCustomerManagedKeyInformer constructs a new informer for AccountCustomerManagedKey type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAccountCustomerManagedKeyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CognitiveV1alpha1().AccountCustomerManagedKeys(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CognitiveV1alpha1().AccountCustomerManagedKeys(namespace).Watch(context.TODO(), options)
			},
		},
		&cognitivev1alpha1.AccountCustomerManagedKey{},
		resyncPeriod,
		indexers,
	)
}

func (f *accountCustomerManagedKeyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAccountCustomerManagedKeyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *accountCustomerManagedKeyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cognitivev1alpha1.AccountCustomerManagedKey{}, f.defaultInformer)
}

func (f *accountCustomerManagedKeyInformer) Lister() v1alpha1.AccountCustomerManagedKeyLister {
	return v1alpha1.NewAccountCustomerManagedKeyLister(f.Informer().GetIndexer())
}
