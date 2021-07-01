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

// FakeGlobalAddresses implements GlobalAddressInterface
type FakeGlobalAddresses struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var globaladdressesResource = schema.GroupVersionResource{Group: "compute.google.kubeform.com", Version: "v1alpha1", Resource: "globaladdresses"}

var globaladdressesKind = schema.GroupVersionKind{Group: "compute.google.kubeform.com", Version: "v1alpha1", Kind: "GlobalAddress"}

// Get takes name of the globalAddress, and returns the corresponding globalAddress object, and an error if there is any.
func (c *FakeGlobalAddresses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GlobalAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(globaladdressesResource, c.ns, name), &v1alpha1.GlobalAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlobalAddress), err
}

// List takes label and field selectors, and returns the list of GlobalAddresses that match those selectors.
func (c *FakeGlobalAddresses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GlobalAddressList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(globaladdressesResource, globaladdressesKind, c.ns, opts), &v1alpha1.GlobalAddressList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GlobalAddressList{ListMeta: obj.(*v1alpha1.GlobalAddressList).ListMeta}
	for _, item := range obj.(*v1alpha1.GlobalAddressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested globalAddresses.
func (c *FakeGlobalAddresses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(globaladdressesResource, c.ns, opts))

}

// Create takes the representation of a globalAddress and creates it.  Returns the server's representation of the globalAddress, and an error, if there is any.
func (c *FakeGlobalAddresses) Create(ctx context.Context, globalAddress *v1alpha1.GlobalAddress, opts v1.CreateOptions) (result *v1alpha1.GlobalAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(globaladdressesResource, c.ns, globalAddress), &v1alpha1.GlobalAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlobalAddress), err
}

// Update takes the representation of a globalAddress and updates it. Returns the server's representation of the globalAddress, and an error, if there is any.
func (c *FakeGlobalAddresses) Update(ctx context.Context, globalAddress *v1alpha1.GlobalAddress, opts v1.UpdateOptions) (result *v1alpha1.GlobalAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(globaladdressesResource, c.ns, globalAddress), &v1alpha1.GlobalAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlobalAddress), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGlobalAddresses) UpdateStatus(ctx context.Context, globalAddress *v1alpha1.GlobalAddress, opts v1.UpdateOptions) (*v1alpha1.GlobalAddress, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(globaladdressesResource, "status", c.ns, globalAddress), &v1alpha1.GlobalAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlobalAddress), err
}

// Delete takes name of the globalAddress and deletes it. Returns an error if one occurs.
func (c *FakeGlobalAddresses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(globaladdressesResource, c.ns, name), &v1alpha1.GlobalAddress{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGlobalAddresses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(globaladdressesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.GlobalAddressList{})
	return err
}

// Patch applies the patch and returns the patched globalAddress.
func (c *FakeGlobalAddresses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GlobalAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(globaladdressesResource, c.ns, name, pt, data, subresources...), &v1alpha1.GlobalAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlobalAddress), err
}
