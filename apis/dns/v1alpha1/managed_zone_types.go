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

type ManagedZone struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedZoneSpec   `json:"spec,omitempty"`
	Status            ManagedZoneStatus `json:"status,omitempty"`
}

type ManagedZoneSpecDnssecConfigDefaultKeySpecs struct {
	// String mnemonic specifying the DNSSEC algorithm of this key Possible values: ["ecdsap256sha256", "ecdsap384sha384", "rsasha1", "rsasha256", "rsasha512"]
	// +optional
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm"`
	// Length of the keys in bits
	// +optional
	KeyLength *int64 `json:"keyLength,omitempty" tf:"key_length"`
	// Specifies whether this is a key signing key (KSK) or a zone
	// signing key (ZSK). Key signing keys have the Secure Entry
	// Point flag set and, when active, will only be used to sign
	// resource record sets of type DNSKEY. Zone signing keys do
	// not have the Secure Entry Point flag set and will be used
	// to sign all other types of resource record sets. Possible values: ["keySigning", "zoneSigning"]
	// +optional
	KeyType *string `json:"keyType,omitempty" tf:"key_type"`
	// Identifies what kind of resource this is
	// +optional
	Kind *string `json:"kind,omitempty" tf:"kind"`
}

type ManagedZoneSpecDnssecConfig struct {
	// Specifies parameters that will be used for generating initial DnsKeys
	// for this ManagedZone. If you provide a spec for keySigning or zoneSigning,
	// you must also provide one for the other.
	// default_key_specs can only be updated when the state is 'off'.
	// +optional
	DefaultKeySpecs []ManagedZoneSpecDnssecConfigDefaultKeySpecs `json:"defaultKeySpecs,omitempty" tf:"default_key_specs"`
	// Identifies what kind of resource this is
	// +optional
	Kind *string `json:"kind,omitempty" tf:"kind"`
	// Specifies the mechanism used to provide authenticated denial-of-existence responses.
	// non_existence can only be updated when the state is 'off'. Possible values: ["nsec", "nsec3"]
	// +optional
	NonExistence *string `json:"nonExistence,omitempty" tf:"non_existence"`
	// Specifies whether DNSSEC is enabled, and what mode it is in Possible values: ["off", "on", "transfer"]
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
}

type ManagedZoneSpecForwardingConfigTargetNameServers struct {
	// Forwarding path for this TargetNameServer. If unset or 'default' Cloud DNS will make forwarding
	// decision based on address ranges, i.e. RFC1918 addresses go to the VPC, Non-RFC1918 addresses go
	// to the Internet. When set to 'private', Cloud DNS will always send queries through VPC for this target Possible values: ["default", "private"]
	// +optional
	ForwardingPath *string `json:"forwardingPath,omitempty" tf:"forwarding_path"`
	// IPv4 address of a target name server.
	Ipv4Address *string `json:"ipv4Address" tf:"ipv4_address"`
}

type ManagedZoneSpecForwardingConfig struct {
	// List of target name servers to forward to. Cloud DNS will
	// select the best available name server if more than
	// one target is given.
	TargetNameServers []ManagedZoneSpecForwardingConfigTargetNameServers `json:"targetNameServers" tf:"target_name_servers"`
}

type ManagedZoneSpecPeeringConfigTargetNetwork struct {
	// The id or fully qualified URL of the VPC network to forward queries to.
	// This should be formatted like 'projects/{project}/global/networks/{network}' or
	// 'https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}'
	NetworkURL *string `json:"networkURL" tf:"network_url"`
}

type ManagedZoneSpecPeeringConfig struct {
	// The network with which to peer.
	TargetNetwork *ManagedZoneSpecPeeringConfigTargetNetwork `json:"targetNetwork" tf:"target_network"`
}

type ManagedZoneSpecPrivateVisibilityConfigNetworks struct {
	// The id or fully qualified URL of the VPC network to bind to.
	// This should be formatted like 'projects/{project}/global/networks/{network}' or
	// 'https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}'
	NetworkURL *string `json:"networkURL" tf:"network_url"`
}

type ManagedZoneSpecPrivateVisibilityConfig struct {
	// The list of VPC networks that can see this zone. Until the provider updates to use the Terraform 0.12 SDK in a future release, you
	// may experience issues with this resource while updating. If you've defined a 'networks' block and
	// add another 'networks' block while keeping the old block, Terraform will see an incorrect diff
	// and apply an incorrect update to the resource. If you encounter this issue, remove all 'networks'
	// blocks in an update and then apply another update adding all of them back simultaneously.
	Networks []ManagedZoneSpecPrivateVisibilityConfigNetworks `json:"networks" tf:"networks"`
}

type ManagedZoneSpec struct {
	State *ManagedZoneSpecResource `json:"state,omitempty" tf:"-"`

	Resource ManagedZoneSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ManagedZoneSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// A textual description field. Defaults to 'Managed by Terraform'.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// The DNS name of this managed zone, for instance "example.com.".
	DnsName *string `json:"dnsName" tf:"dns_name"`
	// DNSSEC configuration
	// +optional
	DnssecConfig *ManagedZoneSpecDnssecConfig `json:"dnssecConfig,omitempty" tf:"dnssec_config"`
	// +optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy"`
	// The presence for this field indicates that outbound forwarding is enabled
	// for this zone. The value of this field contains the set of destinations
	// to forward to.
	// +optional
	ForwardingConfig *ManagedZoneSpecForwardingConfig `json:"forwardingConfig,omitempty" tf:"forwarding_config"`
	// A set of key/value label pairs to assign to this ManagedZone.
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// User assigned name for this resource.
	// Must be unique within the project.
	Name *string `json:"name" tf:"name"`
	// Delegate your managed_zone to these virtual name servers;
	// defined by the server
	// +optional
	NameServers []string `json:"nameServers,omitempty" tf:"name_servers"`
	// The presence of this field indicates that DNS Peering is enabled for this
	// zone. The value of this field contains the network to peer with.
	// +optional
	PeeringConfig *ManagedZoneSpecPeeringConfig `json:"peeringConfig,omitempty" tf:"peering_config"`
	// For privately visible zones, the set of Virtual Private Cloud
	// resources that the zone is visible from.
	// +optional
	PrivateVisibilityConfig *ManagedZoneSpecPrivateVisibilityConfig `json:"privateVisibilityConfig,omitempty" tf:"private_visibility_config"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The zone's visibility: public zones are exposed to the Internet,
	// while private zones are visible only to Virtual Private Cloud resources. Default value: "public" Possible values: ["private", "public"]
	// +optional
	Visibility *string `json:"visibility,omitempty" tf:"visibility"`
}

type ManagedZoneStatus struct {
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

// ManagedZoneList is a list of ManagedZones
type ManagedZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ManagedZone CRD objects
	Items []ManagedZone `json:"items,omitempty"`
}
