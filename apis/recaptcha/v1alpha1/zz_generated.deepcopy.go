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
func (in *EnterpriseKey) DeepCopyInto(out *EnterpriseKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseKey.
func (in *EnterpriseKey) DeepCopy() *EnterpriseKey {
	if in == nil {
		return nil
	}
	out := new(EnterpriseKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnterpriseKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseKeyList) DeepCopyInto(out *EnterpriseKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EnterpriseKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseKeyList.
func (in *EnterpriseKeyList) DeepCopy() *EnterpriseKeyList {
	if in == nil {
		return nil
	}
	out := new(EnterpriseKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnterpriseKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseKeySpec) DeepCopyInto(out *EnterpriseKeySpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(EnterpriseKeySpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseKeySpec.
func (in *EnterpriseKeySpec) DeepCopy() *EnterpriseKeySpec {
	if in == nil {
		return nil
	}
	out := new(EnterpriseKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseKeySpecAndroidSettings) DeepCopyInto(out *EnterpriseKeySpecAndroidSettings) {
	*out = *in
	if in.AllowAllPackageNames != nil {
		in, out := &in.AllowAllPackageNames, &out.AllowAllPackageNames
		*out = new(bool)
		**out = **in
	}
	if in.AllowedPackageNames != nil {
		in, out := &in.AllowedPackageNames, &out.AllowedPackageNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseKeySpecAndroidSettings.
func (in *EnterpriseKeySpecAndroidSettings) DeepCopy() *EnterpriseKeySpecAndroidSettings {
	if in == nil {
		return nil
	}
	out := new(EnterpriseKeySpecAndroidSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseKeySpecIosSettings) DeepCopyInto(out *EnterpriseKeySpecIosSettings) {
	*out = *in
	if in.AllowAllBundleIDS != nil {
		in, out := &in.AllowAllBundleIDS, &out.AllowAllBundleIDS
		*out = new(bool)
		**out = **in
	}
	if in.AllowedBundleIDS != nil {
		in, out := &in.AllowedBundleIDS, &out.AllowedBundleIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseKeySpecIosSettings.
func (in *EnterpriseKeySpecIosSettings) DeepCopy() *EnterpriseKeySpecIosSettings {
	if in == nil {
		return nil
	}
	out := new(EnterpriseKeySpecIosSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseKeySpecResource) DeepCopyInto(out *EnterpriseKeySpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AndroidSettings != nil {
		in, out := &in.AndroidSettings, &out.AndroidSettings
		*out = new(EnterpriseKeySpecAndroidSettings)
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
	if in.IosSettings != nil {
		in, out := &in.IosSettings, &out.IosSettings
		*out = new(EnterpriseKeySpecIosSettings)
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
	if in.TestingOptions != nil {
		in, out := &in.TestingOptions, &out.TestingOptions
		*out = new(EnterpriseKeySpecTestingOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.WebSettings != nil {
		in, out := &in.WebSettings, &out.WebSettings
		*out = new(EnterpriseKeySpecWebSettings)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseKeySpecResource.
func (in *EnterpriseKeySpecResource) DeepCopy() *EnterpriseKeySpecResource {
	if in == nil {
		return nil
	}
	out := new(EnterpriseKeySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseKeySpecTestingOptions) DeepCopyInto(out *EnterpriseKeySpecTestingOptions) {
	*out = *in
	if in.TestingChallenge != nil {
		in, out := &in.TestingChallenge, &out.TestingChallenge
		*out = new(string)
		**out = **in
	}
	if in.TestingScore != nil {
		in, out := &in.TestingScore, &out.TestingScore
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseKeySpecTestingOptions.
func (in *EnterpriseKeySpecTestingOptions) DeepCopy() *EnterpriseKeySpecTestingOptions {
	if in == nil {
		return nil
	}
	out := new(EnterpriseKeySpecTestingOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseKeySpecWebSettings) DeepCopyInto(out *EnterpriseKeySpecWebSettings) {
	*out = *in
	if in.AllowAllDomains != nil {
		in, out := &in.AllowAllDomains, &out.AllowAllDomains
		*out = new(bool)
		**out = **in
	}
	if in.AllowAmpTraffic != nil {
		in, out := &in.AllowAmpTraffic, &out.AllowAmpTraffic
		*out = new(bool)
		**out = **in
	}
	if in.AllowedDomains != nil {
		in, out := &in.AllowedDomains, &out.AllowedDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ChallengeSecurityPreference != nil {
		in, out := &in.ChallengeSecurityPreference, &out.ChallengeSecurityPreference
		*out = new(string)
		**out = **in
	}
	if in.IntegrationType != nil {
		in, out := &in.IntegrationType, &out.IntegrationType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseKeySpecWebSettings.
func (in *EnterpriseKeySpecWebSettings) DeepCopy() *EnterpriseKeySpecWebSettings {
	if in == nil {
		return nil
	}
	out := new(EnterpriseKeySpecWebSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseKeyStatus) DeepCopyInto(out *EnterpriseKeyStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseKeyStatus.
func (in *EnterpriseKeyStatus) DeepCopy() *EnterpriseKeyStatus {
	if in == nil {
		return nil
	}
	out := new(EnterpriseKeyStatus)
	in.DeepCopyInto(out)
	return out
}