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

type EdgeDevice struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EdgeDeviceSpec   `json:"spec,omitempty"`
	Status            EdgeDeviceStatus `json:"status,omitempty"`
}

type EdgeDeviceSpecDeviceProperties struct {
	// +optional
	Capacity *int64 `json:"capacity,omitempty" tf:"capacity"`
	// +optional
	ConfiguredRoleTypes []string `json:"configuredRoleTypes,omitempty" tf:"configured_role_types"`
	// +optional
	Culture *string `json:"culture,omitempty" tf:"culture"`
	// +optional
	HcsVersion *string `json:"hcsVersion,omitempty" tf:"hcs_version"`
	// +optional
	Model *string `json:"model,omitempty" tf:"model"`
	// +optional
	NodeCount *int64 `json:"nodeCount,omitempty" tf:"node_count"`
	// +optional
	SerialNumber *string `json:"serialNumber,omitempty" tf:"serial_number"`
	// +optional
	SoftwareVersion *string `json:"softwareVersion,omitempty" tf:"software_version"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type EdgeDeviceSpec struct {
	State *EdgeDeviceSpecResource `json:"state,omitempty" tf:"-"`

	Resource EdgeDeviceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type EdgeDeviceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	DeviceProperties  []EdgeDeviceSpecDeviceProperties `json:"deviceProperties,omitempty" tf:"device_properties"`
	Location          *string                          `json:"location" tf:"location"`
	Name              *string                          `json:"name" tf:"name"`
	ResourceGroupName *string                          `json:"resourceGroupName" tf:"resource_group_name"`
	SkuName           *string                          `json:"skuName" tf:"sku_name"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type EdgeDeviceStatus struct {
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

// EdgeDeviceList is a list of EdgeDevices
type EdgeDeviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EdgeDevice CRD objects
	Items []EdgeDevice `json:"items,omitempty"`
}
