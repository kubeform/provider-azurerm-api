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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/clientset/versioned/typed/virtual/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeVirtualV1alpha1 struct {
	*testing.Fake
}

func (c *FakeVirtualV1alpha1) DesktopApplications(namespace string) v1alpha1.DesktopApplicationInterface {
	return &FakeDesktopApplications{c, namespace}
}

func (c *FakeVirtualV1alpha1) DesktopApplicationGroups(namespace string) v1alpha1.DesktopApplicationGroupInterface {
	return &FakeDesktopApplicationGroups{c, namespace}
}

func (c *FakeVirtualV1alpha1) DesktopHostPools(namespace string) v1alpha1.DesktopHostPoolInterface {
	return &FakeDesktopHostPools{c, namespace}
}

func (c *FakeVirtualV1alpha1) DesktopWorkspaces(namespace string) v1alpha1.DesktopWorkspaceInterface {
	return &FakeDesktopWorkspaces{c, namespace}
}

func (c *FakeVirtualV1alpha1) DesktopWorkspaceApplicationGroupAssociations(namespace string) v1alpha1.DesktopWorkspaceApplicationGroupAssociationInterface {
	return &FakeDesktopWorkspaceApplicationGroupAssociations{c, namespace}
}

func (c *FakeVirtualV1alpha1) Hubs(namespace string) v1alpha1.HubInterface {
	return &FakeHubs{c, namespace}
}

func (c *FakeVirtualV1alpha1) HubBGPConnections(namespace string) v1alpha1.HubBGPConnectionInterface {
	return &FakeHubBGPConnections{c, namespace}
}

func (c *FakeVirtualV1alpha1) HubConnections(namespace string) v1alpha1.HubConnectionInterface {
	return &FakeHubConnections{c, namespace}
}

func (c *FakeVirtualV1alpha1) HubIPs(namespace string) v1alpha1.HubIPInterface {
	return &FakeHubIPs{c, namespace}
}

func (c *FakeVirtualV1alpha1) HubRouteTables(namespace string) v1alpha1.HubRouteTableInterface {
	return &FakeHubRouteTables{c, namespace}
}

func (c *FakeVirtualV1alpha1) HubSecurityPartnerProviders(namespace string) v1alpha1.HubSecurityPartnerProviderInterface {
	return &FakeHubSecurityPartnerProviders{c, namespace}
}

func (c *FakeVirtualV1alpha1) Machines(namespace string) v1alpha1.MachineInterface {
	return &FakeMachines{c, namespace}
}

func (c *FakeVirtualV1alpha1) MachineConfigurationPolicyAssignments(namespace string) v1alpha1.MachineConfigurationPolicyAssignmentInterface {
	return &FakeMachineConfigurationPolicyAssignments{c, namespace}
}

func (c *FakeVirtualV1alpha1) MachineDataDiskAttachments(namespace string) v1alpha1.MachineDataDiskAttachmentInterface {
	return &FakeMachineDataDiskAttachments{c, namespace}
}

func (c *FakeVirtualV1alpha1) MachineExtensions(namespace string) v1alpha1.MachineExtensionInterface {
	return &FakeMachineExtensions{c, namespace}
}

func (c *FakeVirtualV1alpha1) MachineScaleSets(namespace string) v1alpha1.MachineScaleSetInterface {
	return &FakeMachineScaleSets{c, namespace}
}

func (c *FakeVirtualV1alpha1) MachineScaleSetExtensions(namespace string) v1alpha1.MachineScaleSetExtensionInterface {
	return &FakeMachineScaleSetExtensions{c, namespace}
}

func (c *FakeVirtualV1alpha1) Networks(namespace string) v1alpha1.NetworkInterface {
	return &FakeNetworks{c, namespace}
}

func (c *FakeVirtualV1alpha1) NetworkGateways(namespace string) v1alpha1.NetworkGatewayInterface {
	return &FakeNetworkGateways{c, namespace}
}

func (c *FakeVirtualV1alpha1) NetworkGatewayConnections(namespace string) v1alpha1.NetworkGatewayConnectionInterface {
	return &FakeNetworkGatewayConnections{c, namespace}
}

func (c *FakeVirtualV1alpha1) NetworkPeerings(namespace string) v1alpha1.NetworkPeeringInterface {
	return &FakeNetworkPeerings{c, namespace}
}

func (c *FakeVirtualV1alpha1) Wans(namespace string) v1alpha1.WanInterface {
	return &FakeWans{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeVirtualV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
