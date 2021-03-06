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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "kubeform.dev/provider-azurerm-api/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// GlobalVmShutdownSchedules returns a GlobalVmShutdownScheduleInformer.
	GlobalVmShutdownSchedules() GlobalVmShutdownScheduleInformer
	// Labs returns a LabInformer.
	Labs() LabInformer
	// LinuxVirtualMachines returns a LinuxVirtualMachineInformer.
	LinuxVirtualMachines() LinuxVirtualMachineInformer
	// Policies returns a PolicyInformer.
	Policies() PolicyInformer
	// Schedules returns a ScheduleInformer.
	Schedules() ScheduleInformer
	// VirtualNetworks returns a VirtualNetworkInformer.
	VirtualNetworks() VirtualNetworkInformer
	// WindowsVirtualMachines returns a WindowsVirtualMachineInformer.
	WindowsVirtualMachines() WindowsVirtualMachineInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// GlobalVmShutdownSchedules returns a GlobalVmShutdownScheduleInformer.
func (v *version) GlobalVmShutdownSchedules() GlobalVmShutdownScheduleInformer {
	return &globalVmShutdownScheduleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Labs returns a LabInformer.
func (v *version) Labs() LabInformer {
	return &labInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// LinuxVirtualMachines returns a LinuxVirtualMachineInformer.
func (v *version) LinuxVirtualMachines() LinuxVirtualMachineInformer {
	return &linuxVirtualMachineInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Policies returns a PolicyInformer.
func (v *version) Policies() PolicyInformer {
	return &policyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Schedules returns a ScheduleInformer.
func (v *version) Schedules() ScheduleInformer {
	return &scheduleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VirtualNetworks returns a VirtualNetworkInformer.
func (v *version) VirtualNetworks() VirtualNetworkInformer {
	return &virtualNetworkInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// WindowsVirtualMachines returns a WindowsVirtualMachineInformer.
func (v *version) WindowsVirtualMachines() WindowsVirtualMachineInformer {
	return &windowsVirtualMachineInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
