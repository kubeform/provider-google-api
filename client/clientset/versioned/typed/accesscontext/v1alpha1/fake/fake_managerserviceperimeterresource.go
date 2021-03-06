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

	v1alpha1 "kubeform.dev/provider-google-api/apis/accesscontext/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeManagerServicePerimeterResources implements ManagerServicePerimeterResourceInterface
type FakeManagerServicePerimeterResources struct {
	Fake *FakeAccesscontextV1alpha1
	ns   string
}

var managerserviceperimeterresourcesResource = schema.GroupVersionResource{Group: "accesscontext.google.kubeform.com", Version: "v1alpha1", Resource: "managerserviceperimeterresources"}

var managerserviceperimeterresourcesKind = schema.GroupVersionKind{Group: "accesscontext.google.kubeform.com", Version: "v1alpha1", Kind: "ManagerServicePerimeterResource"}

// Get takes name of the managerServicePerimeterResource, and returns the corresponding managerServicePerimeterResource object, and an error if there is any.
func (c *FakeManagerServicePerimeterResources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ManagerServicePerimeterResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(managerserviceperimeterresourcesResource, c.ns, name), &v1alpha1.ManagerServicePerimeterResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerServicePerimeterResource), err
}

// List takes label and field selectors, and returns the list of ManagerServicePerimeterResources that match those selectors.
func (c *FakeManagerServicePerimeterResources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ManagerServicePerimeterResourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(managerserviceperimeterresourcesResource, managerserviceperimeterresourcesKind, c.ns, opts), &v1alpha1.ManagerServicePerimeterResourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ManagerServicePerimeterResourceList{ListMeta: obj.(*v1alpha1.ManagerServicePerimeterResourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.ManagerServicePerimeterResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managerServicePerimeterResources.
func (c *FakeManagerServicePerimeterResources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(managerserviceperimeterresourcesResource, c.ns, opts))

}

// Create takes the representation of a managerServicePerimeterResource and creates it.  Returns the server's representation of the managerServicePerimeterResource, and an error, if there is any.
func (c *FakeManagerServicePerimeterResources) Create(ctx context.Context, managerServicePerimeterResource *v1alpha1.ManagerServicePerimeterResource, opts v1.CreateOptions) (result *v1alpha1.ManagerServicePerimeterResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(managerserviceperimeterresourcesResource, c.ns, managerServicePerimeterResource), &v1alpha1.ManagerServicePerimeterResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerServicePerimeterResource), err
}

// Update takes the representation of a managerServicePerimeterResource and updates it. Returns the server's representation of the managerServicePerimeterResource, and an error, if there is any.
func (c *FakeManagerServicePerimeterResources) Update(ctx context.Context, managerServicePerimeterResource *v1alpha1.ManagerServicePerimeterResource, opts v1.UpdateOptions) (result *v1alpha1.ManagerServicePerimeterResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(managerserviceperimeterresourcesResource, c.ns, managerServicePerimeterResource), &v1alpha1.ManagerServicePerimeterResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerServicePerimeterResource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagerServicePerimeterResources) UpdateStatus(ctx context.Context, managerServicePerimeterResource *v1alpha1.ManagerServicePerimeterResource, opts v1.UpdateOptions) (*v1alpha1.ManagerServicePerimeterResource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(managerserviceperimeterresourcesResource, "status", c.ns, managerServicePerimeterResource), &v1alpha1.ManagerServicePerimeterResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerServicePerimeterResource), err
}

// Delete takes name of the managerServicePerimeterResource and deletes it. Returns an error if one occurs.
func (c *FakeManagerServicePerimeterResources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(managerserviceperimeterresourcesResource, c.ns, name), &v1alpha1.ManagerServicePerimeterResource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagerServicePerimeterResources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(managerserviceperimeterresourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ManagerServicePerimeterResourceList{})
	return err
}

// Patch applies the patch and returns the patched managerServicePerimeterResource.
func (c *FakeManagerServicePerimeterResources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagerServicePerimeterResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(managerserviceperimeterresourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ManagerServicePerimeterResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerServicePerimeterResource), err
}
