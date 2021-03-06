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

// FakeNetworkEndpointGroups implements NetworkEndpointGroupInterface
type FakeNetworkEndpointGroups struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var networkendpointgroupsResource = schema.GroupVersionResource{Group: "compute.google.kubeform.com", Version: "v1alpha1", Resource: "networkendpointgroups"}

var networkendpointgroupsKind = schema.GroupVersionKind{Group: "compute.google.kubeform.com", Version: "v1alpha1", Kind: "NetworkEndpointGroup"}

// Get takes name of the networkEndpointGroup, and returns the corresponding networkEndpointGroup object, and an error if there is any.
func (c *FakeNetworkEndpointGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkEndpointGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkendpointgroupsResource, c.ns, name), &v1alpha1.NetworkEndpointGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkEndpointGroup), err
}

// List takes label and field selectors, and returns the list of NetworkEndpointGroups that match those selectors.
func (c *FakeNetworkEndpointGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetworkEndpointGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkendpointgroupsResource, networkendpointgroupsKind, c.ns, opts), &v1alpha1.NetworkEndpointGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkEndpointGroupList{ListMeta: obj.(*v1alpha1.NetworkEndpointGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkEndpointGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkEndpointGroups.
func (c *FakeNetworkEndpointGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkendpointgroupsResource, c.ns, opts))

}

// Create takes the representation of a networkEndpointGroup and creates it.  Returns the server's representation of the networkEndpointGroup, and an error, if there is any.
func (c *FakeNetworkEndpointGroups) Create(ctx context.Context, networkEndpointGroup *v1alpha1.NetworkEndpointGroup, opts v1.CreateOptions) (result *v1alpha1.NetworkEndpointGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkendpointgroupsResource, c.ns, networkEndpointGroup), &v1alpha1.NetworkEndpointGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkEndpointGroup), err
}

// Update takes the representation of a networkEndpointGroup and updates it. Returns the server's representation of the networkEndpointGroup, and an error, if there is any.
func (c *FakeNetworkEndpointGroups) Update(ctx context.Context, networkEndpointGroup *v1alpha1.NetworkEndpointGroup, opts v1.UpdateOptions) (result *v1alpha1.NetworkEndpointGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkendpointgroupsResource, c.ns, networkEndpointGroup), &v1alpha1.NetworkEndpointGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkEndpointGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkEndpointGroups) UpdateStatus(ctx context.Context, networkEndpointGroup *v1alpha1.NetworkEndpointGroup, opts v1.UpdateOptions) (*v1alpha1.NetworkEndpointGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkendpointgroupsResource, "status", c.ns, networkEndpointGroup), &v1alpha1.NetworkEndpointGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkEndpointGroup), err
}

// Delete takes name of the networkEndpointGroup and deletes it. Returns an error if one occurs.
func (c *FakeNetworkEndpointGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networkendpointgroupsResource, c.ns, name), &v1alpha1.NetworkEndpointGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkEndpointGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkendpointgroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkEndpointGroupList{})
	return err
}

// Patch applies the patch and returns the patched networkEndpointGroup.
func (c *FakeNetworkEndpointGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkEndpointGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkendpointgroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetworkEndpointGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkEndpointGroup), err
}
