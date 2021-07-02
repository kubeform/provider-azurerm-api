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

type Workspace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkspaceSpec   `json:"spec,omitempty"`
	Status            WorkspaceStatus `json:"status,omitempty"`
}

type WorkspaceSpecAadAdmin struct {
	Login    *string `json:"login" tf:"login"`
	ObjectID *string `json:"objectID" tf:"object_id"`
	TenantID *string `json:"tenantID" tf:"tenant_id"`
}

type WorkspaceSpecAzureDevopsRepo struct {
	AccountName    *string `json:"accountName" tf:"account_name"`
	BranchName     *string `json:"branchName" tf:"branch_name"`
	ProjectName    *string `json:"projectName" tf:"project_name"`
	RepositoryName *string `json:"repositoryName" tf:"repository_name"`
	RootFolder     *string `json:"rootFolder" tf:"root_folder"`
}

type WorkspaceSpecGithubRepo struct {
	AccountName *string `json:"accountName" tf:"account_name"`
	BranchName  *string `json:"branchName" tf:"branch_name"`
	// +optional
	GitURL         *string `json:"gitURL,omitempty" tf:"git_url"`
	RepositoryName *string `json:"repositoryName" tf:"repository_name"`
	RootFolder     *string `json:"rootFolder" tf:"root_folder"`
}

type WorkspaceSpecIdentity struct {
	// +optional
	PrincipalID *string `json:"principalID,omitempty" tf:"principal_id"`
	// +optional
	TenantID *string `json:"tenantID,omitempty" tf:"tenant_id"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type WorkspaceSpec struct {
	KubeformOutput *WorkspaceSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource WorkspaceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type WorkspaceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AadAdmin *WorkspaceSpecAadAdmin `json:"aadAdmin,omitempty" tf:"aad_admin"`
	// +optional
	AzureDevopsRepo *WorkspaceSpecAzureDevopsRepo `json:"azureDevopsRepo,omitempty" tf:"azure_devops_repo"`
	// +optional
	ConnectivityEndpoints *map[string]string `json:"connectivityEndpoints,omitempty" tf:"connectivity_endpoints"`
	// +optional
	CustomerManagedKeyVersionlessID *string `json:"customerManagedKeyVersionlessID,omitempty" tf:"customer_managed_key_versionless_id"`
	// +optional
	DataExfiltrationProtectionEnabled *bool `json:"dataExfiltrationProtectionEnabled,omitempty" tf:"data_exfiltration_protection_enabled"`
	// +optional
	GithubRepo *WorkspaceSpecGithubRepo `json:"githubRepo,omitempty" tf:"github_repo"`
	// +optional
	Identity []WorkspaceSpecIdentity `json:"identity,omitempty" tf:"identity"`
	Location *string                 `json:"location" tf:"location"`
	// +optional
	ManagedResourceGroupName *string `json:"managedResourceGroupName,omitempty" tf:"managed_resource_group_name"`
	// +optional
	ManagedVirtualNetworkEnabled  *bool   `json:"managedVirtualNetworkEnabled,omitempty" tf:"managed_virtual_network_enabled"`
	Name                          *string `json:"name" tf:"name"`
	ResourceGroupName             *string `json:"resourceGroupName" tf:"resource_group_name"`
	SqlAdministratorLogin         *string `json:"sqlAdministratorLogin" tf:"sql_administrator_login"`
	SqlAdministratorLoginPassword *string `json:"-" sensitive:"true" tf:"sql_administrator_login_password"`
	// +optional
	SqlIdentityControlEnabled       *bool   `json:"sqlIdentityControlEnabled,omitempty" tf:"sql_identity_control_enabled"`
	StorageDataLakeGen2FilesystemID *string `json:"storageDataLakeGen2FilesystemID" tf:"storage_data_lake_gen2_filesystem_id"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type WorkspaceStatus struct {
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

// WorkspaceList is a list of Workspaces
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Workspace CRD objects
	Items []Workspace `json:"items,omitempty"`
}