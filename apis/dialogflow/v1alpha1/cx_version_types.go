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

type CxVersion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CxVersionSpec   `json:"spec,omitempty"`
	Status            CxVersionStatus `json:"status,omitempty"`
}

type CxVersionSpecNluSettings struct {
	// To filter out false positive results and still get variety in matched natural language inputs for your agent, you can tune the machine learning classification threshold. If the returned score value is less than the threshold value, then a no-match event will be triggered.
	// The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the default of 0.3 is used.
	// +optional
	ClassificationThreshold *float64 `json:"classificationThreshold,omitempty" tf:"classification_threshold"`
	// Indicates NLU model training mode.
	// * MODEL_TRAINING_MODE_AUTOMATIC: NLU model training is automatically triggered when a flow gets modified. User can also manually trigger model training in this mode.
	// * MODEL_TRAINING_MODE_MANUAL: User needs to manually trigger NLU model training. Best for large flows whose models take long time to train. Possible values: ["MODEL_TRAINING_MODE_AUTOMATIC", "MODEL_TRAINING_MODE_MANUAL"]
	// +optional
	ModelTrainingMode *string `json:"modelTrainingMode,omitempty" tf:"model_training_mode"`
	// Indicates the type of NLU model.
	// * MODEL_TYPE_STANDARD: Use standard NLU model.
	// * MODEL_TYPE_ADVANCED: Use advanced NLU model. Possible values: ["MODEL_TYPE_STANDARD", "MODEL_TYPE_ADVANCED"]
	// +optional
	ModelType *string `json:"modelType,omitempty" tf:"model_type"`
}

type CxVersionSpec struct {
	State *CxVersionSpecResource `json:"state,omitempty" tf:"-"`

	Resource CxVersionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type CxVersionSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	// +optional
	CreateTime *string `json:"createTime,omitempty" tf:"create_time"`
	// The description of the version. The maximum length is 500 characters. If exceeded, the request is rejected.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// The human-readable name of the version. Limit of 64 characters.
	DisplayName *string `json:"displayName" tf:"display_name"`
	// Format: projects//locations//agents//flows//versions/. Version ID is a self-increasing number generated by Dialogflow upon version creation.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The NLU settings of the flow at version creation.
	// +optional
	NluSettings []CxVersionSpecNluSettings `json:"nluSettings,omitempty" tf:"nlu_settings"`
	// The Flow to create an Version for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	// +optional
	Parent *string `json:"parent,omitempty" tf:"parent"`
	// The state of this version.
	// * RUNNING: Version is not ready to serve (e.g. training is running).
	// * SUCCEEDED: Training has succeeded and this version is ready to serve.
	// * FAILED: Version training failed.
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
}

type CxVersionStatus struct {
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

// CxVersionList is a list of CxVersions
type CxVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CxVersion CRD objects
	Items []CxVersion `json:"items,omitempty"`
}
