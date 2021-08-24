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

func (r *InstanceTemplate) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-compute-google-kubeform-com-v1alpha1-instancetemplate,mutating=false,failurePolicy=fail,groups=compute.google.kubeform.com,resources=instancetemplates,versions=v1alpha1,name=instancetemplate.compute.google.kubeform.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &InstanceTemplate{}

var instancetemplateForceNewList = map[string]bool{
	"/advanced_machine_features/*/enable_nested_virtualization": true,
	"/advanced_machine_features/*/threads_per_core":             true,
	"/can_ip_forward": true,
	"/confidential_instance_config/*/enable_confidential_compute": true,
	"/description":        true,
	"/disk/*/auto_delete": true,
	"/disk/*/boot":        true,
	"/disk/*/device_name": true,
	"/disk/*/disk_encryption_key/*/kms_key_self_link": true,
	"/disk/*/disk_name":          true,
	"/disk/*/disk_size_gb":       true,
	"/disk/*/disk_type":          true,
	"/disk/*/interface":          true,
	"/disk/*/labels":             true,
	"/disk/*/mode":               true,
	"/disk/*/resource_policies":  true,
	"/disk/*/source":             true,
	"/disk/*/source_image":       true,
	"/disk/*/type":               true,
	"/enable_display":            true,
	"/guest_accelerator/*/count": true,
	"/guest_accelerator/*/type":  true,
	"/instance_description":      true,
	"/labels":                    true,
	"/machine_type":              true,
	"/metadata":                  true,
	"/metadata_startup_script":   true,
	"/min_cpu_platform":          true,
	"/name":                      true,
	"/name_prefix":               true,
	"/network_interface/*/access_config/*/nat_ip":                 true,
	"/network_interface/*/access_config/*/network_tier":           true,
	"/network_interface/*/alias_ip_range/*/ip_cidr_range":         true,
	"/network_interface/*/alias_ip_range/*/subnetwork_range_name": true,
	"/network_interface/*/network":                                true,
	"/network_interface/*/network_ip":                             true,
	"/network_interface/*/nic_type":                               true,
	"/network_interface/*/subnetwork":                             true,
	"/network_interface/*/subnetwork_project":                     true,
	"/project": true,
	"/region":  true,
	"/reservation_affinity/*/specific_reservation/*/key":      true,
	"/reservation_affinity/*/specific_reservation/*/values":   true,
	"/reservation_affinity/*/type":                            true,
	"/scheduling/*/automatic_restart":                         true,
	"/scheduling/*/on_host_maintenance":                       true,
	"/scheduling/*/preemptible":                               true,
	"/service_account/*/email":                                true,
	"/service_account/*/scopes":                               true,
	"/shielded_instance_config/*/enable_integrity_monitoring": true,
	"/shielded_instance_config/*/enable_secure_boot":          true,
	"/shielded_instance_config/*/enable_vtpm":                 true,
	"/tags": true,
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *InstanceTemplate) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *InstanceTemplate) ValidateUpdate(old runtime.Object) error {
	if r.Spec.Resource.ID == "" {
		return nil
	}
	newObj := r.Spec.Resource
	res := old.(*InstanceTemplate)
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

	for key := range instancetemplateForceNewList {
		keySplit := strings.Split(key, "/*")
		length := len(keySplit)
		checkIfAnyDif := false
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempOld, tempOld, tempNew)
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempNew, tempOld, tempNew)

		if checkIfAnyDif && r.Spec.UpdatePolicy == base.UpdatePolicyDoNotDestroy {
			return fmt.Errorf(`instancetemplate "%v/%v" immutable field can't be updated. To update, change spec.updatePolicy to Destroy`, r.Namespace, r.Name)
		}
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *InstanceTemplate) ValidateDelete() error {
	if r.Spec.TerminationPolicy == base.TerminationPolicyDoNotTerminate {
		return fmt.Errorf(`instancetemplate "%v/%v" can't be terminated. To delete, change spec.terminationPolicy to Delete`, r.Namespace, r.Name)
	}
	return nil
}
