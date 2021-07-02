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

	v1alpha1 "kubeform.dev/provider-google-api/apis/resource/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeManagerLiens implements ManagerLienInterface
type FakeManagerLiens struct {
	Fake *FakeResourceV1alpha1
	ns   string
}

var managerliensResource = schema.GroupVersionResource{Group: "resource.google.kubeform.com", Version: "v1alpha1", Resource: "managerliens"}

var managerliensKind = schema.GroupVersionKind{Group: "resource.google.kubeform.com", Version: "v1alpha1", Kind: "ManagerLien"}

// Get takes name of the managerLien, and returns the corresponding managerLien object, and an error if there is any.
func (c *FakeManagerLiens) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ManagerLien, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(managerliensResource, c.ns, name), &v1alpha1.ManagerLien{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerLien), err
}

// List takes label and field selectors, and returns the list of ManagerLiens that match those selectors.
func (c *FakeManagerLiens) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ManagerLienList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(managerliensResource, managerliensKind, c.ns, opts), &v1alpha1.ManagerLienList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ManagerLienList{ListMeta: obj.(*v1alpha1.ManagerLienList).ListMeta}
	for _, item := range obj.(*v1alpha1.ManagerLienList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managerLiens.
func (c *FakeManagerLiens) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(managerliensResource, c.ns, opts))

}

// Create takes the representation of a managerLien and creates it.  Returns the server's representation of the managerLien, and an error, if there is any.
func (c *FakeManagerLiens) Create(ctx context.Context, managerLien *v1alpha1.ManagerLien, opts v1.CreateOptions) (result *v1alpha1.ManagerLien, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(managerliensResource, c.ns, managerLien), &v1alpha1.ManagerLien{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerLien), err
}

// Update takes the representation of a managerLien and updates it. Returns the server's representation of the managerLien, and an error, if there is any.
func (c *FakeManagerLiens) Update(ctx context.Context, managerLien *v1alpha1.ManagerLien, opts v1.UpdateOptions) (result *v1alpha1.ManagerLien, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(managerliensResource, c.ns, managerLien), &v1alpha1.ManagerLien{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerLien), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagerLiens) UpdateStatus(ctx context.Context, managerLien *v1alpha1.ManagerLien, opts v1.UpdateOptions) (*v1alpha1.ManagerLien, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(managerliensResource, "status", c.ns, managerLien), &v1alpha1.ManagerLien{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerLien), err
}

// Delete takes name of the managerLien and deletes it. Returns an error if one occurs.
func (c *FakeManagerLiens) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(managerliensResource, c.ns, name), &v1alpha1.ManagerLien{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagerLiens) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(managerliensResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ManagerLienList{})
	return err
}

// Patch applies the patch and returns the patched managerLien.
func (c *FakeManagerLiens) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagerLien, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(managerliensResource, c.ns, name, pt, data, subresources...), &v1alpha1.ManagerLien{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerLien), err
}
