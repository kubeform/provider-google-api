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

	v1alpha1 "kubeform.dev/provider-google-api/apis/dialogflow/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCxPages implements CxPageInterface
type FakeCxPages struct {
	Fake *FakeDialogflowV1alpha1
	ns   string
}

var cxpagesResource = schema.GroupVersionResource{Group: "dialogflow.google.kubeform.com", Version: "v1alpha1", Resource: "cxpages"}

var cxpagesKind = schema.GroupVersionKind{Group: "dialogflow.google.kubeform.com", Version: "v1alpha1", Kind: "CxPage"}

// Get takes name of the cxPage, and returns the corresponding cxPage object, and an error if there is any.
func (c *FakeCxPages) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CxPage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cxpagesResource, c.ns, name), &v1alpha1.CxPage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CxPage), err
}

// List takes label and field selectors, and returns the list of CxPages that match those selectors.
func (c *FakeCxPages) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CxPageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cxpagesResource, cxpagesKind, c.ns, opts), &v1alpha1.CxPageList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CxPageList{ListMeta: obj.(*v1alpha1.CxPageList).ListMeta}
	for _, item := range obj.(*v1alpha1.CxPageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cxPages.
func (c *FakeCxPages) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cxpagesResource, c.ns, opts))

}

// Create takes the representation of a cxPage and creates it.  Returns the server's representation of the cxPage, and an error, if there is any.
func (c *FakeCxPages) Create(ctx context.Context, cxPage *v1alpha1.CxPage, opts v1.CreateOptions) (result *v1alpha1.CxPage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cxpagesResource, c.ns, cxPage), &v1alpha1.CxPage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CxPage), err
}

// Update takes the representation of a cxPage and updates it. Returns the server's representation of the cxPage, and an error, if there is any.
func (c *FakeCxPages) Update(ctx context.Context, cxPage *v1alpha1.CxPage, opts v1.UpdateOptions) (result *v1alpha1.CxPage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cxpagesResource, c.ns, cxPage), &v1alpha1.CxPage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CxPage), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCxPages) UpdateStatus(ctx context.Context, cxPage *v1alpha1.CxPage, opts v1.UpdateOptions) (*v1alpha1.CxPage, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cxpagesResource, "status", c.ns, cxPage), &v1alpha1.CxPage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CxPage), err
}

// Delete takes name of the cxPage and deletes it. Returns an error if one occurs.
func (c *FakeCxPages) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cxpagesResource, c.ns, name), &v1alpha1.CxPage{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCxPages) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cxpagesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CxPageList{})
	return err
}

// Patch applies the patch and returns the patched cxPage.
func (c *FakeCxPages) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CxPage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cxpagesResource, c.ns, name, pt, data, subresources...), &v1alpha1.CxPage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CxPage), err
}
