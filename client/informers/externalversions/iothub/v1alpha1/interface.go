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
	internalinterfaces "kubeform.dev/provider-azurerm-api/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ConsumerGroups returns a ConsumerGroupInformer.
	ConsumerGroups() ConsumerGroupInformer
	// Dpses returns a DpsInformer.
	Dpses() DpsInformer
	// DpsCertificates returns a DpsCertificateInformer.
	DpsCertificates() DpsCertificateInformer
	// DpsSharedAccessPolicies returns a DpsSharedAccessPolicyInformer.
	DpsSharedAccessPolicies() DpsSharedAccessPolicyInformer
	// EndpointEventhubs returns a EndpointEventhubInformer.
	EndpointEventhubs() EndpointEventhubInformer
	// EndpointServicebusQueues returns a EndpointServicebusQueueInformer.
	EndpointServicebusQueues() EndpointServicebusQueueInformer
	// EndpointServicebusTopics returns a EndpointServicebusTopicInformer.
	EndpointServicebusTopics() EndpointServicebusTopicInformer
	// EndpointStorageContainers returns a EndpointStorageContainerInformer.
	EndpointStorageContainers() EndpointStorageContainerInformer
	// Enrichments returns a EnrichmentInformer.
	Enrichments() EnrichmentInformer
	// FallbackRoutes returns a FallbackRouteInformer.
	FallbackRoutes() FallbackRouteInformer
	// Iothubs returns a IothubInformer.
	Iothubs() IothubInformer
	// Routes returns a RouteInformer.
	Routes() RouteInformer
	// SharedAccessPolicies returns a SharedAccessPolicyInformer.
	SharedAccessPolicies() SharedAccessPolicyInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ConsumerGroups returns a ConsumerGroupInformer.
func (v *version) ConsumerGroups() ConsumerGroupInformer {
	return &consumerGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Dpses returns a DpsInformer.
func (v *version) Dpses() DpsInformer {
	return &dpsInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DpsCertificates returns a DpsCertificateInformer.
func (v *version) DpsCertificates() DpsCertificateInformer {
	return &dpsCertificateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DpsSharedAccessPolicies returns a DpsSharedAccessPolicyInformer.
func (v *version) DpsSharedAccessPolicies() DpsSharedAccessPolicyInformer {
	return &dpsSharedAccessPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EndpointEventhubs returns a EndpointEventhubInformer.
func (v *version) EndpointEventhubs() EndpointEventhubInformer {
	return &endpointEventhubInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EndpointServicebusQueues returns a EndpointServicebusQueueInformer.
func (v *version) EndpointServicebusQueues() EndpointServicebusQueueInformer {
	return &endpointServicebusQueueInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EndpointServicebusTopics returns a EndpointServicebusTopicInformer.
func (v *version) EndpointServicebusTopics() EndpointServicebusTopicInformer {
	return &endpointServicebusTopicInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EndpointStorageContainers returns a EndpointStorageContainerInformer.
func (v *version) EndpointStorageContainers() EndpointStorageContainerInformer {
	return &endpointStorageContainerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Enrichments returns a EnrichmentInformer.
func (v *version) Enrichments() EnrichmentInformer {
	return &enrichmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FallbackRoutes returns a FallbackRouteInformer.
func (v *version) FallbackRoutes() FallbackRouteInformer {
	return &fallbackRouteInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Iothubs returns a IothubInformer.
func (v *version) Iothubs() IothubInformer {
	return &iothubInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Routes returns a RouteInformer.
func (v *version) Routes() RouteInformer {
	return &routeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SharedAccessPolicies returns a SharedAccessPolicyInformer.
func (v *version) SharedAccessPolicies() SharedAccessPolicyInformer {
	return &sharedAccessPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
