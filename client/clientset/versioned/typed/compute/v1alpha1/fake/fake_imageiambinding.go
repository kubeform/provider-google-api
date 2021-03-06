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

// FakeImageIamBindings implements ImageIamBindingInterface
type FakeImageIamBindings struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var imageiambindingsResource = schema.GroupVersionResource{Group: "compute.google.kubeform.com", Version: "v1alpha1", Resource: "imageiambindings"}

var imageiambindingsKind = schema.GroupVersionKind{Group: "compute.google.kubeform.com", Version: "v1alpha1", Kind: "ImageIamBinding"}

// Get takes name of the imageIamBinding, and returns the corresponding imageIamBinding object, and an error if there is any.
func (c *FakeImageIamBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ImageIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(imageiambindingsResource, c.ns, name), &v1alpha1.ImageIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageIamBinding), err
}

// List takes label and field selectors, and returns the list of ImageIamBindings that match those selectors.
func (c *FakeImageIamBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ImageIamBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(imageiambindingsResource, imageiambindingsKind, c.ns, opts), &v1alpha1.ImageIamBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ImageIamBindingList{ListMeta: obj.(*v1alpha1.ImageIamBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.ImageIamBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested imageIamBindings.
func (c *FakeImageIamBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(imageiambindingsResource, c.ns, opts))

}

// Create takes the representation of a imageIamBinding and creates it.  Returns the server's representation of the imageIamBinding, and an error, if there is any.
func (c *FakeImageIamBindings) Create(ctx context.Context, imageIamBinding *v1alpha1.ImageIamBinding, opts v1.CreateOptions) (result *v1alpha1.ImageIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(imageiambindingsResource, c.ns, imageIamBinding), &v1alpha1.ImageIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageIamBinding), err
}

// Update takes the representation of a imageIamBinding and updates it. Returns the server's representation of the imageIamBinding, and an error, if there is any.
func (c *FakeImageIamBindings) Update(ctx context.Context, imageIamBinding *v1alpha1.ImageIamBinding, opts v1.UpdateOptions) (result *v1alpha1.ImageIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(imageiambindingsResource, c.ns, imageIamBinding), &v1alpha1.ImageIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageIamBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeImageIamBindings) UpdateStatus(ctx context.Context, imageIamBinding *v1alpha1.ImageIamBinding, opts v1.UpdateOptions) (*v1alpha1.ImageIamBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(imageiambindingsResource, "status", c.ns, imageIamBinding), &v1alpha1.ImageIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageIamBinding), err
}

// Delete takes name of the imageIamBinding and deletes it. Returns an error if one occurs.
func (c *FakeImageIamBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(imageiambindingsResource, c.ns, name), &v1alpha1.ImageIamBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeImageIamBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(imageiambindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ImageIamBindingList{})
	return err
}

// Patch applies the patch and returns the patched imageIamBinding.
func (c *FakeImageIamBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ImageIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(imageiambindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ImageIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageIamBinding), err
}
