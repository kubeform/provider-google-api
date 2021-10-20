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

type BucketObject struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketObjectSpec   `json:"spec,omitempty"`
	Status            BucketObjectStatus `json:"status,omitempty"`
}

type BucketObjectSpec struct {
	State *BucketObjectSpecResource `json:"state,omitempty" tf:"-"`

	Resource BucketObjectSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type BucketObjectSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the containing bucket.
	Bucket *string `json:"bucket" tf:"bucket"`
	// Cache-Control directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
	// +optional
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control"`
	// Data as string to be uploaded. Must be defined if source is not. Note: The content field is marked as sensitive. To view the raw contents of the object, please define an output.
	// +optional
	Content *string `json:"-" sensitive:"true" tf:"content"`
	// Content-Disposition of the object data.
	// +optional
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition"`
	// Content-Encoding of the object data.
	// +optional
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding"`
	// Content-Language of the object data.
	// +optional
	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language"`
	// Content-Type of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	// +optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type"`
	// Base 64 CRC32 hash of the uploaded data.
	// +optional
	Crc32c *string `json:"crc32c,omitempty" tf:"crc32c"`
	// +optional
	DetectMd5hash *string `json:"detectMd5hash,omitempty" tf:"detect_md5hash"`
	// Resource name of the Cloud KMS key that will be used to encrypt the object. Overrides the object metadata's kmsKeyName value, if any.
	// +optional
	KmsKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name"`
	// Base 64 MD5 hash of the uploaded data.
	// +optional
	Md5hash *string `json:"md5hash,omitempty" tf:"md5hash"`
	// A url reference to download this object.
	// +optional
	MediaLink *string `json:"mediaLink,omitempty" tf:"media_link"`
	// User-provided metadata, in key/value pairs.
	// +optional
	Metadata *map[string]string `json:"metadata,omitempty" tf:"metadata"`
	// The name of the object. If you're interpolating the name of this object, see output_name instead.
	Name *string `json:"name" tf:"name"`
	// The name of the object. Use this field in interpolations with google_storage_object_acl to recreate google_storage_object_acl resources when your google_storage_bucket_object is recreated.
	// +optional
	OutputName *string `json:"outputName,omitempty" tf:"output_name"`
	// A url reference to this object.
	// +optional
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link"`
	// A path to the data you want to upload. Must be defined if content is not.
	// +optional
	Source *string `json:"source,omitempty" tf:"source"`
	// The StorageClass of the new bucket object. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE. If not provided, this defaults to the bucket's default storage class or to a standard class.
	// +optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class"`
}

type BucketObjectStatus struct {
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

// BucketObjectList is a list of BucketObjects
type BucketObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BucketObject CRD objects
	Items []BucketObject `json:"items,omitempty"`
}
