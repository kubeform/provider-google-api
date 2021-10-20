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

type AiDataset struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AiDatasetSpec   `json:"spec,omitempty"`
	Status            AiDatasetStatus `json:"status,omitempty"`
}

type AiDatasetSpecEncryptionSpec struct {
	// Required. The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource.
	// Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the resource is created.
	// +optional
	KmsKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name"`
}

type AiDatasetSpec struct {
	State *AiDatasetSpecResource `json:"state,omitempty" tf:"-"`

	Resource AiDatasetSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type AiDatasetSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The timestamp of when the workflow was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	// +optional
	CreateTime *string `json:"createTime,omitempty" tf:"create_time"`
	// The user-defined name of the Dataset. The name can be up to 128 characters long and can be consist of any UTF-8 characters.
	DisplayName *string `json:"displayName" tf:"display_name"`
	// Customer-managed encryption key spec for a Dataset. If set, this Dataset and all sub-resources of this Dataset will be secured by this key.
	// +optional
	EncryptionSpec *AiDatasetSpecEncryptionSpec `json:"encryptionSpec,omitempty" tf:"encryption_spec"`
	// A set of key/value label pairs to assign to this Workflow.
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// Points to a YAML file stored on Google Cloud Storage describing additional information about the Dataset. The schema is defined as an OpenAPI 3.0.2 Schema Object. The schema files that can be used here are found in gs://google-cloud-aiplatform/schema/dataset/metadata/.
	MetadataSchemaURI *string `json:"metadataSchemaURI" tf:"metadata_schema_uri"`
	// The resource name of the Dataset. This value is set by Google.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The region of the dataset. eg us-central1
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// The timestamp of when the workflow was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	// +optional
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time"`
}

type AiDatasetStatus struct {
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

// AiDatasetList is a list of AiDatasets
type AiDatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AiDataset CRD objects
	Items []AiDataset `json:"items,omitempty"`
}
