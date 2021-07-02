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
	v1alpha1 "kubeform.dev/provider-google-api/apis/logging/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OrganizationBucketConfigLister helps list OrganizationBucketConfigs.
// All objects returned here must be treated as read-only.
type OrganizationBucketConfigLister interface {
	// List lists all OrganizationBucketConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OrganizationBucketConfig, err error)
	// OrganizationBucketConfigs returns an object that can list and get OrganizationBucketConfigs.
	OrganizationBucketConfigs(namespace string) OrganizationBucketConfigNamespaceLister
	OrganizationBucketConfigListerExpansion
}

// organizationBucketConfigLister implements the OrganizationBucketConfigLister interface.
type organizationBucketConfigLister struct {
	indexer cache.Indexer
}

// NewOrganizationBucketConfigLister returns a new OrganizationBucketConfigLister.
func NewOrganizationBucketConfigLister(indexer cache.Indexer) OrganizationBucketConfigLister {
	return &organizationBucketConfigLister{indexer: indexer}
}

// List lists all OrganizationBucketConfigs in the indexer.
func (s *organizationBucketConfigLister) List(selector labels.Selector) (ret []*v1alpha1.OrganizationBucketConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OrganizationBucketConfig))
	})
	return ret, err
}

// OrganizationBucketConfigs returns an object that can list and get OrganizationBucketConfigs.
func (s *organizationBucketConfigLister) OrganizationBucketConfigs(namespace string) OrganizationBucketConfigNamespaceLister {
	return organizationBucketConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OrganizationBucketConfigNamespaceLister helps list and get OrganizationBucketConfigs.
// All objects returned here must be treated as read-only.
type OrganizationBucketConfigNamespaceLister interface {
	// List lists all OrganizationBucketConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OrganizationBucketConfig, err error)
	// Get retrieves the OrganizationBucketConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.OrganizationBucketConfig, error)
	OrganizationBucketConfigNamespaceListerExpansion
}

// organizationBucketConfigNamespaceLister implements the OrganizationBucketConfigNamespaceLister
// interface.
type organizationBucketConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OrganizationBucketConfigs in the indexer for a given namespace.
func (s organizationBucketConfigNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.OrganizationBucketConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OrganizationBucketConfig))
	})
	return ret, err
}

// Get retrieves the OrganizationBucketConfig from the indexer for a given namespace and name.
func (s organizationBucketConfigNamespaceLister) Get(name string) (*v1alpha1.OrganizationBucketConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("organizationbucketconfig"), name)
	}
	return obj.(*v1alpha1.OrganizationBucketConfig), nil
}
