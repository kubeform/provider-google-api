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

	v1alpha1 "kubeform.dev/provider-google-api/apis/storage/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBucketAccessControls implements BucketAccessControlInterface
type FakeBucketAccessControls struct {
	Fake *FakeStorageV1alpha1
	ns   string
}

var bucketaccesscontrolsResource = schema.GroupVersionResource{Group: "storage.google.kubeform.com", Version: "v1alpha1", Resource: "bucketaccesscontrols"}

var bucketaccesscontrolsKind = schema.GroupVersionKind{Group: "storage.google.kubeform.com", Version: "v1alpha1", Kind: "BucketAccessControl"}

// Get takes name of the bucketAccessControl, and returns the corresponding bucketAccessControl object, and an error if there is any.
func (c *FakeBucketAccessControls) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BucketAccessControl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(bucketaccesscontrolsResource, c.ns, name), &v1alpha1.BucketAccessControl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BucketAccessControl), err
}

// List takes label and field selectors, and returns the list of BucketAccessControls that match those selectors.
func (c *FakeBucketAccessControls) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BucketAccessControlList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(bucketaccesscontrolsResource, bucketaccesscontrolsKind, c.ns, opts), &v1alpha1.BucketAccessControlList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BucketAccessControlList{ListMeta: obj.(*v1alpha1.BucketAccessControlList).ListMeta}
	for _, item := range obj.(*v1alpha1.BucketAccessControlList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bucketAccessControls.
func (c *FakeBucketAccessControls) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(bucketaccesscontrolsResource, c.ns, opts))

}

// Create takes the representation of a bucketAccessControl and creates it.  Returns the server's representation of the bucketAccessControl, and an error, if there is any.
func (c *FakeBucketAccessControls) Create(ctx context.Context, bucketAccessControl *v1alpha1.BucketAccessControl, opts v1.CreateOptions) (result *v1alpha1.BucketAccessControl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(bucketaccesscontrolsResource, c.ns, bucketAccessControl), &v1alpha1.BucketAccessControl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BucketAccessControl), err
}

// Update takes the representation of a bucketAccessControl and updates it. Returns the server's representation of the bucketAccessControl, and an error, if there is any.
func (c *FakeBucketAccessControls) Update(ctx context.Context, bucketAccessControl *v1alpha1.BucketAccessControl, opts v1.UpdateOptions) (result *v1alpha1.BucketAccessControl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(bucketaccesscontrolsResource, c.ns, bucketAccessControl), &v1alpha1.BucketAccessControl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BucketAccessControl), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBucketAccessControls) UpdateStatus(ctx context.Context, bucketAccessControl *v1alpha1.BucketAccessControl, opts v1.UpdateOptions) (*v1alpha1.BucketAccessControl, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(bucketaccesscontrolsResource, "status", c.ns, bucketAccessControl), &v1alpha1.BucketAccessControl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BucketAccessControl), err
}

// Delete takes name of the bucketAccessControl and deletes it. Returns an error if one occurs.
func (c *FakeBucketAccessControls) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(bucketaccesscontrolsResource, c.ns, name), &v1alpha1.BucketAccessControl{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBucketAccessControls) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(bucketaccesscontrolsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BucketAccessControlList{})
	return err
}

// Patch applies the patch and returns the patched bucketAccessControl.
func (c *FakeBucketAccessControls) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BucketAccessControl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bucketaccesscontrolsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BucketAccessControl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BucketAccessControl), err
}
