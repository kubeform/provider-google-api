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

type BillingBudget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BillingBudgetSpec   `json:"spec,omitempty"`
	Status            BillingBudgetStatus `json:"status,omitempty"`
}

type BillingBudgetSpecAllUpdatesRule struct {
	// Boolean. When set to true, disables default notifications sent
	// when a threshold is exceeded. Default recipients are
	// those with Billing Account Administrators and Billing
	// Account Users IAM roles for the target account.
	// +optional
	DisableDefaultIamRecipients *bool `json:"disableDefaultIamRecipients,omitempty" tf:"disable_default_iam_recipients"`
	// The full resource name of a monitoring notification
	// channel in the form
	// projects/{project_id}/notificationChannels/{channel_id}.
	// A maximum of 5 channels are allowed.
	// +optional
	MonitoringNotificationChannels []string `json:"monitoringNotificationChannels,omitempty" tf:"monitoring_notification_channels"`
	// The name of the Cloud Pub/Sub topic where budget related
	// messages will be published, in the form
	// projects/{project_id}/topics/{topic_id}. Updates are sent
	// at regular intervals to the topic.
	// +optional
	PubsubTopic *string `json:"pubsubTopic,omitempty" tf:"pubsub_topic"`
	// The schema version of the notification. Only "1.0" is
	// accepted. It represents the JSON schema as defined in
	// https://cloud.google.com/billing/docs/how-to/budgets#notification_format.
	// +optional
	SchemaVersion *string `json:"schemaVersion,omitempty" tf:"schema_version"`
}

type BillingBudgetSpecAmountSpecifiedAmount struct {
	// The 3-letter currency code defined in ISO 4217.
	// +optional
	CurrencyCode *string `json:"currencyCode,omitempty" tf:"currency_code"`
	// Number of nano (10^-9) units of the amount.
	// The value must be between -999,999,999 and +999,999,999
	// inclusive. If units is positive, nanos must be positive or
	// zero. If units is zero, nanos can be positive, zero, or
	// negative. If units is negative, nanos must be negative or
	// zero. For example $-1.75 is represented as units=-1 and
	// nanos=-750,000,000.
	// +optional
	Nanos *int64 `json:"nanos,omitempty" tf:"nanos"`
	// The whole units of the amount. For example if currencyCode
	// is "USD", then 1 unit is one US dollar.
	// +optional
	Units *string `json:"units,omitempty" tf:"units"`
}

type BillingBudgetSpecAmount struct {
	// Configures a budget amount that is automatically set to 100% of
	// last period's spend.
	// Boolean. Set value to true to use. Do not set to false, instead
	// use the 'specified_amount' block.
	// +optional
	LastPeriodAmount *bool `json:"lastPeriodAmount,omitempty" tf:"last_period_amount"`
	// A specified amount to use as the budget. currencyCode is
	// optional. If specified, it must match the currency of the
	// billing account. The currencyCode is provided on output.
	// +optional
	SpecifiedAmount *BillingBudgetSpecAmountSpecifiedAmount `json:"specifiedAmount,omitempty" tf:"specified_amount"`
}

type BillingBudgetSpecBudgetFilter struct {
	// A set of subaccounts of the form billingAccounts/{account_id},
	// specifying that usage from only this set of subaccounts should
	// be included in the budget. If a subaccount is set to the name of
	// the parent account, usage from the parent account will be included.
	// If the field is omitted, the report will include usage from the parent
	// account and all subaccounts, if they exist.
	// +optional
	CreditTypes []string `json:"creditTypes,omitempty" tf:"credit_types"`
	// Specifies how credits should be treated when determining spend
	// for threshold calculations. Default value: "INCLUDE_ALL_CREDITS" Possible values: ["INCLUDE_ALL_CREDITS", "EXCLUDE_ALL_CREDITS", "INCLUDE_SPECIFIED_CREDITS"]
	// +optional
	CreditTypesTreatment *string `json:"creditTypesTreatment,omitempty" tf:"credit_types_treatment"`
	// A single label and value pair specifying that usage from only
	// this set of labeled resources should be included in the budget.
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// A set of projects of the form projects/{project_number},
	// specifying that usage from only this set of projects should be
	// included in the budget. If omitted, the report will include
	// all usage for the billing account, regardless of which project
	// the usage occurred on.
	// +optional
	Projects []string `json:"projects,omitempty" tf:"projects"`
	// A set of services of the form services/{service_id},
	// specifying that usage from only this set of services should be
	// included in the budget. If omitted, the report will include
	// usage for all the services. The service names are available
	// through the Catalog API:
	// https://cloud.google.com/billing/v1/how-tos/catalog-api.
	// +optional
	Services []string `json:"services,omitempty" tf:"services"`
	// A set of subaccounts of the form billingAccounts/{account_id},
	// specifying that usage from only this set of subaccounts should
	// be included in the budget. If a subaccount is set to the name of
	// the parent account, usage from the parent account will be included.
	// If the field is omitted, the report will include usage from the parent
	// account and all subaccounts, if they exist.
	// +optional
	Subaccounts []string `json:"subaccounts,omitempty" tf:"subaccounts"`
}

type BillingBudgetSpecThresholdRules struct {
	// The type of basis used to determine if spend has passed
	// the threshold. Default value: "CURRENT_SPEND" Possible values: ["CURRENT_SPEND", "FORECASTED_SPEND"]
	// +optional
	SpendBasis *string `json:"spendBasis,omitempty" tf:"spend_basis"`
	// Send an alert when this threshold is exceeded. This is a
	// 1.0-based percentage, so 0.5 = 50%. Must be >= 0.
	ThresholdPercent *float64 `json:"thresholdPercent" tf:"threshold_percent"`
}

type BillingBudgetSpec struct {
	State *BillingBudgetSpecResource `json:"state,omitempty" tf:"-"`

	Resource BillingBudgetSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type BillingBudgetSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Defines notifications that are sent on every update to the
	// billing account's spend, regardless of the thresholds defined
	// using threshold rules.
	// +optional
	AllUpdatesRule *BillingBudgetSpecAllUpdatesRule `json:"allUpdatesRule,omitempty" tf:"all_updates_rule"`
	// The budgeted amount for each usage period.
	Amount *BillingBudgetSpecAmount `json:"amount" tf:"amount"`
	// ID of the billing account to set a budget on.
	BillingAccount *string `json:"billingAccount" tf:"billing_account"`
	// Filters that define which resources are used to compute the actual
	// spend against the budget.
	// +optional
	BudgetFilter *BillingBudgetSpecBudgetFilter `json:"budgetFilter,omitempty" tf:"budget_filter"`
	// User data for display name in UI. Must be <= 60 chars.
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// Resource name of the budget. The resource name
	// implies the scope of a budget. Values are of the form
	// billingAccounts/{billingAccountId}/budgets/{budgetId}.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// Rules that trigger alerts (notifications of thresholds being
	// crossed) when spend exceeds the specified percentages of the
	// budget.
	ThresholdRules []BillingBudgetSpecThresholdRules `json:"thresholdRules" tf:"threshold_rules"`
}

type BillingBudgetStatus struct {
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

// BillingBudgetList is a list of BillingBudgets
type BillingBudgetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BillingBudget CRD objects
	Items []BillingBudget `json:"items,omitempty"`
}
