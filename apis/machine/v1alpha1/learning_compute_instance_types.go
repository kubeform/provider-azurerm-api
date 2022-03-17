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

type LearningComputeInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LearningComputeInstanceSpec   `json:"spec,omitempty"`
	Status            LearningComputeInstanceStatus `json:"status,omitempty"`
}

type LearningComputeInstanceSpecAssignToUser struct {
	// +optional
	ObjectID *string `json:"objectID,omitempty" tf:"object_id"`
	// +optional
	TenantID *string `json:"tenantID,omitempty" tf:"tenant_id"`
}

type LearningComputeInstanceSpecIdentity struct {
	// +optional
	IdentityIDS []string `json:"identityIDS,omitempty" tf:"identity_ids"`
	// +optional
	PrincipalID *string `json:"principalID,omitempty" tf:"principal_id"`
	// +optional
	TenantID *string `json:"tenantID,omitempty" tf:"tenant_id"`
	Type     *string `json:"type" tf:"type"`
}

type LearningComputeInstanceSpecSsh struct {
	// +optional
	Port      *int64  `json:"port,omitempty" tf:"port"`
	PublicKey *string `json:"publicKey" tf:"public_key"`
	// +optional
	Username *string `json:"username,omitempty" tf:"username"`
}

type LearningComputeInstanceSpec struct {
	State *LearningComputeInstanceSpecResource `json:"state,omitempty" tf:"-"`

	Resource LearningComputeInstanceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type LearningComputeInstanceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AssignToUser *LearningComputeInstanceSpecAssignToUser `json:"assignToUser,omitempty" tf:"assign_to_user"`
	// +optional
	AuthorizationType *string `json:"authorizationType,omitempty" tf:"authorization_type"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Identity *LearningComputeInstanceSpecIdentity `json:"identity,omitempty" tf:"identity"`
	// +optional
	LocalAuthEnabled           *bool   `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled"`
	Location                   *string `json:"location" tf:"location"`
	MachineLearningWorkspaceID *string `json:"machineLearningWorkspaceID" tf:"machine_learning_workspace_id"`
	Name                       *string `json:"name" tf:"name"`
	// +optional
	Ssh *LearningComputeInstanceSpecSsh `json:"ssh,omitempty" tf:"ssh"`
	// +optional
	SubnetResourceID *string `json:"subnetResourceID,omitempty" tf:"subnet_resource_id"`
	// +optional
	Tags               *map[string]string `json:"tags,omitempty" tf:"tags"`
	VirtualMachineSize *string            `json:"virtualMachineSize" tf:"virtual_machine_size"`
}

type LearningComputeInstanceStatus struct {
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

// LearningComputeInstanceList is a list of LearningComputeInstances
type LearningComputeInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LearningComputeInstance CRD objects
	Items []LearningComputeInstance `json:"items,omitempty"`
}