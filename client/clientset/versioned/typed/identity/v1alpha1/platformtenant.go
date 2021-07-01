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
	v1alpha1 "kubeform.dev/provider-google-api/apis/identity/v1alpha1"
	scheme "kubeform.dev/provider-google-api/client/clientset/versioned/scheme"
)

// PlatformTenantsGetter has a method to return a PlatformTenantInterface.
// A group's client should implement this interface.
type PlatformTenantsGetter interface {
	PlatformTenants(namespace string) PlatformTenantInterface
}

// PlatformTenantInterface has methods to work with PlatformTenant resources.
type PlatformTenantInterface interface {
	Create(ctx context.Context, platformTenant *v1alpha1.PlatformTenant, opts v1.CreateOptions) (*v1alpha1.PlatformTenant, error)
	Update(ctx context.Context, platformTenant *v1alpha1.PlatformTenant, opts v1.UpdateOptions) (*v1alpha1.PlatformTenant, error)
	UpdateStatus(ctx context.Context, platformTenant *v1alpha1.PlatformTenant, opts v1.UpdateOptions) (*v1alpha1.PlatformTenant, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.PlatformTenant, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.PlatformTenantList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PlatformTenant, err error)
	PlatformTenantExpansion
}

// platformTenants implements PlatformTenantInterface
type platformTenants struct {
	client rest.Interface
	ns     string
}

// newPlatformTenants returns a PlatformTenants
func newPlatformTenants(c *IdentityV1alpha1Client, namespace string) *platformTenants {
	return &platformTenants{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the platformTenant, and returns the corresponding platformTenant object, and an error if there is any.
func (c *platformTenants) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PlatformTenant, err error) {
	result = &v1alpha1.PlatformTenant{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("platformtenants").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PlatformTenants that match those selectors.
func (c *platformTenants) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PlatformTenantList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PlatformTenantList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("platformtenants").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested platformTenants.
func (c *platformTenants) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("platformtenants").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a platformTenant and creates it.  Returns the server's representation of the platformTenant, and an error, if there is any.
func (c *platformTenants) Create(ctx context.Context, platformTenant *v1alpha1.PlatformTenant, opts v1.CreateOptions) (result *v1alpha1.PlatformTenant, err error) {
	result = &v1alpha1.PlatformTenant{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("platformtenants").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(platformTenant).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a platformTenant and updates it. Returns the server's representation of the platformTenant, and an error, if there is any.
func (c *platformTenants) Update(ctx context.Context, platformTenant *v1alpha1.PlatformTenant, opts v1.UpdateOptions) (result *v1alpha1.PlatformTenant, err error) {
	result = &v1alpha1.PlatformTenant{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("platformtenants").
		Name(platformTenant.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(platformTenant).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *platformTenants) UpdateStatus(ctx context.Context, platformTenant *v1alpha1.PlatformTenant, opts v1.UpdateOptions) (result *v1alpha1.PlatformTenant, err error) {
	result = &v1alpha1.PlatformTenant{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("platformtenants").
		Name(platformTenant.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(platformTenant).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the platformTenant and deletes it. Returns an error if one occurs.
func (c *platformTenants) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("platformtenants").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *platformTenants) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("platformtenants").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched platformTenant.
func (c *platformTenants) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PlatformTenant, err error) {
	result = &v1alpha1.PlatformTenant{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("platformtenants").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
