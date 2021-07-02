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

	v1alpha1 "kubeform.dev/provider-google-api/apis/os/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConfigPatchDeployments implements ConfigPatchDeploymentInterface
type FakeConfigPatchDeployments struct {
	Fake *FakeOsV1alpha1
	ns   string
}

var configpatchdeploymentsResource = schema.GroupVersionResource{Group: "os.google.kubeform.com", Version: "v1alpha1", Resource: "configpatchdeployments"}

var configpatchdeploymentsKind = schema.GroupVersionKind{Group: "os.google.kubeform.com", Version: "v1alpha1", Kind: "ConfigPatchDeployment"}

// Get takes name of the configPatchDeployment, and returns the corresponding configPatchDeployment object, and an error if there is any.
func (c *FakeConfigPatchDeployments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ConfigPatchDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(configpatchdeploymentsResource, c.ns, name), &v1alpha1.ConfigPatchDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigPatchDeployment), err
}

// List takes label and field selectors, and returns the list of ConfigPatchDeployments that match those selectors.
func (c *FakeConfigPatchDeployments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ConfigPatchDeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(configpatchdeploymentsResource, configpatchdeploymentsKind, c.ns, opts), &v1alpha1.ConfigPatchDeploymentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ConfigPatchDeploymentList{ListMeta: obj.(*v1alpha1.ConfigPatchDeploymentList).ListMeta}
	for _, item := range obj.(*v1alpha1.ConfigPatchDeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested configPatchDeployments.
func (c *FakeConfigPatchDeployments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(configpatchdeploymentsResource, c.ns, opts))

}

// Create takes the representation of a configPatchDeployment and creates it.  Returns the server's representation of the configPatchDeployment, and an error, if there is any.
func (c *FakeConfigPatchDeployments) Create(ctx context.Context, configPatchDeployment *v1alpha1.ConfigPatchDeployment, opts v1.CreateOptions) (result *v1alpha1.ConfigPatchDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(configpatchdeploymentsResource, c.ns, configPatchDeployment), &v1alpha1.ConfigPatchDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigPatchDeployment), err
}

// Update takes the representation of a configPatchDeployment and updates it. Returns the server's representation of the configPatchDeployment, and an error, if there is any.
func (c *FakeConfigPatchDeployments) Update(ctx context.Context, configPatchDeployment *v1alpha1.ConfigPatchDeployment, opts v1.UpdateOptions) (result *v1alpha1.ConfigPatchDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(configpatchdeploymentsResource, c.ns, configPatchDeployment), &v1alpha1.ConfigPatchDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigPatchDeployment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeConfigPatchDeployments) UpdateStatus(ctx context.Context, configPatchDeployment *v1alpha1.ConfigPatchDeployment, opts v1.UpdateOptions) (*v1alpha1.ConfigPatchDeployment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(configpatchdeploymentsResource, "status", c.ns, configPatchDeployment), &v1alpha1.ConfigPatchDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigPatchDeployment), err
}

// Delete takes name of the configPatchDeployment and deletes it. Returns an error if one occurs.
func (c *FakeConfigPatchDeployments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(configpatchdeploymentsResource, c.ns, name), &v1alpha1.ConfigPatchDeployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConfigPatchDeployments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(configpatchdeploymentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ConfigPatchDeploymentList{})
	return err
}

// Patch applies the patch and returns the patched configPatchDeployment.
func (c *FakeConfigPatchDeployments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ConfigPatchDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(configpatchdeploymentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ConfigPatchDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigPatchDeployment), err
}
