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
	// Environments returns a EnvironmentInformer.
	Environments() EnvironmentInformer
	// Instances returns a InstanceInformer.
	Instances() InstanceInformer
	// InstanceIamBindings returns a InstanceIamBindingInformer.
	InstanceIamBindings() InstanceIamBindingInformer
	// InstanceIamMembers returns a InstanceIamMemberInformer.
	InstanceIamMembers() InstanceIamMemberInformer
	// InstanceIamPolicies returns a InstanceIamPolicyInformer.
	InstanceIamPolicies() InstanceIamPolicyInformer
	// Locations returns a LocationInformer.
	Locations() LocationInformer
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

// Environments returns a EnvironmentInformer.
func (v *version) Environments() EnvironmentInformer {
	return &environmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Instances returns a InstanceInformer.
func (v *version) Instances() InstanceInformer {
	return &instanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InstanceIamBindings returns a InstanceIamBindingInformer.
func (v *version) InstanceIamBindings() InstanceIamBindingInformer {
	return &instanceIamBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InstanceIamMembers returns a InstanceIamMemberInformer.
func (v *version) InstanceIamMembers() InstanceIamMemberInformer {
	return &instanceIamMemberInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InstanceIamPolicies returns a InstanceIamPolicyInformer.
func (v *version) InstanceIamPolicies() InstanceIamPolicyInformer {
	return &instanceIamPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Locations returns a LocationInformer.
func (v *version) Locations() LocationInformer {
	return &locationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
