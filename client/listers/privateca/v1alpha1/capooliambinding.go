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
	v1alpha1 "kubeform.dev/provider-google-api/apis/privateca/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CaPoolIamBindingLister helps list CaPoolIamBindings.
// All objects returned here must be treated as read-only.
type CaPoolIamBindingLister interface {
	// List lists all CaPoolIamBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CaPoolIamBinding, err error)
	// CaPoolIamBindings returns an object that can list and get CaPoolIamBindings.
	CaPoolIamBindings(namespace string) CaPoolIamBindingNamespaceLister
	CaPoolIamBindingListerExpansion
}

// caPoolIamBindingLister implements the CaPoolIamBindingLister interface.
type caPoolIamBindingLister struct {
	indexer cache.Indexer
}

// NewCaPoolIamBindingLister returns a new CaPoolIamBindingLister.
func NewCaPoolIamBindingLister(indexer cache.Indexer) CaPoolIamBindingLister {
	return &caPoolIamBindingLister{indexer: indexer}
}

// List lists all CaPoolIamBindings in the indexer.
func (s *caPoolIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.CaPoolIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CaPoolIamBinding))
	})
	return ret, err
}

// CaPoolIamBindings returns an object that can list and get CaPoolIamBindings.
func (s *caPoolIamBindingLister) CaPoolIamBindings(namespace string) CaPoolIamBindingNamespaceLister {
	return caPoolIamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CaPoolIamBindingNamespaceLister helps list and get CaPoolIamBindings.
// All objects returned here must be treated as read-only.
type CaPoolIamBindingNamespaceLister interface {
	// List lists all CaPoolIamBindings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CaPoolIamBinding, err error)
	// Get retrieves the CaPoolIamBinding from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CaPoolIamBinding, error)
	CaPoolIamBindingNamespaceListerExpansion
}

// caPoolIamBindingNamespaceLister implements the CaPoolIamBindingNamespaceLister
// interface.
type caPoolIamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CaPoolIamBindings in the indexer for a given namespace.
func (s caPoolIamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CaPoolIamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CaPoolIamBinding))
	})
	return ret, err
}

// Get retrieves the CaPoolIamBinding from the indexer for a given namespace and name.
func (s caPoolIamBindingNamespaceLister) Get(name string) (*v1alpha1.CaPoolIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("capooliambinding"), name)
	}
	return obj.(*v1alpha1.CaPoolIamBinding), nil
}
