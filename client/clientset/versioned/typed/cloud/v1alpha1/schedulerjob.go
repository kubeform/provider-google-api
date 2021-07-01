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
	v1alpha1 "kubeform.dev/provider-google-api/apis/cloud/v1alpha1"
	scheme "kubeform.dev/provider-google-api/client/clientset/versioned/scheme"
)

// SchedulerJobsGetter has a method to return a SchedulerJobInterface.
// A group's client should implement this interface.
type SchedulerJobsGetter interface {
	SchedulerJobs(namespace string) SchedulerJobInterface
}

// SchedulerJobInterface has methods to work with SchedulerJob resources.
type SchedulerJobInterface interface {
	Create(ctx context.Context, schedulerJob *v1alpha1.SchedulerJob, opts v1.CreateOptions) (*v1alpha1.SchedulerJob, error)
	Update(ctx context.Context, schedulerJob *v1alpha1.SchedulerJob, opts v1.UpdateOptions) (*v1alpha1.SchedulerJob, error)
	UpdateStatus(ctx context.Context, schedulerJob *v1alpha1.SchedulerJob, opts v1.UpdateOptions) (*v1alpha1.SchedulerJob, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.SchedulerJob, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.SchedulerJobList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SchedulerJob, err error)
	SchedulerJobExpansion
}

// schedulerJobs implements SchedulerJobInterface
type schedulerJobs struct {
	client rest.Interface
	ns     string
}

// newSchedulerJobs returns a SchedulerJobs
func newSchedulerJobs(c *CloudV1alpha1Client, namespace string) *schedulerJobs {
	return &schedulerJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the schedulerJob, and returns the corresponding schedulerJob object, and an error if there is any.
func (c *schedulerJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SchedulerJob, err error) {
	result = &v1alpha1.SchedulerJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("schedulerjobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SchedulerJobs that match those selectors.
func (c *schedulerJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SchedulerJobList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SchedulerJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("schedulerjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested schedulerJobs.
func (c *schedulerJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("schedulerjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a schedulerJob and creates it.  Returns the server's representation of the schedulerJob, and an error, if there is any.
func (c *schedulerJobs) Create(ctx context.Context, schedulerJob *v1alpha1.SchedulerJob, opts v1.CreateOptions) (result *v1alpha1.SchedulerJob, err error) {
	result = &v1alpha1.SchedulerJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("schedulerjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(schedulerJob).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a schedulerJob and updates it. Returns the server's representation of the schedulerJob, and an error, if there is any.
func (c *schedulerJobs) Update(ctx context.Context, schedulerJob *v1alpha1.SchedulerJob, opts v1.UpdateOptions) (result *v1alpha1.SchedulerJob, err error) {
	result = &v1alpha1.SchedulerJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("schedulerjobs").
		Name(schedulerJob.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(schedulerJob).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *schedulerJobs) UpdateStatus(ctx context.Context, schedulerJob *v1alpha1.SchedulerJob, opts v1.UpdateOptions) (result *v1alpha1.SchedulerJob, err error) {
	result = &v1alpha1.SchedulerJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("schedulerjobs").
		Name(schedulerJob.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(schedulerJob).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the schedulerJob and deletes it. Returns an error if one occurs.
func (c *schedulerJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("schedulerjobs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *schedulerJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("schedulerjobs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched schedulerJob.
func (c *schedulerJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SchedulerJob, err error) {
	result = &v1alpha1.SchedulerJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("schedulerjobs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
