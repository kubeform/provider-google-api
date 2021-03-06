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

// FakeAddresses implements AddressInterface
type FakeAddresses struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var addressesResource = schema.GroupVersionResource{Group: "compute.google.kubeform.com", Version: "v1alpha1", Resource: "addresses"}

var addressesKind = schema.GroupVersionKind{Group: "compute.google.kubeform.com", Version: "v1alpha1", Kind: "Address"}

// Get takes name of the address, and returns the corresponding address object, and an error if there is any.
func (c *FakeAddresses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Address, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(addressesResource, c.ns, name), &v1alpha1.Address{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Address), err
}

// List takes label and field selectors, and returns the list of Addresses that match those selectors.
func (c *FakeAddresses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AddressList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(addressesResource, addressesKind, c.ns, opts), &v1alpha1.AddressList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AddressList{ListMeta: obj.(*v1alpha1.AddressList).ListMeta}
	for _, item := range obj.(*v1alpha1.AddressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested addresses.
func (c *FakeAddresses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(addressesResource, c.ns, opts))

}

// Create takes the representation of a address and creates it.  Returns the server's representation of the address, and an error, if there is any.
func (c *FakeAddresses) Create(ctx context.Context, address *v1alpha1.Address, opts v1.CreateOptions) (result *v1alpha1.Address, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(addressesResource, c.ns, address), &v1alpha1.Address{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Address), err
}

// Update takes the representation of a address and updates it. Returns the server's representation of the address, and an error, if there is any.
func (c *FakeAddresses) Update(ctx context.Context, address *v1alpha1.Address, opts v1.UpdateOptions) (result *v1alpha1.Address, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(addressesResource, c.ns, address), &v1alpha1.Address{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Address), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAddresses) UpdateStatus(ctx context.Context, address *v1alpha1.Address, opts v1.UpdateOptions) (*v1alpha1.Address, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(addressesResource, "status", c.ns, address), &v1alpha1.Address{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Address), err
}

// Delete takes name of the address and deletes it. Returns an error if one occurs.
func (c *FakeAddresses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(addressesResource, c.ns, name), &v1alpha1.Address{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAddresses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(addressesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AddressList{})
	return err
}

// Patch applies the patch and returns the patched address.
func (c *FakeAddresses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Address, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(addressesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Address{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Address), err
}
