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

// FakeUsageExportBuckets implements UsageExportBucketInterface
type FakeUsageExportBuckets struct {
	Fake *FakeProjectV1alpha1
	ns   string
}

var usageexportbucketsResource = schema.GroupVersionResource{Group: "project.google.kubeform.com", Version: "v1alpha1", Resource: "usageexportbuckets"}

var usageexportbucketsKind = schema.GroupVersionKind{Group: "project.google.kubeform.com", Version: "v1alpha1", Kind: "UsageExportBucket"}

// Get takes name of the usageExportBucket, and returns the corresponding usageExportBucket object, and an error if there is any.
func (c *FakeUsageExportBuckets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.UsageExportBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(usageexportbucketsResource, c.ns, name), &v1alpha1.UsageExportBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UsageExportBucket), err
}

// List takes label and field selectors, and returns the list of UsageExportBuckets that match those selectors.
func (c *FakeUsageExportBuckets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UsageExportBucketList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(usageexportbucketsResource, usageexportbucketsKind, c.ns, opts), &v1alpha1.UsageExportBucketList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.UsageExportBucketList{ListMeta: obj.(*v1alpha1.UsageExportBucketList).ListMeta}
	for _, item := range obj.(*v1alpha1.UsageExportBucketList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested usageExportBuckets.
func (c *FakeUsageExportBuckets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(usageexportbucketsResource, c.ns, opts))

}

// Create takes the representation of a usageExportBucket and creates it.  Returns the server's representation of the usageExportBucket, and an error, if there is any.
func (c *FakeUsageExportBuckets) Create(ctx context.Context, usageExportBucket *v1alpha1.UsageExportBucket, opts v1.CreateOptions) (result *v1alpha1.UsageExportBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(usageexportbucketsResource, c.ns, usageExportBucket), &v1alpha1.UsageExportBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UsageExportBucket), err
}

// Update takes the representation of a usageExportBucket and updates it. Returns the server's representation of the usageExportBucket, and an error, if there is any.
func (c *FakeUsageExportBuckets) Update(ctx context.Context, usageExportBucket *v1alpha1.UsageExportBucket, opts v1.UpdateOptions) (result *v1alpha1.UsageExportBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(usageexportbucketsResource, c.ns, usageExportBucket), &v1alpha1.UsageExportBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UsageExportBucket), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUsageExportBuckets) UpdateStatus(ctx context.Context, usageExportBucket *v1alpha1.UsageExportBucket, opts v1.UpdateOptions) (*v1alpha1.UsageExportBucket, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(usageexportbucketsResource, "status", c.ns, usageExportBucket), &v1alpha1.UsageExportBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UsageExportBucket), err
}

// Delete takes name of the usageExportBucket and deletes it. Returns an error if one occurs.
func (c *FakeUsageExportBuckets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(usageexportbucketsResource, c.ns, name), &v1alpha1.UsageExportBucket{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUsageExportBuckets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(usageexportbucketsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.UsageExportBucketList{})
	return err
}

// Patch applies the patch and returns the patched usageExportBucket.
func (c *FakeUsageExportBuckets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UsageExportBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(usageexportbucketsResource, c.ns, name, pt, data, subresources...), &v1alpha1.UsageExportBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UsageExportBucket), err
}
