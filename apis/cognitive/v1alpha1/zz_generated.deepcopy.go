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
func (in *Account) DeepCopyInto(out *Account) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Account.
func (in *Account) DeepCopy() *Account {
	if in == nil {
		return nil
	}
	out := new(Account)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Account) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountCustomerManagedKey) DeepCopyInto(out *AccountCustomerManagedKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountCustomerManagedKey.
func (in *AccountCustomerManagedKey) DeepCopy() *AccountCustomerManagedKey {
	if in == nil {
		return nil
	}
	out := new(AccountCustomerManagedKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountCustomerManagedKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountCustomerManagedKeyList) DeepCopyInto(out *AccountCustomerManagedKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccountCustomerManagedKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountCustomerManagedKeyList.
func (in *AccountCustomerManagedKeyList) DeepCopy() *AccountCustomerManagedKeyList {
	if in == nil {
		return nil
	}
	out := new(AccountCustomerManagedKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountCustomerManagedKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountCustomerManagedKeySpec) DeepCopyInto(out *AccountCustomerManagedKeySpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(AccountCustomerManagedKeySpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountCustomerManagedKeySpec.
func (in *AccountCustomerManagedKeySpec) DeepCopy() *AccountCustomerManagedKeySpec {
	if in == nil {
		return nil
	}
	out := new(AccountCustomerManagedKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountCustomerManagedKeySpecResource) DeepCopyInto(out *AccountCustomerManagedKeySpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.CognitiveAccountID != nil {
		in, out := &in.CognitiveAccountID, &out.CognitiveAccountID
		*out = new(string)
		**out = **in
	}
	if in.IdentityClientID != nil {
		in, out := &in.IdentityClientID, &out.IdentityClientID
		*out = new(string)
		**out = **in
	}
	if in.KeyVaultKeyID != nil {
		in, out := &in.KeyVaultKeyID, &out.KeyVaultKeyID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountCustomerManagedKeySpecResource.
func (in *AccountCustomerManagedKeySpecResource) DeepCopy() *AccountCustomerManagedKeySpecResource {
	if in == nil {
		return nil
	}
	out := new(AccountCustomerManagedKeySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountCustomerManagedKeyStatus) DeepCopyInto(out *AccountCustomerManagedKeyStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountCustomerManagedKeyStatus.
func (in *AccountCustomerManagedKeyStatus) DeepCopy() *AccountCustomerManagedKeyStatus {
	if in == nil {
		return nil
	}
	out := new(AccountCustomerManagedKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountList) DeepCopyInto(out *AccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Account, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountList.
func (in *AccountList) DeepCopy() *AccountList {
	if in == nil {
		return nil
	}
	out := new(AccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountSpec) DeepCopyInto(out *AccountSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(AccountSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountSpec.
func (in *AccountSpec) DeepCopy() *AccountSpec {
	if in == nil {
		return nil
	}
	out := new(AccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountSpecIdentity) DeepCopyInto(out *AccountSpecIdentity) {
	*out = *in
	if in.IdentityIDS != nil {
		in, out := &in.IdentityIDS, &out.IdentityIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PrincipalID != nil {
		in, out := &in.PrincipalID, &out.PrincipalID
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountSpecIdentity.
func (in *AccountSpecIdentity) DeepCopy() *AccountSpecIdentity {
	if in == nil {
		return nil
	}
	out := new(AccountSpecIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountSpecNetworkAcls) DeepCopyInto(out *AccountSpecNetworkAcls) {
	*out = *in
	if in.DefaultAction != nil {
		in, out := &in.DefaultAction, &out.DefaultAction
		*out = new(string)
		**out = **in
	}
	if in.IpRules != nil {
		in, out := &in.IpRules, &out.IpRules
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VirtualNetworkRules != nil {
		in, out := &in.VirtualNetworkRules, &out.VirtualNetworkRules
		*out = make([]AccountSpecNetworkAclsVirtualNetworkRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VirtualNetworkSubnetIDS != nil {
		in, out := &in.VirtualNetworkSubnetIDS, &out.VirtualNetworkSubnetIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountSpecNetworkAcls.
func (in *AccountSpecNetworkAcls) DeepCopy() *AccountSpecNetworkAcls {
	if in == nil {
		return nil
	}
	out := new(AccountSpecNetworkAcls)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountSpecNetworkAclsVirtualNetworkRules) DeepCopyInto(out *AccountSpecNetworkAclsVirtualNetworkRules) {
	*out = *in
	if in.IgnoreMissingVnetServiceEndpoint != nil {
		in, out := &in.IgnoreMissingVnetServiceEndpoint, &out.IgnoreMissingVnetServiceEndpoint
		*out = new(bool)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountSpecNetworkAclsVirtualNetworkRules.
func (in *AccountSpecNetworkAclsVirtualNetworkRules) DeepCopy() *AccountSpecNetworkAclsVirtualNetworkRules {
	if in == nil {
		return nil
	}
	out := new(AccountSpecNetworkAclsVirtualNetworkRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountSpecResource) DeepCopyInto(out *AccountSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomSubdomainName != nil {
		in, out := &in.CustomSubdomainName, &out.CustomSubdomainName
		*out = new(string)
		**out = **in
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.Fqdns != nil {
		in, out := &in.Fqdns, &out.Fqdns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(AccountSpecIdentity)
		(*in).DeepCopyInto(*out)
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	if in.LocalAuthEnabled != nil {
		in, out := &in.LocalAuthEnabled, &out.LocalAuthEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MetricsAdvisorAadClientID != nil {
		in, out := &in.MetricsAdvisorAadClientID, &out.MetricsAdvisorAadClientID
		*out = new(string)
		**out = **in
	}
	if in.MetricsAdvisorAadTenantID != nil {
		in, out := &in.MetricsAdvisorAadTenantID, &out.MetricsAdvisorAadTenantID
		*out = new(string)
		**out = **in
	}
	if in.MetricsAdvisorSuperUserName != nil {
		in, out := &in.MetricsAdvisorSuperUserName, &out.MetricsAdvisorSuperUserName
		*out = new(string)
		**out = **in
	}
	if in.MetricsAdvisorWebsiteName != nil {
		in, out := &in.MetricsAdvisorWebsiteName, &out.MetricsAdvisorWebsiteName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkAcls != nil {
		in, out := &in.NetworkAcls, &out.NetworkAcls
		*out = new(AccountSpecNetworkAcls)
		(*in).DeepCopyInto(*out)
	}
	if in.OutboundNetworkAccessRestrited != nil {
		in, out := &in.OutboundNetworkAccessRestrited, &out.OutboundNetworkAccessRestrited
		*out = new(bool)
		**out = **in
	}
	if in.PrimaryAccessKey != nil {
		in, out := &in.PrimaryAccessKey, &out.PrimaryAccessKey
		*out = new(string)
		**out = **in
	}
	if in.PublicNetworkAccessEnabled != nil {
		in, out := &in.PublicNetworkAccessEnabled, &out.PublicNetworkAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.QnaRuntimeEndpoint != nil {
		in, out := &in.QnaRuntimeEndpoint, &out.QnaRuntimeEndpoint
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SecondaryAccessKey != nil {
		in, out := &in.SecondaryAccessKey, &out.SecondaryAccessKey
		*out = new(string)
		**out = **in
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = make([]AccountSpecStorage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountSpecResource.
func (in *AccountSpecResource) DeepCopy() *AccountSpecResource {
	if in == nil {
		return nil
	}
	out := new(AccountSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountSpecStorage) DeepCopyInto(out *AccountSpecStorage) {
	*out = *in
	if in.IdentityClientID != nil {
		in, out := &in.IdentityClientID, &out.IdentityClientID
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountID != nil {
		in, out := &in.StorageAccountID, &out.StorageAccountID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountSpecStorage.
func (in *AccountSpecStorage) DeepCopy() *AccountSpecStorage {
	if in == nil {
		return nil
	}
	out := new(AccountSpecStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountStatus) DeepCopyInto(out *AccountStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountStatus.
func (in *AccountStatus) DeepCopy() *AccountStatus {
	if in == nil {
		return nil
	}
	out := new(AccountStatus)
	in.DeepCopyInto(out)
	return out
}
