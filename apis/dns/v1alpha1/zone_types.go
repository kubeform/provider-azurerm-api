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

type Zone struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ZoneSpec   `json:"spec,omitempty"`
	Status            ZoneStatus `json:"status,omitempty"`
}

type ZoneSpecSoaRecord struct {
	Email *string `json:"email" tf:"email"`
	// +optional
	ExpireTime *int64 `json:"expireTime,omitempty" tf:"expire_time"`
	// +optional
	Fqdn     *string `json:"fqdn,omitempty" tf:"fqdn"`
	HostName *string `json:"hostName" tf:"host_name"`
	// +optional
	MinimumTtl *int64 `json:"minimumTtl,omitempty" tf:"minimum_ttl"`
	// +optional
	RefreshTime *int64 `json:"refreshTime,omitempty" tf:"refresh_time"`
	// +optional
	RetryTime *int64 `json:"retryTime,omitempty" tf:"retry_time"`
	// +optional
	SerialNumber *int64 `json:"serialNumber,omitempty" tf:"serial_number"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	Ttl *int64 `json:"ttl,omitempty" tf:"ttl"`
}

type ZoneSpec struct {
	State *ZoneSpecResource `json:"state,omitempty" tf:"-"`

	Resource ZoneSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ZoneSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	MaxNumberOfRecordSets *int64  `json:"maxNumberOfRecordSets,omitempty" tf:"max_number_of_record_sets"`
	Name                  *string `json:"name" tf:"name"`
	// +optional
	NameServers []string `json:"nameServers,omitempty" tf:"name_servers"`
	// +optional
	NumberOfRecordSets *int64  `json:"numberOfRecordSets,omitempty" tf:"number_of_record_sets"`
	ResourceGroupName  *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SoaRecord *ZoneSpecSoaRecord `json:"soaRecord,omitempty" tf:"soa_record"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type ZoneStatus struct {
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

// ZoneList is a list of Zones
type ZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Zone CRD objects
	Items []Zone `json:"items,omitempty"`
}
