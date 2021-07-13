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

type SecurityRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityRuleSpec   `json:"spec,omitempty"`
	Status            SecurityRuleStatus `json:"status,omitempty"`
}

type SecurityRuleSpec struct {
	State *SecurityRuleSpecResource `json:"state,omitempty" tf:"-"`

	Resource SecurityRuleSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type SecurityRuleSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Access *string `json:"access" tf:"access"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DestinationAddressPrefix *string `json:"destinationAddressPrefix,omitempty" tf:"destination_address_prefix"`
	// +optional
	DestinationAddressPrefixes []string `json:"destinationAddressPrefixes,omitempty" tf:"destination_address_prefixes"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	DestinationApplicationSecurityGroupIDS []string `json:"destinationApplicationSecurityGroupIDS,omitempty" tf:"destination_application_security_group_ids"`
	// +optional
	DestinationPortRange *string `json:"destinationPortRange,omitempty" tf:"destination_port_range"`
	// +optional
	DestinationPortRanges    []string `json:"destinationPortRanges,omitempty" tf:"destination_port_ranges"`
	Direction                *string  `json:"direction" tf:"direction"`
	Name                     *string  `json:"name" tf:"name"`
	NetworkSecurityGroupName *string  `json:"networkSecurityGroupName" tf:"network_security_group_name"`
	Priority                 *int64   `json:"priority" tf:"priority"`
	Protocol                 *string  `json:"protocol" tf:"protocol"`
	ResourceGroupName        *string  `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SourceAddressPrefix *string `json:"sourceAddressPrefix,omitempty" tf:"source_address_prefix"`
	// +optional
	SourceAddressPrefixes []string `json:"sourceAddressPrefixes,omitempty" tf:"source_address_prefixes"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	SourceApplicationSecurityGroupIDS []string `json:"sourceApplicationSecurityGroupIDS,omitempty" tf:"source_application_security_group_ids"`
	// +optional
	SourcePortRange *string `json:"sourcePortRange,omitempty" tf:"source_port_range"`
	// +optional
	SourcePortRanges []string `json:"sourcePortRanges,omitempty" tf:"source_port_ranges"`
}

type SecurityRuleStatus struct {
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

// SecurityRuleList is a list of SecurityRules
type SecurityRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SecurityRule CRD objects
	Items []SecurityRule `json:"items,omitempty"`
}
