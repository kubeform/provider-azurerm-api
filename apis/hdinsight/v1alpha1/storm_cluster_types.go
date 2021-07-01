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

type StormCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StormClusterSpec   `json:"spec,omitempty"`
	Status            StormClusterStatus `json:"status,omitempty"`
}

type StormClusterSpecComponentVersion struct {
	Storm *string `json:"storm" tf:"storm"`
}

type StormClusterSpecGateway struct {
	// +optional
	// Deprecated
	Enabled  *bool   `json:"enabled,omitempty" tf:"enabled"`
	Password *string `json:"-" sensitive:"true" tf:"password"`
	Username *string `json:"username" tf:"username"`
}

type StormClusterSpecMetastoresAmbari struct {
	DatabaseName *string `json:"databaseName" tf:"database_name"`
	Password     *string `json:"-" sensitive:"true" tf:"password"`
	Server       *string `json:"server" tf:"server"`
	Username     *string `json:"username" tf:"username"`
}

type StormClusterSpecMetastoresHive struct {
	DatabaseName *string `json:"databaseName" tf:"database_name"`
	Password     *string `json:"-" sensitive:"true" tf:"password"`
	Server       *string `json:"server" tf:"server"`
	Username     *string `json:"username" tf:"username"`
}

type StormClusterSpecMetastoresOozie struct {
	DatabaseName *string `json:"databaseName" tf:"database_name"`
	Password     *string `json:"-" sensitive:"true" tf:"password"`
	Server       *string `json:"server" tf:"server"`
	Username     *string `json:"username" tf:"username"`
}

type StormClusterSpecMetastores struct {
	// +optional
	Ambari *StormClusterSpecMetastoresAmbari `json:"ambari,omitempty" tf:"ambari"`
	// +optional
	Hive *StormClusterSpecMetastoresHive `json:"hive,omitempty" tf:"hive"`
	// +optional
	Oozie *StormClusterSpecMetastoresOozie `json:"oozie,omitempty" tf:"oozie"`
}

type StormClusterSpecMonitor struct {
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceID" tf:"log_analytics_workspace_id"`
	PrimaryKey              *string `json:"-" sensitive:"true" tf:"primary_key"`
}

type StormClusterSpecRolesHeadNode struct {
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys"`
	// +optional
	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
	Username *string `json:"username" tf:"username"`
	// +optional
	VirtualNetworkID *string `json:"virtualNetworkID,omitempty" tf:"virtual_network_id"`
	VmSize           *string `json:"vmSize" tf:"vm_size"`
}

type StormClusterSpecRolesWorkerNode struct {
	// +optional
	// Deprecated
	MinInstanceCount *int64 `json:"minInstanceCount,omitempty" tf:"min_instance_count"`
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys"`
	// +optional
	SubnetID            *string `json:"subnetID,omitempty" tf:"subnet_id"`
	TargetInstanceCount *int64  `json:"targetInstanceCount" tf:"target_instance_count"`
	Username            *string `json:"username" tf:"username"`
	// +optional
	VirtualNetworkID *string `json:"virtualNetworkID,omitempty" tf:"virtual_network_id"`
	VmSize           *string `json:"vmSize" tf:"vm_size"`
}

type StormClusterSpecRolesZookeeperNode struct {
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys"`
	// +optional
	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
	Username *string `json:"username" tf:"username"`
	// +optional
	VirtualNetworkID *string `json:"virtualNetworkID,omitempty" tf:"virtual_network_id"`
	VmSize           *string `json:"vmSize" tf:"vm_size"`
}

type StormClusterSpecRoles struct {
	HeadNode      *StormClusterSpecRolesHeadNode      `json:"headNode" tf:"head_node"`
	WorkerNode    *StormClusterSpecRolesWorkerNode    `json:"workerNode" tf:"worker_node"`
	ZookeeperNode *StormClusterSpecRolesZookeeperNode `json:"zookeeperNode" tf:"zookeeper_node"`
}

type StormClusterSpecStorageAccount struct {
	IsDefault          *bool   `json:"isDefault" tf:"is_default"`
	StorageAccountKey  *string `json:"-" sensitive:"true" tf:"storage_account_key"`
	StorageContainerID *string `json:"storageContainerID" tf:"storage_container_id"`
}

type StormClusterSpec struct {
	KubeformOutput *StormClusterSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource StormClusterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type StormClusterSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ClusterVersion   *string                           `json:"clusterVersion" tf:"cluster_version"`
	ComponentVersion *StormClusterSpecComponentVersion `json:"componentVersion" tf:"component_version"`
	Gateway          *StormClusterSpecGateway          `json:"gateway" tf:"gateway"`
	// +optional
	HttpsEndpoint *string `json:"httpsEndpoint,omitempty" tf:"https_endpoint"`
	Location      *string `json:"location" tf:"location"`
	// +optional
	Metastores *StormClusterSpecMetastores `json:"metastores,omitempty" tf:"metastores"`
	// +optional
	Monitor           *StormClusterSpecMonitor `json:"monitor,omitempty" tf:"monitor"`
	Name              *string                  `json:"name" tf:"name"`
	ResourceGroupName *string                  `json:"resourceGroupName" tf:"resource_group_name"`
	Roles             *StormClusterSpecRoles   `json:"roles" tf:"roles"`
	// +optional
	SshEndpoint *string `json:"sshEndpoint,omitempty" tf:"ssh_endpoint"`
	// +optional
	StorageAccount []StormClusterSpecStorageAccount `json:"storageAccount,omitempty" tf:"storage_account"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	Tier *string            `json:"tier" tf:"tier"`
	// +optional
	TlsMinVersion *string `json:"tlsMinVersion,omitempty" tf:"tls_min_version"`
}

type StormClusterStatus struct {
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

// StormClusterList is a list of StormClusters
type StormClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StormCluster CRD objects
	Items []StormCluster `json:"items,omitempty"`
}
