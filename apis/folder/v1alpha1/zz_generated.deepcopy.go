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
func (in *AccessApprovalSettings) DeepCopyInto(out *AccessApprovalSettings) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessApprovalSettings.
func (in *AccessApprovalSettings) DeepCopy() *AccessApprovalSettings {
	if in == nil {
		return nil
	}
	out := new(AccessApprovalSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessApprovalSettings) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessApprovalSettingsList) DeepCopyInto(out *AccessApprovalSettingsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccessApprovalSettings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessApprovalSettingsList.
func (in *AccessApprovalSettingsList) DeepCopy() *AccessApprovalSettingsList {
	if in == nil {
		return nil
	}
	out := new(AccessApprovalSettingsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessApprovalSettingsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessApprovalSettingsSpec) DeepCopyInto(out *AccessApprovalSettingsSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(AccessApprovalSettingsSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessApprovalSettingsSpec.
func (in *AccessApprovalSettingsSpec) DeepCopy() *AccessApprovalSettingsSpec {
	if in == nil {
		return nil
	}
	out := new(AccessApprovalSettingsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessApprovalSettingsSpecEnrolledServices) DeepCopyInto(out *AccessApprovalSettingsSpecEnrolledServices) {
	*out = *in
	if in.CloudProduct != nil {
		in, out := &in.CloudProduct, &out.CloudProduct
		*out = new(string)
		**out = **in
	}
	if in.EnrollmentLevel != nil {
		in, out := &in.EnrollmentLevel, &out.EnrollmentLevel
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessApprovalSettingsSpecEnrolledServices.
func (in *AccessApprovalSettingsSpecEnrolledServices) DeepCopy() *AccessApprovalSettingsSpecEnrolledServices {
	if in == nil {
		return nil
	}
	out := new(AccessApprovalSettingsSpecEnrolledServices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessApprovalSettingsSpecResource) DeepCopyInto(out *AccessApprovalSettingsSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.EnrolledAncestor != nil {
		in, out := &in.EnrolledAncestor, &out.EnrolledAncestor
		*out = new(bool)
		**out = **in
	}
	if in.EnrolledServices != nil {
		in, out := &in.EnrolledServices, &out.EnrolledServices
		*out = make([]AccessApprovalSettingsSpecEnrolledServices, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FolderID != nil {
		in, out := &in.FolderID, &out.FolderID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NotificationEmails != nil {
		in, out := &in.NotificationEmails, &out.NotificationEmails
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessApprovalSettingsSpecResource.
func (in *AccessApprovalSettingsSpecResource) DeepCopy() *AccessApprovalSettingsSpecResource {
	if in == nil {
		return nil
	}
	out := new(AccessApprovalSettingsSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessApprovalSettingsStatus) DeepCopyInto(out *AccessApprovalSettingsStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessApprovalSettingsStatus.
func (in *AccessApprovalSettingsStatus) DeepCopy() *AccessApprovalSettingsStatus {
	if in == nil {
		return nil
	}
	out := new(AccessApprovalSettingsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Folder) DeepCopyInto(out *Folder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Folder.
func (in *Folder) DeepCopy() *Folder {
	if in == nil {
		return nil
	}
	out := new(Folder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Folder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FolderList) DeepCopyInto(out *FolderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Folder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FolderList.
func (in *FolderList) DeepCopy() *FolderList {
	if in == nil {
		return nil
	}
	out := new(FolderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FolderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FolderSpec) DeepCopyInto(out *FolderSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(FolderSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FolderSpec.
func (in *FolderSpec) DeepCopy() *FolderSpec {
	if in == nil {
		return nil
	}
	out := new(FolderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FolderSpecResource) DeepCopyInto(out *FolderSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.FolderID != nil {
		in, out := &in.FolderID, &out.FolderID
		*out = new(string)
		**out = **in
	}
	if in.LifecycleState != nil {
		in, out := &in.LifecycleState, &out.LifecycleState
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Parent != nil {
		in, out := &in.Parent, &out.Parent
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FolderSpecResource.
func (in *FolderSpecResource) DeepCopy() *FolderSpecResource {
	if in == nil {
		return nil
	}
	out := new(FolderSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FolderStatus) DeepCopyInto(out *FolderStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FolderStatus.
func (in *FolderStatus) DeepCopy() *FolderStatus {
	if in == nil {
		return nil
	}
	out := new(FolderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamAuditConfig) DeepCopyInto(out *IamAuditConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamAuditConfig.
func (in *IamAuditConfig) DeepCopy() *IamAuditConfig {
	if in == nil {
		return nil
	}
	out := new(IamAuditConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IamAuditConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamAuditConfigList) DeepCopyInto(out *IamAuditConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IamAuditConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamAuditConfigList.
func (in *IamAuditConfigList) DeepCopy() *IamAuditConfigList {
	if in == nil {
		return nil
	}
	out := new(IamAuditConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IamAuditConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamAuditConfigSpec) DeepCopyInto(out *IamAuditConfigSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(IamAuditConfigSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamAuditConfigSpec.
func (in *IamAuditConfigSpec) DeepCopy() *IamAuditConfigSpec {
	if in == nil {
		return nil
	}
	out := new(IamAuditConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamAuditConfigSpecAuditLogConfig) DeepCopyInto(out *IamAuditConfigSpecAuditLogConfig) {
	*out = *in
	if in.ExemptedMembers != nil {
		in, out := &in.ExemptedMembers, &out.ExemptedMembers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LogType != nil {
		in, out := &in.LogType, &out.LogType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamAuditConfigSpecAuditLogConfig.
func (in *IamAuditConfigSpecAuditLogConfig) DeepCopy() *IamAuditConfigSpecAuditLogConfig {
	if in == nil {
		return nil
	}
	out := new(IamAuditConfigSpecAuditLogConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamAuditConfigSpecResource) DeepCopyInto(out *IamAuditConfigSpecResource) {
	*out = *in
	if in.AuditLogConfig != nil {
		in, out := &in.AuditLogConfig, &out.AuditLogConfig
		*out = make([]IamAuditConfigSpecAuditLogConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Folder != nil {
		in, out := &in.Folder, &out.Folder
		*out = new(string)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamAuditConfigSpecResource.
func (in *IamAuditConfigSpecResource) DeepCopy() *IamAuditConfigSpecResource {
	if in == nil {
		return nil
	}
	out := new(IamAuditConfigSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamAuditConfigStatus) DeepCopyInto(out *IamAuditConfigStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamAuditConfigStatus.
func (in *IamAuditConfigStatus) DeepCopy() *IamAuditConfigStatus {
	if in == nil {
		return nil
	}
	out := new(IamAuditConfigStatus)
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
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
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
	if in.Folder != nil {
		in, out := &in.Folder, &out.Folder
		*out = new(string)
		**out = **in
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]string, len(*in))
		copy(*out, *in)
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
		*out = make([]apiv1.Condition, len(*in))
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
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
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
	if in.Folder != nil {
		in, out := &in.Folder, &out.Folder
		*out = new(string)
		**out = **in
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
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
		*out = make([]apiv1.Condition, len(*in))
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
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
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
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Folder != nil {
		in, out := &in.Folder, &out.Folder
		*out = new(string)
		**out = **in
	}
	if in.PolicyData != nil {
		in, out := &in.PolicyData, &out.PolicyData
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
		*out = make([]apiv1.Condition, len(*in))
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationPolicy) DeepCopyInto(out *OrganizationPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationPolicy.
func (in *OrganizationPolicy) DeepCopy() *OrganizationPolicy {
	if in == nil {
		return nil
	}
	out := new(OrganizationPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrganizationPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationPolicyList) DeepCopyInto(out *OrganizationPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OrganizationPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationPolicyList.
func (in *OrganizationPolicyList) DeepCopy() *OrganizationPolicyList {
	if in == nil {
		return nil
	}
	out := new(OrganizationPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrganizationPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationPolicySpec) DeepCopyInto(out *OrganizationPolicySpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(OrganizationPolicySpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationPolicySpec.
func (in *OrganizationPolicySpec) DeepCopy() *OrganizationPolicySpec {
	if in == nil {
		return nil
	}
	out := new(OrganizationPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationPolicySpecBooleanPolicy) DeepCopyInto(out *OrganizationPolicySpecBooleanPolicy) {
	*out = *in
	if in.Enforced != nil {
		in, out := &in.Enforced, &out.Enforced
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationPolicySpecBooleanPolicy.
func (in *OrganizationPolicySpecBooleanPolicy) DeepCopy() *OrganizationPolicySpecBooleanPolicy {
	if in == nil {
		return nil
	}
	out := new(OrganizationPolicySpecBooleanPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationPolicySpecListPolicy) DeepCopyInto(out *OrganizationPolicySpecListPolicy) {
	*out = *in
	if in.Allow != nil {
		in, out := &in.Allow, &out.Allow
		*out = new(OrganizationPolicySpecListPolicyAllow)
		(*in).DeepCopyInto(*out)
	}
	if in.Deny != nil {
		in, out := &in.Deny, &out.Deny
		*out = new(OrganizationPolicySpecListPolicyDeny)
		(*in).DeepCopyInto(*out)
	}
	if in.InheritFromParent != nil {
		in, out := &in.InheritFromParent, &out.InheritFromParent
		*out = new(bool)
		**out = **in
	}
	if in.SuggestedValue != nil {
		in, out := &in.SuggestedValue, &out.SuggestedValue
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationPolicySpecListPolicy.
func (in *OrganizationPolicySpecListPolicy) DeepCopy() *OrganizationPolicySpecListPolicy {
	if in == nil {
		return nil
	}
	out := new(OrganizationPolicySpecListPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationPolicySpecListPolicyAllow) DeepCopyInto(out *OrganizationPolicySpecListPolicyAllow) {
	*out = *in
	if in.All != nil {
		in, out := &in.All, &out.All
		*out = new(bool)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationPolicySpecListPolicyAllow.
func (in *OrganizationPolicySpecListPolicyAllow) DeepCopy() *OrganizationPolicySpecListPolicyAllow {
	if in == nil {
		return nil
	}
	out := new(OrganizationPolicySpecListPolicyAllow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationPolicySpecListPolicyDeny) DeepCopyInto(out *OrganizationPolicySpecListPolicyDeny) {
	*out = *in
	if in.All != nil {
		in, out := &in.All, &out.All
		*out = new(bool)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationPolicySpecListPolicyDeny.
func (in *OrganizationPolicySpecListPolicyDeny) DeepCopy() *OrganizationPolicySpecListPolicyDeny {
	if in == nil {
		return nil
	}
	out := new(OrganizationPolicySpecListPolicyDeny)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationPolicySpecResource) DeepCopyInto(out *OrganizationPolicySpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.BooleanPolicy != nil {
		in, out := &in.BooleanPolicy, &out.BooleanPolicy
		*out = new(OrganizationPolicySpecBooleanPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Constraint != nil {
		in, out := &in.Constraint, &out.Constraint
		*out = new(string)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Folder != nil {
		in, out := &in.Folder, &out.Folder
		*out = new(string)
		**out = **in
	}
	if in.ListPolicy != nil {
		in, out := &in.ListPolicy, &out.ListPolicy
		*out = new(OrganizationPolicySpecListPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.RestorePolicy != nil {
		in, out := &in.RestorePolicy, &out.RestorePolicy
		*out = new(OrganizationPolicySpecRestorePolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationPolicySpecResource.
func (in *OrganizationPolicySpecResource) DeepCopy() *OrganizationPolicySpecResource {
	if in == nil {
		return nil
	}
	out := new(OrganizationPolicySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationPolicySpecRestorePolicy) DeepCopyInto(out *OrganizationPolicySpecRestorePolicy) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationPolicySpecRestorePolicy.
func (in *OrganizationPolicySpecRestorePolicy) DeepCopy() *OrganizationPolicySpecRestorePolicy {
	if in == nil {
		return nil
	}
	out := new(OrganizationPolicySpecRestorePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationPolicyStatus) DeepCopyInto(out *OrganizationPolicyStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationPolicyStatus.
func (in *OrganizationPolicyStatus) DeepCopy() *OrganizationPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(OrganizationPolicyStatus)
	in.DeepCopyInto(out)
	return out
}
