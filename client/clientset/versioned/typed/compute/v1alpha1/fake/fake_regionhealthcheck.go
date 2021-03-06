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

// FakeRegionHealthChecks implements RegionHealthCheckInterface
type FakeRegionHealthChecks struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var regionhealthchecksResource = schema.GroupVersionResource{Group: "compute.google.kubeform.com", Version: "v1alpha1", Resource: "regionhealthchecks"}

var regionhealthchecksKind = schema.GroupVersionKind{Group: "compute.google.kubeform.com", Version: "v1alpha1", Kind: "RegionHealthCheck"}

// Get takes name of the regionHealthCheck, and returns the corresponding regionHealthCheck object, and an error if there is any.
func (c *FakeRegionHealthChecks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RegionHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(regionhealthchecksResource, c.ns, name), &v1alpha1.RegionHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegionHealthCheck), err
}

// List takes label and field selectors, and returns the list of RegionHealthChecks that match those selectors.
func (c *FakeRegionHealthChecks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RegionHealthCheckList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(regionhealthchecksResource, regionhealthchecksKind, c.ns, opts), &v1alpha1.RegionHealthCheckList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RegionHealthCheckList{ListMeta: obj.(*v1alpha1.RegionHealthCheckList).ListMeta}
	for _, item := range obj.(*v1alpha1.RegionHealthCheckList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested regionHealthChecks.
func (c *FakeRegionHealthChecks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(regionhealthchecksResource, c.ns, opts))

}

// Create takes the representation of a regionHealthCheck and creates it.  Returns the server's representation of the regionHealthCheck, and an error, if there is any.
func (c *FakeRegionHealthChecks) Create(ctx context.Context, regionHealthCheck *v1alpha1.RegionHealthCheck, opts v1.CreateOptions) (result *v1alpha1.RegionHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(regionhealthchecksResource, c.ns, regionHealthCheck), &v1alpha1.RegionHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegionHealthCheck), err
}

// Update takes the representation of a regionHealthCheck and updates it. Returns the server's representation of the regionHealthCheck, and an error, if there is any.
func (c *FakeRegionHealthChecks) Update(ctx context.Context, regionHealthCheck *v1alpha1.RegionHealthCheck, opts v1.UpdateOptions) (result *v1alpha1.RegionHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(regionhealthchecksResource, c.ns, regionHealthCheck), &v1alpha1.RegionHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegionHealthCheck), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRegionHealthChecks) UpdateStatus(ctx context.Context, regionHealthCheck *v1alpha1.RegionHealthCheck, opts v1.UpdateOptions) (*v1alpha1.RegionHealthCheck, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(regionhealthchecksResource, "status", c.ns, regionHealthCheck), &v1alpha1.RegionHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegionHealthCheck), err
}

// Delete takes name of the regionHealthCheck and deletes it. Returns an error if one occurs.
func (c *FakeRegionHealthChecks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(regionhealthchecksResource, c.ns, name), &v1alpha1.RegionHealthCheck{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRegionHealthChecks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(regionhealthchecksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RegionHealthCheckList{})
	return err
}

// Patch applies the patch and returns the patched regionHealthCheck.
func (c *FakeRegionHealthChecks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RegionHealthCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(regionhealthchecksResource, c.ns, name, pt, data, subresources...), &v1alpha1.RegionHealthCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegionHealthCheck), err
}
