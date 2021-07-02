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

	v1alpha1 "kubeform.dev/provider-google-api/apis/compute/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHttpsHealthChecks implements HttpsHealthCheckInterface
type FakeHttpsHealthChecks struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var httpshealthchecksResource = schema.GroupVersionResource{Group: "compute.google.kubeform.com", Version: "v1alpha1", Resource: "httpshealthchecks"}

var httpshealthchecksKind = schema.GroupVersionKind{Group: "compute.google.kubeform.com", Version: "v1alpha1", Kind: "HttpsHealthCheck"}

// Get takes name of the httpsHealthCheck, and returns the corresponding httpsHealthCheck object, and an error if there is any.
func (c *FakeHttpsHealthChecks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HttpsHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(httpshealthchecksResource, c.ns, name), &v1alpha1.HttpsHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HttpsHealthCheck), err
}

// List takes label and field selectors, and returns the list of HttpsHealthChecks that match those selectors.
func (c *FakeHttpsHealthChecks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HttpsHealthCheckList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(httpshealthchecksResource, httpshealthchecksKind, c.ns, opts), &v1alpha1.HttpsHealthCheckList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HttpsHealthCheckList{ListMeta: obj.(*v1alpha1.HttpsHealthCheckList).ListMeta}
	for _, item := range obj.(*v1alpha1.HttpsHealthCheckList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested httpsHealthChecks.
func (c *FakeHttpsHealthChecks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(httpshealthchecksResource, c.ns, opts))

}

// Create takes the representation of a httpsHealthCheck and creates it.  Returns the server's representation of the httpsHealthCheck, and an error, if there is any.
func (c *FakeHttpsHealthChecks) Create(ctx context.Context, httpsHealthCheck *v1alpha1.HttpsHealthCheck, opts v1.CreateOptions) (result *v1alpha1.HttpsHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(httpshealthchecksResource, c.ns, httpsHealthCheck), &v1alpha1.HttpsHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HttpsHealthCheck), err
}

// Update takes the representation of a httpsHealthCheck and updates it. Returns the server's representation of the httpsHealthCheck, and an error, if there is any.
func (c *FakeHttpsHealthChecks) Update(ctx context.Context, httpsHealthCheck *v1alpha1.HttpsHealthCheck, opts v1.UpdateOptions) (result *v1alpha1.HttpsHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(httpshealthchecksResource, c.ns, httpsHealthCheck), &v1alpha1.HttpsHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HttpsHealthCheck), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHttpsHealthChecks) UpdateStatus(ctx context.Context, httpsHealthCheck *v1alpha1.HttpsHealthCheck, opts v1.UpdateOptions) (*v1alpha1.HttpsHealthCheck, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(httpshealthchecksResource, "status", c.ns, httpsHealthCheck), &v1alpha1.HttpsHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HttpsHealthCheck), err
}

// Delete takes name of the httpsHealthCheck and deletes it. Returns an error if one occurs.
func (c *FakeHttpsHealthChecks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(httpshealthchecksResource, c.ns, name), &v1alpha1.HttpsHealthCheck{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHttpsHealthChecks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(httpshealthchecksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HttpsHealthCheckList{})
	return err
}

// Patch applies the patch and returns the patched httpsHealthCheck.
func (c *FakeHttpsHealthChecks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HttpsHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(httpshealthchecksResource, c.ns, name, pt, data, subresources...), &v1alpha1.HttpsHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HttpsHealthCheck), err
}
