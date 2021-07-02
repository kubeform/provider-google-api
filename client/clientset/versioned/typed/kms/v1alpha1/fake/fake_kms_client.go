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
	v1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/kms/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeKmsV1alpha1 struct {
	*testing.Fake
}

func (c *FakeKmsV1alpha1) CryptoKeys(namespace string) v1alpha1.CryptoKeyInterface {
	return &FakeCryptoKeys{c, namespace}
}

func (c *FakeKmsV1alpha1) CryptoKeyIamBindings(namespace string) v1alpha1.CryptoKeyIamBindingInterface {
	return &FakeCryptoKeyIamBindings{c, namespace}
}

func (c *FakeKmsV1alpha1) CryptoKeyIamMembers(namespace string) v1alpha1.CryptoKeyIamMemberInterface {
	return &FakeCryptoKeyIamMembers{c, namespace}
}

func (c *FakeKmsV1alpha1) CryptoKeyIamPolicies(namespace string) v1alpha1.CryptoKeyIamPolicyInterface {
	return &FakeCryptoKeyIamPolicies{c, namespace}
}

func (c *FakeKmsV1alpha1) KeyRings(namespace string) v1alpha1.KeyRingInterface {
	return &FakeKeyRings{c, namespace}
}

func (c *FakeKmsV1alpha1) KeyRingIamBindings(namespace string) v1alpha1.KeyRingIamBindingInterface {
	return &FakeKeyRingIamBindings{c, namespace}
}

func (c *FakeKmsV1alpha1) KeyRingIamMembers(namespace string) v1alpha1.KeyRingIamMemberInterface {
	return &FakeKeyRingIamMembers{c, namespace}
}

func (c *FakeKmsV1alpha1) KeyRingIamPolicies(namespace string) v1alpha1.KeyRingIamPolicyInterface {
	return &FakeKeyRingIamPolicies{c, namespace}
}

func (c *FakeKmsV1alpha1) KeyRingImportJobs(namespace string) v1alpha1.KeyRingImportJobInterface {
	return &FakeKeyRingImportJobs{c, namespace}
}

func (c *FakeKmsV1alpha1) SecretCiphertexts(namespace string) v1alpha1.SecretCiphertextInterface {
	return &FakeSecretCiphertexts{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKmsV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
