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

	v1alpha1 "kubeform.dev/provider-google-api/apis/tags/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTagKeyIamBindings implements TagKeyIamBindingInterface
type FakeTagKeyIamBindings struct {
	Fake *FakeTagsV1alpha1
	ns   string
}

var tagkeyiambindingsResource = schema.GroupVersionResource{Group: "tags.google.kubeform.com", Version: "v1alpha1", Resource: "tagkeyiambindings"}

var tagkeyiambindingsKind = schema.GroupVersionKind{Group: "tags.google.kubeform.com", Version: "v1alpha1", Kind: "TagKeyIamBinding"}

// Get takes name of the tagKeyIamBinding, and returns the corresponding tagKeyIamBinding object, and an error if there is any.
func (c *FakeTagKeyIamBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TagKeyIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tagkeyiambindingsResource, c.ns, name), &v1alpha1.TagKeyIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TagKeyIamBinding), err
}

// List takes label and field selectors, and returns the list of TagKeyIamBindings that match those selectors.
func (c *FakeTagKeyIamBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TagKeyIamBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tagkeyiambindingsResource, tagkeyiambindingsKind, c.ns, opts), &v1alpha1.TagKeyIamBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TagKeyIamBindingList{ListMeta: obj.(*v1alpha1.TagKeyIamBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.TagKeyIamBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tagKeyIamBindings.
func (c *FakeTagKeyIamBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tagkeyiambindingsResource, c.ns, opts))

}

// Create takes the representation of a tagKeyIamBinding and creates it.  Returns the server's representation of the tagKeyIamBinding, and an error, if there is any.
func (c *FakeTagKeyIamBindings) Create(ctx context.Context, tagKeyIamBinding *v1alpha1.TagKeyIamBinding, opts v1.CreateOptions) (result *v1alpha1.TagKeyIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tagkeyiambindingsResource, c.ns, tagKeyIamBinding), &v1alpha1.TagKeyIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TagKeyIamBinding), err
}

// Update takes the representation of a tagKeyIamBinding and updates it. Returns the server's representation of the tagKeyIamBinding, and an error, if there is any.
func (c *FakeTagKeyIamBindings) Update(ctx context.Context, tagKeyIamBinding *v1alpha1.TagKeyIamBinding, opts v1.UpdateOptions) (result *v1alpha1.TagKeyIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tagkeyiambindingsResource, c.ns, tagKeyIamBinding), &v1alpha1.TagKeyIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TagKeyIamBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTagKeyIamBindings) UpdateStatus(ctx context.Context, tagKeyIamBinding *v1alpha1.TagKeyIamBinding, opts v1.UpdateOptions) (*v1alpha1.TagKeyIamBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(tagkeyiambindingsResource, "status", c.ns, tagKeyIamBinding), &v1alpha1.TagKeyIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TagKeyIamBinding), err
}

// Delete takes name of the tagKeyIamBinding and deletes it. Returns an error if one occurs.
func (c *FakeTagKeyIamBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tagkeyiambindingsResource, c.ns, name), &v1alpha1.TagKeyIamBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTagKeyIamBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tagkeyiambindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TagKeyIamBindingList{})
	return err
}

// Patch applies the patch and returns the patched tagKeyIamBinding.
func (c *FakeTagKeyIamBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TagKeyIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tagkeyiambindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.TagKeyIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TagKeyIamBinding), err
}
