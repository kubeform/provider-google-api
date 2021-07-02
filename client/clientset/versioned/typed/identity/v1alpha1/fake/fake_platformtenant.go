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

	v1alpha1 "kubeform.dev/provider-google-api/apis/identity/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePlatformTenants implements PlatformTenantInterface
type FakePlatformTenants struct {
	Fake *FakeIdentityV1alpha1
	ns   string
}

var platformtenantsResource = schema.GroupVersionResource{Group: "identity.google.kubeform.com", Version: "v1alpha1", Resource: "platformtenants"}

var platformtenantsKind = schema.GroupVersionKind{Group: "identity.google.kubeform.com", Version: "v1alpha1", Kind: "PlatformTenant"}

// Get takes name of the platformTenant, and returns the corresponding platformTenant object, and an error if there is any.
func (c *FakePlatformTenants) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PlatformTenant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(platformtenantsResource, c.ns, name), &v1alpha1.PlatformTenant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PlatformTenant), err
}

// List takes label and field selectors, and returns the list of PlatformTenants that match those selectors.
func (c *FakePlatformTenants) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PlatformTenantList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(platformtenantsResource, platformtenantsKind, c.ns, opts), &v1alpha1.PlatformTenantList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PlatformTenantList{ListMeta: obj.(*v1alpha1.PlatformTenantList).ListMeta}
	for _, item := range obj.(*v1alpha1.PlatformTenantList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested platformTenants.
func (c *FakePlatformTenants) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(platformtenantsResource, c.ns, opts))

}

// Create takes the representation of a platformTenant and creates it.  Returns the server's representation of the platformTenant, and an error, if there is any.
func (c *FakePlatformTenants) Create(ctx context.Context, platformTenant *v1alpha1.PlatformTenant, opts v1.CreateOptions) (result *v1alpha1.PlatformTenant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(platformtenantsResource, c.ns, platformTenant), &v1alpha1.PlatformTenant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PlatformTenant), err
}

// Update takes the representation of a platformTenant and updates it. Returns the server's representation of the platformTenant, and an error, if there is any.
func (c *FakePlatformTenants) Update(ctx context.Context, platformTenant *v1alpha1.PlatformTenant, opts v1.UpdateOptions) (result *v1alpha1.PlatformTenant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(platformtenantsResource, c.ns, platformTenant), &v1alpha1.PlatformTenant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PlatformTenant), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePlatformTenants) UpdateStatus(ctx context.Context, platformTenant *v1alpha1.PlatformTenant, opts v1.UpdateOptions) (*v1alpha1.PlatformTenant, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(platformtenantsResource, "status", c.ns, platformTenant), &v1alpha1.PlatformTenant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PlatformTenant), err
}

// Delete takes name of the platformTenant and deletes it. Returns an error if one occurs.
func (c *FakePlatformTenants) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(platformtenantsResource, c.ns, name), &v1alpha1.PlatformTenant{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePlatformTenants) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(platformtenantsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PlatformTenantList{})
	return err
}

// Patch applies the patch and returns the patched platformTenant.
func (c *FakePlatformTenants) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PlatformTenant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(platformtenantsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PlatformTenant{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PlatformTenant), err
}
