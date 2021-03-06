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
func (in *Trigger) DeepCopyInto(out *Trigger) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trigger.
func (in *Trigger) DeepCopy() *Trigger {
	if in == nil {
		return nil
	}
	out := new(Trigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Trigger) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerList) DeepCopyInto(out *TriggerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Trigger, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerList.
func (in *TriggerList) DeepCopy() *TriggerList {
	if in == nil {
		return nil
	}
	out := new(TriggerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TriggerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpec) DeepCopyInto(out *TriggerSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(TriggerSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpec.
func (in *TriggerSpec) DeepCopy() *TriggerSpec {
	if in == nil {
		return nil
	}
	out := new(TriggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpecBuild) DeepCopyInto(out *TriggerSpecBuild) {
	*out = *in
	if in.Artifacts != nil {
		in, out := &in.Artifacts, &out.Artifacts
		*out = new(TriggerSpecBuildArtifacts)
		(*in).DeepCopyInto(*out)
	}
	if in.AvailableSecrets != nil {
		in, out := &in.AvailableSecrets, &out.AvailableSecrets
		*out = new(TriggerSpecBuildAvailableSecrets)
		(*in).DeepCopyInto(*out)
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LogsBucket != nil {
		in, out := &in.LogsBucket, &out.LogsBucket
		*out = new(string)
		**out = **in
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = new(TriggerSpecBuildOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.QueueTtl != nil {
		in, out := &in.QueueTtl, &out.QueueTtl
		*out = new(string)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = make([]TriggerSpecBuildSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(TriggerSpecBuildSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Step != nil {
		in, out := &in.Step, &out.Step
		*out = make([]TriggerSpecBuildStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Substitutions != nil {
		in, out := &in.Substitutions, &out.Substitutions
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpecBuild.
func (in *TriggerSpecBuild) DeepCopy() *TriggerSpecBuild {
	if in == nil {
		return nil
	}
	out := new(TriggerSpecBuild)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpecBuildArtifacts) DeepCopyInto(out *TriggerSpecBuildArtifacts) {
	*out = *in
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = new(TriggerSpecBuildArtifactsObjects)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpecBuildArtifacts.
func (in *TriggerSpecBuildArtifacts) DeepCopy() *TriggerSpecBuildArtifacts {
	if in == nil {
		return nil
	}
	out := new(TriggerSpecBuildArtifacts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpecBuildArtifactsObjects) DeepCopyInto(out *TriggerSpecBuildArtifactsObjects) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Timing != nil {
		in, out := &in.Timing, &out.Timing
		*out = make([]TriggerSpecBuildArtifactsObjectsTiming, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpecBuildArtifactsObjects.
func (in *TriggerSpecBuildArtifactsObjects) DeepCopy() *TriggerSpecBuildArtifactsObjects {
	if in == nil {
		return nil
	}
	out := new(TriggerSpecBuildArtifactsObjects)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpecBuildArtifactsObjectsTiming) DeepCopyInto(out *TriggerSpecBuildArtifactsObjectsTiming) {
	*out = *in
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpecBuildArtifactsObjectsTiming.
func (in *TriggerSpecBuildArtifactsObjectsTiming) DeepCopy() *TriggerSpecBuildArtifactsObjectsTiming {
	if in == nil {
		return nil
	}
	out := new(TriggerSpecBuildArtifactsObjectsTiming)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpecBuildAvailableSecrets) DeepCopyInto(out *TriggerSpecBuildAvailableSecrets) {
	*out = *in
	if in.SecretManager != nil {
		in, out := &in.SecretManager, &out.SecretManager
		*out = make([]TriggerSpecBuildAvailableSecretsSecretManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpecBuildAvailableSecrets.
func (in *TriggerSpecBuildAvailableSecrets) DeepCopy() *TriggerSpecBuildAvailableSecrets {
	if in == nil {
		return nil
	}
	out := new(TriggerSpecBuildAvailableSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpecBuildAvailableSecretsSecretManager) DeepCopyInto(out *TriggerSpecBuildAvailableSecretsSecretManager) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = new(string)
		**out = **in
	}
	if in.VersionName != nil {
		in, out := &in.VersionName, &out.VersionName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpecBuildAvailableSecretsSecretManager.
func (in *TriggerSpecBuildAvailableSecretsSecretManager) DeepCopy() *TriggerSpecBuildAvailableSecretsSecretManager {
	if in == nil {
		return nil
	}
	out := new(TriggerSpecBuildAvailableSecretsSecretManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpecBuildOptions) DeepCopyInto(out *TriggerSpecBuildOptions) {
	*out = *in
	if in.DiskSizeGb != nil {
		in, out := &in.DiskSizeGb, &out.DiskSizeGb
		*out = new(int64)
		**out = **in
	}
	if in.DynamicSubstitutions != nil {
		in, out := &in.DynamicSubstitutions, &out.DynamicSubstitutions
		*out = new(bool)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LogStreamingOption != nil {
		in, out := &in.LogStreamingOption, &out.LogStreamingOption
		*out = new(string)
		**out = **in
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(string)
		**out = **in
	}
	if in.MachineType != nil {
		in, out := &in.MachineType, &out.MachineType
		*out = new(string)
		**out = **in
	}
	if in.RequestedVerifyOption != nil {
		in, out := &in.RequestedVerifyOption, &out.RequestedVerifyOption
		*out = new(string)
		**out = **in
	}
	if in.SecretEnv != nil {
		in, out := &in.SecretEnv, &out.SecretEnv
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SourceProvenanceHash != nil {
		in, out := &in.SourceProvenanceHash, &out.SourceProvenanceHash
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubstitutionOption != nil {
		in, out := &in.SubstitutionOption, &out.SubstitutionOption
		*out = new(string)
		**out = **in
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]TriggerSpecBuildOptionsVolumes, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WorkerPool != nil {
		in, out := &in.WorkerPool, &out.WorkerPool
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpecBuildOptions.
func (in *TriggerSpecBuildOptions) DeepCopy() *TriggerSpecBuildOptions {
	if in == nil {
		return nil
	}
	out := new(TriggerSpecBuildOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpecBuildOptionsVolumes) DeepCopyInto(out *TriggerSpecBuildOptionsVolumes) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpecBuildOptionsVolumes.
func (in *TriggerSpecBuildOptionsVolumes) DeepCopy() *TriggerSpecBuildOptionsVolumes {
	if in == nil {
		return nil
	}
	out := new(TriggerSpecBuildOptionsVolumes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpecBuildSecret) DeepCopyInto(out *TriggerSpecBuildSecret) {
	*out = *in
	if in.KmsKeyName != nil {
		in, out := &in.KmsKeyName, &out.KmsKeyName
		*out = new(string)
		**out = **in
	}
	if in.SecretEnv != nil {
		in, out := &in.SecretEnv, &out.SecretEnv
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpecBuildSecret.
func (in *TriggerSpecBuildSecret) DeepCopy() *TriggerSpecBuildSecret {
	if in == nil {
		return nil
	}
	out := new(TriggerSpecBuildSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpecBuildSource) DeepCopyInto(out *TriggerSpecBuildSource) {
	*out = *in
	if in.RepoSource != nil {
		in, out := &in.RepoSource, &out.RepoSource
		*out = new(TriggerSpecBuildSourceRepoSource)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageSource != nil {
		in, out := &in.StorageSource, &out.StorageSource
		*out = new(TriggerSpecBuildSourceStorageSource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpecBuildSource.
func (in *TriggerSpecBuildSource) DeepCopy() *TriggerSpecBuildSource {
	if in == nil {
		return nil
	}
	out := new(TriggerSpecBuildSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpecBuildSourceRepoSource) DeepCopyInto(out *TriggerSpecBuildSourceRepoSource) {
	*out = *in
	if in.BranchName != nil {
		in, out := &in.BranchName, &out.BranchName
		*out = new(string)
		**out = **in
	}
	if in.CommitSha != nil {
		in, out := &in.CommitSha, &out.CommitSha
		*out = new(string)
		**out = **in
	}
	if in.Dir != nil {
		in, out := &in.Dir, &out.Dir
		*out = new(string)
		**out = **in
	}
	if in.InvertRegex != nil {
		in, out := &in.InvertRegex, &out.InvertRegex
		*out = new(bool)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.RepoName != nil {
		in, out := &in.RepoName, &out.RepoName
		*out = new(string)
		**out = **in
	}
	if in.Substitutions != nil {
		in, out := &in.Substitutions, &out.Substitutions
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.TagName != nil {
		in, out := &in.TagName, &out.TagName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpecBuildSourceRepoSource.
func (in *TriggerSpecBuildSourceRepoSource) DeepCopy() *TriggerSpecBuildSourceRepoSource {
	if in == nil {
		return nil
	}
	out := new(TriggerSpecBuildSourceRepoSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpecBuildSourceStorageSource) DeepCopyInto(out *TriggerSpecBuildSourceStorageSource) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.Generation != nil {
		in, out := &in.Generation, &out.Generation
		*out = new(string)
		**out = **in
	}
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpecBuildSourceStorageSource.
func (in *TriggerSpecBuildSourceStorageSource) DeepCopy() *TriggerSpecBuildSourceStorageSource {
	if in == nil {
		return nil
	}
	out := new(TriggerSpecBuildSourceStorageSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpecBuildStep) DeepCopyInto(out *TriggerSpecBuildStep) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Dir != nil {
		in, out := &in.Dir, &out.Dir
		*out = new(string)
		**out = **in
	}
	if in.Entrypoint != nil {
		in, out := &in.Entrypoint, &out.Entrypoint
		*out = new(string)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.SecretEnv != nil {
		in, out := &in.SecretEnv, &out.SecretEnv
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(string)
		**out = **in
	}
	if in.Timing != nil {
		in, out := &in.Timing, &out.Timing
		*out = new(string)
		**out = **in
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]TriggerSpecBuildStepVolumes, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WaitFor != nil {
		in, out := &in.WaitFor, &out.WaitFor
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpecBuildStep.
func (in *TriggerSpecBuildStep) DeepCopy() *TriggerSpecBuildStep {
	if in == nil {
		return nil
	}
	out := new(TriggerSpecBuildStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpecBuildStepVolumes) DeepCopyInto(out *TriggerSpecBuildStepVolumes) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpecBuildStepVolumes.
func (in *TriggerSpecBuildStepVolumes) DeepCopy() *TriggerSpecBuildStepVolumes {
	if in == nil {
		return nil
	}
	out := new(TriggerSpecBuildStepVolumes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpecGithub) DeepCopyInto(out *TriggerSpecGithub) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	if in.PullRequest != nil {
		in, out := &in.PullRequest, &out.PullRequest
		*out = new(TriggerSpecGithubPullRequest)
		(*in).DeepCopyInto(*out)
	}
	if in.Push != nil {
		in, out := &in.Push, &out.Push
		*out = new(TriggerSpecGithubPush)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpecGithub.
func (in *TriggerSpecGithub) DeepCopy() *TriggerSpecGithub {
	if in == nil {
		return nil
	}
	out := new(TriggerSpecGithub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpecGithubPullRequest) DeepCopyInto(out *TriggerSpecGithubPullRequest) {
	*out = *in
	if in.Branch != nil {
		in, out := &in.Branch, &out.Branch
		*out = new(string)
		**out = **in
	}
	if in.CommentControl != nil {
		in, out := &in.CommentControl, &out.CommentControl
		*out = new(string)
		**out = **in
	}
	if in.InvertRegex != nil {
		in, out := &in.InvertRegex, &out.InvertRegex
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpecGithubPullRequest.
func (in *TriggerSpecGithubPullRequest) DeepCopy() *TriggerSpecGithubPullRequest {
	if in == nil {
		return nil
	}
	out := new(TriggerSpecGithubPullRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpecGithubPush) DeepCopyInto(out *TriggerSpecGithubPush) {
	*out = *in
	if in.Branch != nil {
		in, out := &in.Branch, &out.Branch
		*out = new(string)
		**out = **in
	}
	if in.InvertRegex != nil {
		in, out := &in.InvertRegex, &out.InvertRegex
		*out = new(bool)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpecGithubPush.
func (in *TriggerSpecGithubPush) DeepCopy() *TriggerSpecGithubPush {
	if in == nil {
		return nil
	}
	out := new(TriggerSpecGithubPush)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpecPubsubConfig) DeepCopyInto(out *TriggerSpecPubsubConfig) {
	*out = *in
	if in.ServiceAccountEmail != nil {
		in, out := &in.ServiceAccountEmail, &out.ServiceAccountEmail
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Subscription != nil {
		in, out := &in.Subscription, &out.Subscription
		*out = new(string)
		**out = **in
	}
	if in.Topic != nil {
		in, out := &in.Topic, &out.Topic
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpecPubsubConfig.
func (in *TriggerSpecPubsubConfig) DeepCopy() *TriggerSpecPubsubConfig {
	if in == nil {
		return nil
	}
	out := new(TriggerSpecPubsubConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpecResource) DeepCopyInto(out *TriggerSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(TriggerSpecBuild)
		(*in).DeepCopyInto(*out)
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
	if in.Filename != nil {
		in, out := &in.Filename, &out.Filename
		*out = new(string)
		**out = **in
	}
	if in.Github != nil {
		in, out := &in.Github, &out.Github
		*out = new(TriggerSpecGithub)
		(*in).DeepCopyInto(*out)
	}
	if in.IgnoredFiles != nil {
		in, out := &in.IgnoredFiles, &out.IgnoredFiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IncludedFiles != nil {
		in, out := &in.IncludedFiles, &out.IncludedFiles
		*out = make([]string, len(*in))
		copy(*out, *in)
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
	if in.PubsubConfig != nil {
		in, out := &in.PubsubConfig, &out.PubsubConfig
		*out = new(TriggerSpecPubsubConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(string)
		**out = **in
	}
	if in.Substitutions != nil {
		in, out := &in.Substitutions, &out.Substitutions
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TriggerID != nil {
		in, out := &in.TriggerID, &out.TriggerID
		*out = new(string)
		**out = **in
	}
	if in.TriggerTemplate != nil {
		in, out := &in.TriggerTemplate, &out.TriggerTemplate
		*out = new(TriggerSpecTriggerTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.WebhookConfig != nil {
		in, out := &in.WebhookConfig, &out.WebhookConfig
		*out = new(TriggerSpecWebhookConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpecResource.
func (in *TriggerSpecResource) DeepCopy() *TriggerSpecResource {
	if in == nil {
		return nil
	}
	out := new(TriggerSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpecTriggerTemplate) DeepCopyInto(out *TriggerSpecTriggerTemplate) {
	*out = *in
	if in.BranchName != nil {
		in, out := &in.BranchName, &out.BranchName
		*out = new(string)
		**out = **in
	}
	if in.CommitSha != nil {
		in, out := &in.CommitSha, &out.CommitSha
		*out = new(string)
		**out = **in
	}
	if in.Dir != nil {
		in, out := &in.Dir, &out.Dir
		*out = new(string)
		**out = **in
	}
	if in.InvertRegex != nil {
		in, out := &in.InvertRegex, &out.InvertRegex
		*out = new(bool)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.RepoName != nil {
		in, out := &in.RepoName, &out.RepoName
		*out = new(string)
		**out = **in
	}
	if in.TagName != nil {
		in, out := &in.TagName, &out.TagName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpecTriggerTemplate.
func (in *TriggerSpecTriggerTemplate) DeepCopy() *TriggerSpecTriggerTemplate {
	if in == nil {
		return nil
	}
	out := new(TriggerSpecTriggerTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpecWebhookConfig) DeepCopyInto(out *TriggerSpecWebhookConfig) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpecWebhookConfig.
func (in *TriggerSpecWebhookConfig) DeepCopy() *TriggerSpecWebhookConfig {
	if in == nil {
		return nil
	}
	out := new(TriggerSpecWebhookConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerStatus) DeepCopyInto(out *TriggerStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerStatus.
func (in *TriggerStatus) DeepCopy() *TriggerStatus {
	if in == nil {
		return nil
	}
	out := new(TriggerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerPool) DeepCopyInto(out *WorkerPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerPool.
func (in *WorkerPool) DeepCopy() *WorkerPool {
	if in == nil {
		return nil
	}
	out := new(WorkerPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerPoolList) DeepCopyInto(out *WorkerPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkerPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerPoolList.
func (in *WorkerPoolList) DeepCopy() *WorkerPoolList {
	if in == nil {
		return nil
	}
	out := new(WorkerPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkerPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerPoolSpec) DeepCopyInto(out *WorkerPoolSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(WorkerPoolSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerPoolSpec.
func (in *WorkerPoolSpec) DeepCopy() *WorkerPoolSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerPoolSpecNetworkConfig) DeepCopyInto(out *WorkerPoolSpecNetworkConfig) {
	*out = *in
	if in.PeeredNetwork != nil {
		in, out := &in.PeeredNetwork, &out.PeeredNetwork
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerPoolSpecNetworkConfig.
func (in *WorkerPoolSpecNetworkConfig) DeepCopy() *WorkerPoolSpecNetworkConfig {
	if in == nil {
		return nil
	}
	out := new(WorkerPoolSpecNetworkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerPoolSpecResource) DeepCopyInto(out *WorkerPoolSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.DeleteTime != nil {
		in, out := &in.DeleteTime, &out.DeleteTime
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
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
	if in.NetworkConfig != nil {
		in, out := &in.NetworkConfig, &out.NetworkConfig
		*out = new(WorkerPoolSpecNetworkConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Uid != nil {
		in, out := &in.Uid, &out.Uid
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	if in.WorkerConfig != nil {
		in, out := &in.WorkerConfig, &out.WorkerConfig
		*out = new(WorkerPoolSpecWorkerConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerPoolSpecResource.
func (in *WorkerPoolSpecResource) DeepCopy() *WorkerPoolSpecResource {
	if in == nil {
		return nil
	}
	out := new(WorkerPoolSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerPoolSpecWorkerConfig) DeepCopyInto(out *WorkerPoolSpecWorkerConfig) {
	*out = *in
	if in.DiskSizeGb != nil {
		in, out := &in.DiskSizeGb, &out.DiskSizeGb
		*out = new(int64)
		**out = **in
	}
	if in.MachineType != nil {
		in, out := &in.MachineType, &out.MachineType
		*out = new(string)
		**out = **in
	}
	if in.NoExternalIP != nil {
		in, out := &in.NoExternalIP, &out.NoExternalIP
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerPoolSpecWorkerConfig.
func (in *WorkerPoolSpecWorkerConfig) DeepCopy() *WorkerPoolSpecWorkerConfig {
	if in == nil {
		return nil
	}
	out := new(WorkerPoolSpecWorkerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerPoolStatus) DeepCopyInto(out *WorkerPoolStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerPoolStatus.
func (in *WorkerPoolStatus) DeepCopy() *WorkerPoolStatus {
	if in == nil {
		return nil
	}
	out := new(WorkerPoolStatus)
	in.DeepCopyInto(out)
	return out
}
