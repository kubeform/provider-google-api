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

	v1alpha1 "kubeform.dev/provider-google-api/apis/compute/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBackendServices implements BackendServiceInterface
type FakeBackendServices struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var backendservicesResource = schema.GroupVersionResource{Group: "compute.google.kubeform.com", Version: "v1alpha1", Resource: "backendservices"}

var backendservicesKind = schema.GroupVersionKind{Group: "compute.google.kubeform.com", Version: "v1alpha1", Kind: "BackendService"}

// Get takes name of the backendService, and returns the corresponding backendService object, and an error if there is any.
func (c *FakeBackendServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BackendService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(backendservicesResource, c.ns, name), &v1alpha1.BackendService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackendService), err
}

// List takes label and field selectors, and returns the list of BackendServices that match those selectors.
func (c *FakeBackendServices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BackendServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(backendservicesResource, backendservicesKind, c.ns, opts), &v1alpha1.BackendServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BackendServiceList{ListMeta: obj.(*v1alpha1.BackendServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.BackendServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested backendServices.
func (c *FakeBackendServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(backendservicesResource, c.ns, opts))

}

// Create takes the representation of a backendService and creates it.  Returns the server's representation of the backendService, and an error, if there is any.
func (c *FakeBackendServices) Create(ctx context.Context, backendService *v1alpha1.BackendService, opts v1.CreateOptions) (result *v1alpha1.BackendService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(backendservicesResource, c.ns, backendService), &v1alpha1.BackendService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackendService), err
}

// Update takes the representation of a backendService and updates it. Returns the server's representation of the backendService, and an error, if there is any.
func (c *FakeBackendServices) Update(ctx context.Context, backendService *v1alpha1.BackendService, opts v1.UpdateOptions) (result *v1alpha1.BackendService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(backendservicesResource, c.ns, backendService), &v1alpha1.BackendService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackendService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBackendServices) UpdateStatus(ctx context.Context, backendService *v1alpha1.BackendService, opts v1.UpdateOptions) (*v1alpha1.BackendService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(backendservicesResource, "status", c.ns, backendService), &v1alpha1.BackendService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackendService), err
}

// Delete takes name of the backendService and deletes it. Returns an error if one occurs.
func (c *FakeBackendServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(backendservicesResource, c.ns, name), &v1alpha1.BackendService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBackendServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(backendservicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BackendServiceList{})
	return err
}

// Patch applies the patch and returns the patched backendService.
func (c *FakeBackendServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BackendService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(backendservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.BackendService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackendService), err
}
