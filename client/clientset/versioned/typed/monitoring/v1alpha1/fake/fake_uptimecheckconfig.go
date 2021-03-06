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

// FakeUptimeCheckConfigs implements UptimeCheckConfigInterface
type FakeUptimeCheckConfigs struct {
	Fake *FakeMonitoringV1alpha1
	ns   string
}

var uptimecheckconfigsResource = schema.GroupVersionResource{Group: "monitoring.google.kubeform.com", Version: "v1alpha1", Resource: "uptimecheckconfigs"}

var uptimecheckconfigsKind = schema.GroupVersionKind{Group: "monitoring.google.kubeform.com", Version: "v1alpha1", Kind: "UptimeCheckConfig"}

// Get takes name of the uptimeCheckConfig, and returns the corresponding uptimeCheckConfig object, and an error if there is any.
func (c *FakeUptimeCheckConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.UptimeCheckConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(uptimecheckconfigsResource, c.ns, name), &v1alpha1.UptimeCheckConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UptimeCheckConfig), err
}

// List takes label and field selectors, and returns the list of UptimeCheckConfigs that match those selectors.
func (c *FakeUptimeCheckConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UptimeCheckConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(uptimecheckconfigsResource, uptimecheckconfigsKind, c.ns, opts), &v1alpha1.UptimeCheckConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.UptimeCheckConfigList{ListMeta: obj.(*v1alpha1.UptimeCheckConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.UptimeCheckConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested uptimeCheckConfigs.
func (c *FakeUptimeCheckConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(uptimecheckconfigsResource, c.ns, opts))

}

// Create takes the representation of a uptimeCheckConfig and creates it.  Returns the server's representation of the uptimeCheckConfig, and an error, if there is any.
func (c *FakeUptimeCheckConfigs) Create(ctx context.Context, uptimeCheckConfig *v1alpha1.UptimeCheckConfig, opts v1.CreateOptions) (result *v1alpha1.UptimeCheckConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(uptimecheckconfigsResource, c.ns, uptimeCheckConfig), &v1alpha1.UptimeCheckConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UptimeCheckConfig), err
}

// Update takes the representation of a uptimeCheckConfig and updates it. Returns the server's representation of the uptimeCheckConfig, and an error, if there is any.
func (c *FakeUptimeCheckConfigs) Update(ctx context.Context, uptimeCheckConfig *v1alpha1.UptimeCheckConfig, opts v1.UpdateOptions) (result *v1alpha1.UptimeCheckConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(uptimecheckconfigsResource, c.ns, uptimeCheckConfig), &v1alpha1.UptimeCheckConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UptimeCheckConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUptimeCheckConfigs) UpdateStatus(ctx context.Context, uptimeCheckConfig *v1alpha1.UptimeCheckConfig, opts v1.UpdateOptions) (*v1alpha1.UptimeCheckConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(uptimecheckconfigsResource, "status", c.ns, uptimeCheckConfig), &v1alpha1.UptimeCheckConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UptimeCheckConfig), err
}

// Delete takes name of the uptimeCheckConfig and deletes it. Returns an error if one occurs.
func (c *FakeUptimeCheckConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(uptimecheckconfigsResource, c.ns, name), &v1alpha1.UptimeCheckConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUptimeCheckConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(uptimecheckconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.UptimeCheckConfigList{})
	return err
}

// Patch applies the patch and returns the patched uptimeCheckConfig.
func (c *FakeUptimeCheckConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UptimeCheckConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(uptimecheckconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.UptimeCheckConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UptimeCheckConfig), err
}
