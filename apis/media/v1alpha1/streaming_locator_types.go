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

type StreamingLocator struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamingLocatorSpec   `json:"spec,omitempty"`
	Status            StreamingLocatorStatus `json:"status,omitempty"`
}

type StreamingLocatorSpecContentKey struct {
	// +optional
	ContentKeyID *string `json:"contentKeyID,omitempty" tf:"content_key_id"`
	// +optional
	LabelReferenceInStreamingPolicy *string `json:"labelReferenceInStreamingPolicy,omitempty" tf:"label_reference_in_streaming_policy"`
	// +optional
	PolicyName *string `json:"policyName,omitempty" tf:"policy_name"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type StreamingLocatorSpec struct {
	State *StreamingLocatorSpecResource `json:"state,omitempty" tf:"-"`

	Resource StreamingLocatorSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type StreamingLocatorSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AlternativeMediaID *string `json:"alternativeMediaID,omitempty" tf:"alternative_media_id"`
	AssetName          *string `json:"assetName" tf:"asset_name"`
	// +optional
	ContentKey []StreamingLocatorSpecContentKey `json:"contentKey,omitempty" tf:"content_key"`
	// +optional
	DefaultContentKeyPolicyName *string `json:"defaultContentKeyPolicyName,omitempty" tf:"default_content_key_policy_name"`
	// +optional
	EndTime                  *string `json:"endTime,omitempty" tf:"end_time"`
	MediaServicesAccountName *string `json:"mediaServicesAccountName" tf:"media_services_account_name"`
	Name                     *string `json:"name" tf:"name"`
	ResourceGroupName        *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time"`
	// +optional
	StreamingLocatorID  *string `json:"streamingLocatorID,omitempty" tf:"streaming_locator_id"`
	StreamingPolicyName *string `json:"streamingPolicyName" tf:"streaming_policy_name"`
}

type StreamingLocatorStatus struct {
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

// StreamingLocatorList is a list of StreamingLocators
type StreamingLocatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StreamingLocator CRD objects
	Items []StreamingLocator `json:"items,omitempty"`
}
