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

// FakeCatalogTagTemplateIamBindings implements CatalogTagTemplateIamBindingInterface
type FakeCatalogTagTemplateIamBindings struct {
	Fake *FakeDataV1alpha1
	ns   string
}

var catalogtagtemplateiambindingsResource = schema.GroupVersionResource{Group: "data.google.kubeform.com", Version: "v1alpha1", Resource: "catalogtagtemplateiambindings"}

var catalogtagtemplateiambindingsKind = schema.GroupVersionKind{Group: "data.google.kubeform.com", Version: "v1alpha1", Kind: "CatalogTagTemplateIamBinding"}

// Get takes name of the catalogTagTemplateIamBinding, and returns the corresponding catalogTagTemplateIamBinding object, and an error if there is any.
func (c *FakeCatalogTagTemplateIamBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CatalogTagTemplateIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(catalogtagtemplateiambindingsResource, c.ns, name), &v1alpha1.CatalogTagTemplateIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CatalogTagTemplateIamBinding), err
}

// List takes label and field selectors, and returns the list of CatalogTagTemplateIamBindings that match those selectors.
func (c *FakeCatalogTagTemplateIamBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CatalogTagTemplateIamBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(catalogtagtemplateiambindingsResource, catalogtagtemplateiambindingsKind, c.ns, opts), &v1alpha1.CatalogTagTemplateIamBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CatalogTagTemplateIamBindingList{ListMeta: obj.(*v1alpha1.CatalogTagTemplateIamBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.CatalogTagTemplateIamBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested catalogTagTemplateIamBindings.
func (c *FakeCatalogTagTemplateIamBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(catalogtagtemplateiambindingsResource, c.ns, opts))

}

// Create takes the representation of a catalogTagTemplateIamBinding and creates it.  Returns the server's representation of the catalogTagTemplateIamBinding, and an error, if there is any.
func (c *FakeCatalogTagTemplateIamBindings) Create(ctx context.Context, catalogTagTemplateIamBinding *v1alpha1.CatalogTagTemplateIamBinding, opts v1.CreateOptions) (result *v1alpha1.CatalogTagTemplateIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(catalogtagtemplateiambindingsResource, c.ns, catalogTagTemplateIamBinding), &v1alpha1.CatalogTagTemplateIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CatalogTagTemplateIamBinding), err
}

// Update takes the representation of a catalogTagTemplateIamBinding and updates it. Returns the server's representation of the catalogTagTemplateIamBinding, and an error, if there is any.
func (c *FakeCatalogTagTemplateIamBindings) Update(ctx context.Context, catalogTagTemplateIamBinding *v1alpha1.CatalogTagTemplateIamBinding, opts v1.UpdateOptions) (result *v1alpha1.CatalogTagTemplateIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(catalogtagtemplateiambindingsResource, c.ns, catalogTagTemplateIamBinding), &v1alpha1.CatalogTagTemplateIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CatalogTagTemplateIamBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCatalogTagTemplateIamBindings) UpdateStatus(ctx context.Context, catalogTagTemplateIamBinding *v1alpha1.CatalogTagTemplateIamBinding, opts v1.UpdateOptions) (*v1alpha1.CatalogTagTemplateIamBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(catalogtagtemplateiambindingsResource, "status", c.ns, catalogTagTemplateIamBinding), &v1alpha1.CatalogTagTemplateIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CatalogTagTemplateIamBinding), err
}

// Delete takes name of the catalogTagTemplateIamBinding and deletes it. Returns an error if one occurs.
func (c *FakeCatalogTagTemplateIamBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(catalogtagtemplateiambindingsResource, c.ns, name), &v1alpha1.CatalogTagTemplateIamBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCatalogTagTemplateIamBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(catalogtagtemplateiambindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CatalogTagTemplateIamBindingList{})
	return err
}

// Patch applies the patch and returns the patched catalogTagTemplateIamBinding.
func (c *FakeCatalogTagTemplateIamBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CatalogTagTemplateIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(catalogtagtemplateiambindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CatalogTagTemplateIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CatalogTagTemplateIamBinding), err
}
