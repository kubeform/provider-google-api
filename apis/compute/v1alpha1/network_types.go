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

type Network struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkSpec   `json:"spec,omitempty"`
	Status            NetworkStatus `json:"status,omitempty"`
}

type NetworkSpec struct {
	State *NetworkSpecResource `json:"state,omitempty" tf:"-"`

	Resource NetworkSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type NetworkSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// When set to 'true', the network is created in "auto subnet mode" and
	// it will create a subnet for each region automatically across the
	// '10.128.0.0/9' address range.
	//
	// When set to 'false', the network is created in "custom subnet mode" so
	// the user can explicitly connect subnetwork resources.
	// +optional
	AutoCreateSubnetworks *bool `json:"autoCreateSubnetworks,omitempty" tf:"auto_create_subnetworks"`
	// +optional
	DeleteDefaultRoutesOnCreate *bool `json:"deleteDefaultRoutesOnCreate,omitempty" tf:"delete_default_routes_on_create"`
	// An optional description of this resource. The resource must be
	// recreated to modify this field.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// The gateway address for default routing out of the network. This value
	// is selected by GCP.
	// +optional
	GatewayIpv4 *string `json:"gatewayIpv4,omitempty" tf:"gateway_ipv4"`
	// Maximum Transmission Unit in bytes. The minimum value for this field is 1460
	// and the maximum value is 1500 bytes.
	// +optional
	Mtu *int64 `json:"mtu,omitempty" tf:"mtu"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `json:"name" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The network-wide routing mode to use. If set to 'REGIONAL', this
	// network's cloud routers will only advertise routes with subnetworks
	// of this network in the same region as the router. If set to 'GLOBAL',
	// this network's cloud routers will advertise routes with all
	// subnetworks of this network, across regions. Possible values: ["REGIONAL", "GLOBAL"]
	// +optional
	RoutingMode *string `json:"routingMode,omitempty" tf:"routing_mode"`
	// +optional
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link"`
}

type NetworkStatus struct {
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

// NetworkList is a list of Networks
type NetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Network CRD objects
	Items []Network `json:"items,omitempty"`
}
