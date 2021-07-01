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

type Database struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseSpec   `json:"spec,omitempty"`
	Status            DatabaseStatus `json:"status,omitempty"`
}

type DatabaseSpecExtendedAuditingPolicy struct {
	// +optional
	LogMonitoringEnabled *bool `json:"logMonitoringEnabled,omitempty" tf:"log_monitoring_enabled"`
	// +optional
	RetentionInDays *int64 `json:"retentionInDays,omitempty" tf:"retention_in_days"`
	// +optional
	StorageAccountAccessKey *string `json:"-" sensitive:"true" tf:"storage_account_access_key"`
	// +optional
	StorageAccountAccessKeyIsSecondary *bool `json:"storageAccountAccessKeyIsSecondary,omitempty" tf:"storage_account_access_key_is_secondary"`
	// +optional
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint"`
}

type DatabaseSpecLongTermRetentionPolicy struct {
	// +optional
	MonthlyRetention *string `json:"monthlyRetention,omitempty" tf:"monthly_retention"`
	// +optional
	WeekOfYear *int64 `json:"weekOfYear,omitempty" tf:"week_of_year"`
	// +optional
	WeeklyRetention *string `json:"weeklyRetention,omitempty" tf:"weekly_retention"`
	// +optional
	YearlyRetention *string `json:"yearlyRetention,omitempty" tf:"yearly_retention"`
}

type DatabaseSpecShortTermRetentionPolicy struct {
	RetentionDays *int64 `json:"retentionDays" tf:"retention_days"`
}

type DatabaseSpecThreatDetectionPolicy struct {
	// +optional
	DisabledAlerts []string `json:"disabledAlerts,omitempty" tf:"disabled_alerts"`
	// +optional
	EmailAccountAdmins *string `json:"emailAccountAdmins,omitempty" tf:"email_account_admins"`
	// +optional
	EmailAddresses []string `json:"emailAddresses,omitempty" tf:"email_addresses"`
	// +optional
	RetentionDays *int64 `json:"retentionDays,omitempty" tf:"retention_days"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	StorageAccountAccessKey *string `json:"-" sensitive:"true" tf:"storage_account_access_key"`
	// +optional
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint"`
	// +optional
	UseServerDefault *string `json:"useServerDefault,omitempty" tf:"use_server_default"`
}

type DatabaseSpec struct {
	KubeformOutput *DatabaseSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource DatabaseSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type DatabaseSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AutoPauseDelayInMinutes *int64 `json:"autoPauseDelayInMinutes,omitempty" tf:"auto_pause_delay_in_minutes"`
	// +optional
	Collation *string `json:"collation,omitempty" tf:"collation"`
	// +optional
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode"`
	// +optional
	CreationSourceDatabaseID *string `json:"creationSourceDatabaseID,omitempty" tf:"creation_source_database_id"`
	// +optional
	ElasticPoolID *string `json:"elasticPoolID,omitempty" tf:"elastic_pool_id"`
	// +optional
	// Deprecated
	ExtendedAuditingPolicy *DatabaseSpecExtendedAuditingPolicy `json:"extendedAuditingPolicy,omitempty" tf:"extended_auditing_policy"`
	// +optional
	GeoBackupEnabled *bool `json:"geoBackupEnabled,omitempty" tf:"geo_backup_enabled"`
	// +optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type"`
	// +optional
	LongTermRetentionPolicy *DatabaseSpecLongTermRetentionPolicy `json:"longTermRetentionPolicy,omitempty" tf:"long_term_retention_policy"`
	// +optional
	MaxSizeGb *int64 `json:"maxSizeGb,omitempty" tf:"max_size_gb"`
	// +optional
	MinCapacity *float64 `json:"minCapacity,omitempty" tf:"min_capacity"`
	Name        *string  `json:"name" tf:"name"`
	// +optional
	ReadReplicaCount *int64 `json:"readReplicaCount,omitempty" tf:"read_replica_count"`
	// +optional
	ReadScale *bool `json:"readScale,omitempty" tf:"read_scale"`
	// +optional
	RecoverDatabaseID *string `json:"recoverDatabaseID,omitempty" tf:"recover_database_id"`
	// +optional
	RestoreDroppedDatabaseID *string `json:"restoreDroppedDatabaseID,omitempty" tf:"restore_dropped_database_id"`
	// +optional
	RestorePointInTime *string `json:"restorePointInTime,omitempty" tf:"restore_point_in_time"`
	// +optional
	SampleName *string `json:"sampleName,omitempty" tf:"sample_name"`
	ServerID   *string `json:"serverID" tf:"server_id"`
	// +optional
	ShortTermRetentionPolicy *DatabaseSpecShortTermRetentionPolicy `json:"shortTermRetentionPolicy,omitempty" tf:"short_term_retention_policy"`
	// +optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name"`
	// +optional
	StorageAccountType *string `json:"storageAccountType,omitempty" tf:"storage_account_type"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	ThreatDetectionPolicy *DatabaseSpecThreatDetectionPolicy `json:"threatDetectionPolicy,omitempty" tf:"threat_detection_policy"`
	// +optional
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant"`
}

type DatabaseStatus struct {
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

// DatabaseList is a list of Databases
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Database CRD objects
	Items []Database `json:"items,omitempty"`
}
