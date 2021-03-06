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

// TagKeyLister helps list TagKeys.
// All objects returned here must be treated as read-only.
type TagKeyLister interface {
	// List lists all TagKeys in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TagKey, err error)
	// TagKeys returns an object that can list and get TagKeys.
	TagKeys(namespace string) TagKeyNamespaceLister
	TagKeyListerExpansion
}

// tagKeyLister implements the TagKeyLister interface.
type tagKeyLister struct {
	indexer cache.Indexer
}

// NewTagKeyLister returns a new TagKeyLister.
func NewTagKeyLister(indexer cache.Indexer) TagKeyLister {
	return &tagKeyLister{indexer: indexer}
}

// List lists all TagKeys in the indexer.
func (s *tagKeyLister) List(selector labels.Selector) (ret []*v1alpha1.TagKey, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TagKey))
	})
	return ret, err
}

// TagKeys returns an object that can list and get TagKeys.
func (s *tagKeyLister) TagKeys(namespace string) TagKeyNamespaceLister {
	return tagKeyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TagKeyNamespaceLister helps list and get TagKeys.
// All objects returned here must be treated as read-only.
type TagKeyNamespaceLister interface {
	// List lists all TagKeys in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TagKey, err error)
	// Get retrieves the TagKey from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.TagKey, error)
	TagKeyNamespaceListerExpansion
}

// tagKeyNamespaceLister implements the TagKeyNamespaceLister
// interface.
type tagKeyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TagKeys in the indexer for a given namespace.
func (s tagKeyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TagKey, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TagKey))
	})
	return ret, err
}

// Get retrieves the TagKey from the indexer for a given namespace and name.
func (s tagKeyNamespaceLister) Get(name string) (*v1alpha1.TagKey, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("tagkey"), name)
	}
	return obj.(*v1alpha1.TagKey), nil
}
