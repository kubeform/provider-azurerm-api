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

type DpsSharedAccessPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DpsSharedAccessPolicySpec   `json:"spec,omitempty"`
	Status            DpsSharedAccessPolicyStatus `json:"status,omitempty"`
}

type DpsSharedAccessPolicySpec struct {
	State *DpsSharedAccessPolicySpecResource `json:"state,omitempty" tf:"-"`

	Resource DpsSharedAccessPolicySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type DpsSharedAccessPolicySpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	EnrollmentRead *bool `json:"enrollmentRead,omitempty" tf:"enrollment_read"`
	// +optional
	EnrollmentWrite *bool   `json:"enrollmentWrite,omitempty" tf:"enrollment_write"`
	IothubDpsName   *string `json:"iothubDpsName" tf:"iothub_dps_name"`
	Name            *string `json:"name" tf:"name"`
	// +optional
	PrimaryConnectionString *string `json:"-" sensitive:"true" tf:"primary_connection_string"`
	// +optional
	PrimaryKey *string `json:"-" sensitive:"true" tf:"primary_key"`
	// +optional
	RegistrationRead *bool `json:"registrationRead,omitempty" tf:"registration_read"`
	// +optional
	RegistrationWrite *bool   `json:"registrationWrite,omitempty" tf:"registration_write"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryConnectionString *string `json:"-" sensitive:"true" tf:"secondary_connection_string"`
	// +optional
	SecondaryKey *string `json:"-" sensitive:"true" tf:"secondary_key"`
	// +optional
	ServiceConfig *bool `json:"serviceConfig,omitempty" tf:"service_config"`
}

type DpsSharedAccessPolicyStatus struct {
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

// DpsSharedAccessPolicyList is a list of DpsSharedAccessPolicys
type DpsSharedAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DpsSharedAccessPolicy CRD objects
	Items []DpsSharedAccessPolicy `json:"items,omitempty"`
}
