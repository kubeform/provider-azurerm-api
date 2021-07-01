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

type ActionHTTP struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ActionHTTPSpec   `json:"spec,omitempty"`
	Status            ActionHTTPStatus `json:"status,omitempty"`
}

type ActionHTTPSpecRunAfter struct {
	ActionName   *string `json:"actionName" tf:"action_name"`
	ActionResult *string `json:"actionResult" tf:"action_result"`
}

type ActionHTTPSpec struct {
	KubeformOutput *ActionHTTPSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource ActionHTTPSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ActionHTTPSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Body *string `json:"body,omitempty" tf:"body"`
	// +optional
	Headers    *map[string]string `json:"headers,omitempty" tf:"headers"`
	LogicAppID *string            `json:"logicAppID" tf:"logic_app_id"`
	Method     *string            `json:"method" tf:"method"`
	Name       *string            `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	RunAfter []ActionHTTPSpecRunAfter `json:"runAfter,omitempty" tf:"run_after"`
	Uri      *string                  `json:"uri" tf:"uri"`
}

type ActionHTTPStatus struct {
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

// ActionHTTPList is a list of ActionHTTPs
type ActionHTTPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ActionHTTP CRD objects
	Items []ActionHTTP `json:"items,omitempty"`
}
