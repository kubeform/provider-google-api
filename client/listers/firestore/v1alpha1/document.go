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
	v1alpha1 "kubeform.dev/provider-google-api/apis/firestore/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DocumentLister helps list Documents.
// All objects returned here must be treated as read-only.
type DocumentLister interface {
	// List lists all Documents in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Document, err error)
	// Documents returns an object that can list and get Documents.
	Documents(namespace string) DocumentNamespaceLister
	DocumentListerExpansion
}

// documentLister implements the DocumentLister interface.
type documentLister struct {
	indexer cache.Indexer
}

// NewDocumentLister returns a new DocumentLister.
func NewDocumentLister(indexer cache.Indexer) DocumentLister {
	return &documentLister{indexer: indexer}
}

// List lists all Documents in the indexer.
func (s *documentLister) List(selector labels.Selector) (ret []*v1alpha1.Document, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Document))
	})
	return ret, err
}

// Documents returns an object that can list and get Documents.
func (s *documentLister) Documents(namespace string) DocumentNamespaceLister {
	return documentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DocumentNamespaceLister helps list and get Documents.
// All objects returned here must be treated as read-only.
type DocumentNamespaceLister interface {
	// List lists all Documents in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Document, err error)
	// Get retrieves the Document from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Document, error)
	DocumentNamespaceListerExpansion
}

// documentNamespaceLister implements the DocumentNamespaceLister
// interface.
type documentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Documents in the indexer for a given namespace.
func (s documentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Document, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Document))
	})
	return ret, err
}

// Get retrieves the Document from the indexer for a given namespace and name.
func (s documentNamespaceLister) Get(name string) (*v1alpha1.Document, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("document"), name)
	}
	return obj.(*v1alpha1.Document), nil
}
