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
	// LiteSubscriptions returns a LiteSubscriptionInformer.
	LiteSubscriptions() LiteSubscriptionInformer
	// LiteTopics returns a LiteTopicInformer.
	LiteTopics() LiteTopicInformer
	// Schemas returns a SchemaInformer.
	Schemas() SchemaInformer
	// Subscriptions returns a SubscriptionInformer.
	Subscriptions() SubscriptionInformer
	// SubscriptionIamBindings returns a SubscriptionIamBindingInformer.
	SubscriptionIamBindings() SubscriptionIamBindingInformer
	// SubscriptionIamMembers returns a SubscriptionIamMemberInformer.
	SubscriptionIamMembers() SubscriptionIamMemberInformer
	// SubscriptionIamPolicies returns a SubscriptionIamPolicyInformer.
	SubscriptionIamPolicies() SubscriptionIamPolicyInformer
	// Topics returns a TopicInformer.
	Topics() TopicInformer
	// TopicIamBindings returns a TopicIamBindingInformer.
	TopicIamBindings() TopicIamBindingInformer
	// TopicIamMembers returns a TopicIamMemberInformer.
	TopicIamMembers() TopicIamMemberInformer
	// TopicIamPolicies returns a TopicIamPolicyInformer.
	TopicIamPolicies() TopicIamPolicyInformer
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

// LiteSubscriptions returns a LiteSubscriptionInformer.
func (v *version) LiteSubscriptions() LiteSubscriptionInformer {
	return &liteSubscriptionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// LiteTopics returns a LiteTopicInformer.
func (v *version) LiteTopics() LiteTopicInformer {
	return &liteTopicInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Schemas returns a SchemaInformer.
func (v *version) Schemas() SchemaInformer {
	return &schemaInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Subscriptions returns a SubscriptionInformer.
func (v *version) Subscriptions() SubscriptionInformer {
	return &subscriptionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SubscriptionIamBindings returns a SubscriptionIamBindingInformer.
func (v *version) SubscriptionIamBindings() SubscriptionIamBindingInformer {
	return &subscriptionIamBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SubscriptionIamMembers returns a SubscriptionIamMemberInformer.
func (v *version) SubscriptionIamMembers() SubscriptionIamMemberInformer {
	return &subscriptionIamMemberInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SubscriptionIamPolicies returns a SubscriptionIamPolicyInformer.
func (v *version) SubscriptionIamPolicies() SubscriptionIamPolicyInformer {
	return &subscriptionIamPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Topics returns a TopicInformer.
func (v *version) Topics() TopicInformer {
	return &topicInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TopicIamBindings returns a TopicIamBindingInformer.
func (v *version) TopicIamBindings() TopicIamBindingInformer {
	return &topicIamBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TopicIamMembers returns a TopicIamMemberInformer.
func (v *version) TopicIamMembers() TopicIamMemberInformer {
	return &topicIamMemberInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TopicIamPolicies returns a TopicIamPolicyInformer.
func (v *version) TopicIamPolicies() TopicIamPolicyInformer {
	return &topicIamPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
