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
	v1alpha1 "kubeform.dev/provider-google-api/apis/secret/v1alpha1"
)

// FakeManagerSecrets implements ManagerSecretInterface
type FakeManagerSecrets struct {
	Fake *FakeSecretV1alpha1
	ns   string
}

var managersecretsResource = schema.GroupVersionResource{Group: "secret.google.kubeform.com", Version: "v1alpha1", Resource: "managersecrets"}

var managersecretsKind = schema.GroupVersionKind{Group: "secret.google.kubeform.com", Version: "v1alpha1", Kind: "ManagerSecret"}

// Get takes name of the managerSecret, and returns the corresponding managerSecret object, and an error if there is any.
func (c *FakeManagerSecrets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ManagerSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(managersecretsResource, c.ns, name), &v1alpha1.ManagerSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerSecret), err
}

// List takes label and field selectors, and returns the list of ManagerSecrets that match those selectors.
func (c *FakeManagerSecrets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ManagerSecretList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(managersecretsResource, managersecretsKind, c.ns, opts), &v1alpha1.ManagerSecretList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ManagerSecretList{ListMeta: obj.(*v1alpha1.ManagerSecretList).ListMeta}
	for _, item := range obj.(*v1alpha1.ManagerSecretList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managerSecrets.
func (c *FakeManagerSecrets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(managersecretsResource, c.ns, opts))

}

// Create takes the representation of a managerSecret and creates it.  Returns the server's representation of the managerSecret, and an error, if there is any.
func (c *FakeManagerSecrets) Create(ctx context.Context, managerSecret *v1alpha1.ManagerSecret, opts v1.CreateOptions) (result *v1alpha1.ManagerSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(managersecretsResource, c.ns, managerSecret), &v1alpha1.ManagerSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerSecret), err
}

// Update takes the representation of a managerSecret and updates it. Returns the server's representation of the managerSecret, and an error, if there is any.
func (c *FakeManagerSecrets) Update(ctx context.Context, managerSecret *v1alpha1.ManagerSecret, opts v1.UpdateOptions) (result *v1alpha1.ManagerSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(managersecretsResource, c.ns, managerSecret), &v1alpha1.ManagerSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerSecret), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagerSecrets) UpdateStatus(ctx context.Context, managerSecret *v1alpha1.ManagerSecret, opts v1.UpdateOptions) (*v1alpha1.ManagerSecret, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(managersecretsResource, "status", c.ns, managerSecret), &v1alpha1.ManagerSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerSecret), err
}

// Delete takes name of the managerSecret and deletes it. Returns an error if one occurs.
func (c *FakeManagerSecrets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(managersecretsResource, c.ns, name), &v1alpha1.ManagerSecret{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagerSecrets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(managersecretsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ManagerSecretList{})
	return err
}

// Patch applies the patch and returns the patched managerSecret.
func (c *FakeManagerSecrets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagerSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(managersecretsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ManagerSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerSecret), err
}
