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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/provider-google-api/apis/logging/v1alpha1"
)

// FakeFolderBucketConfigs implements FolderBucketConfigInterface
type FakeFolderBucketConfigs struct {
	Fake *FakeLoggingV1alpha1
	ns   string
}

var folderbucketconfigsResource = schema.GroupVersionResource{Group: "logging.google.kubeform.com", Version: "v1alpha1", Resource: "folderbucketconfigs"}

var folderbucketconfigsKind = schema.GroupVersionKind{Group: "logging.google.kubeform.com", Version: "v1alpha1", Kind: "FolderBucketConfig"}

// Get takes name of the folderBucketConfig, and returns the corresponding folderBucketConfig object, and an error if there is any.
func (c *FakeFolderBucketConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FolderBucketConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(folderbucketconfigsResource, c.ns, name), &v1alpha1.FolderBucketConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FolderBucketConfig), err
}

// List takes label and field selectors, and returns the list of FolderBucketConfigs that match those selectors.
func (c *FakeFolderBucketConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FolderBucketConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(folderbucketconfigsResource, folderbucketconfigsKind, c.ns, opts), &v1alpha1.FolderBucketConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FolderBucketConfigList{ListMeta: obj.(*v1alpha1.FolderBucketConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.FolderBucketConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested folderBucketConfigs.
func (c *FakeFolderBucketConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(folderbucketconfigsResource, c.ns, opts))

}

// Create takes the representation of a folderBucketConfig and creates it.  Returns the server's representation of the folderBucketConfig, and an error, if there is any.
func (c *FakeFolderBucketConfigs) Create(ctx context.Context, folderBucketConfig *v1alpha1.FolderBucketConfig, opts v1.CreateOptions) (result *v1alpha1.FolderBucketConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(folderbucketconfigsResource, c.ns, folderBucketConfig), &v1alpha1.FolderBucketConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FolderBucketConfig), err
}

// Update takes the representation of a folderBucketConfig and updates it. Returns the server's representation of the folderBucketConfig, and an error, if there is any.
func (c *FakeFolderBucketConfigs) Update(ctx context.Context, folderBucketConfig *v1alpha1.FolderBucketConfig, opts v1.UpdateOptions) (result *v1alpha1.FolderBucketConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(folderbucketconfigsResource, c.ns, folderBucketConfig), &v1alpha1.FolderBucketConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FolderBucketConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFolderBucketConfigs) UpdateStatus(ctx context.Context, folderBucketConfig *v1alpha1.FolderBucketConfig, opts v1.UpdateOptions) (*v1alpha1.FolderBucketConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(folderbucketconfigsResource, "status", c.ns, folderBucketConfig), &v1alpha1.FolderBucketConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FolderBucketConfig), err
}

// Delete takes name of the folderBucketConfig and deletes it. Returns an error if one occurs.
func (c *FakeFolderBucketConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(folderbucketconfigsResource, c.ns, name), &v1alpha1.FolderBucketConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFolderBucketConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(folderbucketconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FolderBucketConfigList{})
	return err
}

// Patch applies the patch and returns the patched folderBucketConfig.
func (c *FakeFolderBucketConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FolderBucketConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(folderbucketconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.FolderBucketConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FolderBucketConfig), err
}
