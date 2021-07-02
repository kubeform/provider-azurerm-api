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

type Rule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RuleSpec   `json:"spec,omitempty"`
	Status            RuleStatus `json:"status,omitempty"`
}

type RuleSpec struct {
	KubeformOutput *RuleSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource RuleSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type RuleSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	BackendAddressPoolID *string `json:"backendAddressPoolID,omitempty" tf:"backend_address_pool_id"`
	BackendPort          *int64  `json:"backendPort" tf:"backend_port"`
	// +optional
	DisableOutboundSnat *bool `json:"disableOutboundSnat,omitempty" tf:"disable_outbound_snat"`
	// +optional
	EnableFloatingIP *bool `json:"enableFloatingIP,omitempty" tf:"enable_floating_ip"`
	// +optional
	EnableTcpReset *bool `json:"enableTcpReset,omitempty" tf:"enable_tcp_reset"`
	// +optional
	FrontendIPConfigurationID   *string `json:"frontendIPConfigurationID,omitempty" tf:"frontend_ip_configuration_id"`
	FrontendIPConfigurationName *string `json:"frontendIPConfigurationName" tf:"frontend_ip_configuration_name"`
	FrontendPort                *int64  `json:"frontendPort" tf:"frontend_port"`
	// +optional
	IdleTimeoutInMinutes *int64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes"`
	// +optional
	LoadDistribution *string `json:"loadDistribution,omitempty" tf:"load_distribution"`
	LoadbalancerID   *string `json:"loadbalancerID" tf:"loadbalancer_id"`
	Name             *string `json:"name" tf:"name"`
	// +optional
	ProbeID           *string `json:"probeID,omitempty" tf:"probe_id"`
	Protocol          *string `json:"protocol" tf:"protocol"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
}

type RuleStatus struct {
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

// RuleList is a list of Rules
type RuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Rule CRD objects
	Items []Rule `json:"items,omitempty"`
}