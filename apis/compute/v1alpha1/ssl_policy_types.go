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

type SslPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SslPolicySpec   `json:"spec,omitempty"`
	Status            SslPolicyStatus `json:"status,omitempty"`
}

type SslPolicySpec struct {
	State *SslPolicySpecResource `json:"state,omitempty" tf:"-"`

	Resource SslPolicySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SslPolicySpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Creation timestamp in RFC3339 text format.
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp"`
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. This can be one of
	// 'COMPATIBLE', 'MODERN', 'RESTRICTED', or 'CUSTOM'. If using 'CUSTOM',
	// the set of SSL features to enable must be specified in the
	// 'customFeatures' field.
	//
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for which ciphers are available to use. **Note**: this argument
	// *must* be present when using the 'CUSTOM' profile. This argument
	// *must not* be present when using any other profile.
	// +optional
	CustomFeatures []string `json:"customFeatures,omitempty" tf:"custom_features"`
	// An optional description of this resource.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// The list of features enabled in the SSL policy.
	// +optional
	EnabledFeatures []string `json:"enabledFeatures,omitempty" tf:"enabled_features"`
	// Fingerprint of this resource. A hash of the contents stored in this
	// object. This field is used in optimistic locking.
	// +optional
	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint"`
	// The minimum version of SSL protocol that can be used by the clients
	// to establish a connection with the load balancer. Default value: "TLS_1_0" Possible values: ["TLS_1_0", "TLS_1_1", "TLS_1_2"]
	// +optional
	MinTlsVersion *string `json:"minTlsVersion,omitempty" tf:"min_tls_version"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `json:"name" tf:"name"`
	// Profile specifies the set of SSL features that can be used by the
	// load balancer when negotiating SSL with clients. If using 'CUSTOM',
	// the set of SSL features to enable must be specified in the
	// 'customFeatures' field.
	//
	// See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	// for information on what cipher suites each profile provides. If
	// 'CUSTOM' is used, the 'custom_features' attribute **must be set**. Default value: "COMPATIBLE" Possible values: ["COMPATIBLE", "MODERN", "RESTRICTED", "CUSTOM"]
	// +optional
	Profile *string `json:"profile,omitempty" tf:"profile"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// +optional
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link"`
}

type SslPolicyStatus struct {
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

// SslPolicyList is a list of SslPolicys
type SslPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SslPolicy CRD objects
	Items []SslPolicy `json:"items,omitempty"`
}
