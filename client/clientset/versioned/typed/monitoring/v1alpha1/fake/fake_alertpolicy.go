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
	v1alpha1 "kubeform.dev/provider-google-api/apis/monitoring/v1alpha1"
)

// FakeAlertPolicies implements AlertPolicyInterface
type FakeAlertPolicies struct {
	Fake *FakeMonitoringV1alpha1
	ns   string
}

var alertpoliciesResource = schema.GroupVersionResource{Group: "monitoring.google.kubeform.com", Version: "v1alpha1", Resource: "alertpolicies"}

var alertpoliciesKind = schema.GroupVersionKind{Group: "monitoring.google.kubeform.com", Version: "v1alpha1", Kind: "AlertPolicy"}

// Get takes name of the alertPolicy, and returns the corresponding alertPolicy object, and an error if there is any.
func (c *FakeAlertPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AlertPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(alertpoliciesResource, c.ns, name), &v1alpha1.AlertPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlertPolicy), err
}

// List takes label and field selectors, and returns the list of AlertPolicies that match those selectors.
func (c *FakeAlertPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AlertPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(alertpoliciesResource, alertpoliciesKind, c.ns, opts), &v1alpha1.AlertPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AlertPolicyList{ListMeta: obj.(*v1alpha1.AlertPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.AlertPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested alertPolicies.
func (c *FakeAlertPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(alertpoliciesResource, c.ns, opts))

}

// Create takes the representation of a alertPolicy and creates it.  Returns the server's representation of the alertPolicy, and an error, if there is any.
func (c *FakeAlertPolicies) Create(ctx context.Context, alertPolicy *v1alpha1.AlertPolicy, opts v1.CreateOptions) (result *v1alpha1.AlertPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(alertpoliciesResource, c.ns, alertPolicy), &v1alpha1.AlertPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlertPolicy), err
}

// Update takes the representation of a alertPolicy and updates it. Returns the server's representation of the alertPolicy, and an error, if there is any.
func (c *FakeAlertPolicies) Update(ctx context.Context, alertPolicy *v1alpha1.AlertPolicy, opts v1.UpdateOptions) (result *v1alpha1.AlertPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(alertpoliciesResource, c.ns, alertPolicy), &v1alpha1.AlertPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlertPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAlertPolicies) UpdateStatus(ctx context.Context, alertPolicy *v1alpha1.AlertPolicy, opts v1.UpdateOptions) (*v1alpha1.AlertPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(alertpoliciesResource, "status", c.ns, alertPolicy), &v1alpha1.AlertPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlertPolicy), err
}

// Delete takes name of the alertPolicy and deletes it. Returns an error if one occurs.
func (c *FakeAlertPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(alertpoliciesResource, c.ns, name), &v1alpha1.AlertPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAlertPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(alertpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AlertPolicyList{})
	return err
}

// Patch applies the patch and returns the patched alertPolicy.
func (c *FakeAlertPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AlertPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(alertpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AlertPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlertPolicy), err
}
