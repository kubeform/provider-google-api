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
func (in *CloudfunctionsFunction) DeepCopyInto(out *CloudfunctionsFunction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudfunctionsFunction.
func (in *CloudfunctionsFunction) DeepCopy() *CloudfunctionsFunction {
	if in == nil {
		return nil
	}
	out := new(CloudfunctionsFunction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudfunctionsFunction) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudfunctionsFunctionList) DeepCopyInto(out *CloudfunctionsFunctionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudfunctionsFunction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudfunctionsFunctionList.
func (in *CloudfunctionsFunctionList) DeepCopy() *CloudfunctionsFunctionList {
	if in == nil {
		return nil
	}
	out := new(CloudfunctionsFunctionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudfunctionsFunctionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudfunctionsFunctionSpec) DeepCopyInto(out *CloudfunctionsFunctionSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(CloudfunctionsFunctionSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudfunctionsFunctionSpec.
func (in *CloudfunctionsFunctionSpec) DeepCopy() *CloudfunctionsFunctionSpec {
	if in == nil {
		return nil
	}
	out := new(CloudfunctionsFunctionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudfunctionsFunctionSpecEventTrigger) DeepCopyInto(out *CloudfunctionsFunctionSpecEventTrigger) {
	*out = *in
	if in.EventType != nil {
		in, out := &in.EventType, &out.EventType
		*out = new(string)
		**out = **in
	}
	if in.FailurePolicy != nil {
		in, out := &in.FailurePolicy, &out.FailurePolicy
		*out = new(CloudfunctionsFunctionSpecEventTriggerFailurePolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudfunctionsFunctionSpecEventTrigger.
func (in *CloudfunctionsFunctionSpecEventTrigger) DeepCopy() *CloudfunctionsFunctionSpecEventTrigger {
	if in == nil {
		return nil
	}
	out := new(CloudfunctionsFunctionSpecEventTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudfunctionsFunctionSpecEventTriggerFailurePolicy) DeepCopyInto(out *CloudfunctionsFunctionSpecEventTriggerFailurePolicy) {
	*out = *in
	if in.Retry != nil {
		in, out := &in.Retry, &out.Retry
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudfunctionsFunctionSpecEventTriggerFailurePolicy.
func (in *CloudfunctionsFunctionSpecEventTriggerFailurePolicy) DeepCopy() *CloudfunctionsFunctionSpecEventTriggerFailurePolicy {
	if in == nil {
		return nil
	}
	out := new(CloudfunctionsFunctionSpecEventTriggerFailurePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudfunctionsFunctionSpecResource) DeepCopyInto(out *CloudfunctionsFunctionSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AvailableMemoryMb != nil {
		in, out := &in.AvailableMemoryMb, &out.AvailableMemoryMb
		*out = new(int64)
		**out = **in
	}
	if in.BuildEnvironmentVariables != nil {
		in, out := &in.BuildEnvironmentVariables, &out.BuildEnvironmentVariables
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EntryPoint != nil {
		in, out := &in.EntryPoint, &out.EntryPoint
		*out = new(string)
		**out = **in
	}
	if in.EnvironmentVariables != nil {
		in, out := &in.EnvironmentVariables, &out.EnvironmentVariables
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EventTrigger != nil {
		in, out := &in.EventTrigger, &out.EventTrigger
		*out = new(CloudfunctionsFunctionSpecEventTrigger)
		(*in).DeepCopyInto(*out)
	}
	if in.HttpsTriggerURL != nil {
		in, out := &in.HttpsTriggerURL, &out.HttpsTriggerURL
		*out = new(string)
		**out = **in
	}
	if in.IngressSettings != nil {
		in, out := &in.IngressSettings, &out.IngressSettings
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MaxInstances != nil {
		in, out := &in.MaxInstances, &out.MaxInstances
		*out = new(int64)
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
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Runtime != nil {
		in, out := &in.Runtime, &out.Runtime
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountEmail != nil {
		in, out := &in.ServiceAccountEmail, &out.ServiceAccountEmail
		*out = new(string)
		**out = **in
	}
	if in.SourceArchiveBucket != nil {
		in, out := &in.SourceArchiveBucket, &out.SourceArchiveBucket
		*out = new(string)
		**out = **in
	}
	if in.SourceArchiveObject != nil {
		in, out := &in.SourceArchiveObject, &out.SourceArchiveObject
		*out = new(string)
		**out = **in
	}
	if in.SourceRepository != nil {
		in, out := &in.SourceRepository, &out.SourceRepository
		*out = new(CloudfunctionsFunctionSpecSourceRepository)
		(*in).DeepCopyInto(*out)
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int64)
		**out = **in
	}
	if in.TriggerHTTP != nil {
		in, out := &in.TriggerHTTP, &out.TriggerHTTP
		*out = new(bool)
		**out = **in
	}
	if in.VpcConnector != nil {
		in, out := &in.VpcConnector, &out.VpcConnector
		*out = new(string)
		**out = **in
	}
	if in.VpcConnectorEgressSettings != nil {
		in, out := &in.VpcConnectorEgressSettings, &out.VpcConnectorEgressSettings
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudfunctionsFunctionSpecResource.
func (in *CloudfunctionsFunctionSpecResource) DeepCopy() *CloudfunctionsFunctionSpecResource {
	if in == nil {
		return nil
	}
	out := new(CloudfunctionsFunctionSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudfunctionsFunctionSpecSourceRepository) DeepCopyInto(out *CloudfunctionsFunctionSpecSourceRepository) {
	*out = *in
	if in.DeployedURL != nil {
		in, out := &in.DeployedURL, &out.DeployedURL
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudfunctionsFunctionSpecSourceRepository.
func (in *CloudfunctionsFunctionSpecSourceRepository) DeepCopy() *CloudfunctionsFunctionSpecSourceRepository {
	if in == nil {
		return nil
	}
	out := new(CloudfunctionsFunctionSpecSourceRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudfunctionsFunctionStatus) DeepCopyInto(out *CloudfunctionsFunctionStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudfunctionsFunctionStatus.
func (in *CloudfunctionsFunctionStatus) DeepCopy() *CloudfunctionsFunctionStatus {
	if in == nil {
		return nil
	}
	out := new(CloudfunctionsFunctionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamBinding) DeepCopyInto(out *IamBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamBinding.
func (in *IamBinding) DeepCopy() *IamBinding {
	if in == nil {
		return nil
	}
	out := new(IamBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IamBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamBindingList) DeepCopyInto(out *IamBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IamBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamBindingList.
func (in *IamBindingList) DeepCopy() *IamBindingList {
	if in == nil {
		return nil
	}
	out := new(IamBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IamBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamBindingSpec) DeepCopyInto(out *IamBindingSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(IamBindingSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamBindingSpec.
func (in *IamBindingSpec) DeepCopy() *IamBindingSpec {
	if in == nil {
		return nil
	}
	out := new(IamBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamBindingSpecCondition) DeepCopyInto(out *IamBindingSpecCondition) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamBindingSpecCondition.
func (in *IamBindingSpecCondition) DeepCopy() *IamBindingSpecCondition {
	if in == nil {
		return nil
	}
	out := new(IamBindingSpecCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamBindingSpecResource) DeepCopyInto(out *IamBindingSpecResource) {
	*out = *in
	if in.CloudFunction != nil {
		in, out := &in.CloudFunction, &out.CloudFunction
		*out = new(string)
		**out = **in
	}
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(IamBindingSpecCondition)
		(*in).DeepCopyInto(*out)
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]string, len(*in))
		copy(*out, *in)
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
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamBindingSpecResource.
func (in *IamBindingSpecResource) DeepCopy() *IamBindingSpecResource {
	if in == nil {
		return nil
	}
	out := new(IamBindingSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamBindingStatus) DeepCopyInto(out *IamBindingStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamBindingStatus.
func (in *IamBindingStatus) DeepCopy() *IamBindingStatus {
	if in == nil {
		return nil
	}
	out := new(IamBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamMember) DeepCopyInto(out *IamMember) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamMember.
func (in *IamMember) DeepCopy() *IamMember {
	if in == nil {
		return nil
	}
	out := new(IamMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IamMember) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamMemberList) DeepCopyInto(out *IamMemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IamMember, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamMemberList.
func (in *IamMemberList) DeepCopy() *IamMemberList {
	if in == nil {
		return nil
	}
	out := new(IamMemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IamMemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamMemberSpec) DeepCopyInto(out *IamMemberSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(IamMemberSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamMemberSpec.
func (in *IamMemberSpec) DeepCopy() *IamMemberSpec {
	if in == nil {
		return nil
	}
	out := new(IamMemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamMemberSpecCondition) DeepCopyInto(out *IamMemberSpecCondition) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamMemberSpecCondition.
func (in *IamMemberSpecCondition) DeepCopy() *IamMemberSpecCondition {
	if in == nil {
		return nil
	}
	out := new(IamMemberSpecCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamMemberSpecResource) DeepCopyInto(out *IamMemberSpecResource) {
	*out = *in
	if in.CloudFunction != nil {
		in, out := &in.CloudFunction, &out.CloudFunction
		*out = new(string)
		**out = **in
	}
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(IamMemberSpecCondition)
		(*in).DeepCopyInto(*out)
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
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
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamMemberSpecResource.
func (in *IamMemberSpecResource) DeepCopy() *IamMemberSpecResource {
	if in == nil {
		return nil
	}
	out := new(IamMemberSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamMemberStatus) DeepCopyInto(out *IamMemberStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamMemberStatus.
func (in *IamMemberStatus) DeepCopy() *IamMemberStatus {
	if in == nil {
		return nil
	}
	out := new(IamMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamPolicy) DeepCopyInto(out *IamPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamPolicy.
func (in *IamPolicy) DeepCopy() *IamPolicy {
	if in == nil {
		return nil
	}
	out := new(IamPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IamPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamPolicyList) DeepCopyInto(out *IamPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IamPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamPolicyList.
func (in *IamPolicyList) DeepCopy() *IamPolicyList {
	if in == nil {
		return nil
	}
	out := new(IamPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IamPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamPolicySpec) DeepCopyInto(out *IamPolicySpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(IamPolicySpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamPolicySpec.
func (in *IamPolicySpec) DeepCopy() *IamPolicySpec {
	if in == nil {
		return nil
	}
	out := new(IamPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamPolicySpecResource) DeepCopyInto(out *IamPolicySpecResource) {
	*out = *in
	if in.CloudFunction != nil {
		in, out := &in.CloudFunction, &out.CloudFunction
		*out = new(string)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.PolicyData != nil {
		in, out := &in.PolicyData, &out.PolicyData
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamPolicySpecResource.
func (in *IamPolicySpecResource) DeepCopy() *IamPolicySpecResource {
	if in == nil {
		return nil
	}
	out := new(IamPolicySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamPolicyStatus) DeepCopyInto(out *IamPolicyStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamPolicyStatus.
func (in *IamPolicyStatus) DeepCopy() *IamPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(IamPolicyStatus)
	in.DeepCopyInto(out)
	return out
}
