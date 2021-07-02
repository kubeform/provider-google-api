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
func (in *ManagementConnectivityTest) DeepCopyInto(out *ManagementConnectivityTest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementConnectivityTest.
func (in *ManagementConnectivityTest) DeepCopy() *ManagementConnectivityTest {
	if in == nil {
		return nil
	}
	out := new(ManagementConnectivityTest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagementConnectivityTest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementConnectivityTestList) DeepCopyInto(out *ManagementConnectivityTestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagementConnectivityTest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementConnectivityTestList.
func (in *ManagementConnectivityTestList) DeepCopy() *ManagementConnectivityTestList {
	if in == nil {
		return nil
	}
	out := new(ManagementConnectivityTestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagementConnectivityTestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementConnectivityTestSpec) DeepCopyInto(out *ManagementConnectivityTestSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(ManagementConnectivityTestSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementConnectivityTestSpec.
func (in *ManagementConnectivityTestSpec) DeepCopy() *ManagementConnectivityTestSpec {
	if in == nil {
		return nil
	}
	out := new(ManagementConnectivityTestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementConnectivityTestSpecDestination) DeepCopyInto(out *ManagementConnectivityTestSpecDestination) {
	*out = *in
	if in.Instance != nil {
		in, out := &in.Instance, &out.Instance
		*out = new(string)
		**out = **in
	}
	if in.IpAddress != nil {
		in, out := &in.IpAddress, &out.IpAddress
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementConnectivityTestSpecDestination.
func (in *ManagementConnectivityTestSpecDestination) DeepCopy() *ManagementConnectivityTestSpecDestination {
	if in == nil {
		return nil
	}
	out := new(ManagementConnectivityTestSpecDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementConnectivityTestSpecResource) DeepCopyInto(out *ManagementConnectivityTestSpecResource) {
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
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		*out = new(ManagementConnectivityTestSpecDestination)
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
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.RelatedProjects != nil {
		in, out := &in.RelatedProjects, &out.RelatedProjects
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(ManagementConnectivityTestSpecSource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementConnectivityTestSpecResource.
func (in *ManagementConnectivityTestSpecResource) DeepCopy() *ManagementConnectivityTestSpecResource {
	if in == nil {
		return nil
	}
	out := new(ManagementConnectivityTestSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementConnectivityTestSpecSource) DeepCopyInto(out *ManagementConnectivityTestSpecSource) {
	*out = *in
	if in.Instance != nil {
		in, out := &in.Instance, &out.Instance
		*out = new(string)
		**out = **in
	}
	if in.IpAddress != nil {
		in, out := &in.IpAddress, &out.IpAddress
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
	if in.NetworkType != nil {
		in, out := &in.NetworkType, &out.NetworkType
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementConnectivityTestSpecSource.
func (in *ManagementConnectivityTestSpecSource) DeepCopy() *ManagementConnectivityTestSpecSource {
	if in == nil {
		return nil
	}
	out := new(ManagementConnectivityTestSpecSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementConnectivityTestStatus) DeepCopyInto(out *ManagementConnectivityTestStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementConnectivityTestStatus.
func (in *ManagementConnectivityTestStatus) DeepCopy() *ManagementConnectivityTestStatus {
	if in == nil {
		return nil
	}
	out := new(ManagementConnectivityTestStatus)
	in.DeepCopyInto(out)
	return out
}
