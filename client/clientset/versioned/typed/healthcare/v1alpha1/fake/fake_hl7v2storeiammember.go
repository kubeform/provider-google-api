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
	v1alpha1 "kubeform.dev/provider-google-api/apis/healthcare/v1alpha1"
)

// FakeHl7V2StoreIamMembers implements Hl7V2StoreIamMemberInterface
type FakeHl7V2StoreIamMembers struct {
	Fake *FakeHealthcareV1alpha1
	ns   string
}

var hl7v2storeiammembersResource = schema.GroupVersionResource{Group: "healthcare.google.kubeform.com", Version: "v1alpha1", Resource: "hl7v2storeiammembers"}

var hl7v2storeiammembersKind = schema.GroupVersionKind{Group: "healthcare.google.kubeform.com", Version: "v1alpha1", Kind: "Hl7V2StoreIamMember"}

// Get takes name of the hl7V2StoreIamMember, and returns the corresponding hl7V2StoreIamMember object, and an error if there is any.
func (c *FakeHl7V2StoreIamMembers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Hl7V2StoreIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hl7v2storeiammembersResource, c.ns, name), &v1alpha1.Hl7V2StoreIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Hl7V2StoreIamMember), err
}

// List takes label and field selectors, and returns the list of Hl7V2StoreIamMembers that match those selectors.
func (c *FakeHl7V2StoreIamMembers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.Hl7V2StoreIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hl7v2storeiammembersResource, hl7v2storeiammembersKind, c.ns, opts), &v1alpha1.Hl7V2StoreIamMemberList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.Hl7V2StoreIamMemberList{ListMeta: obj.(*v1alpha1.Hl7V2StoreIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.Hl7V2StoreIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hl7V2StoreIamMembers.
func (c *FakeHl7V2StoreIamMembers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hl7v2storeiammembersResource, c.ns, opts))

}

// Create takes the representation of a hl7V2StoreIamMember and creates it.  Returns the server's representation of the hl7V2StoreIamMember, and an error, if there is any.
func (c *FakeHl7V2StoreIamMembers) Create(ctx context.Context, hl7V2StoreIamMember *v1alpha1.Hl7V2StoreIamMember, opts v1.CreateOptions) (result *v1alpha1.Hl7V2StoreIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hl7v2storeiammembersResource, c.ns, hl7V2StoreIamMember), &v1alpha1.Hl7V2StoreIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Hl7V2StoreIamMember), err
}

// Update takes the representation of a hl7V2StoreIamMember and updates it. Returns the server's representation of the hl7V2StoreIamMember, and an error, if there is any.
func (c *FakeHl7V2StoreIamMembers) Update(ctx context.Context, hl7V2StoreIamMember *v1alpha1.Hl7V2StoreIamMember, opts v1.UpdateOptions) (result *v1alpha1.Hl7V2StoreIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hl7v2storeiammembersResource, c.ns, hl7V2StoreIamMember), &v1alpha1.Hl7V2StoreIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Hl7V2StoreIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHl7V2StoreIamMembers) UpdateStatus(ctx context.Context, hl7V2StoreIamMember *v1alpha1.Hl7V2StoreIamMember, opts v1.UpdateOptions) (*v1alpha1.Hl7V2StoreIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(hl7v2storeiammembersResource, "status", c.ns, hl7V2StoreIamMember), &v1alpha1.Hl7V2StoreIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Hl7V2StoreIamMember), err
}

// Delete takes name of the hl7V2StoreIamMember and deletes it. Returns an error if one occurs.
func (c *FakeHl7V2StoreIamMembers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(hl7v2storeiammembersResource, c.ns, name), &v1alpha1.Hl7V2StoreIamMember{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHl7V2StoreIamMembers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hl7v2storeiammembersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.Hl7V2StoreIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched hl7V2StoreIamMember.
func (c *FakeHl7V2StoreIamMembers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Hl7V2StoreIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hl7v2storeiammembersResource, c.ns, name, pt, data, subresources...), &v1alpha1.Hl7V2StoreIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Hl7V2StoreIamMember), err
}
