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

type Route struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteSpec   `json:"spec,omitempty"`
	Status            RouteStatus `json:"status,omitempty"`
}

type RouteSpec struct {
	State *RouteSpecResource `json:"state,omitempty" tf:"-"`

	Resource RouteSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type RouteSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// An optional description of this resource. Provide this property
	// when you create the resource.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// The destination range of outgoing packets that this route applies to.
	// Only IPv4 is supported.
	DestRange *string `json:"destRange" tf:"dest_range"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name *string `json:"name" tf:"name"`
	// The network that this route applies to.
	Network *string `json:"network" tf:"network"`
	// URL to a gateway that should handle matching packets.
	// Currently, you can only specify the internet gateway, using a full or
	// partial valid URL:
	// * 'https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway'
	// * 'projects/project/global/gateways/default-internet-gateway'
	// * 'global/gateways/default-internet-gateway'
	// * The string 'default-internet-gateway'.
	// +optional
	NextHopGateway *string `json:"nextHopGateway,omitempty" tf:"next_hop_gateway"`
	// The IP address or URL to a forwarding rule of type
	// loadBalancingScheme=INTERNAL that should handle matching
	// packets.
	//
	// With the GA provider you can only specify the forwarding
	// rule as a partial or full URL. For example, the following
	// are all valid values:
	// * 10.128.0.56
	// * https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
	// * regions/region/forwardingRules/forwardingRule
	//
	// When the beta provider, you can also specify the IP address
	// of a forwarding rule from the same VPC or any peered VPC.
	//
	// Note that this can only be used when the destinationRange is
	// a public (non-RFC 1918) IP CIDR range.
	// +optional
	NextHopIlb *string `json:"nextHopIlb,omitempty" tf:"next_hop_ilb"`
	// URL to an instance that should handle matching packets.
	// You can specify this as a full or partial URL. For example:
	// * 'https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance'
	// * 'projects/project/zones/zone/instances/instance'
	// * 'zones/zone/instances/instance'
	// * Just the instance name, with the zone in 'next_hop_instance_zone'.
	// +optional
	NextHopInstance *string `json:"nextHopInstance,omitempty" tf:"next_hop_instance"`
	// The zone of the instance specified in next_hop_instance. Omit if next_hop_instance is specified as a URL.
	// +optional
	NextHopInstanceZone *string `json:"nextHopInstanceZone,omitempty" tf:"next_hop_instance_zone"`
	// Network IP address of an instance that should handle matching packets.
	// +optional
	NextHopIP *string `json:"nextHopIP,omitempty" tf:"next_hop_ip"`
	// URL to a Network that should handle matching packets.
	// +optional
	NextHopNetwork *string `json:"nextHopNetwork,omitempty" tf:"next_hop_network"`
	// URL to a VpnTunnel that should handle matching packets.
	// +optional
	NextHopVPNTunnel *string `json:"nextHopVPNTunnel,omitempty" tf:"next_hop_vpn_tunnel"`
	// The priority of this route. Priority is used to break ties in cases
	// where there is more than one matching route of equal prefix length.
	//
	// In the case of two routes with equal prefix length, the one with the
	// lowest-numbered priority value wins.
	//
	// Default value is 1000. Valid range is 0 through 65535.
	// +optional
	Priority *int64 `json:"priority,omitempty" tf:"priority"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// +optional
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link"`
	// A list of instance tags to which this route applies.
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
}

type RouteStatus struct {
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

// RouteList is a list of Routes
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Route CRD objects
	Items []Route `json:"items,omitempty"`
}
