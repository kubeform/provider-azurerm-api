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

type ManagedHardwareSecurityModule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedHardwareSecurityModuleSpec   `json:"spec,omitempty"`
	Status            ManagedHardwareSecurityModuleStatus `json:"status,omitempty"`
}

type ManagedHardwareSecurityModuleSpec struct {
	State *ManagedHardwareSecurityModuleSpecResource `json:"state,omitempty" tf:"-"`

	Resource ManagedHardwareSecurityModuleSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ManagedHardwareSecurityModuleSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AdminObjectIDS []string `json:"adminObjectIDS" tf:"admin_object_ids"`
	// +optional
	HsmURI   *string `json:"hsmURI,omitempty" tf:"hsm_uri"`
	Location *string `json:"location" tf:"location"`
	Name     *string `json:"name" tf:"name"`
	// +optional
	PurgeProtectionEnabled *bool   `json:"purgeProtectionEnabled,omitempty" tf:"purge_protection_enabled"`
	ResourceGroupName      *string `json:"resourceGroupName" tf:"resource_group_name"`
	SkuName                *string `json:"skuName" tf:"sku_name"`
	// +optional
	SoftDeleteRetentionDays *int64 `json:"softDeleteRetentionDays,omitempty" tf:"soft_delete_retention_days"`
	// +optional
	Tags     *map[string]string `json:"tags,omitempty" tf:"tags"`
	TenantID *string            `json:"tenantID" tf:"tenant_id"`
}

type ManagedHardwareSecurityModuleStatus struct {
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

// ManagedHardwareSecurityModuleList is a list of ManagedHardwareSecurityModules
type ManagedHardwareSecurityModuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ManagedHardwareSecurityModule CRD objects
	Items []ManagedHardwareSecurityModule `json:"items,omitempty"`
}
