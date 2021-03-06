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

type Runbook struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RunbookSpec   `json:"spec,omitempty"`
	Status            RunbookStatus `json:"status,omitempty"`
}

type RunbookSpecJobSchedule struct {
	// +optional
	JobScheduleID *string `json:"jobScheduleID,omitempty" tf:"job_schedule_id"`
	// +optional
	Parameters *map[string]string `json:"parameters,omitempty" tf:"parameters"`
	// +optional
	RunOn        *string `json:"runOn,omitempty" tf:"run_on"`
	ScheduleName *string `json:"scheduleName" tf:"schedule_name"`
}

type RunbookSpecPublishContentLinkHash struct {
	Algorithm *string `json:"algorithm" tf:"algorithm"`
	Value     *string `json:"value" tf:"value"`
}

type RunbookSpecPublishContentLink struct {
	// +optional
	Hash *RunbookSpecPublishContentLinkHash `json:"hash,omitempty" tf:"hash"`
	Uri  *string                            `json:"uri" tf:"uri"`
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
}

type RunbookSpec struct {
	State *RunbookSpecResource `json:"state,omitempty" tf:"-"`

	Resource RunbookSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type RunbookSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AutomationAccountName *string `json:"automationAccountName" tf:"automation_account_name"`
	// +optional
	Content *string `json:"content,omitempty" tf:"content"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	JobSchedule []RunbookSpecJobSchedule `json:"jobSchedule,omitempty" tf:"job_schedule"`
	Location    *string                  `json:"location" tf:"location"`
	LogProgress *bool                    `json:"logProgress" tf:"log_progress"`
	LogVerbose  *bool                    `json:"logVerbose" tf:"log_verbose"`
	Name        *string                  `json:"name" tf:"name"`
	// +optional
	PublishContentLink *RunbookSpecPublishContentLink `json:"publishContentLink,omitempty" tf:"publish_content_link"`
	ResourceGroupName  *string                        `json:"resourceGroupName" tf:"resource_group_name"`
	RunbookType        *string                        `json:"runbookType" tf:"runbook_type"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type RunbookStatus struct {
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

// RunbookList is a list of Runbooks
type RunbookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Runbook CRD objects
	Items []Runbook `json:"items,omitempty"`
}
