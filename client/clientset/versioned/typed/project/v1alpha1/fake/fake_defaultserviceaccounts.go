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

	v1alpha1 "kubeform.dev/provider-google-api/apis/project/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDefaultServiceAccountses implements DefaultServiceAccountsInterface
type FakeDefaultServiceAccountses struct {
	Fake *FakeProjectV1alpha1
	ns   string
}

var defaultserviceaccountsesResource = schema.GroupVersionResource{Group: "project.google.kubeform.com", Version: "v1alpha1", Resource: "defaultserviceaccountses"}

var defaultserviceaccountsesKind = schema.GroupVersionKind{Group: "project.google.kubeform.com", Version: "v1alpha1", Kind: "DefaultServiceAccounts"}

// Get takes name of the defaultServiceAccounts, and returns the corresponding defaultServiceAccounts object, and an error if there is any.
func (c *FakeDefaultServiceAccountses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DefaultServiceAccounts, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(defaultserviceaccountsesResource, c.ns, name), &v1alpha1.DefaultServiceAccounts{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DefaultServiceAccounts), err
}

// List takes label and field selectors, and returns the list of DefaultServiceAccountses that match those selectors.
func (c *FakeDefaultServiceAccountses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DefaultServiceAccountsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(defaultserviceaccountsesResource, defaultserviceaccountsesKind, c.ns, opts), &v1alpha1.DefaultServiceAccountsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DefaultServiceAccountsList{ListMeta: obj.(*v1alpha1.DefaultServiceAccountsList).ListMeta}
	for _, item := range obj.(*v1alpha1.DefaultServiceAccountsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested defaultServiceAccountses.
func (c *FakeDefaultServiceAccountses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(defaultserviceaccountsesResource, c.ns, opts))

}

// Create takes the representation of a defaultServiceAccounts and creates it.  Returns the server's representation of the defaultServiceAccounts, and an error, if there is any.
func (c *FakeDefaultServiceAccountses) Create(ctx context.Context, defaultServiceAccounts *v1alpha1.DefaultServiceAccounts, opts v1.CreateOptions) (result *v1alpha1.DefaultServiceAccounts, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(defaultserviceaccountsesResource, c.ns, defaultServiceAccounts), &v1alpha1.DefaultServiceAccounts{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DefaultServiceAccounts), err
}

// Update takes the representation of a defaultServiceAccounts and updates it. Returns the server's representation of the defaultServiceAccounts, and an error, if there is any.
func (c *FakeDefaultServiceAccountses) Update(ctx context.Context, defaultServiceAccounts *v1alpha1.DefaultServiceAccounts, opts v1.UpdateOptions) (result *v1alpha1.DefaultServiceAccounts, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(defaultserviceaccountsesResource, c.ns, defaultServiceAccounts), &v1alpha1.DefaultServiceAccounts{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DefaultServiceAccounts), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDefaultServiceAccountses) UpdateStatus(ctx context.Context, defaultServiceAccounts *v1alpha1.DefaultServiceAccounts, opts v1.UpdateOptions) (*v1alpha1.DefaultServiceAccounts, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(defaultserviceaccountsesResource, "status", c.ns, defaultServiceAccounts), &v1alpha1.DefaultServiceAccounts{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DefaultServiceAccounts), err
}

// Delete takes name of the defaultServiceAccounts and deletes it. Returns an error if one occurs.
func (c *FakeDefaultServiceAccountses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(defaultserviceaccountsesResource, c.ns, name), &v1alpha1.DefaultServiceAccounts{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDefaultServiceAccountses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(defaultserviceaccountsesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DefaultServiceAccountsList{})
	return err
}

// Patch applies the patch and returns the patched defaultServiceAccounts.
func (c *FakeDefaultServiceAccountses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DefaultServiceAccounts, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(defaultserviceaccountsesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DefaultServiceAccounts{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DefaultServiceAccounts), err
}
