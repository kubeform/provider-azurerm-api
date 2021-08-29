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

	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkGateway) DeepCopyInto(out *NetworkGateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkGateway.
func (in *NetworkGateway) DeepCopy() *NetworkGateway {
	if in == nil {
		return nil
	}
	out := new(NetworkGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkGateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkGatewayList) DeepCopyInto(out *NetworkGatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkGateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkGatewayList.
func (in *NetworkGatewayList) DeepCopy() *NetworkGatewayList {
	if in == nil {
		return nil
	}
	out := new(NetworkGatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkGatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkGatewaySpec) DeepCopyInto(out *NetworkGatewaySpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(NetworkGatewaySpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkGatewaySpec.
func (in *NetworkGatewaySpec) DeepCopy() *NetworkGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(NetworkGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkGatewaySpecBgpSettings) DeepCopyInto(out *NetworkGatewaySpecBgpSettings) {
	*out = *in
	if in.Asn != nil {
		in, out := &in.Asn, &out.Asn
		*out = new(int64)
		**out = **in
	}
	if in.BgpPeeringAddress != nil {
		in, out := &in.BgpPeeringAddress, &out.BgpPeeringAddress
		*out = new(string)
		**out = **in
	}
	if in.PeerWeight != nil {
		in, out := &in.PeerWeight, &out.PeerWeight
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkGatewaySpecBgpSettings.
func (in *NetworkGatewaySpecBgpSettings) DeepCopy() *NetworkGatewaySpecBgpSettings {
	if in == nil {
		return nil
	}
	out := new(NetworkGatewaySpecBgpSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkGatewaySpecResource) DeepCopyInto(out *NetworkGatewaySpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AddressSpace != nil {
		in, out := &in.AddressSpace, &out.AddressSpace
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BgpSettings != nil {
		in, out := &in.BgpSettings, &out.BgpSettings
		*out = new(NetworkGatewaySpecBgpSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.GatewayAddress != nil {
		in, out := &in.GatewayAddress, &out.GatewayAddress
		*out = new(string)
		**out = **in
	}
	if in.GatewayFqdn != nil {
		in, out := &in.GatewayFqdn, &out.GatewayFqdn
		*out = new(string)
		**out = **in
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkGatewaySpecResource.
func (in *NetworkGatewaySpecResource) DeepCopy() *NetworkGatewaySpecResource {
	if in == nil {
		return nil
	}
	out := new(NetworkGatewaySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkGatewayStatus) DeepCopyInto(out *NetworkGatewayStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkGatewayStatus.
func (in *NetworkGatewayStatus) DeepCopy() *NetworkGatewayStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkGatewayStatus)
	in.DeepCopyInto(out)
	return out
}
