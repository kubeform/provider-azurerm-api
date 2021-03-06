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
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/clientset/versioned/typed/devtest/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeDevtestV1alpha1 struct {
	*testing.Fake
}

func (c *FakeDevtestV1alpha1) GlobalVmShutdownSchedules(namespace string) v1alpha1.GlobalVmShutdownScheduleInterface {
	return &FakeGlobalVmShutdownSchedules{c, namespace}
}

func (c *FakeDevtestV1alpha1) Labs(namespace string) v1alpha1.LabInterface {
	return &FakeLabs{c, namespace}
}

func (c *FakeDevtestV1alpha1) LinuxVirtualMachines(namespace string) v1alpha1.LinuxVirtualMachineInterface {
	return &FakeLinuxVirtualMachines{c, namespace}
}

func (c *FakeDevtestV1alpha1) Policies(namespace string) v1alpha1.PolicyInterface {
	return &FakePolicies{c, namespace}
}

func (c *FakeDevtestV1alpha1) Schedules(namespace string) v1alpha1.ScheduleInterface {
	return &FakeSchedules{c, namespace}
}

func (c *FakeDevtestV1alpha1) VirtualNetworks(namespace string) v1alpha1.VirtualNetworkInterface {
	return &FakeVirtualNetworks{c, namespace}
}

func (c *FakeDevtestV1alpha1) WindowsVirtualMachines(namespace string) v1alpha1.WindowsVirtualMachineInterface {
	return &FakeWindowsVirtualMachines{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeDevtestV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
