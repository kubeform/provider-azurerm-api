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
func (in *Snapshot) DeepCopyInto(out *Snapshot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Snapshot.
func (in *Snapshot) DeepCopy() *Snapshot {
	if in == nil {
		return nil
	}
	out := new(Snapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Snapshot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotList) DeepCopyInto(out *SnapshotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Snapshot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotList.
func (in *SnapshotList) DeepCopy() *SnapshotList {
	if in == nil {
		return nil
	}
	out := new(SnapshotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnapshotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotSpec) DeepCopyInto(out *SnapshotSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(SnapshotSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotSpec.
func (in *SnapshotSpec) DeepCopy() *SnapshotSpec {
	if in == nil {
		return nil
	}
	out := new(SnapshotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotSpecEncryptionSettings) DeepCopyInto(out *SnapshotSpecEncryptionSettings) {
	*out = *in
	if in.DiskEncryptionKey != nil {
		in, out := &in.DiskEncryptionKey, &out.DiskEncryptionKey
		*out = new(SnapshotSpecEncryptionSettingsDiskEncryptionKey)
		(*in).DeepCopyInto(*out)
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.KeyEncryptionKey != nil {
		in, out := &in.KeyEncryptionKey, &out.KeyEncryptionKey
		*out = new(SnapshotSpecEncryptionSettingsKeyEncryptionKey)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotSpecEncryptionSettings.
func (in *SnapshotSpecEncryptionSettings) DeepCopy() *SnapshotSpecEncryptionSettings {
	if in == nil {
		return nil
	}
	out := new(SnapshotSpecEncryptionSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotSpecEncryptionSettingsDiskEncryptionKey) DeepCopyInto(out *SnapshotSpecEncryptionSettingsDiskEncryptionKey) {
	*out = *in
	if in.SecretURL != nil {
		in, out := &in.SecretURL, &out.SecretURL
		*out = new(string)
		**out = **in
	}
	if in.SourceVaultID != nil {
		in, out := &in.SourceVaultID, &out.SourceVaultID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotSpecEncryptionSettingsDiskEncryptionKey.
func (in *SnapshotSpecEncryptionSettingsDiskEncryptionKey) DeepCopy() *SnapshotSpecEncryptionSettingsDiskEncryptionKey {
	if in == nil {
		return nil
	}
	out := new(SnapshotSpecEncryptionSettingsDiskEncryptionKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotSpecEncryptionSettingsKeyEncryptionKey) DeepCopyInto(out *SnapshotSpecEncryptionSettingsKeyEncryptionKey) {
	*out = *in
	if in.KeyURL != nil {
		in, out := &in.KeyURL, &out.KeyURL
		*out = new(string)
		**out = **in
	}
	if in.SourceVaultID != nil {
		in, out := &in.SourceVaultID, &out.SourceVaultID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotSpecEncryptionSettingsKeyEncryptionKey.
func (in *SnapshotSpecEncryptionSettingsKeyEncryptionKey) DeepCopy() *SnapshotSpecEncryptionSettingsKeyEncryptionKey {
	if in == nil {
		return nil
	}
	out := new(SnapshotSpecEncryptionSettingsKeyEncryptionKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotSpecResource) DeepCopyInto(out *SnapshotSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.CreateOption != nil {
		in, out := &in.CreateOption, &out.CreateOption
		*out = new(string)
		**out = **in
	}
	if in.DiskSizeGb != nil {
		in, out := &in.DiskSizeGb, &out.DiskSizeGb
		*out = new(int64)
		**out = **in
	}
	if in.EncryptionSettings != nil {
		in, out := &in.EncryptionSettings, &out.EncryptionSettings
		*out = new(SnapshotSpecEncryptionSettings)
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
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SourceResourceID != nil {
		in, out := &in.SourceResourceID, &out.SourceResourceID
		*out = new(string)
		**out = **in
	}
	if in.SourceURI != nil {
		in, out := &in.SourceURI, &out.SourceURI
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotSpecResource.
func (in *SnapshotSpecResource) DeepCopy() *SnapshotSpecResource {
	if in == nil {
		return nil
	}
	out := new(SnapshotSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotStatus) DeepCopyInto(out *SnapshotStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotStatus.
func (in *SnapshotStatus) DeepCopy() *SnapshotStatus {
	if in == nil {
		return nil
	}
	out := new(SnapshotStatus)
	in.DeepCopyInto(out)
	return out
}
