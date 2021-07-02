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

	v1alpha1 "kubeform.dev/provider-google-api/apis/runtimeconfig/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConfigIamBindings implements ConfigIamBindingInterface
type FakeConfigIamBindings struct {
	Fake *FakeRuntimeconfigV1alpha1
	ns   string
}

var configiambindingsResource = schema.GroupVersionResource{Group: "runtimeconfig.google.kubeform.com", Version: "v1alpha1", Resource: "configiambindings"}

var configiambindingsKind = schema.GroupVersionKind{Group: "runtimeconfig.google.kubeform.com", Version: "v1alpha1", Kind: "ConfigIamBinding"}

// Get takes name of the configIamBinding, and returns the corresponding configIamBinding object, and an error if there is any.
func (c *FakeConfigIamBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ConfigIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(configiambindingsResource, c.ns, name), &v1alpha1.ConfigIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigIamBinding), err
}

// List takes label and field selectors, and returns the list of ConfigIamBindings that match those selectors.
func (c *FakeConfigIamBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ConfigIamBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(configiambindingsResource, configiambindingsKind, c.ns, opts), &v1alpha1.ConfigIamBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ConfigIamBindingList{ListMeta: obj.(*v1alpha1.ConfigIamBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.ConfigIamBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested configIamBindings.
func (c *FakeConfigIamBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(configiambindingsResource, c.ns, opts))

}

// Create takes the representation of a configIamBinding and creates it.  Returns the server's representation of the configIamBinding, and an error, if there is any.
func (c *FakeConfigIamBindings) Create(ctx context.Context, configIamBinding *v1alpha1.ConfigIamBinding, opts v1.CreateOptions) (result *v1alpha1.ConfigIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(configiambindingsResource, c.ns, configIamBinding), &v1alpha1.ConfigIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigIamBinding), err
}

// Update takes the representation of a configIamBinding and updates it. Returns the server's representation of the configIamBinding, and an error, if there is any.
func (c *FakeConfigIamBindings) Update(ctx context.Context, configIamBinding *v1alpha1.ConfigIamBinding, opts v1.UpdateOptions) (result *v1alpha1.ConfigIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(configiambindingsResource, c.ns, configIamBinding), &v1alpha1.ConfigIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigIamBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeConfigIamBindings) UpdateStatus(ctx context.Context, configIamBinding *v1alpha1.ConfigIamBinding, opts v1.UpdateOptions) (*v1alpha1.ConfigIamBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(configiambindingsResource, "status", c.ns, configIamBinding), &v1alpha1.ConfigIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigIamBinding), err
}

// Delete takes name of the configIamBinding and deletes it. Returns an error if one occurs.
func (c *FakeConfigIamBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(configiambindingsResource, c.ns, name), &v1alpha1.ConfigIamBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConfigIamBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(configiambindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ConfigIamBindingList{})
	return err
}

// Patch applies the patch and returns the patched configIamBinding.
func (c *FakeConfigIamBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ConfigIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(configiambindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ConfigIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigIamBinding), err
}
