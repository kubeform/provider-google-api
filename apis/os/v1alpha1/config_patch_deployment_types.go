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

type ConfigPatchDeployment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigPatchDeploymentSpec   `json:"spec,omitempty"`
	Status            ConfigPatchDeploymentStatus `json:"status,omitempty"`
}

type ConfigPatchDeploymentSpecInstanceFilterGroupLabels struct {
	// Compute Engine instance labels that must be present for a VM instance to be targeted by this filter
	Labels *map[string]string `json:"labels" tf:"labels"`
}

type ConfigPatchDeploymentSpecInstanceFilter struct {
	// Target all VM instances in the project. If true, no other criteria is permitted.
	// +optional
	All *bool `json:"all,omitempty" tf:"all"`
	// Targets VM instances matching ANY of these GroupLabels. This allows targeting of disparate groups of VM instances.
	// +optional
	GroupLabels []ConfigPatchDeploymentSpecInstanceFilterGroupLabels `json:"groupLabels,omitempty" tf:"group_labels"`
	// Targets VMs whose name starts with one of these prefixes. Similar to labels, this is another way to group
	// VMs when targeting configs, for example prefix="prod-".
	// +optional
	InstanceNamePrefixes []string `json:"instanceNamePrefixes,omitempty" tf:"instance_name_prefixes"`
	// Targets any of the VM instances specified. Instances are specified by their URI in the 'form zones/{{zone}}/instances/{{instance_name}}',
	// 'projects/{{project_id}}/zones/{{zone}}/instances/{{instance_name}}', or
	// 'https://www.googleapis.com/compute/v1/projects/{{project_id}}/zones/{{zone}}/instances/{{instance_name}}'
	// +optional
	Instances []string `json:"instances,omitempty" tf:"instances"`
	// Targets VM instances in ANY of these zones. Leave empty to target VM instances in any zone.
	// +optional
	Zones []string `json:"zones,omitempty" tf:"zones"`
}

type ConfigPatchDeploymentSpecOneTimeSchedule struct {
	// The desired patch job execution time. A timestamp in RFC3339 UTC "Zulu" format,
	// accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	ExecuteTime *string `json:"executeTime" tf:"execute_time"`
}

type ConfigPatchDeploymentSpecPatchConfigApt struct {
	// List of packages to exclude from update. These packages will be excluded.
	// +optional
	Excludes []string `json:"excludes,omitempty" tf:"excludes"`
	// An exclusive list of packages to be updated. These are the only packages that will be updated.
	// If these packages are not installed, they will be ignored. This field cannot be specified with
	// any other patch configuration fields.
	// +optional
	ExclusivePackages []string `json:"exclusivePackages,omitempty" tf:"exclusive_packages"`
	// By changing the type to DIST, the patching is performed using apt-get dist-upgrade instead. Possible values: ["DIST", "UPGRADE"]
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type ConfigPatchDeploymentSpecPatchConfigGoo struct {
	// goo update settings. Use this setting to override the default goo patch rules.
	Enabled *bool `json:"enabled" tf:"enabled"`
}

type ConfigPatchDeploymentSpecPatchConfigPostStepLinuxExecStepConfigGcsObject struct {
	// Bucket of the Cloud Storage object.
	Bucket *string `json:"bucket" tf:"bucket"`
	// Generation number of the Cloud Storage object. This is used to ensure that the ExecStep specified by this PatchJob does not change.
	GenerationNumber *string `json:"generationNumber" tf:"generation_number"`
	// Name of the Cloud Storage object.
	Object *string `json:"object" tf:"object"`
}

type ConfigPatchDeploymentSpecPatchConfigPostStepLinuxExecStepConfig struct {
	// Defaults to [0]. A list of possible return values that the execution can return to indicate a success.
	// +optional
	AllowedSuccessCodes []int64 `json:"allowedSuccessCodes,omitempty" tf:"allowed_success_codes"`
	// A Cloud Storage object containing the executable.
	// +optional
	GcsObject *ConfigPatchDeploymentSpecPatchConfigPostStepLinuxExecStepConfigGcsObject `json:"gcsObject,omitempty" tf:"gcs_object"`
	// The script interpreter to use to run the script. If no interpreter is specified the script will
	// be executed directly, which will likely only succeed for scripts with shebang lines. Possible values: ["SHELL", "POWERSHELL"]
	// +optional
	Interpreter *string `json:"interpreter,omitempty" tf:"interpreter"`
	// An absolute path to the executable on the VM.
	// +optional
	LocalPath *string `json:"localPath,omitempty" tf:"local_path"`
}

type ConfigPatchDeploymentSpecPatchConfigPostStepWindowsExecStepConfigGcsObject struct {
	// Bucket of the Cloud Storage object.
	Bucket *string `json:"bucket" tf:"bucket"`
	// Generation number of the Cloud Storage object. This is used to ensure that the ExecStep specified by this PatchJob does not change.
	GenerationNumber *string `json:"generationNumber" tf:"generation_number"`
	// Name of the Cloud Storage object.
	Object *string `json:"object" tf:"object"`
}

type ConfigPatchDeploymentSpecPatchConfigPostStepWindowsExecStepConfig struct {
	// Defaults to [0]. A list of possible return values that the execution can return to indicate a success.
	// +optional
	AllowedSuccessCodes []int64 `json:"allowedSuccessCodes,omitempty" tf:"allowed_success_codes"`
	// A Cloud Storage object containing the executable.
	// +optional
	GcsObject *ConfigPatchDeploymentSpecPatchConfigPostStepWindowsExecStepConfigGcsObject `json:"gcsObject,omitempty" tf:"gcs_object"`
	// The script interpreter to use to run the script. If no interpreter is specified the script will
	// be executed directly, which will likely only succeed for scripts with shebang lines. Possible values: ["SHELL", "POWERSHELL"]
	// +optional
	Interpreter *string `json:"interpreter,omitempty" tf:"interpreter"`
	// An absolute path to the executable on the VM.
	// +optional
	LocalPath *string `json:"localPath,omitempty" tf:"local_path"`
}

type ConfigPatchDeploymentSpecPatchConfigPostStep struct {
	// The ExecStepConfig for all Linux VMs targeted by the PatchJob.
	// +optional
	LinuxExecStepConfig *ConfigPatchDeploymentSpecPatchConfigPostStepLinuxExecStepConfig `json:"linuxExecStepConfig,omitempty" tf:"linux_exec_step_config"`
	// The ExecStepConfig for all Windows VMs targeted by the PatchJob.
	// +optional
	WindowsExecStepConfig *ConfigPatchDeploymentSpecPatchConfigPostStepWindowsExecStepConfig `json:"windowsExecStepConfig,omitempty" tf:"windows_exec_step_config"`
}

type ConfigPatchDeploymentSpecPatchConfigPreStepLinuxExecStepConfigGcsObject struct {
	// Bucket of the Cloud Storage object.
	Bucket *string `json:"bucket" tf:"bucket"`
	// Generation number of the Cloud Storage object. This is used to ensure that the ExecStep specified by this PatchJob does not change.
	GenerationNumber *string `json:"generationNumber" tf:"generation_number"`
	// Name of the Cloud Storage object.
	Object *string `json:"object" tf:"object"`
}

type ConfigPatchDeploymentSpecPatchConfigPreStepLinuxExecStepConfig struct {
	// Defaults to [0]. A list of possible return values that the execution can return to indicate a success.
	// +optional
	AllowedSuccessCodes []int64 `json:"allowedSuccessCodes,omitempty" tf:"allowed_success_codes"`
	// A Cloud Storage object containing the executable.
	// +optional
	GcsObject *ConfigPatchDeploymentSpecPatchConfigPreStepLinuxExecStepConfigGcsObject `json:"gcsObject,omitempty" tf:"gcs_object"`
	// The script interpreter to use to run the script. If no interpreter is specified the script will
	// be executed directly, which will likely only succeed for scripts with shebang lines. Possible values: ["SHELL", "POWERSHELL"]
	// +optional
	Interpreter *string `json:"interpreter,omitempty" tf:"interpreter"`
	// An absolute path to the executable on the VM.
	// +optional
	LocalPath *string `json:"localPath,omitempty" tf:"local_path"`
}

type ConfigPatchDeploymentSpecPatchConfigPreStepWindowsExecStepConfigGcsObject struct {
	// Bucket of the Cloud Storage object.
	Bucket *string `json:"bucket" tf:"bucket"`
	// Generation number of the Cloud Storage object. This is used to ensure that the ExecStep specified by this PatchJob does not change.
	GenerationNumber *string `json:"generationNumber" tf:"generation_number"`
	// Name of the Cloud Storage object.
	Object *string `json:"object" tf:"object"`
}

type ConfigPatchDeploymentSpecPatchConfigPreStepWindowsExecStepConfig struct {
	// Defaults to [0]. A list of possible return values that the execution can return to indicate a success.
	// +optional
	AllowedSuccessCodes []int64 `json:"allowedSuccessCodes,omitempty" tf:"allowed_success_codes"`
	// A Cloud Storage object containing the executable.
	// +optional
	GcsObject *ConfigPatchDeploymentSpecPatchConfigPreStepWindowsExecStepConfigGcsObject `json:"gcsObject,omitempty" tf:"gcs_object"`
	// The script interpreter to use to run the script. If no interpreter is specified the script will
	// be executed directly, which will likely only succeed for scripts with shebang lines. Possible values: ["SHELL", "POWERSHELL"]
	// +optional
	Interpreter *string `json:"interpreter,omitempty" tf:"interpreter"`
	// An absolute path to the executable on the VM.
	// +optional
	LocalPath *string `json:"localPath,omitempty" tf:"local_path"`
}

type ConfigPatchDeploymentSpecPatchConfigPreStep struct {
	// The ExecStepConfig for all Linux VMs targeted by the PatchJob.
	// +optional
	LinuxExecStepConfig *ConfigPatchDeploymentSpecPatchConfigPreStepLinuxExecStepConfig `json:"linuxExecStepConfig,omitempty" tf:"linux_exec_step_config"`
	// The ExecStepConfig for all Windows VMs targeted by the PatchJob.
	// +optional
	WindowsExecStepConfig *ConfigPatchDeploymentSpecPatchConfigPreStepWindowsExecStepConfig `json:"windowsExecStepConfig,omitempty" tf:"windows_exec_step_config"`
}

type ConfigPatchDeploymentSpecPatchConfigWindowsUpdate struct {
	// Only apply updates of these windows update classifications. If empty, all updates are applied. Possible values: ["CRITICAL", "SECURITY", "DEFINITION", "DRIVER", "FEATURE_PACK", "SERVICE_PACK", "TOOL", "UPDATE_ROLLUP", "UPDATE"]
	// +optional
	Classifications []string `json:"classifications,omitempty" tf:"classifications"`
	// List of KBs to exclude from update.
	// +optional
	Excludes []string `json:"excludes,omitempty" tf:"excludes"`
	// An exclusive list of kbs to be updated. These are the only patches that will be updated.
	// This field must not be used with other patch configurations.
	// +optional
	ExclusivePatches []string `json:"exclusivePatches,omitempty" tf:"exclusive_patches"`
}

type ConfigPatchDeploymentSpecPatchConfigYum struct {
	// List of packages to exclude from update. These packages will be excluded.
	// +optional
	Excludes []string `json:"excludes,omitempty" tf:"excludes"`
	// An exclusive list of packages to be updated. These are the only packages that will be updated.
	// If these packages are not installed, they will be ignored. This field cannot be specified with
	// any other patch configuration fields.
	// +optional
	ExclusivePackages []string `json:"exclusivePackages,omitempty" tf:"exclusive_packages"`
	// Will cause patch to run yum update-minimal instead.
	// +optional
	Minimal *bool `json:"minimal,omitempty" tf:"minimal"`
	// Adds the --security flag to yum update. Not supported on all platforms.
	// +optional
	Security *bool `json:"security,omitempty" tf:"security"`
}

type ConfigPatchDeploymentSpecPatchConfigZypper struct {
	// Install only patches with these categories. Common categories include security, recommended, and feature.
	// +optional
	Categories []string `json:"categories,omitempty" tf:"categories"`
	// List of packages to exclude from update.
	// +optional
	Excludes []string `json:"excludes,omitempty" tf:"excludes"`
	// An exclusive list of patches to be updated. These are the only patches that will be installed using 'zypper patch patch:' command.
	// This field must not be used with any other patch configuration fields.
	// +optional
	ExclusivePatches []string `json:"exclusivePatches,omitempty" tf:"exclusive_patches"`
	// Install only patches with these severities. Common severities include critical, important, moderate, and low.
	// +optional
	Severities []string `json:"severities,omitempty" tf:"severities"`
	// Adds the --with-optional flag to zypper patch.
	// +optional
	WithOptional *bool `json:"withOptional,omitempty" tf:"with_optional"`
	// Adds the --with-update flag, to zypper patch.
	// +optional
	WithUpdate *bool `json:"withUpdate,omitempty" tf:"with_update"`
}

type ConfigPatchDeploymentSpecPatchConfig struct {
	// Apt update settings. Use this setting to override the default apt patch rules.
	// +optional
	Apt *ConfigPatchDeploymentSpecPatchConfigApt `json:"apt,omitempty" tf:"apt"`
	// goo update settings. Use this setting to override the default goo patch rules.
	// +optional
	Goo *ConfigPatchDeploymentSpecPatchConfigGoo `json:"goo,omitempty" tf:"goo"`
	// The ExecStep to run after the patch update.
	// +optional
	PostStep *ConfigPatchDeploymentSpecPatchConfigPostStep `json:"postStep,omitempty" tf:"post_step"`
	// The ExecStep to run before the patch update.
	// +optional
	PreStep *ConfigPatchDeploymentSpecPatchConfigPreStep `json:"preStep,omitempty" tf:"pre_step"`
	// Post-patch reboot settings. Possible values: ["DEFAULT", "ALWAYS", "NEVER"]
	// +optional
	RebootConfig *string `json:"rebootConfig,omitempty" tf:"reboot_config"`
	// Windows update settings. Use this setting to override the default Windows patch rules.
	// +optional
	WindowsUpdate *ConfigPatchDeploymentSpecPatchConfigWindowsUpdate `json:"windowsUpdate,omitempty" tf:"windows_update"`
	// Yum update settings. Use this setting to override the default yum patch rules.
	// +optional
	Yum *ConfigPatchDeploymentSpecPatchConfigYum `json:"yum,omitempty" tf:"yum"`
	// zypper update settings. Use this setting to override the default zypper patch rules.
	// +optional
	Zypper *ConfigPatchDeploymentSpecPatchConfigZypper `json:"zypper,omitempty" tf:"zypper"`
}

type ConfigPatchDeploymentSpecRecurringScheduleMonthlyWeekDayOfMonth struct {
	// A day of the week. Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]
	DayOfWeek *string `json:"dayOfWeek" tf:"day_of_week"`
	// Week number in a month. 1-4 indicates the 1st to 4th week of the month. -1 indicates the last week of the month.
	WeekOrdinal *int64 `json:"weekOrdinal" tf:"week_ordinal"`
}

type ConfigPatchDeploymentSpecRecurringScheduleMonthly struct {
	// One day of the month. 1-31 indicates the 1st to the 31st day. -1 indicates the last day of the month.
	// Months without the target day will be skipped. For example, a schedule to run "every month on the 31st"
	// will not run in February, April, June, etc.
	// +optional
	MonthDay *int64 `json:"monthDay,omitempty" tf:"month_day"`
	// Week day in a month.
	// +optional
	WeekDayOfMonth *ConfigPatchDeploymentSpecRecurringScheduleMonthlyWeekDayOfMonth `json:"weekDayOfMonth,omitempty" tf:"week_day_of_month"`
}

type ConfigPatchDeploymentSpecRecurringScheduleTimeOfDay struct {
	// Hours of day in 24 hour format. Should be from 0 to 23.
	// An API may choose to allow the value "24:00:00" for scenarios like business closing time.
	// +optional
	Hours *int64 `json:"hours,omitempty" tf:"hours"`
	// Minutes of hour of day. Must be from 0 to 59.
	// +optional
	Minutes *int64 `json:"minutes,omitempty" tf:"minutes"`
	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	// +optional
	Nanos *int64 `json:"nanos,omitempty" tf:"nanos"`
	// Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.
	// +optional
	Seconds *int64 `json:"seconds,omitempty" tf:"seconds"`
}

type ConfigPatchDeploymentSpecRecurringScheduleTimeZone struct {
	// IANA Time Zone Database time zone, e.g. "America/New_York".
	ID *string `json:"ID" tf:"id"`
	// IANA Time Zone Database version number, e.g. "2019a".
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
}

type ConfigPatchDeploymentSpecRecurringScheduleWeekly struct {
	// IANA Time Zone Database time zone, e.g. "America/New_York". Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]
	DayOfWeek *string `json:"dayOfWeek" tf:"day_of_week"`
}

type ConfigPatchDeploymentSpecRecurringSchedule struct {
	// The end time at which a recurring patch deployment schedule is no longer active.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	// +optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time"`
	// The time the last patch job ran successfully.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	// +optional
	LastExecuteTime *string `json:"lastExecuteTime,omitempty" tf:"last_execute_time"`
	// Schedule with monthly executions.
	// +optional
	Monthly *ConfigPatchDeploymentSpecRecurringScheduleMonthly `json:"monthly,omitempty" tf:"monthly"`
	// The time the next patch job is scheduled to run.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	// +optional
	NextExecuteTime *string `json:"nextExecuteTime,omitempty" tf:"next_execute_time"`
	// The time that the recurring schedule becomes effective. Defaults to createTime of the patch deployment.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	// +optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time"`
	// Time of the day to run a recurring deployment.
	TimeOfDay *ConfigPatchDeploymentSpecRecurringScheduleTimeOfDay `json:"timeOfDay" tf:"time_of_day"`
	// Defines the time zone that timeOfDay is relative to. The rules for daylight saving time are
	// determined by the chosen time zone.
	TimeZone *ConfigPatchDeploymentSpecRecurringScheduleTimeZone `json:"timeZone" tf:"time_zone"`
	// Schedule with weekly executions.
	// +optional
	Weekly *ConfigPatchDeploymentSpecRecurringScheduleWeekly `json:"weekly,omitempty" tf:"weekly"`
}

type ConfigPatchDeploymentSpecRolloutDisruptionBudget struct {
	// Specifies a fixed value.
	// +optional
	Fixed *int64 `json:"fixed,omitempty" tf:"fixed"`
	// Specifies the relative value defined as a percentage, which will be multiplied by a reference value.
	// +optional
	Percentage *int64 `json:"percentage,omitempty" tf:"percentage"`
}

type ConfigPatchDeploymentSpecRollout struct {
	// The maximum number (or percentage) of VMs per zone to disrupt at any given moment. The number of VMs calculated from multiplying the percentage by the total number of VMs in a zone is rounded up.
	// During patching, a VM is considered disrupted from the time the agent is notified to begin until patching has completed. This disruption time includes the time to complete reboot and any post-patch steps.
	// A VM contributes to the disruption budget if its patching operation fails either when applying the patches, running pre or post patch steps, or if it fails to respond with a success notification before timing out. VMs that are not running or do not have an active agent do not count toward this disruption budget.
	// For zone-by-zone rollouts, if the disruption budget in a zone is exceeded, the patch job stops, because continuing to the next zone requires completion of the patch process in the previous zone.
	// For example, if the disruption budget has a fixed value of 10, and 8 VMs fail to patch in the current zone, the patch job continues to patch 2 VMs at a time until the zone is completed. When that zone is completed successfully, patching begins with 10 VMs at a time in the next zone. If 10 VMs in the next zone fail to patch, the patch job stops.
	DisruptionBudget *ConfigPatchDeploymentSpecRolloutDisruptionBudget `json:"disruptionBudget" tf:"disruption_budget"`
	// Mode of the patch rollout. Possible values: ["ZONE_BY_ZONE", "CONCURRENT_ZONES"]
	Mode *string `json:"mode" tf:"mode"`
}

type ConfigPatchDeploymentSpec struct {
	KubeformOutput *ConfigPatchDeploymentSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource ConfigPatchDeploymentSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ConfigPatchDeploymentSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Time the patch deployment was created. Timestamp is in RFC3339 text format.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	// +optional
	CreateTime *string `json:"createTime,omitempty" tf:"create_time"`
	// Description of the patch deployment. Length of the description is limited to 1024 characters.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Duration of the patch. After the duration ends, the patch times out.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
	// +optional
	Duration *string `json:"duration,omitempty" tf:"duration"`
	// VM instances to patch.
	InstanceFilter *ConfigPatchDeploymentSpecInstanceFilter `json:"instanceFilter" tf:"instance_filter"`
	// The last time a patch job was started by this deployment. Timestamp is in RFC3339 text format.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	// +optional
	LastExecuteTime *string `json:"lastExecuteTime,omitempty" tf:"last_execute_time"`
	// Unique name for the patch deployment resource in a project.
	// The patch deployment name is in the form: projects/{project_id}/patchDeployments/{patchDeploymentId}.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// Schedule a one-time execution.
	// +optional
	OneTimeSchedule *ConfigPatchDeploymentSpecOneTimeSchedule `json:"oneTimeSchedule,omitempty" tf:"one_time_schedule"`
	// Patch configuration that is applied.
	// +optional
	PatchConfig *ConfigPatchDeploymentSpecPatchConfig `json:"patchConfig,omitempty" tf:"patch_config"`
	// A name for the patch deployment in the project. When creating a name the following rules apply:
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the project.
	PatchDeploymentID *string `json:"patchDeploymentID" tf:"patch_deployment_id"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// Schedule recurring executions.
	// +optional
	RecurringSchedule *ConfigPatchDeploymentSpecRecurringSchedule `json:"recurringSchedule,omitempty" tf:"recurring_schedule"`
	// Rollout strategy of the patch job.
	// +optional
	Rollout *ConfigPatchDeploymentSpecRollout `json:"rollout,omitempty" tf:"rollout"`
	// Time the patch deployment was last updated. Timestamp is in RFC3339 text format.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	// +optional
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time"`
}

type ConfigPatchDeploymentStatus struct {
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

// ConfigPatchDeploymentList is a list of ConfigPatchDeployments
type ConfigPatchDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ConfigPatchDeployment CRD objects
	Items []ConfigPatchDeployment `json:"items,omitempty"`
}
