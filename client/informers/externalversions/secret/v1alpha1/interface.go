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
	// ManagerSecrets returns a ManagerSecretInformer.
	ManagerSecrets() ManagerSecretInformer
	// ManagerSecretIamBindings returns a ManagerSecretIamBindingInformer.
	ManagerSecretIamBindings() ManagerSecretIamBindingInformer
	// ManagerSecretIamMembers returns a ManagerSecretIamMemberInformer.
	ManagerSecretIamMembers() ManagerSecretIamMemberInformer
	// ManagerSecretIamPolicies returns a ManagerSecretIamPolicyInformer.
	ManagerSecretIamPolicies() ManagerSecretIamPolicyInformer
	// ManagerSecretVersions returns a ManagerSecretVersionInformer.
	ManagerSecretVersions() ManagerSecretVersionInformer
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

// ManagerSecrets returns a ManagerSecretInformer.
func (v *version) ManagerSecrets() ManagerSecretInformer {
	return &managerSecretInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerSecretIamBindings returns a ManagerSecretIamBindingInformer.
func (v *version) ManagerSecretIamBindings() ManagerSecretIamBindingInformer {
	return &managerSecretIamBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerSecretIamMembers returns a ManagerSecretIamMemberInformer.
func (v *version) ManagerSecretIamMembers() ManagerSecretIamMemberInformer {
	return &managerSecretIamMemberInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerSecretIamPolicies returns a ManagerSecretIamPolicyInformer.
func (v *version) ManagerSecretIamPolicies() ManagerSecretIamPolicyInformer {
	return &managerSecretIamPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerSecretVersions returns a ManagerSecretVersionInformer.
func (v *version) ManagerSecretVersions() ManagerSecretVersionInformer {
	return &managerSecretVersionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
