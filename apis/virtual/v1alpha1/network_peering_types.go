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

type NetworkPeering struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkPeeringSpec   `json:"spec,omitempty"`
	Status            NetworkPeeringStatus `json:"status,omitempty"`
}

type NetworkPeeringSpec struct {
	State *NetworkPeeringSpecResource `json:"state,omitempty" tf:"-"`

	Resource NetworkPeeringSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type NetworkPeeringSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AllowForwardedTraffic *bool `json:"allowForwardedTraffic,omitempty" tf:"allow_forwarded_traffic"`
	// +optional
	AllowGatewayTransit *bool `json:"allowGatewayTransit,omitempty" tf:"allow_gateway_transit"`
	// +optional
	AllowVirtualNetworkAccess *bool   `json:"allowVirtualNetworkAccess,omitempty" tf:"allow_virtual_network_access"`
	Name                      *string `json:"name" tf:"name"`
	RemoteVirtualNetworkID    *string `json:"remoteVirtualNetworkID" tf:"remote_virtual_network_id"`
	ResourceGroupName         *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	UseRemoteGateways  *bool   `json:"useRemoteGateways,omitempty" tf:"use_remote_gateways"`
	VirtualNetworkName *string `json:"virtualNetworkName" tf:"virtual_network_name"`
}

type NetworkPeeringStatus struct {
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

// NetworkPeeringList is a list of NetworkPeerings
type NetworkPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkPeering CRD objects
	Items []NetworkPeering `json:"items,omitempty"`
}
