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

// AzureNodePoolLister helps list AzureNodePools.
// All objects returned here must be treated as read-only.
type AzureNodePoolLister interface {
	// List lists all AzureNodePools in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AzureNodePool, err error)
	// AzureNodePools returns an object that can list and get AzureNodePools.
	AzureNodePools(namespace string) AzureNodePoolNamespaceLister
	AzureNodePoolListerExpansion
}

// azureNodePoolLister implements the AzureNodePoolLister interface.
type azureNodePoolLister struct {
	indexer cache.Indexer
}

// NewAzureNodePoolLister returns a new AzureNodePoolLister.
func NewAzureNodePoolLister(indexer cache.Indexer) AzureNodePoolLister {
	return &azureNodePoolLister{indexer: indexer}
}

// List lists all AzureNodePools in the indexer.
func (s *azureNodePoolLister) List(selector labels.Selector) (ret []*v1alpha1.AzureNodePool, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AzureNodePool))
	})
	return ret, err
}

// AzureNodePools returns an object that can list and get AzureNodePools.
func (s *azureNodePoolLister) AzureNodePools(namespace string) AzureNodePoolNamespaceLister {
	return azureNodePoolNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AzureNodePoolNamespaceLister helps list and get AzureNodePools.
// All objects returned here must be treated as read-only.
type AzureNodePoolNamespaceLister interface {
	// List lists all AzureNodePools in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AzureNodePool, err error)
	// Get retrieves the AzureNodePool from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AzureNodePool, error)
	AzureNodePoolNamespaceListerExpansion
}

// azureNodePoolNamespaceLister implements the AzureNodePoolNamespaceLister
// interface.
type azureNodePoolNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AzureNodePools in the indexer for a given namespace.
func (s azureNodePoolNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AzureNodePool, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AzureNodePool))
	})
	return ret, err
}

// Get retrieves the AzureNodePool from the indexer for a given namespace and name.
func (s azureNodePoolNamespaceLister) Get(name string) (*v1alpha1.AzureNodePool, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("azurenodepool"), name)
	}
	return obj.(*v1alpha1.AzureNodePool), nil
}