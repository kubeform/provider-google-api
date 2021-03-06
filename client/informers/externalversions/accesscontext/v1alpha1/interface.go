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
	// ManagerAccessLevels returns a ManagerAccessLevelInformer.
	ManagerAccessLevels() ManagerAccessLevelInformer
	// ManagerAccessLevelBatches returns a ManagerAccessLevelBatchInformer.
	ManagerAccessLevelBatches() ManagerAccessLevelBatchInformer
	// ManagerAccessLevelConditions returns a ManagerAccessLevelConditionInformer.
	ManagerAccessLevelConditions() ManagerAccessLevelConditionInformer
	// ManagerAccessPolicies returns a ManagerAccessPolicyInformer.
	ManagerAccessPolicies() ManagerAccessPolicyInformer
	// ManagerGcpUserAccessBindings returns a ManagerGcpUserAccessBindingInformer.
	ManagerGcpUserAccessBindings() ManagerGcpUserAccessBindingInformer
	// ManagerServicePerimeters returns a ManagerServicePerimeterInformer.
	ManagerServicePerimeters() ManagerServicePerimeterInformer
	// ManagerServicePerimeterBatches returns a ManagerServicePerimeterBatchInformer.
	ManagerServicePerimeterBatches() ManagerServicePerimeterBatchInformer
	// ManagerServicePerimeterResources returns a ManagerServicePerimeterResourceInformer.
	ManagerServicePerimeterResources() ManagerServicePerimeterResourceInformer
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

// ManagerAccessLevels returns a ManagerAccessLevelInformer.
func (v *version) ManagerAccessLevels() ManagerAccessLevelInformer {
	return &managerAccessLevelInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerAccessLevelBatches returns a ManagerAccessLevelBatchInformer.
func (v *version) ManagerAccessLevelBatches() ManagerAccessLevelBatchInformer {
	return &managerAccessLevelBatchInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerAccessLevelConditions returns a ManagerAccessLevelConditionInformer.
func (v *version) ManagerAccessLevelConditions() ManagerAccessLevelConditionInformer {
	return &managerAccessLevelConditionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerAccessPolicies returns a ManagerAccessPolicyInformer.
func (v *version) ManagerAccessPolicies() ManagerAccessPolicyInformer {
	return &managerAccessPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerGcpUserAccessBindings returns a ManagerGcpUserAccessBindingInformer.
func (v *version) ManagerGcpUserAccessBindings() ManagerGcpUserAccessBindingInformer {
	return &managerGcpUserAccessBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerServicePerimeters returns a ManagerServicePerimeterInformer.
func (v *version) ManagerServicePerimeters() ManagerServicePerimeterInformer {
	return &managerServicePerimeterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerServicePerimeterBatches returns a ManagerServicePerimeterBatchInformer.
func (v *version) ManagerServicePerimeterBatches() ManagerServicePerimeterBatchInformer {
	return &managerServicePerimeterBatchInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerServicePerimeterResources returns a ManagerServicePerimeterResourceInformer.
func (v *version) ManagerServicePerimeterResources() ManagerServicePerimeterResourceInformer {
	return &managerServicePerimeterResourceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
