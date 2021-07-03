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

	loganalyticsv1alpha1 "kubeform.dev/provider-azurerm-api/apis/loganalytics/v1alpha1"
	versioned "kubeform.dev/provider-azurerm-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-azurerm-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/listers/loganalytics/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterCustomerManagedKeyInformer provides access to a shared informer and lister for
// ClusterCustomerManagedKeys.
type ClusterCustomerManagedKeyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterCustomerManagedKeyLister
}

type clusterCustomerManagedKeyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewClusterCustomerManagedKeyInformer constructs a new informer for ClusterCustomerManagedKey type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterCustomerManagedKeyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterCustomerManagedKeyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredClusterCustomerManagedKeyInformer constructs a new informer for ClusterCustomerManagedKey type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterCustomerManagedKeyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LoganalyticsV1alpha1().ClusterCustomerManagedKeys(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LoganalyticsV1alpha1().ClusterCustomerManagedKeys(namespace).Watch(context.TODO(), options)
			},
		},
		&loganalyticsv1alpha1.ClusterCustomerManagedKey{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterCustomerManagedKeyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterCustomerManagedKeyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterCustomerManagedKeyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&loganalyticsv1alpha1.ClusterCustomerManagedKey{}, f.defaultInformer)
}

func (f *clusterCustomerManagedKeyInformer) Lister() v1alpha1.ClusterCustomerManagedKeyLister {
	return v1alpha1.NewClusterCustomerManagedKeyLister(f.Informer().GetIndexer())
}
