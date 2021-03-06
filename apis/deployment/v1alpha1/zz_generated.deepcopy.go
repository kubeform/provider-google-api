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
func (in *ManagerDeployment) DeepCopyInto(out *ManagerDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerDeployment.
func (in *ManagerDeployment) DeepCopy() *ManagerDeployment {
	if in == nil {
		return nil
	}
	out := new(ManagerDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagerDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerDeploymentList) DeepCopyInto(out *ManagerDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagerDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerDeploymentList.
func (in *ManagerDeploymentList) DeepCopy() *ManagerDeploymentList {
	if in == nil {
		return nil
	}
	out := new(ManagerDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagerDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerDeploymentSpec) DeepCopyInto(out *ManagerDeploymentSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ManagerDeploymentSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerDeploymentSpec.
func (in *ManagerDeploymentSpec) DeepCopy() *ManagerDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(ManagerDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerDeploymentSpecLabels) DeepCopyInto(out *ManagerDeploymentSpecLabels) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerDeploymentSpecLabels.
func (in *ManagerDeploymentSpecLabels) DeepCopy() *ManagerDeploymentSpecLabels {
	if in == nil {
		return nil
	}
	out := new(ManagerDeploymentSpecLabels)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerDeploymentSpecResource) DeepCopyInto(out *ManagerDeploymentSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.CreatePolicy != nil {
		in, out := &in.CreatePolicy, &out.CreatePolicy
		*out = new(string)
		**out = **in
	}
	if in.DeletePolicy != nil {
		in, out := &in.DeletePolicy, &out.DeletePolicy
		*out = new(string)
		**out = **in
	}
	if in.DeploymentID != nil {
		in, out := &in.DeploymentID, &out.DeploymentID
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]ManagerDeploymentSpecLabels, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Manifest != nil {
		in, out := &in.Manifest, &out.Manifest
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Preview != nil {
		in, out := &in.Preview, &out.Preview
		*out = new(bool)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.SelfLink != nil {
		in, out := &in.SelfLink, &out.SelfLink
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(ManagerDeploymentSpecTarget)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerDeploymentSpecResource.
func (in *ManagerDeploymentSpecResource) DeepCopy() *ManagerDeploymentSpecResource {
	if in == nil {
		return nil
	}
	out := new(ManagerDeploymentSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerDeploymentSpecTarget) DeepCopyInto(out *ManagerDeploymentSpecTarget) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(ManagerDeploymentSpecTargetConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Imports != nil {
		in, out := &in.Imports, &out.Imports
		*out = make([]ManagerDeploymentSpecTargetImports, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerDeploymentSpecTarget.
func (in *ManagerDeploymentSpecTarget) DeepCopy() *ManagerDeploymentSpecTarget {
	if in == nil {
		return nil
	}
	out := new(ManagerDeploymentSpecTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerDeploymentSpecTargetConfig) DeepCopyInto(out *ManagerDeploymentSpecTargetConfig) {
	*out = *in
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerDeploymentSpecTargetConfig.
func (in *ManagerDeploymentSpecTargetConfig) DeepCopy() *ManagerDeploymentSpecTargetConfig {
	if in == nil {
		return nil
	}
	out := new(ManagerDeploymentSpecTargetConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerDeploymentSpecTargetImports) DeepCopyInto(out *ManagerDeploymentSpecTargetImports) {
	*out = *in
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerDeploymentSpecTargetImports.
func (in *ManagerDeploymentSpecTargetImports) DeepCopy() *ManagerDeploymentSpecTargetImports {
	if in == nil {
		return nil
	}
	out := new(ManagerDeploymentSpecTargetImports)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerDeploymentStatus) DeepCopyInto(out *ManagerDeploymentStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerDeploymentStatus.
func (in *ManagerDeploymentStatus) DeepCopy() *ManagerDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(ManagerDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}
