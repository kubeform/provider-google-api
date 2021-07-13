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

type BillingAccountBucketConfig struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BillingAccountBucketConfigSpec   `json:"spec,omitempty"`
	Status            BillingAccountBucketConfigStatus `json:"status,omitempty"`
}

type BillingAccountBucketConfigSpec struct {
	State *BillingAccountBucketConfigSpecResource `json:"state,omitempty" tf:"-"`

	Resource BillingAccountBucketConfigSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type BillingAccountBucketConfigSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The parent resource that contains the logging bucket.
	BillingAccount *string `json:"billingAccount" tf:"billing_account"`
	// The name of the logging bucket. Logging automatically creates two log buckets: _Required and _Default.
	BucketID *string `json:"bucketID" tf:"bucket_id"`
	// An optional description for this bucket.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// The bucket's lifecycle such as active or deleted.
	// +optional
	LifecycleState *string `json:"lifecycleState,omitempty" tf:"lifecycle_state"`
	// The location of the bucket.
	Location *string `json:"location" tf:"location"`
	// The resource name of the bucket
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	// +optional
	RetentionDays *int64 `json:"retentionDays,omitempty" tf:"retention_days"`
}

type BillingAccountBucketConfigStatus struct {
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

// BillingAccountBucketConfigList is a list of BillingAccountBucketConfigs
type BillingAccountBucketConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BillingAccountBucketConfig CRD objects
	Items []BillingAccountBucketConfig `json:"items,omitempty"`
}
