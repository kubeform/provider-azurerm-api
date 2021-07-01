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

type ShareDatasetBlobStorage struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ShareDatasetBlobStorageSpec   `json:"spec,omitempty"`
	Status            ShareDatasetBlobStorageStatus `json:"status,omitempty"`
}

type ShareDatasetBlobStorageSpecStorageAccount struct {
	Name              *string `json:"name" tf:"name"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	SubscriptionID    *string `json:"subscriptionID" tf:"subscription_id"`
}

type ShareDatasetBlobStorageSpec struct {
	KubeformOutput *ShareDatasetBlobStorageSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource ShareDatasetBlobStorageSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ShareDatasetBlobStorageSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ContainerName *string `json:"containerName" tf:"container_name"`
	DataShareID   *string `json:"dataShareID" tf:"data_share_id"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	FilePath *string `json:"filePath,omitempty" tf:"file_path"`
	// +optional
	FolderPath     *string                                    `json:"folderPath,omitempty" tf:"folder_path"`
	Name           *string                                    `json:"name" tf:"name"`
	StorageAccount *ShareDatasetBlobStorageSpecStorageAccount `json:"storageAccount" tf:"storage_account"`
}

type ShareDatasetBlobStorageStatus struct {
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

// ShareDatasetBlobStorageList is a list of ShareDatasetBlobStorages
type ShareDatasetBlobStorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ShareDatasetBlobStorage CRD objects
	Items []ShareDatasetBlobStorage `json:"items,omitempty"`
}
