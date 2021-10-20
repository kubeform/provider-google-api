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

type Job struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              JobSpec   `json:"spec,omitempty"`
	Status            JobStatus `json:"status,omitempty"`
}

type JobSpec struct {
	State *JobSpecResource `json:"state,omitempty" tf:"-"`

	Resource JobSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type JobSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// List of experiments that should be used by the job. An example value is ["enable_stackdriver_agent_metrics"].
	// +optional
	AdditionalExperiments []string `json:"additionalExperiments,omitempty" tf:"additional_experiments"`
	// Indicates if the job should use the streaming engine feature.
	// +optional
	EnableStreamingEngine *bool `json:"enableStreamingEngine,omitempty" tf:"enable_streaming_engine"`
	// The configuration for VM IPs. Options are "WORKER_IP_PUBLIC" or "WORKER_IP_PRIVATE".
	// +optional
	IpConfiguration *string `json:"ipConfiguration,omitempty" tf:"ip_configuration"`
	// The unique ID of this job.
	// +optional
	JobID *string `json:"jobID,omitempty" tf:"job_id"`
	// The name for the Cloud KMS key for the job. Key format is: projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY
	// +optional
	KmsKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name"`
	// User labels to be specified for the job. Keys and values should follow the restrictions specified in the labeling restrictions page. NOTE: Google-provided Dataflow templates often provide default labels that begin with goog-dataflow-provided. Unless explicitly set in config, these labels will be ignored to prevent diffs on re-apply.
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels"`
	// The machine type to use for the job.
	// +optional
	MachineType *string `json:"machineType,omitempty" tf:"machine_type"`
	// The number of workers permitted to work on the job. More workers may improve processing speed at additional cost.
	// +optional
	MaxWorkers *int64 `json:"maxWorkers,omitempty" tf:"max_workers"`
	// A unique name for the resource, required by Dataflow.
	Name *string `json:"name" tf:"name"`
	// The network to which VMs will be assigned. If it is not provided, "default" will be used.
	// +optional
	Network *string `json:"network,omitempty" tf:"network"`
	// One of "drain" or "cancel". Specifies behavior of deletion during terraform destroy.
	// +optional
	OnDelete *string `json:"onDelete,omitempty" tf:"on_delete"`
	// Key/Value pairs to be passed to the Dataflow job (as used in the template).
	// +optional
	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters"`
	// The project in which the resource belongs.
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The region in which the created job should run.
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// The Service Account email used to create the job.
	// +optional
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty" tf:"service_account_email"`
	// The current state of the resource, selected from the JobState enum.
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
	// +optional
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork"`
	// A writeable location on Google Cloud Storage for the Dataflow job to dump its temporary data.
	TempGcsLocation *string `json:"tempGcsLocation" tf:"temp_gcs_location"`
	// The Google Cloud Storage path to the Dataflow job template.
	TemplateGcsPath *string `json:"templateGcsPath" tf:"template_gcs_path"`
	// Only applicable when updating a pipeline. Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job.
	// +optional
	TransformNameMapping map[string]string `json:"transformNameMapping,omitempty" tf:"transform_name_mapping"`
	// The type of this job, selected from the JobType enum.
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// The zone in which the created job should run. If it is not provided, the provider zone is used.
	// +optional
	Zone *string `json:"zone,omitempty" tf:"zone"`
}

type JobStatus struct {
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

// JobList is a list of Jobs
type JobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Job CRD objects
	Items []Job `json:"items,omitempty"`
}
