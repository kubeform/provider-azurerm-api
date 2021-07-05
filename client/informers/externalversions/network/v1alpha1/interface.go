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
	"context"
	time "time"

	networkv1alpha1 "kubeform.dev/provider-azurerm-api/apis/network/v1alpha1"
	versioned "kubeform.dev/provider-azurerm-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-azurerm-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-azurerm-api/client/listers/network/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ConnectionMonitors returns a ConnectionMonitorInformer.
	ConnectionMonitors() ConnectionMonitorInformer
	// DdosProtectionPlans returns a DdosProtectionPlanInformer.
	DdosProtectionPlans() DdosProtectionPlanInformer
	// Interfaces returns a InterfaceInformer.
	Interfaces() InterfaceInformer
	// InterfaceApplicationGatewayBackendAddressPoolAssociations returns a InterfaceApplicationGatewayBackendAddressPoolAssociationInformer.
	InterfaceApplicationGatewayBackendAddressPoolAssociations() InterfaceApplicationGatewayBackendAddressPoolAssociationInformer
	// InterfaceApplicationSecurityGroupAssociations returns a InterfaceApplicationSecurityGroupAssociationInformer.
	InterfaceApplicationSecurityGroupAssociations() InterfaceApplicationSecurityGroupAssociationInformer
	// InterfaceBackendAddressPoolAssociations returns a InterfaceBackendAddressPoolAssociationInformer.
	InterfaceBackendAddressPoolAssociations() InterfaceBackendAddressPoolAssociationInformer
	// InterfaceNATRuleAssociations returns a InterfaceNATRuleAssociationInformer.
	InterfaceNATRuleAssociations() InterfaceNATRuleAssociationInformer
	// InterfaceSecurityGroupAssociations returns a InterfaceSecurityGroupAssociationInformer.
	InterfaceSecurityGroupAssociations() InterfaceSecurityGroupAssociationInformer
	// PacketCaptures returns a PacketCaptureInformer.
	PacketCaptures() PacketCaptureInformer
	// Profiles returns a ProfileInformer.
	Profiles() ProfileInformer
	// SecurityGroups returns a SecurityGroupInformer.
	SecurityGroups() SecurityGroupInformer
	// SecurityRules returns a SecurityRuleInformer.
	SecurityRules() SecurityRuleInformer
	// Watchers returns a WatcherInformer.
	Watchers() WatcherInformer
	// WatcherFlowLogs returns a WatcherFlowLogInformer.
	WatcherFlowLogs() WatcherFlowLogInformer
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

// ConnectionMonitors returns a ConnectionMonitorInformer.
func (v *version) ConnectionMonitors() ConnectionMonitorInformer {
	return &connectionMonitorInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DdosProtectionPlans returns a DdosProtectionPlanInformer.
func (v *version) DdosProtectionPlans() DdosProtectionPlanInformer {
	return &ddosProtectionPlanInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Interfaces returns a InterfaceInformer.
func (v *version) Interfaces() InterfaceInformer {
	return &interfaceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InterfaceApplicationGatewayBackendAddressPoolAssociations returns a InterfaceApplicationGatewayBackendAddressPoolAssociationInformer.
func (v *version) InterfaceApplicationGatewayBackendAddressPoolAssociations() InterfaceApplicationGatewayBackendAddressPoolAssociationInformer {
	return &interfaceApplicationGatewayBackendAddressPoolAssociationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InterfaceApplicationSecurityGroupAssociations returns a InterfaceApplicationSecurityGroupAssociationInformer.
func (v *version) InterfaceApplicationSecurityGroupAssociations() InterfaceApplicationSecurityGroupAssociationInformer {
	return &interfaceApplicationSecurityGroupAssociationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InterfaceBackendAddressPoolAssociations returns a InterfaceBackendAddressPoolAssociationInformer.
func (v *version) InterfaceBackendAddressPoolAssociations() InterfaceBackendAddressPoolAssociationInformer {
	return &interfaceBackendAddressPoolAssociationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InterfaceNATRuleAssociations returns a InterfaceNATRuleAssociationInformer.
func (v *version) InterfaceNATRuleAssociations() InterfaceNATRuleAssociationInformer {
	return &interfaceNATRuleAssociationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InterfaceSecurityGroupAssociations returns a InterfaceSecurityGroupAssociationInformer.
func (v *version) InterfaceSecurityGroupAssociations() InterfaceSecurityGroupAssociationInformer {
	return &interfaceSecurityGroupAssociationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PacketCaptures returns a PacketCaptureInformer.
func (v *version) PacketCaptures() PacketCaptureInformer {
	return &packetCaptureInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Profiles returns a ProfileInformer.
func (v *version) Profiles() ProfileInformer {
	return &profileInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SecurityGroups returns a SecurityGroupInformer.
func (v *version) SecurityGroups() SecurityGroupInformer {
	return &securityGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SecurityRules returns a SecurityRuleInformer.
func (v *version) SecurityRules() SecurityRuleInformer {
	return &securityRuleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Watchers returns a WatcherInformer.
func (v *version) Watchers() WatcherInformer {
	return &watcherInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// WatcherFlowLogs returns a WatcherFlowLogInformer.
func (v *version) WatcherFlowLogs() WatcherFlowLogInformer {
	return &watcherFlowLogInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InterfaceInformer provides access to a shared informer and lister for
// Interfaces.
type InterfaceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.InterfaceLister
}

type interfaceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewInterfaceInformer constructs a new informer for Interface type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewInterfaceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredInterfaceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredInterfaceInformer constructs a new informer for Interface type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredInterfaceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1alpha1().Interfaces(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1alpha1().Interfaces(namespace).Watch(context.TODO(), options)
			},
		},
		&networkv1alpha1.Interface{},
		resyncPeriod,
		indexers,
	)
}

func (f *interfaceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredInterfaceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *interfaceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkv1alpha1.Interface{}, f.defaultInformer)
}

func (f *interfaceInformer) Lister() v1alpha1.InterfaceLister {
	return v1alpha1.NewInterfaceLister(f.Informer().GetIndexer())
}