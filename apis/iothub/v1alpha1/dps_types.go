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

type Dps struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DpsSpec   `json:"spec,omitempty"`
	Status            DpsStatus `json:"status,omitempty"`
}

type DpsSpecIpFilterRule struct {
	Action *string `json:"action" tf:"action"`
	IpMask *string `json:"ipMask" tf:"ip_mask"`
	Name   *string `json:"name" tf:"name"`
	// +optional
	Target *string `json:"target,omitempty" tf:"target"`
}

type DpsSpecLinkedHub struct {
	// +optional
	AllocationWeight *int64 `json:"allocationWeight,omitempty" tf:"allocation_weight"`
	// +optional
	ApplyAllocationPolicy *bool   `json:"applyAllocationPolicy,omitempty" tf:"apply_allocation_policy"`
	ConnectionString      *string `json:"-" sensitive:"true" tf:"connection_string"`
	// +optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname"`
	Location *string `json:"location" tf:"location"`
}

type DpsSpecSku struct {
	Capacity *int64  `json:"capacity" tf:"capacity"`
	Name     *string `json:"name" tf:"name"`
}

type DpsSpec struct {
	State *DpsSpecResource `json:"state,omitempty" tf:"-"`

	Resource DpsSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type DpsSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AllocationPolicy *string `json:"allocationPolicy,omitempty" tf:"allocation_policy"`
	// +optional
	DeviceProvisioningHostName *string `json:"deviceProvisioningHostName,omitempty" tf:"device_provisioning_host_name"`
	// +optional
	IDScope *string `json:"IDScope,omitempty" tf:"id_scope"`
	// +optional
	IpFilterRule []DpsSpecIpFilterRule `json:"ipFilterRule,omitempty" tf:"ip_filter_rule"`
	// +optional
	LinkedHub []DpsSpecLinkedHub `json:"linkedHub,omitempty" tf:"linked_hub"`
	Location  *string            `json:"location" tf:"location"`
	Name      *string            `json:"name" tf:"name"`
	// +optional
	PublicNetworkAccessEnabled *bool   `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled"`
	ResourceGroupName          *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	ServiceOperationsHostName *string     `json:"serviceOperationsHostName,omitempty" tf:"service_operations_host_name"`
	Sku                       *DpsSpecSku `json:"sku" tf:"sku"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type DpsStatus struct {
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

// DpsList is a list of Dpss
type DpsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Dps CRD objects
	Items []Dps `json:"items,omitempty"`
}
