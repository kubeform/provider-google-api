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

	v1alpha1 "kubeform.dev/provider-google-api/apis/network/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConnectivityHubs implements ConnectivityHubInterface
type FakeConnectivityHubs struct {
	Fake *FakeNetworkV1alpha1
	ns   string
}

var connectivityhubsResource = schema.GroupVersionResource{Group: "network.google.kubeform.com", Version: "v1alpha1", Resource: "connectivityhubs"}

var connectivityhubsKind = schema.GroupVersionKind{Group: "network.google.kubeform.com", Version: "v1alpha1", Kind: "ConnectivityHub"}

// Get takes name of the connectivityHub, and returns the corresponding connectivityHub object, and an error if there is any.
func (c *FakeConnectivityHubs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ConnectivityHub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(connectivityhubsResource, c.ns, name), &v1alpha1.ConnectivityHub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConnectivityHub), err
}

// List takes label and field selectors, and returns the list of ConnectivityHubs that match those selectors.
func (c *FakeConnectivityHubs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ConnectivityHubList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(connectivityhubsResource, connectivityhubsKind, c.ns, opts), &v1alpha1.ConnectivityHubList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ConnectivityHubList{ListMeta: obj.(*v1alpha1.ConnectivityHubList).ListMeta}
	for _, item := range obj.(*v1alpha1.ConnectivityHubList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested connectivityHubs.
func (c *FakeConnectivityHubs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(connectivityhubsResource, c.ns, opts))

}

// Create takes the representation of a connectivityHub and creates it.  Returns the server's representation of the connectivityHub, and an error, if there is any.
func (c *FakeConnectivityHubs) Create(ctx context.Context, connectivityHub *v1alpha1.ConnectivityHub, opts v1.CreateOptions) (result *v1alpha1.ConnectivityHub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(connectivityhubsResource, c.ns, connectivityHub), &v1alpha1.ConnectivityHub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConnectivityHub), err
}

// Update takes the representation of a connectivityHub and updates it. Returns the server's representation of the connectivityHub, and an error, if there is any.
func (c *FakeConnectivityHubs) Update(ctx context.Context, connectivityHub *v1alpha1.ConnectivityHub, opts v1.UpdateOptions) (result *v1alpha1.ConnectivityHub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(connectivityhubsResource, c.ns, connectivityHub), &v1alpha1.ConnectivityHub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConnectivityHub), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeConnectivityHubs) UpdateStatus(ctx context.Context, connectivityHub *v1alpha1.ConnectivityHub, opts v1.UpdateOptions) (*v1alpha1.ConnectivityHub, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(connectivityhubsResource, "status", c.ns, connectivityHub), &v1alpha1.ConnectivityHub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConnectivityHub), err
}

// Delete takes name of the connectivityHub and deletes it. Returns an error if one occurs.
func (c *FakeConnectivityHubs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(connectivityhubsResource, c.ns, name), &v1alpha1.ConnectivityHub{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConnectivityHubs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(connectivityhubsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ConnectivityHubList{})
	return err
}

// Patch applies the patch and returns the patched connectivityHub.
func (c *FakeConnectivityHubs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ConnectivityHub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(connectivityhubsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ConnectivityHub{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConnectivityHub), err
}
