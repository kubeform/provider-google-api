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
	v1alpha1 "kubeform.dev/provider-google-api/apis/tags/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TagValueLister helps list TagValues.
// All objects returned here must be treated as read-only.
type TagValueLister interface {
	// List lists all TagValues in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TagValue, err error)
	// TagValues returns an object that can list and get TagValues.
	TagValues(namespace string) TagValueNamespaceLister
	TagValueListerExpansion
}

// tagValueLister implements the TagValueLister interface.
type tagValueLister struct {
	indexer cache.Indexer
}

// NewTagValueLister returns a new TagValueLister.
func NewTagValueLister(indexer cache.Indexer) TagValueLister {
	return &tagValueLister{indexer: indexer}
}

// List lists all TagValues in the indexer.
func (s *tagValueLister) List(selector labels.Selector) (ret []*v1alpha1.TagValue, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TagValue))
	})
	return ret, err
}

// TagValues returns an object that can list and get TagValues.
func (s *tagValueLister) TagValues(namespace string) TagValueNamespaceLister {
	return tagValueNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TagValueNamespaceLister helps list and get TagValues.
// All objects returned here must be treated as read-only.
type TagValueNamespaceLister interface {
	// List lists all TagValues in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TagValue, err error)
	// Get retrieves the TagValue from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.TagValue, error)
	TagValueNamespaceListerExpansion
}

// tagValueNamespaceLister implements the TagValueNamespaceLister
// interface.
type tagValueNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TagValues in the indexer for a given namespace.
func (s tagValueNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TagValue, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TagValue))
	})
	return ret, err
}

// Get retrieves the TagValue from the indexer for a given namespace and name.
func (s tagValueNamespaceLister) Get(name string) (*v1alpha1.TagValue, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("tagvalue"), name)
	}
	return obj.(*v1alpha1.TagValue), nil
}
