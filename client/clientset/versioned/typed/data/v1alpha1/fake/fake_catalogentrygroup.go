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

// FakeCatalogEntryGroups implements CatalogEntryGroupInterface
type FakeCatalogEntryGroups struct {
	Fake *FakeDataV1alpha1
	ns   string
}

var catalogentrygroupsResource = schema.GroupVersionResource{Group: "data.google.kubeform.com", Version: "v1alpha1", Resource: "catalogentrygroups"}

var catalogentrygroupsKind = schema.GroupVersionKind{Group: "data.google.kubeform.com", Version: "v1alpha1", Kind: "CatalogEntryGroup"}

// Get takes name of the catalogEntryGroup, and returns the corresponding catalogEntryGroup object, and an error if there is any.
func (c *FakeCatalogEntryGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CatalogEntryGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(catalogentrygroupsResource, c.ns, name), &v1alpha1.CatalogEntryGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CatalogEntryGroup), err
}

// List takes label and field selectors, and returns the list of CatalogEntryGroups that match those selectors.
func (c *FakeCatalogEntryGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CatalogEntryGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(catalogentrygroupsResource, catalogentrygroupsKind, c.ns, opts), &v1alpha1.CatalogEntryGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CatalogEntryGroupList{ListMeta: obj.(*v1alpha1.CatalogEntryGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.CatalogEntryGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested catalogEntryGroups.
func (c *FakeCatalogEntryGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(catalogentrygroupsResource, c.ns, opts))

}

// Create takes the representation of a catalogEntryGroup and creates it.  Returns the server's representation of the catalogEntryGroup, and an error, if there is any.
func (c *FakeCatalogEntryGroups) Create(ctx context.Context, catalogEntryGroup *v1alpha1.CatalogEntryGroup, opts v1.CreateOptions) (result *v1alpha1.CatalogEntryGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(catalogentrygroupsResource, c.ns, catalogEntryGroup), &v1alpha1.CatalogEntryGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CatalogEntryGroup), err
}

// Update takes the representation of a catalogEntryGroup and updates it. Returns the server's representation of the catalogEntryGroup, and an error, if there is any.
func (c *FakeCatalogEntryGroups) Update(ctx context.Context, catalogEntryGroup *v1alpha1.CatalogEntryGroup, opts v1.UpdateOptions) (result *v1alpha1.CatalogEntryGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(catalogentrygroupsResource, c.ns, catalogEntryGroup), &v1alpha1.CatalogEntryGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CatalogEntryGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCatalogEntryGroups) UpdateStatus(ctx context.Context, catalogEntryGroup *v1alpha1.CatalogEntryGroup, opts v1.UpdateOptions) (*v1alpha1.CatalogEntryGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(catalogentrygroupsResource, "status", c.ns, catalogEntryGroup), &v1alpha1.CatalogEntryGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CatalogEntryGroup), err
}

// Delete takes name of the catalogEntryGroup and deletes it. Returns an error if one occurs.
func (c *FakeCatalogEntryGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(catalogentrygroupsResource, c.ns, name), &v1alpha1.CatalogEntryGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCatalogEntryGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(catalogentrygroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CatalogEntryGroupList{})
	return err
}

// Patch applies the patch and returns the patched catalogEntryGroup.
func (c *FakeCatalogEntryGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CatalogEntryGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(catalogentrygroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CatalogEntryGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CatalogEntryGroup), err
}
