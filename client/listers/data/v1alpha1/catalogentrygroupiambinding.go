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
	v1alpha1 "kubeform.dev/provider-google-api/apis/data/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CatalogEntryGroupIamBindingLister helps list CatalogEntryGroupIamBindings.
// All objects returned here must be treated as read-only.
type CatalogEntryGroupIamBindingLister interface {
	// List lists all CatalogEntryGroupIamBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CatalogEntryGroupIamBinding, err error)
	// CatalogEntryGroupIamBindings returns an object that can list and get CatalogEntryGroupIamBindings.
	CatalogEntryGroupIamBindings(namespace string) CatalogEntryGroupIamBindingNamespaceLister
	CatalogEntryGroupIamBindingListerExpansion
}

// catalogEntryGroupIamBindingLister implements the CatalogEntryGroupIamBindingLister interface.
type catalogEntryGroupIamBindingLister struct {
	indexer cache.Indexer
}

// NewCatalogEntryGroupIamBindingLister returns a new CatalogEntryGroupIamBindingLister.
func NewCatalogEntryGroupIamBindingLister(indexer cache.Indexer) CatalogEntryGroupIamBindingLister {
	return &catalogEntryGroupIamBindingLister{indexer: indexer}
}

// List lists all CatalogEntryGroupIamBindings in the indexer.
func (s *catalogEntryGroupIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.CatalogEntryGroupIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CatalogEntryGroupIamBinding))
	})
	return ret, err
}

// CatalogEntryGroupIamBindings returns an object that can list and get CatalogEntryGroupIamBindings.
func (s *catalogEntryGroupIamBindingLister) CatalogEntryGroupIamBindings(namespace string) CatalogEntryGroupIamBindingNamespaceLister {
	return catalogEntryGroupIamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CatalogEntryGroupIamBindingNamespaceLister helps list and get CatalogEntryGroupIamBindings.
// All objects returned here must be treated as read-only.
type CatalogEntryGroupIamBindingNamespaceLister interface {
	// List lists all CatalogEntryGroupIamBindings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CatalogEntryGroupIamBinding, err error)
	// Get retrieves the CatalogEntryGroupIamBinding from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CatalogEntryGroupIamBinding, error)
	CatalogEntryGroupIamBindingNamespaceListerExpansion
}

// catalogEntryGroupIamBindingNamespaceLister implements the CatalogEntryGroupIamBindingNamespaceLister
// interface.
type catalogEntryGroupIamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CatalogEntryGroupIamBindings in the indexer for a given namespace.
func (s catalogEntryGroupIamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CatalogEntryGroupIamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CatalogEntryGroupIamBinding))
	})
	return ret, err
}

// Get retrieves the CatalogEntryGroupIamBinding from the indexer for a given namespace and name.
func (s catalogEntryGroupIamBindingNamespaceLister) Get(name string) (*v1alpha1.CatalogEntryGroupIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("catalogentrygroupiambinding"), name)
	}
	return obj.(*v1alpha1.CatalogEntryGroupIamBinding), nil
}
