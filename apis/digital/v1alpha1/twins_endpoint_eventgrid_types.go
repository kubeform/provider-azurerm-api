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

type TwinsEndpointEventgrid struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TwinsEndpointEventgridSpec   `json:"spec,omitempty"`
	Status            TwinsEndpointEventgridStatus `json:"status,omitempty"`
}

type TwinsEndpointEventgridSpec struct {
	KubeformOutput *TwinsEndpointEventgridSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource TwinsEndpointEventgridSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type TwinsEndpointEventgridSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	DeadLetterStorageSecret          *string `json:"deadLetterStorageSecret,omitempty" tf:"dead_letter_storage_secret"`
	DigitalTwinsID                   *string `json:"digitalTwinsID" tf:"digital_twins_id"`
	EventgridTopicEndpoint           *string `json:"eventgridTopicEndpoint" tf:"eventgrid_topic_endpoint"`
	EventgridTopicPrimaryAccessKey   *string `json:"eventgridTopicPrimaryAccessKey" tf:"eventgrid_topic_primary_access_key"`
	EventgridTopicSecondaryAccessKey *string `json:"eventgridTopicSecondaryAccessKey" tf:"eventgrid_topic_secondary_access_key"`
	Name                             *string `json:"name" tf:"name"`
}

type TwinsEndpointEventgridStatus struct {
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

// TwinsEndpointEventgridList is a list of TwinsEndpointEventgrids
type TwinsEndpointEventgridList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of TwinsEndpointEventgrid CRD objects
	Items []TwinsEndpointEventgrid `json:"items,omitempty"`
}
