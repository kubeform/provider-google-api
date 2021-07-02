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

	v1alpha1 "kubeform.dev/provider-google-api/apis/notebooks/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLocations implements LocationInterface
type FakeLocations struct {
	Fake *FakeNotebooksV1alpha1
	ns   string
}

var locationsResource = schema.GroupVersionResource{Group: "notebooks.google.kubeform.com", Version: "v1alpha1", Resource: "locations"}

var locationsKind = schema.GroupVersionKind{Group: "notebooks.google.kubeform.com", Version: "v1alpha1", Kind: "Location"}

// Get takes name of the location, and returns the corresponding location object, and an error if there is any.
func (c *FakeLocations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Location, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(locationsResource, c.ns, name), &v1alpha1.Location{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Location), err
}

// List takes label and field selectors, and returns the list of Locations that match those selectors.
func (c *FakeLocations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LocationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(locationsResource, locationsKind, c.ns, opts), &v1alpha1.LocationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LocationList{ListMeta: obj.(*v1alpha1.LocationList).ListMeta}
	for _, item := range obj.(*v1alpha1.LocationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested locations.
func (c *FakeLocations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(locationsResource, c.ns, opts))

}

// Create takes the representation of a location and creates it.  Returns the server's representation of the location, and an error, if there is any.
func (c *FakeLocations) Create(ctx context.Context, location *v1alpha1.Location, opts v1.CreateOptions) (result *v1alpha1.Location, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(locationsResource, c.ns, location), &v1alpha1.Location{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Location), err
}

// Update takes the representation of a location and updates it. Returns the server's representation of the location, and an error, if there is any.
func (c *FakeLocations) Update(ctx context.Context, location *v1alpha1.Location, opts v1.UpdateOptions) (result *v1alpha1.Location, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(locationsResource, c.ns, location), &v1alpha1.Location{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Location), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLocations) UpdateStatus(ctx context.Context, location *v1alpha1.Location, opts v1.UpdateOptions) (*v1alpha1.Location, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(locationsResource, "status", c.ns, location), &v1alpha1.Location{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Location), err
}

// Delete takes name of the location and deletes it. Returns an error if one occurs.
func (c *FakeLocations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(locationsResource, c.ns, name), &v1alpha1.Location{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLocations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(locationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LocationList{})
	return err
}

// Patch applies the patch and returns the patched location.
func (c *FakeLocations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Location, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(locationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Location{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Location), err
}
