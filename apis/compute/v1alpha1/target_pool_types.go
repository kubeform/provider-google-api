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

type TargetPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TargetPoolSpec   `json:"spec,omitempty"`
	Status            TargetPoolStatus `json:"status,omitempty"`
}

type TargetPoolSpec struct {
	State *TargetPoolSpecResource `json:"state,omitempty" tf:"-"`

	Resource TargetPoolSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type TargetPoolSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// URL to the backup target pool. Must also set failover_ratio.
	// +optional
	BackupPool *string `json:"backupPool,omitempty" tf:"backup_pool"`
	// Textual description field.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Ratio (0 to 1) of failed nodes before using the backup pool (which must also be set).
	// +optional
	FailoverRatio *float64 `json:"failoverRatio,omitempty" tf:"failover_ratio"`
	// List of zero or one health check name or self_link. Only legacy google_compute_http_health_check is supported.
	// +optional
	HealthChecks []string `json:"healthChecks,omitempty" tf:"health_checks"`
	// List of instances in the pool. They can be given as URLs, or in the form of "zone/name". Note that the instances need not exist at the time of target pool creation, so there is no need to use the Terraform interpolators to create a dependency on the instances from the target pool.
	// +optional
	Instances []string `json:"instances,omitempty" tf:"instances"`
	// A unique name for the resource, required by GCE. Changing this forces a new resource to be created.
	Name *string `json:"name" tf:"name"`
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// Where the target pool resides. Defaults to project region.
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// The URI of the created resource.
	// +optional
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link"`
	// How to distribute load. Options are "NONE" (no affinity). "CLIENT_IP" (hash of the source/dest addresses / ports), and "CLIENT_IP_PROTO" also includes the protocol (default "NONE").
	// +optional
	SessionAffinity *string `json:"sessionAffinity,omitempty" tf:"session_affinity"`
}

type TargetPoolStatus struct {
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

// TargetPoolList is a list of TargetPools
type TargetPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of TargetPool CRD objects
	Items []TargetPool `json:"items,omitempty"`
}
