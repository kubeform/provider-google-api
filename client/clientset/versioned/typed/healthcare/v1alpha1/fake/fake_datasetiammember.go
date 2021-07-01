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

// FakeDatasetIamMembers implements DatasetIamMemberInterface
type FakeDatasetIamMembers struct {
	Fake *FakeHealthcareV1alpha1
	ns   string
}

var datasetiammembersResource = schema.GroupVersionResource{Group: "healthcare.google.kubeform.com", Version: "v1alpha1", Resource: "datasetiammembers"}

var datasetiammembersKind = schema.GroupVersionKind{Group: "healthcare.google.kubeform.com", Version: "v1alpha1", Kind: "DatasetIamMember"}

// Get takes name of the datasetIamMember, and returns the corresponding datasetIamMember object, and an error if there is any.
func (c *FakeDatasetIamMembers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DatasetIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datasetiammembersResource, c.ns, name), &v1alpha1.DatasetIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasetIamMember), err
}

// List takes label and field selectors, and returns the list of DatasetIamMembers that match those selectors.
func (c *FakeDatasetIamMembers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DatasetIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datasetiammembersResource, datasetiammembersKind, c.ns, opts), &v1alpha1.DatasetIamMemberList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DatasetIamMemberList{ListMeta: obj.(*v1alpha1.DatasetIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.DatasetIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested datasetIamMembers.
func (c *FakeDatasetIamMembers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datasetiammembersResource, c.ns, opts))

}

// Create takes the representation of a datasetIamMember and creates it.  Returns the server's representation of the datasetIamMember, and an error, if there is any.
func (c *FakeDatasetIamMembers) Create(ctx context.Context, datasetIamMember *v1alpha1.DatasetIamMember, opts v1.CreateOptions) (result *v1alpha1.DatasetIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datasetiammembersResource, c.ns, datasetIamMember), &v1alpha1.DatasetIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasetIamMember), err
}

// Update takes the representation of a datasetIamMember and updates it. Returns the server's representation of the datasetIamMember, and an error, if there is any.
func (c *FakeDatasetIamMembers) Update(ctx context.Context, datasetIamMember *v1alpha1.DatasetIamMember, opts v1.UpdateOptions) (result *v1alpha1.DatasetIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datasetiammembersResource, c.ns, datasetIamMember), &v1alpha1.DatasetIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasetIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDatasetIamMembers) UpdateStatus(ctx context.Context, datasetIamMember *v1alpha1.DatasetIamMember, opts v1.UpdateOptions) (*v1alpha1.DatasetIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datasetiammembersResource, "status", c.ns, datasetIamMember), &v1alpha1.DatasetIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasetIamMember), err
}

// Delete takes name of the datasetIamMember and deletes it. Returns an error if one occurs.
func (c *FakeDatasetIamMembers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(datasetiammembersResource, c.ns, name), &v1alpha1.DatasetIamMember{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDatasetIamMembers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datasetiammembersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DatasetIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched datasetIamMember.
func (c *FakeDatasetIamMembers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DatasetIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datasetiammembersResource, c.ns, name, pt, data, subresources...), &v1alpha1.DatasetIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DatasetIamMember), err
}
