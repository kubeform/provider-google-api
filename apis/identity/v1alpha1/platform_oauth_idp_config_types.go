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

type PlatformOauthIdpConfig struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PlatformOauthIdpConfigSpec   `json:"spec,omitempty"`
	Status            PlatformOauthIdpConfigStatus `json:"status,omitempty"`
}

type PlatformOauthIdpConfigSpec struct {
	State *PlatformOauthIdpConfigSpecResource `json:"state,omitempty" tf:"-"`

	Resource PlatformOauthIdpConfigSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type PlatformOauthIdpConfigSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The client id of an OAuth client.
	ClientID *string `json:"clientID" tf:"client_id"`
	// The client secret of the OAuth client, to enable OIDC code flow.
	// +optional
	ClientSecret *string `json:"clientSecret,omitempty" tf:"client_secret"`
	// Human friendly display name.
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// If this config allows users to sign in with the provider.
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// For OIDC Idps, the issuer identifier.
	Issuer *string `json:"issuer" tf:"issuer"`
	// The name of the OauthIdpConfig. Must start with 'oidc.'.
	Name *string `json:"name" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
}

type PlatformOauthIdpConfigStatus struct {
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

// PlatformOauthIdpConfigList is a list of PlatformOauthIdpConfigs
type PlatformOauthIdpConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PlatformOauthIdpConfig CRD objects
	Items []PlatformOauthIdpConfig `json:"items,omitempty"`
}
