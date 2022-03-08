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
	v1alpha1 "kubeform.dev/provider-google-api/apis/notebooks/v1alpha1"
	"kubeform.dev/provider-google-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type NotebooksV1alpha1Interface interface {
	RESTClient() rest.Interface
	EnvironmentsGetter
	InstancesGetter
	InstanceIamBindingsGetter
	InstanceIamMembersGetter
	InstanceIamPoliciesGetter
	LocationsGetter
	RuntimesGetter
	RuntimeIamBindingsGetter
	RuntimeIamMembersGetter
	RuntimeIamPoliciesGetter
}

// NotebooksV1alpha1Client is used to interact with features provided by the notebooks.google.kubeform.com group.
type NotebooksV1alpha1Client struct {
	restClient rest.Interface
}

func (c *NotebooksV1alpha1Client) Environments(namespace string) EnvironmentInterface {
	return newEnvironments(c, namespace)
}

func (c *NotebooksV1alpha1Client) Instances(namespace string) InstanceInterface {
	return newInstances(c, namespace)
}

func (c *NotebooksV1alpha1Client) InstanceIamBindings(namespace string) InstanceIamBindingInterface {
	return newInstanceIamBindings(c, namespace)
}

func (c *NotebooksV1alpha1Client) InstanceIamMembers(namespace string) InstanceIamMemberInterface {
	return newInstanceIamMembers(c, namespace)
}

func (c *NotebooksV1alpha1Client) InstanceIamPolicies(namespace string) InstanceIamPolicyInterface {
	return newInstanceIamPolicies(c, namespace)
}

func (c *NotebooksV1alpha1Client) Locations(namespace string) LocationInterface {
	return newLocations(c, namespace)
}

func (c *NotebooksV1alpha1Client) Runtimes(namespace string) RuntimeInterface {
	return newRuntimes(c, namespace)
}

func (c *NotebooksV1alpha1Client) RuntimeIamBindings(namespace string) RuntimeIamBindingInterface {
	return newRuntimeIamBindings(c, namespace)
}

func (c *NotebooksV1alpha1Client) RuntimeIamMembers(namespace string) RuntimeIamMemberInterface {
	return newRuntimeIamMembers(c, namespace)
}

func (c *NotebooksV1alpha1Client) RuntimeIamPolicies(namespace string) RuntimeIamPolicyInterface {
	return newRuntimeIamPolicies(c, namespace)
}

// NewForConfig creates a new NotebooksV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*NotebooksV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &NotebooksV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new NotebooksV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *NotebooksV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new NotebooksV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *NotebooksV1alpha1Client {
	return &NotebooksV1alpha1Client{c}
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
func (c *NotebooksV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
