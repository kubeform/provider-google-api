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

	v1alpha1 "kubeform.dev/provider-google-api/apis/scc/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNotificationConfigs implements NotificationConfigInterface
type FakeNotificationConfigs struct {
	Fake *FakeSccV1alpha1
	ns   string
}

var notificationconfigsResource = schema.GroupVersionResource{Group: "scc.google.kubeform.com", Version: "v1alpha1", Resource: "notificationconfigs"}

var notificationconfigsKind = schema.GroupVersionKind{Group: "scc.google.kubeform.com", Version: "v1alpha1", Kind: "NotificationConfig"}

// Get takes name of the notificationConfig, and returns the corresponding notificationConfig object, and an error if there is any.
func (c *FakeNotificationConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NotificationConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(notificationconfigsResource, c.ns, name), &v1alpha1.NotificationConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NotificationConfig), err
}

// List takes label and field selectors, and returns the list of NotificationConfigs that match those selectors.
func (c *FakeNotificationConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NotificationConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(notificationconfigsResource, notificationconfigsKind, c.ns, opts), &v1alpha1.NotificationConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NotificationConfigList{ListMeta: obj.(*v1alpha1.NotificationConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.NotificationConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested notificationConfigs.
func (c *FakeNotificationConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(notificationconfigsResource, c.ns, opts))

}

// Create takes the representation of a notificationConfig and creates it.  Returns the server's representation of the notificationConfig, and an error, if there is any.
func (c *FakeNotificationConfigs) Create(ctx context.Context, notificationConfig *v1alpha1.NotificationConfig, opts v1.CreateOptions) (result *v1alpha1.NotificationConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(notificationconfigsResource, c.ns, notificationConfig), &v1alpha1.NotificationConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NotificationConfig), err
}

// Update takes the representation of a notificationConfig and updates it. Returns the server's representation of the notificationConfig, and an error, if there is any.
func (c *FakeNotificationConfigs) Update(ctx context.Context, notificationConfig *v1alpha1.NotificationConfig, opts v1.UpdateOptions) (result *v1alpha1.NotificationConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(notificationconfigsResource, c.ns, notificationConfig), &v1alpha1.NotificationConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NotificationConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNotificationConfigs) UpdateStatus(ctx context.Context, notificationConfig *v1alpha1.NotificationConfig, opts v1.UpdateOptions) (*v1alpha1.NotificationConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(notificationconfigsResource, "status", c.ns, notificationConfig), &v1alpha1.NotificationConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NotificationConfig), err
}

// Delete takes name of the notificationConfig and deletes it. Returns an error if one occurs.
func (c *FakeNotificationConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(notificationconfigsResource, c.ns, name), &v1alpha1.NotificationConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNotificationConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(notificationconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NotificationConfigList{})
	return err
}

// Patch applies the patch and returns the patched notificationConfig.
func (c *FakeNotificationConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NotificationConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(notificationconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.NotificationConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NotificationConfig), err
}