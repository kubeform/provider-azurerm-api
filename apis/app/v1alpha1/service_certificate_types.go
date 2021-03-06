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

type ServiceCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceCertificateSpec   `json:"spec,omitempty"`
	Status            ServiceCertificateStatus `json:"status,omitempty"`
}

type ServiceCertificateSpec struct {
	State *ServiceCertificateSpecResource `json:"state,omitempty" tf:"-"`

	Resource ServiceCertificateSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ServiceCertificateSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AppServicePlanID *string `json:"appServicePlanID,omitempty" tf:"app_service_plan_id"`
	// +optional
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date"`
	// +optional
	FriendlyName *string `json:"friendlyName,omitempty" tf:"friendly_name"`
	// +optional
	HostNames []string `json:"hostNames,omitempty" tf:"host_names"`
	// +optional
	// Deprecated
	HostingEnvironmentProfileID *string `json:"hostingEnvironmentProfileID,omitempty" tf:"hosting_environment_profile_id"`
	// +optional
	IssueDate *string `json:"issueDate,omitempty" tf:"issue_date"`
	// +optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer"`
	// +optional
	KeyVaultSecretID *string `json:"keyVaultSecretID,omitempty" tf:"key_vault_secret_id"`
	Location         *string `json:"location" tf:"location"`
	Name             *string `json:"name" tf:"name"`
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	PfxBlob           *string `json:"-" sensitive:"true" tf:"pfx_blob"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SubjectName *string `json:"subjectName,omitempty" tf:"subject_name"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint"`
}

type ServiceCertificateStatus struct {
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

// ServiceCertificateList is a list of ServiceCertificates
type ServiceCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServiceCertificate CRD objects
	Items []ServiceCertificate `json:"items,omitempty"`
}
