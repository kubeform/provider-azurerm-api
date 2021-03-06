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

type Port struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PortSpec   `json:"spec,omitempty"`
	Status            PortStatus `json:"status,omitempty"`
}

type PortSpecIdentity struct {
	IdentityIDS []string `json:"identityIDS" tf:"identity_ids"`
	Type        *string  `json:"type" tf:"type"`
}

type PortSpecLink1 struct {
	// +optional
	AdminEnabled *bool `json:"adminEnabled,omitempty" tf:"admin_enabled"`
	// +optional
	ConnectorType *string `json:"connectorType,omitempty" tf:"connector_type"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	InterfaceName *string `json:"interfaceName,omitempty" tf:"interface_name"`
	// +optional
	MacsecCakKeyvaultSecretID *string `json:"macsecCakKeyvaultSecretID,omitempty" tf:"macsec_cak_keyvault_secret_id"`
	// +optional
	MacsecCipher *string `json:"macsecCipher,omitempty" tf:"macsec_cipher"`
	// +optional
	MacsecCknKeyvaultSecretID *string `json:"macsecCknKeyvaultSecretID,omitempty" tf:"macsec_ckn_keyvault_secret_id"`
	// +optional
	PatchPanelID *string `json:"patchPanelID,omitempty" tf:"patch_panel_id"`
	// +optional
	RackID *string `json:"rackID,omitempty" tf:"rack_id"`
	// +optional
	RouterName *string `json:"routerName,omitempty" tf:"router_name"`
}

type PortSpecLink2 struct {
	// +optional
	AdminEnabled *bool `json:"adminEnabled,omitempty" tf:"admin_enabled"`
	// +optional
	ConnectorType *string `json:"connectorType,omitempty" tf:"connector_type"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	InterfaceName *string `json:"interfaceName,omitempty" tf:"interface_name"`
	// +optional
	MacsecCakKeyvaultSecretID *string `json:"macsecCakKeyvaultSecretID,omitempty" tf:"macsec_cak_keyvault_secret_id"`
	// +optional
	MacsecCipher *string `json:"macsecCipher,omitempty" tf:"macsec_cipher"`
	// +optional
	MacsecCknKeyvaultSecretID *string `json:"macsecCknKeyvaultSecretID,omitempty" tf:"macsec_ckn_keyvault_secret_id"`
	// +optional
	PatchPanelID *string `json:"patchPanelID,omitempty" tf:"patch_panel_id"`
	// +optional
	RackID *string `json:"rackID,omitempty" tf:"rack_id"`
	// +optional
	RouterName *string `json:"routerName,omitempty" tf:"router_name"`
}

type PortSpec struct {
	State *PortSpecResource `json:"state,omitempty" tf:"-"`

	Resource PortSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type PortSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	BandwidthInGbps *int64  `json:"bandwidthInGbps" tf:"bandwidth_in_gbps"`
	Encapsulation   *string `json:"encapsulation" tf:"encapsulation"`
	// +optional
	Ethertype *string `json:"ethertype,omitempty" tf:"ethertype"`
	// +optional
	Guid *string `json:"guid,omitempty" tf:"guid"`
	// +optional
	Identity *PortSpecIdentity `json:"identity,omitempty" tf:"identity"`
	// +optional
	Link1 *PortSpecLink1 `json:"link1,omitempty" tf:"link1"`
	// +optional
	Link2    *PortSpecLink2 `json:"link2,omitempty" tf:"link2"`
	Location *string        `json:"location" tf:"location"`
	// +optional
	Mtu               *string `json:"mtu,omitempty" tf:"mtu"`
	Name              *string `json:"name" tf:"name"`
	PeeringLocation   *string `json:"peeringLocation" tf:"peering_location"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type PortStatus struct {
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

// PortList is a list of Ports
type PortList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Port CRD objects
	Items []Port `json:"items,omitempty"`
}
