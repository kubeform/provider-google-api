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

	v1alpha1 "kubeform.dev/provider-google-api/apis/bigtable/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTableIamMembers implements TableIamMemberInterface
type FakeTableIamMembers struct {
	Fake *FakeBigtableV1alpha1
	ns   string
}

var tableiammembersResource = schema.GroupVersionResource{Group: "bigtable.google.kubeform.com", Version: "v1alpha1", Resource: "tableiammembers"}

var tableiammembersKind = schema.GroupVersionKind{Group: "bigtable.google.kubeform.com", Version: "v1alpha1", Kind: "TableIamMember"}

// Get takes name of the tableIamMember, and returns the corresponding tableIamMember object, and an error if there is any.
func (c *FakeTableIamMembers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TableIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tableiammembersResource, c.ns, name), &v1alpha1.TableIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TableIamMember), err
}

// List takes label and field selectors, and returns the list of TableIamMembers that match those selectors.
func (c *FakeTableIamMembers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TableIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tableiammembersResource, tableiammembersKind, c.ns, opts), &v1alpha1.TableIamMemberList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TableIamMemberList{ListMeta: obj.(*v1alpha1.TableIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.TableIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tableIamMembers.
func (c *FakeTableIamMembers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tableiammembersResource, c.ns, opts))

}

// Create takes the representation of a tableIamMember and creates it.  Returns the server's representation of the tableIamMember, and an error, if there is any.
func (c *FakeTableIamMembers) Create(ctx context.Context, tableIamMember *v1alpha1.TableIamMember, opts v1.CreateOptions) (result *v1alpha1.TableIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tableiammembersResource, c.ns, tableIamMember), &v1alpha1.TableIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TableIamMember), err
}

// Update takes the representation of a tableIamMember and updates it. Returns the server's representation of the tableIamMember, and an error, if there is any.
func (c *FakeTableIamMembers) Update(ctx context.Context, tableIamMember *v1alpha1.TableIamMember, opts v1.UpdateOptions) (result *v1alpha1.TableIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tableiammembersResource, c.ns, tableIamMember), &v1alpha1.TableIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TableIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTableIamMembers) UpdateStatus(ctx context.Context, tableIamMember *v1alpha1.TableIamMember, opts v1.UpdateOptions) (*v1alpha1.TableIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(tableiammembersResource, "status", c.ns, tableIamMember), &v1alpha1.TableIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TableIamMember), err
}

// Delete takes name of the tableIamMember and deletes it. Returns an error if one occurs.
func (c *FakeTableIamMembers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tableiammembersResource, c.ns, name), &v1alpha1.TableIamMember{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTableIamMembers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tableiammembersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TableIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched tableIamMember.
func (c *FakeTableIamMembers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TableIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tableiammembersResource, c.ns, name, pt, data, subresources...), &v1alpha1.TableIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TableIamMember), err
}
