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

// FakeRouterInterfaces implements RouterInterfaceInterface
type FakeRouterInterfaces struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var routerinterfacesResource = schema.GroupVersionResource{Group: "compute.google.kubeform.com", Version: "v1alpha1", Resource: "routerinterfaces"}

var routerinterfacesKind = schema.GroupVersionKind{Group: "compute.google.kubeform.com", Version: "v1alpha1", Kind: "RouterInterface"}

// Get takes name of the routerInterface, and returns the corresponding routerInterface object, and an error if there is any.
func (c *FakeRouterInterfaces) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RouterInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(routerinterfacesResource, c.ns, name), &v1alpha1.RouterInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RouterInterface), err
}

// List takes label and field selectors, and returns the list of RouterInterfaces that match those selectors.
func (c *FakeRouterInterfaces) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RouterInterfaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(routerinterfacesResource, routerinterfacesKind, c.ns, opts), &v1alpha1.RouterInterfaceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RouterInterfaceList{ListMeta: obj.(*v1alpha1.RouterInterfaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.RouterInterfaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested routerInterfaces.
func (c *FakeRouterInterfaces) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(routerinterfacesResource, c.ns, opts))

}

// Create takes the representation of a routerInterface and creates it.  Returns the server's representation of the routerInterface, and an error, if there is any.
func (c *FakeRouterInterfaces) Create(ctx context.Context, routerInterface *v1alpha1.RouterInterface, opts v1.CreateOptions) (result *v1alpha1.RouterInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(routerinterfacesResource, c.ns, routerInterface), &v1alpha1.RouterInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RouterInterface), err
}

// Update takes the representation of a routerInterface and updates it. Returns the server's representation of the routerInterface, and an error, if there is any.
func (c *FakeRouterInterfaces) Update(ctx context.Context, routerInterface *v1alpha1.RouterInterface, opts v1.UpdateOptions) (result *v1alpha1.RouterInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(routerinterfacesResource, c.ns, routerInterface), &v1alpha1.RouterInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RouterInterface), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRouterInterfaces) UpdateStatus(ctx context.Context, routerInterface *v1alpha1.RouterInterface, opts v1.UpdateOptions) (*v1alpha1.RouterInterface, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(routerinterfacesResource, "status", c.ns, routerInterface), &v1alpha1.RouterInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RouterInterface), err
}

// Delete takes name of the routerInterface and deletes it. Returns an error if one occurs.
func (c *FakeRouterInterfaces) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(routerinterfacesResource, c.ns, name), &v1alpha1.RouterInterface{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRouterInterfaces) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(routerinterfacesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RouterInterfaceList{})
	return err
}

// Patch applies the patch and returns the patched routerInterface.
func (c *FakeRouterInterfaces) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RouterInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(routerinterfacesResource, c.ns, name, pt, data, subresources...), &v1alpha1.RouterInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RouterInterface), err
}
