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

	sentinelv1alpha1 "kubeform.dev/provider-azurerm-api/apis/sentinel/v1alpha1"
	versioned "kubeform.dev/provider-azurerm-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-azurerm-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/listers/sentinel/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AlertRuleMachineLearningBehaviorAnalyticsInformer provides access to a shared informer and lister for
// AlertRuleMachineLearningBehaviorAnalyticses.
type AlertRuleMachineLearningBehaviorAnalyticsInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AlertRuleMachineLearningBehaviorAnalyticsLister
}

type alertRuleMachineLearningBehaviorAnalyticsInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAlertRuleMachineLearningBehaviorAnalyticsInformer constructs a new informer for AlertRuleMachineLearningBehaviorAnalytics type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAlertRuleMachineLearningBehaviorAnalyticsInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAlertRuleMachineLearningBehaviorAnalyticsInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAlertRuleMachineLearningBehaviorAnalyticsInformer constructs a new informer for AlertRuleMachineLearningBehaviorAnalytics type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAlertRuleMachineLearningBehaviorAnalyticsInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SentinelV1alpha1().AlertRuleMachineLearningBehaviorAnalyticses(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SentinelV1alpha1().AlertRuleMachineLearningBehaviorAnalyticses(namespace).Watch(context.TODO(), options)
			},
		},
		&sentinelv1alpha1.AlertRuleMachineLearningBehaviorAnalytics{},
		resyncPeriod,
		indexers,
	)
}

func (f *alertRuleMachineLearningBehaviorAnalyticsInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAlertRuleMachineLearningBehaviorAnalyticsInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *alertRuleMachineLearningBehaviorAnalyticsInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&sentinelv1alpha1.AlertRuleMachineLearningBehaviorAnalytics{}, f.defaultInformer)
}

func (f *alertRuleMachineLearningBehaviorAnalyticsInformer) Lister() v1alpha1.AlertRuleMachineLearningBehaviorAnalyticsLister {
	return v1alpha1.NewAlertRuleMachineLearningBehaviorAnalyticsLister(f.Informer().GetIndexer())
}
