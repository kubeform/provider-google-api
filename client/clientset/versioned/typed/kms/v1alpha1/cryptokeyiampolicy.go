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
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubeform.dev/provider-google-api/apis/kms/v1alpha1"
	scheme "kubeform.dev/provider-google-api/client/clientset/versioned/scheme"
)

// CryptoKeyIamPoliciesGetter has a method to return a CryptoKeyIamPolicyInterface.
// A group's client should implement this interface.
type CryptoKeyIamPoliciesGetter interface {
	CryptoKeyIamPolicies(namespace string) CryptoKeyIamPolicyInterface
}

// CryptoKeyIamPolicyInterface has methods to work with CryptoKeyIamPolicy resources.
type CryptoKeyIamPolicyInterface interface {
	Create(ctx context.Context, cryptoKeyIamPolicy *v1alpha1.CryptoKeyIamPolicy, opts v1.CreateOptions) (*v1alpha1.CryptoKeyIamPolicy, error)
	Update(ctx context.Context, cryptoKeyIamPolicy *v1alpha1.CryptoKeyIamPolicy, opts v1.UpdateOptions) (*v1alpha1.CryptoKeyIamPolicy, error)
	UpdateStatus(ctx context.Context, cryptoKeyIamPolicy *v1alpha1.CryptoKeyIamPolicy, opts v1.UpdateOptions) (*v1alpha1.CryptoKeyIamPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.CryptoKeyIamPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CryptoKeyIamPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CryptoKeyIamPolicy, err error)
	CryptoKeyIamPolicyExpansion
}

// cryptoKeyIamPolicies implements CryptoKeyIamPolicyInterface
type cryptoKeyIamPolicies struct {
	client rest.Interface
	ns     string
}

// newCryptoKeyIamPolicies returns a CryptoKeyIamPolicies
func newCryptoKeyIamPolicies(c *KmsV1alpha1Client, namespace string) *cryptoKeyIamPolicies {
	return &cryptoKeyIamPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cryptoKeyIamPolicy, and returns the corresponding cryptoKeyIamPolicy object, and an error if there is any.
func (c *cryptoKeyIamPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CryptoKeyIamPolicy, err error) {
	result = &v1alpha1.CryptoKeyIamPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cryptokeyiampolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CryptoKeyIamPolicies that match those selectors.
func (c *cryptoKeyIamPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CryptoKeyIamPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CryptoKeyIamPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cryptokeyiampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cryptoKeyIamPolicies.
func (c *cryptoKeyIamPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cryptokeyiampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cryptoKeyIamPolicy and creates it.  Returns the server's representation of the cryptoKeyIamPolicy, and an error, if there is any.
func (c *cryptoKeyIamPolicies) Create(ctx context.Context, cryptoKeyIamPolicy *v1alpha1.CryptoKeyIamPolicy, opts v1.CreateOptions) (result *v1alpha1.CryptoKeyIamPolicy, err error) {
	result = &v1alpha1.CryptoKeyIamPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cryptokeyiampolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cryptoKeyIamPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cryptoKeyIamPolicy and updates it. Returns the server's representation of the cryptoKeyIamPolicy, and an error, if there is any.
func (c *cryptoKeyIamPolicies) Update(ctx context.Context, cryptoKeyIamPolicy *v1alpha1.CryptoKeyIamPolicy, opts v1.UpdateOptions) (result *v1alpha1.CryptoKeyIamPolicy, err error) {
	result = &v1alpha1.CryptoKeyIamPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cryptokeyiampolicies").
		Name(cryptoKeyIamPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cryptoKeyIamPolicy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *cryptoKeyIamPolicies) UpdateStatus(ctx context.Context, cryptoKeyIamPolicy *v1alpha1.CryptoKeyIamPolicy, opts v1.UpdateOptions) (result *v1alpha1.CryptoKeyIamPolicy, err error) {
	result = &v1alpha1.CryptoKeyIamPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cryptokeyiampolicies").
		Name(cryptoKeyIamPolicy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cryptoKeyIamPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cryptoKeyIamPolicy and deletes it. Returns an error if one occurs.
func (c *cryptoKeyIamPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cryptokeyiampolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cryptoKeyIamPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cryptokeyiampolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cryptoKeyIamPolicy.
func (c *cryptoKeyIamPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CryptoKeyIamPolicy, err error) {
	result = &v1alpha1.CryptoKeyIamPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cryptokeyiampolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
