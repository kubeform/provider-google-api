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

	v1alpha1 "kubeform.dev/provider-google-api/apis/service/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAccountIamMembers implements AccountIamMemberInterface
type FakeAccountIamMembers struct {
	Fake *FakeServiceV1alpha1
	ns   string
}

var accountiammembersResource = schema.GroupVersionResource{Group: "service.google.kubeform.com", Version: "v1alpha1", Resource: "accountiammembers"}

var accountiammembersKind = schema.GroupVersionKind{Group: "service.google.kubeform.com", Version: "v1alpha1", Kind: "AccountIamMember"}

// Get takes name of the accountIamMember, and returns the corresponding accountIamMember object, and an error if there is any.
func (c *FakeAccountIamMembers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AccountIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(accountiammembersResource, c.ns, name), &v1alpha1.AccountIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountIamMember), err
}

// List takes label and field selectors, and returns the list of AccountIamMembers that match those selectors.
func (c *FakeAccountIamMembers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AccountIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(accountiammembersResource, accountiammembersKind, c.ns, opts), &v1alpha1.AccountIamMemberList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AccountIamMemberList{ListMeta: obj.(*v1alpha1.AccountIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.AccountIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested accountIamMembers.
func (c *FakeAccountIamMembers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(accountiammembersResource, c.ns, opts))

}

// Create takes the representation of a accountIamMember and creates it.  Returns the server's representation of the accountIamMember, and an error, if there is any.
func (c *FakeAccountIamMembers) Create(ctx context.Context, accountIamMember *v1alpha1.AccountIamMember, opts v1.CreateOptions) (result *v1alpha1.AccountIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(accountiammembersResource, c.ns, accountIamMember), &v1alpha1.AccountIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountIamMember), err
}

// Update takes the representation of a accountIamMember and updates it. Returns the server's representation of the accountIamMember, and an error, if there is any.
func (c *FakeAccountIamMembers) Update(ctx context.Context, accountIamMember *v1alpha1.AccountIamMember, opts v1.UpdateOptions) (result *v1alpha1.AccountIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(accountiammembersResource, c.ns, accountIamMember), &v1alpha1.AccountIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAccountIamMembers) UpdateStatus(ctx context.Context, accountIamMember *v1alpha1.AccountIamMember, opts v1.UpdateOptions) (*v1alpha1.AccountIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(accountiammembersResource, "status", c.ns, accountIamMember), &v1alpha1.AccountIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountIamMember), err
}

// Delete takes name of the accountIamMember and deletes it. Returns an error if one occurs.
func (c *FakeAccountIamMembers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(accountiammembersResource, c.ns, name), &v1alpha1.AccountIamMember{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAccountIamMembers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(accountiammembersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AccountIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched accountIamMember.
func (c *FakeAccountIamMembers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AccountIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(accountiammembersResource, c.ns, name, pt, data, subresources...), &v1alpha1.AccountIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountIamMember), err
}
