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
	// AadDiagnosticSettings returns a AadDiagnosticSettingInformer.
	AadDiagnosticSettings() AadDiagnosticSettingInformer
	// ActionGroups returns a ActionGroupInformer.
	ActionGroups() ActionGroupInformer
	// ActionRuleActionGroups returns a ActionRuleActionGroupInformer.
	ActionRuleActionGroups() ActionRuleActionGroupInformer
	// ActionRuleSuppressions returns a ActionRuleSuppressionInformer.
	ActionRuleSuppressions() ActionRuleSuppressionInformer
	// ActivityLogAlerts returns a ActivityLogAlertInformer.
	ActivityLogAlerts() ActivityLogAlertInformer
	// AutoscaleSettings returns a AutoscaleSettingInformer.
	AutoscaleSettings() AutoscaleSettingInformer
	// DiagnosticSettings returns a DiagnosticSettingInformer.
	DiagnosticSettings() DiagnosticSettingInformer
	// LogProfiles returns a LogProfileInformer.
	LogProfiles() LogProfileInformer
	// MetricAlerts returns a MetricAlertInformer.
	MetricAlerts() MetricAlertInformer
	// PrivateLinkScopes returns a PrivateLinkScopeInformer.
	PrivateLinkScopes() PrivateLinkScopeInformer
	// PrivateLinkScopedServices returns a PrivateLinkScopedServiceInformer.
	PrivateLinkScopedServices() PrivateLinkScopedServiceInformer
	// ScheduledQueryRulesAlerts returns a ScheduledQueryRulesAlertInformer.
	ScheduledQueryRulesAlerts() ScheduledQueryRulesAlertInformer
	// ScheduledQueryRulesLogs returns a ScheduledQueryRulesLogInformer.
	ScheduledQueryRulesLogs() ScheduledQueryRulesLogInformer
	// SmartDetectorAlertRules returns a SmartDetectorAlertRuleInformer.
	SmartDetectorAlertRules() SmartDetectorAlertRuleInformer
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

// AadDiagnosticSettings returns a AadDiagnosticSettingInformer.
func (v *version) AadDiagnosticSettings() AadDiagnosticSettingInformer {
	return &aadDiagnosticSettingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ActionGroups returns a ActionGroupInformer.
func (v *version) ActionGroups() ActionGroupInformer {
	return &actionGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ActionRuleActionGroups returns a ActionRuleActionGroupInformer.
func (v *version) ActionRuleActionGroups() ActionRuleActionGroupInformer {
	return &actionRuleActionGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ActionRuleSuppressions returns a ActionRuleSuppressionInformer.
func (v *version) ActionRuleSuppressions() ActionRuleSuppressionInformer {
	return &actionRuleSuppressionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ActivityLogAlerts returns a ActivityLogAlertInformer.
func (v *version) ActivityLogAlerts() ActivityLogAlertInformer {
	return &activityLogAlertInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AutoscaleSettings returns a AutoscaleSettingInformer.
func (v *version) AutoscaleSettings() AutoscaleSettingInformer {
	return &autoscaleSettingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DiagnosticSettings returns a DiagnosticSettingInformer.
func (v *version) DiagnosticSettings() DiagnosticSettingInformer {
	return &diagnosticSettingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// LogProfiles returns a LogProfileInformer.
func (v *version) LogProfiles() LogProfileInformer {
	return &logProfileInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MetricAlerts returns a MetricAlertInformer.
func (v *version) MetricAlerts() MetricAlertInformer {
	return &metricAlertInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PrivateLinkScopes returns a PrivateLinkScopeInformer.
func (v *version) PrivateLinkScopes() PrivateLinkScopeInformer {
	return &privateLinkScopeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PrivateLinkScopedServices returns a PrivateLinkScopedServiceInformer.
func (v *version) PrivateLinkScopedServices() PrivateLinkScopedServiceInformer {
	return &privateLinkScopedServiceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ScheduledQueryRulesAlerts returns a ScheduledQueryRulesAlertInformer.
func (v *version) ScheduledQueryRulesAlerts() ScheduledQueryRulesAlertInformer {
	return &scheduledQueryRulesAlertInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ScheduledQueryRulesLogs returns a ScheduledQueryRulesLogInformer.
func (v *version) ScheduledQueryRulesLogs() ScheduledQueryRulesLogInformer {
	return &scheduledQueryRulesLogInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SmartDetectorAlertRules returns a SmartDetectorAlertRuleInformer.
func (v *version) SmartDetectorAlertRules() SmartDetectorAlertRuleInformer {
	return &smartDetectorAlertRuleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
