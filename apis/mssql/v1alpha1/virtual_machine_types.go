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

type VirtualMachine struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualMachineSpec   `json:"spec,omitempty"`
	Status            VirtualMachineStatus `json:"status,omitempty"`
}

type VirtualMachineSpecAutoBackupManualSchedule struct {
	FullBackupFrequency         *string `json:"fullBackupFrequency" tf:"full_backup_frequency"`
	FullBackupStartHour         *int64  `json:"fullBackupStartHour" tf:"full_backup_start_hour"`
	FullBackupWindowInHours     *int64  `json:"fullBackupWindowInHours" tf:"full_backup_window_in_hours"`
	LogBackupFrequencyInMinutes *int64  `json:"logBackupFrequencyInMinutes" tf:"log_backup_frequency_in_minutes"`
}

type VirtualMachineSpecAutoBackup struct {
	// +optional
	EncryptionEnabled *bool `json:"encryptionEnabled,omitempty" tf:"encryption_enabled"`
	// +optional
	EncryptionPassword *string `json:"-" sensitive:"true" tf:"encryption_password"`
	// +optional
	ManualSchedule          *VirtualMachineSpecAutoBackupManualSchedule `json:"manualSchedule,omitempty" tf:"manual_schedule"`
	RetentionPeriodInDays   *int64                                      `json:"retentionPeriodInDays" tf:"retention_period_in_days"`
	StorageAccountAccessKey *string                                     `json:"storageAccountAccessKey" tf:"storage_account_access_key"`
	StorageBlobEndpoint     *string                                     `json:"storageBlobEndpoint" tf:"storage_blob_endpoint"`
	// +optional
	SystemDatabasesBackupEnabled *bool `json:"systemDatabasesBackupEnabled,omitempty" tf:"system_databases_backup_enabled"`
}

type VirtualMachineSpecAutoPatching struct {
	DayOfWeek                          *string `json:"dayOfWeek" tf:"day_of_week"`
	MaintenanceWindowDurationInMinutes *int64  `json:"maintenanceWindowDurationInMinutes" tf:"maintenance_window_duration_in_minutes"`
	MaintenanceWindowStartingHour      *int64  `json:"maintenanceWindowStartingHour" tf:"maintenance_window_starting_hour"`
}

type VirtualMachineSpecKeyVaultCredential struct {
	KeyVaultURL            *string `json:"-" sensitive:"true" tf:"key_vault_url"`
	Name                   *string `json:"name" tf:"name"`
	ServicePrincipalName   *string `json:"-" sensitive:"true" tf:"service_principal_name"`
	ServicePrincipalSecret *string `json:"-" sensitive:"true" tf:"service_principal_secret"`
}

type VirtualMachineSpecStorageConfigurationDataSettings struct {
	DefaultFilePath *string `json:"defaultFilePath" tf:"default_file_path"`
	Luns            []int64 `json:"luns" tf:"luns"`
}

type VirtualMachineSpecStorageConfigurationLogSettings struct {
	DefaultFilePath *string `json:"defaultFilePath" tf:"default_file_path"`
	Luns            []int64 `json:"luns" tf:"luns"`
}

type VirtualMachineSpecStorageConfigurationTempDbSettings struct {
	DefaultFilePath *string `json:"defaultFilePath" tf:"default_file_path"`
	Luns            []int64 `json:"luns" tf:"luns"`
}

type VirtualMachineSpecStorageConfiguration struct {
	// +optional
	DataSettings *VirtualMachineSpecStorageConfigurationDataSettings `json:"dataSettings,omitempty" tf:"data_settings"`
	DiskType     *string                                             `json:"diskType" tf:"disk_type"`
	// +optional
	LogSettings         *VirtualMachineSpecStorageConfigurationLogSettings `json:"logSettings,omitempty" tf:"log_settings"`
	StorageWorkloadType *string                                            `json:"storageWorkloadType" tf:"storage_workload_type"`
	// +optional
	TempDbSettings *VirtualMachineSpecStorageConfigurationTempDbSettings `json:"tempDbSettings,omitempty" tf:"temp_db_settings"`
}

type VirtualMachineSpec struct {
	State *VirtualMachineSpecResource `json:"state,omitempty" tf:"-"`

	Resource VirtualMachineSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type VirtualMachineSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AutoBackup *VirtualMachineSpecAutoBackup `json:"autoBackup,omitempty" tf:"auto_backup"`
	// +optional
	AutoPatching *VirtualMachineSpecAutoPatching `json:"autoPatching,omitempty" tf:"auto_patching"`
	// +optional
	KeyVaultCredential *VirtualMachineSpecKeyVaultCredential `json:"keyVaultCredential,omitempty" tf:"key_vault_credential"`
	// +optional
	RServicesEnabled *bool `json:"rServicesEnabled,omitempty" tf:"r_services_enabled"`
	// +optional
	SqlConnectivityPort *int64 `json:"sqlConnectivityPort,omitempty" tf:"sql_connectivity_port"`
	// +optional
	SqlConnectivityType *string `json:"sqlConnectivityType,omitempty" tf:"sql_connectivity_type"`
	// +optional
	SqlConnectivityUpdatePassword *string `json:"-" sensitive:"true" tf:"sql_connectivity_update_password"`
	// +optional
	SqlConnectivityUpdateUsername *string `json:"-" sensitive:"true" tf:"sql_connectivity_update_username"`
	SqlLicenseType                *string `json:"sqlLicenseType" tf:"sql_license_type"`
	// +optional
	StorageConfiguration *VirtualMachineSpecStorageConfiguration `json:"storageConfiguration,omitempty" tf:"storage_configuration"`
	// +optional
	Tags             *map[string]string `json:"tags,omitempty" tf:"tags"`
	VirtualMachineID *string            `json:"virtualMachineID" tf:"virtual_machine_id"`
}

type VirtualMachineStatus struct {
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

// VirtualMachineList is a list of VirtualMachines
type VirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VirtualMachine CRD objects
	Items []VirtualMachine `json:"items,omitempty"`
}
