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

type Assignment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AssignmentSpec   `json:"spec,omitempty"`
	Status            AssignmentStatus `json:"status,omitempty"`
}

type AssignmentSpecIdentity struct {
	// +optional
	PrincipalID *string `json:"principalID,omitempty" tf:"principal_id"`
	// +optional
	TenantID *string `json:"tenantID,omitempty" tf:"tenant_id"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type AssignmentSpec struct {
	KubeformOutput *AssignmentSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource AssignmentSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type AssignmentSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	EnforcementMode *bool `json:"enforcementMode,omitempty" tf:"enforcement_mode"`
	// +optional
	Identity *AssignmentSpecIdentity `json:"identity,omitempty" tf:"identity"`
	// +optional
	Location *string `json:"location,omitempty" tf:"location"`
	// +optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata"`
	Name     *string `json:"name" tf:"name"`
	// +optional
	NotScopes []string `json:"notScopes,omitempty" tf:"not_scopes"`
	// +optional
	Parameters         *string `json:"parameters,omitempty" tf:"parameters"`
	PolicyDefinitionID *string `json:"policyDefinitionID" tf:"policy_definition_id"`
	Scope              *string `json:"scope" tf:"scope"`
}

type AssignmentStatus struct {
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

// AssignmentList is a list of Assignments
type AssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Assignment CRD objects
	Items []Assignment `json:"items,omitempty"`
}
