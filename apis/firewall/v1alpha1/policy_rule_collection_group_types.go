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

type PolicyRuleCollectionGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicyRuleCollectionGroupSpec   `json:"spec,omitempty"`
	Status            PolicyRuleCollectionGroupStatus `json:"status,omitempty"`
}

type PolicyRuleCollectionGroupSpecApplicationRuleCollectionRuleProtocols struct {
	Port *int64  `json:"port" tf:"port"`
	Type *string `json:"type" tf:"type"`
}

type PolicyRuleCollectionGroupSpecApplicationRuleCollectionRule struct {
	// +optional
	DestinationFqdnTags []string `json:"destinationFqdnTags,omitempty" tf:"destination_fqdn_tags"`
	// +optional
	DestinationFqdns []string                                                              `json:"destinationFqdns,omitempty" tf:"destination_fqdns"`
	Name             *string                                                               `json:"name" tf:"name"`
	Protocols        []PolicyRuleCollectionGroupSpecApplicationRuleCollectionRuleProtocols `json:"protocols" tf:"protocols"`
	// +optional
	SourceAddresses []string `json:"sourceAddresses,omitempty" tf:"source_addresses"`
	// +optional
	SourceIPGroups []string `json:"sourceIPGroups,omitempty" tf:"source_ip_groups"`
}

type PolicyRuleCollectionGroupSpecApplicationRuleCollection struct {
	Action   *string `json:"action" tf:"action"`
	Name     *string `json:"name" tf:"name"`
	Priority *int64  `json:"priority" tf:"priority"`
	// +kubebuilder:validation:MinItems=1
	Rule []PolicyRuleCollectionGroupSpecApplicationRuleCollectionRule `json:"rule" tf:"rule"`
}

type PolicyRuleCollectionGroupSpecNatRuleCollectionRule struct {
	// +optional
	DestinationAddress *string `json:"destinationAddress,omitempty" tf:"destination_address"`
	// +optional
	DestinationPorts []string `json:"destinationPorts,omitempty" tf:"destination_ports"`
	Name             *string  `json:"name" tf:"name"`
	Protocols        []string `json:"protocols" tf:"protocols"`
	// +optional
	SourceAddresses []string `json:"sourceAddresses,omitempty" tf:"source_addresses"`
	// +optional
	SourceIPGroups    []string `json:"sourceIPGroups,omitempty" tf:"source_ip_groups"`
	TranslatedAddress *string  `json:"translatedAddress" tf:"translated_address"`
	TranslatedPort    *int64   `json:"translatedPort" tf:"translated_port"`
}

type PolicyRuleCollectionGroupSpecNatRuleCollection struct {
	Action   *string `json:"action" tf:"action"`
	Name     *string `json:"name" tf:"name"`
	Priority *int64  `json:"priority" tf:"priority"`
	// +kubebuilder:validation:MinItems=1
	Rule []PolicyRuleCollectionGroupSpecNatRuleCollectionRule `json:"rule" tf:"rule"`
}

type PolicyRuleCollectionGroupSpecNetworkRuleCollectionRule struct {
	// +optional
	DestinationAddresses []string `json:"destinationAddresses,omitempty" tf:"destination_addresses"`
	// +optional
	DestinationFqdns []string `json:"destinationFqdns,omitempty" tf:"destination_fqdns"`
	// +optional
	DestinationIPGroups []string `json:"destinationIPGroups,omitempty" tf:"destination_ip_groups"`
	DestinationPorts    []string `json:"destinationPorts" tf:"destination_ports"`
	Name                *string  `json:"name" tf:"name"`
	Protocols           []string `json:"protocols" tf:"protocols"`
	// +optional
	SourceAddresses []string `json:"sourceAddresses,omitempty" tf:"source_addresses"`
	// +optional
	SourceIPGroups []string `json:"sourceIPGroups,omitempty" tf:"source_ip_groups"`
}

type PolicyRuleCollectionGroupSpecNetworkRuleCollection struct {
	Action   *string `json:"action" tf:"action"`
	Name     *string `json:"name" tf:"name"`
	Priority *int64  `json:"priority" tf:"priority"`
	// +kubebuilder:validation:MinItems=1
	Rule []PolicyRuleCollectionGroupSpecNetworkRuleCollectionRule `json:"rule" tf:"rule"`
}

type PolicyRuleCollectionGroupSpec struct {
	KubeformOutput *PolicyRuleCollectionGroupSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource PolicyRuleCollectionGroupSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type PolicyRuleCollectionGroupSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MinItems=1
	ApplicationRuleCollection []PolicyRuleCollectionGroupSpecApplicationRuleCollection `json:"applicationRuleCollection,omitempty" tf:"application_rule_collection"`
	FirewallPolicyID          *string                                                  `json:"firewallPolicyID" tf:"firewall_policy_id"`
	Name                      *string                                                  `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	NatRuleCollection []PolicyRuleCollectionGroupSpecNatRuleCollection `json:"natRuleCollection,omitempty" tf:"nat_rule_collection"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	NetworkRuleCollection []PolicyRuleCollectionGroupSpecNetworkRuleCollection `json:"networkRuleCollection,omitempty" tf:"network_rule_collection"`
	Priority              *int64                                               `json:"priority" tf:"priority"`
}

type PolicyRuleCollectionGroupStatus struct {
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

// PolicyRuleCollectionGroupList is a list of PolicyRuleCollectionGroups
type PolicyRuleCollectionGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PolicyRuleCollectionGroup CRD objects
	Items []PolicyRuleCollectionGroup `json:"items,omitempty"`
}
