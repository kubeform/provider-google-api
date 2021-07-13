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

type DirectoryDomainTrust struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DirectoryDomainTrustSpec   `json:"spec,omitempty"`
	Status            DirectoryDomainTrustStatus `json:"status,omitempty"`
}

type DirectoryDomainTrustSpec struct {
	State *DirectoryDomainTrustSpecResource `json:"state,omitempty" tf:"-"`

	Resource DirectoryDomainTrustSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type DirectoryDomainTrustSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
	// https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
	Domain *string `json:"domain" tf:"domain"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// Whether the trusted side has forest/domain wide access or selective access to an approved set of resources.
	// +optional
	SelectiveAuthentication *bool `json:"selectiveAuthentication,omitempty" tf:"selective_authentication"`
	// The target DNS server IP addresses which can resolve the remote domain involved in the trust.
	TargetDNSIPAddresses []string `json:"targetDNSIPAddresses" tf:"target_dns_ip_addresses"`
	// The fully qualified target domain name which will be in trust with the current domain.
	TargetDomainName *string `json:"targetDomainName" tf:"target_domain_name"`
	// The trust direction, which decides if the current domain is trusted, trusting, or both. Possible values: ["INBOUND", "OUTBOUND", "BIDIRECTIONAL"]
	TrustDirection *string `json:"trustDirection" tf:"trust_direction"`
	// The trust secret used for the handshake with the target domain. This will not be stored.
	TrustHandshakeSecret *string `json:"-" sensitive:"true" tf:"trust_handshake_secret"`
	// The type of trust represented by the trust resource. Possible values: ["FOREST", "EXTERNAL"]
	TrustType *string `json:"trustType" tf:"trust_type"`
}

type DirectoryDomainTrustStatus struct {
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

// DirectoryDomainTrustList is a list of DirectoryDomainTrusts
type DirectoryDomainTrustList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DirectoryDomainTrust CRD objects
	Items []DirectoryDomainTrust `json:"items,omitempty"`
}
