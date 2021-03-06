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

// FakeRouterNATs implements RouterNATInterface
type FakeRouterNATs struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var routernatsResource = schema.GroupVersionResource{Group: "compute.google.kubeform.com", Version: "v1alpha1", Resource: "routernats"}

var routernatsKind = schema.GroupVersionKind{Group: "compute.google.kubeform.com", Version: "v1alpha1", Kind: "RouterNAT"}

// Get takes name of the routerNAT, and returns the corresponding routerNAT object, and an error if there is any.
func (c *FakeRouterNATs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RouterNAT, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(routernatsResource, c.ns, name), &v1alpha1.RouterNAT{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RouterNAT), err
}

// List takes label and field selectors, and returns the list of RouterNATs that match those selectors.
func (c *FakeRouterNATs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RouterNATList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(routernatsResource, routernatsKind, c.ns, opts), &v1alpha1.RouterNATList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RouterNATList{ListMeta: obj.(*v1alpha1.RouterNATList).ListMeta}
	for _, item := range obj.(*v1alpha1.RouterNATList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested routerNATs.
func (c *FakeRouterNATs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(routernatsResource, c.ns, opts))

}

// Create takes the representation of a routerNAT and creates it.  Returns the server's representation of the routerNAT, and an error, if there is any.
func (c *FakeRouterNATs) Create(ctx context.Context, routerNAT *v1alpha1.RouterNAT, opts v1.CreateOptions) (result *v1alpha1.RouterNAT, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(routernatsResource, c.ns, routerNAT), &v1alpha1.RouterNAT{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RouterNAT), err
}

// Update takes the representation of a routerNAT and updates it. Returns the server's representation of the routerNAT, and an error, if there is any.
func (c *FakeRouterNATs) Update(ctx context.Context, routerNAT *v1alpha1.RouterNAT, opts v1.UpdateOptions) (result *v1alpha1.RouterNAT, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(routernatsResource, c.ns, routerNAT), &v1alpha1.RouterNAT{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RouterNAT), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRouterNATs) UpdateStatus(ctx context.Context, routerNAT *v1alpha1.RouterNAT, opts v1.UpdateOptions) (*v1alpha1.RouterNAT, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(routernatsResource, "status", c.ns, routerNAT), &v1alpha1.RouterNAT{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RouterNAT), err
}

// Delete takes name of the routerNAT and deletes it. Returns an error if one occurs.
func (c *FakeRouterNATs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(routernatsResource, c.ns, name), &v1alpha1.RouterNAT{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRouterNATs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(routernatsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RouterNATList{})
	return err
}

// Patch applies the patch and returns the patched routerNAT.
func (c *FakeRouterNATs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RouterNAT, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(routernatsResource, c.ns, name, pt, data, subresources...), &v1alpha1.RouterNAT{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RouterNAT), err
}
