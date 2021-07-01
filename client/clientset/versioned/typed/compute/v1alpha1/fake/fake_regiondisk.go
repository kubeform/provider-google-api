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

// FakeRegionDisks implements RegionDiskInterface
type FakeRegionDisks struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var regiondisksResource = schema.GroupVersionResource{Group: "compute.google.kubeform.com", Version: "v1alpha1", Resource: "regiondisks"}

var regiondisksKind = schema.GroupVersionKind{Group: "compute.google.kubeform.com", Version: "v1alpha1", Kind: "RegionDisk"}

// Get takes name of the regionDisk, and returns the corresponding regionDisk object, and an error if there is any.
func (c *FakeRegionDisks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RegionDisk, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(regiondisksResource, c.ns, name), &v1alpha1.RegionDisk{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegionDisk), err
}

// List takes label and field selectors, and returns the list of RegionDisks that match those selectors.
func (c *FakeRegionDisks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RegionDiskList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(regiondisksResource, regiondisksKind, c.ns, opts), &v1alpha1.RegionDiskList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RegionDiskList{ListMeta: obj.(*v1alpha1.RegionDiskList).ListMeta}
	for _, item := range obj.(*v1alpha1.RegionDiskList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested regionDisks.
func (c *FakeRegionDisks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(regiondisksResource, c.ns, opts))

}

// Create takes the representation of a regionDisk and creates it.  Returns the server's representation of the regionDisk, and an error, if there is any.
func (c *FakeRegionDisks) Create(ctx context.Context, regionDisk *v1alpha1.RegionDisk, opts v1.CreateOptions) (result *v1alpha1.RegionDisk, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(regiondisksResource, c.ns, regionDisk), &v1alpha1.RegionDisk{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegionDisk), err
}

// Update takes the representation of a regionDisk and updates it. Returns the server's representation of the regionDisk, and an error, if there is any.
func (c *FakeRegionDisks) Update(ctx context.Context, regionDisk *v1alpha1.RegionDisk, opts v1.UpdateOptions) (result *v1alpha1.RegionDisk, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(regiondisksResource, c.ns, regionDisk), &v1alpha1.RegionDisk{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegionDisk), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRegionDisks) UpdateStatus(ctx context.Context, regionDisk *v1alpha1.RegionDisk, opts v1.UpdateOptions) (*v1alpha1.RegionDisk, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(regiondisksResource, "status", c.ns, regionDisk), &v1alpha1.RegionDisk{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegionDisk), err
}

// Delete takes name of the regionDisk and deletes it. Returns an error if one occurs.
func (c *FakeRegionDisks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(regiondisksResource, c.ns, name), &v1alpha1.RegionDisk{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRegionDisks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(regiondisksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RegionDiskList{})
	return err
}

// Patch applies the patch and returns the patched regionDisk.
func (c *FakeRegionDisks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RegionDisk, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(regiondisksResource, c.ns, name, pt, data, subresources...), &v1alpha1.RegionDisk{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegionDisk), err
}
