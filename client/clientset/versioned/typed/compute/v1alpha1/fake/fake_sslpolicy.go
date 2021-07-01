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
	v1alpha1 "kubeform.dev/provider-google-api/apis/compute/v1alpha1"
)

// FakeSslPolicies implements SslPolicyInterface
type FakeSslPolicies struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var sslpoliciesResource = schema.GroupVersionResource{Group: "compute.google.kubeform.com", Version: "v1alpha1", Resource: "sslpolicies"}

var sslpoliciesKind = schema.GroupVersionKind{Group: "compute.google.kubeform.com", Version: "v1alpha1", Kind: "SslPolicy"}

// Get takes name of the sslPolicy, and returns the corresponding sslPolicy object, and an error if there is any.
func (c *FakeSslPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SslPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sslpoliciesResource, c.ns, name), &v1alpha1.SslPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SslPolicy), err
}

// List takes label and field selectors, and returns the list of SslPolicies that match those selectors.
func (c *FakeSslPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SslPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sslpoliciesResource, sslpoliciesKind, c.ns, opts), &v1alpha1.SslPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SslPolicyList{ListMeta: obj.(*v1alpha1.SslPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.SslPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sslPolicies.
func (c *FakeSslPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sslpoliciesResource, c.ns, opts))

}

// Create takes the representation of a sslPolicy and creates it.  Returns the server's representation of the sslPolicy, and an error, if there is any.
func (c *FakeSslPolicies) Create(ctx context.Context, sslPolicy *v1alpha1.SslPolicy, opts v1.CreateOptions) (result *v1alpha1.SslPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sslpoliciesResource, c.ns, sslPolicy), &v1alpha1.SslPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SslPolicy), err
}

// Update takes the representation of a sslPolicy and updates it. Returns the server's representation of the sslPolicy, and an error, if there is any.
func (c *FakeSslPolicies) Update(ctx context.Context, sslPolicy *v1alpha1.SslPolicy, opts v1.UpdateOptions) (result *v1alpha1.SslPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sslpoliciesResource, c.ns, sslPolicy), &v1alpha1.SslPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SslPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSslPolicies) UpdateStatus(ctx context.Context, sslPolicy *v1alpha1.SslPolicy, opts v1.UpdateOptions) (*v1alpha1.SslPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sslpoliciesResource, "status", c.ns, sslPolicy), &v1alpha1.SslPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SslPolicy), err
}

// Delete takes name of the sslPolicy and deletes it. Returns an error if one occurs.
func (c *FakeSslPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sslpoliciesResource, c.ns, name), &v1alpha1.SslPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSslPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sslpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SslPolicyList{})
	return err
}

// Patch applies the patch and returns the patched sslPolicy.
func (c *FakeSslPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SslPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sslpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SslPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SslPolicy), err
}
