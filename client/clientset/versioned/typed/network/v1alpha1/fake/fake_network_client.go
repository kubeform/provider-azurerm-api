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
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/clientset/versioned/typed/network/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeNetworkV1alpha1 struct {
	*testing.Fake
}

func (c *FakeNetworkV1alpha1) ConnectionMonitors(namespace string) v1alpha1.ConnectionMonitorInterface {
	return &FakeConnectionMonitors{c, namespace}
}

func (c *FakeNetworkV1alpha1) DdosProtectionPlans(namespace string) v1alpha1.DdosProtectionPlanInterface {
	return &FakeDdosProtectionPlans{c, namespace}
}

func (c *FakeNetworkV1alpha1) Interfaces(namespace string) v1alpha1.InterfaceInterface {
	return &FakeInterfaces{c, namespace}
}

func (c *FakeNetworkV1alpha1) InterfaceApplicationGatewayBackendAddressPoolAssociations(namespace string) v1alpha1.InterfaceApplicationGatewayBackendAddressPoolAssociationInterface {
	return &FakeInterfaceApplicationGatewayBackendAddressPoolAssociations{c, namespace}
}

func (c *FakeNetworkV1alpha1) InterfaceApplicationSecurityGroupAssociations(namespace string) v1alpha1.InterfaceApplicationSecurityGroupAssociationInterface {
	return &FakeInterfaceApplicationSecurityGroupAssociations{c, namespace}
}

func (c *FakeNetworkV1alpha1) InterfaceBackendAddressPoolAssociations(namespace string) v1alpha1.InterfaceBackendAddressPoolAssociationInterface {
	return &FakeInterfaceBackendAddressPoolAssociations{c, namespace}
}

func (c *FakeNetworkV1alpha1) InterfaceNATRuleAssociations(namespace string) v1alpha1.InterfaceNATRuleAssociationInterface {
	return &FakeInterfaceNATRuleAssociations{c, namespace}
}

func (c *FakeNetworkV1alpha1) InterfaceSecurityGroupAssociations(namespace string) v1alpha1.InterfaceSecurityGroupAssociationInterface {
	return &FakeInterfaceSecurityGroupAssociations{c, namespace}
}

func (c *FakeNetworkV1alpha1) PacketCaptures(namespace string) v1alpha1.PacketCaptureInterface {
	return &FakePacketCaptures{c, namespace}
}

func (c *FakeNetworkV1alpha1) Profiles(namespace string) v1alpha1.ProfileInterface {
	return &FakeProfiles{c, namespace}
}

func (c *FakeNetworkV1alpha1) SecurityGroups(namespace string) v1alpha1.SecurityGroupInterface {
	return &FakeSecurityGroups{c, namespace}
}

func (c *FakeNetworkV1alpha1) SecurityRules(namespace string) v1alpha1.SecurityRuleInterface {
	return &FakeSecurityRules{c, namespace}
}

func (c *FakeNetworkV1alpha1) Watchers(namespace string) v1alpha1.WatcherInterface {
	return &FakeWatchers{c, namespace}
}

func (c *FakeNetworkV1alpha1) WatcherFlowLogs(namespace string) v1alpha1.WatcherFlowLogInterface {
	return &FakeWatcherFlowLogs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeNetworkV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
