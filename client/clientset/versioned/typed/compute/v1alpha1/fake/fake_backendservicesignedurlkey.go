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
	v1alpha1 "kubeform.dev/provider-google-api/apis/compute/v1alpha1"
)

// FakeBackendServiceSignedURLKeys implements BackendServiceSignedURLKeyInterface
type FakeBackendServiceSignedURLKeys struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var backendservicesignedurlkeysResource = schema.GroupVersionResource{Group: "compute.google.kubeform.com", Version: "v1alpha1", Resource: "backendservicesignedurlkeys"}

var backendservicesignedurlkeysKind = schema.GroupVersionKind{Group: "compute.google.kubeform.com", Version: "v1alpha1", Kind: "BackendServiceSignedURLKey"}

// Get takes name of the backendServiceSignedURLKey, and returns the corresponding backendServiceSignedURLKey object, and an error if there is any.
func (c *FakeBackendServiceSignedURLKeys) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BackendServiceSignedURLKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(backendservicesignedurlkeysResource, c.ns, name), &v1alpha1.BackendServiceSignedURLKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackendServiceSignedURLKey), err
}

// List takes label and field selectors, and returns the list of BackendServiceSignedURLKeys that match those selectors.
func (c *FakeBackendServiceSignedURLKeys) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BackendServiceSignedURLKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(backendservicesignedurlkeysResource, backendservicesignedurlkeysKind, c.ns, opts), &v1alpha1.BackendServiceSignedURLKeyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BackendServiceSignedURLKeyList{ListMeta: obj.(*v1alpha1.BackendServiceSignedURLKeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.BackendServiceSignedURLKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested backendServiceSignedURLKeys.
func (c *FakeBackendServiceSignedURLKeys) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(backendservicesignedurlkeysResource, c.ns, opts))

}

// Create takes the representation of a backendServiceSignedURLKey and creates it.  Returns the server's representation of the backendServiceSignedURLKey, and an error, if there is any.
func (c *FakeBackendServiceSignedURLKeys) Create(ctx context.Context, backendServiceSignedURLKey *v1alpha1.BackendServiceSignedURLKey, opts v1.CreateOptions) (result *v1alpha1.BackendServiceSignedURLKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(backendservicesignedurlkeysResource, c.ns, backendServiceSignedURLKey), &v1alpha1.BackendServiceSignedURLKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackendServiceSignedURLKey), err
}

// Update takes the representation of a backendServiceSignedURLKey and updates it. Returns the server's representation of the backendServiceSignedURLKey, and an error, if there is any.
func (c *FakeBackendServiceSignedURLKeys) Update(ctx context.Context, backendServiceSignedURLKey *v1alpha1.BackendServiceSignedURLKey, opts v1.UpdateOptions) (result *v1alpha1.BackendServiceSignedURLKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(backendservicesignedurlkeysResource, c.ns, backendServiceSignedURLKey), &v1alpha1.BackendServiceSignedURLKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackendServiceSignedURLKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBackendServiceSignedURLKeys) UpdateStatus(ctx context.Context, backendServiceSignedURLKey *v1alpha1.BackendServiceSignedURLKey, opts v1.UpdateOptions) (*v1alpha1.BackendServiceSignedURLKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(backendservicesignedurlkeysResource, "status", c.ns, backendServiceSignedURLKey), &v1alpha1.BackendServiceSignedURLKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackendServiceSignedURLKey), err
}

// Delete takes name of the backendServiceSignedURLKey and deletes it. Returns an error if one occurs.
func (c *FakeBackendServiceSignedURLKeys) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(backendservicesignedurlkeysResource, c.ns, name), &v1alpha1.BackendServiceSignedURLKey{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBackendServiceSignedURLKeys) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(backendservicesignedurlkeysResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BackendServiceSignedURLKeyList{})
	return err
}

// Patch applies the patch and returns the patched backendServiceSignedURLKey.
func (c *FakeBackendServiceSignedURLKeys) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BackendServiceSignedURLKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(backendservicesignedurlkeysResource, c.ns, name, pt, data, subresources...), &v1alpha1.BackendServiceSignedURLKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackendServiceSignedURLKey), err
}
