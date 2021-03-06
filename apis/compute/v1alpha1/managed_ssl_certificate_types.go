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

type ManagedSslCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedSslCertificateSpec   `json:"spec,omitempty"`
	Status            ManagedSslCertificateStatus `json:"status,omitempty"`
}

type ManagedSslCertificateSpecManaged struct {
	// Domains for which a managed SSL certificate will be valid.  Currently,
	// there can be up to 100 domains in this list.
	// +kubebuilder:validation:MaxItems=100
	Domains []string `json:"domains" tf:"domains"`
}

type ManagedSslCertificateSpec struct {
	State *ManagedSslCertificateSpecResource `json:"state,omitempty" tf:"-"`

	Resource ManagedSslCertificateSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ManagedSslCertificateSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The unique identifier for the resource.
	// +optional
	CertificateID *int64 `json:"certificateID,omitempty" tf:"certificate_id"`
	// Creation timestamp in RFC3339 text format.
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp"`
	// An optional description of this resource.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Expire time of the certificate.
	// +optional
	ExpireTime *string `json:"expireTime,omitempty" tf:"expire_time"`
	// Properties relevant to a managed certificate.  These will be used if the
	// certificate is managed (as indicated by a value of 'MANAGED' in 'type').
	// +optional
	Managed *ManagedSslCertificateSpecManaged `json:"managed,omitempty" tf:"managed"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	//
	// These are in the same namespace as the managed SSL certificates.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// +optional
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link"`
	// Domains associated with the certificate via Subject Alternative Name.
	// +optional
	SubjectAlternativeNames []string `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names"`
	// Enum field whose value is always 'MANAGED' - used to signal to the API
	// which type this is. Default value: "MANAGED" Possible values: ["MANAGED"]
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type ManagedSslCertificateStatus struct {
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

// ManagedSslCertificateList is a list of ManagedSslCertificates
type ManagedSslCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ManagedSslCertificate CRD objects
	Items []ManagedSslCertificate `json:"items,omitempty"`
}
