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
	v1alpha1 "kubeform.dev/provider-google-api/apis/pubsub/v1alpha1"
)

// FakeSubscriptionIamMembers implements SubscriptionIamMemberInterface
type FakeSubscriptionIamMembers struct {
	Fake *FakePubsubV1alpha1
	ns   string
}

var subscriptioniammembersResource = schema.GroupVersionResource{Group: "pubsub.google.kubeform.com", Version: "v1alpha1", Resource: "subscriptioniammembers"}

var subscriptioniammembersKind = schema.GroupVersionKind{Group: "pubsub.google.kubeform.com", Version: "v1alpha1", Kind: "SubscriptionIamMember"}

// Get takes name of the subscriptionIamMember, and returns the corresponding subscriptionIamMember object, and an error if there is any.
func (c *FakeSubscriptionIamMembers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SubscriptionIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(subscriptioniammembersResource, c.ns, name), &v1alpha1.SubscriptionIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubscriptionIamMember), err
}

// List takes label and field selectors, and returns the list of SubscriptionIamMembers that match those selectors.
func (c *FakeSubscriptionIamMembers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SubscriptionIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(subscriptioniammembersResource, subscriptioniammembersKind, c.ns, opts), &v1alpha1.SubscriptionIamMemberList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SubscriptionIamMemberList{ListMeta: obj.(*v1alpha1.SubscriptionIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.SubscriptionIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested subscriptionIamMembers.
func (c *FakeSubscriptionIamMembers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(subscriptioniammembersResource, c.ns, opts))

}

// Create takes the representation of a subscriptionIamMember and creates it.  Returns the server's representation of the subscriptionIamMember, and an error, if there is any.
func (c *FakeSubscriptionIamMembers) Create(ctx context.Context, subscriptionIamMember *v1alpha1.SubscriptionIamMember, opts v1.CreateOptions) (result *v1alpha1.SubscriptionIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(subscriptioniammembersResource, c.ns, subscriptionIamMember), &v1alpha1.SubscriptionIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubscriptionIamMember), err
}

// Update takes the representation of a subscriptionIamMember and updates it. Returns the server's representation of the subscriptionIamMember, and an error, if there is any.
func (c *FakeSubscriptionIamMembers) Update(ctx context.Context, subscriptionIamMember *v1alpha1.SubscriptionIamMember, opts v1.UpdateOptions) (result *v1alpha1.SubscriptionIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(subscriptioniammembersResource, c.ns, subscriptionIamMember), &v1alpha1.SubscriptionIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubscriptionIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSubscriptionIamMembers) UpdateStatus(ctx context.Context, subscriptionIamMember *v1alpha1.SubscriptionIamMember, opts v1.UpdateOptions) (*v1alpha1.SubscriptionIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(subscriptioniammembersResource, "status", c.ns, subscriptionIamMember), &v1alpha1.SubscriptionIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubscriptionIamMember), err
}

// Delete takes name of the subscriptionIamMember and deletes it. Returns an error if one occurs.
func (c *FakeSubscriptionIamMembers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(subscriptioniammembersResource, c.ns, name), &v1alpha1.SubscriptionIamMember{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSubscriptionIamMembers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(subscriptioniammembersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SubscriptionIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched subscriptionIamMember.
func (c *FakeSubscriptionIamMembers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SubscriptionIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(subscriptioniammembersResource, c.ns, name, pt, data, subresources...), &v1alpha1.SubscriptionIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubscriptionIamMember), err
}
