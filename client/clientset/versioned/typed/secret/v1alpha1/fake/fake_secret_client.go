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
	v1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/secret/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSecretV1alpha1 struct {
	*testing.Fake
}

func (c *FakeSecretV1alpha1) ManagerSecrets(namespace string) v1alpha1.ManagerSecretInterface {
	return &FakeManagerSecrets{c, namespace}
}

func (c *FakeSecretV1alpha1) ManagerSecretIamBindings(namespace string) v1alpha1.ManagerSecretIamBindingInterface {
	return &FakeManagerSecretIamBindings{c, namespace}
}

func (c *FakeSecretV1alpha1) ManagerSecretIamMembers(namespace string) v1alpha1.ManagerSecretIamMemberInterface {
	return &FakeManagerSecretIamMembers{c, namespace}
}

func (c *FakeSecretV1alpha1) ManagerSecretIamPolicies(namespace string) v1alpha1.ManagerSecretIamPolicyInterface {
	return &FakeManagerSecretIamPolicies{c, namespace}
}

func (c *FakeSecretV1alpha1) ManagerSecretVersions(namespace string) v1alpha1.ManagerSecretVersionInterface {
	return &FakeManagerSecretVersions{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSecretV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
