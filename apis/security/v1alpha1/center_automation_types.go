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

type CenterAutomation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CenterAutomationSpec   `json:"spec,omitempty"`
	Status            CenterAutomationStatus `json:"status,omitempty"`
}

type CenterAutomationSpecAction struct {
	// +optional
	ConnectionString *string `json:"-" sensitive:"true" tf:"connection_string"`
	ResourceID       *string `json:"resourceID" tf:"resource_id"`
	// +optional
	TriggerURL *string `json:"-" sensitive:"true" tf:"trigger_url"`
	Type       *string `json:"type" tf:"type"`
}

type CenterAutomationSpecSourceRuleSetRule struct {
	ExpectedValue *string `json:"expectedValue" tf:"expected_value"`
	Operator      *string `json:"operator" tf:"operator"`
	PropertyPath  *string `json:"propertyPath" tf:"property_path"`
	PropertyType  *string `json:"propertyType" tf:"property_type"`
}

type CenterAutomationSpecSourceRuleSet struct {
	Rule []CenterAutomationSpecSourceRuleSetRule `json:"rule" tf:"rule"`
}

type CenterAutomationSpecSource struct {
	EventSource *string `json:"eventSource" tf:"event_source"`
	// +optional
	RuleSet []CenterAutomationSpecSourceRuleSet `json:"ruleSet,omitempty" tf:"rule_set"`
}

type CenterAutomationSpec struct {
	KubeformOutput *CenterAutomationSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource CenterAutomationSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type CenterAutomationSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:MinItems=1
	Action []CenterAutomationSpecAction `json:"action" tf:"action"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Enabled           *bool   `json:"enabled,omitempty" tf:"enabled"`
	Location          *string `json:"location" tf:"location"`
	Name              *string `json:"name" tf:"name"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MinItems=1
	Scopes []string `json:"scopes" tf:"scopes"`
	// +kubebuilder:validation:MinItems=1
	Source []CenterAutomationSpecSource `json:"source" tf:"source"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type CenterAutomationStatus struct {
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

// CenterAutomationList is a list of CenterAutomations
type CenterAutomationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CenterAutomation CRD objects
	Items []CenterAutomation `json:"items,omitempty"`
}
