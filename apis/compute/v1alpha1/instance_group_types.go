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

type InstanceGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceGroupSpec   `json:"spec,omitempty"`
	Status            InstanceGroupStatus `json:"status,omitempty"`
}

type InstanceGroupSpecNamedPort struct {
	// The name which the port will be mapped to.
	Name *string `json:"name" tf:"name"`
	// The port number to map the name to.
	Port *int64 `json:"port" tf:"port"`
}

type InstanceGroupSpec struct {
	State *InstanceGroupSpecResource `json:"state,omitempty" tf:"-"`

	Resource InstanceGroupSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type InstanceGroupSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// An optional textual description of the instance group.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// List of instances in the group. They should be given as self_link URLs. When adding instances they must all be in the same network and zone as the instance group.
	// +optional
	Instances []string `json:"instances,omitempty" tf:"instances"`
	// The name of the instance group. Must be 1-63 characters long and comply with RFC1035. Supported characters include lowercase letters, numbers, and hyphens.
	Name *string `json:"name" tf:"name"`
	// The named port configuration.
	// +optional
	NamedPort []InstanceGroupSpecNamedPort `json:"namedPort,omitempty" tf:"named_port"`
	// The URL of the network the instance group is in. If this is different from the network where the instances are in, the creation fails. Defaults to the network where the instances are in (if neither network nor instances is specified, this field will be blank).
	// +optional
	Network *string `json:"network,omitempty" tf:"network"`
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The URI of the created resource.
	// +optional
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link"`
	// The number of instances in the group.
	// +optional
	Size *int64 `json:"size,omitempty" tf:"size"`
	// The zone that this instance group should be created in.
	// +optional
	Zone *string `json:"zone,omitempty" tf:"zone"`
}

type InstanceGroupStatus struct {
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

// InstanceGroupList is a list of InstanceGroups
type InstanceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of InstanceGroup CRD objects
	Items []InstanceGroup `json:"items,omitempty"`
}
