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

type SslCert struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SslCertSpec   `json:"spec,omitempty"`
	Status            SslCertStatus `json:"status,omitempty"`
}

type SslCertSpec struct {
	State *SslCertSpecResource `json:"state,omitempty" tf:"-"`

	Resource SslCertSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SslCertSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The actual certificate data for this client certificate.
	// +optional
	Cert *string `json:"cert,omitempty" tf:"cert"`
	// The serial number extracted from the certificate data.
	// +optional
	CertSerialNumber *string `json:"certSerialNumber,omitempty" tf:"cert_serial_number"`
	// The common name to be used in the certificate to identify the client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
	CommonName *string `json:"commonName" tf:"common_name"`
	// The time when the certificate was created in RFC 3339 format, for example 2012-11-15T16:19:00.094Z.
	// +optional
	CreateTime *string `json:"createTime,omitempty" tf:"create_time"`
	// The time when the certificate expires in RFC 3339 format, for example 2012-11-15T16:19:00.094Z.
	// +optional
	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time"`
	// The name of the Cloud SQL instance. Changing this forces a new resource to be created.
	Instance *string `json:"instance" tf:"instance"`
	// The private key associated with the client certificate.
	// +optional
	PrivateKey *string `json:"-" sensitive:"true" tf:"private_key"`
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The CA cert of the server this client cert was generated from.
	// +optional
	ServerCaCert *string `json:"serverCaCert,omitempty" tf:"server_ca_cert"`
	// The SHA1 Fingerprint of the certificate.
	// +optional
	Sha1Fingerprint *string `json:"sha1Fingerprint,omitempty" tf:"sha1_fingerprint"`
}

type SslCertStatus struct {
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

// SslCertList is a list of SslCerts
type SslCertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SslCert CRD objects
	Items []SslCert `json:"items,omitempty"`
}
