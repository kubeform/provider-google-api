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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-google-api/apis/container/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AwsNodePoolLister helps list AwsNodePools.
// All objects returned here must be treated as read-only.
type AwsNodePoolLister interface {
	// List lists all AwsNodePools in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AwsNodePool, err error)
	// AwsNodePools returns an object that can list and get AwsNodePools.
	AwsNodePools(namespace string) AwsNodePoolNamespaceLister
	AwsNodePoolListerExpansion
}

// awsNodePoolLister implements the AwsNodePoolLister interface.
type awsNodePoolLister struct {
	indexer cache.Indexer
}

// NewAwsNodePoolLister returns a new AwsNodePoolLister.
func NewAwsNodePoolLister(indexer cache.Indexer) AwsNodePoolLister {
	return &awsNodePoolLister{indexer: indexer}
}

// List lists all AwsNodePools in the indexer.
func (s *awsNodePoolLister) List(selector labels.Selector) (ret []*v1alpha1.AwsNodePool, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsNodePool))
	})
	return ret, err
}

// AwsNodePools returns an object that can list and get AwsNodePools.
func (s *awsNodePoolLister) AwsNodePools(namespace string) AwsNodePoolNamespaceLister {
	return awsNodePoolNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsNodePoolNamespaceLister helps list and get AwsNodePools.
// All objects returned here must be treated as read-only.
type AwsNodePoolNamespaceLister interface {
	// List lists all AwsNodePools in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AwsNodePool, err error)
	// Get retrieves the AwsNodePool from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AwsNodePool, error)
	AwsNodePoolNamespaceListerExpansion
}

// awsNodePoolNamespaceLister implements the AwsNodePoolNamespaceLister
// interface.
type awsNodePoolNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsNodePools in the indexer for a given namespace.
func (s awsNodePoolNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AwsNodePool, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsNodePool))
	})
	return ret, err
}

// Get retrieves the AwsNodePool from the indexer for a given namespace and name.
func (s awsNodePoolNamespaceLister) Get(name string) (*v1alpha1.AwsNodePool, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsnodepool"), name)
	}
	return obj.(*v1alpha1.AwsNodePool), nil
}
