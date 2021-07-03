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
	// Caches returns a CacheInformer.
	Caches() CacheInformer
	// CacheAccessPolicies returns a CacheAccessPolicyInformer.
	CacheAccessPolicies() CacheAccessPolicyInformer
	// CacheBlobNfsTargets returns a CacheBlobNfsTargetInformer.
	CacheBlobNfsTargets() CacheBlobNfsTargetInformer
	// CacheBlobTargets returns a CacheBlobTargetInformer.
	CacheBlobTargets() CacheBlobTargetInformer
	// CacheNfsTargets returns a CacheNfsTargetInformer.
	CacheNfsTargets() CacheNfsTargetInformer
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

// Caches returns a CacheInformer.
func (v *version) Caches() CacheInformer {
	return &cacheInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CacheAccessPolicies returns a CacheAccessPolicyInformer.
func (v *version) CacheAccessPolicies() CacheAccessPolicyInformer {
	return &cacheAccessPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CacheBlobNfsTargets returns a CacheBlobNfsTargetInformer.
func (v *version) CacheBlobNfsTargets() CacheBlobNfsTargetInformer {
	return &cacheBlobNfsTargetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CacheBlobTargets returns a CacheBlobTargetInformer.
func (v *version) CacheBlobTargets() CacheBlobTargetInformer {
	return &cacheBlobTargetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CacheNfsTargets returns a CacheNfsTargetInformer.
func (v *version) CacheNfsTargets() CacheNfsTargetInformer {
	return &cacheNfsTargetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
