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
	"fmt"
	"strings"

	base "kubeform.dev/apimachinery/api/v1alpha1"
	"kubeform.dev/apimachinery/pkg/util"

	jsoniter "github.com/json-iterator/go"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (r *ConfigPatchDeployment) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-os-google-kubeform-com-v1alpha1-configpatchdeployment,mutating=false,failurePolicy=fail,groups=os.google.kubeform.com,resources=configpatchdeployments,versions=v1alpha1,name=configpatchdeployment.os.google.kubeform.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &ConfigPatchDeployment{}

var configpatchdeploymentForceNewList = map[string]bool{
	"/description":           true,
	"/duration":              true,
	"/instance_filter/*/all": true,
	"/instance_filter/*/group_labels/*/labels":                                              true,
	"/instance_filter/*/instance_name_prefixes":                                             true,
	"/instance_filter/*/instances":                                                          true,
	"/instance_filter/*/zones":                                                              true,
	"/one_time_schedule/*/execute_time":                                                     true,
	"/patch_config/*/apt/*/excludes":                                                        true,
	"/patch_config/*/apt/*/exclusive_packages":                                              true,
	"/patch_config/*/apt/*/type":                                                            true,
	"/patch_config/*/goo/*/enabled":                                                         true,
	"/patch_config/*/post_step/*/linux_exec_step_config/*/allowed_success_codes":            true,
	"/patch_config/*/post_step/*/linux_exec_step_config/*/gcs_object/*/bucket":              true,
	"/patch_config/*/post_step/*/linux_exec_step_config/*/gcs_object/*/generation_number":   true,
	"/patch_config/*/post_step/*/linux_exec_step_config/*/gcs_object/*/object":              true,
	"/patch_config/*/post_step/*/linux_exec_step_config/*/interpreter":                      true,
	"/patch_config/*/post_step/*/linux_exec_step_config/*/local_path":                       true,
	"/patch_config/*/post_step/*/windows_exec_step_config/*/allowed_success_codes":          true,
	"/patch_config/*/post_step/*/windows_exec_step_config/*/gcs_object/*/bucket":            true,
	"/patch_config/*/post_step/*/windows_exec_step_config/*/gcs_object/*/generation_number": true,
	"/patch_config/*/post_step/*/windows_exec_step_config/*/gcs_object/*/object":            true,
	"/patch_config/*/post_step/*/windows_exec_step_config/*/interpreter":                    true,
	"/patch_config/*/post_step/*/windows_exec_step_config/*/local_path":                     true,
	"/patch_config/*/pre_step/*/linux_exec_step_config/*/allowed_success_codes":             true,
	"/patch_config/*/pre_step/*/linux_exec_step_config/*/gcs_object/*/bucket":               true,
	"/patch_config/*/pre_step/*/linux_exec_step_config/*/gcs_object/*/generation_number":    true,
	"/patch_config/*/pre_step/*/linux_exec_step_config/*/gcs_object/*/object":               true,
	"/patch_config/*/pre_step/*/linux_exec_step_config/*/interpreter":                       true,
	"/patch_config/*/pre_step/*/linux_exec_step_config/*/local_path":                        true,
	"/patch_config/*/pre_step/*/windows_exec_step_config/*/allowed_success_codes":           true,
	"/patch_config/*/pre_step/*/windows_exec_step_config/*/gcs_object/*/bucket":             true,
	"/patch_config/*/pre_step/*/windows_exec_step_config/*/gcs_object/*/generation_number":  true,
	"/patch_config/*/pre_step/*/windows_exec_step_config/*/gcs_object/*/object":             true,
	"/patch_config/*/pre_step/*/windows_exec_step_config/*/interpreter":                     true,
	"/patch_config/*/pre_step/*/windows_exec_step_config/*/local_path":                      true,
	"/patch_config/*/reboot_config":                                                         true,
	"/patch_config/*/windows_update/*/classifications":                                      true,
	"/patch_config/*/windows_update/*/excludes":                                             true,
	"/patch_config/*/windows_update/*/exclusive_patches":                                    true,
	"/patch_config/*/yum/*/excludes":                                                        true,
	"/patch_config/*/yum/*/exclusive_packages":                                              true,
	"/patch_config/*/yum/*/minimal":                                                         true,
	"/patch_config/*/yum/*/security":                                                        true,
	"/patch_config/*/zypper/*/categories":                                                   true,
	"/patch_config/*/zypper/*/excludes":                                                     true,
	"/patch_config/*/zypper/*/exclusive_patches":                                            true,
	"/patch_config/*/zypper/*/severities":                                                   true,
	"/patch_config/*/zypper/*/with_optional":                                                true,
	"/patch_config/*/zypper/*/with_update":                                                  true,
	"/patch_deployment_id":                                                                  true,
	"/project":                                                                              true,
	"/recurring_schedule/*/end_time":                                                        true,
	"/recurring_schedule/*/monthly/*/month_day":                                             true,
	"/recurring_schedule/*/monthly/*/week_day_of_month/*/day_of_week":                       true,
	"/recurring_schedule/*/monthly/*/week_day_of_month/*/week_ordinal":                      true,
	"/recurring_schedule/*/start_time":                                                      true,
	"/recurring_schedule/*/time_of_day/*/hours":                                             true,
	"/recurring_schedule/*/time_of_day/*/minutes":                                           true,
	"/recurring_schedule/*/time_of_day/*/nanos":                                             true,
	"/recurring_schedule/*/time_of_day/*/seconds":                                           true,
	"/recurring_schedule/*/time_zone/*/id":                                                  true,
	"/recurring_schedule/*/time_zone/*/version":                                             true,
	"/recurring_schedule/*/weekly/*/day_of_week":                                            true,
	"/rollout/*/disruption_budget/*/fixed":                                                  true,
	"/rollout/*/disruption_budget/*/percentage":                                             true,
	"/rollout/*/mode": true,
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *ConfigPatchDeployment) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *ConfigPatchDeployment) ValidateUpdate(old runtime.Object) error {
	if r.Spec.Resource.ID == "" {
		return nil
	}
	newObj := r.Spec.Resource
	res := old.(*ConfigPatchDeployment)
	oldObj := res.Spec.Resource

	jsnitr := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		TagKey:                 "tf",
		ValidateJsonRawMessage: true,
		TypeEncoders:           GetEncoder(),
		TypeDecoders:           GetDecoder(),
	}.Froze()

	byteNew, err := jsnitr.Marshal(newObj)
	if err != nil {
		return err
	}
	tempNew := make(map[string]interface{})
	err = jsnitr.Unmarshal(byteNew, &tempNew)
	if err != nil {
		return err
	}

	byteOld, err := jsnitr.Marshal(oldObj)
	if err != nil {
		return err
	}
	tempOld := make(map[string]interface{})
	err = jsnitr.Unmarshal(byteOld, &tempOld)
	if err != nil {
		return err
	}

	for key, _ := range configpatchdeploymentForceNewList {
		keySplit := strings.Split(key, "/*")
		length := len(keySplit)
		checkIfAnyDif := false
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempOld, tempOld, tempNew)
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempNew, tempOld, tempNew)

		if checkIfAnyDif && r.Spec.UpdatePolicy == base.UpdatePolicyDoNotDestroy {
			return fmt.Errorf(`configpatchdeployment "%v/%v" immutable field can't be updated. To update, change spec.updatePolicy to Destroy`, r.Namespace, r.Name)
		}
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *ConfigPatchDeployment) ValidateDelete() error {
	if r.Spec.TerminationPolicy == base.TerminationPolicyDoNotTerminate {
		return fmt.Errorf(`configpatchdeployment "%v/%v" can't be terminated. To delete, change spec.terminationPolicy to Delete`, r.Namespace, r.Name)
	}
	return nil
}
