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
	v1alpha1 "kubeform.dev/provider-google-api/apis/billingsubaccount/v1alpha1"
)

// FakeBillingSubaccounts implements BillingSubaccountInterface
type FakeBillingSubaccounts struct {
	Fake *FakeBillingsubaccountV1alpha1
	ns   string
}

var billingsubaccountsResource = schema.GroupVersionResource{Group: "billingsubaccount.google.kubeform.com", Version: "v1alpha1", Resource: "billingsubaccounts"}

var billingsubaccountsKind = schema.GroupVersionKind{Group: "billingsubaccount.google.kubeform.com", Version: "v1alpha1", Kind: "BillingSubaccount"}

// Get takes name of the billingSubaccount, and returns the corresponding billingSubaccount object, and an error if there is any.
func (c *FakeBillingSubaccounts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BillingSubaccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(billingsubaccountsResource, c.ns, name), &v1alpha1.BillingSubaccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BillingSubaccount), err
}

// List takes label and field selectors, and returns the list of BillingSubaccounts that match those selectors.
func (c *FakeBillingSubaccounts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BillingSubaccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(billingsubaccountsResource, billingsubaccountsKind, c.ns, opts), &v1alpha1.BillingSubaccountList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BillingSubaccountList{ListMeta: obj.(*v1alpha1.BillingSubaccountList).ListMeta}
	for _, item := range obj.(*v1alpha1.BillingSubaccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested billingSubaccounts.
func (c *FakeBillingSubaccounts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(billingsubaccountsResource, c.ns, opts))

}

// Create takes the representation of a billingSubaccount and creates it.  Returns the server's representation of the billingSubaccount, and an error, if there is any.
func (c *FakeBillingSubaccounts) Create(ctx context.Context, billingSubaccount *v1alpha1.BillingSubaccount, opts v1.CreateOptions) (result *v1alpha1.BillingSubaccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(billingsubaccountsResource, c.ns, billingSubaccount), &v1alpha1.BillingSubaccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BillingSubaccount), err
}

// Update takes the representation of a billingSubaccount and updates it. Returns the server's representation of the billingSubaccount, and an error, if there is any.
func (c *FakeBillingSubaccounts) Update(ctx context.Context, billingSubaccount *v1alpha1.BillingSubaccount, opts v1.UpdateOptions) (result *v1alpha1.BillingSubaccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(billingsubaccountsResource, c.ns, billingSubaccount), &v1alpha1.BillingSubaccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BillingSubaccount), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBillingSubaccounts) UpdateStatus(ctx context.Context, billingSubaccount *v1alpha1.BillingSubaccount, opts v1.UpdateOptions) (*v1alpha1.BillingSubaccount, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(billingsubaccountsResource, "status", c.ns, billingSubaccount), &v1alpha1.BillingSubaccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BillingSubaccount), err
}

// Delete takes name of the billingSubaccount and deletes it. Returns an error if one occurs.
func (c *FakeBillingSubaccounts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(billingsubaccountsResource, c.ns, name), &v1alpha1.BillingSubaccount{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBillingSubaccounts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(billingsubaccountsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BillingSubaccountList{})
	return err
}

// Patch applies the patch and returns the patched billingSubaccount.
func (c *FakeBillingSubaccounts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BillingSubaccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(billingsubaccountsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BillingSubaccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BillingSubaccount), err
}
