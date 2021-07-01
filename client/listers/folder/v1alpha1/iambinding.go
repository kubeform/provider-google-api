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
	v1alpha1 "kubeform.dev/provider-google-api/apis/folder/v1alpha1"
)

// IamBindingLister helps list IamBindings.
// All objects returned here must be treated as read-only.
type IamBindingLister interface {
	// List lists all IamBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IamBinding, err error)
	// IamBindings returns an object that can list and get IamBindings.
	IamBindings(namespace string) IamBindingNamespaceLister
	IamBindingListerExpansion
}

// iamBindingLister implements the IamBindingLister interface.
type iamBindingLister struct {
	indexer cache.Indexer
}

// NewIamBindingLister returns a new IamBindingLister.
func NewIamBindingLister(indexer cache.Indexer) IamBindingLister {
	return &iamBindingLister{indexer: indexer}
}

// List lists all IamBindings in the indexer.
func (s *iamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.IamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamBinding))
	})
	return ret, err
}

// IamBindings returns an object that can list and get IamBindings.
func (s *iamBindingLister) IamBindings(namespace string) IamBindingNamespaceLister {
	return iamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IamBindingNamespaceLister helps list and get IamBindings.
// All objects returned here must be treated as read-only.
type IamBindingNamespaceLister interface {
	// List lists all IamBindings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IamBinding, err error)
	// Get retrieves the IamBinding from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.IamBinding, error)
	IamBindingNamespaceListerExpansion
}

// iamBindingNamespaceLister implements the IamBindingNamespaceLister
// interface.
type iamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IamBindings in the indexer for a given namespace.
func (s iamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamBinding))
	})
	return ret, err
}

// Get retrieves the IamBinding from the indexer for a given namespace and name.
func (s iamBindingNamespaceLister) Get(name string) (*v1alpha1.IamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iambinding"), name)
	}
	return obj.(*v1alpha1.IamBinding), nil
}
