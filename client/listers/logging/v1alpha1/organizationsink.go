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

// OrganizationSinkLister helps list OrganizationSinks.
// All objects returned here must be treated as read-only.
type OrganizationSinkLister interface {
	// List lists all OrganizationSinks in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OrganizationSink, err error)
	// OrganizationSinks returns an object that can list and get OrganizationSinks.
	OrganizationSinks(namespace string) OrganizationSinkNamespaceLister
	OrganizationSinkListerExpansion
}

// organizationSinkLister implements the OrganizationSinkLister interface.
type organizationSinkLister struct {
	indexer cache.Indexer
}

// NewOrganizationSinkLister returns a new OrganizationSinkLister.
func NewOrganizationSinkLister(indexer cache.Indexer) OrganizationSinkLister {
	return &organizationSinkLister{indexer: indexer}
}

// List lists all OrganizationSinks in the indexer.
func (s *organizationSinkLister) List(selector labels.Selector) (ret []*v1alpha1.OrganizationSink, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OrganizationSink))
	})
	return ret, err
}

// OrganizationSinks returns an object that can list and get OrganizationSinks.
func (s *organizationSinkLister) OrganizationSinks(namespace string) OrganizationSinkNamespaceLister {
	return organizationSinkNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OrganizationSinkNamespaceLister helps list and get OrganizationSinks.
// All objects returned here must be treated as read-only.
type OrganizationSinkNamespaceLister interface {
	// List lists all OrganizationSinks in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OrganizationSink, err error)
	// Get retrieves the OrganizationSink from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.OrganizationSink, error)
	OrganizationSinkNamespaceListerExpansion
}

// organizationSinkNamespaceLister implements the OrganizationSinkNamespaceLister
// interface.
type organizationSinkNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OrganizationSinks in the indexer for a given namespace.
func (s organizationSinkNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.OrganizationSink, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OrganizationSink))
	})
	return ret, err
}

// Get retrieves the OrganizationSink from the indexer for a given namespace and name.
func (s organizationSinkNamespaceLister) Get(name string) (*v1alpha1.OrganizationSink, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("organizationsink"), name)
	}
	return obj.(*v1alpha1.OrganizationSink), nil
}
