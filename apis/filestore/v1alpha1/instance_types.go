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

type Instance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec,omitempty"`
	Status            InstanceStatus `json:"status,omitempty"`
}

type InstanceSpecFileShares struct {
	// File share capacity in GiB. This must be at least 1024 GiB
	// for the standard tier, or 2560 GiB for the premium tier.
	CapacityGb *int64 `json:"capacityGb" tf:"capacity_gb"`
	// The name of the fileshare (16 characters or less)
	Name *string `json:"name" tf:"name"`
}

type InstanceSpecNetworks struct {
	// A list of IPv4 or IPv6 addresses.
	// +optional
	IpAddresses []string `json:"ipAddresses,omitempty" tf:"ip_addresses"`
	// IP versions for which the instance has
	// IP addresses assigned. Possible values: ["ADDRESS_MODE_UNSPECIFIED", "MODE_IPV4", "MODE_IPV6"]
	Modes []string `json:"modes" tf:"modes"`
	// The name of the GCE VPC network to which the
	// instance is connected.
	Network *string `json:"network" tf:"network"`
	// A /29 CIDR block that identifies the range of IP
	// addresses reserved for this instance.
	// +optional
	ReservedIPRange *string `json:"reservedIPRange,omitempty" tf:"reserved_ip_range"`
}

type InstanceSpec struct {
	State *InstanceSpecResource `json:"state,omitempty" tf:"-"`

	Resource InstanceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type InstanceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Creation timestamp in RFC3339 text format.
	// +optional
	CreateTime *string `json:"createTime,omitempty" tf:"create_time"`
	// A description of the instance.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Server-specified ETag for the instance resource to prevent
	// simultaneous updates from overwriting each other.
	// +optional
	Etag *string `json:"etag,omitempty" tf:"etag"`
	// File system shares on the instance. For this version, only a
	// single file share is supported.
	FileShares *InstanceSpecFileShares `json:"fileShares" tf:"file_shares"`
	// Resource labels to represent user-provided metadata.
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
	// +optional
	Location *string `json:"location,omitempty" tf:"location"`
	// The resource name of the instance.
	Name *string `json:"name" tf:"name"`
	// VPC networks to which the instance is connected. For this version,
	// only a single network is supported.
	// +kubebuilder:validation:MinItems=1
	Networks []InstanceSpecNetworks `json:"networks" tf:"networks"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The service tier of the instance.
	// Possible values include: STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD and ENTERPRISE (beta only)
	Tier *string `json:"tier" tf:"tier"`
	// The name of the Filestore zone of the instance.
	// +optional
	// Deprecated
	Zone *string `json:"zone,omitempty" tf:"zone"`
}

type InstanceStatus struct {
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

// InstanceList is a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Instance CRD objects
	Items []Instance `json:"items,omitempty"`
}
