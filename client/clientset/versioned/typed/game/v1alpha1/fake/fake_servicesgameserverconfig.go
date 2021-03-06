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

	v1alpha1 "kubeform.dev/provider-google-api/apis/game/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServicesGameServerConfigs implements ServicesGameServerConfigInterface
type FakeServicesGameServerConfigs struct {
	Fake *FakeGameV1alpha1
	ns   string
}

var servicesgameserverconfigsResource = schema.GroupVersionResource{Group: "game.google.kubeform.com", Version: "v1alpha1", Resource: "servicesgameserverconfigs"}

var servicesgameserverconfigsKind = schema.GroupVersionKind{Group: "game.google.kubeform.com", Version: "v1alpha1", Kind: "ServicesGameServerConfig"}

// Get takes name of the servicesGameServerConfig, and returns the corresponding servicesGameServerConfig object, and an error if there is any.
func (c *FakeServicesGameServerConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServicesGameServerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicesgameserverconfigsResource, c.ns, name), &v1alpha1.ServicesGameServerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicesGameServerConfig), err
}

// List takes label and field selectors, and returns the list of ServicesGameServerConfigs that match those selectors.
func (c *FakeServicesGameServerConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServicesGameServerConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicesgameserverconfigsResource, servicesgameserverconfigsKind, c.ns, opts), &v1alpha1.ServicesGameServerConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServicesGameServerConfigList{ListMeta: obj.(*v1alpha1.ServicesGameServerConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServicesGameServerConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested servicesGameServerConfigs.
func (c *FakeServicesGameServerConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicesgameserverconfigsResource, c.ns, opts))

}

// Create takes the representation of a servicesGameServerConfig and creates it.  Returns the server's representation of the servicesGameServerConfig, and an error, if there is any.
func (c *FakeServicesGameServerConfigs) Create(ctx context.Context, servicesGameServerConfig *v1alpha1.ServicesGameServerConfig, opts v1.CreateOptions) (result *v1alpha1.ServicesGameServerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicesgameserverconfigsResource, c.ns, servicesGameServerConfig), &v1alpha1.ServicesGameServerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicesGameServerConfig), err
}

// Update takes the representation of a servicesGameServerConfig and updates it. Returns the server's representation of the servicesGameServerConfig, and an error, if there is any.
func (c *FakeServicesGameServerConfigs) Update(ctx context.Context, servicesGameServerConfig *v1alpha1.ServicesGameServerConfig, opts v1.UpdateOptions) (result *v1alpha1.ServicesGameServerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicesgameserverconfigsResource, c.ns, servicesGameServerConfig), &v1alpha1.ServicesGameServerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicesGameServerConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServicesGameServerConfigs) UpdateStatus(ctx context.Context, servicesGameServerConfig *v1alpha1.ServicesGameServerConfig, opts v1.UpdateOptions) (*v1alpha1.ServicesGameServerConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicesgameserverconfigsResource, "status", c.ns, servicesGameServerConfig), &v1alpha1.ServicesGameServerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicesGameServerConfig), err
}

// Delete takes name of the servicesGameServerConfig and deletes it. Returns an error if one occurs.
func (c *FakeServicesGameServerConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicesgameserverconfigsResource, c.ns, name), &v1alpha1.ServicesGameServerConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServicesGameServerConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicesgameserverconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServicesGameServerConfigList{})
	return err
}

// Patch applies the patch and returns the patched servicesGameServerConfig.
func (c *FakeServicesGameServerConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServicesGameServerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicesgameserverconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ServicesGameServerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicesGameServerConfig), err
}
