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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "kubeform.dev/provider-google-api/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// EndpointAttachments returns a EndpointAttachmentInformer.
	EndpointAttachments() EndpointAttachmentInformer
	// Envgroups returns a EnvgroupInformer.
	Envgroups() EnvgroupInformer
	// EnvgroupAttachments returns a EnvgroupAttachmentInformer.
	EnvgroupAttachments() EnvgroupAttachmentInformer
	// Environments returns a EnvironmentInformer.
	Environments() EnvironmentInformer
	// EnvironmentIamBindings returns a EnvironmentIamBindingInformer.
	EnvironmentIamBindings() EnvironmentIamBindingInformer
	// EnvironmentIamMembers returns a EnvironmentIamMemberInformer.
	EnvironmentIamMembers() EnvironmentIamMemberInformer
	// EnvironmentIamPolicies returns a EnvironmentIamPolicyInformer.
	EnvironmentIamPolicies() EnvironmentIamPolicyInformer
	// Instances returns a InstanceInformer.
	Instances() InstanceInformer
	// InstanceAttachments returns a InstanceAttachmentInformer.
	InstanceAttachments() InstanceAttachmentInformer
	// Organizations returns a OrganizationInformer.
	Organizations() OrganizationInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// EndpointAttachments returns a EndpointAttachmentInformer.
func (v *version) EndpointAttachments() EndpointAttachmentInformer {
	return &endpointAttachmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Envgroups returns a EnvgroupInformer.
func (v *version) Envgroups() EnvgroupInformer {
	return &envgroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EnvgroupAttachments returns a EnvgroupAttachmentInformer.
func (v *version) EnvgroupAttachments() EnvgroupAttachmentInformer {
	return &envgroupAttachmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Environments returns a EnvironmentInformer.
func (v *version) Environments() EnvironmentInformer {
	return &environmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EnvironmentIamBindings returns a EnvironmentIamBindingInformer.
func (v *version) EnvironmentIamBindings() EnvironmentIamBindingInformer {
	return &environmentIamBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EnvironmentIamMembers returns a EnvironmentIamMemberInformer.
func (v *version) EnvironmentIamMembers() EnvironmentIamMemberInformer {
	return &environmentIamMemberInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EnvironmentIamPolicies returns a EnvironmentIamPolicyInformer.
func (v *version) EnvironmentIamPolicies() EnvironmentIamPolicyInformer {
	return &environmentIamPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Instances returns a InstanceInformer.
func (v *version) Instances() InstanceInformer {
	return &instanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InstanceAttachments returns a InstanceAttachmentInformer.
func (v *version) InstanceAttachments() InstanceAttachmentInformer {
	return &instanceAttachmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Organizations returns a OrganizationInformer.
func (v *version) Organizations() OrganizationInformer {
	return &organizationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
