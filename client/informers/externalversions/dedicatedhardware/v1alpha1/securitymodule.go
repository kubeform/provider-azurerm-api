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

	dedicatedhardwarev1alpha1 "kubeform.dev/provider-azurerm-api/apis/dedicatedhardware/v1alpha1"
	versioned "kubeform.dev/provider-azurerm-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-azurerm-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/listers/dedicatedhardware/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SecurityModuleInformer provides access to a shared informer and lister for
// SecurityModules.
type SecurityModuleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.SecurityModuleLister
}

type securityModuleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSecurityModuleInformer constructs a new informer for SecurityModule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSecurityModuleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSecurityModuleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSecurityModuleInformer constructs a new informer for SecurityModule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSecurityModuleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DedicatedhardwareV1alpha1().SecurityModules(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DedicatedhardwareV1alpha1().SecurityModules(namespace).Watch(context.TODO(), options)
			},
		},
		&dedicatedhardwarev1alpha1.SecurityModule{},
		resyncPeriod,
		indexers,
	)
}

func (f *securityModuleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSecurityModuleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *securityModuleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&dedicatedhardwarev1alpha1.SecurityModule{}, f.defaultInformer)
}

func (f *securityModuleInformer) Lister() v1alpha1.SecurityModuleLister {
	return v1alpha1.NewSecurityModuleLister(f.Informer().GetIndexer())
}
