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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/provider-google-api/apis/compute/v1alpha1"
)

// FakeHttpHealthChecks implements HttpHealthCheckInterface
type FakeHttpHealthChecks struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var httphealthchecksResource = schema.GroupVersionResource{Group: "compute.google.kubeform.com", Version: "v1alpha1", Resource: "httphealthchecks"}

var httphealthchecksKind = schema.GroupVersionKind{Group: "compute.google.kubeform.com", Version: "v1alpha1", Kind: "HttpHealthCheck"}

// Get takes name of the httpHealthCheck, and returns the corresponding httpHealthCheck object, and an error if there is any.
func (c *FakeHttpHealthChecks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HttpHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(httphealthchecksResource, c.ns, name), &v1alpha1.HttpHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HttpHealthCheck), err
}

// List takes label and field selectors, and returns the list of HttpHealthChecks that match those selectors.
func (c *FakeHttpHealthChecks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HttpHealthCheckList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(httphealthchecksResource, httphealthchecksKind, c.ns, opts), &v1alpha1.HttpHealthCheckList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HttpHealthCheckList{ListMeta: obj.(*v1alpha1.HttpHealthCheckList).ListMeta}
	for _, item := range obj.(*v1alpha1.HttpHealthCheckList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested httpHealthChecks.
func (c *FakeHttpHealthChecks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(httphealthchecksResource, c.ns, opts))

}

// Create takes the representation of a httpHealthCheck and creates it.  Returns the server's representation of the httpHealthCheck, and an error, if there is any.
func (c *FakeHttpHealthChecks) Create(ctx context.Context, httpHealthCheck *v1alpha1.HttpHealthCheck, opts v1.CreateOptions) (result *v1alpha1.HttpHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(httphealthchecksResource, c.ns, httpHealthCheck), &v1alpha1.HttpHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HttpHealthCheck), err
}

// Update takes the representation of a httpHealthCheck and updates it. Returns the server's representation of the httpHealthCheck, and an error, if there is any.
func (c *FakeHttpHealthChecks) Update(ctx context.Context, httpHealthCheck *v1alpha1.HttpHealthCheck, opts v1.UpdateOptions) (result *v1alpha1.HttpHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(httphealthchecksResource, c.ns, httpHealthCheck), &v1alpha1.HttpHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HttpHealthCheck), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHttpHealthChecks) UpdateStatus(ctx context.Context, httpHealthCheck *v1alpha1.HttpHealthCheck, opts v1.UpdateOptions) (*v1alpha1.HttpHealthCheck, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(httphealthchecksResource, "status", c.ns, httpHealthCheck), &v1alpha1.HttpHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HttpHealthCheck), err
}

// Delete takes name of the httpHealthCheck and deletes it. Returns an error if one occurs.
func (c *FakeHttpHealthChecks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(httphealthchecksResource, c.ns, name), &v1alpha1.HttpHealthCheck{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHttpHealthChecks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(httphealthchecksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HttpHealthCheckList{})
	return err
}

// Patch applies the patch and returns the patched httpHealthCheck.
func (c *FakeHttpHealthChecks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HttpHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(httphealthchecksResource, c.ns, name, pt, data, subresources...), &v1alpha1.HttpHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HttpHealthCheck), err
}
