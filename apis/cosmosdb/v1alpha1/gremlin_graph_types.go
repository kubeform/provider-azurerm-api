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

type GremlinGraph struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GremlinGraphSpec   `json:"spec,omitempty"`
	Status            GremlinGraphStatus `json:"status,omitempty"`
}

type GremlinGraphSpecAutoscaleSettings struct {
	// +optional
	MaxThroughput *int64 `json:"maxThroughput,omitempty" tf:"max_throughput"`
}

type GremlinGraphSpecConflictResolutionPolicy struct {
	// +optional
	ConflictResolutionPath *string `json:"conflictResolutionPath,omitempty" tf:"conflict_resolution_path"`
	// +optional
	ConflictResolutionProcedure *string `json:"conflictResolutionProcedure,omitempty" tf:"conflict_resolution_procedure"`
	Mode                        *string `json:"mode" tf:"mode"`
}

type GremlinGraphSpecIndexPolicyCompositeIndexIndex struct {
	Order *string `json:"order" tf:"order"`
	Path  *string `json:"path" tf:"path"`
}

type GremlinGraphSpecIndexPolicyCompositeIndex struct {
	// +kubebuilder:validation:MinItems=1
	Index []GremlinGraphSpecIndexPolicyCompositeIndexIndex `json:"index" tf:"index"`
}

type GremlinGraphSpecIndexPolicySpatialIndex struct {
	Path *string `json:"path" tf:"path"`
	// +optional
	Types []string `json:"types,omitempty" tf:"types"`
}

type GremlinGraphSpecIndexPolicy struct {
	// +optional
	Automatic *bool `json:"automatic,omitempty" tf:"automatic"`
	// +optional
	CompositeIndex []GremlinGraphSpecIndexPolicyCompositeIndex `json:"compositeIndex,omitempty" tf:"composite_index"`
	// +optional
	ExcludedPaths []string `json:"excludedPaths,omitempty" tf:"excluded_paths"`
	// +optional
	IncludedPaths []string `json:"includedPaths,omitempty" tf:"included_paths"`
	IndexingMode  *string  `json:"indexingMode" tf:"indexing_mode"`
	// +optional
	SpatialIndex []GremlinGraphSpecIndexPolicySpatialIndex `json:"spatialIndex,omitempty" tf:"spatial_index"`
}

type GremlinGraphSpecUniqueKey struct {
	Paths []string `json:"paths" tf:"paths"`
}

type GremlinGraphSpec struct {
	KubeformOutput *GremlinGraphSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource GremlinGraphSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type GremlinGraphSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AccountName *string `json:"accountName" tf:"account_name"`
	// +optional
	AutoscaleSettings *GremlinGraphSpecAutoscaleSettings `json:"autoscaleSettings,omitempty" tf:"autoscale_settings"`
	// +optional
	ConflictResolutionPolicy *GremlinGraphSpecConflictResolutionPolicy `json:"conflictResolutionPolicy,omitempty" tf:"conflict_resolution_policy"`
	DatabaseName             *string                                   `json:"databaseName" tf:"database_name"`
	// +optional
	DefaultTtl *int64 `json:"defaultTtl,omitempty" tf:"default_ttl"`
	// +optional
	IndexPolicy      *GremlinGraphSpecIndexPolicy `json:"indexPolicy,omitempty" tf:"index_policy"`
	Name             *string                      `json:"name" tf:"name"`
	PartitionKeyPath *string                      `json:"partitionKeyPath" tf:"partition_key_path"`
	// +optional
	PartitionKeyVersion *int64  `json:"partitionKeyVersion,omitempty" tf:"partition_key_version"`
	ResourceGroupName   *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Throughput *int64 `json:"throughput,omitempty" tf:"throughput"`
	// +optional
	UniqueKey []GremlinGraphSpecUniqueKey `json:"uniqueKey,omitempty" tf:"unique_key"`
}

type GremlinGraphStatus struct {
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

// GremlinGraphList is a list of GremlinGraphs
type GremlinGraphList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GremlinGraph CRD objects
	Items []GremlinGraph `json:"items,omitempty"`
}