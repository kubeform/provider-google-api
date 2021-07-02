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

// FakeRouters implements RouterInterface
type FakeRouters struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var routersResource = schema.GroupVersionResource{Group: "compute.google.kubeform.com", Version: "v1alpha1", Resource: "routers"}

var routersKind = schema.GroupVersionKind{Group: "compute.google.kubeform.com", Version: "v1alpha1", Kind: "Router"}

// Get takes name of the router, and returns the corresponding router object, and an error if there is any.
func (c *FakeRouters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Router, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(routersResource, c.ns, name), &v1alpha1.Router{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Router), err
}

// List takes label and field selectors, and returns the list of Routers that match those selectors.
func (c *FakeRouters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RouterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(routersResource, routersKind, c.ns, opts), &v1alpha1.RouterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RouterList{ListMeta: obj.(*v1alpha1.RouterList).ListMeta}
	for _, item := range obj.(*v1alpha1.RouterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested routers.
func (c *FakeRouters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(routersResource, c.ns, opts))

}

// Create takes the representation of a router and creates it.  Returns the server's representation of the router, and an error, if there is any.
func (c *FakeRouters) Create(ctx context.Context, router *v1alpha1.Router, opts v1.CreateOptions) (result *v1alpha1.Router, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(routersResource, c.ns, router), &v1alpha1.Router{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Router), err
}

// Update takes the representation of a router and updates it. Returns the server's representation of the router, and an error, if there is any.
func (c *FakeRouters) Update(ctx context.Context, router *v1alpha1.Router, opts v1.UpdateOptions) (result *v1alpha1.Router, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(routersResource, c.ns, router), &v1alpha1.Router{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Router), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRouters) UpdateStatus(ctx context.Context, router *v1alpha1.Router, opts v1.UpdateOptions) (*v1alpha1.Router, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(routersResource, "status", c.ns, router), &v1alpha1.Router{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Router), err
}

// Delete takes name of the router and deletes it. Returns an error if one occurs.
func (c *FakeRouters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(routersResource, c.ns, name), &v1alpha1.Router{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRouters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(routersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RouterList{})
	return err
}

// Patch applies the patch and returns the patched router.
func (c *FakeRouters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Router, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(routersResource, c.ns, name, pt, data, subresources...), &v1alpha1.Router{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Router), err
}
