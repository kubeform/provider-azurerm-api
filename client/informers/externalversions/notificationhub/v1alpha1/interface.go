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
	// AuthorizationRules returns a AuthorizationRuleInformer.
	AuthorizationRules() AuthorizationRuleInformer
	// Namespaces returns a NamespaceInformer.
	Namespaces() NamespaceInformer
	// NotificationHubs returns a NotificationHubInformer.
	NotificationHubs() NotificationHubInformer
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

// AuthorizationRules returns a AuthorizationRuleInformer.
func (v *version) AuthorizationRules() AuthorizationRuleInformer {
	return &authorizationRuleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Namespaces returns a NamespaceInformer.
func (v *version) Namespaces() NamespaceInformer {
	return &namespaceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NotificationHubs returns a NotificationHubInformer.
func (v *version) NotificationHubs() NotificationHubInformer {
	return &notificationHubInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
