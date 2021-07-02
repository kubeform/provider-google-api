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

	v1alpha1 "kubeform.dev/provider-google-api/apis/bigquery/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTables implements TableInterface
type FakeTables struct {
	Fake *FakeBigqueryV1alpha1
	ns   string
}

var tablesResource = schema.GroupVersionResource{Group: "bigquery.google.kubeform.com", Version: "v1alpha1", Resource: "tables"}

var tablesKind = schema.GroupVersionKind{Group: "bigquery.google.kubeform.com", Version: "v1alpha1", Kind: "Table"}

// Get takes name of the table, and returns the corresponding table object, and an error if there is any.
func (c *FakeTables) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Table, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tablesResource, c.ns, name), &v1alpha1.Table{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Table), err
}

// List takes label and field selectors, and returns the list of Tables that match those selectors.
func (c *FakeTables) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TableList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tablesResource, tablesKind, c.ns, opts), &v1alpha1.TableList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TableList{ListMeta: obj.(*v1alpha1.TableList).ListMeta}
	for _, item := range obj.(*v1alpha1.TableList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tables.
func (c *FakeTables) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tablesResource, c.ns, opts))

}

// Create takes the representation of a table and creates it.  Returns the server's representation of the table, and an error, if there is any.
func (c *FakeTables) Create(ctx context.Context, table *v1alpha1.Table, opts v1.CreateOptions) (result *v1alpha1.Table, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tablesResource, c.ns, table), &v1alpha1.Table{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Table), err
}

// Update takes the representation of a table and updates it. Returns the server's representation of the table, and an error, if there is any.
func (c *FakeTables) Update(ctx context.Context, table *v1alpha1.Table, opts v1.UpdateOptions) (result *v1alpha1.Table, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tablesResource, c.ns, table), &v1alpha1.Table{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Table), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTables) UpdateStatus(ctx context.Context, table *v1alpha1.Table, opts v1.UpdateOptions) (*v1alpha1.Table, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(tablesResource, "status", c.ns, table), &v1alpha1.Table{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Table), err
}

// Delete takes name of the table and deletes it. Returns an error if one occurs.
func (c *FakeTables) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tablesResource, c.ns, name), &v1alpha1.Table{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTables) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tablesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TableList{})
	return err
}

// Patch applies the patch and returns the patched table.
func (c *FakeTables) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Table, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tablesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Table{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Table), err
}
