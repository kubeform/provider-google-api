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
	v1alpha1 "kubeform.dev/provider-google-api/apis/healthcare/v1alpha1"
	scheme "kubeform.dev/provider-google-api/client/clientset/versioned/scheme"
)

// Hl7V2StoreIamMembersGetter has a method to return a Hl7V2StoreIamMemberInterface.
// A group's client should implement this interface.
type Hl7V2StoreIamMembersGetter interface {
	Hl7V2StoreIamMembers(namespace string) Hl7V2StoreIamMemberInterface
}

// Hl7V2StoreIamMemberInterface has methods to work with Hl7V2StoreIamMember resources.
type Hl7V2StoreIamMemberInterface interface {
	Create(ctx context.Context, hl7V2StoreIamMember *v1alpha1.Hl7V2StoreIamMember, opts v1.CreateOptions) (*v1alpha1.Hl7V2StoreIamMember, error)
	Update(ctx context.Context, hl7V2StoreIamMember *v1alpha1.Hl7V2StoreIamMember, opts v1.UpdateOptions) (*v1alpha1.Hl7V2StoreIamMember, error)
	UpdateStatus(ctx context.Context, hl7V2StoreIamMember *v1alpha1.Hl7V2StoreIamMember, opts v1.UpdateOptions) (*v1alpha1.Hl7V2StoreIamMember, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Hl7V2StoreIamMember, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.Hl7V2StoreIamMemberList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Hl7V2StoreIamMember, err error)
	Hl7V2StoreIamMemberExpansion
}

// hl7V2StoreIamMembers implements Hl7V2StoreIamMemberInterface
type hl7V2StoreIamMembers struct {
	client rest.Interface
	ns     string
}

// newHl7V2StoreIamMembers returns a Hl7V2StoreIamMembers
func newHl7V2StoreIamMembers(c *HealthcareV1alpha1Client, namespace string) *hl7V2StoreIamMembers {
	return &hl7V2StoreIamMembers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the hl7V2StoreIamMember, and returns the corresponding hl7V2StoreIamMember object, and an error if there is any.
func (c *hl7V2StoreIamMembers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Hl7V2StoreIamMember, err error) {
	result = &v1alpha1.Hl7V2StoreIamMember{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hl7v2storeiammembers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Hl7V2StoreIamMembers that match those selectors.
func (c *hl7V2StoreIamMembers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.Hl7V2StoreIamMemberList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.Hl7V2StoreIamMemberList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hl7v2storeiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hl7V2StoreIamMembers.
func (c *hl7V2StoreIamMembers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("hl7v2storeiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a hl7V2StoreIamMember and creates it.  Returns the server's representation of the hl7V2StoreIamMember, and an error, if there is any.
func (c *hl7V2StoreIamMembers) Create(ctx context.Context, hl7V2StoreIamMember *v1alpha1.Hl7V2StoreIamMember, opts v1.CreateOptions) (result *v1alpha1.Hl7V2StoreIamMember, err error) {
	result = &v1alpha1.Hl7V2StoreIamMember{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("hl7v2storeiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hl7V2StoreIamMember).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a hl7V2StoreIamMember and updates it. Returns the server's representation of the hl7V2StoreIamMember, and an error, if there is any.
func (c *hl7V2StoreIamMembers) Update(ctx context.Context, hl7V2StoreIamMember *v1alpha1.Hl7V2StoreIamMember, opts v1.UpdateOptions) (result *v1alpha1.Hl7V2StoreIamMember, err error) {
	result = &v1alpha1.Hl7V2StoreIamMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hl7v2storeiammembers").
		Name(hl7V2StoreIamMember.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hl7V2StoreIamMember).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *hl7V2StoreIamMembers) UpdateStatus(ctx context.Context, hl7V2StoreIamMember *v1alpha1.Hl7V2StoreIamMember, opts v1.UpdateOptions) (result *v1alpha1.Hl7V2StoreIamMember, err error) {
	result = &v1alpha1.Hl7V2StoreIamMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hl7v2storeiammembers").
		Name(hl7V2StoreIamMember.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hl7V2StoreIamMember).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the hl7V2StoreIamMember and deletes it. Returns an error if one occurs.
func (c *hl7V2StoreIamMembers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hl7v2storeiammembers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hl7V2StoreIamMembers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hl7v2storeiammembers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched hl7V2StoreIamMember.
func (c *hl7V2StoreIamMembers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Hl7V2StoreIamMember, err error) {
	result = &v1alpha1.Hl7V2StoreIamMember{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("hl7v2storeiammembers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
