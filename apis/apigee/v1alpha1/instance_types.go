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

type Instance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec,omitempty"`
	Status            InstanceStatus `json:"status,omitempty"`
}

type InstanceSpec struct {
	State *InstanceSpecResource `json:"state,omitempty" tf:"-"`

	Resource InstanceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type InstanceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Description of the instance.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
	// Use the following format: 'projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)'
	// +optional
	DiskEncryptionKeyName *string `json:"diskEncryptionKeyName,omitempty" tf:"disk_encryption_key_name"`
	// Display name of the instance.
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// Output only. Hostname or IP address of the exposed Apigee endpoint used by clients to connect to the service.
	// +optional
	Host *string `json:"host,omitempty" tf:"host"`
	// IP range represents the customer-provided CIDR block of length 22 that will be used for
	// the Apigee instance creation. This optional range, if provided, should be freely
	// available as part of larger named range the customer has allocated to the Service
	// Networking peering. If this is not provided, Apigee will automatically request for any
	// available /22 CIDR block from Service Networking. The customer should use this CIDR block
	// for configuring their firewall needs to allow traffic from Apigee.
	// Input format: "a.b.c.d/22"
	// +optional
	IpRange *string `json:"ipRange,omitempty" tf:"ip_range"`
	// Compute Engine location where the instance resides. For trial organization
	// subscriptions, the location must be a Compute Engine zone. For paid organization
	// subscriptions, it should correspond to a Compute Engine region.
	Location *string `json:"location" tf:"location"`
	// Resource ID of the instance.
	Name *string `json:"name" tf:"name"`
	// The Apigee Organization associated with the Apigee instance,
	// in the format 'organizations/{{org_name}}'.
	OrgID *string `json:"orgID" tf:"org_id"`
	// The size of the CIDR block range that will be reserved by the instance. For valid values,
	// see [CidrRange](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances#CidrRange) on the documentation.
	// +optional
	PeeringCIDRRange *string `json:"peeringCIDRRange,omitempty" tf:"peering_cidr_range"`
	// Output only. Port number of the exposed Apigee endpoint.
	// +optional
	Port *string `json:"port,omitempty" tf:"port"`
}

type InstanceStatus struct {
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

// InstanceList is a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Instance CRD objects
	Items []Instance `json:"items,omitempty"`
}
