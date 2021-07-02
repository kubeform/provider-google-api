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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/healthcare/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeHealthcareV1alpha1 struct {
	*testing.Fake
}

func (c *FakeHealthcareV1alpha1) ConsentStores(namespace string) v1alpha1.ConsentStoreInterface {
	return &FakeConsentStores{c, namespace}
}

func (c *FakeHealthcareV1alpha1) ConsentStoreIamBindings(namespace string) v1alpha1.ConsentStoreIamBindingInterface {
	return &FakeConsentStoreIamBindings{c, namespace}
}

func (c *FakeHealthcareV1alpha1) ConsentStoreIamMembers(namespace string) v1alpha1.ConsentStoreIamMemberInterface {
	return &FakeConsentStoreIamMembers{c, namespace}
}

func (c *FakeHealthcareV1alpha1) ConsentStoreIamPolicies(namespace string) v1alpha1.ConsentStoreIamPolicyInterface {
	return &FakeConsentStoreIamPolicies{c, namespace}
}

func (c *FakeHealthcareV1alpha1) Datasets(namespace string) v1alpha1.DatasetInterface {
	return &FakeDatasets{c, namespace}
}

func (c *FakeHealthcareV1alpha1) DatasetIamBindings(namespace string) v1alpha1.DatasetIamBindingInterface {
	return &FakeDatasetIamBindings{c, namespace}
}

func (c *FakeHealthcareV1alpha1) DatasetIamMembers(namespace string) v1alpha1.DatasetIamMemberInterface {
	return &FakeDatasetIamMembers{c, namespace}
}

func (c *FakeHealthcareV1alpha1) DatasetIamPolicies(namespace string) v1alpha1.DatasetIamPolicyInterface {
	return &FakeDatasetIamPolicies{c, namespace}
}

func (c *FakeHealthcareV1alpha1) DicomStores(namespace string) v1alpha1.DicomStoreInterface {
	return &FakeDicomStores{c, namespace}
}

func (c *FakeHealthcareV1alpha1) DicomStoreIamBindings(namespace string) v1alpha1.DicomStoreIamBindingInterface {
	return &FakeDicomStoreIamBindings{c, namespace}
}

func (c *FakeHealthcareV1alpha1) DicomStoreIamMembers(namespace string) v1alpha1.DicomStoreIamMemberInterface {
	return &FakeDicomStoreIamMembers{c, namespace}
}

func (c *FakeHealthcareV1alpha1) DicomStoreIamPolicies(namespace string) v1alpha1.DicomStoreIamPolicyInterface {
	return &FakeDicomStoreIamPolicies{c, namespace}
}

func (c *FakeHealthcareV1alpha1) FhirStores(namespace string) v1alpha1.FhirStoreInterface {
	return &FakeFhirStores{c, namespace}
}

func (c *FakeHealthcareV1alpha1) FhirStoreIamBindings(namespace string) v1alpha1.FhirStoreIamBindingInterface {
	return &FakeFhirStoreIamBindings{c, namespace}
}

func (c *FakeHealthcareV1alpha1) FhirStoreIamMembers(namespace string) v1alpha1.FhirStoreIamMemberInterface {
	return &FakeFhirStoreIamMembers{c, namespace}
}

func (c *FakeHealthcareV1alpha1) FhirStoreIamPolicies(namespace string) v1alpha1.FhirStoreIamPolicyInterface {
	return &FakeFhirStoreIamPolicies{c, namespace}
}

func (c *FakeHealthcareV1alpha1) Hl7V2Stores(namespace string) v1alpha1.Hl7V2StoreInterface {
	return &FakeHl7V2Stores{c, namespace}
}

func (c *FakeHealthcareV1alpha1) Hl7V2StoreIamBindings(namespace string) v1alpha1.Hl7V2StoreIamBindingInterface {
	return &FakeHl7V2StoreIamBindings{c, namespace}
}

func (c *FakeHealthcareV1alpha1) Hl7V2StoreIamMembers(namespace string) v1alpha1.Hl7V2StoreIamMemberInterface {
	return &FakeHl7V2StoreIamMembers{c, namespace}
}

func (c *FakeHealthcareV1alpha1) Hl7V2StoreIamPolicies(namespace string) v1alpha1.Hl7V2StoreIamPolicyInterface {
	return &FakeHl7V2StoreIamPolicies{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeHealthcareV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
