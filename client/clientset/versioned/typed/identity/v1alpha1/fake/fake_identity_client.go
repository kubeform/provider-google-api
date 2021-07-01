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
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/provider-google-api/client/clientset/versioned/typed/identity/v1alpha1"
)

type FakeIdentityV1alpha1 struct {
	*testing.Fake
}

func (c *FakeIdentityV1alpha1) PlatformDefaultSupportedIdpConfigs(namespace string) v1alpha1.PlatformDefaultSupportedIdpConfigInterface {
	return &FakePlatformDefaultSupportedIdpConfigs{c, namespace}
}

func (c *FakeIdentityV1alpha1) PlatformInboundSamlConfigs(namespace string) v1alpha1.PlatformInboundSamlConfigInterface {
	return &FakePlatformInboundSamlConfigs{c, namespace}
}

func (c *FakeIdentityV1alpha1) PlatformOauthIdpConfigs(namespace string) v1alpha1.PlatformOauthIdpConfigInterface {
	return &FakePlatformOauthIdpConfigs{c, namespace}
}

func (c *FakeIdentityV1alpha1) PlatformTenants(namespace string) v1alpha1.PlatformTenantInterface {
	return &FakePlatformTenants{c, namespace}
}

func (c *FakeIdentityV1alpha1) PlatformTenantDefaultSupportedIdpConfigs(namespace string) v1alpha1.PlatformTenantDefaultSupportedIdpConfigInterface {
	return &FakePlatformTenantDefaultSupportedIdpConfigs{c, namespace}
}

func (c *FakeIdentityV1alpha1) PlatformTenantInboundSamlConfigs(namespace string) v1alpha1.PlatformTenantInboundSamlConfigInterface {
	return &FakePlatformTenantInboundSamlConfigs{c, namespace}
}

func (c *FakeIdentityV1alpha1) PlatformTenantOauthIdpConfigs(namespace string) v1alpha1.PlatformTenantOauthIdpConfigInterface {
	return &FakePlatformTenantOauthIdpConfigs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeIdentityV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
