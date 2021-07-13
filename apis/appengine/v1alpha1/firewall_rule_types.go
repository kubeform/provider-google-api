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

type FirewallRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallRuleSpec   `json:"spec,omitempty"`
	Status            FirewallRuleStatus `json:"status,omitempty"`
}

type FirewallRuleSpec struct {
	State *FirewallRuleSpecResource `json:"state,omitempty" tf:"-"`

	Resource FirewallRuleSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type FirewallRuleSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The action to take if this rule matches. Possible values: ["UNSPECIFIED_ACTION", "ALLOW", "DENY"]
	Action *string `json:"action" tf:"action"`
	// An optional string description of this rule.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// A positive integer that defines the order of rule evaluation.
	// Rules with the lowest priority are evaluated first.
	//
	// A default rule at priority Int32.MaxValue matches all IPv4 and
	// IPv6 traffic when no previous rule matches. Only the action of
	// this rule can be modified by the user.
	// +optional
	Priority *int64 `json:"priority,omitempty" tf:"priority"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// IP address or range, defined using CIDR notation, of requests that this rule applies to.
	SourceRange *string `json:"sourceRange" tf:"source_range"`
}

type FirewallRuleStatus struct {
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

// FirewallRuleList is a list of FirewallRules
type FirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FirewallRule CRD objects
	Items []FirewallRule `json:"items,omitempty"`
}
