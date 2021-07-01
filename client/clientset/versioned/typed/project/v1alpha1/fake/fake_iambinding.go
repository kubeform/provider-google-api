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
	v1alpha1 "kubeform.dev/provider-google-api/apis/project/v1alpha1"
)

// FakeIamBindings implements IamBindingInterface
type FakeIamBindings struct {
	Fake *FakeProjectV1alpha1
	ns   string
}

var iambindingsResource = schema.GroupVersionResource{Group: "project.google.kubeform.com", Version: "v1alpha1", Resource: "iambindings"}

var iambindingsKind = schema.GroupVersionKind{Group: "project.google.kubeform.com", Version: "v1alpha1", Kind: "IamBinding"}

// Get takes name of the iamBinding, and returns the corresponding iamBinding object, and an error if there is any.
func (c *FakeIamBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iambindingsResource, c.ns, name), &v1alpha1.IamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamBinding), err
}

// List takes label and field selectors, and returns the list of IamBindings that match those selectors.
func (c *FakeIamBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IamBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iambindingsResource, iambindingsKind, c.ns, opts), &v1alpha1.IamBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IamBindingList{ListMeta: obj.(*v1alpha1.IamBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.IamBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iamBindings.
func (c *FakeIamBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iambindingsResource, c.ns, opts))

}

// Create takes the representation of a iamBinding and creates it.  Returns the server's representation of the iamBinding, and an error, if there is any.
func (c *FakeIamBindings) Create(ctx context.Context, iamBinding *v1alpha1.IamBinding, opts v1.CreateOptions) (result *v1alpha1.IamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iambindingsResource, c.ns, iamBinding), &v1alpha1.IamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamBinding), err
}

// Update takes the representation of a iamBinding and updates it. Returns the server's representation of the iamBinding, and an error, if there is any.
func (c *FakeIamBindings) Update(ctx context.Context, iamBinding *v1alpha1.IamBinding, opts v1.UpdateOptions) (result *v1alpha1.IamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iambindingsResource, c.ns, iamBinding), &v1alpha1.IamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIamBindings) UpdateStatus(ctx context.Context, iamBinding *v1alpha1.IamBinding, opts v1.UpdateOptions) (*v1alpha1.IamBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iambindingsResource, "status", c.ns, iamBinding), &v1alpha1.IamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamBinding), err
}

// Delete takes name of the iamBinding and deletes it. Returns an error if one occurs.
func (c *FakeIamBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iambindingsResource, c.ns, name), &v1alpha1.IamBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIamBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iambindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IamBindingList{})
	return err
}

// Patch applies the patch and returns the patched iamBinding.
func (c *FakeIamBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iambindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.IamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamBinding), err
}
