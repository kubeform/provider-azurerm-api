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

type EndpointStorageContainer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointStorageContainerSpec   `json:"spec,omitempty"`
	Status            EndpointStorageContainerStatus `json:"status,omitempty"`
}

type EndpointStorageContainerSpec struct {
	State *EndpointStorageContainerSpecResource `json:"state,omitempty" tf:"-"`

	Resource EndpointStorageContainerSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type EndpointStorageContainerSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type"`
	// +optional
	BatchFrequencyInSeconds *int64 `json:"batchFrequencyInSeconds,omitempty" tf:"batch_frequency_in_seconds"`
	// +optional
	ConnectionString *string `json:"-" sensitive:"true" tf:"connection_string"`
	ContainerName    *string `json:"containerName" tf:"container_name"`
	// +optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding"`
	// +optional
	EndpointURI *string `json:"endpointURI,omitempty" tf:"endpoint_uri"`
	// +optional
	FileNameFormat *string `json:"fileNameFormat,omitempty" tf:"file_name_format"`
	// +optional
	IdentityID *string `json:"identityID,omitempty" tf:"identity_id"`
	// +optional
	IothubID *string `json:"iothubID,omitempty" tf:"iothub_id"`
	// +optional
	// Deprecated
	IothubName *string `json:"iothubName,omitempty" tf:"iothub_name"`
	// +optional
	MaxChunkSizeInBytes *int64  `json:"maxChunkSizeInBytes,omitempty" tf:"max_chunk_size_in_bytes"`
	Name                *string `json:"name" tf:"name"`
	ResourceGroupName   *string `json:"resourceGroupName" tf:"resource_group_name"`
}

type EndpointStorageContainerStatus struct {
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

// EndpointStorageContainerList is a list of EndpointStorageContainers
type EndpointStorageContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EndpointStorageContainer CRD objects
	Items []EndpointStorageContainer `json:"items,omitempty"`
}
