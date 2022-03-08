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

type RunService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RunServiceSpec   `json:"spec,omitempty"`
	Status            RunServiceStatus `json:"status,omitempty"`
}

type RunServiceSpecMetadata struct {
	// Annotations is a key value map stored with a resource that
	// may be set by external tools to store and retrieve arbitrary metadata. More
	// info: http://kubernetes.io/docs/user-guide/annotations
	//
	// **Note**: The Cloud Run API may add additional annotations that were not provided in your config.
	// If terraform plan shows a diff where a server-side annotation is added, you can add it to your config
	// or apply the lifecycle.ignore_changes rule to the metadata.0.annotations field.
	//
	// Cloud Run (fully managed) uses the following annotation keys to configure features on a Service:
	//
	// - 'run.googleapis.com/ingress' sets the [ingress settings](https://cloud.google.com/sdk/gcloud/reference/run/deploy#--ingress)
	//   for the Service. For example, '"run.googleapis.com/ingress" = "all"'.
	// +optional
	Annotations *map[string]string `json:"annotations,omitempty" tf:"annotations"`
	// A sequence number representing a specific generation of the desired state.
	// +optional
	Generation *int64 `json:"generation,omitempty" tf:"generation"`
	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. May match selectors of replication controllers
	// and routes.
	// More info: http://kubernetes.io/docs/user-guide/labels
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// In Cloud Run the namespace must be equal to either the
	// project ID or project number.
	// +optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace"`
	// An opaque value that represents the internal version of this object that
	// can be used by clients to determine when objects have changed. May be used
	// for optimistic concurrency, change detection, and the watch operation on a
	// resource or set of resources. They may only be valid for a
	// particular resource or set of resources.
	//
	// More info:
	// https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
	// +optional
	ResourceVersion *string `json:"resourceVersion,omitempty" tf:"resource_version"`
	// SelfLink is a URL representing this object.
	// +optional
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link"`
	// UID is a unique id generated by the server on successful creation of a resource and is not
	// allowed to change on PUT operations.
	//
	// More info: http://kubernetes.io/docs/user-guide/identifiers#uids
	// +optional
	Uid *string `json:"uid,omitempty" tf:"uid"`
}

type RunServiceSpecStatusConditions struct {
	// Human readable message indicating details about the current status.
	// +optional
	Message *string `json:"message,omitempty" tf:"message"`
	// One-word CamelCase reason for the condition's current status.
	// +optional
	Reason *string `json:"reason,omitempty" tf:"reason"`
	// Status of the condition, one of True, False, Unknown.
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// Type of domain mapping condition.
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type RunServiceSpecStatus struct {
	// Array of observed Service Conditions, indicating the current ready state of the service.
	// +optional
	Conditions []RunServiceSpecStatusConditions `json:"conditions,omitempty" tf:"conditions"`
	// From ConfigurationStatus. LatestCreatedRevisionName is the last revision that was created
	// from this Service's Configuration. It might not be ready yet, for that use
	// LatestReadyRevisionName.
	// +optional
	LatestCreatedRevisionName *string `json:"latestCreatedRevisionName,omitempty" tf:"latest_created_revision_name"`
	// From ConfigurationStatus. LatestReadyRevisionName holds the name of the latest Revision
	// stamped out from this Service's Configuration that has had its "Ready" condition become
	// "True".
	// +optional
	LatestReadyRevisionName *string `json:"latestReadyRevisionName,omitempty" tf:"latest_ready_revision_name"`
	// ObservedGeneration is the 'Generation' of the Route that was last processed by the
	// controller.
	//
	// Clients polling for completed reconciliation should poll until observedGeneration =
	// metadata.generation and the Ready condition's status is True or False.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty" tf:"observed_generation"`
	// From RouteStatus. URL holds the url that will distribute traffic over the provided traffic
	// targets. It generally has the form
	// https://{route-hash}-{project-hash}-{cluster-level-suffix}.a.run.app
	// +optional
	Url *string `json:"url,omitempty" tf:"url"`
}

type RunServiceSpecTemplateMetadata struct {
	// Annotations is a key value map stored with a resource that
	// may be set by external tools to store and retrieve arbitrary metadata. More
	// info: http://kubernetes.io/docs/user-guide/annotations
	//
	// **Note**: The Cloud Run API may add additional annotations that were not provided in your config.
	// If terraform plan shows a diff where a server-side annotation is added, you can add it to your config
	// or apply the lifecycle.ignore_changes rule to the metadata.0.annotations field.
	// +optional
	Annotations *map[string]string `json:"annotations,omitempty" tf:"annotations"`
	// A sequence number representing a specific generation of the desired state.
	// +optional
	Generation *int64 `json:"generation,omitempty" tf:"generation"`
	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. May match selectors of replication controllers
	// and routes.
	// More info: http://kubernetes.io/docs/user-guide/labels
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// Name must be unique within a namespace, within a Cloud Run region.
	// Is required when creating resources. Name is primarily intended
	// for creation idempotence and configuration definition. Cannot be updated.
	// More info: http://kubernetes.io/docs/user-guide/identifiers#names
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// In Cloud Run the namespace must be equal to either the
	// project ID or project number. It will default to the resource's project.
	// +optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace"`
	// An opaque value that represents the internal version of this object that
	// can be used by clients to determine when objects have changed. May be used
	// for optimistic concurrency, change detection, and the watch operation on a
	// resource or set of resources. They may only be valid for a
	// particular resource or set of resources.
	//
	// More info:
	// https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
	// +optional
	ResourceVersion *string `json:"resourceVersion,omitempty" tf:"resource_version"`
	// SelfLink is a URL representing this object.
	// +optional
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link"`
	// UID is a unique id generated by the server on successful creation of a resource and is not
	// allowed to change on PUT operations.
	//
	// More info: http://kubernetes.io/docs/user-guide/identifiers#uids
	// +optional
	Uid *string `json:"uid,omitempty" tf:"uid"`
}

type RunServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef struct {
	// A Cloud Secret Manager secret version. Must be 'latest' for the latest
	// version or an integer for a specific version.
	Key *string `json:"key" tf:"key"`
	// The name of the secret in Cloud Secret Manager. By default, the secret
	// is assumed to be in the same project.
	// If the secret is in another project, you must define an alias.
	// You set the <alias> in this field, and create an annotation with the
	// following structure
	// "run.googleapis.com/secrets" = "<alias>:projects/<project-id|project-number>/secrets/<secret-name>".
	// If multiple alias definitions are needed, they must be separated by
	// commas in the annotation field.
	Name *string `json:"name" tf:"name"`
}

type RunServiceSpecTemplateSpecContainersEnvValueFrom struct {
	// Selects a key (version) of a secret in Secret Manager.
	SecretKeyRef *RunServiceSpecTemplateSpecContainersEnvValueFromSecretKeyRef `json:"secretKeyRef" tf:"secret_key_ref"`
}

type RunServiceSpecTemplateSpecContainersEnv struct {
	// Name of the environment variable.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// Variable references $(VAR_NAME) are expanded
	// using the previous defined environment variables in the container and
	// any route environment variables. If a variable cannot be resolved,
	// the reference in the input string will be unchanged. The $(VAR_NAME)
	// syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped
	// references will never be expanded, regardless of whether the variable
	// exists or not.
	// Defaults to "".
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
	// Source for the environment variable's value. Only supports secret_key_ref.
	// +optional
	ValueFrom *RunServiceSpecTemplateSpecContainersEnvValueFrom `json:"valueFrom,omitempty" tf:"value_from"`
}

type RunServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference struct {
	// Name of the referent.
	// More info:
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `json:"name" tf:"name"`
}

type RunServiceSpecTemplateSpecContainersEnvFromConfigMapRef struct {
	// The ConfigMap to select from.
	// +optional
	LocalObjectReference *RunServiceSpecTemplateSpecContainersEnvFromConfigMapRefLocalObjectReference `json:"localObjectReference,omitempty" tf:"local_object_reference"`
	// Specify whether the ConfigMap must be defined
	// +optional
	Optional *bool `json:"optional,omitempty" tf:"optional"`
}

type RunServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference struct {
	// Name of the referent.
	// More info:
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `json:"name" tf:"name"`
}

type RunServiceSpecTemplateSpecContainersEnvFromSecretRef struct {
	// The Secret to select from.
	// +optional
	LocalObjectReference *RunServiceSpecTemplateSpecContainersEnvFromSecretRefLocalObjectReference `json:"localObjectReference,omitempty" tf:"local_object_reference"`
	// Specify whether the Secret must be defined
	// +optional
	Optional *bool `json:"optional,omitempty" tf:"optional"`
}

type RunServiceSpecTemplateSpecContainersEnvFrom struct {
	// The ConfigMap to select from.
	// +optional
	ConfigMapRef *RunServiceSpecTemplateSpecContainersEnvFromConfigMapRef `json:"configMapRef,omitempty" tf:"config_map_ref"`
	// An optional identifier to prepend to each key in the ConfigMap.
	// +optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
	// The Secret to select from.
	// +optional
	SecretRef *RunServiceSpecTemplateSpecContainersEnvFromSecretRef `json:"secretRef,omitempty" tf:"secret_ref"`
}

type RunServiceSpecTemplateSpecContainersPorts struct {
	// Port number the container listens on. This must be a valid port number, 0 < x < 65536.
	// +optional
	ContainerPort *int64 `json:"containerPort,omitempty" tf:"container_port"`
	// If specified, used to specify which protocol to use. Allowed values are "http1" and "h2c".
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// Protocol for port. Must be "TCP". Defaults to "TCP".
	// +optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`
}

type RunServiceSpecTemplateSpecContainersResources struct {
	// Limits describes the maximum amount of compute resources allowed.
	// The values of the map is string form of the 'quantity' k8s type:
	// https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go
	// +optional
	Limits *map[string]string `json:"limits,omitempty" tf:"limits"`
	// Requests describes the minimum amount of compute resources required.
	// If Requests is omitted for a container, it defaults to Limits if that is
	// explicitly specified, otherwise to an implementation-defined value.
	// The values of the map is string form of the 'quantity' k8s type:
	// https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go
	// +optional
	Requests *map[string]string `json:"requests,omitempty" tf:"requests"`
}

type RunServiceSpecTemplateSpecContainersVolumeMounts struct {
	// Path within the container at which the volume should be mounted.  Must
	// not contain ':'.
	MountPath *string `json:"mountPath" tf:"mount_path"`
	// This must match the Name of a Volume.
	Name *string `json:"name" tf:"name"`
}

type RunServiceSpecTemplateSpecContainers struct {
	// Arguments to the entrypoint.
	// The docker image's CMD is used if this is not provided.
	// Variable references $(VAR_NAME) are expanded using the container's
	// environment. If a variable cannot be resolved, the reference in the input
	// string will be unchanged. The $(VAR_NAME) syntax can be escaped with a
	// double $$, ie: $$(VAR_NAME). Escaped references will never be expanded,
	// regardless of whether the variable exists or not.
	// More info:
	// https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	// +optional
	Args []string `json:"args,omitempty" tf:"args"`
	// Entrypoint array. Not executed within a shell.
	// The docker image's ENTRYPOINT is used if this is not provided.
	// Variable references $(VAR_NAME) are expanded using the container's
	// environment. If a variable cannot be resolved, the reference in the input
	// string will be unchanged. The $(VAR_NAME) syntax can be escaped with a
	// double $$, ie: $$(VAR_NAME). Escaped references will never be expanded,
	// regardless of whether the variable exists or not.
	// More info:
	// https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	// +optional
	Command []string `json:"command,omitempty" tf:"command"`
	// List of environment variables to set in the container.
	// +optional
	Env []RunServiceSpecTemplateSpecContainersEnv `json:"env,omitempty" tf:"env"`
	// List of sources to populate environment variables in the container.
	// All invalid keys will be reported as an event when the container is starting.
	// When a key exists in multiple sources, the value associated with the last source will
	// take precedence. Values defined by an Env with a duplicate key will take
	// precedence.
	// +optional
	// Deprecated
	EnvFrom []RunServiceSpecTemplateSpecContainersEnvFrom `json:"envFrom,omitempty" tf:"env_from"`
	// Docker image name. This is most often a reference to a container located
	// in the container registry, such as gcr.io/cloudrun/hello
	// More info: https://kubernetes.io/docs/concepts/containers/images
	Image *string `json:"image" tf:"image"`
	// List of open ports in the container.
	// More Info:
	// https://cloud.google.com/run/docs/reference/rest/v1/RevisionSpec#ContainerPort
	// +optional
	Ports []RunServiceSpecTemplateSpecContainersPorts `json:"ports,omitempty" tf:"ports"`
	// Compute Resources required by this container. Used to set values such as max memory
	// More info:
	// https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#requests-and-limits
	// +optional
	Resources *RunServiceSpecTemplateSpecContainersResources `json:"resources,omitempty" tf:"resources"`
	// Volume to mount into the container's filesystem.
	// Only supports SecretVolumeSources.
	// +optional
	VolumeMounts []RunServiceSpecTemplateSpecContainersVolumeMounts `json:"volumeMounts,omitempty" tf:"volume_mounts"`
	// Container's working directory.
	// If not specified, the container runtime's default will be used, which
	// might be configured in the container image.
	// +optional
	// Deprecated
	WorkingDir *string `json:"workingDir,omitempty" tf:"working_dir"`
}

type RunServiceSpecTemplateSpecVolumesSecretItems struct {
	// The Cloud Secret Manager secret version.
	// Can be 'latest' for the latest value or an integer for a specific version.
	Key *string `json:"key" tf:"key"`
	// Mode bits to use on this file, must be a value between 0000 and 0777. If
	// not specified, the volume defaultMode will be used. This might be in
	// conflict with other options that affect the file mode, like fsGroup, and
	// the result can be other mode bits set.
	// +optional
	Mode *int64 `json:"mode,omitempty" tf:"mode"`
	// The relative path of the file to map the key to.
	// May not be an absolute path.
	// May not contain the path element '..'.
	// May not start with the string '..'.
	Path *string `json:"path" tf:"path"`
}

type RunServiceSpecTemplateSpecVolumesSecret struct {
	// Mode bits to use on created files by default. Must be a value between 0000
	// and 0777. Defaults to 0644. Directories within the path are not affected by
	// this setting. This might be in conflict with other options that affect the
	// file mode, like fsGroup, and the result can be other mode bits set.
	// +optional
	DefaultMode *int64 `json:"defaultMode,omitempty" tf:"default_mode"`
	// If unspecified, the volume will expose a file whose name is the
	// secret_name.
	// If specified, the key will be used as the version to fetch from Cloud
	// Secret Manager and the path will be the name of the file exposed in the
	// volume. When items are defined, they must specify a key and a path.
	// +optional
	Items []RunServiceSpecTemplateSpecVolumesSecretItems `json:"items,omitempty" tf:"items"`
	// The name of the secret in Cloud Secret Manager. By default, the secret
	// is assumed to be in the same project.
	// If the secret is in another project, you must define an alias.
	// An alias definition has the form:
	// <alias>:projects/<project-id|project-number>/secrets/<secret-name>.
	// If multiple alias definitions are needed, they must be separated by
	// commas.
	// The alias definitions must be set on the run.googleapis.com/secrets
	// annotation.
	SecretName *string `json:"secretName" tf:"secret_name"`
}

type RunServiceSpecTemplateSpecVolumes struct {
	// Volume's name.
	Name *string `json:"name" tf:"name"`
	// The secret's value will be presented as the content of a file whose
	// name is defined in the item path. If no items are defined, the name of
	// the file is the secret_name.
	Secret *RunServiceSpecTemplateSpecVolumesSecret `json:"secret" tf:"secret"`
}

type RunServiceSpecTemplateSpec struct {
	// ContainerConcurrency specifies the maximum allowed in-flight (concurrent)
	// requests per container of the Revision. Values are:
	// - '0' thread-safe, the system should manage the max concurrency. This is
	//     the default value.
	// - '1' not-thread-safe. Single concurrency
	// - '2-N' thread-safe, max concurrency of N
	// +optional
	ContainerConcurrency *int64 `json:"containerConcurrency,omitempty" tf:"container_concurrency"`
	// Container defines the unit of execution for this Revision.
	// In the context of a Revision, we disallow a number of the fields of
	// this Container, including: name, ports, and volumeMounts.
	// The runtime contract is documented here:
	// https://github.com/knative/serving/blob/master/docs/runtime-contract.md
	// +optional
	Containers []RunServiceSpecTemplateSpecContainers `json:"containers,omitempty" tf:"containers"`
	// Email address of the IAM service account associated with the revision of the
	// service. The service account represents the identity of the running revision,
	// and determines what permissions the revision has. If not provided, the revision
	// will use the project's default service account.
	// +optional
	ServiceAccountName *string `json:"serviceAccountName,omitempty" tf:"service_account_name"`
	// ServingState holds a value describing the state the resources
	// are in for this Revision.
	// It is expected
	// that the system will manipulate this based on routability and load.
	// +optional
	// Deprecated
	ServingState *string `json:"servingState,omitempty" tf:"serving_state"`
	// TimeoutSeconds holds the max duration the instance is allowed for responding to a request.
	// +optional
	TimeoutSeconds *int64 `json:"timeoutSeconds,omitempty" tf:"timeout_seconds"`
	// Volume represents a named volume in a container.
	// +optional
	Volumes []RunServiceSpecTemplateSpecVolumes `json:"volumes,omitempty" tf:"volumes"`
}

type RunServiceSpecTemplate struct {
	// Optional metadata for this Revision, including labels and annotations.
	// Name will be generated by the Configuration. To set minimum instances
	// for this revision, use the "autoscaling.knative.dev/minScale" annotation
	// key. To set maximum instances for this revision, use the
	// "autoscaling.knative.dev/maxScale" annotation key. To set Cloud SQL
	// connections for the revision, use the "run.googleapis.com/cloudsql-instances"
	// annotation key.
	// +optional
	Metadata *RunServiceSpecTemplateMetadata `json:"metadata,omitempty" tf:"metadata"`
	// RevisionSpec holds the desired state of the Revision (from the client).
	// +optional
	Spec *RunServiceSpecTemplateSpec `json:"spec,omitempty" tf:"spec"`
}

type RunServiceSpecTraffic struct {
	// LatestRevision may be optionally provided to indicate that the latest ready
	// Revision of the Configuration should be used for this traffic target. When
	// provided LatestRevision must be true if RevisionName is empty; it must be
	// false when RevisionName is non-empty.
	// +optional
	LatestRevision *bool `json:"latestRevision,omitempty" tf:"latest_revision"`
	// Percent specifies percent of the traffic to this Revision or Configuration.
	Percent *int64 `json:"percent" tf:"percent"`
	// RevisionName of a specific revision to which to send this portion of traffic.
	// +optional
	RevisionName *string `json:"revisionName,omitempty" tf:"revision_name"`
}

type RunServiceSpec struct {
	State *RunServiceSpecResource `json:"state,omitempty" tf:"-"`

	Resource RunServiceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type RunServiceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AutogenerateRevisionName *bool `json:"autogenerateRevisionName,omitempty" tf:"autogenerate_revision_name"`
	// The location of the cloud run instance. eg us-central1
	Location *string `json:"location" tf:"location"`
	// Metadata associated with this Service, including name, namespace, labels,
	// and annotations.
	// +optional
	Metadata *RunServiceSpecMetadata `json:"metadata,omitempty" tf:"metadata"`
	// Name must be unique within a namespace, within a Cloud Run region.
	// Is required when creating resources. Name is primarily intended
	// for creation idempotence and configuration definition. Cannot be updated.
	// More info: http://kubernetes.io/docs/user-guide/identifiers#names
	Name *string `json:"name" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The current status of the Service.
	// +optional
	Status []RunServiceSpecStatus `json:"status,omitempty" tf:"status"`
	// template holds the latest specification for the Revision to
	// be stamped out. The template references the container image, and may also
	// include labels and annotations that should be attached to the Revision.
	// To correlate a Revision, and/or to force a Revision to be created when the
	// spec doesn't otherwise change, a nonce label may be provided in the
	// template metadata. For more details, see:
	// https://github.com/knative/serving/blob/master/docs/client-conventions.md#associate-modifications-with-revisions
	//
	// Cloud Run does not currently support referencing a build that is
	// responsible for materializing the container image from source.
	// +optional
	Template *RunServiceSpecTemplate `json:"template,omitempty" tf:"template"`
	// Traffic specifies how to distribute traffic over a collection of Knative Revisions
	// and Configurations
	// +optional
	Traffic []RunServiceSpecTraffic `json:"traffic,omitempty" tf:"traffic"`
}

type RunServiceStatus struct {
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

// RunServiceList is a list of RunServices
type RunServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RunService CRD objects
	Items []RunService `json:"items,omitempty"`
}
