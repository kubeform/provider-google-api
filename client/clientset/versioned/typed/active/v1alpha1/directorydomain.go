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

	v1alpha1 "kubeform.dev/provider-google-api/apis/active/v1alpha1"
	scheme "kubeform.dev/provider-google-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DirectoryDomainsGetter has a method to return a DirectoryDomainInterface.
// A group's client should implement this interface.
type DirectoryDomainsGetter interface {
	DirectoryDomains(namespace string) DirectoryDomainInterface
}

// DirectoryDomainInterface has methods to work with DirectoryDomain resources.
type DirectoryDomainInterface interface {
	Create(ctx context.Context, directoryDomain *v1alpha1.DirectoryDomain, opts v1.CreateOptions) (*v1alpha1.DirectoryDomain, error)
	Update(ctx context.Context, directoryDomain *v1alpha1.DirectoryDomain, opts v1.UpdateOptions) (*v1alpha1.DirectoryDomain, error)
	UpdateStatus(ctx context.Context, directoryDomain *v1alpha1.DirectoryDomain, opts v1.UpdateOptions) (*v1alpha1.DirectoryDomain, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.DirectoryDomain, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DirectoryDomainList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DirectoryDomain, err error)
	DirectoryDomainExpansion
}

// directoryDomains implements DirectoryDomainInterface
type directoryDomains struct {
	client rest.Interface
	ns     string
}

// newDirectoryDomains returns a DirectoryDomains
func newDirectoryDomains(c *ActiveV1alpha1Client, namespace string) *directoryDomains {
	return &directoryDomains{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the directoryDomain, and returns the corresponding directoryDomain object, and an error if there is any.
func (c *directoryDomains) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DirectoryDomain, err error) {
	result = &v1alpha1.DirectoryDomain{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("directorydomains").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DirectoryDomains that match those selectors.
func (c *directoryDomains) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DirectoryDomainList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DirectoryDomainList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("directorydomains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested directoryDomains.
func (c *directoryDomains) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("directorydomains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a directoryDomain and creates it.  Returns the server's representation of the directoryDomain, and an error, if there is any.
func (c *directoryDomains) Create(ctx context.Context, directoryDomain *v1alpha1.DirectoryDomain, opts v1.CreateOptions) (result *v1alpha1.DirectoryDomain, err error) {
	result = &v1alpha1.DirectoryDomain{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("directorydomains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(directoryDomain).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a directoryDomain and updates it. Returns the server's representation of the directoryDomain, and an error, if there is any.
func (c *directoryDomains) Update(ctx context.Context, directoryDomain *v1alpha1.DirectoryDomain, opts v1.UpdateOptions) (result *v1alpha1.DirectoryDomain, err error) {
	result = &v1alpha1.DirectoryDomain{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("directorydomains").
		Name(directoryDomain.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(directoryDomain).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *directoryDomains) UpdateStatus(ctx context.Context, directoryDomain *v1alpha1.DirectoryDomain, opts v1.UpdateOptions) (result *v1alpha1.DirectoryDomain, err error) {
	result = &v1alpha1.DirectoryDomain{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("directorydomains").
		Name(directoryDomain.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(directoryDomain).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the directoryDomain and deletes it. Returns an error if one occurs.
func (c *directoryDomains) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("directorydomains").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *directoryDomains) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("directorydomains").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched directoryDomain.
func (c *directoryDomains) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DirectoryDomain, err error) {
	result = &v1alpha1.DirectoryDomain{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("directorydomains").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
