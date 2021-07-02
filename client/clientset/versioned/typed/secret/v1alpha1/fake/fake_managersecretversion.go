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

	v1alpha1 "kubeform.dev/provider-google-api/apis/secret/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeManagerSecretVersions implements ManagerSecretVersionInterface
type FakeManagerSecretVersions struct {
	Fake *FakeSecretV1alpha1
	ns   string
}

var managersecretversionsResource = schema.GroupVersionResource{Group: "secret.google.kubeform.com", Version: "v1alpha1", Resource: "managersecretversions"}

var managersecretversionsKind = schema.GroupVersionKind{Group: "secret.google.kubeform.com", Version: "v1alpha1", Kind: "ManagerSecretVersion"}

// Get takes name of the managerSecretVersion, and returns the corresponding managerSecretVersion object, and an error if there is any.
func (c *FakeManagerSecretVersions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ManagerSecretVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(managersecretversionsResource, c.ns, name), &v1alpha1.ManagerSecretVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerSecretVersion), err
}

// List takes label and field selectors, and returns the list of ManagerSecretVersions that match those selectors.
func (c *FakeManagerSecretVersions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ManagerSecretVersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(managersecretversionsResource, managersecretversionsKind, c.ns, opts), &v1alpha1.ManagerSecretVersionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ManagerSecretVersionList{ListMeta: obj.(*v1alpha1.ManagerSecretVersionList).ListMeta}
	for _, item := range obj.(*v1alpha1.ManagerSecretVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managerSecretVersions.
func (c *FakeManagerSecretVersions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(managersecretversionsResource, c.ns, opts))

}

// Create takes the representation of a managerSecretVersion and creates it.  Returns the server's representation of the managerSecretVersion, and an error, if there is any.
func (c *FakeManagerSecretVersions) Create(ctx context.Context, managerSecretVersion *v1alpha1.ManagerSecretVersion, opts v1.CreateOptions) (result *v1alpha1.ManagerSecretVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(managersecretversionsResource, c.ns, managerSecretVersion), &v1alpha1.ManagerSecretVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerSecretVersion), err
}

// Update takes the representation of a managerSecretVersion and updates it. Returns the server's representation of the managerSecretVersion, and an error, if there is any.
func (c *FakeManagerSecretVersions) Update(ctx context.Context, managerSecretVersion *v1alpha1.ManagerSecretVersion, opts v1.UpdateOptions) (result *v1alpha1.ManagerSecretVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(managersecretversionsResource, c.ns, managerSecretVersion), &v1alpha1.ManagerSecretVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerSecretVersion), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagerSecretVersions) UpdateStatus(ctx context.Context, managerSecretVersion *v1alpha1.ManagerSecretVersion, opts v1.UpdateOptions) (*v1alpha1.ManagerSecretVersion, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(managersecretversionsResource, "status", c.ns, managerSecretVersion), &v1alpha1.ManagerSecretVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerSecretVersion), err
}

// Delete takes name of the managerSecretVersion and deletes it. Returns an error if one occurs.
func (c *FakeManagerSecretVersions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(managersecretversionsResource, c.ns, name), &v1alpha1.ManagerSecretVersion{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagerSecretVersions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(managersecretversionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ManagerSecretVersionList{})
	return err
}

// Patch applies the patch and returns the patched managerSecretVersion.
func (c *FakeManagerSecretVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagerSecretVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(managersecretversionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ManagerSecretVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerSecretVersion), err
}
