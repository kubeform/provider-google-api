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

	v1alpha1 "kubeform.dev/provider-google-api/apis/monitoring/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSlos implements SloInterface
type FakeSlos struct {
	Fake *FakeMonitoringV1alpha1
	ns   string
}

var slosResource = schema.GroupVersionResource{Group: "monitoring.google.kubeform.com", Version: "v1alpha1", Resource: "slos"}

var slosKind = schema.GroupVersionKind{Group: "monitoring.google.kubeform.com", Version: "v1alpha1", Kind: "Slo"}

// Get takes name of the slo, and returns the corresponding slo object, and an error if there is any.
func (c *FakeSlos) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Slo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(slosResource, c.ns, name), &v1alpha1.Slo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Slo), err
}

// List takes label and field selectors, and returns the list of Slos that match those selectors.
func (c *FakeSlos) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SloList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(slosResource, slosKind, c.ns, opts), &v1alpha1.SloList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SloList{ListMeta: obj.(*v1alpha1.SloList).ListMeta}
	for _, item := range obj.(*v1alpha1.SloList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested slos.
func (c *FakeSlos) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(slosResource, c.ns, opts))

}

// Create takes the representation of a slo and creates it.  Returns the server's representation of the slo, and an error, if there is any.
func (c *FakeSlos) Create(ctx context.Context, slo *v1alpha1.Slo, opts v1.CreateOptions) (result *v1alpha1.Slo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(slosResource, c.ns, slo), &v1alpha1.Slo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Slo), err
}

// Update takes the representation of a slo and updates it. Returns the server's representation of the slo, and an error, if there is any.
func (c *FakeSlos) Update(ctx context.Context, slo *v1alpha1.Slo, opts v1.UpdateOptions) (result *v1alpha1.Slo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(slosResource, c.ns, slo), &v1alpha1.Slo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Slo), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSlos) UpdateStatus(ctx context.Context, slo *v1alpha1.Slo, opts v1.UpdateOptions) (*v1alpha1.Slo, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(slosResource, "status", c.ns, slo), &v1alpha1.Slo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Slo), err
}

// Delete takes name of the slo and deletes it. Returns an error if one occurs.
func (c *FakeSlos) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(slosResource, c.ns, name), &v1alpha1.Slo{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSlos) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(slosResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SloList{})
	return err
}

// Patch applies the patch and returns the patched slo.
func (c *FakeSlos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Slo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(slosResource, c.ns, name, pt, data, subresources...), &v1alpha1.Slo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Slo), err
}
