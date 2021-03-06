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

type AlertRuleMsSecurityIncident struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlertRuleMsSecurityIncidentSpec   `json:"spec,omitempty"`
	Status            AlertRuleMsSecurityIncidentStatus `json:"status,omitempty"`
}

type AlertRuleMsSecurityIncidentSpec struct {
	State *AlertRuleMsSecurityIncidentSpecResource `json:"state,omitempty" tf:"-"`

	Resource AlertRuleMsSecurityIncidentSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type AlertRuleMsSecurityIncidentSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AlertRuleTemplateGuid *string `json:"alertRuleTemplateGuid,omitempty" tf:"alert_rule_template_guid"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	DisplayName *string `json:"displayName" tf:"display_name"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	DisplayNameExcludeFilter []string `json:"displayNameExcludeFilter,omitempty" tf:"display_name_exclude_filter"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	DisplayNameFilter []string `json:"displayNameFilter,omitempty" tf:"display_name_filter"`
	// +optional
	Enabled                 *bool   `json:"enabled,omitempty" tf:"enabled"`
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceID" tf:"log_analytics_workspace_id"`
	Name                    *string `json:"name" tf:"name"`
	ProductFilter           *string `json:"productFilter" tf:"product_filter"`
	// +kubebuilder:validation:MinItems=1
	SeverityFilter []string `json:"severityFilter" tf:"severity_filter"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	// Deprecated
	TextWhitelist []string `json:"textWhitelist,omitempty" tf:"text_whitelist"`
}

type AlertRuleMsSecurityIncidentStatus struct {
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

// AlertRuleMsSecurityIncidentList is a list of AlertRuleMsSecurityIncidents
type AlertRuleMsSecurityIncidentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AlertRuleMsSecurityIncident CRD objects
	Items []AlertRuleMsSecurityIncident `json:"items,omitempty"`
}
