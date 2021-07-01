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

type NotificationHub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NotificationHubSpec   `json:"spec,omitempty"`
	Status            NotificationHubStatus `json:"status,omitempty"`
}

type NotificationHubSpecApnsCredential struct {
	ApplicationMode *string `json:"applicationMode" tf:"application_mode"`
	BundleID        *string `json:"bundleID" tf:"bundle_id"`
	KeyID           *string `json:"keyID" tf:"key_id"`
	TeamID          *string `json:"teamID" tf:"team_id"`
	Token           *string `json:"-" sensitive:"true" tf:"token"`
}

type NotificationHubSpecGcmCredential struct {
	ApiKey *string `json:"-" sensitive:"true" tf:"api_key"`
}

type NotificationHubSpec struct {
	KubeformOutput *NotificationHubSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource NotificationHubSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type NotificationHubSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ApnsCredential *NotificationHubSpecApnsCredential `json:"apnsCredential,omitempty" tf:"apns_credential"`
	// +optional
	GcmCredential     *NotificationHubSpecGcmCredential `json:"gcmCredential,omitempty" tf:"gcm_credential"`
	Location          *string                           `json:"location" tf:"location"`
	Name              *string                           `json:"name" tf:"name"`
	NamespaceName     *string                           `json:"namespaceName" tf:"namespace_name"`
	ResourceGroupName *string                           `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type NotificationHubStatus struct {
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

// NotificationHubList is a list of NotificationHubs
type NotificationHubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NotificationHub CRD objects
	Items []NotificationHub `json:"items,omitempty"`
}
