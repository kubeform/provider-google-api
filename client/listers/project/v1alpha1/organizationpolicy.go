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
	v1alpha1 "kubeform.dev/provider-google-api/apis/project/v1alpha1"
)

// OrganizationPolicyLister helps list OrganizationPolicies.
// All objects returned here must be treated as read-only.
type OrganizationPolicyLister interface {
	// List lists all OrganizationPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OrganizationPolicy, err error)
	// OrganizationPolicies returns an object that can list and get OrganizationPolicies.
	OrganizationPolicies(namespace string) OrganizationPolicyNamespaceLister
	OrganizationPolicyListerExpansion
}

// organizationPolicyLister implements the OrganizationPolicyLister interface.
type organizationPolicyLister struct {
	indexer cache.Indexer
}

// NewOrganizationPolicyLister returns a new OrganizationPolicyLister.
func NewOrganizationPolicyLister(indexer cache.Indexer) OrganizationPolicyLister {
	return &organizationPolicyLister{indexer: indexer}
}

// List lists all OrganizationPolicies in the indexer.
func (s *organizationPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.OrganizationPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OrganizationPolicy))
	})
	return ret, err
}

// OrganizationPolicies returns an object that can list and get OrganizationPolicies.
func (s *organizationPolicyLister) OrganizationPolicies(namespace string) OrganizationPolicyNamespaceLister {
	return organizationPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OrganizationPolicyNamespaceLister helps list and get OrganizationPolicies.
// All objects returned here must be treated as read-only.
type OrganizationPolicyNamespaceLister interface {
	// List lists all OrganizationPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OrganizationPolicy, err error)
	// Get retrieves the OrganizationPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.OrganizationPolicy, error)
	OrganizationPolicyNamespaceListerExpansion
}

// organizationPolicyNamespaceLister implements the OrganizationPolicyNamespaceLister
// interface.
type organizationPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OrganizationPolicies in the indexer for a given namespace.
func (s organizationPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.OrganizationPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OrganizationPolicy))
	})
	return ret, err
}

// Get retrieves the OrganizationPolicy from the indexer for a given namespace and name.
func (s organizationPolicyNamespaceLister) Get(name string) (*v1alpha1.OrganizationPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("organizationpolicy"), name)
	}
	return obj.(*v1alpha1.OrganizationPolicy), nil
}
