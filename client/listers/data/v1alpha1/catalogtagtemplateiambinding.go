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

// CatalogTagTemplateIamBindingLister helps list CatalogTagTemplateIamBindings.
// All objects returned here must be treated as read-only.
type CatalogTagTemplateIamBindingLister interface {
	// List lists all CatalogTagTemplateIamBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CatalogTagTemplateIamBinding, err error)
	// CatalogTagTemplateIamBindings returns an object that can list and get CatalogTagTemplateIamBindings.
	CatalogTagTemplateIamBindings(namespace string) CatalogTagTemplateIamBindingNamespaceLister
	CatalogTagTemplateIamBindingListerExpansion
}

// catalogTagTemplateIamBindingLister implements the CatalogTagTemplateIamBindingLister interface.
type catalogTagTemplateIamBindingLister struct {
	indexer cache.Indexer
}

// NewCatalogTagTemplateIamBindingLister returns a new CatalogTagTemplateIamBindingLister.
func NewCatalogTagTemplateIamBindingLister(indexer cache.Indexer) CatalogTagTemplateIamBindingLister {
	return &catalogTagTemplateIamBindingLister{indexer: indexer}
}

// List lists all CatalogTagTemplateIamBindings in the indexer.
func (s *catalogTagTemplateIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.CatalogTagTemplateIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CatalogTagTemplateIamBinding))
	})
	return ret, err
}

// CatalogTagTemplateIamBindings returns an object that can list and get CatalogTagTemplateIamBindings.
func (s *catalogTagTemplateIamBindingLister) CatalogTagTemplateIamBindings(namespace string) CatalogTagTemplateIamBindingNamespaceLister {
	return catalogTagTemplateIamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CatalogTagTemplateIamBindingNamespaceLister helps list and get CatalogTagTemplateIamBindings.
// All objects returned here must be treated as read-only.
type CatalogTagTemplateIamBindingNamespaceLister interface {
	// List lists all CatalogTagTemplateIamBindings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CatalogTagTemplateIamBinding, err error)
	// Get retrieves the CatalogTagTemplateIamBinding from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CatalogTagTemplateIamBinding, error)
	CatalogTagTemplateIamBindingNamespaceListerExpansion
}

// catalogTagTemplateIamBindingNamespaceLister implements the CatalogTagTemplateIamBindingNamespaceLister
// interface.
type catalogTagTemplateIamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CatalogTagTemplateIamBindings in the indexer for a given namespace.
func (s catalogTagTemplateIamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CatalogTagTemplateIamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CatalogTagTemplateIamBinding))
	})
	return ret, err
}

// Get retrieves the CatalogTagTemplateIamBinding from the indexer for a given namespace and name.
func (s catalogTagTemplateIamBindingNamespaceLister) Get(name string) (*v1alpha1.CatalogTagTemplateIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("catalogtagtemplateiambinding"), name)
	}
	return obj.(*v1alpha1.CatalogTagTemplateIamBinding), nil
}
