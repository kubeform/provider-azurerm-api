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

type DatabasePrincipalAssignment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabasePrincipalAssignmentSpec   `json:"spec,omitempty"`
	Status            DatabasePrincipalAssignmentStatus `json:"status,omitempty"`
}

type DatabasePrincipalAssignmentSpec struct {
	State *DatabasePrincipalAssignmentSpecResource `json:"state,omitempty" tf:"-"`

	Resource DatabasePrincipalAssignmentSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DatabasePrincipalAssignmentSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ClusterName  *string `json:"clusterName" tf:"cluster_name"`
	DatabaseName *string `json:"databaseName" tf:"database_name"`
	Name         *string `json:"name" tf:"name"`
	PrincipalID  *string `json:"principalID" tf:"principal_id"`
	// +optional
	PrincipalName     *string `json:"principalName,omitempty" tf:"principal_name"`
	PrincipalType     *string `json:"principalType" tf:"principal_type"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	Role              *string `json:"role" tf:"role"`
	TenantID          *string `json:"tenantID" tf:"tenant_id"`
	// +optional
	TenantName *string `json:"tenantName,omitempty" tf:"tenant_name"`
}

type DatabasePrincipalAssignmentStatus struct {
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

// DatabasePrincipalAssignmentList is a list of DatabasePrincipalAssignments
type DatabasePrincipalAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DatabasePrincipalAssignment CRD objects
	Items []DatabasePrincipalAssignment `json:"items,omitempty"`
}
