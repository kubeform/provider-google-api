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
	// TagBindings returns a TagBindingInformer.
	TagBindings() TagBindingInformer
	// TagKeys returns a TagKeyInformer.
	TagKeys() TagKeyInformer
	// TagKeyIamBindings returns a TagKeyIamBindingInformer.
	TagKeyIamBindings() TagKeyIamBindingInformer
	// TagKeyIamMembers returns a TagKeyIamMemberInformer.
	TagKeyIamMembers() TagKeyIamMemberInformer
	// TagKeyIamPolicies returns a TagKeyIamPolicyInformer.
	TagKeyIamPolicies() TagKeyIamPolicyInformer
	// TagValues returns a TagValueInformer.
	TagValues() TagValueInformer
	// TagValueIamBindings returns a TagValueIamBindingInformer.
	TagValueIamBindings() TagValueIamBindingInformer
	// TagValueIamMembers returns a TagValueIamMemberInformer.
	TagValueIamMembers() TagValueIamMemberInformer
	// TagValueIamPolicies returns a TagValueIamPolicyInformer.
	TagValueIamPolicies() TagValueIamPolicyInformer
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

// TagBindings returns a TagBindingInformer.
func (v *version) TagBindings() TagBindingInformer {
	return &tagBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TagKeys returns a TagKeyInformer.
func (v *version) TagKeys() TagKeyInformer {
	return &tagKeyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TagKeyIamBindings returns a TagKeyIamBindingInformer.
func (v *version) TagKeyIamBindings() TagKeyIamBindingInformer {
	return &tagKeyIamBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TagKeyIamMembers returns a TagKeyIamMemberInformer.
func (v *version) TagKeyIamMembers() TagKeyIamMemberInformer {
	return &tagKeyIamMemberInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TagKeyIamPolicies returns a TagKeyIamPolicyInformer.
func (v *version) TagKeyIamPolicies() TagKeyIamPolicyInformer {
	return &tagKeyIamPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TagValues returns a TagValueInformer.
func (v *version) TagValues() TagValueInformer {
	return &tagValueInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TagValueIamBindings returns a TagValueIamBindingInformer.
func (v *version) TagValueIamBindings() TagValueIamBindingInformer {
	return &tagValueIamBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TagValueIamMembers returns a TagValueIamMemberInformer.
func (v *version) TagValueIamMembers() TagValueIamMemberInformer {
	return &tagValueIamMemberInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TagValueIamPolicies returns a TagValueIamPolicyInformer.
func (v *version) TagValueIamPolicies() TagValueIamPolicyInformer {
	return &tagValueIamPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
