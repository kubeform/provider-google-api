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

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-google-api/apis/apigee/v1alpha1"
	"kubeform.dev/provider-google-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type ApigeeV1alpha1Interface interface {
	RESTClient() rest.Interface
	EndpointAttachmentsGetter
	EnvgroupsGetter
	EnvgroupAttachmentsGetter
	EnvironmentsGetter
	EnvironmentIamBindingsGetter
	EnvironmentIamMembersGetter
	EnvironmentIamPoliciesGetter
	InstancesGetter
	InstanceAttachmentsGetter
	OrganizationsGetter
}

// ApigeeV1alpha1Client is used to interact with features provided by the apigee.google.kubeform.com group.
type ApigeeV1alpha1Client struct {
	restClient rest.Interface
}

func (c *ApigeeV1alpha1Client) EndpointAttachments(namespace string) EndpointAttachmentInterface {
	return newEndpointAttachments(c, namespace)
}

func (c *ApigeeV1alpha1Client) Envgroups(namespace string) EnvgroupInterface {
	return newEnvgroups(c, namespace)
}

func (c *ApigeeV1alpha1Client) EnvgroupAttachments(namespace string) EnvgroupAttachmentInterface {
	return newEnvgroupAttachments(c, namespace)
}

func (c *ApigeeV1alpha1Client) Environments(namespace string) EnvironmentInterface {
	return newEnvironments(c, namespace)
}

func (c *ApigeeV1alpha1Client) EnvironmentIamBindings(namespace string) EnvironmentIamBindingInterface {
	return newEnvironmentIamBindings(c, namespace)
}

func (c *ApigeeV1alpha1Client) EnvironmentIamMembers(namespace string) EnvironmentIamMemberInterface {
	return newEnvironmentIamMembers(c, namespace)
}

func (c *ApigeeV1alpha1Client) EnvironmentIamPolicies(namespace string) EnvironmentIamPolicyInterface {
	return newEnvironmentIamPolicies(c, namespace)
}

func (c *ApigeeV1alpha1Client) Instances(namespace string) InstanceInterface {
	return newInstances(c, namespace)
}

func (c *ApigeeV1alpha1Client) InstanceAttachments(namespace string) InstanceAttachmentInterface {
	return newInstanceAttachments(c, namespace)
}

func (c *ApigeeV1alpha1Client) Organizations(namespace string) OrganizationInterface {
	return newOrganizations(c, namespace)
}

// NewForConfig creates a new ApigeeV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*ApigeeV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ApigeeV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new ApigeeV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ApigeeV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ApigeeV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *ApigeeV1alpha1Client {
	return &ApigeeV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ApigeeV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
