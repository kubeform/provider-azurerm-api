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

type DatabasePrincipal struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabasePrincipalSpec   `json:"spec,omitempty"`
	Status            DatabasePrincipalStatus `json:"status,omitempty"`
}

type DatabasePrincipalSpec struct {
	State *DatabasePrincipalSpecResource `json:"state,omitempty" tf:"-"`

	Resource DatabasePrincipalSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DatabasePrincipalSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AppID        *string `json:"appID,omitempty" tf:"app_id"`
	ClientID     *string `json:"clientID" tf:"client_id"`
	ClusterName  *string `json:"clusterName" tf:"cluster_name"`
	DatabaseName *string `json:"databaseName" tf:"database_name"`
	// +optional
	Email *string `json:"email,omitempty" tf:"email"`
	// +optional
	FullyQualifiedName *string `json:"fullyQualifiedName,omitempty" tf:"fully_qualified_name"`
	// +optional
	Name              *string `json:"name,omitempty" tf:"name"`
	ObjectID          *string `json:"objectID" tf:"object_id"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	Role              *string `json:"role" tf:"role"`
	Type              *string `json:"type" tf:"type"`
}

type DatabasePrincipalStatus struct {
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

// DatabasePrincipalList is a list of DatabasePrincipals
type DatabasePrincipalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DatabasePrincipal CRD objects
	Items []DatabasePrincipal `json:"items,omitempty"`
}
