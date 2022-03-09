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

type ManagedStorageAccountSasTokenDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedStorageAccountSasTokenDefinitionSpec   `json:"spec,omitempty"`
	Status            ManagedStorageAccountSasTokenDefinitionStatus `json:"status,omitempty"`
}

type ManagedStorageAccountSasTokenDefinitionSpec struct {
	State *ManagedStorageAccountSasTokenDefinitionSpecResource `json:"state,omitempty" tf:"-"`

	Resource ManagedStorageAccountSasTokenDefinitionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ManagedStorageAccountSasTokenDefinitionSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ManagedStorageAccountID *string `json:"managedStorageAccountID" tf:"managed_storage_account_id"`
	Name                    *string `json:"name" tf:"name"`
	SasTemplateURI          *string `json:"sasTemplateURI" tf:"sas_template_uri"`
	SasType                 *string `json:"sasType" tf:"sas_type"`
	// +optional
	SecretID *string `json:"secretID,omitempty" tf:"secret_id"`
	// +optional
	Tags           *map[string]string `json:"tags,omitempty" tf:"tags"`
	ValidityPeriod *string            `json:"validityPeriod" tf:"validity_period"`
}

type ManagedStorageAccountSasTokenDefinitionStatus struct {
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

// ManagedStorageAccountSasTokenDefinitionList is a list of ManagedStorageAccountSasTokenDefinitions
type ManagedStorageAccountSasTokenDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ManagedStorageAccountSasTokenDefinition CRD objects
	Items []ManagedStorageAccountSasTokenDefinition `json:"items,omitempty"`
}
