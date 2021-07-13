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

type AppProfile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppProfileSpec   `json:"spec,omitempty"`
	Status            AppProfileStatus `json:"status,omitempty"`
}

type AppProfileSpecSingleClusterRouting struct {
	// If true, CheckAndMutateRow and ReadModifyWriteRow requests are allowed by this app profile.
	// It is unsafe to send these requests to the same table/row/column in multiple clusters.
	// +optional
	AllowTransactionalWrites *bool `json:"allowTransactionalWrites,omitempty" tf:"allow_transactional_writes"`
	// The cluster to which read/write requests should be routed.
	ClusterID *string `json:"clusterID" tf:"cluster_id"`
}

type AppProfileSpec struct {
	State *AppProfileSpecResource `json:"state,omitempty" tf:"-"`

	Resource AppProfileSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type AppProfileSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The unique name of the app profile in the form '[_a-zA-Z0-9][-_.a-zA-Z0-9]*'.
	AppProfileID *string `json:"appProfileID" tf:"app_profile_id"`
	// Long form description of the use case for this app profile.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// If true, ignore safety checks when deleting/updating the app profile.
	// +optional
	IgnoreWarnings *bool `json:"ignoreWarnings,omitempty" tf:"ignore_warnings"`
	// The name of the instance to create the app profile within.
	// +optional
	Instance *string `json:"instance,omitempty" tf:"instance"`
	// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
	// in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
	// consistency to improve availability.
	// +optional
	MultiClusterRoutingUseAny *bool `json:"multiClusterRoutingUseAny,omitempty" tf:"multi_cluster_routing_use_any"`
	// The unique name of the requested app profile. Values are of the form 'projects/<project>/instances/<instance>/appProfiles/<appProfileId>'.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// Use a single-cluster routing policy.
	// +optional
	SingleClusterRouting *AppProfileSpecSingleClusterRouting `json:"singleClusterRouting,omitempty" tf:"single_cluster_routing"`
}

type AppProfileStatus struct {
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

// AppProfileList is a list of AppProfiles
type AppProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppProfile CRD objects
	Items []AppProfile `json:"items,omitempty"`
}
