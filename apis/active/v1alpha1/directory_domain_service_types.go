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

type DirectoryDomainService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DirectoryDomainServiceSpec   `json:"spec,omitempty"`
	Status            DirectoryDomainServiceStatus `json:"status,omitempty"`
}

type DirectoryDomainServiceSpecInitialReplicaSet struct {
	// +optional
	DomainControllerIPAddresses []string `json:"domainControllerIPAddresses,omitempty" tf:"domain_controller_ip_addresses"`
	// +optional
	ExternalAccessIPAddress *string `json:"externalAccessIPAddress,omitempty" tf:"external_access_ip_address"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Location *string `json:"location,omitempty" tf:"location"`
	// +optional
	ServiceStatus *string `json:"serviceStatus,omitempty" tf:"service_status"`
	SubnetID      *string `json:"subnetID" tf:"subnet_id"`
}

type DirectoryDomainServiceSpecNotifications struct {
	// +optional
	AdditionalRecipients []string `json:"additionalRecipients,omitempty" tf:"additional_recipients"`
	// +optional
	NotifyDcAdmins *bool `json:"notifyDcAdmins,omitempty" tf:"notify_dc_admins"`
	// +optional
	NotifyGlobalAdmins *bool `json:"notifyGlobalAdmins,omitempty" tf:"notify_global_admins"`
}

type DirectoryDomainServiceSpecSecureLdap struct {
	// +optional
	CertificateExpiry *string `json:"certificateExpiry,omitempty" tf:"certificate_expiry"`
	// +optional
	CertificateThumbprint *string `json:"certificateThumbprint,omitempty" tf:"certificate_thumbprint"`
	Enabled               *bool   `json:"enabled" tf:"enabled"`
	// +optional
	ExternalAccessEnabled  *bool   `json:"externalAccessEnabled,omitempty" tf:"external_access_enabled"`
	PfxCertificate         *string `json:"-" sensitive:"true" tf:"pfx_certificate"`
	PfxCertificatePassword *string `json:"-" sensitive:"true" tf:"pfx_certificate_password"`
	// +optional
	PublicCertificate *string `json:"publicCertificate,omitempty" tf:"public_certificate"`
}

type DirectoryDomainServiceSpecSecurity struct {
	// +optional
	NtlmV1Enabled *bool `json:"ntlmV1Enabled,omitempty" tf:"ntlm_v1_enabled"`
	// +optional
	SyncKerberosPasswords *bool `json:"syncKerberosPasswords,omitempty" tf:"sync_kerberos_passwords"`
	// +optional
	SyncNtlmPasswords *bool `json:"syncNtlmPasswords,omitempty" tf:"sync_ntlm_passwords"`
	// +optional
	SyncOnPremPasswords *bool `json:"syncOnPremPasswords,omitempty" tf:"sync_on_prem_passwords"`
	// +optional
	TlsV1Enabled *bool `json:"tlsV1Enabled,omitempty" tf:"tls_v1_enabled"`
}

type DirectoryDomainServiceSpec struct {
	State *DirectoryDomainServiceSpecResource `json:"state,omitempty" tf:"-"`

	Resource DirectoryDomainServiceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type DirectoryDomainServiceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	DeploymentID *string `json:"deploymentID,omitempty" tf:"deployment_id"`
	DomainName   *string `json:"domainName" tf:"domain_name"`
	// +optional
	FilteredSyncEnabled *bool                                        `json:"filteredSyncEnabled,omitempty" tf:"filtered_sync_enabled"`
	InitialReplicaSet   *DirectoryDomainServiceSpecInitialReplicaSet `json:"initialReplicaSet" tf:"initial_replica_set"`
	Location            *string                                      `json:"location" tf:"location"`
	Name                *string                                      `json:"name" tf:"name"`
	// +optional
	Notifications     *DirectoryDomainServiceSpecNotifications `json:"notifications,omitempty" tf:"notifications"`
	ResourceGroupName *string                                  `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	ResourceID *string `json:"resourceID,omitempty" tf:"resource_id"`
	// +optional
	SecureLdap *DirectoryDomainServiceSpecSecureLdap `json:"secureLdap,omitempty" tf:"secure_ldap"`
	// +optional
	Security *DirectoryDomainServiceSpecSecurity `json:"security,omitempty" tf:"security"`
	Sku      *string                             `json:"sku" tf:"sku"`
	// +optional
	SyncOwner *string `json:"syncOwner,omitempty" tf:"sync_owner"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TenantID *string `json:"tenantID,omitempty" tf:"tenant_id"`
	// +optional
	Version *int64 `json:"version,omitempty" tf:"version"`
}

type DirectoryDomainServiceStatus struct {
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

// DirectoryDomainServiceList is a list of DirectoryDomainServices
type DirectoryDomainServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DirectoryDomainService CRD objects
	Items []DirectoryDomainService `json:"items,omitempty"`
}
