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

type StreamingPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamingPolicySpec   `json:"spec,omitempty"`
	Status            StreamingPolicyStatus `json:"status,omitempty"`
}

type StreamingPolicySpecCommonEncryptionCbcsDefaultContentKey struct {
	// +optional
	Label *string `json:"label,omitempty" tf:"label"`
	// +optional
	PolicyName *string `json:"policyName,omitempty" tf:"policy_name"`
}

type StreamingPolicySpecCommonEncryptionCbcsDrmFairplay struct {
	// +optional
	AllowPersistentLicense *bool `json:"allowPersistentLicense,omitempty" tf:"allow_persistent_license"`
	// +optional
	CustomLicenseAcquisitionURLTemplate *string `json:"customLicenseAcquisitionURLTemplate,omitempty" tf:"custom_license_acquisition_url_template"`
}

type StreamingPolicySpecCommonEncryptionCbcsEnabledProtocols struct {
	// +optional
	Dash *bool `json:"dash,omitempty" tf:"dash"`
	// +optional
	Download *bool `json:"download,omitempty" tf:"download"`
	// +optional
	Hls *bool `json:"hls,omitempty" tf:"hls"`
	// +optional
	SmoothStreaming *bool `json:"smoothStreaming,omitempty" tf:"smooth_streaming"`
}

type StreamingPolicySpecCommonEncryptionCbcs struct {
	// +optional
	DefaultContentKey *StreamingPolicySpecCommonEncryptionCbcsDefaultContentKey `json:"defaultContentKey,omitempty" tf:"default_content_key"`
	// +optional
	DrmFairplay *StreamingPolicySpecCommonEncryptionCbcsDrmFairplay `json:"drmFairplay,omitempty" tf:"drm_fairplay"`
	// +optional
	EnabledProtocols *StreamingPolicySpecCommonEncryptionCbcsEnabledProtocols `json:"enabledProtocols,omitempty" tf:"enabled_protocols"`
}

type StreamingPolicySpecCommonEncryptionCencDefaultContentKey struct {
	// +optional
	Label *string `json:"label,omitempty" tf:"label"`
	// +optional
	PolicyName *string `json:"policyName,omitempty" tf:"policy_name"`
}

type StreamingPolicySpecCommonEncryptionCencDrmPlayready struct {
	// +optional
	CustomAttributes *string `json:"customAttributes,omitempty" tf:"custom_attributes"`
	// +optional
	CustomLicenseAcquisitionURLTemplate *string `json:"customLicenseAcquisitionURLTemplate,omitempty" tf:"custom_license_acquisition_url_template"`
}

type StreamingPolicySpecCommonEncryptionCencEnabledProtocols struct {
	// +optional
	Dash *bool `json:"dash,omitempty" tf:"dash"`
	// +optional
	Download *bool `json:"download,omitempty" tf:"download"`
	// +optional
	Hls *bool `json:"hls,omitempty" tf:"hls"`
	// +optional
	SmoothStreaming *bool `json:"smoothStreaming,omitempty" tf:"smooth_streaming"`
}

type StreamingPolicySpecCommonEncryptionCenc struct {
	// +optional
	DefaultContentKey *StreamingPolicySpecCommonEncryptionCencDefaultContentKey `json:"defaultContentKey,omitempty" tf:"default_content_key"`
	// +optional
	DrmPlayready *StreamingPolicySpecCommonEncryptionCencDrmPlayready `json:"drmPlayready,omitempty" tf:"drm_playready"`
	// +optional
	DrmWidevineCustomLicenseAcquisitionURLTemplate *string `json:"drmWidevineCustomLicenseAcquisitionURLTemplate,omitempty" tf:"drm_widevine_custom_license_acquisition_url_template"`
	// +optional
	EnabledProtocols *StreamingPolicySpecCommonEncryptionCencEnabledProtocols `json:"enabledProtocols,omitempty" tf:"enabled_protocols"`
}

type StreamingPolicySpecNoEncryptionEnabledProtocols struct {
	// +optional
	Dash *bool `json:"dash,omitempty" tf:"dash"`
	// +optional
	Download *bool `json:"download,omitempty" tf:"download"`
	// +optional
	Hls *bool `json:"hls,omitempty" tf:"hls"`
	// +optional
	SmoothStreaming *bool `json:"smoothStreaming,omitempty" tf:"smooth_streaming"`
}

type StreamingPolicySpec struct {
	State *StreamingPolicySpecResource `json:"state,omitempty" tf:"-"`

	Resource StreamingPolicySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type StreamingPolicySpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CommonEncryptionCbcs *StreamingPolicySpecCommonEncryptionCbcs `json:"commonEncryptionCbcs,omitempty" tf:"common_encryption_cbcs"`
	// +optional
	CommonEncryptionCenc *StreamingPolicySpecCommonEncryptionCenc `json:"commonEncryptionCenc,omitempty" tf:"common_encryption_cenc"`
	// +optional
	DefaultContentKeyPolicyName *string `json:"defaultContentKeyPolicyName,omitempty" tf:"default_content_key_policy_name"`
	MediaServicesAccountName    *string `json:"mediaServicesAccountName" tf:"media_services_account_name"`
	Name                        *string `json:"name" tf:"name"`
	// +optional
	NoEncryptionEnabledProtocols *StreamingPolicySpecNoEncryptionEnabledProtocols `json:"noEncryptionEnabledProtocols,omitempty" tf:"no_encryption_enabled_protocols"`
	ResourceGroupName            *string                                          `json:"resourceGroupName" tf:"resource_group_name"`
}

type StreamingPolicyStatus struct {
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

// StreamingPolicyList is a list of StreamingPolicys
type StreamingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StreamingPolicy CRD objects
	Items []StreamingPolicy `json:"items,omitempty"`
}
