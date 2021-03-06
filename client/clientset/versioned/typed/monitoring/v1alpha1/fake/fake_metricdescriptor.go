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

// FakeMetricDescriptors implements MetricDescriptorInterface
type FakeMetricDescriptors struct {
	Fake *FakeMonitoringV1alpha1
	ns   string
}

var metricdescriptorsResource = schema.GroupVersionResource{Group: "monitoring.google.kubeform.com", Version: "v1alpha1", Resource: "metricdescriptors"}

var metricdescriptorsKind = schema.GroupVersionKind{Group: "monitoring.google.kubeform.com", Version: "v1alpha1", Kind: "MetricDescriptor"}

// Get takes name of the metricDescriptor, and returns the corresponding metricDescriptor object, and an error if there is any.
func (c *FakeMetricDescriptors) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MetricDescriptor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(metricdescriptorsResource, c.ns, name), &v1alpha1.MetricDescriptor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricDescriptor), err
}

// List takes label and field selectors, and returns the list of MetricDescriptors that match those selectors.
func (c *FakeMetricDescriptors) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MetricDescriptorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(metricdescriptorsResource, metricdescriptorsKind, c.ns, opts), &v1alpha1.MetricDescriptorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MetricDescriptorList{ListMeta: obj.(*v1alpha1.MetricDescriptorList).ListMeta}
	for _, item := range obj.(*v1alpha1.MetricDescriptorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested metricDescriptors.
func (c *FakeMetricDescriptors) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(metricdescriptorsResource, c.ns, opts))

}

// Create takes the representation of a metricDescriptor and creates it.  Returns the server's representation of the metricDescriptor, and an error, if there is any.
func (c *FakeMetricDescriptors) Create(ctx context.Context, metricDescriptor *v1alpha1.MetricDescriptor, opts v1.CreateOptions) (result *v1alpha1.MetricDescriptor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(metricdescriptorsResource, c.ns, metricDescriptor), &v1alpha1.MetricDescriptor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricDescriptor), err
}

// Update takes the representation of a metricDescriptor and updates it. Returns the server's representation of the metricDescriptor, and an error, if there is any.
func (c *FakeMetricDescriptors) Update(ctx context.Context, metricDescriptor *v1alpha1.MetricDescriptor, opts v1.UpdateOptions) (result *v1alpha1.MetricDescriptor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(metricdescriptorsResource, c.ns, metricDescriptor), &v1alpha1.MetricDescriptor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricDescriptor), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMetricDescriptors) UpdateStatus(ctx context.Context, metricDescriptor *v1alpha1.MetricDescriptor, opts v1.UpdateOptions) (*v1alpha1.MetricDescriptor, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(metricdescriptorsResource, "status", c.ns, metricDescriptor), &v1alpha1.MetricDescriptor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricDescriptor), err
}

// Delete takes name of the metricDescriptor and deletes it. Returns an error if one occurs.
func (c *FakeMetricDescriptors) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(metricdescriptorsResource, c.ns, name), &v1alpha1.MetricDescriptor{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMetricDescriptors) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(metricdescriptorsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MetricDescriptorList{})
	return err
}

// Patch applies the patch and returns the patched metricDescriptor.
func (c *FakeMetricDescriptors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MetricDescriptor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(metricdescriptorsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MetricDescriptor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricDescriptor), err
}
