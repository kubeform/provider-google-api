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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubeform.dev/provider-google-api/apis/logging/v1alpha1"
)

// BillingAccountBucketConfigLister helps list BillingAccountBucketConfigs.
// All objects returned here must be treated as read-only.
type BillingAccountBucketConfigLister interface {
	// List lists all BillingAccountBucketConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BillingAccountBucketConfig, err error)
	// BillingAccountBucketConfigs returns an object that can list and get BillingAccountBucketConfigs.
	BillingAccountBucketConfigs(namespace string) BillingAccountBucketConfigNamespaceLister
	BillingAccountBucketConfigListerExpansion
}

// billingAccountBucketConfigLister implements the BillingAccountBucketConfigLister interface.
type billingAccountBucketConfigLister struct {
	indexer cache.Indexer
}

// NewBillingAccountBucketConfigLister returns a new BillingAccountBucketConfigLister.
func NewBillingAccountBucketConfigLister(indexer cache.Indexer) BillingAccountBucketConfigLister {
	return &billingAccountBucketConfigLister{indexer: indexer}
}

// List lists all BillingAccountBucketConfigs in the indexer.
func (s *billingAccountBucketConfigLister) List(selector labels.Selector) (ret []*v1alpha1.BillingAccountBucketConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BillingAccountBucketConfig))
	})
	return ret, err
}

// BillingAccountBucketConfigs returns an object that can list and get BillingAccountBucketConfigs.
func (s *billingAccountBucketConfigLister) BillingAccountBucketConfigs(namespace string) BillingAccountBucketConfigNamespaceLister {
	return billingAccountBucketConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BillingAccountBucketConfigNamespaceLister helps list and get BillingAccountBucketConfigs.
// All objects returned here must be treated as read-only.
type BillingAccountBucketConfigNamespaceLister interface {
	// List lists all BillingAccountBucketConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BillingAccountBucketConfig, err error)
	// Get retrieves the BillingAccountBucketConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BillingAccountBucketConfig, error)
	BillingAccountBucketConfigNamespaceListerExpansion
}

// billingAccountBucketConfigNamespaceLister implements the BillingAccountBucketConfigNamespaceLister
// interface.
type billingAccountBucketConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BillingAccountBucketConfigs in the indexer for a given namespace.
func (s billingAccountBucketConfigNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BillingAccountBucketConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BillingAccountBucketConfig))
	})
	return ret, err
}

// Get retrieves the BillingAccountBucketConfig from the indexer for a given namespace and name.
func (s billingAccountBucketConfigNamespaceLister) Get(name string) (*v1alpha1.BillingAccountBucketConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("billingaccountbucketconfig"), name)
	}
	return obj.(*v1alpha1.BillingAccountBucketConfig), nil
}
