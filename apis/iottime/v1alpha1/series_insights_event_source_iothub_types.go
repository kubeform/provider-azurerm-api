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

type SeriesInsightsEventSourceIothub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SeriesInsightsEventSourceIothubSpec   `json:"spec,omitempty"`
	Status            SeriesInsightsEventSourceIothubStatus `json:"status,omitempty"`
}

type SeriesInsightsEventSourceIothubSpec struct {
	State *SeriesInsightsEventSourceIothubSpecResource `json:"state,omitempty" tf:"-"`

	Resource SeriesInsightsEventSourceIothubSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SeriesInsightsEventSourceIothubSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ConsumerGroupName     *string `json:"consumerGroupName" tf:"consumer_group_name"`
	EnvironmentID         *string `json:"environmentID" tf:"environment_id"`
	EventSourceResourceID *string `json:"eventSourceResourceID" tf:"event_source_resource_id"`
	IothubName            *string `json:"iothubName" tf:"iothub_name"`
	Location              *string `json:"location" tf:"location"`
	Name                  *string `json:"name" tf:"name"`
	SharedAccessKey       *string `json:"sharedAccessKey" tf:"shared_access_key"`
	SharedAccessKeyName   *string `json:"sharedAccessKeyName" tf:"shared_access_key_name"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TimestampPropertyName *string `json:"timestampPropertyName,omitempty" tf:"timestamp_property_name"`
}

type SeriesInsightsEventSourceIothubStatus struct {
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

// SeriesInsightsEventSourceIothubList is a list of SeriesInsightsEventSourceIothubs
type SeriesInsightsEventSourceIothubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SeriesInsightsEventSourceIothub CRD objects
	Items []SeriesInsightsEventSourceIothub `json:"items,omitempty"`
}
