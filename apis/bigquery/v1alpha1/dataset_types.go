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

type Dataset struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatasetSpec   `json:"spec,omitempty"`
	Status            DatasetStatus `json:"status,omitempty"`
}

type DatasetSpecAccessView struct {
	// The ID of the dataset containing this table.
	DatasetID *string `json:"datasetID" tf:"dataset_id"`
	// The ID of the project containing this table.
	ProjectID *string `json:"projectID" tf:"project_id"`
	// The ID of the table. The ID must contain only letters (a-z,
	// A-Z), numbers (0-9), or underscores (_). The maximum length
	// is 1,024 characters.
	TableID *string `json:"tableID" tf:"table_id"`
}

type DatasetSpecAccess struct {
	// A domain to grant access to. Any users signed in with the
	// domain specified will be granted the specified access
	// +optional
	Domain *string `json:"domain,omitempty" tf:"domain"`
	// An email address of a Google Group to grant access to.
	// +optional
	GroupByEmail *string `json:"groupByEmail,omitempty" tf:"group_by_email"`
	// Describes the rights granted to the user specified by the other
	// member of the access object. Basic, predefined, and custom roles
	// are supported. Predefined roles that have equivalent basic roles
	// are swapped by the API to their basic counterparts. See
	// [official docs](https://cloud.google.com/bigquery/docs/access-control).
	// +optional
	Role *string `json:"role,omitempty" tf:"role"`
	// A special group to grant access to. Possible values include:
	//
	//
	// * 'projectOwners': Owners of the enclosing project.
	//
	//
	// * 'projectReaders': Readers of the enclosing project.
	//
	//
	// * 'projectWriters': Writers of the enclosing project.
	//
	//
	// * 'allAuthenticatedUsers': All authenticated BigQuery users.
	// +optional
	SpecialGroup *string `json:"specialGroup,omitempty" tf:"special_group"`
	// An email address of a user to grant access to. For example:
	// fred@example.com
	// +optional
	UserByEmail *string `json:"userByEmail,omitempty" tf:"user_by_email"`
	// A view from a different dataset to grant access to. Queries
	// executed against that view will have read access to tables in
	// this dataset. The role field is not required when this field is
	// set. If that view is updated by any user, access to the view
	// needs to be granted again via an update operation.
	// +optional
	View *DatasetSpecAccessView `json:"view,omitempty" tf:"view"`
}

type DatasetSpecDefaultEncryptionConfiguration struct {
	// Describes the Cloud KMS encryption key that will be used to protect destination
	// BigQuery table. The BigQuery Service Account associated with your project requires
	// access to this encryption key.
	KmsKeyName *string `json:"kmsKeyName" tf:"kms_key_name"`
}

type DatasetSpec struct {
	State *DatasetSpecResource `json:"state,omitempty" tf:"-"`

	Resource DatasetSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DatasetSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// An array of objects that define dataset access for one or more entities.
	// +optional
	Access []DatasetSpecAccess `json:"access,omitempty" tf:"access"`
	// The time when this dataset was created, in milliseconds since the
	// epoch.
	// +optional
	CreationTime *int64 `json:"creationTime,omitempty" tf:"creation_time"`
	// A unique ID for this dataset, without the project name. The ID
	// must contain only letters (a-z, A-Z), numbers (0-9), or
	// underscores (_). The maximum length is 1,024 characters.
	DatasetID *string `json:"datasetID" tf:"dataset_id"`
	// The default encryption key for all tables in the dataset. Once this property is set,
	// all newly-created partitioned tables in the dataset will have encryption key set to
	// this value, unless table creation request (or query) overrides the key.
	// +optional
	DefaultEncryptionConfiguration *DatasetSpecDefaultEncryptionConfiguration `json:"defaultEncryptionConfiguration,omitempty" tf:"default_encryption_configuration"`
	// The default partition expiration for all partitioned tables in
	// the dataset, in milliseconds.
	//
	//
	// Once this property is set, all newly-created partitioned tables in
	// the dataset will have an 'expirationMs' property in the 'timePartitioning'
	// settings set to this value, and changing the value will only
	// affect new tables, not existing ones. The storage in a partition will
	// have an expiration time of its partition time plus this value.
	// Setting this property overrides the use of 'defaultTableExpirationMs'
	// for partitioned tables: only one of 'defaultTableExpirationMs' and
	// 'defaultPartitionExpirationMs' will be used for any new partitioned
	// table. If you provide an explicit 'timePartitioning.expirationMs' when
	// creating or updating a partitioned table, that value takes precedence
	// over the default partition expiration time indicated by this property.
	// +optional
	DefaultPartitionExpirationMs *int64 `json:"defaultPartitionExpirationMs,omitempty" tf:"default_partition_expiration_ms"`
	// The default lifetime of all tables in the dataset, in milliseconds.
	// The minimum value is 3600000 milliseconds (one hour).
	//
	//
	// Once this property is set, all newly-created tables in the dataset
	// will have an 'expirationTime' property set to the creation time plus
	// the value in this property, and changing the value will only affect
	// new tables, not existing ones. When the 'expirationTime' for a given
	// table is reached, that table will be deleted automatically.
	// If a table's 'expirationTime' is modified or removed before the
	// table expires, or if you provide an explicit 'expirationTime' when
	// creating a table, that value takes precedence over the default
	// expiration time indicated by this property.
	// +optional
	DefaultTableExpirationMs *int64 `json:"defaultTableExpirationMs,omitempty" tf:"default_table_expiration_ms"`
	// +optional
	DeleteContentsOnDestroy *bool `json:"deleteContentsOnDestroy,omitempty" tf:"delete_contents_on_destroy"`
	// A user-friendly description of the dataset
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// A hash of the resource.
	// +optional
	Etag *string `json:"etag,omitempty" tf:"etag"`
	// A descriptive name for the dataset
	// +optional
	FriendlyName *string `json:"friendlyName,omitempty" tf:"friendly_name"`
	// The labels associated with this dataset. You can use these to
	// organize and group your datasets
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// The date when this dataset or any of its tables was last modified, in
	// milliseconds since the epoch.
	// +optional
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty" tf:"last_modified_time"`
	// The geographic location where the dataset should reside.
	// See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
	//
	//
	// There are two types of locations, regional or multi-regional. A regional
	// location is a specific geographic place, such as Tokyo, and a multi-regional
	// location is a large geographic area, such as the United States, that
	// contains at least two geographic places.
	//
	//
	// The default value is multi-regional location 'US'.
	// Changing this forces a new resource to be created.
	// +optional
	Location *string `json:"location,omitempty" tf:"location"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// +optional
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link"`
}

type DatasetStatus struct {
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

// DatasetList is a list of Datasets
type DatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Dataset CRD objects
	Items []Dataset `json:"items,omitempty"`
}
