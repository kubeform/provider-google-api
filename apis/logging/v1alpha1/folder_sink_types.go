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

type FolderSink struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FolderSinkSpec   `json:"spec,omitempty"`
	Status            FolderSinkStatus `json:"status,omitempty"`
}

type FolderSinkSpecBigqueryOptions struct {
	// Whether to use BigQuery's partition tables. By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned tables the date suffix is no longer present and special query syntax has to be used instead. In both cases, tables are sharded based on UTC timezone.
	UsePartitionedTables *bool `json:"usePartitionedTables" tf:"use_partitioned_tables"`
}

type FolderSinkSpecExclusions struct {
	// A description of this exclusion.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// If set to True, then this exclusion is disabled and it does not exclude any log entries
	// +optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled"`
	// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries
	Filter *string `json:"filter" tf:"filter"`
	// A client-assigned identifier, such as "load-balancer-exclusion". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name *string `json:"name" tf:"name"`
}

type FolderSinkSpec struct {
	State *FolderSinkSpecResource `json:"state,omitempty" tf:"-"`

	Resource FolderSinkSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type FolderSinkSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Options that affect sinks exporting data to BigQuery.
	// +optional
	BigqueryOptions *FolderSinkSpecBigqueryOptions `json:"bigqueryOptions,omitempty" tf:"bigquery_options"`
	// A description of this sink. The maximum length of the description is 8000 characters.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// The destination of the sink (or, in other words, where logs are written to). Can be a Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples: "storage.googleapis.com/[GCS_BUCKET]" "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]" "pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]" The writer associated with the sink must have access to write to the above resource.
	Destination *string `json:"destination" tf:"destination"`
	// If set to True, then this sink is disabled and it does not export any log entries.
	// +optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled"`
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusion_filters it will not be exported.
	// +optional
	Exclusions []FolderSinkSpecExclusions `json:"exclusions,omitempty" tf:"exclusions"`
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// +optional
	Filter *string `json:"filter,omitempty" tf:"filter"`
	// The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is accepted.
	Folder *string `json:"folder" tf:"folder"`
	// Whether or not to include children folders in the sink export. If true, logs associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
	// +optional
	IncludeChildren *bool `json:"includeChildren,omitempty" tf:"include_children"`
	// The name of the logging sink.
	Name *string `json:"name" tf:"name"`
	// The identity associated with this sink. This identity must be granted write access to the configured destination.
	// +optional
	WriterIdentity *string `json:"writerIdentity,omitempty" tf:"writer_identity"`
}

type FolderSinkStatus struct {
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

// FolderSinkList is a list of FolderSinks
type FolderSinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FolderSink CRD objects
	Items []FolderSink `json:"items,omitempty"`
}
