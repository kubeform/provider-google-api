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

type PerInstanceConfig struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PerInstanceConfigSpec   `json:"spec,omitempty"`
	Status            PerInstanceConfigStatus `json:"status,omitempty"`
}

type PerInstanceConfigSpecPreservedStateDisk struct {
	// A value that prescribes what should happen to the stateful disk when the VM instance is deleted.
	// The available options are 'NEVER' and 'ON_PERMANENT_INSTANCE_DELETION'.
	// 'NEVER' - detach the disk when the VM is deleted, but do not delete the disk.
	// 'ON_PERMANENT_INSTANCE_DELETION' will delete the stateful disk when the VM is permanently
	// deleted from the instance group. Default value: "NEVER" Possible values: ["NEVER", "ON_PERMANENT_INSTANCE_DELETION"]
	// +optional
	DeleteRule *string `json:"deleteRule,omitempty" tf:"delete_rule"`
	// A unique device name that is reflected into the /dev/ tree of a Linux operating system running within the instance.
	DeviceName *string `json:"deviceName" tf:"device_name"`
	// The mode of the disk. Default value: "READ_WRITE" Possible values: ["READ_ONLY", "READ_WRITE"]
	// +optional
	Mode *string `json:"mode,omitempty" tf:"mode"`
	// The URI of an existing persistent disk to attach under the specified device-name in the format
	// 'projects/project-id/zones/zone/disks/disk-name'.
	Source *string `json:"source" tf:"source"`
}

type PerInstanceConfigSpecPreservedState struct {
	// Stateful disks for the instance.
	// +optional
	Disk []PerInstanceConfigSpecPreservedStateDisk `json:"disk,omitempty" tf:"disk"`
	// Preserved metadata defined for this instance. This is a list of key->value pairs.
	// +optional
	Metadata *map[string]string `json:"metadata,omitempty" tf:"metadata"`
}

type PerInstanceConfigSpec struct {
	KubeformOutput *PerInstanceConfigSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource PerInstanceConfigSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type PerInstanceConfigSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The instance group manager this instance config is part of.
	InstanceGroupManager *string `json:"instanceGroupManager" tf:"instance_group_manager"`
	// +optional
	MinimalAction *string `json:"minimalAction,omitempty" tf:"minimal_action"`
	// +optional
	MostDisruptiveAllowedAction *string `json:"mostDisruptiveAllowedAction,omitempty" tf:"most_disruptive_allowed_action"`
	// The name for this per-instance config and its corresponding instance.
	Name *string `json:"name" tf:"name"`
	// The preserved state for this instance.
	// +optional
	PreservedState *PerInstanceConfigSpecPreservedState `json:"preservedState,omitempty" tf:"preserved_state"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// +optional
	RemoveInstanceStateOnDestroy *bool `json:"removeInstanceStateOnDestroy,omitempty" tf:"remove_instance_state_on_destroy"`
	// Zone where the containing instance group manager is located
	// +optional
	Zone *string `json:"zone,omitempty" tf:"zone"`
}

type PerInstanceConfigStatus struct {
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

// PerInstanceConfigList is a list of PerInstanceConfigs
type PerInstanceConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PerInstanceConfig CRD objects
	Items []PerInstanceConfig `json:"items,omitempty"`
}
