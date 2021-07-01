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

type Trigger struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TriggerSpec   `json:"spec,omitempty"`
	Status            TriggerStatus `json:"status,omitempty"`
}

type TriggerSpecDestinationCloudRunService struct {
	// +optional
	Path *string `json:"path,omitempty" tf:"path"`
	// +optional
	Region  *string `json:"region,omitempty" tf:"region"`
	Service *string `json:"service" tf:"service"`
}

type TriggerSpecDestination struct {
	// +optional
	CloudFunction *string `json:"cloudFunction,omitempty" tf:"cloud_function"`
	// +optional
	CloudRunService *TriggerSpecDestinationCloudRunService `json:"cloudRunService,omitempty" tf:"cloud_run_service"`
}

type TriggerSpecMatchingCriteria struct {
	Attribute *string `json:"attribute" tf:"attribute"`
	Value     *string `json:"value" tf:"value"`
}

type TriggerSpecTransportPubsub struct {
	// +optional
	Subscription *string `json:"subscription,omitempty" tf:"subscription"`
	// +optional
	Topic *string `json:"topic,omitempty" tf:"topic"`
}

type TriggerSpecTransport struct {
	// +optional
	Pubsub *TriggerSpecTransportPubsub `json:"pubsub,omitempty" tf:"pubsub"`
}

type TriggerSpec struct {
	KubeformOutput *TriggerSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource TriggerSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type TriggerSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CreateTime  *string                 `json:"createTime,omitempty" tf:"create_time"`
	Destination *TriggerSpecDestination `json:"destination" tf:"destination"`
	// +optional
	Etag *string `json:"etag,omitempty" tf:"etag"`
	// +optional
	Labels           *map[string]string            `json:"labels,omitempty" tf:"labels"`
	Location         *string                       `json:"location" tf:"location"`
	MatchingCriteria []TriggerSpecMatchingCriteria `json:"matchingCriteria" tf:"matching_criteria"`
	Name             *string                       `json:"name" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// +optional
	ServiceAccount *string `json:"serviceAccount,omitempty" tf:"service_account"`
	// +optional
	Transport *TriggerSpecTransport `json:"transport,omitempty" tf:"transport"`
	// +optional
	Uid *string `json:"uid,omitempty" tf:"uid"`
	// +optional
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time"`
}

type TriggerStatus struct {
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

// TriggerList is a list of Triggers
type TriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Trigger CRD objects
	Items []Trigger `json:"items,omitempty"`
}
