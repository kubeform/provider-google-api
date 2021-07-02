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

// FakeManagerServicePerimeterBatches implements ManagerServicePerimeterBatchInterface
type FakeManagerServicePerimeterBatches struct {
	Fake *FakeAccesscontextV1alpha1
	ns   string
}

var managerserviceperimeterbatchesResource = schema.GroupVersionResource{Group: "accesscontext.google.kubeform.com", Version: "v1alpha1", Resource: "managerserviceperimeterbatches"}

var managerserviceperimeterbatchesKind = schema.GroupVersionKind{Group: "accesscontext.google.kubeform.com", Version: "v1alpha1", Kind: "ManagerServicePerimeterBatch"}

// Get takes name of the managerServicePerimeterBatch, and returns the corresponding managerServicePerimeterBatch object, and an error if there is any.
func (c *FakeManagerServicePerimeterBatches) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ManagerServicePerimeterBatch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(managerserviceperimeterbatchesResource, c.ns, name), &v1alpha1.ManagerServicePerimeterBatch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerServicePerimeterBatch), err
}

// List takes label and field selectors, and returns the list of ManagerServicePerimeterBatches that match those selectors.
func (c *FakeManagerServicePerimeterBatches) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ManagerServicePerimeterBatchList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(managerserviceperimeterbatchesResource, managerserviceperimeterbatchesKind, c.ns, opts), &v1alpha1.ManagerServicePerimeterBatchList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ManagerServicePerimeterBatchList{ListMeta: obj.(*v1alpha1.ManagerServicePerimeterBatchList).ListMeta}
	for _, item := range obj.(*v1alpha1.ManagerServicePerimeterBatchList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managerServicePerimeterBatches.
func (c *FakeManagerServicePerimeterBatches) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(managerserviceperimeterbatchesResource, c.ns, opts))

}

// Create takes the representation of a managerServicePerimeterBatch and creates it.  Returns the server's representation of the managerServicePerimeterBatch, and an error, if there is any.
func (c *FakeManagerServicePerimeterBatches) Create(ctx context.Context, managerServicePerimeterBatch *v1alpha1.ManagerServicePerimeterBatch, opts v1.CreateOptions) (result *v1alpha1.ManagerServicePerimeterBatch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(managerserviceperimeterbatchesResource, c.ns, managerServicePerimeterBatch), &v1alpha1.ManagerServicePerimeterBatch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerServicePerimeterBatch), err
}

// Update takes the representation of a managerServicePerimeterBatch and updates it. Returns the server's representation of the managerServicePerimeterBatch, and an error, if there is any.
func (c *FakeManagerServicePerimeterBatches) Update(ctx context.Context, managerServicePerimeterBatch *v1alpha1.ManagerServicePerimeterBatch, opts v1.UpdateOptions) (result *v1alpha1.ManagerServicePerimeterBatch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(managerserviceperimeterbatchesResource, c.ns, managerServicePerimeterBatch), &v1alpha1.ManagerServicePerimeterBatch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerServicePerimeterBatch), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagerServicePerimeterBatches) UpdateStatus(ctx context.Context, managerServicePerimeterBatch *v1alpha1.ManagerServicePerimeterBatch, opts v1.UpdateOptions) (*v1alpha1.ManagerServicePerimeterBatch, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(managerserviceperimeterbatchesResource, "status", c.ns, managerServicePerimeterBatch), &v1alpha1.ManagerServicePerimeterBatch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerServicePerimeterBatch), err
}

// Delete takes name of the managerServicePerimeterBatch and deletes it. Returns an error if one occurs.
func (c *FakeManagerServicePerimeterBatches) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(managerserviceperimeterbatchesResource, c.ns, name), &v1alpha1.ManagerServicePerimeterBatch{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagerServicePerimeterBatches) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(managerserviceperimeterbatchesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ManagerServicePerimeterBatchList{})
	return err
}

// Patch applies the patch and returns the patched managerServicePerimeterBatch.
func (c *FakeManagerServicePerimeterBatches) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagerServicePerimeterBatch, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(managerserviceperimeterbatchesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ManagerServicePerimeterBatch{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerServicePerimeterBatch), err
}
