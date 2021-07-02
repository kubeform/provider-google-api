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

	v1alpha1 "kubeform.dev/provider-google-api/apis/organization/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAccessApprovalSettingses implements AccessApprovalSettingsInterface
type FakeAccessApprovalSettingses struct {
	Fake *FakeOrganizationV1alpha1
	ns   string
}

var accessapprovalsettingsesResource = schema.GroupVersionResource{Group: "organization.google.kubeform.com", Version: "v1alpha1", Resource: "accessapprovalsettingses"}

var accessapprovalsettingsesKind = schema.GroupVersionKind{Group: "organization.google.kubeform.com", Version: "v1alpha1", Kind: "AccessApprovalSettings"}

// Get takes name of the accessApprovalSettings, and returns the corresponding accessApprovalSettings object, and an error if there is any.
func (c *FakeAccessApprovalSettingses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AccessApprovalSettings, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(accessapprovalsettingsesResource, c.ns, name), &v1alpha1.AccessApprovalSettings{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessApprovalSettings), err
}

// List takes label and field selectors, and returns the list of AccessApprovalSettingses that match those selectors.
func (c *FakeAccessApprovalSettingses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AccessApprovalSettingsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(accessapprovalsettingsesResource, accessapprovalsettingsesKind, c.ns, opts), &v1alpha1.AccessApprovalSettingsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AccessApprovalSettingsList{ListMeta: obj.(*v1alpha1.AccessApprovalSettingsList).ListMeta}
	for _, item := range obj.(*v1alpha1.AccessApprovalSettingsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested accessApprovalSettingses.
func (c *FakeAccessApprovalSettingses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(accessapprovalsettingsesResource, c.ns, opts))

}

// Create takes the representation of a accessApprovalSettings and creates it.  Returns the server's representation of the accessApprovalSettings, and an error, if there is any.
func (c *FakeAccessApprovalSettingses) Create(ctx context.Context, accessApprovalSettings *v1alpha1.AccessApprovalSettings, opts v1.CreateOptions) (result *v1alpha1.AccessApprovalSettings, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(accessapprovalsettingsesResource, c.ns, accessApprovalSettings), &v1alpha1.AccessApprovalSettings{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessApprovalSettings), err
}

// Update takes the representation of a accessApprovalSettings and updates it. Returns the server's representation of the accessApprovalSettings, and an error, if there is any.
func (c *FakeAccessApprovalSettingses) Update(ctx context.Context, accessApprovalSettings *v1alpha1.AccessApprovalSettings, opts v1.UpdateOptions) (result *v1alpha1.AccessApprovalSettings, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(accessapprovalsettingsesResource, c.ns, accessApprovalSettings), &v1alpha1.AccessApprovalSettings{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessApprovalSettings), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAccessApprovalSettingses) UpdateStatus(ctx context.Context, accessApprovalSettings *v1alpha1.AccessApprovalSettings, opts v1.UpdateOptions) (*v1alpha1.AccessApprovalSettings, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(accessapprovalsettingsesResource, "status", c.ns, accessApprovalSettings), &v1alpha1.AccessApprovalSettings{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessApprovalSettings), err
}

// Delete takes name of the accessApprovalSettings and deletes it. Returns an error if one occurs.
func (c *FakeAccessApprovalSettingses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(accessapprovalsettingsesResource, c.ns, name), &v1alpha1.AccessApprovalSettings{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAccessApprovalSettingses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(accessapprovalsettingsesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AccessApprovalSettingsList{})
	return err
}

// Patch applies the patch and returns the patched accessApprovalSettings.
func (c *FakeAccessApprovalSettingses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AccessApprovalSettings, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(accessapprovalsettingsesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AccessApprovalSettings{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessApprovalSettings), err
}
