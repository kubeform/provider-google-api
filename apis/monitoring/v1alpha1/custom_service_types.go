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

type CustomService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CustomServiceSpec   `json:"spec,omitempty"`
	Status            CustomServiceStatus `json:"status,omitempty"`
}

type CustomServiceSpecTelemetry struct {
	// The full name of the resource that defines this service.
	// Formatted as described in
	// https://cloud.google.com/apis/design/resource_names.
	// +optional
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name"`
}

type CustomServiceSpec struct {
	State *CustomServiceSpecResource `json:"state,omitempty" tf:"-"`

	Resource CustomServiceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type CustomServiceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Name used for UI elements listing this Service.
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// The full resource name for this service. The syntax is:
	// projects/[PROJECT_ID]/services/[SERVICE_ID].
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// An optional service ID to use. If not given, the server will generate a
	// service ID.
	// +optional
	ServiceID *string `json:"serviceID,omitempty" tf:"service_id"`
	// Configuration for how to query telemetry on a Service.
	// +optional
	Telemetry *CustomServiceSpecTelemetry `json:"telemetry,omitempty" tf:"telemetry"`
}

type CustomServiceStatus struct {
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

// CustomServiceList is a list of CustomServices
type CustomServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CustomService CRD objects
	Items []CustomService `json:"items,omitempty"`
}
