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

type RouterPeer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouterPeerSpec   `json:"spec,omitempty"`
	Status            RouterPeerStatus `json:"status,omitempty"`
}

type RouterPeerSpecAdvertisedIPRanges struct {
	// User-specified description for the IP range.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// The IP range to advertise. The value must be a
	// CIDR-formatted string.
	Range *string `json:"range" tf:"range"`
}

type RouterPeerSpecBfd struct {
	// The minimum interval, in milliseconds, between BFD control packets
	// received from the peer router. The actual value is negotiated
	// between the two routers and is equal to the greater of this value
	// and the transmit interval of the other router. If set, this value
	// must be between 1000 and 30000.
	// +optional
	MinReceiveInterval *int64 `json:"minReceiveInterval,omitempty" tf:"min_receive_interval"`
	// The minimum interval, in milliseconds, between BFD control packets
	// transmitted to the peer router. The actual value is negotiated
	// between the two routers and is equal to the greater of this value
	// and the corresponding receive interval of the other router. If set,
	// this value must be between 1000 and 30000.
	// +optional
	MinTransmitInterval *int64 `json:"minTransmitInterval,omitempty" tf:"min_transmit_interval"`
	// The number of consecutive BFD packets that must be missed before
	// BFD declares that a peer is unavailable. If set, the value must
	// be a value between 5 and 16.
	// +optional
	Multiplier *int64 `json:"multiplier,omitempty" tf:"multiplier"`
	// The BFD session initialization mode for this BGP peer.
	// If set to 'ACTIVE', the Cloud Router will initiate the BFD session
	// for this BGP peer. If set to 'PASSIVE', the Cloud Router will wait
	// for the peer router to initiate the BFD session for this BGP peer.
	// If set to 'DISABLED', BFD is disabled for this BGP peer. Possible values: ["ACTIVE", "DISABLED", "PASSIVE"]
	SessionInitializationMode *string `json:"sessionInitializationMode" tf:"session_initialization_mode"`
}

type RouterPeerSpec struct {
	State *RouterPeerSpecResource `json:"state,omitempty" tf:"-"`

	Resource RouterPeerSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type RouterPeerSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// User-specified flag to indicate which mode to use for advertisement.
	// Valid values of this enum field are: 'DEFAULT', 'CUSTOM' Default value: "DEFAULT" Possible values: ["DEFAULT", "CUSTOM"]
	// +optional
	AdvertiseMode *string `json:"advertiseMode,omitempty" tf:"advertise_mode"`
	// User-specified list of prefix groups to advertise in custom
	// mode, which can take one of the following options:
	//
	// * 'ALL_SUBNETS': Advertises all available subnets, including peer VPC subnets.
	// * 'ALL_VPC_SUBNETS': Advertises the router's own VPC subnets.
	// * 'ALL_PEER_VPC_SUBNETS': Advertises peer subnets of the router's VPC network.
	//
	//
	// Note that this field can only be populated if advertiseMode is 'CUSTOM'
	// and overrides the list defined for the router (in the "bgp" message).
	// These groups are advertised in addition to any specified prefixes.
	// Leave this field blank to advertise no custom groups.
	// +optional
	AdvertisedGroups []string `json:"advertisedGroups,omitempty" tf:"advertised_groups"`
	// User-specified list of individual IP ranges to advertise in
	// custom mode. This field can only be populated if advertiseMode
	// is 'CUSTOM' and is advertised to all peers of the router. These IP
	// ranges will be advertised in addition to any specified groups.
	// Leave this field blank to advertise no custom IP ranges.
	// +optional
	AdvertisedIPRanges []RouterPeerSpecAdvertisedIPRanges `json:"advertisedIPRanges,omitempty" tf:"advertised_ip_ranges"`
	// The priority of routes advertised to this BGP peer.
	// Where there is more than one matching route of maximum
	// length, the routes with the lowest priority value win.
	// +optional
	AdvertisedRoutePriority *int64 `json:"advertisedRoutePriority,omitempty" tf:"advertised_route_priority"`
	// BFD configuration for the BGP peering.
	// +optional
	Bfd *RouterPeerSpecBfd `json:"bfd,omitempty" tf:"bfd"`
	// The status of the BGP peer connection. If set to false, any active session
	// with the peer is terminated and all associated routing information is removed.
	// If set to true, the peer connection can be established with routing information.
	// The default is true.
	// +optional
	Enable *bool `json:"enable,omitempty" tf:"enable"`
	// Name of the interface the BGP peer is associated with.
	Interface *string `json:"interface" tf:"interface"`
	// IP address of the interface inside Google Cloud Platform.
	// Only IPv4 is supported.
	// +optional
	IpAddress *string `json:"ipAddress,omitempty" tf:"ip_address"`
	// The resource that configures and manages this BGP peer.
	//
	// * 'MANAGED_BY_USER' is the default value and can be managed by
	// you or other users
	// * 'MANAGED_BY_ATTACHMENT' is a BGP peer that is configured and
	// managed by Cloud Interconnect, specifically by an
	// InterconnectAttachment of type PARTNER. Google automatically
	// creates, updates, and deletes this type of BGP peer when the
	// PARTNER InterconnectAttachment is created, updated,
	// or deleted.
	// +optional
	ManagementType *string `json:"managementType,omitempty" tf:"management_type"`
	// Name of this BGP peer. The name must be 1-63 characters long,
	// and comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name *string `json:"name" tf:"name"`
	// Peer BGP Autonomous System Number (ASN).
	// Each BGP interface may use a different value.
	PeerAsn *int64 `json:"peerAsn" tf:"peer_asn"`
	// IP address of the BGP interface outside Google Cloud Platform.
	// Only IPv4 is supported.
	PeerIPAddress *string `json:"peerIPAddress" tf:"peer_ip_address"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// Region where the router and BgpPeer reside.
	// If it is not provided, the provider region is used.
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// The name of the Cloud Router in which this BgpPeer will be configured.
	Router *string `json:"router" tf:"router"`
}

type RouterPeerStatus struct {
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

// RouterPeerList is a list of RouterPeers
type RouterPeerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RouterPeer CRD objects
	Items []RouterPeer `json:"items,omitempty"`
}
