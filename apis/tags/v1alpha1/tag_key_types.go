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

type TagKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TagKeySpec   `json:"spec,omitempty"`
	Status            TagKeyStatus `json:"status,omitempty"`
}

type TagKeySpec struct {
	State *TagKeySpecResource `json:"state,omitempty" tf:"-"`

	Resource TagKeySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type TagKeySpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Output only. Creation time.
	//
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	// +optional
	CreateTime *string `json:"createTime,omitempty" tf:"create_time"`
	// User-assigned description of the TagKey. Must not exceed 256 characters.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// The generated numeric id for the TagKey.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// Output only. Namespaced name of the TagKey.
	// +optional
	NamespacedName *string `json:"namespacedName,omitempty" tf:"namespaced_name"`
	// Input only. The resource name of the new TagKey's parent. Must be of the form organizations/{org_id}.
	Parent *string `json:"parent" tf:"parent"`
	// Input only. The user friendly name for a TagKey. The short name should be unique for TagKeys within the same tag namespace.
	//
	// The short name must be 1-63 characters, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
	ShortName *string `json:"shortName" tf:"short_name"`
	// Output only. Update time.
	//
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	// +optional
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time"`
}

type TagKeyStatus struct {
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

// TagKeyList is a list of TagKeys
type TagKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of TagKey CRD objects
	Items []TagKey `json:"items,omitempty"`
}
