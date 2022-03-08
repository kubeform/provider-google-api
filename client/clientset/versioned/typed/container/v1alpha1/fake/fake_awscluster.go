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

	v1alpha1 "kubeform.dev/provider-google-api/apis/container/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAwsClusters implements AwsClusterInterface
type FakeAwsClusters struct {
	Fake *FakeContainerV1alpha1
	ns   string
}

var awsclustersResource = schema.GroupVersionResource{Group: "container.google.kubeform.com", Version: "v1alpha1", Resource: "awsclusters"}

var awsclustersKind = schema.GroupVersionKind{Group: "container.google.kubeform.com", Version: "v1alpha1", Kind: "AwsCluster"}

// Get takes name of the awsCluster, and returns the corresponding awsCluster object, and an error if there is any.
func (c *FakeAwsClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AwsCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsclustersResource, c.ns, name), &v1alpha1.AwsCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCluster), err
}

// List takes label and field selectors, and returns the list of AwsClusters that match those selectors.
func (c *FakeAwsClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AwsClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsclustersResource, awsclustersKind, c.ns, opts), &v1alpha1.AwsClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsClusterList{ListMeta: obj.(*v1alpha1.AwsClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsClusters.
func (c *FakeAwsClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsclustersResource, c.ns, opts))

}

// Create takes the representation of a awsCluster and creates it.  Returns the server's representation of the awsCluster, and an error, if there is any.
func (c *FakeAwsClusters) Create(ctx context.Context, awsCluster *v1alpha1.AwsCluster, opts v1.CreateOptions) (result *v1alpha1.AwsCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsclustersResource, c.ns, awsCluster), &v1alpha1.AwsCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCluster), err
}

// Update takes the representation of a awsCluster and updates it. Returns the server's representation of the awsCluster, and an error, if there is any.
func (c *FakeAwsClusters) Update(ctx context.Context, awsCluster *v1alpha1.AwsCluster, opts v1.UpdateOptions) (result *v1alpha1.AwsCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsclustersResource, c.ns, awsCluster), &v1alpha1.AwsCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsClusters) UpdateStatus(ctx context.Context, awsCluster *v1alpha1.AwsCluster, opts v1.UpdateOptions) (*v1alpha1.AwsCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(awsclustersResource, "status", c.ns, awsCluster), &v1alpha1.AwsCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCluster), err
}

// Delete takes name of the awsCluster and deletes it. Returns an error if one occurs.
func (c *FakeAwsClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsclustersResource, c.ns, name), &v1alpha1.AwsCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsClusterList{})
	return err
}

// Patch applies the patch and returns the patched awsCluster.
func (c *FakeAwsClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AwsCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.AwsCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsCluster), err
}