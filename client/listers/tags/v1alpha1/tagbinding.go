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
	v1alpha1 "kubeform.dev/provider-google-api/apis/tags/v1alpha1"
)

// TagBindingLister helps list TagBindings.
// All objects returned here must be treated as read-only.
type TagBindingLister interface {
	// List lists all TagBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TagBinding, err error)
	// TagBindings returns an object that can list and get TagBindings.
	TagBindings(namespace string) TagBindingNamespaceLister
	TagBindingListerExpansion
}

// tagBindingLister implements the TagBindingLister interface.
type tagBindingLister struct {
	indexer cache.Indexer
}

// NewTagBindingLister returns a new TagBindingLister.
func NewTagBindingLister(indexer cache.Indexer) TagBindingLister {
	return &tagBindingLister{indexer: indexer}
}

// List lists all TagBindings in the indexer.
func (s *tagBindingLister) List(selector labels.Selector) (ret []*v1alpha1.TagBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TagBinding))
	})
	return ret, err
}

// TagBindings returns an object that can list and get TagBindings.
func (s *tagBindingLister) TagBindings(namespace string) TagBindingNamespaceLister {
	return tagBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TagBindingNamespaceLister helps list and get TagBindings.
// All objects returned here must be treated as read-only.
type TagBindingNamespaceLister interface {
	// List lists all TagBindings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TagBinding, err error)
	// Get retrieves the TagBinding from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.TagBinding, error)
	TagBindingNamespaceListerExpansion
}

// tagBindingNamespaceLister implements the TagBindingNamespaceLister
// interface.
type tagBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TagBindings in the indexer for a given namespace.
func (s tagBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TagBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TagBinding))
	})
	return ret, err
}

// Get retrieves the TagBinding from the indexer for a given namespace and name.
func (s tagBindingNamespaceLister) Get(name string) (*v1alpha1.TagBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("tagbinding"), name)
	}
	return obj.(*v1alpha1.TagBinding), nil
}
