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

type SqlPoolWorkloadGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlPoolWorkloadGroupSpec   `json:"spec,omitempty"`
	Status            SqlPoolWorkloadGroupStatus `json:"status,omitempty"`
}

type SqlPoolWorkloadGroupSpec struct {
	State *SqlPoolWorkloadGroupSpecResource `json:"state,omitempty" tf:"-"`

	Resource SqlPoolWorkloadGroupSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SqlPoolWorkloadGroupSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Importance         *string `json:"importance,omitempty" tf:"importance"`
	MaxResourcePercent *int64  `json:"maxResourcePercent" tf:"max_resource_percent"`
	// +optional
	MaxResourcePercentPerRequest *float64 `json:"maxResourcePercentPerRequest,omitempty" tf:"max_resource_percent_per_request"`
	MinResourcePercent           *int64   `json:"minResourcePercent" tf:"min_resource_percent"`
	// +optional
	MinResourcePercentPerRequest *float64 `json:"minResourcePercentPerRequest,omitempty" tf:"min_resource_percent_per_request"`
	Name                         *string  `json:"name" tf:"name"`
	// +optional
	QueryExecutionTimeoutInSeconds *int64  `json:"queryExecutionTimeoutInSeconds,omitempty" tf:"query_execution_timeout_in_seconds"`
	SqlPoolID                      *string `json:"sqlPoolID" tf:"sql_pool_id"`
}

type SqlPoolWorkloadGroupStatus struct {
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

// SqlPoolWorkloadGroupList is a list of SqlPoolWorkloadGroups
type SqlPoolWorkloadGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SqlPoolWorkloadGroup CRD objects
	Items []SqlPoolWorkloadGroup `json:"items,omitempty"`
}
