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
	v1alpha1 "kubeform.dev/provider-google-api/apis/notebooks/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RuntimeIamBindingLister helps list RuntimeIamBindings.
// All objects returned here must be treated as read-only.
type RuntimeIamBindingLister interface {
	// List lists all RuntimeIamBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RuntimeIamBinding, err error)
	// RuntimeIamBindings returns an object that can list and get RuntimeIamBindings.
	RuntimeIamBindings(namespace string) RuntimeIamBindingNamespaceLister
	RuntimeIamBindingListerExpansion
}

// runtimeIamBindingLister implements the RuntimeIamBindingLister interface.
type runtimeIamBindingLister struct {
	indexer cache.Indexer
}

// NewRuntimeIamBindingLister returns a new RuntimeIamBindingLister.
func NewRuntimeIamBindingLister(indexer cache.Indexer) RuntimeIamBindingLister {
	return &runtimeIamBindingLister{indexer: indexer}
}

// List lists all RuntimeIamBindings in the indexer.
func (s *runtimeIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.RuntimeIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RuntimeIamBinding))
	})
	return ret, err
}

// RuntimeIamBindings returns an object that can list and get RuntimeIamBindings.
func (s *runtimeIamBindingLister) RuntimeIamBindings(namespace string) RuntimeIamBindingNamespaceLister {
	return runtimeIamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RuntimeIamBindingNamespaceLister helps list and get RuntimeIamBindings.
// All objects returned here must be treated as read-only.
type RuntimeIamBindingNamespaceLister interface {
	// List lists all RuntimeIamBindings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RuntimeIamBinding, err error)
	// Get retrieves the RuntimeIamBinding from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.RuntimeIamBinding, error)
	RuntimeIamBindingNamespaceListerExpansion
}

// runtimeIamBindingNamespaceLister implements the RuntimeIamBindingNamespaceLister
// interface.
type runtimeIamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RuntimeIamBindings in the indexer for a given namespace.
func (s runtimeIamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RuntimeIamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RuntimeIamBinding))
	})
	return ret, err
}

// Get retrieves the RuntimeIamBinding from the indexer for a given namespace and name.
func (s runtimeIamBindingNamespaceLister) Get(name string) (*v1alpha1.RuntimeIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("runtimeiambinding"), name)
	}
	return obj.(*v1alpha1.RuntimeIamBinding), nil
}