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

type CenterAssessmentPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CenterAssessmentPolicySpec   `json:"spec,omitempty"`
	Status            CenterAssessmentPolicyStatus `json:"status,omitempty"`
}

type CenterAssessmentPolicySpec struct {
	State *CenterAssessmentPolicySpecResource `json:"state,omitempty" tf:"-"`

	Resource CenterAssessmentPolicySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type CenterAssessmentPolicySpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Categories  []string `json:"categories,omitempty" tf:"categories"`
	Description *string  `json:"description" tf:"description"`
	DisplayName *string  `json:"displayName" tf:"display_name"`
	// +optional
	ImplementationEffort *string `json:"implementationEffort,omitempty" tf:"implementation_effort"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	RemediationDescription *string `json:"remediationDescription,omitempty" tf:"remediation_description"`
	// +optional
	Severity *string `json:"severity,omitempty" tf:"severity"`
	// +optional
	Threats []string `json:"threats,omitempty" tf:"threats"`
	// +optional
	UserImpact *string `json:"userImpact,omitempty" tf:"user_impact"`
}

type CenterAssessmentPolicyStatus struct {
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

// CenterAssessmentPolicyList is a list of CenterAssessmentPolicys
type CenterAssessmentPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CenterAssessmentPolicy CRD objects
	Items []CenterAssessmentPolicy `json:"items,omitempty"`
}
