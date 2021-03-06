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

type Runtime struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RuntimeSpec   `json:"spec,omitempty"`
	Status            RuntimeStatus `json:"status,omitempty"`
}

type RuntimeSpecAccessConfig struct {
	// The type of access mode this instance. For valid values, see
	// 'https://cloud.google.com/vertex-ai/docs/workbench/reference/
	// rest/v1/projects.locations.runtimes#RuntimeAccessType'.
	// +optional
	AccessType *string `json:"accessType,omitempty" tf:"access_type"`
	// The proxy endpoint that is used to access the runtime.
	// +optional
	ProxyURI *string `json:"proxyURI,omitempty" tf:"proxy_uri"`
	// The owner of this runtime after creation. Format: 'alias@example.com'.
	// Currently supports one owner only.
	// +optional
	RuntimeOwner *string `json:"runtimeOwner,omitempty" tf:"runtime_owner"`
}

type RuntimeSpecMetrics struct {
	// Contains runtime daemon metrics, such as OS and kernels and
	// sessions stats.
	// +optional
	SystemMetrics *map[string]string `json:"systemMetrics,omitempty" tf:"system_metrics"`
}

type RuntimeSpecSoftwareConfig struct {
	// Specify a custom Cloud Storage path where the GPU driver is stored.
	// If not specified, we'll automatically choose from official GPU drivers.
	// +optional
	CustomGpuDriverPath *string `json:"customGpuDriverPath,omitempty" tf:"custom_gpu_driver_path"`
	// Verifies core internal services are running. Default: True.
	// +optional
	EnableHealthMonitoring *bool `json:"enableHealthMonitoring,omitempty" tf:"enable_health_monitoring"`
	// Runtime will automatically shutdown after idle_shutdown_time.
	// Default: True
	// +optional
	IdleShutdown *bool `json:"idleShutdown,omitempty" tf:"idle_shutdown"`
	// Time in minutes to wait before shuting down runtime.
	// Default: 180 minutes
	// +optional
	IdleShutdownTimeout *int64 `json:"idleShutdownTimeout,omitempty" tf:"idle_shutdown_timeout"`
	// Install Nvidia Driver automatically.
	// +optional
	InstallGpuDriver *bool `json:"installGpuDriver,omitempty" tf:"install_gpu_driver"`
	// Cron expression in UTC timezone for schedule instance auto upgrade.
	// Please follow the [cron format](https://en.wikipedia.org/wiki/Cron).
	// +optional
	NotebookUpgradeSchedule *string `json:"notebookUpgradeSchedule,omitempty" tf:"notebook_upgrade_schedule"`
	// Path to a Bash script that automatically runs after a notebook instance
	// fully boots up. The path must be a URL or
	// Cloud Storage path (gs://path-to-file/file-name).
	// +optional
	PostStartupScript *string `json:"postStartupScript,omitempty" tf:"post_startup_script"`
}

type RuntimeSpecVirtualMachineVirtualMachineConfigAcceleratorConfig struct {
	// Count of cores of this accelerator.
	// +optional
	CoreCount *int64 `json:"coreCount,omitempty" tf:"core_count"`
	// Accelerator model. For valid values, see
	// 'https://cloud.google.com/vertex-ai/docs/workbench/reference/
	// rest/v1/projects.locations.runtimes#AcceleratorType'
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type RuntimeSpecVirtualMachineVirtualMachineConfigContainerImages struct {
	// The path to the container image repository.
	// For example: gcr.io/{project_id}/{imageName}
	Repository *string `json:"repository" tf:"repository"`
	// The tag of the container image. If not specified, this defaults to the latest tag.
	// +optional
	Tag *string `json:"tag,omitempty" tf:"tag"`
}

type RuntimeSpecVirtualMachineVirtualMachineConfigDataDiskInitializeParams struct {
	// Provide this property when creating the disk.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Specifies the disk name. If not specified, the default is
	// to use the name of the instance. If the disk with the
	// instance name exists already in the given zone/region, a
	// new name will be automatically generated.
	// +optional
	DiskName *string `json:"diskName,omitempty" tf:"disk_name"`
	// Specifies the size of the disk in base-2 GB. If not
	// specified, the disk will be the same size as the image
	// (usually 10GB). If specified, the size must be equal to
	// or larger than 10GB. Default 100 GB.
	// +optional
	DiskSizeGb *int64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb"`
	// The type of the boot disk attached to this runtime,
	// defaults to standard persistent disk. For valid values,
	// see 'https://cloud.google.com/vertex-ai/docs/workbench/
	// reference/rest/v1/projects.locations.runtimes#disktype'
	// +optional
	DiskType *string `json:"diskType,omitempty" tf:"disk_type"`
	// Labels to apply to this disk. These can be later modified
	// by the disks.setLabels method. This field is only
	// applicable for persistent disks.
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
}

type RuntimeSpecVirtualMachineVirtualMachineConfigDataDisk struct {
	// Optional. Specifies whether the disk will be auto-deleted
	// when the instance is deleted (but not when the disk is
	// detached from the instance).
	// +optional
	AutoDelete *bool `json:"autoDelete,omitempty" tf:"auto_delete"`
	// Optional. Indicates that this is a boot disk. The virtual
	// machine will use the first partition of the disk for its
	// root filesystem.
	// +optional
	Boot *bool `json:"boot,omitempty" tf:"boot"`
	// Optional. Specifies a unique device name of your choice
	// that is reflected into the /dev/disk/by-id/google-* tree
	// of a Linux operating system running within the instance.
	// This name can be used to reference the device for mounting,
	// resizing, and so on, from within the instance.
	// If not specified, the server chooses a default device name
	// to apply to this disk, in the form persistent-disk-x, where
	// x is a number assigned by Google Compute Engine. This field
	// is only applicable for persistent disks.
	// +optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name"`
	// Indicates a list of features to enable on the guest operating
	// system. Applicable only for bootable images. To see a list of
	// available features, read 'https://cloud.google.com/compute/docs/
	// images/create-delete-deprecate-private-images#guest-os-features'
	// options. ''
	// +optional
	GuestOsFeatures []string `json:"guestOsFeatures,omitempty" tf:"guest_os_features"`
	// Output only. A zero-based index to this disk, where 0 is
	// reserved for the boot disk. If you have many disks attached
	// to an instance, each disk would have a unique index number.
	// +optional
	Index *int64 `json:"index,omitempty" tf:"index"`
	// Input only. Specifies the parameters for a new disk that will
	// be created alongside the new instance. Use initialization
	// parameters to create boot disks or local SSDs attached to the
	// new instance. This property is mutually exclusive with the
	// source property; you can only define one or the other, but not
	// both.
	// +optional
	InitializeParams *RuntimeSpecVirtualMachineVirtualMachineConfigDataDiskInitializeParams `json:"initializeParams,omitempty" tf:"initialize_params"`
	// "Specifies the disk interface to use for attaching this disk,
	// which is either SCSI or NVME. The default is SCSI. Persistent
	// disks must always use SCSI and the request will fail if you attempt
	// to attach a persistent disk in any other format than SCSI. Local SSDs
	// can use either NVME or SCSI. For performance characteristics of SCSI
	// over NVMe, see Local SSD performance. Valid values: * NVME * SCSI".
	// +optional
	Interface *string `json:"interface,omitempty" tf:"interface"`
	// Type of the resource. Always compute#attachedDisk for attached
	// disks.
	// +optional
	Kind *string `json:"kind,omitempty" tf:"kind"`
	// Output only. Any valid publicly visible licenses.
	// +optional
	Licenses []string `json:"licenses,omitempty" tf:"licenses"`
	// The mode in which to attach this disk, either READ_WRITE
	// or READ_ONLY. If not specified, the default is to attach
	// the disk in READ_WRITE mode.
	// +optional
	Mode *string `json:"mode,omitempty" tf:"mode"`
	// Specifies a valid partial or full URL to an existing
	// Persistent Disk resource.
	// +optional
	Source *string `json:"source,omitempty" tf:"source"`
	// Specifies the type of the disk, either SCRATCH or PERSISTENT.
	// If not specified, the default is PERSISTENT.
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type RuntimeSpecVirtualMachineVirtualMachineConfigEncryptionConfig struct {
	// The Cloud KMS resource identifier of the customer-managed
	// encryption key used to protect a resource, such as a disks.
	// It has the following format:
	// 'projects/{PROJECT_ID}/locations/{REGION}/keyRings/
	// {KEY_RING_NAME}/cryptoKeys/{KEY_NAME}'
	// +optional
	KmsKey *string `json:"kmsKey,omitempty" tf:"kms_key"`
}

type RuntimeSpecVirtualMachineVirtualMachineConfigShieldedInstanceConfig struct {
	// Defines whether the instance has integrity monitoring enabled.
	// Enables monitoring and attestation of the boot integrity of
	// the instance. The attestation is performed against the
	// integrity policy baseline. This baseline is initially derived
	// from the implicitly trusted boot image when the instance is
	// created. Enabled by default.
	// +optional
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty" tf:"enable_integrity_monitoring"`
	// Defines whether the instance has Secure Boot enabled.Secure
	// Boot helps ensure that the system only runs authentic software
	// by verifying the digital signature of all boot components, and
	// halting the boot process if signature verification fails.
	// Disabled by default.
	// +optional
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty" tf:"enable_secure_boot"`
	// Defines whether the instance has the vTPM enabled. Enabled by
	// default.
	// +optional
	EnableVtpm *bool `json:"enableVtpm,omitempty" tf:"enable_vtpm"`
}

type RuntimeSpecVirtualMachineVirtualMachineConfig struct {
	// The Compute Engine accelerator configuration for this runtime.
	// +optional
	AcceleratorConfig *RuntimeSpecVirtualMachineVirtualMachineConfigAcceleratorConfig `json:"acceleratorConfig,omitempty" tf:"accelerator_config"`
	// Use a list of container images to start the notebook instance.
	// +optional
	ContainerImages []RuntimeSpecVirtualMachineVirtualMachineConfigContainerImages `json:"containerImages,omitempty" tf:"container_images"`
	// Data disk option configuration settings.
	DataDisk *RuntimeSpecVirtualMachineVirtualMachineConfigDataDisk `json:"dataDisk" tf:"data_disk"`
	// Encryption settings for virtual machine data disk.
	// +optional
	EncryptionConfig *RuntimeSpecVirtualMachineVirtualMachineConfigEncryptionConfig `json:"encryptionConfig,omitempty" tf:"encryption_config"`
	// The Compute Engine guest attributes. (see [Project and instance
	// guest attributes](https://cloud.google.com/compute/docs/
	// storing-retrieving-metadata#guest_attributes)).
	// +optional
	GuestAttributes *map[string]string `json:"guestAttributes,omitempty" tf:"guest_attributes"`
	// If true, runtime will only have internal IP addresses. By default,
	// runtimes are not restricted to internal IP addresses, and will
	// have ephemeral external IP addresses assigned to each vm. This
	// 'internal_ip_only' restriction can only be enabled for subnetwork
	// enabled networks, and all dependencies must be configured to be
	// accessible without external IP addresses.
	// +optional
	InternalIPOnly *bool `json:"internalIPOnly,omitempty" tf:"internal_ip_only"`
	// The labels to associate with this runtime. Label **keys** must
	// contain 1 to 63 characters, and must conform to [RFC 1035]
	// (https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be
	// empty, but, if present, must contain 1 to 63 characters, and must
	// conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No
	// more than 32 labels can be associated with a cluster.
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// The Compute Engine machine type used for runtimes.
	MachineType *string `json:"machineType" tf:"machine_type"`
	// The Compute Engine metadata entries to add to virtual machine.
	// (see [Project and instance metadata](https://cloud.google.com
	// /compute/docs/storing-retrieving-metadata#project_and_instance
	// _metadata)).
	// +optional
	Metadata *map[string]string `json:"metadata,omitempty" tf:"metadata"`
	// The Compute Engine network to be used for machine communications.
	// Cannot be specified with subnetwork. If neither 'network' nor
	// 'subnet' is specified, the "default" network of the project is
	// used, if it exists. A full URL or partial URI. Examples:
	//   * 'https://www.googleapis.com/compute/v1/projects/[project_id]/
	//   regions/global/default'
	//   * 'projects/[project_id]/regions/global/default'
	// Runtimes are managed resources inside Google Infrastructure.
	// Runtimes support the following network configurations:
	//   * Google Managed Network (Network & subnet are empty)
	//   * Consumer Project VPC (network & subnet are required). Requires
	//   configuring Private Service Access.
	//   * Shared VPC (network & subnet are required). Requires
	//   configuring Private Service Access.
	// +optional
	Network *string `json:"network,omitempty" tf:"network"`
	// The type of vNIC to be used on this interface. This may be gVNIC
	// or VirtioNet. Possible values: ["UNSPECIFIED_NIC_TYPE", "VIRTIO_NET", "GVNIC"]
	// +optional
	NicType *string `json:"nicType,omitempty" tf:"nic_type"`
	// Shielded VM Instance configuration settings.
	// +optional
	ShieldedInstanceConfig *RuntimeSpecVirtualMachineVirtualMachineConfigShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty" tf:"shielded_instance_config"`
	// The Compute Engine subnetwork to be used for machine
	// communications. Cannot be specified with network. A full URL or
	// partial URI are valid. Examples:
	//   * 'https://www.googleapis.com/compute/v1/projects/[project_id]/
	//   regions/us-east1/subnetworks/sub0'
	//   * 'projects/[project_id]/regions/us-east1/subnetworks/sub0'
	// +optional
	Subnet *string `json:"subnet,omitempty" tf:"subnet"`
	// The Compute Engine tags to add to runtime (see [Tagging instances]
	// (https://cloud.google.com/compute/docs/
	// label-or-tag-resources#tags)).
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// The zone where the virtual machine is located.
	// +optional
	Zone *string `json:"zone,omitempty" tf:"zone"`
}

type RuntimeSpecVirtualMachine struct {
	// The unique identifier of the Managed Compute Engine instance.
	// +optional
	InstanceID *string `json:"instanceID,omitempty" tf:"instance_id"`
	// The user-friendly name of the Managed Compute Engine instance.
	// +optional
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name"`
	// Virtual Machine configuration settings.
	// +optional
	VirtualMachineConfig *RuntimeSpecVirtualMachineVirtualMachineConfig `json:"virtualMachineConfig,omitempty" tf:"virtual_machine_config"`
}

type RuntimeSpec struct {
	State *RuntimeSpecResource `json:"state,omitempty" tf:"-"`

	Resource RuntimeSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type RuntimeSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The config settings for accessing runtime.
	// +optional
	AccessConfig *RuntimeSpecAccessConfig `json:"accessConfig,omitempty" tf:"access_config"`
	// The health state of this runtime. For a list of possible output
	// values, see 'https://cloud.google.com/vertex-ai/docs/workbench/
	// reference/rest/v1/projects.locations.runtimes#healthstate'.
	// +optional
	HealthState *string `json:"healthState,omitempty" tf:"health_state"`
	// A reference to the zone where the machine resides.
	Location *string `json:"location" tf:"location"`
	// Contains Runtime daemon metrics such as Service status and JupyterLab
	// status
	// +optional
	Metrics []RuntimeSpecMetrics `json:"metrics,omitempty" tf:"metrics"`
	// The name specified for the Notebook instance.
	Name *string `json:"name" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The config settings for software inside the runtime.
	// +optional
	SoftwareConfig *RuntimeSpecSoftwareConfig `json:"softwareConfig,omitempty" tf:"software_config"`
	// The state of this runtime.
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// Use a Compute Engine VM image to start the managed notebook instance.
	// +optional
	VirtualMachine *RuntimeSpecVirtualMachine `json:"virtualMachine,omitempty" tf:"virtual_machine"`
}

type RuntimeStatus struct {
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

// RuntimeList is a list of Runtimes
type RuntimeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Runtime CRD objects
	Items []Runtime `json:"items,omitempty"`
}
