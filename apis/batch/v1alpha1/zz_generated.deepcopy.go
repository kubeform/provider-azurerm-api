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
func (in *AccountSpecKeyVaultReference) DeepCopyInto(out *AccountSpecKeyVaultReference) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Url != nil {
		in, out := &in.Url, &out.Url
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountSpecKeyVaultReference.
func (in *AccountSpecKeyVaultReference) DeepCopy() *AccountSpecKeyVaultReference {
	if in == nil {
		return nil
	}
	out := new(AccountSpecKeyVaultReference)
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
	if in.AccountEndpoint != nil {
		in, out := &in.AccountEndpoint, &out.AccountEndpoint
		*out = new(string)
		**out = **in
	}
	if in.KeyVaultReference != nil {
		in, out := &in.KeyVaultReference, &out.KeyVaultReference
		*out = new(AccountSpecKeyVaultReference)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PoolAllocationMode != nil {
		in, out := &in.PoolAllocationMode, &out.PoolAllocationMode
		*out = new(string)
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
	if in.StorageAccountID != nil {
		in, out := &in.StorageAccountID, &out.StorageAccountID
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Application) DeepCopyInto(out *Application) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Application.
func (in *Application) DeepCopy() *Application {
	if in == nil {
		return nil
	}
	out := new(Application)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Application) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationList) DeepCopyInto(out *ApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Application, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationList.
func (in *ApplicationList) DeepCopy() *ApplicationList {
	if in == nil {
		return nil
	}
	out := new(ApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpec) DeepCopyInto(out *ApplicationSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ApplicationSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSpec.
func (in *ApplicationSpec) DeepCopy() *ApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpecResource) DeepCopyInto(out *ApplicationSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AccountName != nil {
		in, out := &in.AccountName, &out.AccountName
		*out = new(string)
		**out = **in
	}
	if in.AllowUpdates != nil {
		in, out := &in.AllowUpdates, &out.AllowUpdates
		*out = new(bool)
		**out = **in
	}
	if in.DefaultVersion != nil {
		in, out := &in.DefaultVersion, &out.DefaultVersion
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSpecResource.
func (in *ApplicationSpecResource) DeepCopy() *ApplicationSpecResource {
	if in == nil {
		return nil
	}
	out := new(ApplicationSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationStatus) DeepCopyInto(out *ApplicationStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationStatus.
func (in *ApplicationStatus) DeepCopy() *ApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Certificate) DeepCopyInto(out *Certificate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Certificate.
func (in *Certificate) DeepCopy() *Certificate {
	if in == nil {
		return nil
	}
	out := new(Certificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Certificate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateList) DeepCopyInto(out *CertificateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Certificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateList.
func (in *CertificateList) DeepCopy() *CertificateList {
	if in == nil {
		return nil
	}
	out := new(CertificateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSpec) DeepCopyInto(out *CertificateSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(CertificateSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSpec.
func (in *CertificateSpec) DeepCopy() *CertificateSpec {
	if in == nil {
		return nil
	}
	out := new(CertificateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSpecResource) DeepCopyInto(out *CertificateSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AccountName != nil {
		in, out := &in.AccountName, &out.AccountName
		*out = new(string)
		**out = **in
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.PublicData != nil {
		in, out := &in.PublicData, &out.PublicData
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.Thumbprint != nil {
		in, out := &in.Thumbprint, &out.Thumbprint
		*out = new(string)
		**out = **in
	}
	if in.ThumbprintAlgorithm != nil {
		in, out := &in.ThumbprintAlgorithm, &out.ThumbprintAlgorithm
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSpecResource.
func (in *CertificateSpecResource) DeepCopy() *CertificateSpecResource {
	if in == nil {
		return nil
	}
	out := new(CertificateSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateStatus) DeepCopyInto(out *CertificateStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateStatus.
func (in *CertificateStatus) DeepCopy() *CertificateStatus {
	if in == nil {
		return nil
	}
	out := new(CertificateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pool) DeepCopyInto(out *Pool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pool.
func (in *Pool) DeepCopy() *Pool {
	if in == nil {
		return nil
	}
	out := new(Pool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolList) DeepCopyInto(out *PoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolList.
func (in *PoolList) DeepCopy() *PoolList {
	if in == nil {
		return nil
	}
	out := new(PoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolSpec) DeepCopyInto(out *PoolSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(PoolSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolSpec.
func (in *PoolSpec) DeepCopy() *PoolSpec {
	if in == nil {
		return nil
	}
	out := new(PoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolSpecAutoScale) DeepCopyInto(out *PoolSpecAutoScale) {
	*out = *in
	if in.EvaluationInterval != nil {
		in, out := &in.EvaluationInterval, &out.EvaluationInterval
		*out = new(string)
		**out = **in
	}
	if in.Formula != nil {
		in, out := &in.Formula, &out.Formula
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolSpecAutoScale.
func (in *PoolSpecAutoScale) DeepCopy() *PoolSpecAutoScale {
	if in == nil {
		return nil
	}
	out := new(PoolSpecAutoScale)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolSpecCertificate) DeepCopyInto(out *PoolSpecCertificate) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.StoreLocation != nil {
		in, out := &in.StoreLocation, &out.StoreLocation
		*out = new(string)
		**out = **in
	}
	if in.StoreName != nil {
		in, out := &in.StoreName, &out.StoreName
		*out = new(string)
		**out = **in
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolSpecCertificate.
func (in *PoolSpecCertificate) DeepCopy() *PoolSpecCertificate {
	if in == nil {
		return nil
	}
	out := new(PoolSpecCertificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolSpecContainerConfiguration) DeepCopyInto(out *PoolSpecContainerConfiguration) {
	*out = *in
	if in.ContainerImageNames != nil {
		in, out := &in.ContainerImageNames, &out.ContainerImageNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ContainerRegistries != nil {
		in, out := &in.ContainerRegistries, &out.ContainerRegistries
		*out = make([]PoolSpecContainerConfigurationContainerRegistries, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolSpecContainerConfiguration.
func (in *PoolSpecContainerConfiguration) DeepCopy() *PoolSpecContainerConfiguration {
	if in == nil {
		return nil
	}
	out := new(PoolSpecContainerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolSpecContainerConfigurationContainerRegistries) DeepCopyInto(out *PoolSpecContainerConfigurationContainerRegistries) {
	*out = *in
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.RegistryServer != nil {
		in, out := &in.RegistryServer, &out.RegistryServer
		*out = new(string)
		**out = **in
	}
	if in.UserName != nil {
		in, out := &in.UserName, &out.UserName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolSpecContainerConfigurationContainerRegistries.
func (in *PoolSpecContainerConfigurationContainerRegistries) DeepCopy() *PoolSpecContainerConfigurationContainerRegistries {
	if in == nil {
		return nil
	}
	out := new(PoolSpecContainerConfigurationContainerRegistries)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolSpecFixedScale) DeepCopyInto(out *PoolSpecFixedScale) {
	*out = *in
	if in.ResizeTimeout != nil {
		in, out := &in.ResizeTimeout, &out.ResizeTimeout
		*out = new(string)
		**out = **in
	}
	if in.TargetDedicatedNodes != nil {
		in, out := &in.TargetDedicatedNodes, &out.TargetDedicatedNodes
		*out = new(int64)
		**out = **in
	}
	if in.TargetLowPriorityNodes != nil {
		in, out := &in.TargetLowPriorityNodes, &out.TargetLowPriorityNodes
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolSpecFixedScale.
func (in *PoolSpecFixedScale) DeepCopy() *PoolSpecFixedScale {
	if in == nil {
		return nil
	}
	out := new(PoolSpecFixedScale)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolSpecNetworkConfiguration) DeepCopyInto(out *PoolSpecNetworkConfiguration) {
	*out = *in
	if in.EndpointConfiguration != nil {
		in, out := &in.EndpointConfiguration, &out.EndpointConfiguration
		*out = make([]PoolSpecNetworkConfigurationEndpointConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PublicAddressProvisioningType != nil {
		in, out := &in.PublicAddressProvisioningType, &out.PublicAddressProvisioningType
		*out = new(string)
		**out = **in
	}
	if in.PublicIPS != nil {
		in, out := &in.PublicIPS, &out.PublicIPS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolSpecNetworkConfiguration.
func (in *PoolSpecNetworkConfiguration) DeepCopy() *PoolSpecNetworkConfiguration {
	if in == nil {
		return nil
	}
	out := new(PoolSpecNetworkConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolSpecNetworkConfigurationEndpointConfiguration) DeepCopyInto(out *PoolSpecNetworkConfigurationEndpointConfiguration) {
	*out = *in
	if in.BackendPort != nil {
		in, out := &in.BackendPort, &out.BackendPort
		*out = new(int64)
		**out = **in
	}
	if in.FrontendPortRange != nil {
		in, out := &in.FrontendPortRange, &out.FrontendPortRange
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkSecurityGroupRules != nil {
		in, out := &in.NetworkSecurityGroupRules, &out.NetworkSecurityGroupRules
		*out = make([]PoolSpecNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolSpecNetworkConfigurationEndpointConfiguration.
func (in *PoolSpecNetworkConfigurationEndpointConfiguration) DeepCopy() *PoolSpecNetworkConfigurationEndpointConfiguration {
	if in == nil {
		return nil
	}
	out := new(PoolSpecNetworkConfigurationEndpointConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolSpecNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRules) DeepCopyInto(out *PoolSpecNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRules) {
	*out = *in
	if in.Access != nil {
		in, out := &in.Access, &out.Access
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int64)
		**out = **in
	}
	if in.SourceAddressPrefix != nil {
		in, out := &in.SourceAddressPrefix, &out.SourceAddressPrefix
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolSpecNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRules.
func (in *PoolSpecNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRules) DeepCopy() *PoolSpecNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRules {
	if in == nil {
		return nil
	}
	out := new(PoolSpecNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolSpecResource) DeepCopyInto(out *PoolSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AccountName != nil {
		in, out := &in.AccountName, &out.AccountName
		*out = new(string)
		**out = **in
	}
	if in.AutoScale != nil {
		in, out := &in.AutoScale, &out.AutoScale
		*out = new(PoolSpecAutoScale)
		(*in).DeepCopyInto(*out)
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = make([]PoolSpecCertificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ContainerConfiguration != nil {
		in, out := &in.ContainerConfiguration, &out.ContainerConfiguration
		*out = new(PoolSpecContainerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.FixedScale != nil {
		in, out := &in.FixedScale, &out.FixedScale
		*out = new(PoolSpecFixedScale)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxTasksPerNode != nil {
		in, out := &in.MaxTasksPerNode, &out.MaxTasksPerNode
		*out = new(int64)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkConfiguration != nil {
		in, out := &in.NetworkConfiguration, &out.NetworkConfiguration
		*out = new(PoolSpecNetworkConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeAgentSkuID != nil {
		in, out := &in.NodeAgentSkuID, &out.NodeAgentSkuID
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.StartTask != nil {
		in, out := &in.StartTask, &out.StartTask
		*out = new(PoolSpecStartTask)
		(*in).DeepCopyInto(*out)
	}
	if in.StopPendingResizeOperation != nil {
		in, out := &in.StopPendingResizeOperation, &out.StopPendingResizeOperation
		*out = new(bool)
		**out = **in
	}
	if in.StorageImageReference != nil {
		in, out := &in.StorageImageReference, &out.StorageImageReference
		*out = new(PoolSpecStorageImageReference)
		(*in).DeepCopyInto(*out)
	}
	if in.VmSize != nil {
		in, out := &in.VmSize, &out.VmSize
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolSpecResource.
func (in *PoolSpecResource) DeepCopy() *PoolSpecResource {
	if in == nil {
		return nil
	}
	out := new(PoolSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolSpecStartTask) DeepCopyInto(out *PoolSpecStartTask) {
	*out = *in
	if in.CommandLine != nil {
		in, out := &in.CommandLine, &out.CommandLine
		*out = new(string)
		**out = **in
	}
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.MaxTaskRetryCount != nil {
		in, out := &in.MaxTaskRetryCount, &out.MaxTaskRetryCount
		*out = new(int64)
		**out = **in
	}
	if in.ResourceFile != nil {
		in, out := &in.ResourceFile, &out.ResourceFile
		*out = make([]PoolSpecStartTaskResourceFile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UserIdentity != nil {
		in, out := &in.UserIdentity, &out.UserIdentity
		*out = new(PoolSpecStartTaskUserIdentity)
		(*in).DeepCopyInto(*out)
	}
	if in.WaitForSuccess != nil {
		in, out := &in.WaitForSuccess, &out.WaitForSuccess
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolSpecStartTask.
func (in *PoolSpecStartTask) DeepCopy() *PoolSpecStartTask {
	if in == nil {
		return nil
	}
	out := new(PoolSpecStartTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolSpecStartTaskResourceFile) DeepCopyInto(out *PoolSpecStartTaskResourceFile) {
	*out = *in
	if in.AutoStorageContainerName != nil {
		in, out := &in.AutoStorageContainerName, &out.AutoStorageContainerName
		*out = new(string)
		**out = **in
	}
	if in.BlobPrefix != nil {
		in, out := &in.BlobPrefix, &out.BlobPrefix
		*out = new(string)
		**out = **in
	}
	if in.FileMode != nil {
		in, out := &in.FileMode, &out.FileMode
		*out = new(string)
		**out = **in
	}
	if in.FilePath != nil {
		in, out := &in.FilePath, &out.FilePath
		*out = new(string)
		**out = **in
	}
	if in.HttpURL != nil {
		in, out := &in.HttpURL, &out.HttpURL
		*out = new(string)
		**out = **in
	}
	if in.StorageContainerURL != nil {
		in, out := &in.StorageContainerURL, &out.StorageContainerURL
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolSpecStartTaskResourceFile.
func (in *PoolSpecStartTaskResourceFile) DeepCopy() *PoolSpecStartTaskResourceFile {
	if in == nil {
		return nil
	}
	out := new(PoolSpecStartTaskResourceFile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolSpecStartTaskUserIdentity) DeepCopyInto(out *PoolSpecStartTaskUserIdentity) {
	*out = *in
	if in.AutoUser != nil {
		in, out := &in.AutoUser, &out.AutoUser
		*out = new(PoolSpecStartTaskUserIdentityAutoUser)
		(*in).DeepCopyInto(*out)
	}
	if in.UserName != nil {
		in, out := &in.UserName, &out.UserName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolSpecStartTaskUserIdentity.
func (in *PoolSpecStartTaskUserIdentity) DeepCopy() *PoolSpecStartTaskUserIdentity {
	if in == nil {
		return nil
	}
	out := new(PoolSpecStartTaskUserIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolSpecStartTaskUserIdentityAutoUser) DeepCopyInto(out *PoolSpecStartTaskUserIdentityAutoUser) {
	*out = *in
	if in.ElevationLevel != nil {
		in, out := &in.ElevationLevel, &out.ElevationLevel
		*out = new(string)
		**out = **in
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolSpecStartTaskUserIdentityAutoUser.
func (in *PoolSpecStartTaskUserIdentityAutoUser) DeepCopy() *PoolSpecStartTaskUserIdentityAutoUser {
	if in == nil {
		return nil
	}
	out := new(PoolSpecStartTaskUserIdentityAutoUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolSpecStorageImageReference) DeepCopyInto(out *PoolSpecStorageImageReference) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Offer != nil {
		in, out := &in.Offer, &out.Offer
		*out = new(string)
		**out = **in
	}
	if in.Publisher != nil {
		in, out := &in.Publisher, &out.Publisher
		*out = new(string)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolSpecStorageImageReference.
func (in *PoolSpecStorageImageReference) DeepCopy() *PoolSpecStorageImageReference {
	if in == nil {
		return nil
	}
	out := new(PoolSpecStorageImageReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolStatus) DeepCopyInto(out *PoolStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolStatus.
func (in *PoolStatus) DeepCopy() *PoolStatus {
	if in == nil {
		return nil
	}
	out := new(PoolStatus)
	in.DeepCopyInto(out)
	return out
}
