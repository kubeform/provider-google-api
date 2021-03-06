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

	v1alpha1 "kubeform.dev/provider-google-api/apis/container/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAzureNodePools implements AzureNodePoolInterface
type FakeAzureNodePools struct {
	Fake *FakeContainerV1alpha1
	ns   string
}

var azurenodepoolsResource = schema.GroupVersionResource{Group: "container.google.kubeform.com", Version: "v1alpha1", Resource: "azurenodepools"}

var azurenodepoolsKind = schema.GroupVersionKind{Group: "container.google.kubeform.com", Version: "v1alpha1", Kind: "AzureNodePool"}

// Get takes name of the azureNodePool, and returns the corresponding azureNodePool object, and an error if there is any.
func (c *FakeAzureNodePools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AzureNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(azurenodepoolsResource, c.ns, name), &v1alpha1.AzureNodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureNodePool), err
}

// List takes label and field selectors, and returns the list of AzureNodePools that match those selectors.
func (c *FakeAzureNodePools) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AzureNodePoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(azurenodepoolsResource, azurenodepoolsKind, c.ns, opts), &v1alpha1.AzureNodePoolList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzureNodePoolList{ListMeta: obj.(*v1alpha1.AzureNodePoolList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzureNodePoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azureNodePools.
func (c *FakeAzureNodePools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(azurenodepoolsResource, c.ns, opts))

}

// Create takes the representation of a azureNodePool and creates it.  Returns the server's representation of the azureNodePool, and an error, if there is any.
func (c *FakeAzureNodePools) Create(ctx context.Context, azureNodePool *v1alpha1.AzureNodePool, opts v1.CreateOptions) (result *v1alpha1.AzureNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(azurenodepoolsResource, c.ns, azureNodePool), &v1alpha1.AzureNodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureNodePool), err
}

// Update takes the representation of a azureNodePool and updates it. Returns the server's representation of the azureNodePool, and an error, if there is any.
func (c *FakeAzureNodePools) Update(ctx context.Context, azureNodePool *v1alpha1.AzureNodePool, opts v1.UpdateOptions) (result *v1alpha1.AzureNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(azurenodepoolsResource, c.ns, azureNodePool), &v1alpha1.AzureNodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureNodePool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzureNodePools) UpdateStatus(ctx context.Context, azureNodePool *v1alpha1.AzureNodePool, opts v1.UpdateOptions) (*v1alpha1.AzureNodePool, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(azurenodepoolsResource, "status", c.ns, azureNodePool), &v1alpha1.AzureNodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureNodePool), err
}

// Delete takes name of the azureNodePool and deletes it. Returns an error if one occurs.
func (c *FakeAzureNodePools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(azurenodepoolsResource, c.ns, name), &v1alpha1.AzureNodePool{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzureNodePools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(azurenodepoolsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzureNodePoolList{})
	return err
}

// Patch applies the patch and returns the patched azureNodePool.
func (c *FakeAzureNodePools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AzureNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(azurenodepoolsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AzureNodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureNodePool), err
}
