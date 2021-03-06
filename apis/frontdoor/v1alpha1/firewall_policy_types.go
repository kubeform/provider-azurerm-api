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

type FirewallPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallPolicySpec   `json:"spec,omitempty"`
	Status            FirewallPolicyStatus `json:"status,omitempty"`
}

type FirewallPolicySpecCustomRuleMatchCondition struct {
	// +kubebuilder:validation:MaxItems=600
	MatchValues   []string `json:"matchValues" tf:"match_values"`
	MatchVariable *string  `json:"matchVariable" tf:"match_variable"`
	// +optional
	NegationCondition *bool   `json:"negationCondition,omitempty" tf:"negation_condition"`
	Operator          *string `json:"operator" tf:"operator"`
	// +optional
	Selector *string `json:"selector,omitempty" tf:"selector"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	Transforms []string `json:"transforms,omitempty" tf:"transforms"`
}

type FirewallPolicySpecCustomRule struct {
	Action *string `json:"action" tf:"action"`
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	MatchCondition []FirewallPolicySpecCustomRuleMatchCondition `json:"matchCondition,omitempty" tf:"match_condition"`
	Name           *string                                      `json:"name" tf:"name"`
	// +optional
	Priority *int64 `json:"priority,omitempty" tf:"priority"`
	// +optional
	RateLimitDurationInMinutes *int64 `json:"rateLimitDurationInMinutes,omitempty" tf:"rate_limit_duration_in_minutes"`
	// +optional
	RateLimitThreshold *int64  `json:"rateLimitThreshold,omitempty" tf:"rate_limit_threshold"`
	Type               *string `json:"type" tf:"type"`
}

type FirewallPolicySpecManagedRuleExclusion struct {
	MatchVariable *string `json:"matchVariable" tf:"match_variable"`
	Operator      *string `json:"operator" tf:"operator"`
	Selector      *string `json:"selector" tf:"selector"`
}

type FirewallPolicySpecManagedRuleOverrideExclusion struct {
	MatchVariable *string `json:"matchVariable" tf:"match_variable"`
	Operator      *string `json:"operator" tf:"operator"`
	Selector      *string `json:"selector" tf:"selector"`
}

type FirewallPolicySpecManagedRuleOverrideRuleExclusion struct {
	MatchVariable *string `json:"matchVariable" tf:"match_variable"`
	Operator      *string `json:"operator" tf:"operator"`
	Selector      *string `json:"selector" tf:"selector"`
}

type FirewallPolicySpecManagedRuleOverrideRule struct {
	Action *string `json:"action" tf:"action"`
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	Exclusion []FirewallPolicySpecManagedRuleOverrideRuleExclusion `json:"exclusion,omitempty" tf:"exclusion"`
	RuleID    *string                                              `json:"ruleID" tf:"rule_id"`
}

type FirewallPolicySpecManagedRuleOverride struct {
	// +optional
	// +kubebuilder:validation:MaxItems=100
	Exclusion []FirewallPolicySpecManagedRuleOverrideExclusion `json:"exclusion,omitempty" tf:"exclusion"`
	// +optional
	// +kubebuilder:validation:MaxItems=1000
	Rule          []FirewallPolicySpecManagedRuleOverrideRule `json:"rule,omitempty" tf:"rule"`
	RuleGroupName *string                                     `json:"ruleGroupName" tf:"rule_group_name"`
}

type FirewallPolicySpecManagedRule struct {
	// +optional
	// +kubebuilder:validation:MaxItems=100
	Exclusion []FirewallPolicySpecManagedRuleExclusion `json:"exclusion,omitempty" tf:"exclusion"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	Override []FirewallPolicySpecManagedRuleOverride `json:"override,omitempty" tf:"override"`
	Type     *string                                 `json:"type" tf:"type"`
	Version  *string                                 `json:"version" tf:"version"`
}

type FirewallPolicySpec struct {
	State *FirewallPolicySpecResource `json:"state,omitempty" tf:"-"`

	Resource FirewallPolicySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type FirewallPolicySpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CustomBlockResponseBody *string `json:"customBlockResponseBody,omitempty" tf:"custom_block_response_body"`
	// +optional
	CustomBlockResponseStatusCode *int64 `json:"customBlockResponseStatusCode,omitempty" tf:"custom_block_response_status_code"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	CustomRule []FirewallPolicySpecCustomRule `json:"customRule,omitempty" tf:"custom_rule"`
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	FrontendEndpointIDS []string `json:"frontendEndpointIDS,omitempty" tf:"frontend_endpoint_ids"`
	// +optional
	Location *string `json:"location,omitempty" tf:"location"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	ManagedRule []FirewallPolicySpecManagedRule `json:"managedRule,omitempty" tf:"managed_rule"`
	// +optional
	Mode *string `json:"mode,omitempty" tf:"mode"`
	Name *string `json:"name" tf:"name"`
	// +optional
	RedirectURL       *string `json:"redirectURL,omitempty" tf:"redirect_url"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type FirewallPolicyStatus struct {
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

// FirewallPolicyList is a list of FirewallPolicys
type FirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FirewallPolicy CRD objects
	Items []FirewallPolicy `json:"items,omitempty"`
}
