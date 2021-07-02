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

// Code generated by Kubeform. DO NOT EDIT.

package provider

type AzurermSpecFeaturesKeyVault struct {
	// +optional
	PurgeSoftDeleteOnDestroy *bool `json:"purgeSoftDeleteOnDestroy,omitempty" tf:"purge_soft_delete_on_destroy"`
	// +optional
	RecoverSoftDeletedKeyVaults *bool `json:"recoverSoftDeletedKeyVaults,omitempty" tf:"recover_soft_deleted_key_vaults"`
}

type AzurermSpecFeaturesLogAnalyticsWorkspace struct {
	PermanentlyDeleteOnDestroy *bool `json:"permanentlyDeleteOnDestroy" tf:"permanently_delete_on_destroy"`
}

type AzurermSpecFeaturesNetwork struct {
	RelaxedLocking *bool `json:"relaxedLocking" tf:"relaxed_locking"`
}

type AzurermSpecFeaturesTemplateDeployment struct {
	DeleteNestedItemsDuringDeletion *bool `json:"deleteNestedItemsDuringDeletion" tf:"delete_nested_items_during_deletion"`
}

type AzurermSpecFeaturesVirtualMachine struct {
	// +optional
	DeleteOsDiskOnDeletion *bool `json:"deleteOsDiskOnDeletion,omitempty" tf:"delete_os_disk_on_deletion"`
	// +optional
	GracefulShutdown *bool `json:"gracefulShutdown,omitempty" tf:"graceful_shutdown"`
	// +optional
	SkipShutdownAndForceDelete *bool `json:"skipShutdownAndForceDelete,omitempty" tf:"skip_shutdown_and_force_delete"`
}

type AzurermSpecFeaturesVirtualMachineScaleSet struct {
	// +optional
	ForceDelete               *bool `json:"forceDelete,omitempty" tf:"force_delete"`
	RollInstancesWhenRequired *bool `json:"rollInstancesWhenRequired" tf:"roll_instances_when_required"`
}

type AzurermSpecFeatures struct {
	// +optional
	KeyVault *AzurermSpecFeaturesKeyVault `json:"keyVault,omitempty" tf:"key_vault"`
	// +optional
	LogAnalyticsWorkspace *AzurermSpecFeaturesLogAnalyticsWorkspace `json:"logAnalyticsWorkspace,omitempty" tf:"log_analytics_workspace"`
	// +optional
	Network *AzurermSpecFeaturesNetwork `json:"network,omitempty" tf:"network"`
	// +optional
	TemplateDeployment *AzurermSpecFeaturesTemplateDeployment `json:"templateDeployment,omitempty" tf:"template_deployment"`
	// +optional
	VirtualMachine *AzurermSpecFeaturesVirtualMachine `json:"virtualMachine,omitempty" tf:"virtual_machine"`
	// +optional
	VirtualMachineScaleSet *AzurermSpecFeaturesVirtualMachineScaleSet `json:"virtualMachineScaleSet,omitempty" tf:"virtual_machine_scale_set"`
}

type AzurermSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=3
	AuxiliaryTenantIDS []string `json:"auxiliaryTenantIDS,omitempty" tf:"auxiliary_tenant_ids"`
	// The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client Certificate
	// +optional
	ClientCertificatePassword *string `json:"clientCertificatePassword,omitempty" tf:"client_certificate_password"`
	// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service Principal using a Client Certificate.
	// +optional
	ClientCertificatePath *string `json:"clientCertificatePath,omitempty" tf:"client_certificate_path"`
	// The Client ID which should be used.
	// +optional
	ClientID *string `json:"clientID,omitempty" tf:"client_id"`
	// The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
	// +optional
	ClientSecret *string `json:"clientSecret,omitempty" tf:"client_secret"`
	// This will disable the x-ms-correlation-request-id header.
	// +optional
	DisableCorrelationRequestID *bool `json:"disableCorrelationRequestID,omitempty" tf:"disable_correlation_request_id"`
	// This will disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
	// +optional
	DisableTerraformPartnerID *bool `json:"disableTerraformPartnerID,omitempty" tf:"disable_terraform_partner_id"`
	// The Cloud Environment which should be used. Possible values are public, usgovernment, german, and china. Defaults to public.
	Environment *string              `json:"environment" tf:"environment"`
	Features    *AzurermSpecFeatures `json:"features" tf:"features"`
	// The Hostname which should be used for the Azure Metadata Service.
	MetadataHost *string `json:"metadataHost" tf:"metadata_host"`
	// Deprecated - replaced by `metadata_host`.
	// +optional
	// Deprecated
	MetadataURL *string `json:"metadataURL,omitempty" tf:"metadata_url"`
	// The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected automatically.
	// +optional
	MsiEndpoint *string `json:"msiEndpoint,omitempty" tf:"msi_endpoint"`
	// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
	// +optional
	PartnerID *string `json:"partnerID,omitempty" tf:"partner_id"`
	// [DEPRECATED] This will cause the AzureRM Provider to skip verifying the credentials being used are valid.
	// +optional
	// Deprecated
	SkipCredentialsValidation *bool `json:"skipCredentialsValidation,omitempty" tf:"skip_credentials_validation"`
	// Should the AzureRM Provider skip registering all of the Resource Providers that it supports, if they're not already registered?
	// +optional
	SkipProviderRegistration *bool `json:"skipProviderRegistration,omitempty" tf:"skip_provider_registration"`
	// Should the AzureRM Provider use AzureAD to access the Storage Data Plane API's?
	// +optional
	StorageUseAzuread *bool `json:"storageUseAzuread,omitempty" tf:"storage_use_azuread"`
	// The Subscription ID which should be used.
	// +optional
	SubscriptionID *string `json:"subscriptionID,omitempty" tf:"subscription_id"`
	// The Tenant ID which should be used.
	// +optional
	TenantID *string `json:"tenantID,omitempty" tf:"tenant_id"`
	// Allowed Managed Service Identity be used for Authentication.
	// +optional
	UseMsi *bool `json:"useMsi,omitempty" tf:"use_msi"`
}