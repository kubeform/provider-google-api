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
	"kubeform.dev/provider-google-api/util"

	jsoniter "github.com/json-iterator/go"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (r *JobIamBinding) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-dataproc-google-kubeform-com-v1alpha1-jobiambinding,mutating=false,failurePolicy=fail,groups=dataproc.google.kubeform.com,resources=jobiambindings,versions=v1alpha1,name=jobiambinding.dataproc.google.kubeform.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &JobIamBinding{}

var jobiambindingForceNewList = map[string]bool{
	"/condition/*/description": true,
	"/condition/*/expression":  true,
	"/condition/*/title":       true,
	"/job_id":                  true,
	"/project":                 true,
	"/region":                  true,
	"/role":                    true,
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *JobIamBinding) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *JobIamBinding) ValidateUpdate(old runtime.Object) error {
	if r.Spec.Resource.ID == "" {
		return nil
	}
	newObj := r.Spec.Resource
	res := old.(*JobIamBinding)
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

	for key := range jobiambindingForceNewList {
		keySplit := strings.Split(key, "/*")
		length := len(keySplit)
		checkIfAnyDif := false
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempOld, tempOld, tempNew)
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempNew, tempOld, tempNew)

		if checkIfAnyDif && r.Spec.UpdatePolicy == base.UpdatePolicyDoNotDestroy {
			return fmt.Errorf(`jobiambinding "%v/%v" immutable field can't be updated. To update, change spec.updatePolicy to Destroy`, r.Namespace, r.Name)
		}
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *JobIamBinding) ValidateDelete() error {
	if r.Spec.TerminationPolicy == base.TerminationPolicyDoNotTerminate {
		return fmt.Errorf(`jobiambinding "%v/%v" can't be terminated. To delete, change spec.terminationPolicy to Delete`, r.Namespace, r.Name)
	}
	return nil
}
