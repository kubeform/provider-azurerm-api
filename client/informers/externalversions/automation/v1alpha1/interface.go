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
	// Accounts returns a AccountInformer.
	Accounts() AccountInformer
	// Certificates returns a CertificateInformer.
	Certificates() CertificateInformer
	// Connections returns a ConnectionInformer.
	Connections() ConnectionInformer
	// ConnectionCertificates returns a ConnectionCertificateInformer.
	ConnectionCertificates() ConnectionCertificateInformer
	// ConnectionClassicCertificates returns a ConnectionClassicCertificateInformer.
	ConnectionClassicCertificates() ConnectionClassicCertificateInformer
	// ConnectionServicePrincipals returns a ConnectionServicePrincipalInformer.
	ConnectionServicePrincipals() ConnectionServicePrincipalInformer
	// Credentials returns a CredentialInformer.
	Credentials() CredentialInformer
	// DscConfigurations returns a DscConfigurationInformer.
	DscConfigurations() DscConfigurationInformer
	// DscNodeconfigurations returns a DscNodeconfigurationInformer.
	DscNodeconfigurations() DscNodeconfigurationInformer
	// JobSchedules returns a JobScheduleInformer.
	JobSchedules() JobScheduleInformer
	// Modules returns a ModuleInformer.
	Modules() ModuleInformer
	// Runbooks returns a RunbookInformer.
	Runbooks() RunbookInformer
	// Schedules returns a ScheduleInformer.
	Schedules() ScheduleInformer
	// VariableBools returns a VariableBoolInformer.
	VariableBools() VariableBoolInformer
	// VariableDatetimes returns a VariableDatetimeInformer.
	VariableDatetimes() VariableDatetimeInformer
	// VariableInts returns a VariableIntInformer.
	VariableInts() VariableIntInformer
	// VariableStrings returns a VariableStringInformer.
	VariableStrings() VariableStringInformer
	// Webhooks returns a WebhookInformer.
	Webhooks() WebhookInformer
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

// Accounts returns a AccountInformer.
func (v *version) Accounts() AccountInformer {
	return &accountInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Certificates returns a CertificateInformer.
func (v *version) Certificates() CertificateInformer {
	return &certificateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Connections returns a ConnectionInformer.
func (v *version) Connections() ConnectionInformer {
	return &connectionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ConnectionCertificates returns a ConnectionCertificateInformer.
func (v *version) ConnectionCertificates() ConnectionCertificateInformer {
	return &connectionCertificateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ConnectionClassicCertificates returns a ConnectionClassicCertificateInformer.
func (v *version) ConnectionClassicCertificates() ConnectionClassicCertificateInformer {
	return &connectionClassicCertificateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ConnectionServicePrincipals returns a ConnectionServicePrincipalInformer.
func (v *version) ConnectionServicePrincipals() ConnectionServicePrincipalInformer {
	return &connectionServicePrincipalInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Credentials returns a CredentialInformer.
func (v *version) Credentials() CredentialInformer {
	return &credentialInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DscConfigurations returns a DscConfigurationInformer.
func (v *version) DscConfigurations() DscConfigurationInformer {
	return &dscConfigurationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DscNodeconfigurations returns a DscNodeconfigurationInformer.
func (v *version) DscNodeconfigurations() DscNodeconfigurationInformer {
	return &dscNodeconfigurationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// JobSchedules returns a JobScheduleInformer.
func (v *version) JobSchedules() JobScheduleInformer {
	return &jobScheduleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Modules returns a ModuleInformer.
func (v *version) Modules() ModuleInformer {
	return &moduleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Runbooks returns a RunbookInformer.
func (v *version) Runbooks() RunbookInformer {
	return &runbookInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Schedules returns a ScheduleInformer.
func (v *version) Schedules() ScheduleInformer {
	return &scheduleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VariableBools returns a VariableBoolInformer.
func (v *version) VariableBools() VariableBoolInformer {
	return &variableBoolInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VariableDatetimes returns a VariableDatetimeInformer.
func (v *version) VariableDatetimes() VariableDatetimeInformer {
	return &variableDatetimeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VariableInts returns a VariableIntInformer.
func (v *version) VariableInts() VariableIntInformer {
	return &variableIntInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VariableStrings returns a VariableStringInformer.
func (v *version) VariableStrings() VariableStringInformer {
	return &variableStringInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Webhooks returns a WebhookInformer.
func (v *version) Webhooks() WebhookInformer {
	return &webhookInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
