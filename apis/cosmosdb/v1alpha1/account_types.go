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

type Account struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccountSpec   `json:"spec,omitempty"`
	Status            AccountStatus `json:"status,omitempty"`
}

type AccountSpecBackup struct {
	// +optional
	IntervalInMinutes *int64 `json:"intervalInMinutes,omitempty" tf:"interval_in_minutes"`
	// +optional
	RetentionInHours *int64  `json:"retentionInHours,omitempty" tf:"retention_in_hours"`
	Type             *string `json:"type" tf:"type"`
}

type AccountSpecCapabilities struct {
	Name *string `json:"name" tf:"name"`
}

type AccountSpecConsistencyPolicy struct {
	ConsistencyLevel *string `json:"consistencyLevel" tf:"consistency_level"`
	// +optional
	MaxIntervalInSeconds *int64 `json:"maxIntervalInSeconds,omitempty" tf:"max_interval_in_seconds"`
	// +optional
	MaxStalenessPrefix *int64 `json:"maxStalenessPrefix,omitempty" tf:"max_staleness_prefix"`
}

type AccountSpecCorsRule struct {
	// +kubebuilder:validation:MaxItems=64
	AllowedHeaders []string `json:"allowedHeaders" tf:"allowed_headers"`
	// +kubebuilder:validation:MaxItems=64
	AllowedMethods []string `json:"allowedMethods" tf:"allowed_methods"`
	// +kubebuilder:validation:MaxItems=64
	AllowedOrigins []string `json:"allowedOrigins" tf:"allowed_origins"`
	// +kubebuilder:validation:MaxItems=64
	ExposedHeaders  []string `json:"exposedHeaders" tf:"exposed_headers"`
	MaxAgeInSeconds *int64   `json:"maxAgeInSeconds" tf:"max_age_in_seconds"`
}

type AccountSpecGeoLocation struct {
	FailoverPriority *int64 `json:"failoverPriority" tf:"failover_priority"`
	// +optional
	ID       *string `json:"ID,omitempty" tf:"id"`
	Location *string `json:"location" tf:"location"`
	// +optional
	// Deprecated
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
	// +optional
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant"`
}

type AccountSpecIdentity struct {
	// +optional
	PrincipalID *string `json:"principalID,omitempty" tf:"principal_id"`
	// +optional
	TenantID *string `json:"tenantID,omitempty" tf:"tenant_id"`
	Type     *string `json:"type" tf:"type"`
}

type AccountSpecVirtualNetworkRule struct {
	ID *string `json:"ID" tf:"id"`
	// +optional
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty" tf:"ignore_missing_vnet_service_endpoint"`
}

type AccountSpec struct {
	State *AccountSpecResource `json:"state,omitempty" tf:"-"`

	Resource AccountSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type AccountSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AccessKeyMetadataWritesEnabled *bool `json:"accessKeyMetadataWritesEnabled,omitempty" tf:"access_key_metadata_writes_enabled"`
	// +optional
	AnalyticalStorageEnabled *bool `json:"analyticalStorageEnabled,omitempty" tf:"analytical_storage_enabled"`
	// +optional
	Backup *AccountSpecBackup `json:"backup,omitempty" tf:"backup"`
	// +optional
	Capabilities []AccountSpecCapabilities `json:"capabilities,omitempty" tf:"capabilities"`
	// +optional
	ConsistencyPolicy *AccountSpecConsistencyPolicy `json:"consistencyPolicy" tf:"consistency_policy"`
	// +optional
	CorsRule *AccountSpecCorsRule `json:"corsRule,omitempty" tf:"cors_rule"`
	// +optional
	EnableAutomaticFailover *bool `json:"enableAutomaticFailover,omitempty" tf:"enable_automatic_failover"`
	// +optional
	EnableFreeTier *bool `json:"enableFreeTier,omitempty" tf:"enable_free_tier"`
	// +optional
	EnableMultipleWriteLocations *bool `json:"enableMultipleWriteLocations,omitempty" tf:"enable_multiple_write_locations"`
	// +optional
	Endpoint    *string                  `json:"endpoint,omitempty" tf:"endpoint"`
	GeoLocation []AccountSpecGeoLocation `json:"geoLocation" tf:"geo_location"`
	// +optional
	Identity *AccountSpecIdentity `json:"identity,omitempty" tf:"identity"`
	// +optional
	IpRangeFilter *string `json:"ipRangeFilter,omitempty" tf:"ip_range_filter"`
	// +optional
	IsVirtualNetworkFilterEnabled *bool `json:"isVirtualNetworkFilterEnabled,omitempty" tf:"is_virtual_network_filter_enabled"`
	// +optional
	KeyVaultKeyID *string `json:"keyVaultKeyID,omitempty" tf:"key_vault_key_id"`
	// +optional
	Kind     *string `json:"kind,omitempty" tf:"kind"`
	Location *string `json:"location" tf:"location"`
	// +optional
	MongoServerVersion *string `json:"mongoServerVersion,omitempty" tf:"mongo_server_version"`
	Name               *string `json:"name" tf:"name"`
	// +optional
	NetworkACLBypassForAzureServices *bool `json:"networkACLBypassForAzureServices,omitempty" tf:"network_acl_bypass_for_azure_services"`
	// +optional
	NetworkACLBypassIDS []string `json:"networkACLBypassIDS,omitempty" tf:"network_acl_bypass_ids"`
	OfferType           *string  `json:"offerType" tf:"offer_type"`
	// +optional
	PrimaryKey *string `json:"-" sensitive:"true" tf:"primary_key"`
	// +optional
	PrimaryMasterKey *string `json:"-" sensitive:"true" tf:"primary_master_key"`
	// +optional
	PrimaryReadonlyKey *string `json:"-" sensitive:"true" tf:"primary_readonly_key"`
	// +optional
	PrimaryReadonlyMasterKey *string `json:"-" sensitive:"true" tf:"primary_readonly_master_key"`
	// +optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled"`
	// +optional
	ReadEndpoints     []string `json:"readEndpoints,omitempty" tf:"read_endpoints"`
	ResourceGroupName *string  `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryKey *string `json:"-" sensitive:"true" tf:"secondary_key"`
	// +optional
	SecondaryMasterKey *string `json:"-" sensitive:"true" tf:"secondary_master_key"`
	// +optional
	SecondaryReadonlyKey *string `json:"-" sensitive:"true" tf:"secondary_readonly_key"`
	// +optional
	SecondaryReadonlyMasterKey *string `json:"-" sensitive:"true" tf:"secondary_readonly_master_key"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	VirtualNetworkRule []AccountSpecVirtualNetworkRule `json:"virtualNetworkRule,omitempty" tf:"virtual_network_rule"`
	// +optional
	WriteEndpoints []string `json:"writeEndpoints,omitempty" tf:"write_endpoints"`
}

type AccountStatus struct {
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

// AccountList is a list of Accounts
type AccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Account CRD objects
	Items []Account `json:"items,omitempty"`
}
