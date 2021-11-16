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

type ServicesEdgeCacheOrigin struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicesEdgeCacheOriginSpec   `json:"spec,omitempty"`
	Status            ServicesEdgeCacheOriginStatus `json:"status,omitempty"`
}

type ServicesEdgeCacheOriginSpecTimeout struct {
	// The maximum duration to wait for the origin connection to be established, including DNS lookup, TLS handshake and TCP/QUIC connection establishment.
	//
	// Defaults to 5 seconds. The timeout must be a value between 1s and 15s.
	// +optional
	ConnectTimeout *string `json:"connectTimeout,omitempty" tf:"connect_timeout"`
	// The maximum time across all connection attempts to the origin, including failover origins, before returning an error to the client. A HTTP 503 will be returned if the timeout is reached before a response is returned.
	//
	// Defaults to 5 seconds. The timeout must be a value between 1s and 15s.
	// +optional
	MaxAttemptsTimeout *string `json:"maxAttemptsTimeout,omitempty" tf:"max_attempts_timeout"`
	// The maximum duration to wait for data to arrive when reading from the HTTP connection/stream.
	//
	// Defaults to 5 seconds. The timeout must be a value between 1s and 30s.
	// +optional
	ResponseTimeout *string `json:"responseTimeout,omitempty" tf:"response_timeout"`
}

type ServicesEdgeCacheOriginSpec struct {
	State *ServicesEdgeCacheOriginSpecResource `json:"state,omitempty" tf:"-"`

	Resource ServicesEdgeCacheOriginSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ServicesEdgeCacheOriginSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// A human-readable description of the resource.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// The Origin resource to try when the current origin cannot be reached.
	// After maxAttempts is reached, the configured failoverOrigin will be used to fulfil the request.
	//
	// The value of timeout.maxAttemptsTimeout dictates the timeout across all origins.
	// A reference to a Topic resource.
	// +optional
	FailoverOrigin *string `json:"failoverOrigin,omitempty" tf:"failover_origin"`
	// Set of label tags associated with the EdgeCache resource.
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// The maximum number of attempts to cache fill from this origin. Another attempt is made when a cache fill fails with one of the retryConditions.
	//
	// Once maxAttempts to this origin have failed the failoverOrigin will be used, if one is specified. That failoverOrigin may specify its own maxAttempts,
	// retryConditions and failoverOrigin to control its own cache fill failures.
	//
	// The total number of allowed attempts to cache fill across this and failover origins is limited to four.
	// The total time allowed for cache fill attempts across this and failover origins can be controlled with maxAttemptsTimeout.
	//
	// The last valid response from an origin will be returned to the client.
	// If no origin returns a valid response, an HTTP 503 will be returned to the client.
	//
	// Defaults to 1. Must be a value greater than 0 and less than 4.
	// +optional
	MaxAttempts *int64 `json:"maxAttempts,omitempty" tf:"max_attempts"`
	// Name of the resource; provided by the client when the resource is created.
	// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
	// and all following characters must be a dash, underscore, letter or digit.
	Name *string `json:"name" tf:"name"`
	// A fully qualified domain name (FQDN) or IP address reachable over the public Internet, or the address of a Google Cloud Storage bucket.
	//
	// This address will be used as the origin for cache requests - e.g. FQDN: media-backend.example.com IPv4:35.218.1.1 IPv6:[2607:f8b0:4012:809::200e] Cloud Storage: gs://bucketname
	//
	// When providing an FQDN (hostname), it must be publicly resolvable (e.g. via Google public DNS) and IP addresses must be publicly routable.
	// If a Cloud Storage bucket is provided, it must be in the canonical "gs://bucketname" format. Other forms, such as "storage.googleapis.com", will be rejected.
	OriginAddress *string `json:"originAddress" tf:"origin_address"`
	// The port to connect to the origin on.
	// Defaults to port 443 for HTTP2 and HTTPS protocols, and port 80 for HTTP.
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The protocol to use to connect to the configured origin. Defaults to HTTP2, and it is strongly recommended that users use HTTP2 for both security & performance.
	//
	// When using HTTP2 or HTTPS as the protocol, a valid, publicly-signed, unexpired TLS (SSL) certificate must be presented by the origin server. Possible values: ["HTTP2", "HTTPS", "HTTP"]
	// +optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`
	// Specifies one or more retry conditions for the configured origin.
	//
	// If the failure mode during a connection attempt to the origin matches the configured retryCondition(s),
	// the origin request will be retried up to maxAttempts times. The failoverOrigin, if configured, will then be used to satisfy the request.
	//
	// The default retryCondition is "CONNECT_FAILURE".
	//
	// retryConditions apply to this origin, and not subsequent failoverOrigin(s),
	// which may specify their own retryConditions and maxAttempts.
	//
	// Valid values are:
	//
	// - CONNECT_FAILURE: Retry on failures connecting to origins, for example due to connection timeouts.
	// - HTTP_5XX: Retry if the origin responds with any 5xx response code, or if the origin does not respond at all, example: disconnects, reset, read timeout, connection failure, and refused streams.
	// - GATEWAY_ERROR: Similar to 5xx, but only applies to response codes 502, 503 or 504.
	// - RETRIABLE_4XX: Retry for retriable 4xx response codes, which include HTTP 409 (Conflict) and HTTP 429 (Too Many Requests)
	// - NOT_FOUND: Retry if the origin returns a HTTP 404 (Not Found). This can be useful when generating video content, and the segment is not available yet. Possible values: ["CONNECT_FAILURE", "HTTP_5XX", "GATEWAY_ERROR", "RETRIABLE_4XX", "NOT_FOUND"]
	// +optional
	RetryConditions []string `json:"retryConditions,omitempty" tf:"retry_conditions"`
	// The connection and HTTP timeout configuration for this origin.
	// +optional
	Timeout *ServicesEdgeCacheOriginSpecTimeout `json:"timeout,omitempty" tf:"timeout"`
}

type ServicesEdgeCacheOriginStatus struct {
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

// ServicesEdgeCacheOriginList is a list of ServicesEdgeCacheOrigins
type ServicesEdgeCacheOriginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServicesEdgeCacheOrigin CRD objects
	Items []ServicesEdgeCacheOrigin `json:"items,omitempty"`
}