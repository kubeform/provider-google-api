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
	v1alpha1 "kubeform.dev/provider-google-api/apis/secret/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ManagerSecretIamBindingLister helps list ManagerSecretIamBindings.
// All objects returned here must be treated as read-only.
type ManagerSecretIamBindingLister interface {
	// List lists all ManagerSecretIamBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagerSecretIamBinding, err error)
	// ManagerSecretIamBindings returns an object that can list and get ManagerSecretIamBindings.
	ManagerSecretIamBindings(namespace string) ManagerSecretIamBindingNamespaceLister
	ManagerSecretIamBindingListerExpansion
}

// managerSecretIamBindingLister implements the ManagerSecretIamBindingLister interface.
type managerSecretIamBindingLister struct {
	indexer cache.Indexer
}

// NewManagerSecretIamBindingLister returns a new ManagerSecretIamBindingLister.
func NewManagerSecretIamBindingLister(indexer cache.Indexer) ManagerSecretIamBindingLister {
	return &managerSecretIamBindingLister{indexer: indexer}
}

// List lists all ManagerSecretIamBindings in the indexer.
func (s *managerSecretIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.ManagerSecretIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagerSecretIamBinding))
	})
	return ret, err
}

// ManagerSecretIamBindings returns an object that can list and get ManagerSecretIamBindings.
func (s *managerSecretIamBindingLister) ManagerSecretIamBindings(namespace string) ManagerSecretIamBindingNamespaceLister {
	return managerSecretIamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ManagerSecretIamBindingNamespaceLister helps list and get ManagerSecretIamBindings.
// All objects returned here must be treated as read-only.
type ManagerSecretIamBindingNamespaceLister interface {
	// List lists all ManagerSecretIamBindings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagerSecretIamBinding, err error)
	// Get retrieves the ManagerSecretIamBinding from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ManagerSecretIamBinding, error)
	ManagerSecretIamBindingNamespaceListerExpansion
}

// managerSecretIamBindingNamespaceLister implements the ManagerSecretIamBindingNamespaceLister
// interface.
type managerSecretIamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ManagerSecretIamBindings in the indexer for a given namespace.
func (s managerSecretIamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ManagerSecretIamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagerSecretIamBinding))
	})
	return ret, err
}

// Get retrieves the ManagerSecretIamBinding from the indexer for a given namespace and name.
func (s managerSecretIamBindingNamespaceLister) Get(name string) (*v1alpha1.ManagerSecretIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("managersecretiambinding"), name)
	}
	return obj.(*v1alpha1.ManagerSecretIamBinding), nil
}
