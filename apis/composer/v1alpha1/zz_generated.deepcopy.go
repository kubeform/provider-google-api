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
func (in *Environment) DeepCopyInto(out *Environment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Environment.
func (in *Environment) DeepCopy() *Environment {
	if in == nil {
		return nil
	}
	out := new(Environment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Environment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentList) DeepCopyInto(out *EnvironmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Environment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentList.
func (in *EnvironmentList) DeepCopy() *EnvironmentList {
	if in == nil {
		return nil
	}
	out := new(EnvironmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvironmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpec) DeepCopyInto(out *EnvironmentSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(EnvironmentSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpec.
func (in *EnvironmentSpec) DeepCopy() *EnvironmentSpec {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpecConfig) DeepCopyInto(out *EnvironmentSpecConfig) {
	*out = *in
	if in.AirflowURI != nil {
		in, out := &in.AirflowURI, &out.AirflowURI
		*out = new(string)
		**out = **in
	}
	if in.DagGcsPrefix != nil {
		in, out := &in.DagGcsPrefix, &out.DagGcsPrefix
		*out = new(string)
		**out = **in
	}
	if in.GkeCluster != nil {
		in, out := &in.GkeCluster, &out.GkeCluster
		*out = new(string)
		**out = **in
	}
	if in.NodeConfig != nil {
		in, out := &in.NodeConfig, &out.NodeConfig
		*out = new(EnvironmentSpecConfigNodeConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeCount != nil {
		in, out := &in.NodeCount, &out.NodeCount
		*out = new(int64)
		**out = **in
	}
	if in.PrivateEnvironmentConfig != nil {
		in, out := &in.PrivateEnvironmentConfig, &out.PrivateEnvironmentConfig
		*out = new(EnvironmentSpecConfigPrivateEnvironmentConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.SoftwareConfig != nil {
		in, out := &in.SoftwareConfig, &out.SoftwareConfig
		*out = new(EnvironmentSpecConfigSoftwareConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpecConfig.
func (in *EnvironmentSpecConfig) DeepCopy() *EnvironmentSpecConfig {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpecConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpecConfigNodeConfig) DeepCopyInto(out *EnvironmentSpecConfigNodeConfig) {
	*out = *in
	if in.DiskSizeGb != nil {
		in, out := &in.DiskSizeGb, &out.DiskSizeGb
		*out = new(int64)
		**out = **in
	}
	if in.IpAllocationPolicy != nil {
		in, out := &in.IpAllocationPolicy, &out.IpAllocationPolicy
		*out = new(EnvironmentSpecConfigNodeConfigIpAllocationPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.MachineType != nil {
		in, out := &in.MachineType, &out.MachineType
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
	if in.OauthScopes != nil {
		in, out := &in.OauthScopes, &out.OauthScopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(string)
		**out = **in
	}
	if in.Subnetwork != nil {
		in, out := &in.Subnetwork, &out.Subnetwork
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpecConfigNodeConfig.
func (in *EnvironmentSpecConfigNodeConfig) DeepCopy() *EnvironmentSpecConfigNodeConfig {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpecConfigNodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpecConfigNodeConfigIpAllocationPolicy) DeepCopyInto(out *EnvironmentSpecConfigNodeConfigIpAllocationPolicy) {
	*out = *in
	if in.ClusterIpv4CIDRBlock != nil {
		in, out := &in.ClusterIpv4CIDRBlock, &out.ClusterIpv4CIDRBlock
		*out = new(string)
		**out = **in
	}
	if in.ClusterSecondaryRangeName != nil {
		in, out := &in.ClusterSecondaryRangeName, &out.ClusterSecondaryRangeName
		*out = new(string)
		**out = **in
	}
	if in.ServicesIpv4CIDRBlock != nil {
		in, out := &in.ServicesIpv4CIDRBlock, &out.ServicesIpv4CIDRBlock
		*out = new(string)
		**out = **in
	}
	if in.ServicesSecondaryRangeName != nil {
		in, out := &in.ServicesSecondaryRangeName, &out.ServicesSecondaryRangeName
		*out = new(string)
		**out = **in
	}
	if in.UseIPAliases != nil {
		in, out := &in.UseIPAliases, &out.UseIPAliases
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpecConfigNodeConfigIpAllocationPolicy.
func (in *EnvironmentSpecConfigNodeConfigIpAllocationPolicy) DeepCopy() *EnvironmentSpecConfigNodeConfigIpAllocationPolicy {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpecConfigNodeConfigIpAllocationPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpecConfigPrivateEnvironmentConfig) DeepCopyInto(out *EnvironmentSpecConfigPrivateEnvironmentConfig) {
	*out = *in
	if in.CloudSQLIpv4CIDRBlock != nil {
		in, out := &in.CloudSQLIpv4CIDRBlock, &out.CloudSQLIpv4CIDRBlock
		*out = new(string)
		**out = **in
	}
	if in.EnablePrivateEndpoint != nil {
		in, out := &in.EnablePrivateEndpoint, &out.EnablePrivateEndpoint
		*out = new(bool)
		**out = **in
	}
	if in.MasterIpv4CIDRBlock != nil {
		in, out := &in.MasterIpv4CIDRBlock, &out.MasterIpv4CIDRBlock
		*out = new(string)
		**out = **in
	}
	if in.WebServerIpv4CIDRBlock != nil {
		in, out := &in.WebServerIpv4CIDRBlock, &out.WebServerIpv4CIDRBlock
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpecConfigPrivateEnvironmentConfig.
func (in *EnvironmentSpecConfigPrivateEnvironmentConfig) DeepCopy() *EnvironmentSpecConfigPrivateEnvironmentConfig {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpecConfigPrivateEnvironmentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpecConfigSoftwareConfig) DeepCopyInto(out *EnvironmentSpecConfigSoftwareConfig) {
	*out = *in
	if in.AirflowConfigOverrides != nil {
		in, out := &in.AirflowConfigOverrides, &out.AirflowConfigOverrides
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.EnvVariables != nil {
		in, out := &in.EnvVariables, &out.EnvVariables
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.ImageVersion != nil {
		in, out := &in.ImageVersion, &out.ImageVersion
		*out = new(string)
		**out = **in
	}
	if in.PypiPackages != nil {
		in, out := &in.PypiPackages, &out.PypiPackages
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.PythonVersion != nil {
		in, out := &in.PythonVersion, &out.PythonVersion
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpecConfigSoftwareConfig.
func (in *EnvironmentSpecConfigSoftwareConfig) DeepCopy() *EnvironmentSpecConfigSoftwareConfig {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpecConfigSoftwareConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpecResource) DeepCopyInto(out *EnvironmentSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(EnvironmentSpecConfig)
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
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpecResource.
func (in *EnvironmentSpecResource) DeepCopy() *EnvironmentSpecResource {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentStatus) DeepCopyInto(out *EnvironmentStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentStatus.
func (in *EnvironmentStatus) DeepCopy() *EnvironmentStatus {
	if in == nil {
		return nil
	}
	out := new(EnvironmentStatus)
	in.DeepCopyInto(out)
	return out
}
