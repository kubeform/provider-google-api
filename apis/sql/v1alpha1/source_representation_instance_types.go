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

type SourceRepresentationInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SourceRepresentationInstanceSpec   `json:"spec,omitempty"`
	Status            SourceRepresentationInstanceStatus `json:"status,omitempty"`
}

type SourceRepresentationInstanceSpec struct {
	State *SourceRepresentationInstanceSpecResource `json:"state,omitempty" tf:"-"`

	Resource SourceRepresentationInstanceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SourceRepresentationInstanceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The MySQL version running on your source database server. Possible values: ["MYSQL_5_5", "MYSQL_5_6", "MYSQL_5_7", "MYSQL_8_0"]
	DatabaseVersion *string `json:"databaseVersion" tf:"database_version"`
	// The externally accessible IPv4 address for the source database server.
	Host *string `json:"host" tf:"host"`
	// The name of the source representation instance. Use any valid Cloud SQL instance name.
	Name *string `json:"name" tf:"name"`
	// The externally accessible port for the source database server.
	// Defaults to 3306.
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The Region in which the created instance should reside.
	// If it is not provided, the provider region is used.
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
}

type SourceRepresentationInstanceStatus struct {
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

// SourceRepresentationInstanceList is a list of SourceRepresentationInstances
type SourceRepresentationInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SourceRepresentationInstance CRD objects
	Items []SourceRepresentationInstance `json:"items,omitempty"`
}
