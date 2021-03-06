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
	// Apis returns a ApiInformer.
	Apis() ApiInformer
	// ApiDiagnostics returns a ApiDiagnosticInformer.
	ApiDiagnostics() ApiDiagnosticInformer
	// ApiManagements returns a ApiManagementInformer.
	ApiManagements() ApiManagementInformer
	// ApiOperations returns a ApiOperationInformer.
	ApiOperations() ApiOperationInformer
	// ApiOperationPolicies returns a ApiOperationPolicyInformer.
	ApiOperationPolicies() ApiOperationPolicyInformer
	// ApiOperationTags returns a ApiOperationTagInformer.
	ApiOperationTags() ApiOperationTagInformer
	// ApiPolicies returns a ApiPolicyInformer.
	ApiPolicies() ApiPolicyInformer
	// ApiReleases returns a ApiReleaseInformer.
	ApiReleases() ApiReleaseInformer
	// ApiSchemas returns a ApiSchemaInformer.
	ApiSchemas() ApiSchemaInformer
	// ApiTags returns a ApiTagInformer.
	ApiTags() ApiTagInformer
	// ApiVersionSets returns a ApiVersionSetInformer.
	ApiVersionSets() ApiVersionSetInformer
	// AuthorizationServers returns a AuthorizationServerInformer.
	AuthorizationServers() AuthorizationServerInformer
	// Backends returns a BackendInformer.
	Backends() BackendInformer
	// Certificates returns a CertificateInformer.
	Certificates() CertificateInformer
	// CustomDomains returns a CustomDomainInformer.
	CustomDomains() CustomDomainInformer
	// Diagnostics returns a DiagnosticInformer.
	Diagnostics() DiagnosticInformer
	// EmailTemplates returns a EmailTemplateInformer.
	EmailTemplates() EmailTemplateInformer
	// Gateways returns a GatewayInformer.
	Gateways() GatewayInformer
	// GatewayAPIs returns a GatewayAPIInformer.
	GatewayAPIs() GatewayAPIInformer
	// Groups returns a GroupInformer.
	Groups() GroupInformer
	// GroupUsers returns a GroupUserInformer.
	GroupUsers() GroupUserInformer
	// IdentityProviderAads returns a IdentityProviderAadInformer.
	IdentityProviderAads() IdentityProviderAadInformer
	// IdentityProviderAadb2cs returns a IdentityProviderAadb2cInformer.
	IdentityProviderAadb2cs() IdentityProviderAadb2cInformer
	// IdentityProviderFacebooks returns a IdentityProviderFacebookInformer.
	IdentityProviderFacebooks() IdentityProviderFacebookInformer
	// IdentityProviderGoogles returns a IdentityProviderGoogleInformer.
	IdentityProviderGoogles() IdentityProviderGoogleInformer
	// IdentityProviderMicrosofts returns a IdentityProviderMicrosoftInformer.
	IdentityProviderMicrosofts() IdentityProviderMicrosoftInformer
	// IdentityProviderTwitters returns a IdentityProviderTwitterInformer.
	IdentityProviderTwitters() IdentityProviderTwitterInformer
	// Loggers returns a LoggerInformer.
	Loggers() LoggerInformer
	// NamedValues returns a NamedValueInformer.
	NamedValues() NamedValueInformer
	// NotificationRecipientEmails returns a NotificationRecipientEmailInformer.
	NotificationRecipientEmails() NotificationRecipientEmailInformer
	// NotificationRecipientUsers returns a NotificationRecipientUserInformer.
	NotificationRecipientUsers() NotificationRecipientUserInformer
	// OpenidConnectProviders returns a OpenidConnectProviderInformer.
	OpenidConnectProviders() OpenidConnectProviderInformer
	// Policies returns a PolicyInformer.
	Policies() PolicyInformer
	// Products returns a ProductInformer.
	Products() ProductInformer
	// ProductAPIs returns a ProductAPIInformer.
	ProductAPIs() ProductAPIInformer
	// ProductGroups returns a ProductGroupInformer.
	ProductGroups() ProductGroupInformer
	// ProductPolicies returns a ProductPolicyInformer.
	ProductPolicies() ProductPolicyInformer
	// Properties returns a PropertyInformer.
	Properties() PropertyInformer
	// RedisCaches returns a RedisCacheInformer.
	RedisCaches() RedisCacheInformer
	// Subscriptions returns a SubscriptionInformer.
	Subscriptions() SubscriptionInformer
	// Tags returns a TagInformer.
	Tags() TagInformer
	// Users returns a UserInformer.
	Users() UserInformer
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

// Apis returns a ApiInformer.
func (v *version) Apis() ApiInformer {
	return &apiInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApiDiagnostics returns a ApiDiagnosticInformer.
func (v *version) ApiDiagnostics() ApiDiagnosticInformer {
	return &apiDiagnosticInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApiManagements returns a ApiManagementInformer.
func (v *version) ApiManagements() ApiManagementInformer {
	return &apiManagementInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApiOperations returns a ApiOperationInformer.
func (v *version) ApiOperations() ApiOperationInformer {
	return &apiOperationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApiOperationPolicies returns a ApiOperationPolicyInformer.
func (v *version) ApiOperationPolicies() ApiOperationPolicyInformer {
	return &apiOperationPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApiOperationTags returns a ApiOperationTagInformer.
func (v *version) ApiOperationTags() ApiOperationTagInformer {
	return &apiOperationTagInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApiPolicies returns a ApiPolicyInformer.
func (v *version) ApiPolicies() ApiPolicyInformer {
	return &apiPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApiReleases returns a ApiReleaseInformer.
func (v *version) ApiReleases() ApiReleaseInformer {
	return &apiReleaseInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApiSchemas returns a ApiSchemaInformer.
func (v *version) ApiSchemas() ApiSchemaInformer {
	return &apiSchemaInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApiTags returns a ApiTagInformer.
func (v *version) ApiTags() ApiTagInformer {
	return &apiTagInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApiVersionSets returns a ApiVersionSetInformer.
func (v *version) ApiVersionSets() ApiVersionSetInformer {
	return &apiVersionSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AuthorizationServers returns a AuthorizationServerInformer.
func (v *version) AuthorizationServers() AuthorizationServerInformer {
	return &authorizationServerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Backends returns a BackendInformer.
func (v *version) Backends() BackendInformer {
	return &backendInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Certificates returns a CertificateInformer.
func (v *version) Certificates() CertificateInformer {
	return &certificateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CustomDomains returns a CustomDomainInformer.
func (v *version) CustomDomains() CustomDomainInformer {
	return &customDomainInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Diagnostics returns a DiagnosticInformer.
func (v *version) Diagnostics() DiagnosticInformer {
	return &diagnosticInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EmailTemplates returns a EmailTemplateInformer.
func (v *version) EmailTemplates() EmailTemplateInformer {
	return &emailTemplateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Gateways returns a GatewayInformer.
func (v *version) Gateways() GatewayInformer {
	return &gatewayInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GatewayAPIs returns a GatewayAPIInformer.
func (v *version) GatewayAPIs() GatewayAPIInformer {
	return &gatewayAPIInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Groups returns a GroupInformer.
func (v *version) Groups() GroupInformer {
	return &groupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GroupUsers returns a GroupUserInformer.
func (v *version) GroupUsers() GroupUserInformer {
	return &groupUserInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// IdentityProviderAads returns a IdentityProviderAadInformer.
func (v *version) IdentityProviderAads() IdentityProviderAadInformer {
	return &identityProviderAadInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// IdentityProviderAadb2cs returns a IdentityProviderAadb2cInformer.
func (v *version) IdentityProviderAadb2cs() IdentityProviderAadb2cInformer {
	return &identityProviderAadb2cInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// IdentityProviderFacebooks returns a IdentityProviderFacebookInformer.
func (v *version) IdentityProviderFacebooks() IdentityProviderFacebookInformer {
	return &identityProviderFacebookInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// IdentityProviderGoogles returns a IdentityProviderGoogleInformer.
func (v *version) IdentityProviderGoogles() IdentityProviderGoogleInformer {
	return &identityProviderGoogleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// IdentityProviderMicrosofts returns a IdentityProviderMicrosoftInformer.
func (v *version) IdentityProviderMicrosofts() IdentityProviderMicrosoftInformer {
	return &identityProviderMicrosoftInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// IdentityProviderTwitters returns a IdentityProviderTwitterInformer.
func (v *version) IdentityProviderTwitters() IdentityProviderTwitterInformer {
	return &identityProviderTwitterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Loggers returns a LoggerInformer.
func (v *version) Loggers() LoggerInformer {
	return &loggerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NamedValues returns a NamedValueInformer.
func (v *version) NamedValues() NamedValueInformer {
	return &namedValueInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NotificationRecipientEmails returns a NotificationRecipientEmailInformer.
func (v *version) NotificationRecipientEmails() NotificationRecipientEmailInformer {
	return &notificationRecipientEmailInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NotificationRecipientUsers returns a NotificationRecipientUserInformer.
func (v *version) NotificationRecipientUsers() NotificationRecipientUserInformer {
	return &notificationRecipientUserInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// OpenidConnectProviders returns a OpenidConnectProviderInformer.
func (v *version) OpenidConnectProviders() OpenidConnectProviderInformer {
	return &openidConnectProviderInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Policies returns a PolicyInformer.
func (v *version) Policies() PolicyInformer {
	return &policyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Products returns a ProductInformer.
func (v *version) Products() ProductInformer {
	return &productInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ProductAPIs returns a ProductAPIInformer.
func (v *version) ProductAPIs() ProductAPIInformer {
	return &productAPIInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ProductGroups returns a ProductGroupInformer.
func (v *version) ProductGroups() ProductGroupInformer {
	return &productGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ProductPolicies returns a ProductPolicyInformer.
func (v *version) ProductPolicies() ProductPolicyInformer {
	return &productPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Properties returns a PropertyInformer.
func (v *version) Properties() PropertyInformer {
	return &propertyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// RedisCaches returns a RedisCacheInformer.
func (v *version) RedisCaches() RedisCacheInformer {
	return &redisCacheInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Subscriptions returns a SubscriptionInformer.
func (v *version) Subscriptions() SubscriptionInformer {
	return &subscriptionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Tags returns a TagInformer.
func (v *version) Tags() TagInformer {
	return &tagInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Users returns a UserInformer.
func (v *version) Users() UserInformer {
	return &userInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
