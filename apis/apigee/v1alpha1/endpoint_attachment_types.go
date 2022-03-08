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

type EndpointAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointAttachmentSpec   `json:"spec,omitempty"`
	Status            EndpointAttachmentStatus `json:"status,omitempty"`
}

type EndpointAttachmentSpec struct {
	State *EndpointAttachmentSpecResource `json:"state,omitempty" tf:"-"`

	Resource EndpointAttachmentSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type EndpointAttachmentSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the endpoint attachment.
	EndpointAttachmentID *string `json:"endpointAttachmentID" tf:"endpoint_attachment_id"`
	// Host that can be used in either HTTP Target Endpoint directly, or as the host in Target Server.
	// +optional
	Host *string `json:"host,omitempty" tf:"host"`
	// Location of the endpoint attachment.
	Location *string `json:"location" tf:"location"`
	// Name of the Endpoint Attachment in the following format:
	// organizations/{organization}/endpointAttachments/{endpointAttachment}.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The Apigee Organization associated with the Apigee instance,
	// in the format 'organizations/{{org_name}}'.
	OrgID *string `json:"orgID" tf:"org_id"`
	// Format: projects/*/regions/*/serviceAttachments/*
	ServiceAttachment *string `json:"serviceAttachment" tf:"service_attachment"`
}

type EndpointAttachmentStatus struct {
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

// EndpointAttachmentList is a list of EndpointAttachments
type EndpointAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EndpointAttachment CRD objects
	Items []EndpointAttachment `json:"items,omitempty"`
}
