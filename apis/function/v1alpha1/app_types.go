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

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type App struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppSpec   `json:"spec,omitempty"`
	Status            AppStatus `json:"status,omitempty"`
}

type AppSpecAuthSettingsActiveDirectory struct {
	// +optional
	AllowedAudiences []string `json:"allowedAudiences,omitempty" tf:"allowed_audiences"`
	ClientID         *string  `json:"clientID" tf:"client_id"`
	// +optional
	ClientSecret *string `json:"-" sensitive:"true" tf:"client_secret"`
}

type AppSpecAuthSettingsFacebook struct {
	AppID     *string `json:"appID" tf:"app_id"`
	AppSecret *string `json:"-" sensitive:"true" tf:"app_secret"`
	// +optional
	OauthScopes []string `json:"oauthScopes,omitempty" tf:"oauth_scopes"`
}

type AppSpecAuthSettingsGoogle struct {
	ClientID     *string `json:"clientID" tf:"client_id"`
	ClientSecret *string `json:"-" sensitive:"true" tf:"client_secret"`
	// +optional
	OauthScopes []string `json:"oauthScopes,omitempty" tf:"oauth_scopes"`
}

type AppSpecAuthSettingsMicrosoft struct {
	ClientID     *string `json:"clientID" tf:"client_id"`
	ClientSecret *string `json:"-" sensitive:"true" tf:"client_secret"`
	// +optional
	OauthScopes []string `json:"oauthScopes,omitempty" tf:"oauth_scopes"`
}

type AppSpecAuthSettingsTwitter struct {
	ConsumerKey    *string `json:"consumerKey" tf:"consumer_key"`
	ConsumerSecret *string `json:"-" sensitive:"true" tf:"consumer_secret"`
}

type AppSpecAuthSettings struct {
	// +optional
	ActiveDirectory *AppSpecAuthSettingsActiveDirectory `json:"activeDirectory,omitempty" tf:"active_directory"`
	// +optional
	AdditionalLoginParams *map[string]string `json:"additionalLoginParams,omitempty" tf:"additional_login_params"`
	// +optional
	AllowedExternalRedirectUrls []string `json:"allowedExternalRedirectUrls,omitempty" tf:"allowed_external_redirect_urls"`
	// +optional
	DefaultProvider *string `json:"defaultProvider,omitempty" tf:"default_provider"`
	Enabled         *bool   `json:"enabled" tf:"enabled"`
	// +optional
	Facebook *AppSpecAuthSettingsFacebook `json:"facebook,omitempty" tf:"facebook"`
	// +optional
	Google *AppSpecAuthSettingsGoogle `json:"google,omitempty" tf:"google"`
	// +optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer"`
	// +optional
	Microsoft *AppSpecAuthSettingsMicrosoft `json:"microsoft,omitempty" tf:"microsoft"`
	// +optional
	RuntimeVersion *string `json:"runtimeVersion,omitempty" tf:"runtime_version"`
	// +optional
	TokenRefreshExtensionHours *float64 `json:"tokenRefreshExtensionHours,omitempty" tf:"token_refresh_extension_hours"`
	// +optional
	TokenStoreEnabled *bool `json:"tokenStoreEnabled,omitempty" tf:"token_store_enabled"`
	// +optional
	Twitter *AppSpecAuthSettingsTwitter `json:"twitter,omitempty" tf:"twitter"`
	// +optional
	UnauthenticatedClientAction *string `json:"unauthenticatedClientAction,omitempty" tf:"unauthenticated_client_action"`
}

type AppSpecConnectionString struct {
	Name  *string `json:"name" tf:"name"`
	Type  *string `json:"type" tf:"type"`
	Value *string `json:"-" sensitive:"true" tf:"value"`
}

type AppSpecIdentity struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	IdentityIDS []string `json:"identityIDS,omitempty" tf:"identity_ids"`
	// +optional
	PrincipalID *string `json:"principalID,omitempty" tf:"principal_id"`
	// +optional
	TenantID *string `json:"tenantID,omitempty" tf:"tenant_id"`
	Type     *string `json:"type" tf:"type"`
}

type AppSpecSiteConfigCors struct {
	AllowedOrigins []string `json:"allowedOrigins" tf:"allowed_origins"`
	// +optional
	SupportCredentials *bool `json:"supportCredentials,omitempty" tf:"support_credentials"`
}

type AppSpecSiteConfigIpRestrictionHeaders struct {
	// +optional
	// +kubebuilder:validation:MaxItems=8
	XAzureFdid []string `json:"xAzureFdid,omitempty" tf:"x_azure_fdid"`
	// +optional
	XFdHealthProbe []string `json:"xFdHealthProbe,omitempty" tf:"x_fd_health_probe"`
	// +optional
	// +kubebuilder:validation:MaxItems=8
	XForwardedFor []string `json:"xForwardedFor,omitempty" tf:"x_forwarded_for"`
	// +optional
	// +kubebuilder:validation:MaxItems=8
	XForwardedHost []string `json:"xForwardedHost,omitempty" tf:"x_forwarded_host"`
}

type AppSpecSiteConfigIpRestriction struct {
	// +optional
	Action *string `json:"action,omitempty" tf:"action"`
	// +optional
	Headers *AppSpecSiteConfigIpRestrictionHeaders `json:"headers,omitempty" tf:"headers"`
	// +optional
	IpAddress *string `json:"ipAddress,omitempty" tf:"ip_address"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Priority *int64 `json:"priority,omitempty" tf:"priority"`
	// +optional
	ServiceTag *string `json:"serviceTag,omitempty" tf:"service_tag"`
	// +optional
	VirtualNetworkSubnetID *string `json:"virtualNetworkSubnetID,omitempty" tf:"virtual_network_subnet_id"`
}

type AppSpecSiteConfigScmIPRestrictionHeaders struct {
	// +optional
	// +kubebuilder:validation:MaxItems=8
	XAzureFdid []string `json:"xAzureFdid,omitempty" tf:"x_azure_fdid"`
	// +optional
	XFdHealthProbe []string `json:"xFdHealthProbe,omitempty" tf:"x_fd_health_probe"`
	// +optional
	// +kubebuilder:validation:MaxItems=8
	XForwardedFor []string `json:"xForwardedFor,omitempty" tf:"x_forwarded_for"`
	// +optional
	// +kubebuilder:validation:MaxItems=8
	XForwardedHost []string `json:"xForwardedHost,omitempty" tf:"x_forwarded_host"`
}

type AppSpecSiteConfigScmIPRestriction struct {
	// +optional
	Action *string `json:"action,omitempty" tf:"action"`
	// +optional
	Headers *AppSpecSiteConfigScmIPRestrictionHeaders `json:"headers,omitempty" tf:"headers"`
	// +optional
	IpAddress *string `json:"ipAddress,omitempty" tf:"ip_address"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Priority *int64 `json:"priority,omitempty" tf:"priority"`
	// +optional
	ServiceTag *string `json:"serviceTag,omitempty" tf:"service_tag"`
	// +optional
	VirtualNetworkSubnetID *string `json:"virtualNetworkSubnetID,omitempty" tf:"virtual_network_subnet_id"`
}

type AppSpecSiteConfig struct {
	// +optional
	AlwaysOn *bool `json:"alwaysOn,omitempty" tf:"always_on"`
	// +optional
	AutoSwapSlotName *string `json:"autoSwapSlotName,omitempty" tf:"auto_swap_slot_name"`
	// +optional
	Cors *AppSpecSiteConfigCors `json:"cors,omitempty" tf:"cors"`
	// +optional
	FtpsState *string `json:"ftpsState,omitempty" tf:"ftps_state"`
	// +optional
	HealthCheckPath *string `json:"healthCheckPath,omitempty" tf:"health_check_path"`
	// +optional
	Http2Enabled *bool `json:"http2Enabled,omitempty" tf:"http2_enabled"`
	// +optional
	IpRestriction []AppSpecSiteConfigIpRestriction `json:"ipRestriction,omitempty" tf:"ip_restriction"`
	// +optional
	JavaVersion *string `json:"javaVersion,omitempty" tf:"java_version"`
	// +optional
	LinuxFxVersion *string `json:"linuxFxVersion,omitempty" tf:"linux_fx_version"`
	// +optional
	MinTlsVersion *string `json:"minTlsVersion,omitempty" tf:"min_tls_version"`
	// +optional
	PreWarmedInstanceCount *int64 `json:"preWarmedInstanceCount,omitempty" tf:"pre_warmed_instance_count"`
	// +optional
	ScmIPRestriction []AppSpecSiteConfigScmIPRestriction `json:"scmIPRestriction,omitempty" tf:"scm_ip_restriction"`
	// +optional
	ScmType *string `json:"scmType,omitempty" tf:"scm_type"`
	// +optional
	ScmUseMainIPRestriction *bool `json:"scmUseMainIPRestriction,omitempty" tf:"scm_use_main_ip_restriction"`
	// +optional
	Use32BitWorkerProcess *bool `json:"use32BitWorkerProcess,omitempty" tf:"use_32_bit_worker_process"`
	// +optional
	WebsocketsEnabled *bool `json:"websocketsEnabled,omitempty" tf:"websockets_enabled"`
}

type AppSpecSiteCredential struct {
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	Username *string `json:"username,omitempty" tf:"username"`
}

type AppSpecSourceControl struct {
	// +optional
	Branch *string `json:"branch,omitempty" tf:"branch"`
	// +optional
	ManualIntegration *bool `json:"manualIntegration,omitempty" tf:"manual_integration"`
	// +optional
	RepoURL *string `json:"repoURL,omitempty" tf:"repo_url"`
	// +optional
	RollbackEnabled *bool `json:"rollbackEnabled,omitempty" tf:"rollback_enabled"`
	// +optional
	UseMercurial *bool `json:"useMercurial,omitempty" tf:"use_mercurial"`
}

type AppSpec struct {
	State *AppSpecResource `json:"state,omitempty" tf:"-"`

	Resource AppSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type AppSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AppServicePlanID *string `json:"appServicePlanID" tf:"app_service_plan_id"`
	// +optional
	AppSettings *map[string]string `json:"appSettings,omitempty" tf:"app_settings"`
	// +optional
	AuthSettings *AppSpecAuthSettings `json:"authSettings,omitempty" tf:"auth_settings"`
	// +optional
	ClientAffinityEnabled *bool `json:"clientAffinityEnabled,omitempty" tf:"client_affinity_enabled"`
	// +optional
	ClientCertMode *string `json:"clientCertMode,omitempty" tf:"client_cert_mode"`
	// +optional
	ConnectionString []AppSpecConnectionString `json:"connectionString,omitempty" tf:"connection_string"`
	// +optional
	CustomDomainVerificationID *string `json:"customDomainVerificationID,omitempty" tf:"custom_domain_verification_id"`
	// +optional
	DailyMemoryTimeQuota *int64 `json:"dailyMemoryTimeQuota,omitempty" tf:"daily_memory_time_quota"`
	// +optional
	DefaultHostname *string `json:"defaultHostname,omitempty" tf:"default_hostname"`
	// +optional
	EnableBuiltinLogging *bool `json:"enableBuiltinLogging,omitempty" tf:"enable_builtin_logging"`
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	HttpsOnly *bool `json:"httpsOnly,omitempty" tf:"https_only"`
	// +optional
	Identity *AppSpecIdentity `json:"identity,omitempty" tf:"identity"`
	// +optional
	Kind     *string `json:"kind,omitempty" tf:"kind"`
	Location *string `json:"location" tf:"location"`
	Name     *string `json:"name" tf:"name"`
	// +optional
	OsType *string `json:"osType,omitempty" tf:"os_type"`
	// +optional
	OutboundIPAddresses *string `json:"outboundIPAddresses,omitempty" tf:"outbound_ip_addresses"`
	// +optional
	PossibleOutboundIPAddresses *string `json:"possibleOutboundIPAddresses,omitempty" tf:"possible_outbound_ip_addresses"`
	ResourceGroupName           *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SiteConfig *AppSpecSiteConfig `json:"siteConfig,omitempty" tf:"site_config"`
	// +optional
	SiteCredential []AppSpecSiteCredential `json:"siteCredential,omitempty" tf:"site_credential"`
	// +optional
	SourceControl *AppSpecSourceControl `json:"sourceControl,omitempty" tf:"source_control"`
	// +optional
	StorageAccountAccessKey *string `json:"-" sensitive:"true" tf:"storage_account_access_key"`
	// +optional
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name"`
	// +optional
	StorageConnectionString *string `json:"-" sensitive:"true" tf:"storage_connection_string"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
}

type AppStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppList is a list of Apps
type AppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of App CRD objects
	Items []App `json:"items,omitempty"`
}
