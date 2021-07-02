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

type SqlPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlPoolSpec   `json:"spec,omitempty"`
	Status            SqlPoolStatus `json:"status,omitempty"`
}

type SqlPoolSpecRestore struct {
	PointInTime      *string `json:"pointInTime" tf:"point_in_time"`
	SourceDatabaseID *string `json:"sourceDatabaseID" tf:"source_database_id"`
}

type SqlPoolSpec struct {
	KubeformOutput *SqlPoolSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource SqlPoolSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type SqlPoolSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Collation *string `json:"collation,omitempty" tf:"collation"`
	// +optional
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode"`
	// +optional
	DataEncrypted *bool   `json:"dataEncrypted,omitempty" tf:"data_encrypted"`
	Name          *string `json:"name" tf:"name"`
	// +optional
	RecoveryDatabaseID *string `json:"recoveryDatabaseID,omitempty" tf:"recovery_database_id"`
	// +optional
	Restore            *SqlPoolSpecRestore `json:"restore,omitempty" tf:"restore"`
	SkuName            *string             `json:"skuName" tf:"sku_name"`
	SynapseWorkspaceID *string             `json:"synapseWorkspaceID" tf:"synapse_workspace_id"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type SqlPoolStatus struct {
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

// SqlPoolList is a list of SqlPools
type SqlPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SqlPool CRD objects
	Items []SqlPool `json:"items,omitempty"`
}