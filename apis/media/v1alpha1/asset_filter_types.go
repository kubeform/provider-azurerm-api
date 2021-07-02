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

type AssetFilter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AssetFilterSpec   `json:"spec,omitempty"`
	Status            AssetFilterStatus `json:"status,omitempty"`
}

type AssetFilterSpecPresentationTimeRange struct {
	// +optional
	EndInUnits *int64 `json:"endInUnits,omitempty" tf:"end_in_units"`
	// +optional
	ForceEnd *bool `json:"forceEnd,omitempty" tf:"force_end"`
	// +optional
	LiveBackoffInUnits *int64 `json:"liveBackoffInUnits,omitempty" tf:"live_backoff_in_units"`
	// +optional
	PresentationWindowInUnits *int64 `json:"presentationWindowInUnits,omitempty" tf:"presentation_window_in_units"`
	// +optional
	StartInUnits *int64 `json:"startInUnits,omitempty" tf:"start_in_units"`
	// +optional
	UnitTimescaleInMiliseconds *int64 `json:"unitTimescaleInMiliseconds,omitempty" tf:"unit_timescale_in_miliseconds"`
}

type AssetFilterSpecTrackSelectionCondition struct {
	// +optional
	Operation *string `json:"operation,omitempty" tf:"operation"`
	// +optional
	Property *string `json:"property,omitempty" tf:"property"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type AssetFilterSpecTrackSelection struct {
	Condition []AssetFilterSpecTrackSelectionCondition `json:"condition" tf:"condition"`
}

type AssetFilterSpec struct {
	KubeformOutput *AssetFilterSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource AssetFilterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type AssetFilterSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AssetID *string `json:"assetID" tf:"asset_id"`
	// +optional
	FirstQualityBitrate *int64  `json:"firstQualityBitrate,omitempty" tf:"first_quality_bitrate"`
	Name                *string `json:"name" tf:"name"`
	// +optional
	PresentationTimeRange *AssetFilterSpecPresentationTimeRange `json:"presentationTimeRange,omitempty" tf:"presentation_time_range"`
	// +optional
	TrackSelection []AssetFilterSpecTrackSelection `json:"trackSelection,omitempty" tf:"track_selection"`
}

type AssetFilterStatus struct {
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

// AssetFilterList is a list of AssetFilters
type AssetFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AssetFilter CRD objects
	Items []AssetFilter `json:"items,omitempty"`
}