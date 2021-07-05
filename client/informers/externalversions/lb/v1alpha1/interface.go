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
	// BackendAddressPools returns a BackendAddressPoolInformer.
	BackendAddressPools() BackendAddressPoolInformer
	// BackendAddressPoolAddresses returns a BackendAddressPoolAddressInformer.
	BackendAddressPoolAddresses() BackendAddressPoolAddressInformer
	// Lbs returns a LbInformer.
	Lbs() LbInformer
	// NatPools returns a NatPoolInformer.
	NatPools() NatPoolInformer
	// NatRules returns a NatRuleInformer.
	NatRules() NatRuleInformer
	// OutboundRules returns a OutboundRuleInformer.
	OutboundRules() OutboundRuleInformer
	// Probes returns a ProbeInformer.
	Probes() ProbeInformer
	// Rules returns a RuleInformer.
	Rules() RuleInformer
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

// BackendAddressPools returns a BackendAddressPoolInformer.
func (v *version) BackendAddressPools() BackendAddressPoolInformer {
	return &backendAddressPoolInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BackendAddressPoolAddresses returns a BackendAddressPoolAddressInformer.
func (v *version) BackendAddressPoolAddresses() BackendAddressPoolAddressInformer {
	return &backendAddressPoolAddressInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Lbs returns a LbInformer.
func (v *version) Lbs() LbInformer {
	return &lbInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NatPools returns a NatPoolInformer.
func (v *version) NatPools() NatPoolInformer {
	return &natPoolInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NatRules returns a NatRuleInformer.
func (v *version) NatRules() NatRuleInformer {
	return &natRuleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// OutboundRules returns a OutboundRuleInformer.
func (v *version) OutboundRules() OutboundRuleInformer {
	return &outboundRuleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Probes returns a ProbeInformer.
func (v *version) Probes() ProbeInformer {
	return &probeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Rules returns a RuleInformer.
func (v *version) Rules() RuleInformer {
	return &ruleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}