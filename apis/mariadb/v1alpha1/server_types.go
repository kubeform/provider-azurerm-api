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

type Server struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServerSpec   `json:"spec,omitempty"`
	Status            ServerStatus `json:"status,omitempty"`
}

type ServerSpecStorageProfile struct {
	// +optional
	// Deprecated
	AutoGrow *string `json:"autoGrow,omitempty" tf:"auto_grow"`
	// +optional
	// Deprecated
	BackupRetentionDays *int64 `json:"backupRetentionDays,omitempty" tf:"backup_retention_days"`
	// +optional
	// Deprecated
	GeoRedundantBackup *string `json:"geoRedundantBackup,omitempty" tf:"geo_redundant_backup"`
	// +optional
	// Deprecated
	StorageMb *int64 `json:"storageMb,omitempty" tf:"storage_mb"`
}

type ServerSpec struct {
	State *ServerSpecResource `json:"state,omitempty" tf:"-"`

	Resource ServerSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type ServerSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login"`
	// +optional
	AdministratorLoginPassword *string `json:"-" sensitive:"true" tf:"administrator_login_password"`
	// +optional
	AutoGrowEnabled *bool `json:"autoGrowEnabled,omitempty" tf:"auto_grow_enabled"`
	// +optional
	BackupRetentionDays *int64 `json:"backupRetentionDays,omitempty" tf:"backup_retention_days"`
	// +optional
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode"`
	// +optional
	CreationSourceServerID *string `json:"creationSourceServerID,omitempty" tf:"creation_source_server_id"`
	// +optional
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn"`
	// +optional
	GeoRedundantBackupEnabled *bool   `json:"geoRedundantBackupEnabled,omitempty" tf:"geo_redundant_backup_enabled"`
	Location                  *string `json:"location" tf:"location"`
	Name                      *string `json:"name" tf:"name"`
	// +optional
	PublicNetworkAccessEnabled *bool   `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled"`
	ResourceGroupName          *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	RestorePointInTime *string `json:"restorePointInTime,omitempty" tf:"restore_point_in_time"`
	SkuName            *string `json:"skuName" tf:"sku_name"`
	// +optional
	// Deprecated
	SslEnforcement *string `json:"sslEnforcement,omitempty" tf:"ssl_enforcement"`
	// +optional
	SslEnforcementEnabled *bool `json:"sslEnforcementEnabled,omitempty" tf:"ssl_enforcement_enabled"`
	// +optional
	StorageMb *int64 `json:"storageMb,omitempty" tf:"storage_mb"`
	// +optional
	// Deprecated
	StorageProfile *ServerSpecStorageProfile `json:"storageProfile,omitempty" tf:"storage_profile"`
	// +optional
	Tags    *map[string]string `json:"tags,omitempty" tf:"tags"`
	Version *string            `json:"version" tf:"version"`
}

type ServerStatus struct {
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

// ServerList is a list of Servers
type ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Server CRD objects
	Items []Server `json:"items,omitempty"`
}
