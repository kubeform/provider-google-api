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
	v1alpha1 "kubeform.dev/provider-google-api/apis/endpoints/v1alpha1"
)

// FakeServiceIamMembers implements ServiceIamMemberInterface
type FakeServiceIamMembers struct {
	Fake *FakeEndpointsV1alpha1
	ns   string
}

var serviceiammembersResource = schema.GroupVersionResource{Group: "endpoints.google.kubeform.com", Version: "v1alpha1", Resource: "serviceiammembers"}

var serviceiammembersKind = schema.GroupVersionKind{Group: "endpoints.google.kubeform.com", Version: "v1alpha1", Kind: "ServiceIamMember"}

// Get takes name of the serviceIamMember, and returns the corresponding serviceIamMember object, and an error if there is any.
func (c *FakeServiceIamMembers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServiceIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(serviceiammembersResource, c.ns, name), &v1alpha1.ServiceIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceIamMember), err
}

// List takes label and field selectors, and returns the list of ServiceIamMembers that match those selectors.
func (c *FakeServiceIamMembers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServiceIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(serviceiammembersResource, serviceiammembersKind, c.ns, opts), &v1alpha1.ServiceIamMemberList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServiceIamMemberList{ListMeta: obj.(*v1alpha1.ServiceIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServiceIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceIamMembers.
func (c *FakeServiceIamMembers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(serviceiammembersResource, c.ns, opts))

}

// Create takes the representation of a serviceIamMember and creates it.  Returns the server's representation of the serviceIamMember, and an error, if there is any.
func (c *FakeServiceIamMembers) Create(ctx context.Context, serviceIamMember *v1alpha1.ServiceIamMember, opts v1.CreateOptions) (result *v1alpha1.ServiceIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(serviceiammembersResource, c.ns, serviceIamMember), &v1alpha1.ServiceIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceIamMember), err
}

// Update takes the representation of a serviceIamMember and updates it. Returns the server's representation of the serviceIamMember, and an error, if there is any.
func (c *FakeServiceIamMembers) Update(ctx context.Context, serviceIamMember *v1alpha1.ServiceIamMember, opts v1.UpdateOptions) (result *v1alpha1.ServiceIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(serviceiammembersResource, c.ns, serviceIamMember), &v1alpha1.ServiceIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceIamMembers) UpdateStatus(ctx context.Context, serviceIamMember *v1alpha1.ServiceIamMember, opts v1.UpdateOptions) (*v1alpha1.ServiceIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(serviceiammembersResource, "status", c.ns, serviceIamMember), &v1alpha1.ServiceIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceIamMember), err
}

// Delete takes name of the serviceIamMember and deletes it. Returns an error if one occurs.
func (c *FakeServiceIamMembers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(serviceiammembersResource, c.ns, name), &v1alpha1.ServiceIamMember{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceIamMembers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(serviceiammembersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServiceIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched serviceIamMember.
func (c *FakeServiceIamMembers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(serviceiammembersResource, c.ns, name, pt, data, subresources...), &v1alpha1.ServiceIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceIamMember), err
}
