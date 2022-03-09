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
	// Databases returns a DatabaseInformer.
	Databases() DatabaseInformer
	// DatabaseExtendedAuditingPolicies returns a DatabaseExtendedAuditingPolicyInformer.
	DatabaseExtendedAuditingPolicies() DatabaseExtendedAuditingPolicyInformer
	// DatabaseVulnerabilityAssessmentRuleBaselines returns a DatabaseVulnerabilityAssessmentRuleBaselineInformer.
	DatabaseVulnerabilityAssessmentRuleBaselines() DatabaseVulnerabilityAssessmentRuleBaselineInformer
	// Elasticpools returns a ElasticpoolInformer.
	Elasticpools() ElasticpoolInformer
	// FailoverGroups returns a FailoverGroupInformer.
	FailoverGroups() FailoverGroupInformer
	// FirewallRules returns a FirewallRuleInformer.
	FirewallRules() FirewallRuleInformer
	// JobAgents returns a JobAgentInformer.
	JobAgents() JobAgentInformer
	// JobCredentials returns a JobCredentialInformer.
	JobCredentials() JobCredentialInformer
	// ManagedDatabases returns a ManagedDatabaseInformer.
	ManagedDatabases() ManagedDatabaseInformer
	// ManagedInstances returns a ManagedInstanceInformer.
	ManagedInstances() ManagedInstanceInformer
	// ManagedInstanceActiveDirectoryAdministrators returns a ManagedInstanceActiveDirectoryAdministratorInformer.
	ManagedInstanceActiveDirectoryAdministrators() ManagedInstanceActiveDirectoryAdministratorInformer
	// ManagedInstanceFailoverGroups returns a ManagedInstanceFailoverGroupInformer.
	ManagedInstanceFailoverGroups() ManagedInstanceFailoverGroupInformer
	// OutboundFirewallRules returns a OutboundFirewallRuleInformer.
	OutboundFirewallRules() OutboundFirewallRuleInformer
	// Servers returns a ServerInformer.
	Servers() ServerInformer
	// ServerExtendedAuditingPolicies returns a ServerExtendedAuditingPolicyInformer.
	ServerExtendedAuditingPolicies() ServerExtendedAuditingPolicyInformer
	// ServerSecurityAlertPolicies returns a ServerSecurityAlertPolicyInformer.
	ServerSecurityAlertPolicies() ServerSecurityAlertPolicyInformer
	// ServerTransparentDataEncryptions returns a ServerTransparentDataEncryptionInformer.
	ServerTransparentDataEncryptions() ServerTransparentDataEncryptionInformer
	// ServerVulnerabilityAssessments returns a ServerVulnerabilityAssessmentInformer.
	ServerVulnerabilityAssessments() ServerVulnerabilityAssessmentInformer
	// VirtualMachines returns a VirtualMachineInformer.
	VirtualMachines() VirtualMachineInformer
	// VirtualNetworkRules returns a VirtualNetworkRuleInformer.
	VirtualNetworkRules() VirtualNetworkRuleInformer
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

// Databases returns a DatabaseInformer.
func (v *version) Databases() DatabaseInformer {
	return &databaseInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DatabaseExtendedAuditingPolicies returns a DatabaseExtendedAuditingPolicyInformer.
func (v *version) DatabaseExtendedAuditingPolicies() DatabaseExtendedAuditingPolicyInformer {
	return &databaseExtendedAuditingPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DatabaseVulnerabilityAssessmentRuleBaselines returns a DatabaseVulnerabilityAssessmentRuleBaselineInformer.
func (v *version) DatabaseVulnerabilityAssessmentRuleBaselines() DatabaseVulnerabilityAssessmentRuleBaselineInformer {
	return &databaseVulnerabilityAssessmentRuleBaselineInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Elasticpools returns a ElasticpoolInformer.
func (v *version) Elasticpools() ElasticpoolInformer {
	return &elasticpoolInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FailoverGroups returns a FailoverGroupInformer.
func (v *version) FailoverGroups() FailoverGroupInformer {
	return &failoverGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FirewallRules returns a FirewallRuleInformer.
func (v *version) FirewallRules() FirewallRuleInformer {
	return &firewallRuleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// JobAgents returns a JobAgentInformer.
func (v *version) JobAgents() JobAgentInformer {
	return &jobAgentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// JobCredentials returns a JobCredentialInformer.
func (v *version) JobCredentials() JobCredentialInformer {
	return &jobCredentialInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagedDatabases returns a ManagedDatabaseInformer.
func (v *version) ManagedDatabases() ManagedDatabaseInformer {
	return &managedDatabaseInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagedInstances returns a ManagedInstanceInformer.
func (v *version) ManagedInstances() ManagedInstanceInformer {
	return &managedInstanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagedInstanceActiveDirectoryAdministrators returns a ManagedInstanceActiveDirectoryAdministratorInformer.
func (v *version) ManagedInstanceActiveDirectoryAdministrators() ManagedInstanceActiveDirectoryAdministratorInformer {
	return &managedInstanceActiveDirectoryAdministratorInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagedInstanceFailoverGroups returns a ManagedInstanceFailoverGroupInformer.
func (v *version) ManagedInstanceFailoverGroups() ManagedInstanceFailoverGroupInformer {
	return &managedInstanceFailoverGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// OutboundFirewallRules returns a OutboundFirewallRuleInformer.
func (v *version) OutboundFirewallRules() OutboundFirewallRuleInformer {
	return &outboundFirewallRuleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Servers returns a ServerInformer.
func (v *version) Servers() ServerInformer {
	return &serverInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServerExtendedAuditingPolicies returns a ServerExtendedAuditingPolicyInformer.
func (v *version) ServerExtendedAuditingPolicies() ServerExtendedAuditingPolicyInformer {
	return &serverExtendedAuditingPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServerSecurityAlertPolicies returns a ServerSecurityAlertPolicyInformer.
func (v *version) ServerSecurityAlertPolicies() ServerSecurityAlertPolicyInformer {
	return &serverSecurityAlertPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServerTransparentDataEncryptions returns a ServerTransparentDataEncryptionInformer.
func (v *version) ServerTransparentDataEncryptions() ServerTransparentDataEncryptionInformer {
	return &serverTransparentDataEncryptionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServerVulnerabilityAssessments returns a ServerVulnerabilityAssessmentInformer.
func (v *version) ServerVulnerabilityAssessments() ServerVulnerabilityAssessmentInformer {
	return &serverVulnerabilityAssessmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VirtualMachines returns a VirtualMachineInformer.
func (v *version) VirtualMachines() VirtualMachineInformer {
	return &virtualMachineInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VirtualNetworkRules returns a VirtualNetworkRuleInformer.
func (v *version) VirtualNetworkRules() VirtualNetworkRuleInformer {
	return &virtualNetworkRuleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
