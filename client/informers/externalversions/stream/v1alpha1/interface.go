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
	// AnalyticsClusters returns a AnalyticsClusterInformer.
	AnalyticsClusters() AnalyticsClusterInformer
	// AnalyticsFunctionJavascriptUdves returns a AnalyticsFunctionJavascriptUdfInformer.
	AnalyticsFunctionJavascriptUdves() AnalyticsFunctionJavascriptUdfInformer
	// AnalyticsJobs returns a AnalyticsJobInformer.
	AnalyticsJobs() AnalyticsJobInformer
	// AnalyticsManagedPrivateEndpoints returns a AnalyticsManagedPrivateEndpointInformer.
	AnalyticsManagedPrivateEndpoints() AnalyticsManagedPrivateEndpointInformer
	// AnalyticsOutputBlobs returns a AnalyticsOutputBlobInformer.
	AnalyticsOutputBlobs() AnalyticsOutputBlobInformer
	// AnalyticsOutputEventhubs returns a AnalyticsOutputEventhubInformer.
	AnalyticsOutputEventhubs() AnalyticsOutputEventhubInformer
	// AnalyticsOutputFunctions returns a AnalyticsOutputFunctionInformer.
	AnalyticsOutputFunctions() AnalyticsOutputFunctionInformer
	// AnalyticsOutputMssqls returns a AnalyticsOutputMssqlInformer.
	AnalyticsOutputMssqls() AnalyticsOutputMssqlInformer
	// AnalyticsOutputServicebusQueues returns a AnalyticsOutputServicebusQueueInformer.
	AnalyticsOutputServicebusQueues() AnalyticsOutputServicebusQueueInformer
	// AnalyticsOutputServicebusTopics returns a AnalyticsOutputServicebusTopicInformer.
	AnalyticsOutputServicebusTopics() AnalyticsOutputServicebusTopicInformer
	// AnalyticsOutputSynapses returns a AnalyticsOutputSynapseInformer.
	AnalyticsOutputSynapses() AnalyticsOutputSynapseInformer
	// AnalyticsOutputTables returns a AnalyticsOutputTableInformer.
	AnalyticsOutputTables() AnalyticsOutputTableInformer
	// AnalyticsReferenceInputBlobs returns a AnalyticsReferenceInputBlobInformer.
	AnalyticsReferenceInputBlobs() AnalyticsReferenceInputBlobInformer
	// AnalyticsReferenceInputMssqls returns a AnalyticsReferenceInputMssqlInformer.
	AnalyticsReferenceInputMssqls() AnalyticsReferenceInputMssqlInformer
	// AnalyticsStreamInputBlobs returns a AnalyticsStreamInputBlobInformer.
	AnalyticsStreamInputBlobs() AnalyticsStreamInputBlobInformer
	// AnalyticsStreamInputEventhubs returns a AnalyticsStreamInputEventhubInformer.
	AnalyticsStreamInputEventhubs() AnalyticsStreamInputEventhubInformer
	// AnalyticsStreamInputIothubs returns a AnalyticsStreamInputIothubInformer.
	AnalyticsStreamInputIothubs() AnalyticsStreamInputIothubInformer
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

// AnalyticsClusters returns a AnalyticsClusterInformer.
func (v *version) AnalyticsClusters() AnalyticsClusterInformer {
	return &analyticsClusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AnalyticsFunctionJavascriptUdves returns a AnalyticsFunctionJavascriptUdfInformer.
func (v *version) AnalyticsFunctionJavascriptUdves() AnalyticsFunctionJavascriptUdfInformer {
	return &analyticsFunctionJavascriptUdfInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AnalyticsJobs returns a AnalyticsJobInformer.
func (v *version) AnalyticsJobs() AnalyticsJobInformer {
	return &analyticsJobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AnalyticsManagedPrivateEndpoints returns a AnalyticsManagedPrivateEndpointInformer.
func (v *version) AnalyticsManagedPrivateEndpoints() AnalyticsManagedPrivateEndpointInformer {
	return &analyticsManagedPrivateEndpointInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AnalyticsOutputBlobs returns a AnalyticsOutputBlobInformer.
func (v *version) AnalyticsOutputBlobs() AnalyticsOutputBlobInformer {
	return &analyticsOutputBlobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AnalyticsOutputEventhubs returns a AnalyticsOutputEventhubInformer.
func (v *version) AnalyticsOutputEventhubs() AnalyticsOutputEventhubInformer {
	return &analyticsOutputEventhubInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AnalyticsOutputFunctions returns a AnalyticsOutputFunctionInformer.
func (v *version) AnalyticsOutputFunctions() AnalyticsOutputFunctionInformer {
	return &analyticsOutputFunctionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AnalyticsOutputMssqls returns a AnalyticsOutputMssqlInformer.
func (v *version) AnalyticsOutputMssqls() AnalyticsOutputMssqlInformer {
	return &analyticsOutputMssqlInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AnalyticsOutputServicebusQueues returns a AnalyticsOutputServicebusQueueInformer.
func (v *version) AnalyticsOutputServicebusQueues() AnalyticsOutputServicebusQueueInformer {
	return &analyticsOutputServicebusQueueInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AnalyticsOutputServicebusTopics returns a AnalyticsOutputServicebusTopicInformer.
func (v *version) AnalyticsOutputServicebusTopics() AnalyticsOutputServicebusTopicInformer {
	return &analyticsOutputServicebusTopicInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AnalyticsOutputSynapses returns a AnalyticsOutputSynapseInformer.
func (v *version) AnalyticsOutputSynapses() AnalyticsOutputSynapseInformer {
	return &analyticsOutputSynapseInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AnalyticsOutputTables returns a AnalyticsOutputTableInformer.
func (v *version) AnalyticsOutputTables() AnalyticsOutputTableInformer {
	return &analyticsOutputTableInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AnalyticsReferenceInputBlobs returns a AnalyticsReferenceInputBlobInformer.
func (v *version) AnalyticsReferenceInputBlobs() AnalyticsReferenceInputBlobInformer {
	return &analyticsReferenceInputBlobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AnalyticsReferenceInputMssqls returns a AnalyticsReferenceInputMssqlInformer.
func (v *version) AnalyticsReferenceInputMssqls() AnalyticsReferenceInputMssqlInformer {
	return &analyticsReferenceInputMssqlInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AnalyticsStreamInputBlobs returns a AnalyticsStreamInputBlobInformer.
func (v *version) AnalyticsStreamInputBlobs() AnalyticsStreamInputBlobInformer {
	return &analyticsStreamInputBlobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AnalyticsStreamInputEventhubs returns a AnalyticsStreamInputEventhubInformer.
func (v *version) AnalyticsStreamInputEventhubs() AnalyticsStreamInputEventhubInformer {
	return &analyticsStreamInputEventhubInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AnalyticsStreamInputIothubs returns a AnalyticsStreamInputIothubInformer.
func (v *version) AnalyticsStreamInputIothubs() AnalyticsStreamInputIothubInformer {
	return &analyticsStreamInputIothubInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
