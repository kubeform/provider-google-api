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
	v1alpha1 "kubeform.dev/provider-google-api/apis/active/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DirectoryDomainLister helps list DirectoryDomains.
// All objects returned here must be treated as read-only.
type DirectoryDomainLister interface {
	// List lists all DirectoryDomains in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DirectoryDomain, err error)
	// DirectoryDomains returns an object that can list and get DirectoryDomains.
	DirectoryDomains(namespace string) DirectoryDomainNamespaceLister
	DirectoryDomainListerExpansion
}

// directoryDomainLister implements the DirectoryDomainLister interface.
type directoryDomainLister struct {
	indexer cache.Indexer
}

// NewDirectoryDomainLister returns a new DirectoryDomainLister.
func NewDirectoryDomainLister(indexer cache.Indexer) DirectoryDomainLister {
	return &directoryDomainLister{indexer: indexer}
}

// List lists all DirectoryDomains in the indexer.
func (s *directoryDomainLister) List(selector labels.Selector) (ret []*v1alpha1.DirectoryDomain, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DirectoryDomain))
	})
	return ret, err
}

// DirectoryDomains returns an object that can list and get DirectoryDomains.
func (s *directoryDomainLister) DirectoryDomains(namespace string) DirectoryDomainNamespaceLister {
	return directoryDomainNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DirectoryDomainNamespaceLister helps list and get DirectoryDomains.
// All objects returned here must be treated as read-only.
type DirectoryDomainNamespaceLister interface {
	// List lists all DirectoryDomains in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DirectoryDomain, err error)
	// Get retrieves the DirectoryDomain from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DirectoryDomain, error)
	DirectoryDomainNamespaceListerExpansion
}

// directoryDomainNamespaceLister implements the DirectoryDomainNamespaceLister
// interface.
type directoryDomainNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DirectoryDomains in the indexer for a given namespace.
func (s directoryDomainNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DirectoryDomain, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DirectoryDomain))
	})
	return ret, err
}

// Get retrieves the DirectoryDomain from the indexer for a given namespace and name.
func (s directoryDomainNamespaceLister) Get(name string) (*v1alpha1.DirectoryDomain, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("directorydomain"), name)
	}
	return obj.(*v1alpha1.DirectoryDomain), nil
}
