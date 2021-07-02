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

	v1alpha1 "kubeform.dev/provider-google-api/apis/billingaccount/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIamMembers implements IamMemberInterface
type FakeIamMembers struct {
	Fake *FakeBillingaccountV1alpha1
	ns   string
}

var iammembersResource = schema.GroupVersionResource{Group: "billingaccount.google.kubeform.com", Version: "v1alpha1", Resource: "iammembers"}

var iammembersKind = schema.GroupVersionKind{Group: "billingaccount.google.kubeform.com", Version: "v1alpha1", Kind: "IamMember"}

// Get takes name of the iamMember, and returns the corresponding iamMember object, and an error if there is any.
func (c *FakeIamMembers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iammembersResource, c.ns, name), &v1alpha1.IamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamMember), err
}

// List takes label and field selectors, and returns the list of IamMembers that match those selectors.
func (c *FakeIamMembers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iammembersResource, iammembersKind, c.ns, opts), &v1alpha1.IamMemberList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IamMemberList{ListMeta: obj.(*v1alpha1.IamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.IamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iamMembers.
func (c *FakeIamMembers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iammembersResource, c.ns, opts))

}

// Create takes the representation of a iamMember and creates it.  Returns the server's representation of the iamMember, and an error, if there is any.
func (c *FakeIamMembers) Create(ctx context.Context, iamMember *v1alpha1.IamMember, opts v1.CreateOptions) (result *v1alpha1.IamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iammembersResource, c.ns, iamMember), &v1alpha1.IamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamMember), err
}

// Update takes the representation of a iamMember and updates it. Returns the server's representation of the iamMember, and an error, if there is any.
func (c *FakeIamMembers) Update(ctx context.Context, iamMember *v1alpha1.IamMember, opts v1.UpdateOptions) (result *v1alpha1.IamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iammembersResource, c.ns, iamMember), &v1alpha1.IamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIamMembers) UpdateStatus(ctx context.Context, iamMember *v1alpha1.IamMember, opts v1.UpdateOptions) (*v1alpha1.IamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iammembersResource, "status", c.ns, iamMember), &v1alpha1.IamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamMember), err
}

// Delete takes name of the iamMember and deletes it. Returns an error if one occurs.
func (c *FakeIamMembers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iammembersResource, c.ns, name), &v1alpha1.IamMember{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIamMembers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iammembersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IamMemberList{})
	return err
}

// Patch applies the patch and returns the patched iamMember.
func (c *FakeIamMembers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iammembersResource, c.ns, name, pt, data, subresources...), &v1alpha1.IamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamMember), err
}
