//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	apiv1alpha1 "kubeform.dev/apimachinery/api/v1alpha1"

	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Monitor) DeepCopyInto(out *Monitor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Monitor.
func (in *Monitor) DeepCopy() *Monitor {
	if in == nil {
		return nil
	}
	out := new(Monitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Monitor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorList) DeepCopyInto(out *MonitorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Monitor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorList.
func (in *MonitorList) DeepCopy() *MonitorList {
	if in == nil {
		return nil
	}
	out := new(MonitorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorSpec) DeepCopyInto(out *MonitorSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(MonitorSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorSpec.
func (in *MonitorSpec) DeepCopy() *MonitorSpec {
	if in == nil {
		return nil
	}
	out := new(MonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorSpecPlan) DeepCopyInto(out *MonitorSpecPlan) {
	*out = *in
	if in.BillingCycle != nil {
		in, out := &in.BillingCycle, &out.BillingCycle
		*out = new(string)
		**out = **in
	}
	if in.EffectiveDate != nil {
		in, out := &in.EffectiveDate, &out.EffectiveDate
		*out = new(string)
		**out = **in
	}
	if in.PlanID != nil {
		in, out := &in.PlanID, &out.PlanID
		*out = new(string)
		**out = **in
	}
	if in.UsageType != nil {
		in, out := &in.UsageType, &out.UsageType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorSpecPlan.
func (in *MonitorSpecPlan) DeepCopy() *MonitorSpecPlan {
	if in == nil {
		return nil
	}
	out := new(MonitorSpecPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorSpecResource) DeepCopyInto(out *MonitorSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.CompanyName != nil {
		in, out := &in.CompanyName, &out.CompanyName
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EnterpriseAppID != nil {
		in, out := &in.EnterpriseAppID, &out.EnterpriseAppID
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.LogzOrganizationID != nil {
		in, out := &in.LogzOrganizationID, &out.LogzOrganizationID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Plan != nil {
		in, out := &in.Plan, &out.Plan
		*out = new(MonitorSpecPlan)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SingleSignOnURL != nil {
		in, out := &in.SingleSignOnURL, &out.SingleSignOnURL
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(MonitorSpecUser)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorSpecResource.
func (in *MonitorSpecResource) DeepCopy() *MonitorSpecResource {
	if in == nil {
		return nil
	}
	out := new(MonitorSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorSpecUser) DeepCopyInto(out *MonitorSpecUser) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.FirstName != nil {
		in, out := &in.FirstName, &out.FirstName
		*out = new(string)
		**out = **in
	}
	if in.LastName != nil {
		in, out := &in.LastName, &out.LastName
		*out = new(string)
		**out = **in
	}
	if in.PhoneNumber != nil {
		in, out := &in.PhoneNumber, &out.PhoneNumber
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorSpecUser.
func (in *MonitorSpecUser) DeepCopy() *MonitorSpecUser {
	if in == nil {
		return nil
	}
	out := new(MonitorSpecUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorStatus) DeepCopyInto(out *MonitorStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorStatus.
func (in *MonitorStatus) DeepCopy() *MonitorStatus {
	if in == nil {
		return nil
	}
	out := new(MonitorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagRule) DeepCopyInto(out *TagRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagRule.
func (in *TagRule) DeepCopy() *TagRule {
	if in == nil {
		return nil
	}
	out := new(TagRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TagRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagRuleList) DeepCopyInto(out *TagRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TagRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagRuleList.
func (in *TagRuleList) DeepCopy() *TagRuleList {
	if in == nil {
		return nil
	}
	out := new(TagRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TagRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagRuleSpec) DeepCopyInto(out *TagRuleSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(TagRuleSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagRuleSpec.
func (in *TagRuleSpec) DeepCopy() *TagRuleSpec {
	if in == nil {
		return nil
	}
	out := new(TagRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagRuleSpecResource) DeepCopyInto(out *TagRuleSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.LogzMonitorID != nil {
		in, out := &in.LogzMonitorID, &out.LogzMonitorID
		*out = new(string)
		**out = **in
	}
	if in.SendAadLogs != nil {
		in, out := &in.SendAadLogs, &out.SendAadLogs
		*out = new(bool)
		**out = **in
	}
	if in.SendActivityLogs != nil {
		in, out := &in.SendActivityLogs, &out.SendActivityLogs
		*out = new(bool)
		**out = **in
	}
	if in.SendSubscriptionLogs != nil {
		in, out := &in.SendSubscriptionLogs, &out.SendSubscriptionLogs
		*out = new(bool)
		**out = **in
	}
	if in.TagFilter != nil {
		in, out := &in.TagFilter, &out.TagFilter
		*out = make([]TagRuleSpecTagFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagRuleSpecResource.
func (in *TagRuleSpecResource) DeepCopy() *TagRuleSpecResource {
	if in == nil {
		return nil
	}
	out := new(TagRuleSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagRuleSpecTagFilter) DeepCopyInto(out *TagRuleSpecTagFilter) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagRuleSpecTagFilter.
func (in *TagRuleSpecTagFilter) DeepCopy() *TagRuleSpecTagFilter {
	if in == nil {
		return nil
	}
	out := new(TagRuleSpecTagFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagRuleStatus) DeepCopyInto(out *TagRuleStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagRuleStatus.
func (in *TagRuleStatus) DeepCopy() *TagRuleStatus {
	if in == nil {
		return nil
	}
	out := new(TagRuleStatus)
	in.DeepCopyInto(out)
	return out
}
