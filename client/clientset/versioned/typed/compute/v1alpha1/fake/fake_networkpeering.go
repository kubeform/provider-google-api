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

// FakeNetworkPeerings implements NetworkPeeringInterface
type FakeNetworkPeerings struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var networkpeeringsResource = schema.GroupVersionResource{Group: "compute.google.kubeform.com", Version: "v1alpha1", Resource: "networkpeerings"}

var networkpeeringsKind = schema.GroupVersionKind{Group: "compute.google.kubeform.com", Version: "v1alpha1", Kind: "NetworkPeering"}

// Get takes name of the networkPeering, and returns the corresponding networkPeering object, and an error if there is any.
func (c *FakeNetworkPeerings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkPeering, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkpeeringsResource, c.ns, name), &v1alpha1.NetworkPeering{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkPeering), err
}

// List takes label and field selectors, and returns the list of NetworkPeerings that match those selectors.
func (c *FakeNetworkPeerings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetworkPeeringList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkpeeringsResource, networkpeeringsKind, c.ns, opts), &v1alpha1.NetworkPeeringList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkPeeringList{ListMeta: obj.(*v1alpha1.NetworkPeeringList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkPeeringList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkPeerings.
func (c *FakeNetworkPeerings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkpeeringsResource, c.ns, opts))

}

// Create takes the representation of a networkPeering and creates it.  Returns the server's representation of the networkPeering, and an error, if there is any.
func (c *FakeNetworkPeerings) Create(ctx context.Context, networkPeering *v1alpha1.NetworkPeering, opts v1.CreateOptions) (result *v1alpha1.NetworkPeering, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkpeeringsResource, c.ns, networkPeering), &v1alpha1.NetworkPeering{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkPeering), err
}

// Update takes the representation of a networkPeering and updates it. Returns the server's representation of the networkPeering, and an error, if there is any.
func (c *FakeNetworkPeerings) Update(ctx context.Context, networkPeering *v1alpha1.NetworkPeering, opts v1.UpdateOptions) (result *v1alpha1.NetworkPeering, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkpeeringsResource, c.ns, networkPeering), &v1alpha1.NetworkPeering{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkPeering), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkPeerings) UpdateStatus(ctx context.Context, networkPeering *v1alpha1.NetworkPeering, opts v1.UpdateOptions) (*v1alpha1.NetworkPeering, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkpeeringsResource, "status", c.ns, networkPeering), &v1alpha1.NetworkPeering{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkPeering), err
}

// Delete takes name of the networkPeering and deletes it. Returns an error if one occurs.
func (c *FakeNetworkPeerings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networkpeeringsResource, c.ns, name), &v1alpha1.NetworkPeering{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkPeerings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkpeeringsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkPeeringList{})
	return err
}

// Patch applies the patch and returns the patched networkPeering.
func (c *FakeNetworkPeerings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkPeering, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkpeeringsResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetworkPeering{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkPeering), err
}
