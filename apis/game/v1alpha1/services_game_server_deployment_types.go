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

type ServicesGameServerDeployment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicesGameServerDeploymentSpec   `json:"spec,omitempty"`
	Status            ServicesGameServerDeploymentStatus `json:"status,omitempty"`
}

type ServicesGameServerDeploymentSpec struct {
	State *ServicesGameServerDeploymentSpecResource `json:"state,omitempty" tf:"-"`

	Resource ServicesGameServerDeploymentSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ServicesGameServerDeploymentSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// A unique id for the deployment.
	DeploymentID *string `json:"deploymentID" tf:"deployment_id"`
	// Human readable description of the game server deployment.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// The labels associated with this game server deployment. Each label is a
	// key-value pair.
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// Location of the Deployment.
	// +optional
	Location *string `json:"location,omitempty" tf:"location"`
	// The resource id of the game server deployment, eg:
	//
	// 'projects/{project_id}/locations/{location}/gameServerDeployments/{deployment_id}'.
	// For example,
	//
	// 'projects/my-project/locations/{location}/gameServerDeployments/my-deployment'.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
}

type ServicesGameServerDeploymentStatus struct {
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

// ServicesGameServerDeploymentList is a list of ServicesGameServerDeployments
type ServicesGameServerDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServicesGameServerDeployment CRD objects
	Items []ServicesGameServerDeployment `json:"items,omitempty"`
}
