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
	v1alpha1 "kubeform.dev/provider-google-api/apis/iap/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WebIamBindingLister helps list WebIamBindings.
// All objects returned here must be treated as read-only.
type WebIamBindingLister interface {
	// List lists all WebIamBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WebIamBinding, err error)
	// WebIamBindings returns an object that can list and get WebIamBindings.
	WebIamBindings(namespace string) WebIamBindingNamespaceLister
	WebIamBindingListerExpansion
}

// webIamBindingLister implements the WebIamBindingLister interface.
type webIamBindingLister struct {
	indexer cache.Indexer
}

// NewWebIamBindingLister returns a new WebIamBindingLister.
func NewWebIamBindingLister(indexer cache.Indexer) WebIamBindingLister {
	return &webIamBindingLister{indexer: indexer}
}

// List lists all WebIamBindings in the indexer.
func (s *webIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.WebIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WebIamBinding))
	})
	return ret, err
}

// WebIamBindings returns an object that can list and get WebIamBindings.
func (s *webIamBindingLister) WebIamBindings(namespace string) WebIamBindingNamespaceLister {
	return webIamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WebIamBindingNamespaceLister helps list and get WebIamBindings.
// All objects returned here must be treated as read-only.
type WebIamBindingNamespaceLister interface {
	// List lists all WebIamBindings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WebIamBinding, err error)
	// Get retrieves the WebIamBinding from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.WebIamBinding, error)
	WebIamBindingNamespaceListerExpansion
}

// webIamBindingNamespaceLister implements the WebIamBindingNamespaceLister
// interface.
type webIamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WebIamBindings in the indexer for a given namespace.
func (s webIamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WebIamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WebIamBinding))
	})
	return ret, err
}

// Get retrieves the WebIamBinding from the indexer for a given namespace and name.
func (s webIamBindingNamespaceLister) Get(name string) (*v1alpha1.WebIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("webiambinding"), name)
	}
	return obj.(*v1alpha1.WebIamBinding), nil
}
