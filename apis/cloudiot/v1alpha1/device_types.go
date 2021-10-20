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

type Device struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeviceSpec   `json:"spec,omitempty"`
	Status            DeviceStatus `json:"status,omitempty"`
}

type DeviceSpecConfig struct {
	// The device configuration data.
	// +optional
	BinaryData *string `json:"binaryData,omitempty" tf:"binary_data"`
	// The time at which this configuration version was updated in Cloud IoT Core.
	// +optional
	CloudUpdateTime *string `json:"cloudUpdateTime,omitempty" tf:"cloud_update_time"`
	// The time at which Cloud IoT Core received the acknowledgment from the device,
	// indicating that the device has received this configuration version.
	// +optional
	DeviceACKTime *string `json:"deviceACKTime,omitempty" tf:"device_ack_time"`
	// The version of this update.
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
}

type DeviceSpecCredentialsPublicKey struct {
	// The format of the key. Possible values: ["RSA_PEM", "RSA_X509_PEM", "ES256_PEM", "ES256_X509_PEM"]
	Format *string `json:"format" tf:"format"`
	// The key data.
	Key *string `json:"key" tf:"key"`
}

type DeviceSpecCredentials struct {
	// The time at which this credential becomes invalid.
	// +optional
	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time"`
	// A public key used to verify the signature of JSON Web Tokens (JWTs).
	PublicKey *DeviceSpecCredentialsPublicKey `json:"publicKey" tf:"public_key"`
}

type DeviceSpecGatewayConfig struct {
	// Indicates whether the device is a gateway. Possible values: ["ASSOCIATION_ONLY", "DEVICE_AUTH_TOKEN_ONLY", "ASSOCIATION_AND_DEVICE_AUTH_TOKEN"]
	// +optional
	GatewayAuthMethod *string `json:"gatewayAuthMethod,omitempty" tf:"gateway_auth_method"`
	// Indicates whether the device is a gateway. Default value: "NON_GATEWAY" Possible values: ["GATEWAY", "NON_GATEWAY"]
	// +optional
	GatewayType *string `json:"gatewayType,omitempty" tf:"gateway_type"`
	// The ID of the gateway the device accessed most recently.
	// +optional
	LastAccessedGatewayID *string `json:"lastAccessedGatewayID,omitempty" tf:"last_accessed_gateway_id"`
	// The most recent time at which the device accessed the gateway specified in last_accessed_gateway.
	// +optional
	LastAccessedGatewayTime *string `json:"lastAccessedGatewayTime,omitempty" tf:"last_accessed_gateway_time"`
}

type DeviceSpecLastErrorStatus struct {
	// A list of messages that carry the error details.
	// +optional
	// A developer-facing error message, which should be in English.
	// +optional
	Message *string `json:"message,omitempty" tf:"message"`
	// The status code, which should be an enum value of google.rpc.Code.
	// +optional
	Number *int64 `json:"number,omitempty" tf:"number"`
}

type DeviceSpecState struct {
	// The device state data.
	// +optional
	BinaryData *string `json:"binaryData,omitempty" tf:"binary_data"`
	// The time at which this state version was updated in Cloud IoT Core.
	// +optional
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time"`
}

type DeviceSpec struct {
	State *DeviceSpecResource `json:"state,omitempty" tf:"-"`

	Resource DeviceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type DeviceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// If a device is blocked, connections or requests from this device will fail.
	// +optional
	Blocked *bool `json:"blocked,omitempty" tf:"blocked"`
	// The most recent device configuration, which is eventually sent from Cloud IoT Core to the device.
	// +optional
	Config []DeviceSpecConfig `json:"config,omitempty" tf:"config"`
	// The credentials used to authenticate this device.
	// +optional
	// +kubebuilder:validation:MaxItems=3
	Credentials []DeviceSpecCredentials `json:"credentials,omitempty" tf:"credentials"`
	// Gateway-related configuration and state.
	// +optional
	GatewayConfig *DeviceSpecGatewayConfig `json:"gatewayConfig,omitempty" tf:"gateway_config"`
	// The last time a cloud-to-device config version acknowledgment was received from the device.
	// +optional
	LastConfigACKTime *string `json:"lastConfigACKTime,omitempty" tf:"last_config_ack_time"`
	// The last time a cloud-to-device config version was sent to the device.
	// +optional
	LastConfigSendTime *string `json:"lastConfigSendTime,omitempty" tf:"last_config_send_time"`
	// The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub.
	// +optional
	LastErrorStatus []DeviceSpecLastErrorStatus `json:"lastErrorStatus,omitempty" tf:"last_error_status"`
	// The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub.
	// +optional
	LastErrorTime *string `json:"lastErrorTime,omitempty" tf:"last_error_time"`
	// The last time a telemetry event was received.
	// +optional
	LastEventTime *string `json:"lastEventTime,omitempty" tf:"last_event_time"`
	// The last time an MQTT PINGREQ was received.
	// +optional
	LastHeartbeatTime *string `json:"lastHeartbeatTime,omitempty" tf:"last_heartbeat_time"`
	// The last time a state event was received.
	// +optional
	LastStateTime *string `json:"lastStateTime,omitempty" tf:"last_state_time"`
	// The logging verbosity for device activity. Possible values: ["NONE", "ERROR", "INFO", "DEBUG"]
	// +optional
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level"`
	// The metadata key-value pairs assigned to the device.
	// +optional
	Metadata *map[string]string `json:"metadata,omitempty" tf:"metadata"`
	// A unique name for the resource.
	Name *string `json:"name" tf:"name"`
	// A server-defined unique numeric ID for the device.
	// This is a more compact way to identify devices, and it is globally unique.
	// +optional
	NumID *string `json:"numID,omitempty" tf:"num_id"`
	// The name of the device registry where this device should be created.
	Registry *string `json:"registry" tf:"registry"`
	// The state most recently received from the device.
	// +optional
	State []DeviceSpecState `json:"state,omitempty" tf:"state"`
}

type DeviceStatus struct {
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

// DeviceList is a list of Devices
type DeviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Device CRD objects
	Items []Device `json:"items,omitempty"`
}
