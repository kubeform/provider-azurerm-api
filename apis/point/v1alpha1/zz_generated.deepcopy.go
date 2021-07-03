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
func (in *ToSiteVPNGateway) DeepCopyInto(out *ToSiteVPNGateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToSiteVPNGateway.
func (in *ToSiteVPNGateway) DeepCopy() *ToSiteVPNGateway {
	if in == nil {
		return nil
	}
	out := new(ToSiteVPNGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ToSiteVPNGateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToSiteVPNGatewayList) DeepCopyInto(out *ToSiteVPNGatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ToSiteVPNGateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToSiteVPNGatewayList.
func (in *ToSiteVPNGatewayList) DeepCopy() *ToSiteVPNGatewayList {
	if in == nil {
		return nil
	}
	out := new(ToSiteVPNGatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ToSiteVPNGatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToSiteVPNGatewaySpec) DeepCopyInto(out *ToSiteVPNGatewaySpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(ToSiteVPNGatewaySpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToSiteVPNGatewaySpec.
func (in *ToSiteVPNGatewaySpec) DeepCopy() *ToSiteVPNGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(ToSiteVPNGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToSiteVPNGatewaySpecConnectionConfiguration) DeepCopyInto(out *ToSiteVPNGatewaySpecConnectionConfiguration) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		*out = new(ToSiteVPNGatewaySpecConnectionConfigurationRoute)
		(*in).DeepCopyInto(*out)
	}
	if in.VpnClientAddressPool != nil {
		in, out := &in.VpnClientAddressPool, &out.VpnClientAddressPool
		*out = new(ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToSiteVPNGatewaySpecConnectionConfiguration.
func (in *ToSiteVPNGatewaySpecConnectionConfiguration) DeepCopy() *ToSiteVPNGatewaySpecConnectionConfiguration {
	if in == nil {
		return nil
	}
	out := new(ToSiteVPNGatewaySpecConnectionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToSiteVPNGatewaySpecConnectionConfigurationRoute) DeepCopyInto(out *ToSiteVPNGatewaySpecConnectionConfigurationRoute) {
	*out = *in
	if in.AssociatedRouteTableID != nil {
		in, out := &in.AssociatedRouteTableID, &out.AssociatedRouteTableID
		*out = new(string)
		**out = **in
	}
	if in.PropagatedRouteTable != nil {
		in, out := &in.PropagatedRouteTable, &out.PropagatedRouteTable
		*out = new(ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToSiteVPNGatewaySpecConnectionConfigurationRoute.
func (in *ToSiteVPNGatewaySpecConnectionConfigurationRoute) DeepCopy() *ToSiteVPNGatewaySpecConnectionConfigurationRoute {
	if in == nil {
		return nil
	}
	out := new(ToSiteVPNGatewaySpecConnectionConfigurationRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable) DeepCopyInto(out *ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable) {
	*out = *in
	if in.Ids != nil {
		in, out := &in.Ids, &out.Ids
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable.
func (in *ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable) DeepCopy() *ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable {
	if in == nil {
		return nil
	}
	out := new(ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool) DeepCopyInto(out *ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool) {
	*out = *in
	if in.AddressPrefixes != nil {
		in, out := &in.AddressPrefixes, &out.AddressPrefixes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool.
func (in *ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool) DeepCopy() *ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool {
	if in == nil {
		return nil
	}
	out := new(ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToSiteVPNGatewaySpecResource) DeepCopyInto(out *ToSiteVPNGatewaySpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.ConnectionConfiguration != nil {
		in, out := &in.ConnectionConfiguration, &out.ConnectionConfiguration
		*out = new(ToSiteVPNGatewaySpecConnectionConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.DnsServers != nil {
		in, out := &in.DnsServers, &out.DnsServers
		*out = make([]string, len(*in))
		copy(*out, *in)
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
	if in.ScaleUnit != nil {
		in, out := &in.ScaleUnit, &out.ScaleUnit
		*out = new(int64)
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
	if in.VirtualHubID != nil {
		in, out := &in.VirtualHubID, &out.VirtualHubID
		*out = new(string)
		**out = **in
	}
	if in.VpnServerConfigurationID != nil {
		in, out := &in.VpnServerConfigurationID, &out.VpnServerConfigurationID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToSiteVPNGatewaySpecResource.
func (in *ToSiteVPNGatewaySpecResource) DeepCopy() *ToSiteVPNGatewaySpecResource {
	if in == nil {
		return nil
	}
	out := new(ToSiteVPNGatewaySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToSiteVPNGatewayStatus) DeepCopyInto(out *ToSiteVPNGatewayStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToSiteVPNGatewayStatus.
func (in *ToSiteVPNGatewayStatus) DeepCopy() *ToSiteVPNGatewayStatus {
	if in == nil {
		return nil
	}
	out := new(ToSiteVPNGatewayStatus)
	in.DeepCopyInto(out)
	return out
}
