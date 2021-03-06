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

type Monitor struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorSpec   `json:"spec,omitempty"`
	Status            MonitorStatus `json:"status,omitempty"`
}

type MonitorSpecPlan struct {
	BillingCycle  *string `json:"billingCycle" tf:"billing_cycle"`
	EffectiveDate *string `json:"effectiveDate" tf:"effective_date"`
	PlanID        *string `json:"planID" tf:"plan_id"`
	UsageType     *string `json:"usageType" tf:"usage_type"`
}

type MonitorSpecUser struct {
	Email       *string `json:"email" tf:"email"`
	FirstName   *string `json:"firstName" tf:"first_name"`
	LastName    *string `json:"lastName" tf:"last_name"`
	PhoneNumber *string `json:"phoneNumber" tf:"phone_number"`
}

type MonitorSpec struct {
	State *MonitorSpecResource `json:"state,omitempty" tf:"-"`

	Resource MonitorSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type MonitorSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CompanyName *string `json:"companyName,omitempty" tf:"company_name"`
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	EnterpriseAppID *string `json:"enterpriseAppID,omitempty" tf:"enterprise_app_id"`
	Location        *string `json:"location" tf:"location"`
	// +optional
	LogzOrganizationID *string          `json:"logzOrganizationID,omitempty" tf:"logz_organization_id"`
	Name               *string          `json:"name" tf:"name"`
	Plan               *MonitorSpecPlan `json:"plan" tf:"plan"`
	ResourceGroupName  *string          `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SingleSignOnURL *string `json:"singleSignOnURL,omitempty" tf:"single_sign_on_url"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	User *MonitorSpecUser   `json:"user" tf:"user"`
}

type MonitorStatus struct {
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

// MonitorList is a list of Monitors
type MonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Monitor CRD objects
	Items []Monitor `json:"items,omitempty"`
}
