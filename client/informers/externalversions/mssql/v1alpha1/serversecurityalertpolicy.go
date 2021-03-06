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

	mssqlv1alpha1 "kubeform.dev/provider-azurerm-api/apis/mssql/v1alpha1"
	versioned "kubeform.dev/provider-azurerm-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-azurerm-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/listers/mssql/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServerSecurityAlertPolicyInformer provides access to a shared informer and lister for
// ServerSecurityAlertPolicies.
type ServerSecurityAlertPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ServerSecurityAlertPolicyLister
}

type serverSecurityAlertPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServerSecurityAlertPolicyInformer constructs a new informer for ServerSecurityAlertPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServerSecurityAlertPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServerSecurityAlertPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServerSecurityAlertPolicyInformer constructs a new informer for ServerSecurityAlertPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServerSecurityAlertPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MssqlV1alpha1().ServerSecurityAlertPolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MssqlV1alpha1().ServerSecurityAlertPolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&mssqlv1alpha1.ServerSecurityAlertPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *serverSecurityAlertPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServerSecurityAlertPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serverSecurityAlertPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&mssqlv1alpha1.ServerSecurityAlertPolicy{}, f.defaultInformer)
}

func (f *serverSecurityAlertPolicyInformer) Lister() v1alpha1.ServerSecurityAlertPolicyLister {
	return v1alpha1.NewServerSecurityAlertPolicyLister(f.Informer().GetIndexer())
}
