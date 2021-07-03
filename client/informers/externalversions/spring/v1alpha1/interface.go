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
	// CloudActiveDeployments returns a CloudActiveDeploymentInformer.
	CloudActiveDeployments() CloudActiveDeploymentInformer
	// CloudApps returns a CloudAppInformer.
	CloudApps() CloudAppInformer
	// CloudAppCosmosdbAssociations returns a CloudAppCosmosdbAssociationInformer.
	CloudAppCosmosdbAssociations() CloudAppCosmosdbAssociationInformer
	// CloudAppMysqlAssociations returns a CloudAppMysqlAssociationInformer.
	CloudAppMysqlAssociations() CloudAppMysqlAssociationInformer
	// CloudAppRedisAssociations returns a CloudAppRedisAssociationInformer.
	CloudAppRedisAssociations() CloudAppRedisAssociationInformer
	// CloudCertificates returns a CloudCertificateInformer.
	CloudCertificates() CloudCertificateInformer
	// CloudCustomDomains returns a CloudCustomDomainInformer.
	CloudCustomDomains() CloudCustomDomainInformer
	// CloudJavaDeployments returns a CloudJavaDeploymentInformer.
	CloudJavaDeployments() CloudJavaDeploymentInformer
	// CloudServices returns a CloudServiceInformer.
	CloudServices() CloudServiceInformer
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

// CloudActiveDeployments returns a CloudActiveDeploymentInformer.
func (v *version) CloudActiveDeployments() CloudActiveDeploymentInformer {
	return &cloudActiveDeploymentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CloudApps returns a CloudAppInformer.
func (v *version) CloudApps() CloudAppInformer {
	return &cloudAppInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CloudAppCosmosdbAssociations returns a CloudAppCosmosdbAssociationInformer.
func (v *version) CloudAppCosmosdbAssociations() CloudAppCosmosdbAssociationInformer {
	return &cloudAppCosmosdbAssociationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CloudAppMysqlAssociations returns a CloudAppMysqlAssociationInformer.
func (v *version) CloudAppMysqlAssociations() CloudAppMysqlAssociationInformer {
	return &cloudAppMysqlAssociationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CloudAppRedisAssociations returns a CloudAppRedisAssociationInformer.
func (v *version) CloudAppRedisAssociations() CloudAppRedisAssociationInformer {
	return &cloudAppRedisAssociationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CloudCertificates returns a CloudCertificateInformer.
func (v *version) CloudCertificates() CloudCertificateInformer {
	return &cloudCertificateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CloudCustomDomains returns a CloudCustomDomainInformer.
func (v *version) CloudCustomDomains() CloudCustomDomainInformer {
	return &cloudCustomDomainInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CloudJavaDeployments returns a CloudJavaDeploymentInformer.
func (v *version) CloudJavaDeployments() CloudJavaDeploymentInformer {
	return &cloudJavaDeploymentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CloudServices returns a CloudServiceInformer.
func (v *version) CloudServices() CloudServiceInformer {
	return &cloudServiceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
