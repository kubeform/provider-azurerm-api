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

type SubscriptionRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubscriptionRuleSpec   `json:"spec,omitempty"`
	Status            SubscriptionRuleStatus `json:"status,omitempty"`
}

type SubscriptionRuleSpecCorrelationFilter struct {
	// +optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type"`
	// +optional
	CorrelationID *string `json:"correlationID,omitempty" tf:"correlation_id"`
	// +optional
	Label *string `json:"label,omitempty" tf:"label"`
	// +optional
	MessageID *string `json:"messageID,omitempty" tf:"message_id"`
	// +optional
	Properties *map[string]string `json:"properties,omitempty" tf:"properties"`
	// +optional
	ReplyTo *string `json:"replyTo,omitempty" tf:"reply_to"`
	// +optional
	ReplyToSessionID *string `json:"replyToSessionID,omitempty" tf:"reply_to_session_id"`
	// +optional
	SessionID *string `json:"sessionID,omitempty" tf:"session_id"`
	// +optional
	To *string `json:"to,omitempty" tf:"to"`
}

type SubscriptionRuleSpec struct {
	State *SubscriptionRuleSpecResource `json:"state,omitempty" tf:"-"`

	Resource SubscriptionRuleSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SubscriptionRuleSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Action *string `json:"action,omitempty" tf:"action"`
	// +optional
	CorrelationFilter *SubscriptionRuleSpecCorrelationFilter `json:"correlationFilter,omitempty" tf:"correlation_filter"`
	FilterType        *string                                `json:"filterType" tf:"filter_type"`
	Name              *string                                `json:"name" tf:"name"`
	// +optional
	// Deprecated
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name"`
	// +optional
	// Deprecated
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name"`
	// +optional
	SqlFilter *string `json:"sqlFilter,omitempty" tf:"sql_filter"`
	// +optional
	SubscriptionID *string `json:"subscriptionID,omitempty" tf:"subscription_id"`
	// +optional
	// Deprecated
	SubscriptionName *string `json:"subscriptionName,omitempty" tf:"subscription_name"`
	// +optional
	// Deprecated
	TopicName *string `json:"topicName,omitempty" tf:"topic_name"`
}

type SubscriptionRuleStatus struct {
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

// SubscriptionRuleList is a list of SubscriptionRules
type SubscriptionRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SubscriptionRule CRD objects
	Items []SubscriptionRule `json:"items,omitempty"`
}
