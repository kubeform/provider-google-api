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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type PolicyPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicyPolicySpec   `json:"spec,omitempty"`
	Status            PolicyPolicyStatus `json:"status,omitempty"`
}

type PolicyPolicySpecSpecRulesCondition struct {
	// Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Textual representation of an expression in Common Expression Language syntax.
	// +optional
	Expression *string `json:"expression,omitempty" tf:"expression"`
	// Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	// +optional
	Location *string `json:"location,omitempty" tf:"location"`
	// Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
	// +optional
	Title *string `json:"title,omitempty" tf:"title"`
}

type PolicyPolicySpecSpecRulesValues struct {
	// List of values allowed at this resource.
	// +optional
	AllowedValues []string `json:"allowedValues,omitempty" tf:"allowed_values"`
	// List of values denied at this resource.
	// +optional
	DeniedValues []string `json:"deniedValues,omitempty" tf:"denied_values"`
}

type PolicyPolicySpecSpecRules struct {
	// Setting this to true means that all values are allowed. This field can be set only in Policies for list constraints.
	// +optional
	AllowAll *string `json:"allowAll,omitempty" tf:"allow_all"`
	// A condition which determines whether this rule is used in the evaluation of the policy. When set, the `expression` field in the `Expr' must include from 1 to 10 subexpressions, joined by the "||" or "&&" operators. Each subexpression must be of the form "resource.matchTag('/tag_key_short_name, 'tag_value_short_name')". or "resource.matchTagId('tagKeys/key_id', 'tagValues/value_id')". where key_name and value_name are the resource names for Label Keys and Values. These names are available from the Tag Manager Service. An example expression is: "resource.matchTag('123456789/environment, 'prod')". or "resource.matchTagId('tagKeys/123', 'tagValues/456')".
	// +optional
	Condition *PolicyPolicySpecSpecRulesCondition `json:"condition,omitempty" tf:"condition"`
	// Setting this to true means that all values are denied. This field can be set only in Policies for list constraints.
	// +optional
	DenyAll *string `json:"denyAll,omitempty" tf:"deny_all"`
	// If `true`, then the `Policy` is enforced. If `false`, then any configuration is acceptable. This field can be set only in Policies for boolean constraints.
	// +optional
	Enforce *string `json:"enforce,omitempty" tf:"enforce"`
	// List of values to be used for this PolicyRule. This field can be set only in Policies for list constraints.
	// +optional
	Values *PolicyPolicySpecSpecRulesValues `json:"values,omitempty" tf:"values"`
}

type PolicyPolicySpecSpec struct {
	// An opaque tag indicating the current version of the `Policy`, used for concurrency control. This field is ignored if used in a `CreatePolicy` request. When the `Policy` is returned from either a `GetPolicy` or a `ListPolicies` request, this `etag` indicates the version of the current `Policy` to use when executing a read-modify-write loop. When the `Policy` is returned from a `GetEffectivePolicy` request, the `etag` will be unset.
	// +optional
	Etag *string `json:"etag,omitempty" tf:"etag"`
	// Determines the inheritance behavior for this `Policy`. If `inherit_from_parent` is true, PolicyRules set higher up in the hierarchy (up to the closest root) are inherited and present in the effective policy. If it is false, then no rules are inherited, and this Policy becomes the new root for evaluation. This field can be set only for Policies which configure list constraints.
	// +optional
	InheritFromParent *bool `json:"inheritFromParent,omitempty" tf:"inherit_from_parent"`
	// Ignores policies set above this resource and restores the `constraint_default` enforcement behavior of the specific `Constraint` at this resource. This field can be set in policies for either list or boolean constraints. If set, `rules` must be empty and `inherit_from_parent` must be set to false.
	// +optional
	Reset *bool `json:"reset,omitempty" tf:"reset"`
	// Up to 10 PolicyRules are allowed. In Policies for boolean constraints, the following requirements apply: - There must be one and only one PolicyRule where condition is unset. - BooleanPolicyRules with conditions must set `enforced` to the opposite of the PolicyRule without a condition. - During policy evaluation, PolicyRules with conditions that are true for a target resource take precedence.
	// +optional
	Rules []PolicyPolicySpecSpecRules `json:"rules,omitempty" tf:"rules"`
	// Output only. The time stamp this was previously updated. This represents the last time a call to `CreatePolicy` or `UpdatePolicy` was made for that `Policy`.
	// +optional
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time"`
}

type PolicyPolicySpec struct {
	State *PolicyPolicySpecResource `json:"state,omitempty" tf:"-"`

	Resource PolicyPolicySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type PolicyPolicySpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the constraint which this Policy configures: * `projects/{project_number}/policies/{constraint_name}` * `folders/{folder_id}/policies/{constraint_name}` * `organizations/{organization_id}/policies/{constraint_name}` For example, "projects/123/policies/compute.disableSerialPortAccess". Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
	Name *string `json:"name" tf:"name"`
	// The parent of the resource.
	Parent *string `json:"parent" tf:"parent"`
	// Basic information about the Organization Policy.
	// +optional
	Spec *PolicyPolicySpecSpec `json:"spec,omitempty" tf:"spec"`
}

type PolicyPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PolicyPolicyList is a list of PolicyPolicys
type PolicyPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PolicyPolicy CRD objects
	Items []PolicyPolicy `json:"items,omitempty"`
}