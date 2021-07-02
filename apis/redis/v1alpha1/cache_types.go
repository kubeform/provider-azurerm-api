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

type Cache struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CacheSpec   `json:"spec,omitempty"`
	Status            CacheStatus `json:"status,omitempty"`
}

type CacheSpecPatchSchedule struct {
	DayOfWeek *string `json:"dayOfWeek" tf:"day_of_week"`
	// +optional
	StartHourUtc *int64 `json:"startHourUtc,omitempty" tf:"start_hour_utc"`
}

type CacheSpecRedisConfiguration struct {
	// +optional
	AofBackupEnabled *bool `json:"aofBackupEnabled,omitempty" tf:"aof_backup_enabled"`
	// +optional
	AofStorageConnectionString0 *string `json:"-" sensitive:"true" tf:"aof_storage_connection_string_0"`
	// +optional
	AofStorageConnectionString1 *string `json:"-" sensitive:"true" tf:"aof_storage_connection_string_1"`
	// +optional
	EnableAuthentication *bool `json:"enableAuthentication,omitempty" tf:"enable_authentication"`
	// +optional
	Maxclients *int64 `json:"maxclients,omitempty" tf:"maxclients"`
	// +optional
	MaxfragmentationmemoryReserved *int64 `json:"maxfragmentationmemoryReserved,omitempty" tf:"maxfragmentationmemory_reserved"`
	// +optional
	MaxmemoryDelta *int64 `json:"maxmemoryDelta,omitempty" tf:"maxmemory_delta"`
	// +optional
	MaxmemoryPolicy *string `json:"maxmemoryPolicy,omitempty" tf:"maxmemory_policy"`
	// +optional
	MaxmemoryReserved *int64 `json:"maxmemoryReserved,omitempty" tf:"maxmemory_reserved"`
	// +optional
	NotifyKeyspaceEvents *string `json:"notifyKeyspaceEvents,omitempty" tf:"notify_keyspace_events"`
	// +optional
	RdbBackupEnabled *bool `json:"rdbBackupEnabled,omitempty" tf:"rdb_backup_enabled"`
	// +optional
	RdbBackupFrequency *int64 `json:"rdbBackupFrequency,omitempty" tf:"rdb_backup_frequency"`
	// +optional
	RdbBackupMaxSnapshotCount *int64 `json:"rdbBackupMaxSnapshotCount,omitempty" tf:"rdb_backup_max_snapshot_count"`
	// +optional
	RdbStorageConnectionString *string `json:"-" sensitive:"true" tf:"rdb_storage_connection_string"`
}

type CacheSpec struct {
	KubeformOutput *CacheSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource CacheSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type CacheSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Capacity *int64 `json:"capacity" tf:"capacity"`
	// +optional
	EnableNonSslPort *bool   `json:"enableNonSslPort,omitempty" tf:"enable_non_ssl_port"`
	Family           *string `json:"family" tf:"family"`
	// +optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname"`
	Location *string `json:"location" tf:"location"`
	// +optional
	MinimumTlsVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version"`
	Name              *string `json:"name" tf:"name"`
	// +optional
	PatchSchedule []CacheSpecPatchSchedule `json:"patchSchedule,omitempty" tf:"patch_schedule"`
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	PrimaryAccessKey *string `json:"-" sensitive:"true" tf:"primary_access_key"`
	// +optional
	PrimaryConnectionString *string `json:"-" sensitive:"true" tf:"primary_connection_string"`
	// +optional
	PrivateStaticIPAddress *string `json:"privateStaticIPAddress,omitempty" tf:"private_static_ip_address"`
	// +optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled"`
	// +optional
	RedisConfiguration *CacheSpecRedisConfiguration `json:"redisConfiguration,omitempty" tf:"redis_configuration"`
	// +optional
	ReplicasPerMaster *int64  `json:"replicasPerMaster,omitempty" tf:"replicas_per_master"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryAccessKey *string `json:"-" sensitive:"true" tf:"secondary_access_key"`
	// +optional
	SecondaryConnectionString *string `json:"-" sensitive:"true" tf:"secondary_connection_string"`
	// +optional
	ShardCount *int64  `json:"shardCount,omitempty" tf:"shard_count"`
	SkuName    *string `json:"skuName" tf:"sku_name"`
	// +optional
	SslPort *int64 `json:"sslPort,omitempty" tf:"ssl_port"`
	// +optional
	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	Zones []string `json:"zones,omitempty" tf:"zones"`
}

type CacheStatus struct {
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

// CacheList is a list of Caches
type CacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Cache CRD objects
	Items []Cache `json:"items,omitempty"`
}