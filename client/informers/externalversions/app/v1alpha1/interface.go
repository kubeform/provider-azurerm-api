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
	// Configurations returns a ConfigurationInformer.
	Configurations() ConfigurationInformer
	// Services returns a ServiceInformer.
	Services() ServiceInformer
	// ServiceActiveSlots returns a ServiceActiveSlotInformer.
	ServiceActiveSlots() ServiceActiveSlotInformer
	// ServiceCertificates returns a ServiceCertificateInformer.
	ServiceCertificates() ServiceCertificateInformer
	// ServiceCertificateBindings returns a ServiceCertificateBindingInformer.
	ServiceCertificateBindings() ServiceCertificateBindingInformer
	// ServiceCertificateOrders returns a ServiceCertificateOrderInformer.
	ServiceCertificateOrders() ServiceCertificateOrderInformer
	// ServiceCustomHostnameBindings returns a ServiceCustomHostnameBindingInformer.
	ServiceCustomHostnameBindings() ServiceCustomHostnameBindingInformer
	// ServiceEnvironments returns a ServiceEnvironmentInformer.
	ServiceEnvironments() ServiceEnvironmentInformer
	// ServiceEnvironmentV3s returns a ServiceEnvironmentV3Informer.
	ServiceEnvironmentV3s() ServiceEnvironmentV3Informer
	// ServiceHybridConnections returns a ServiceHybridConnectionInformer.
	ServiceHybridConnections() ServiceHybridConnectionInformer
	// ServiceManagedCertificates returns a ServiceManagedCertificateInformer.
	ServiceManagedCertificates() ServiceManagedCertificateInformer
	// ServicePlans returns a ServicePlanInformer.
	ServicePlans() ServicePlanInformer
	// ServiceSlots returns a ServiceSlotInformer.
	ServiceSlots() ServiceSlotInformer
	// ServiceSlotVirtualNetworkSwiftConnections returns a ServiceSlotVirtualNetworkSwiftConnectionInformer.
	ServiceSlotVirtualNetworkSwiftConnections() ServiceSlotVirtualNetworkSwiftConnectionInformer
	// ServiceSourceControlTokens returns a ServiceSourceControlTokenInformer.
	ServiceSourceControlTokens() ServiceSourceControlTokenInformer
	// ServiceVirtualNetworkSwiftConnections returns a ServiceVirtualNetworkSwiftConnectionInformer.
	ServiceVirtualNetworkSwiftConnections() ServiceVirtualNetworkSwiftConnectionInformer
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

// Configurations returns a ConfigurationInformer.
func (v *version) Configurations() ConfigurationInformer {
	return &configurationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Services returns a ServiceInformer.
func (v *version) Services() ServiceInformer {
	return &serviceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceActiveSlots returns a ServiceActiveSlotInformer.
func (v *version) ServiceActiveSlots() ServiceActiveSlotInformer {
	return &serviceActiveSlotInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceCertificates returns a ServiceCertificateInformer.
func (v *version) ServiceCertificates() ServiceCertificateInformer {
	return &serviceCertificateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceCertificateBindings returns a ServiceCertificateBindingInformer.
func (v *version) ServiceCertificateBindings() ServiceCertificateBindingInformer {
	return &serviceCertificateBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceCertificateOrders returns a ServiceCertificateOrderInformer.
func (v *version) ServiceCertificateOrders() ServiceCertificateOrderInformer {
	return &serviceCertificateOrderInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceCustomHostnameBindings returns a ServiceCustomHostnameBindingInformer.
func (v *version) ServiceCustomHostnameBindings() ServiceCustomHostnameBindingInformer {
	return &serviceCustomHostnameBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceEnvironments returns a ServiceEnvironmentInformer.
func (v *version) ServiceEnvironments() ServiceEnvironmentInformer {
	return &serviceEnvironmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceEnvironmentV3s returns a ServiceEnvironmentV3Informer.
func (v *version) ServiceEnvironmentV3s() ServiceEnvironmentV3Informer {
	return &serviceEnvironmentV3Informer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceHybridConnections returns a ServiceHybridConnectionInformer.
func (v *version) ServiceHybridConnections() ServiceHybridConnectionInformer {
	return &serviceHybridConnectionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceManagedCertificates returns a ServiceManagedCertificateInformer.
func (v *version) ServiceManagedCertificates() ServiceManagedCertificateInformer {
	return &serviceManagedCertificateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServicePlans returns a ServicePlanInformer.
func (v *version) ServicePlans() ServicePlanInformer {
	return &servicePlanInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceSlots returns a ServiceSlotInformer.
func (v *version) ServiceSlots() ServiceSlotInformer {
	return &serviceSlotInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceSlotVirtualNetworkSwiftConnections returns a ServiceSlotVirtualNetworkSwiftConnectionInformer.
func (v *version) ServiceSlotVirtualNetworkSwiftConnections() ServiceSlotVirtualNetworkSwiftConnectionInformer {
	return &serviceSlotVirtualNetworkSwiftConnectionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceSourceControlTokens returns a ServiceSourceControlTokenInformer.
func (v *version) ServiceSourceControlTokens() ServiceSourceControlTokenInformer {
	return &serviceSourceControlTokenInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceVirtualNetworkSwiftConnections returns a ServiceVirtualNetworkSwiftConnectionInformer.
func (v *version) ServiceVirtualNetworkSwiftConnections() ServiceVirtualNetworkSwiftConnectionInformer {
	return &serviceVirtualNetworkSwiftConnectionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
