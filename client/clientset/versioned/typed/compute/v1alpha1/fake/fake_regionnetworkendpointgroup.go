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

// FakeRegionNetworkEndpointGroups implements RegionNetworkEndpointGroupInterface
type FakeRegionNetworkEndpointGroups struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var regionnetworkendpointgroupsResource = schema.GroupVersionResource{Group: "compute.google.kubeform.com", Version: "v1alpha1", Resource: "regionnetworkendpointgroups"}

var regionnetworkendpointgroupsKind = schema.GroupVersionKind{Group: "compute.google.kubeform.com", Version: "v1alpha1", Kind: "RegionNetworkEndpointGroup"}

// Get takes name of the regionNetworkEndpointGroup, and returns the corresponding regionNetworkEndpointGroup object, and an error if there is any.
func (c *FakeRegionNetworkEndpointGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RegionNetworkEndpointGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(regionnetworkendpointgroupsResource, c.ns, name), &v1alpha1.RegionNetworkEndpointGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegionNetworkEndpointGroup), err
}

// List takes label and field selectors, and returns the list of RegionNetworkEndpointGroups that match those selectors.
func (c *FakeRegionNetworkEndpointGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RegionNetworkEndpointGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(regionnetworkendpointgroupsResource, regionnetworkendpointgroupsKind, c.ns, opts), &v1alpha1.RegionNetworkEndpointGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RegionNetworkEndpointGroupList{ListMeta: obj.(*v1alpha1.RegionNetworkEndpointGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.RegionNetworkEndpointGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested regionNetworkEndpointGroups.
func (c *FakeRegionNetworkEndpointGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(regionnetworkendpointgroupsResource, c.ns, opts))

}

// Create takes the representation of a regionNetworkEndpointGroup and creates it.  Returns the server's representation of the regionNetworkEndpointGroup, and an error, if there is any.
func (c *FakeRegionNetworkEndpointGroups) Create(ctx context.Context, regionNetworkEndpointGroup *v1alpha1.RegionNetworkEndpointGroup, opts v1.CreateOptions) (result *v1alpha1.RegionNetworkEndpointGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(regionnetworkendpointgroupsResource, c.ns, regionNetworkEndpointGroup), &v1alpha1.RegionNetworkEndpointGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegionNetworkEndpointGroup), err
}

// Update takes the representation of a regionNetworkEndpointGroup and updates it. Returns the server's representation of the regionNetworkEndpointGroup, and an error, if there is any.
func (c *FakeRegionNetworkEndpointGroups) Update(ctx context.Context, regionNetworkEndpointGroup *v1alpha1.RegionNetworkEndpointGroup, opts v1.UpdateOptions) (result *v1alpha1.RegionNetworkEndpointGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(regionnetworkendpointgroupsResource, c.ns, regionNetworkEndpointGroup), &v1alpha1.RegionNetworkEndpointGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegionNetworkEndpointGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRegionNetworkEndpointGroups) UpdateStatus(ctx context.Context, regionNetworkEndpointGroup *v1alpha1.RegionNetworkEndpointGroup, opts v1.UpdateOptions) (*v1alpha1.RegionNetworkEndpointGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(regionnetworkendpointgroupsResource, "status", c.ns, regionNetworkEndpointGroup), &v1alpha1.RegionNetworkEndpointGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegionNetworkEndpointGroup), err
}

// Delete takes name of the regionNetworkEndpointGroup and deletes it. Returns an error if one occurs.
func (c *FakeRegionNetworkEndpointGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(regionnetworkendpointgroupsResource, c.ns, name), &v1alpha1.RegionNetworkEndpointGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRegionNetworkEndpointGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(regionnetworkendpointgroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RegionNetworkEndpointGroupList{})
	return err
}

// Patch applies the patch and returns the patched regionNetworkEndpointGroup.
func (c *FakeRegionNetworkEndpointGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RegionNetworkEndpointGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(regionnetworkendpointgroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.RegionNetworkEndpointGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegionNetworkEndpointGroup), err
}
