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

	v1alpha1 "kubeform.dev/provider-google-api/apis/data/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCatalogTagTemplateIamMembers implements CatalogTagTemplateIamMemberInterface
type FakeCatalogTagTemplateIamMembers struct {
	Fake *FakeDataV1alpha1
	ns   string
}

var catalogtagtemplateiammembersResource = schema.GroupVersionResource{Group: "data.google.kubeform.com", Version: "v1alpha1", Resource: "catalogtagtemplateiammembers"}

var catalogtagtemplateiammembersKind = schema.GroupVersionKind{Group: "data.google.kubeform.com", Version: "v1alpha1", Kind: "CatalogTagTemplateIamMember"}

// Get takes name of the catalogTagTemplateIamMember, and returns the corresponding catalogTagTemplateIamMember object, and an error if there is any.
func (c *FakeCatalogTagTemplateIamMembers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CatalogTagTemplateIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(catalogtagtemplateiammembersResource, c.ns, name), &v1alpha1.CatalogTagTemplateIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CatalogTagTemplateIamMember), err
}

// List takes label and field selectors, and returns the list of CatalogTagTemplateIamMembers that match those selectors.
func (c *FakeCatalogTagTemplateIamMembers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CatalogTagTemplateIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(catalogtagtemplateiammembersResource, catalogtagtemplateiammembersKind, c.ns, opts), &v1alpha1.CatalogTagTemplateIamMemberList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CatalogTagTemplateIamMemberList{ListMeta: obj.(*v1alpha1.CatalogTagTemplateIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.CatalogTagTemplateIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested catalogTagTemplateIamMembers.
func (c *FakeCatalogTagTemplateIamMembers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(catalogtagtemplateiammembersResource, c.ns, opts))

}

// Create takes the representation of a catalogTagTemplateIamMember and creates it.  Returns the server's representation of the catalogTagTemplateIamMember, and an error, if there is any.
func (c *FakeCatalogTagTemplateIamMembers) Create(ctx context.Context, catalogTagTemplateIamMember *v1alpha1.CatalogTagTemplateIamMember, opts v1.CreateOptions) (result *v1alpha1.CatalogTagTemplateIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(catalogtagtemplateiammembersResource, c.ns, catalogTagTemplateIamMember), &v1alpha1.CatalogTagTemplateIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CatalogTagTemplateIamMember), err
}

// Update takes the representation of a catalogTagTemplateIamMember and updates it. Returns the server's representation of the catalogTagTemplateIamMember, and an error, if there is any.
func (c *FakeCatalogTagTemplateIamMembers) Update(ctx context.Context, catalogTagTemplateIamMember *v1alpha1.CatalogTagTemplateIamMember, opts v1.UpdateOptions) (result *v1alpha1.CatalogTagTemplateIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(catalogtagtemplateiammembersResource, c.ns, catalogTagTemplateIamMember), &v1alpha1.CatalogTagTemplateIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CatalogTagTemplateIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCatalogTagTemplateIamMembers) UpdateStatus(ctx context.Context, catalogTagTemplateIamMember *v1alpha1.CatalogTagTemplateIamMember, opts v1.UpdateOptions) (*v1alpha1.CatalogTagTemplateIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(catalogtagtemplateiammembersResource, "status", c.ns, catalogTagTemplateIamMember), &v1alpha1.CatalogTagTemplateIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CatalogTagTemplateIamMember), err
}

// Delete takes name of the catalogTagTemplateIamMember and deletes it. Returns an error if one occurs.
func (c *FakeCatalogTagTemplateIamMembers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(catalogtagtemplateiammembersResource, c.ns, name), &v1alpha1.CatalogTagTemplateIamMember{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCatalogTagTemplateIamMembers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(catalogtagtemplateiammembersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CatalogTagTemplateIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched catalogTagTemplateIamMember.
func (c *FakeCatalogTagTemplateIamMembers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CatalogTagTemplateIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(catalogtagtemplateiammembersResource, c.ns, name, pt, data, subresources...), &v1alpha1.CatalogTagTemplateIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CatalogTagTemplateIamMember), err
}
