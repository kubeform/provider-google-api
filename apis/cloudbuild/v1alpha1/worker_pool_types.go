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

type WorkerPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkerPoolSpec   `json:"spec,omitempty"`
	Status            WorkerPoolStatus `json:"status,omitempty"`
}

type WorkerPoolSpecNetworkConfig struct {
	// Required. Immutable. The network definition that the workers are peered to. If this section is left empty, the workers will be peered to `WorkerPool.project_id` on the service producer network. Must be in the format `projects/{project}/global/networks/{network}`, where `{project}` is a project number, such as `12345`, and `{network}` is the name of a VPC network in the project. See [Understanding network configuration options](https://cloud.google.com/cloud-build/docs/custom-workers/set-up-custom-worker-pool-environment#understanding_the_network_configuration_options)
	PeeredNetwork *string `json:"peeredNetwork" tf:"peered_network"`
}

type WorkerPoolSpecWorkerConfig struct {
	// Size of the disk attached to the worker, in GB. See [Worker pool config file](https://cloud.google.com/cloud-build/docs/custom-workers/worker-pool-config-file). Specify a value of up to 1000. If `0` is specified, Cloud Build will use a standard disk size.
	// +optional
	DiskSizeGb *int64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb"`
	// Machine type of a worker, such as `n1-standard-1`. See [Worker pool config file](https://cloud.google.com/cloud-build/docs/custom-workers/worker-pool-config-file). If left blank, Cloud Build will use `n1-standard-1`.
	// +optional
	MachineType *string `json:"machineType,omitempty" tf:"machine_type"`
	// If true, workers are created without any public address, which prevents network egress to public IPs.
	// +optional
	NoExternalIP *bool `json:"noExternalIP,omitempty" tf:"no_external_ip"`
}

type WorkerPoolSpec struct {
	State *WorkerPoolSpecResource `json:"state,omitempty" tf:"-"`

	Resource WorkerPoolSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type WorkerPoolSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// User specified annotations. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
	// +optional
	Annotations *map[string]string `json:"annotations,omitempty" tf:"annotations"`
	// Output only. Time at which the request to create the `WorkerPool` was received.
	// +optional
	CreateTime *string `json:"createTime,omitempty" tf:"create_time"`
	// Output only. Time at which the request to delete the `WorkerPool` was received.
	// +optional
	DeleteTime *string `json:"deleteTime,omitempty" tf:"delete_time"`
	// A user-specified, human-readable name for the `WorkerPool`. If provided, this value must be 1-63 characters.
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// The location for the resource
	Location *string `json:"location" tf:"location"`
	// User-defined name of the `WorkerPool`.
	Name *string `json:"name" tf:"name"`
	// Network configuration for the `WorkerPool`.
	// +optional
	NetworkConfig *WorkerPoolSpecNetworkConfig `json:"networkConfig,omitempty" tf:"network_config"`
	// The project for the resource
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// Output only. `WorkerPool` state. Possible values: STATE_UNSPECIFIED, PENDING, APPROVED, REJECTED, CANCELLED
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// Output only. A unique identifier for the `WorkerPool`.
	// +optional
	Uid *string `json:"uid,omitempty" tf:"uid"`
	// Output only. Time at which the request to update the `WorkerPool` was received.
	// +optional
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time"`
	// Configuration to be used for a creating workers in the `WorkerPool`.
	// +optional
	WorkerConfig *WorkerPoolSpecWorkerConfig `json:"workerConfig,omitempty" tf:"worker_config"`
}

type WorkerPoolStatus struct {
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

// WorkerPoolList is a list of WorkerPools
type WorkerPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WorkerPool CRD objects
	Items []WorkerPool `json:"items,omitempty"`
}
