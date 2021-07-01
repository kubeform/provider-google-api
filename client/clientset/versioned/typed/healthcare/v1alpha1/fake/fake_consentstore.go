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
	v1alpha1 "kubeform.dev/provider-google-api/apis/healthcare/v1alpha1"
)

// FakeConsentStores implements ConsentStoreInterface
type FakeConsentStores struct {
	Fake *FakeHealthcareV1alpha1
	ns   string
}

var consentstoresResource = schema.GroupVersionResource{Group: "healthcare.google.kubeform.com", Version: "v1alpha1", Resource: "consentstores"}

var consentstoresKind = schema.GroupVersionKind{Group: "healthcare.google.kubeform.com", Version: "v1alpha1", Kind: "ConsentStore"}

// Get takes name of the consentStore, and returns the corresponding consentStore object, and an error if there is any.
func (c *FakeConsentStores) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ConsentStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(consentstoresResource, c.ns, name), &v1alpha1.ConsentStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConsentStore), err
}

// List takes label and field selectors, and returns the list of ConsentStores that match those selectors.
func (c *FakeConsentStores) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ConsentStoreList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(consentstoresResource, consentstoresKind, c.ns, opts), &v1alpha1.ConsentStoreList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ConsentStoreList{ListMeta: obj.(*v1alpha1.ConsentStoreList).ListMeta}
	for _, item := range obj.(*v1alpha1.ConsentStoreList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested consentStores.
func (c *FakeConsentStores) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(consentstoresResource, c.ns, opts))

}

// Create takes the representation of a consentStore and creates it.  Returns the server's representation of the consentStore, and an error, if there is any.
func (c *FakeConsentStores) Create(ctx context.Context, consentStore *v1alpha1.ConsentStore, opts v1.CreateOptions) (result *v1alpha1.ConsentStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(consentstoresResource, c.ns, consentStore), &v1alpha1.ConsentStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConsentStore), err
}

// Update takes the representation of a consentStore and updates it. Returns the server's representation of the consentStore, and an error, if there is any.
func (c *FakeConsentStores) Update(ctx context.Context, consentStore *v1alpha1.ConsentStore, opts v1.UpdateOptions) (result *v1alpha1.ConsentStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(consentstoresResource, c.ns, consentStore), &v1alpha1.ConsentStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConsentStore), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeConsentStores) UpdateStatus(ctx context.Context, consentStore *v1alpha1.ConsentStore, opts v1.UpdateOptions) (*v1alpha1.ConsentStore, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(consentstoresResource, "status", c.ns, consentStore), &v1alpha1.ConsentStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConsentStore), err
}

// Delete takes name of the consentStore and deletes it. Returns an error if one occurs.
func (c *FakeConsentStores) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(consentstoresResource, c.ns, name), &v1alpha1.ConsentStore{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConsentStores) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(consentstoresResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ConsentStoreList{})
	return err
}

// Patch applies the patch and returns the patched consentStore.
func (c *FakeConsentStores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ConsentStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(consentstoresResource, c.ns, name, pt, data, subresources...), &v1alpha1.ConsentStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConsentStore), err
}
