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

type PacketCapture struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PacketCaptureSpec   `json:"spec,omitempty"`
	Status            PacketCaptureStatus `json:"status,omitempty"`
}

type PacketCaptureSpecFilter struct {
	// +optional
	LocalIPAddress *string `json:"localIPAddress,omitempty" tf:"local_ip_address"`
	// +optional
	LocalPort *string `json:"localPort,omitempty" tf:"local_port"`
	Protocol  *string `json:"protocol" tf:"protocol"`
	// +optional
	RemoteIPAddress *string `json:"remoteIPAddress,omitempty" tf:"remote_ip_address"`
	// +optional
	RemotePort *string `json:"remotePort,omitempty" tf:"remote_port"`
}

type PacketCaptureSpecStorageLocation struct {
	// +optional
	FilePath *string `json:"filePath,omitempty" tf:"file_path"`
	// +optional
	StorageAccountID *string `json:"storageAccountID,omitempty" tf:"storage_account_id"`
	// +optional
	StoragePath *string `json:"storagePath,omitempty" tf:"storage_path"`
}

type PacketCaptureSpec struct {
	State *PacketCaptureSpecResource `json:"state,omitempty" tf:"-"`

	Resource PacketCaptureSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type PacketCaptureSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Filter []PacketCaptureSpecFilter `json:"filter,omitempty" tf:"filter"`
	// +optional
	MaximumBytesPerPacket *int64 `json:"maximumBytesPerPacket,omitempty" tf:"maximum_bytes_per_packet"`
	// +optional
	MaximumBytesPerSession *int64 `json:"maximumBytesPerSession,omitempty" tf:"maximum_bytes_per_session"`
	// +optional
	MaximumCaptureDuration *int64                            `json:"maximumCaptureDuration,omitempty" tf:"maximum_capture_duration"`
	Name                   *string                           `json:"name" tf:"name"`
	NetworkWatcherName     *string                           `json:"networkWatcherName" tf:"network_watcher_name"`
	ResourceGroupName      *string                           `json:"resourceGroupName" tf:"resource_group_name"`
	StorageLocation        *PacketCaptureSpecStorageLocation `json:"storageLocation" tf:"storage_location"`
	TargetResourceID       *string                           `json:"targetResourceID" tf:"target_resource_id"`
}

type PacketCaptureStatus struct {
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

// PacketCaptureList is a list of PacketCaptures
type PacketCaptureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PacketCapture CRD objects
	Items []PacketCapture `json:"items,omitempty"`
}
