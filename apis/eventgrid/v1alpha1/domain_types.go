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

type Domain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainSpec   `json:"spec,omitempty"`
	Status            DomainStatus `json:"status,omitempty"`
}

type DomainSpecInboundIPRule struct {
	// +optional
	Action *string `json:"action,omitempty" tf:"action"`
	IpMask *string `json:"ipMask" tf:"ip_mask"`
}

type DomainSpecInputMappingDefaultValues struct {
	// +optional
	DataVersion *string `json:"dataVersion,omitempty" tf:"data_version"`
	// +optional
	EventType *string `json:"eventType,omitempty" tf:"event_type"`
	// +optional
	Subject *string `json:"subject,omitempty" tf:"subject"`
}

type DomainSpecInputMappingFields struct {
	// +optional
	DataVersion *string `json:"dataVersion,omitempty" tf:"data_version"`
	// +optional
	EventTime *string `json:"eventTime,omitempty" tf:"event_time"`
	// +optional
	EventType *string `json:"eventType,omitempty" tf:"event_type"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Subject *string `json:"subject,omitempty" tf:"subject"`
	// +optional
	Topic *string `json:"topic,omitempty" tf:"topic"`
}

type DomainSpec struct {
	KubeformOutput *DomainSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource DomainSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type DomainSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint"`
	// +optional
	// +kubebuilder:validation:MaxItems=128
	InboundIPRule []DomainSpecInboundIPRule `json:"inboundIPRule,omitempty" tf:"inbound_ip_rule"`
	// +optional
	InputMappingDefaultValues *DomainSpecInputMappingDefaultValues `json:"inputMappingDefaultValues,omitempty" tf:"input_mapping_default_values"`
	// +optional
	InputMappingFields *DomainSpecInputMappingFields `json:"inputMappingFields,omitempty" tf:"input_mapping_fields"`
	// +optional
	InputSchema *string `json:"inputSchema,omitempty" tf:"input_schema"`
	Location    *string `json:"location" tf:"location"`
	Name        *string `json:"name" tf:"name"`
	// +optional
	PrimaryAccessKey *string `json:"-" sensitive:"true" tf:"primary_access_key"`
	// +optional
	PublicNetworkAccessEnabled *bool   `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled"`
	ResourceGroupName          *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryAccessKey *string `json:"-" sensitive:"true" tf:"secondary_access_key"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type DomainStatus struct {
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

// DomainList is a list of Domains
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Domain CRD objects
	Items []Domain `json:"items,omitempty"`
}
