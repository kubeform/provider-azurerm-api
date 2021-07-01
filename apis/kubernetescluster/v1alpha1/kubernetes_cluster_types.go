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

type KubernetesCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubernetesClusterSpec   `json:"spec,omitempty"`
	Status            KubernetesClusterStatus `json:"status,omitempty"`
}

type KubernetesClusterSpecAddonProfileAciConnectorLinux struct {
	Enabled *bool `json:"enabled" tf:"enabled"`
	// +optional
	SubnetName *string `json:"subnetName,omitempty" tf:"subnet_name"`
}

type KubernetesClusterSpecAddonProfileAzurePolicy struct {
	Enabled *bool `json:"enabled" tf:"enabled"`
}

type KubernetesClusterSpecAddonProfileHttpApplicationRouting struct {
	Enabled *bool `json:"enabled" tf:"enabled"`
	// +optional
	HttpApplicationRoutingZoneName *string `json:"httpApplicationRoutingZoneName,omitempty" tf:"http_application_routing_zone_name"`
}

type KubernetesClusterSpecAddonProfileIngressApplicationGatewayIngressApplicationGatewayIdentity struct {
	// +optional
	ClientID *string `json:"clientID,omitempty" tf:"client_id"`
	// +optional
	ObjectID *string `json:"objectID,omitempty" tf:"object_id"`
	// +optional
	UserAssignedIdentityID *string `json:"userAssignedIdentityID,omitempty" tf:"user_assigned_identity_id"`
}

type KubernetesClusterSpecAddonProfileIngressApplicationGateway struct {
	// +optional
	EffectiveGatewayID *string `json:"effectiveGatewayID,omitempty" tf:"effective_gateway_id"`
	Enabled            *bool   `json:"enabled" tf:"enabled"`
	// +optional
	GatewayID *string `json:"gatewayID,omitempty" tf:"gateway_id"`
	// +optional
	GatewayName *string `json:"gatewayName,omitempty" tf:"gateway_name"`
	// +optional
	IngressApplicationGatewayIdentity []KubernetesClusterSpecAddonProfileIngressApplicationGatewayIngressApplicationGatewayIdentity `json:"ingressApplicationGatewayIdentity,omitempty" tf:"ingress_application_gateway_identity"`
	// +optional
	SubnetCIDR *string `json:"subnetCIDR,omitempty" tf:"subnet_cidr"`
	// +optional
	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
}

type KubernetesClusterSpecAddonProfileKubeDashboard struct {
	Enabled *bool `json:"enabled" tf:"enabled"`
}

type KubernetesClusterSpecAddonProfileOmsAgentOmsAgentIdentity struct {
	// +optional
	ClientID *string `json:"clientID,omitempty" tf:"client_id"`
	// +optional
	ObjectID *string `json:"objectID,omitempty" tf:"object_id"`
	// +optional
	UserAssignedIdentityID *string `json:"userAssignedIdentityID,omitempty" tf:"user_assigned_identity_id"`
}

type KubernetesClusterSpecAddonProfileOmsAgent struct {
	Enabled *bool `json:"enabled" tf:"enabled"`
	// +optional
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceID,omitempty" tf:"log_analytics_workspace_id"`
	// +optional
	OmsAgentIdentity []KubernetesClusterSpecAddonProfileOmsAgentOmsAgentIdentity `json:"omsAgentIdentity,omitempty" tf:"oms_agent_identity"`
}

type KubernetesClusterSpecAddonProfile struct {
	// +optional
	AciConnectorLinux *KubernetesClusterSpecAddonProfileAciConnectorLinux `json:"aciConnectorLinux,omitempty" tf:"aci_connector_linux"`
	// +optional
	AzurePolicy *KubernetesClusterSpecAddonProfileAzurePolicy `json:"azurePolicy,omitempty" tf:"azure_policy"`
	// +optional
	HttpApplicationRouting *KubernetesClusterSpecAddonProfileHttpApplicationRouting `json:"httpApplicationRouting,omitempty" tf:"http_application_routing"`
	// +optional
	IngressApplicationGateway *KubernetesClusterSpecAddonProfileIngressApplicationGateway `json:"ingressApplicationGateway,omitempty" tf:"ingress_application_gateway"`
	// +optional
	KubeDashboard *KubernetesClusterSpecAddonProfileKubeDashboard `json:"kubeDashboard,omitempty" tf:"kube_dashboard"`
	// +optional
	OmsAgent *KubernetesClusterSpecAddonProfileOmsAgent `json:"omsAgent,omitempty" tf:"oms_agent"`
}

type KubernetesClusterSpecAutoScalerProfile struct {
	// +optional
	BalanceSimilarNodeGroups *bool `json:"balanceSimilarNodeGroups,omitempty" tf:"balance_similar_node_groups"`
	// +optional
	EmptyBulkDeleteMax *string `json:"emptyBulkDeleteMax,omitempty" tf:"empty_bulk_delete_max"`
	// +optional
	Expander *string `json:"expander,omitempty" tf:"expander"`
	// +optional
	MaxGracefulTerminationSec *string `json:"maxGracefulTerminationSec,omitempty" tf:"max_graceful_termination_sec"`
	// +optional
	MaxNodeProvisioningTime *string `json:"maxNodeProvisioningTime,omitempty" tf:"max_node_provisioning_time"`
	// +optional
	MaxUnreadyNodes *int64 `json:"maxUnreadyNodes,omitempty" tf:"max_unready_nodes"`
	// +optional
	MaxUnreadyPercentage *float64 `json:"maxUnreadyPercentage,omitempty" tf:"max_unready_percentage"`
	// +optional
	NewPodScaleUpDelay *string `json:"newPodScaleUpDelay,omitempty" tf:"new_pod_scale_up_delay"`
	// +optional
	ScaleDownDelayAfterAdd *string `json:"scaleDownDelayAfterAdd,omitempty" tf:"scale_down_delay_after_add"`
	// +optional
	ScaleDownDelayAfterDelete *string `json:"scaleDownDelayAfterDelete,omitempty" tf:"scale_down_delay_after_delete"`
	// +optional
	ScaleDownDelayAfterFailure *string `json:"scaleDownDelayAfterFailure,omitempty" tf:"scale_down_delay_after_failure"`
	// +optional
	ScaleDownUnneeded *string `json:"scaleDownUnneeded,omitempty" tf:"scale_down_unneeded"`
	// +optional
	ScaleDownUnready *string `json:"scaleDownUnready,omitempty" tf:"scale_down_unready"`
	// +optional
	ScaleDownUtilizationThreshold *string `json:"scaleDownUtilizationThreshold,omitempty" tf:"scale_down_utilization_threshold"`
	// +optional
	ScanInterval *string `json:"scanInterval,omitempty" tf:"scan_interval"`
	// +optional
	SkipNodesWithLocalStorage *bool `json:"skipNodesWithLocalStorage,omitempty" tf:"skip_nodes_with_local_storage"`
	// +optional
	SkipNodesWithSystemPods *bool `json:"skipNodesWithSystemPods,omitempty" tf:"skip_nodes_with_system_pods"`
}

type KubernetesClusterSpecDefaultNodePoolKubeletConfig struct {
	// +optional
	AllowedUnsafeSysctls []string `json:"allowedUnsafeSysctls,omitempty" tf:"allowed_unsafe_sysctls"`
	// +optional
	ContainerLogMaxLine *int64 `json:"containerLogMaxLine,omitempty" tf:"container_log_max_line"`
	// +optional
	ContainerLogMaxSizeMb *int64 `json:"containerLogMaxSizeMb,omitempty" tf:"container_log_max_size_mb"`
	// +optional
	CpuCfsQuotaEnabled *bool `json:"cpuCfsQuotaEnabled,omitempty" tf:"cpu_cfs_quota_enabled"`
	// +optional
	CpuCfsQuotaPeriod *string `json:"cpuCfsQuotaPeriod,omitempty" tf:"cpu_cfs_quota_period"`
	// +optional
	CpuManagerPolicy *string `json:"cpuManagerPolicy,omitempty" tf:"cpu_manager_policy"`
	// +optional
	ImageGcHighThreshold *int64 `json:"imageGcHighThreshold,omitempty" tf:"image_gc_high_threshold"`
	// +optional
	ImageGcLowThreshold *int64 `json:"imageGcLowThreshold,omitempty" tf:"image_gc_low_threshold"`
	// +optional
	PodMaxPid *int64 `json:"podMaxPid,omitempty" tf:"pod_max_pid"`
	// +optional
	TopologyManagerPolicy *string `json:"topologyManagerPolicy,omitempty" tf:"topology_manager_policy"`
}

type KubernetesClusterSpecDefaultNodePoolLinuxOsConfigSysctlConfig struct {
	// +optional
	FsAioMaxNr *int64 `json:"fsAioMaxNr,omitempty" tf:"fs_aio_max_nr"`
	// +optional
	FsFileMax *int64 `json:"fsFileMax,omitempty" tf:"fs_file_max"`
	// +optional
	FsInotifyMaxUserWatches *int64 `json:"fsInotifyMaxUserWatches,omitempty" tf:"fs_inotify_max_user_watches"`
	// +optional
	FsNrOpen *int64 `json:"fsNrOpen,omitempty" tf:"fs_nr_open"`
	// +optional
	KernelThreadsMax *int64 `json:"kernelThreadsMax,omitempty" tf:"kernel_threads_max"`
	// +optional
	NetCoreNetdevMaxBacklog *int64 `json:"netCoreNetdevMaxBacklog,omitempty" tf:"net_core_netdev_max_backlog"`
	// +optional
	NetCoreOptmemMax *int64 `json:"netCoreOptmemMax,omitempty" tf:"net_core_optmem_max"`
	// +optional
	NetCoreRmemDefault *int64 `json:"netCoreRmemDefault,omitempty" tf:"net_core_rmem_default"`
	// +optional
	NetCoreRmemMax *int64 `json:"netCoreRmemMax,omitempty" tf:"net_core_rmem_max"`
	// +optional
	NetCoreSomaxconn *int64 `json:"netCoreSomaxconn,omitempty" tf:"net_core_somaxconn"`
	// +optional
	NetCoreWmemDefault *int64 `json:"netCoreWmemDefault,omitempty" tf:"net_core_wmem_default"`
	// +optional
	NetCoreWmemMax *int64 `json:"netCoreWmemMax,omitempty" tf:"net_core_wmem_max"`
	// +optional
	NetIpv4IPLocalPortRangeMax *int64 `json:"netIpv4IPLocalPortRangeMax,omitempty" tf:"net_ipv4_ip_local_port_range_max"`
	// +optional
	NetIpv4IPLocalPortRangeMin *int64 `json:"netIpv4IPLocalPortRangeMin,omitempty" tf:"net_ipv4_ip_local_port_range_min"`
	// +optional
	NetIpv4NeighDefaultGcThresh1 *int64 `json:"netIpv4NeighDefaultGcThresh1,omitempty" tf:"net_ipv4_neigh_default_gc_thresh1"`
	// +optional
	NetIpv4NeighDefaultGcThresh2 *int64 `json:"netIpv4NeighDefaultGcThresh2,omitempty" tf:"net_ipv4_neigh_default_gc_thresh2"`
	// +optional
	NetIpv4NeighDefaultGcThresh3 *int64 `json:"netIpv4NeighDefaultGcThresh3,omitempty" tf:"net_ipv4_neigh_default_gc_thresh3"`
	// +optional
	NetIpv4TcpFinTimeout *int64 `json:"netIpv4TcpFinTimeout,omitempty" tf:"net_ipv4_tcp_fin_timeout"`
	// +optional
	NetIpv4TcpKeepaliveIntvl *int64 `json:"netIpv4TcpKeepaliveIntvl,omitempty" tf:"net_ipv4_tcp_keepalive_intvl"`
	// +optional
	NetIpv4TcpKeepaliveProbes *int64 `json:"netIpv4TcpKeepaliveProbes,omitempty" tf:"net_ipv4_tcp_keepalive_probes"`
	// +optional
	NetIpv4TcpKeepaliveTime *int64 `json:"netIpv4TcpKeepaliveTime,omitempty" tf:"net_ipv4_tcp_keepalive_time"`
	// +optional
	NetIpv4TcpMaxSYNBacklog *int64 `json:"netIpv4TcpMaxSYNBacklog,omitempty" tf:"net_ipv4_tcp_max_syn_backlog"`
	// +optional
	NetIpv4TcpMaxTwBuckets *int64 `json:"netIpv4TcpMaxTwBuckets,omitempty" tf:"net_ipv4_tcp_max_tw_buckets"`
	// +optional
	NetIpv4TcpTwReuse *bool `json:"netIpv4TcpTwReuse,omitempty" tf:"net_ipv4_tcp_tw_reuse"`
	// +optional
	NetNetfilterNfConntrackBuckets *int64 `json:"netNetfilterNfConntrackBuckets,omitempty" tf:"net_netfilter_nf_conntrack_buckets"`
	// +optional
	NetNetfilterNfConntrackMax *int64 `json:"netNetfilterNfConntrackMax,omitempty" tf:"net_netfilter_nf_conntrack_max"`
	// +optional
	VmMaxMapCount *int64 `json:"vmMaxMapCount,omitempty" tf:"vm_max_map_count"`
	// +optional
	VmSwappiness *int64 `json:"vmSwappiness,omitempty" tf:"vm_swappiness"`
	// +optional
	VmVfsCachePressure *int64 `json:"vmVfsCachePressure,omitempty" tf:"vm_vfs_cache_pressure"`
}

type KubernetesClusterSpecDefaultNodePoolLinuxOsConfig struct {
	// +optional
	SwapFileSizeMb *int64 `json:"swapFileSizeMb,omitempty" tf:"swap_file_size_mb"`
	// +optional
	SysctlConfig *KubernetesClusterSpecDefaultNodePoolLinuxOsConfigSysctlConfig `json:"sysctlConfig,omitempty" tf:"sysctl_config"`
	// +optional
	TransparentHugePageDefrag *string `json:"transparentHugePageDefrag,omitempty" tf:"transparent_huge_page_defrag"`
	// +optional
	TransparentHugePageEnabled *string `json:"transparentHugePageEnabled,omitempty" tf:"transparent_huge_page_enabled"`
}

type KubernetesClusterSpecDefaultNodePoolUpgradeSettings struct {
	MaxSurge *string `json:"maxSurge" tf:"max_surge"`
}

type KubernetesClusterSpecDefaultNodePool struct {
	// +optional
	AvailabilityZones []string `json:"availabilityZones,omitempty" tf:"availability_zones"`
	// +optional
	EnableAutoScaling *bool `json:"enableAutoScaling,omitempty" tf:"enable_auto_scaling"`
	// +optional
	EnableHostEncryption *bool `json:"enableHostEncryption,omitempty" tf:"enable_host_encryption"`
	// +optional
	EnableNodePublicIP *bool `json:"enableNodePublicIP,omitempty" tf:"enable_node_public_ip"`
	// +optional
	KubeletConfig *KubernetesClusterSpecDefaultNodePoolKubeletConfig `json:"kubeletConfig,omitempty" tf:"kubelet_config"`
	// +optional
	LinuxOsConfig *KubernetesClusterSpecDefaultNodePoolLinuxOsConfig `json:"linuxOsConfig,omitempty" tf:"linux_os_config"`
	// +optional
	MaxCount *int64 `json:"maxCount,omitempty" tf:"max_count"`
	// +optional
	MaxPods *int64 `json:"maxPods,omitempty" tf:"max_pods"`
	// +optional
	MinCount *int64  `json:"minCount,omitempty" tf:"min_count"`
	Name     *string `json:"name" tf:"name"`
	// +optional
	NodeCount *int64 `json:"nodeCount,omitempty" tf:"node_count"`
	// +optional
	NodeLabels *map[string]string `json:"nodeLabels,omitempty" tf:"node_labels"`
	// +optional
	NodePublicIPPrefixID *string `json:"nodePublicIPPrefixID,omitempty" tf:"node_public_ip_prefix_id"`
	// +optional
	NodeTaints []string `json:"nodeTaints,omitempty" tf:"node_taints"`
	// +optional
	OnlyCriticalAddonsEnabled *bool `json:"onlyCriticalAddonsEnabled,omitempty" tf:"only_critical_addons_enabled"`
	// +optional
	OrchestratorVersion *string `json:"orchestratorVersion,omitempty" tf:"orchestrator_version"`
	// +optional
	OsDiskSizeGb *int64 `json:"osDiskSizeGb,omitempty" tf:"os_disk_size_gb"`
	// +optional
	OsDiskType *string `json:"osDiskType,omitempty" tf:"os_disk_type"`
	// +optional
	ProximityPlacementGroupID *string `json:"proximityPlacementGroupID,omitempty" tf:"proximity_placement_group_id"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// +optional
	UpgradeSettings *KubernetesClusterSpecDefaultNodePoolUpgradeSettings `json:"upgradeSettings,omitempty" tf:"upgrade_settings"`
	VmSize          *string                                              `json:"vmSize" tf:"vm_size"`
	// +optional
	VnetSubnetID *string `json:"vnetSubnetID,omitempty" tf:"vnet_subnet_id"`
}

type KubernetesClusterSpecIdentity struct {
	// +optional
	PrincipalID *string `json:"principalID,omitempty" tf:"principal_id"`
	// +optional
	TenantID *string `json:"tenantID,omitempty" tf:"tenant_id"`
	Type     *string `json:"type" tf:"type"`
	// +optional
	UserAssignedIdentityID *string `json:"userAssignedIdentityID,omitempty" tf:"user_assigned_identity_id"`
}

type KubernetesClusterSpecKubeAdminConfig struct {
	// +optional
	ClientCertificate *string `json:"-" sensitive:"true" tf:"client_certificate"`
	// +optional
	ClientKey *string `json:"-" sensitive:"true" tf:"client_key"`
	// +optional
	ClusterCaCertificate *string `json:"-" sensitive:"true" tf:"cluster_ca_certificate"`
	// +optional
	Host *string `json:"host,omitempty" tf:"host"`
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	Username *string `json:"username,omitempty" tf:"username"`
}

type KubernetesClusterSpecKubeConfig struct {
	// +optional
	ClientCertificate *string `json:"-" sensitive:"true" tf:"client_certificate"`
	// +optional
	ClientKey *string `json:"-" sensitive:"true" tf:"client_key"`
	// +optional
	ClusterCaCertificate *string `json:"-" sensitive:"true" tf:"cluster_ca_certificate"`
	// +optional
	Host *string `json:"host,omitempty" tf:"host"`
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	Username *string `json:"username,omitempty" tf:"username"`
}

type KubernetesClusterSpecKubeletIdentity struct {
	// +optional
	ClientID *string `json:"clientID,omitempty" tf:"client_id"`
	// +optional
	ObjectID *string `json:"objectID,omitempty" tf:"object_id"`
	// +optional
	UserAssignedIdentityID *string `json:"userAssignedIdentityID,omitempty" tf:"user_assigned_identity_id"`
}

type KubernetesClusterSpecLinuxProfileSshKey struct {
	KeyData *string `json:"keyData" tf:"key_data"`
}

type KubernetesClusterSpecLinuxProfile struct {
	AdminUsername *string                                  `json:"adminUsername" tf:"admin_username"`
	SshKey        *KubernetesClusterSpecLinuxProfileSshKey `json:"sshKey" tf:"ssh_key"`
}

type KubernetesClusterSpecNetworkProfileLoadBalancerProfile struct {
	// +optional
	EffectiveOutboundIPS []string `json:"effectiveOutboundIPS,omitempty" tf:"effective_outbound_ips"`
	// +optional
	IdleTimeoutInMinutes *int64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes"`
	// +optional
	ManagedOutboundIPCount *int64 `json:"managedOutboundIPCount,omitempty" tf:"managed_outbound_ip_count"`
	// +optional
	OutboundIPAddressIDS []string `json:"outboundIPAddressIDS,omitempty" tf:"outbound_ip_address_ids"`
	// +optional
	OutboundIPPrefixIDS []string `json:"outboundIPPrefixIDS,omitempty" tf:"outbound_ip_prefix_ids"`
	// +optional
	OutboundPortsAllocated *int64 `json:"outboundPortsAllocated,omitempty" tf:"outbound_ports_allocated"`
}

type KubernetesClusterSpecNetworkProfile struct {
	// +optional
	DnsServiceIP *string `json:"dnsServiceIP,omitempty" tf:"dns_service_ip"`
	// +optional
	DockerBridgeCIDR *string `json:"dockerBridgeCIDR,omitempty" tf:"docker_bridge_cidr"`
	// +optional
	LoadBalancerProfile *KubernetesClusterSpecNetworkProfileLoadBalancerProfile `json:"loadBalancerProfile,omitempty" tf:"load_balancer_profile"`
	// +optional
	LoadBalancerSku *string `json:"loadBalancerSku,omitempty" tf:"load_balancer_sku"`
	// +optional
	NetworkMode   *string `json:"networkMode,omitempty" tf:"network_mode"`
	NetworkPlugin *string `json:"networkPlugin" tf:"network_plugin"`
	// +optional
	NetworkPolicy *string `json:"networkPolicy,omitempty" tf:"network_policy"`
	// +optional
	OutboundType *string `json:"outboundType,omitempty" tf:"outbound_type"`
	// +optional
	PodCIDR *string `json:"podCIDR,omitempty" tf:"pod_cidr"`
	// +optional
	ServiceCIDR *string `json:"serviceCIDR,omitempty" tf:"service_cidr"`
}

type KubernetesClusterSpecRoleBasedAccessControlAzureActiveDirectory struct {
	// +optional
	AdminGroupObjectIDS []string `json:"adminGroupObjectIDS,omitempty" tf:"admin_group_object_ids"`
	// +optional
	AzureRbacEnabled *bool `json:"azureRbacEnabled,omitempty" tf:"azure_rbac_enabled"`
	// +optional
	ClientAppID *string `json:"clientAppID,omitempty" tf:"client_app_id"`
	// +optional
	Managed *bool `json:"managed,omitempty" tf:"managed"`
	// +optional
	ServerAppID *string `json:"serverAppID,omitempty" tf:"server_app_id"`
	// +optional
	ServerAppSecret *string `json:"-" sensitive:"true" tf:"server_app_secret"`
	// +optional
	TenantID *string `json:"tenantID,omitempty" tf:"tenant_id"`
}

type KubernetesClusterSpecRoleBasedAccessControl struct {
	// +optional
	AzureActiveDirectory *KubernetesClusterSpecRoleBasedAccessControlAzureActiveDirectory `json:"azureActiveDirectory,omitempty" tf:"azure_active_directory"`
	Enabled              *bool                                                            `json:"enabled" tf:"enabled"`
}

type KubernetesClusterSpecServicePrincipal struct {
	ClientID     *string `json:"clientID" tf:"client_id"`
	ClientSecret *string `json:"-" sensitive:"true" tf:"client_secret"`
}

type KubernetesClusterSpecWindowsProfile struct {
	// +optional
	AdminPassword *string `json:"-" sensitive:"true" tf:"admin_password"`
	AdminUsername *string `json:"adminUsername" tf:"admin_username"`
}

type KubernetesClusterSpec struct {
	KubeformOutput *KubernetesClusterSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource KubernetesClusterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type KubernetesClusterSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AddonProfile *KubernetesClusterSpecAddonProfile `json:"addonProfile,omitempty" tf:"addon_profile"`
	// +optional
	ApiServerAuthorizedIPRanges []string `json:"apiServerAuthorizedIPRanges,omitempty" tf:"api_server_authorized_ip_ranges"`
	// +optional
	AutoScalerProfile *KubernetesClusterSpecAutoScalerProfile `json:"autoScalerProfile,omitempty" tf:"auto_scaler_profile"`
	// +optional
	AutomaticChannelUpgrade *string                               `json:"automaticChannelUpgrade,omitempty" tf:"automatic_channel_upgrade"`
	DefaultNodePool         *KubernetesClusterSpecDefaultNodePool `json:"defaultNodePool" tf:"default_node_pool"`
	// +optional
	DiskEncryptionSetID *string `json:"diskEncryptionSetID,omitempty" tf:"disk_encryption_set_id"`
	// +optional
	DnsPrefix *string `json:"dnsPrefix,omitempty" tf:"dns_prefix"`
	// +optional
	DnsPrefixPrivateCluster *string `json:"dnsPrefixPrivateCluster,omitempty" tf:"dns_prefix_private_cluster"`
	// +optional
	EnablePodSecurityPolicy *bool `json:"enablePodSecurityPolicy,omitempty" tf:"enable_pod_security_policy"`
	// +optional
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn"`
	// +optional
	Identity *KubernetesClusterSpecIdentity `json:"identity,omitempty" tf:"identity"`
	// +optional
	KubeAdminConfig []KubernetesClusterSpecKubeAdminConfig `json:"kubeAdminConfig,omitempty" tf:"kube_admin_config"`
	// +optional
	KubeAdminConfigRaw *string `json:"-" sensitive:"true" tf:"kube_admin_config_raw"`
	// +optional
	KubeConfig []KubernetesClusterSpecKubeConfig `json:"kubeConfig,omitempty" tf:"kube_config"`
	// +optional
	KubeConfigRaw *string `json:"-" sensitive:"true" tf:"kube_config_raw"`
	// +optional
	KubeletIdentity *KubernetesClusterSpecKubeletIdentity `json:"kubeletIdentity,omitempty" tf:"kubelet_identity"`
	// +optional
	KubernetesVersion *string `json:"kubernetesVersion,omitempty" tf:"kubernetes_version"`
	// +optional
	LinuxProfile *KubernetesClusterSpecLinuxProfile `json:"linuxProfile,omitempty" tf:"linux_profile"`
	Location     *string                            `json:"location" tf:"location"`
	Name         *string                            `json:"name" tf:"name"`
	// +optional
	NetworkProfile *KubernetesClusterSpecNetworkProfile `json:"networkProfile,omitempty" tf:"network_profile"`
	// +optional
	NodeResourceGroup *string `json:"nodeResourceGroup,omitempty" tf:"node_resource_group"`
	// +optional
	PrivateClusterEnabled *bool `json:"privateClusterEnabled,omitempty" tf:"private_cluster_enabled"`
	// +optional
	PrivateDNSZoneID *string `json:"privateDNSZoneID,omitempty" tf:"private_dns_zone_id"`
	// +optional
	PrivateFqdn *string `json:"privateFqdn,omitempty" tf:"private_fqdn"`
	// +optional
	// Deprecated
	PrivateLinkEnabled *bool   `json:"privateLinkEnabled,omitempty" tf:"private_link_enabled"`
	ResourceGroupName  *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	RoleBasedAccessControl *KubernetesClusterSpecRoleBasedAccessControl `json:"roleBasedAccessControl,omitempty" tf:"role_based_access_control"`
	// +optional
	ServicePrincipal *KubernetesClusterSpecServicePrincipal `json:"servicePrincipal,omitempty" tf:"service_principal"`
	// +optional
	SkuTier *string `json:"skuTier,omitempty" tf:"sku_tier"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	WindowsProfile *KubernetesClusterSpecWindowsProfile `json:"windowsProfile,omitempty" tf:"windows_profile"`
}

type KubernetesClusterStatus struct {
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

// KubernetesClusterList is a list of KubernetesClusters
type KubernetesClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubernetesCluster CRD objects
	Items []KubernetesCluster `json:"items,omitempty"`
}
