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

func (r *Certificate) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-privateca-google-kubeform-com-v1alpha1-certificate,mutating=false,failurePolicy=fail,groups=privateca.google.kubeform.com,resources=certificates,versions=v1alpha1,name=certificate.privateca.google.kubeform.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &Certificate{}

var certificateForceNewList = map[string]bool{
	"/certificate_authority":                                                           true,
	"/certificate_template":                                                            true,
	"/config/*/public_key/*/format":                                                    true,
	"/config/*/public_key/*/key":                                                       true,
	"/config/*/subject_config/*/subject/*/common_name":                                 true,
	"/config/*/subject_config/*/subject/*/country_code":                                true,
	"/config/*/subject_config/*/subject/*/locality":                                    true,
	"/config/*/subject_config/*/subject/*/organization":                                true,
	"/config/*/subject_config/*/subject/*/organizational_unit":                         true,
	"/config/*/subject_config/*/subject/*/postal_code":                                 true,
	"/config/*/subject_config/*/subject/*/province":                                    true,
	"/config/*/subject_config/*/subject/*/street_address":                              true,
	"/config/*/subject_config/*/subject_alt_name/*/dns_names":                          true,
	"/config/*/subject_config/*/subject_alt_name/*/email_addresses":                    true,
	"/config/*/subject_config/*/subject_alt_name/*/ip_addresses":                       true,
	"/config/*/subject_config/*/subject_alt_name/*/uris":                               true,
	"/config/*/x509_config/*/additional_extensions/*/critical":                         true,
	"/config/*/x509_config/*/additional_extensions/*/object_id/*/object_id_path":       true,
	"/config/*/x509_config/*/additional_extensions/*/value":                            true,
	"/config/*/x509_config/*/aia_ocsp_servers":                                         true,
	"/config/*/x509_config/*/ca_options/*/is_ca":                                       true,
	"/config/*/x509_config/*/ca_options/*/max_issuer_path_length":                      true,
	"/config/*/x509_config/*/ca_options/*/non_ca":                                      true,
	"/config/*/x509_config/*/ca_options/*/zero_max_issuer_path_length":                 true,
	"/config/*/x509_config/*/key_usage/*/base_key_usage/*/cert_sign":                   true,
	"/config/*/x509_config/*/key_usage/*/base_key_usage/*/content_commitment":          true,
	"/config/*/x509_config/*/key_usage/*/base_key_usage/*/crl_sign":                    true,
	"/config/*/x509_config/*/key_usage/*/base_key_usage/*/data_encipherment":           true,
	"/config/*/x509_config/*/key_usage/*/base_key_usage/*/decipher_only":               true,
	"/config/*/x509_config/*/key_usage/*/base_key_usage/*/digital_signature":           true,
	"/config/*/x509_config/*/key_usage/*/base_key_usage/*/encipher_only":               true,
	"/config/*/x509_config/*/key_usage/*/base_key_usage/*/key_agreement":               true,
	"/config/*/x509_config/*/key_usage/*/base_key_usage/*/key_encipherment":            true,
	"/config/*/x509_config/*/key_usage/*/extended_key_usage/*/client_auth":             true,
	"/config/*/x509_config/*/key_usage/*/extended_key_usage/*/code_signing":            true,
	"/config/*/x509_config/*/key_usage/*/extended_key_usage/*/email_protection":        true,
	"/config/*/x509_config/*/key_usage/*/extended_key_usage/*/ocsp_signing":            true,
	"/config/*/x509_config/*/key_usage/*/extended_key_usage/*/server_auth":             true,
	"/config/*/x509_config/*/key_usage/*/extended_key_usage/*/time_stamping":           true,
	"/config/*/x509_config/*/key_usage/*/unknown_extended_key_usages/*/object_id_path": true,
	"/config/*/x509_config/*/policy_ids/*/object_id_path":                              true,
	"/labels":   true,
	"/lifetime": true,
	"/location": true,
	"/name":     true,
	"/pem_csr":  true,
	"/pool":     true,
	"/project":  true,
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Certificate) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Certificate) ValidateUpdate(old runtime.Object) error {
	if r.Spec.Resource.ID == "" {
		return nil
	}
	newObj := r.Spec.Resource
	res := old.(*Certificate)
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

	for key, _ := range certificateForceNewList {
		keySplit := strings.Split(key, "/*")
		length := len(keySplit)
		checkIfAnyDif := false
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempOld, tempOld, tempNew)
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempNew, tempOld, tempNew)

		if checkIfAnyDif && r.Spec.UpdatePolicy == base.UpdatePolicyDoNotDestroy {
			return fmt.Errorf(`certificate "%v/%v" immutable field can't be updated. To update, change spec.updatePolicy to Destroy`, r.Namespace, r.Name)
		}
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Certificate) ValidateDelete() error {
	if r.Spec.TerminationPolicy == base.TerminationPolicyDoNotTerminate {
		return fmt.Errorf(`certificate "%v/%v" can't be terminated. To delete, change spec.terminationPolicy to Delete`, r.Namespace, r.Name)
	}
	return nil
}
