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
func (in *ManagedZone) DeepCopyInto(out *ManagedZone) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedZone.
func (in *ManagedZone) DeepCopy() *ManagedZone {
	if in == nil {
		return nil
	}
	out := new(ManagedZone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedZone) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedZoneList) DeepCopyInto(out *ManagedZoneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedZone, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedZoneList.
func (in *ManagedZoneList) DeepCopy() *ManagedZoneList {
	if in == nil {
		return nil
	}
	out := new(ManagedZoneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedZoneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedZoneSpec) DeepCopyInto(out *ManagedZoneSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ManagedZoneSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedZoneSpec.
func (in *ManagedZoneSpec) DeepCopy() *ManagedZoneSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedZoneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedZoneSpecDnssecConfig) DeepCopyInto(out *ManagedZoneSpecDnssecConfig) {
	*out = *in
	if in.DefaultKeySpecs != nil {
		in, out := &in.DefaultKeySpecs, &out.DefaultKeySpecs
		*out = make([]ManagedZoneSpecDnssecConfigDefaultKeySpecs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	if in.NonExistence != nil {
		in, out := &in.NonExistence, &out.NonExistence
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedZoneSpecDnssecConfig.
func (in *ManagedZoneSpecDnssecConfig) DeepCopy() *ManagedZoneSpecDnssecConfig {
	if in == nil {
		return nil
	}
	out := new(ManagedZoneSpecDnssecConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedZoneSpecDnssecConfigDefaultKeySpecs) DeepCopyInto(out *ManagedZoneSpecDnssecConfigDefaultKeySpecs) {
	*out = *in
	if in.Algorithm != nil {
		in, out := &in.Algorithm, &out.Algorithm
		*out = new(string)
		**out = **in
	}
	if in.KeyLength != nil {
		in, out := &in.KeyLength, &out.KeyLength
		*out = new(int64)
		**out = **in
	}
	if in.KeyType != nil {
		in, out := &in.KeyType, &out.KeyType
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedZoneSpecDnssecConfigDefaultKeySpecs.
func (in *ManagedZoneSpecDnssecConfigDefaultKeySpecs) DeepCopy() *ManagedZoneSpecDnssecConfigDefaultKeySpecs {
	if in == nil {
		return nil
	}
	out := new(ManagedZoneSpecDnssecConfigDefaultKeySpecs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedZoneSpecForwardingConfig) DeepCopyInto(out *ManagedZoneSpecForwardingConfig) {
	*out = *in
	if in.TargetNameServers != nil {
		in, out := &in.TargetNameServers, &out.TargetNameServers
		*out = make([]ManagedZoneSpecForwardingConfigTargetNameServers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedZoneSpecForwardingConfig.
func (in *ManagedZoneSpecForwardingConfig) DeepCopy() *ManagedZoneSpecForwardingConfig {
	if in == nil {
		return nil
	}
	out := new(ManagedZoneSpecForwardingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedZoneSpecForwardingConfigTargetNameServers) DeepCopyInto(out *ManagedZoneSpecForwardingConfigTargetNameServers) {
	*out = *in
	if in.ForwardingPath != nil {
		in, out := &in.ForwardingPath, &out.ForwardingPath
		*out = new(string)
		**out = **in
	}
	if in.Ipv4Address != nil {
		in, out := &in.Ipv4Address, &out.Ipv4Address
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedZoneSpecForwardingConfigTargetNameServers.
func (in *ManagedZoneSpecForwardingConfigTargetNameServers) DeepCopy() *ManagedZoneSpecForwardingConfigTargetNameServers {
	if in == nil {
		return nil
	}
	out := new(ManagedZoneSpecForwardingConfigTargetNameServers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedZoneSpecPeeringConfig) DeepCopyInto(out *ManagedZoneSpecPeeringConfig) {
	*out = *in
	if in.TargetNetwork != nil {
		in, out := &in.TargetNetwork, &out.TargetNetwork
		*out = new(ManagedZoneSpecPeeringConfigTargetNetwork)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedZoneSpecPeeringConfig.
func (in *ManagedZoneSpecPeeringConfig) DeepCopy() *ManagedZoneSpecPeeringConfig {
	if in == nil {
		return nil
	}
	out := new(ManagedZoneSpecPeeringConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedZoneSpecPeeringConfigTargetNetwork) DeepCopyInto(out *ManagedZoneSpecPeeringConfigTargetNetwork) {
	*out = *in
	if in.NetworkURL != nil {
		in, out := &in.NetworkURL, &out.NetworkURL
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedZoneSpecPeeringConfigTargetNetwork.
func (in *ManagedZoneSpecPeeringConfigTargetNetwork) DeepCopy() *ManagedZoneSpecPeeringConfigTargetNetwork {
	if in == nil {
		return nil
	}
	out := new(ManagedZoneSpecPeeringConfigTargetNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedZoneSpecPrivateVisibilityConfig) DeepCopyInto(out *ManagedZoneSpecPrivateVisibilityConfig) {
	*out = *in
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]ManagedZoneSpecPrivateVisibilityConfigNetworks, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedZoneSpecPrivateVisibilityConfig.
func (in *ManagedZoneSpecPrivateVisibilityConfig) DeepCopy() *ManagedZoneSpecPrivateVisibilityConfig {
	if in == nil {
		return nil
	}
	out := new(ManagedZoneSpecPrivateVisibilityConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedZoneSpecPrivateVisibilityConfigNetworks) DeepCopyInto(out *ManagedZoneSpecPrivateVisibilityConfigNetworks) {
	*out = *in
	if in.NetworkURL != nil {
		in, out := &in.NetworkURL, &out.NetworkURL
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedZoneSpecPrivateVisibilityConfigNetworks.
func (in *ManagedZoneSpecPrivateVisibilityConfigNetworks) DeepCopy() *ManagedZoneSpecPrivateVisibilityConfigNetworks {
	if in == nil {
		return nil
	}
	out := new(ManagedZoneSpecPrivateVisibilityConfigNetworks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedZoneSpecResource) DeepCopyInto(out *ManagedZoneSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DnsName != nil {
		in, out := &in.DnsName, &out.DnsName
		*out = new(string)
		**out = **in
	}
	if in.DnssecConfig != nil {
		in, out := &in.DnssecConfig, &out.DnssecConfig
		*out = new(ManagedZoneSpecDnssecConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ForceDestroy != nil {
		in, out := &in.ForceDestroy, &out.ForceDestroy
		*out = new(bool)
		**out = **in
	}
	if in.ForwardingConfig != nil {
		in, out := &in.ForwardingConfig, &out.ForwardingConfig
		*out = new(ManagedZoneSpecForwardingConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
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
	if in.NameServers != nil {
		in, out := &in.NameServers, &out.NameServers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PeeringConfig != nil {
		in, out := &in.PeeringConfig, &out.PeeringConfig
		*out = new(ManagedZoneSpecPeeringConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivateVisibilityConfig != nil {
		in, out := &in.PrivateVisibilityConfig, &out.PrivateVisibilityConfig
		*out = new(ManagedZoneSpecPrivateVisibilityConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedZoneSpecResource.
func (in *ManagedZoneSpecResource) DeepCopy() *ManagedZoneSpecResource {
	if in == nil {
		return nil
	}
	out := new(ManagedZoneSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedZoneStatus) DeepCopyInto(out *ManagedZoneStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedZoneStatus.
func (in *ManagedZoneStatus) DeepCopy() *ManagedZoneStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedZoneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Policy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyList) DeepCopyInto(out *PolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Policy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyList.
func (in *PolicyList) DeepCopy() *PolicyList {
	if in == nil {
		return nil
	}
	out := new(PolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpec) DeepCopyInto(out *PolicySpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(PolicySpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpec.
func (in *PolicySpec) DeepCopy() *PolicySpec {
	if in == nil {
		return nil
	}
	out := new(PolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpecAlternativeNameServerConfig) DeepCopyInto(out *PolicySpecAlternativeNameServerConfig) {
	*out = *in
	if in.TargetNameServers != nil {
		in, out := &in.TargetNameServers, &out.TargetNameServers
		*out = make([]PolicySpecAlternativeNameServerConfigTargetNameServers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpecAlternativeNameServerConfig.
func (in *PolicySpecAlternativeNameServerConfig) DeepCopy() *PolicySpecAlternativeNameServerConfig {
	if in == nil {
		return nil
	}
	out := new(PolicySpecAlternativeNameServerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpecAlternativeNameServerConfigTargetNameServers) DeepCopyInto(out *PolicySpecAlternativeNameServerConfigTargetNameServers) {
	*out = *in
	if in.ForwardingPath != nil {
		in, out := &in.ForwardingPath, &out.ForwardingPath
		*out = new(string)
		**out = **in
	}
	if in.Ipv4Address != nil {
		in, out := &in.Ipv4Address, &out.Ipv4Address
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpecAlternativeNameServerConfigTargetNameServers.
func (in *PolicySpecAlternativeNameServerConfigTargetNameServers) DeepCopy() *PolicySpecAlternativeNameServerConfigTargetNameServers {
	if in == nil {
		return nil
	}
	out := new(PolicySpecAlternativeNameServerConfigTargetNameServers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpecNetworks) DeepCopyInto(out *PolicySpecNetworks) {
	*out = *in
	if in.NetworkURL != nil {
		in, out := &in.NetworkURL, &out.NetworkURL
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpecNetworks.
func (in *PolicySpecNetworks) DeepCopy() *PolicySpecNetworks {
	if in == nil {
		return nil
	}
	out := new(PolicySpecNetworks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpecResource) DeepCopyInto(out *PolicySpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AlternativeNameServerConfig != nil {
		in, out := &in.AlternativeNameServerConfig, &out.AlternativeNameServerConfig
		*out = new(PolicySpecAlternativeNameServerConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EnableInboundForwarding != nil {
		in, out := &in.EnableInboundForwarding, &out.EnableInboundForwarding
		*out = new(bool)
		**out = **in
	}
	if in.EnableLogging != nil {
		in, out := &in.EnableLogging, &out.EnableLogging
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]PolicySpecNetworks, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpecResource.
func (in *PolicySpecResource) DeepCopy() *PolicySpecResource {
	if in == nil {
		return nil
	}
	out := new(PolicySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyStatus) DeepCopyInto(out *PolicyStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyStatus.
func (in *PolicyStatus) DeepCopy() *PolicyStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordSet) DeepCopyInto(out *RecordSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordSet.
func (in *RecordSet) DeepCopy() *RecordSet {
	if in == nil {
		return nil
	}
	out := new(RecordSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RecordSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordSetList) DeepCopyInto(out *RecordSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RecordSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordSetList.
func (in *RecordSetList) DeepCopy() *RecordSetList {
	if in == nil {
		return nil
	}
	out := new(RecordSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RecordSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordSetSpec) DeepCopyInto(out *RecordSetSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(RecordSetSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordSetSpec.
func (in *RecordSetSpec) DeepCopy() *RecordSetSpec {
	if in == nil {
		return nil
	}
	out := new(RecordSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordSetSpecResource) DeepCopyInto(out *RecordSetSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.ManagedZone != nil {
		in, out := &in.ManagedZone, &out.ManagedZone
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Rrdatas != nil {
		in, out := &in.Rrdatas, &out.Rrdatas
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ttl != nil {
		in, out := &in.Ttl, &out.Ttl
		*out = new(int64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordSetSpecResource.
func (in *RecordSetSpecResource) DeepCopy() *RecordSetSpecResource {
	if in == nil {
		return nil
	}
	out := new(RecordSetSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordSetStatus) DeepCopyInto(out *RecordSetStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordSetStatus.
func (in *RecordSetStatus) DeepCopy() *RecordSetStatus {
	if in == nil {
		return nil
	}
	out := new(RecordSetStatus)
	in.DeepCopyInto(out)
	return out
}
