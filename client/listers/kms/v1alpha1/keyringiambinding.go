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
	v1alpha1 "kubeform.dev/provider-google-api/apis/kms/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KeyRingIamBindingLister helps list KeyRingIamBindings.
// All objects returned here must be treated as read-only.
type KeyRingIamBindingLister interface {
	// List lists all KeyRingIamBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.KeyRingIamBinding, err error)
	// KeyRingIamBindings returns an object that can list and get KeyRingIamBindings.
	KeyRingIamBindings(namespace string) KeyRingIamBindingNamespaceLister
	KeyRingIamBindingListerExpansion
}

// keyRingIamBindingLister implements the KeyRingIamBindingLister interface.
type keyRingIamBindingLister struct {
	indexer cache.Indexer
}

// NewKeyRingIamBindingLister returns a new KeyRingIamBindingLister.
func NewKeyRingIamBindingLister(indexer cache.Indexer) KeyRingIamBindingLister {
	return &keyRingIamBindingLister{indexer: indexer}
}

// List lists all KeyRingIamBindings in the indexer.
func (s *keyRingIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.KeyRingIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KeyRingIamBinding))
	})
	return ret, err
}

// KeyRingIamBindings returns an object that can list and get KeyRingIamBindings.
func (s *keyRingIamBindingLister) KeyRingIamBindings(namespace string) KeyRingIamBindingNamespaceLister {
	return keyRingIamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KeyRingIamBindingNamespaceLister helps list and get KeyRingIamBindings.
// All objects returned here must be treated as read-only.
type KeyRingIamBindingNamespaceLister interface {
	// List lists all KeyRingIamBindings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.KeyRingIamBinding, err error)
	// Get retrieves the KeyRingIamBinding from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.KeyRingIamBinding, error)
	KeyRingIamBindingNamespaceListerExpansion
}

// keyRingIamBindingNamespaceLister implements the KeyRingIamBindingNamespaceLister
// interface.
type keyRingIamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KeyRingIamBindings in the indexer for a given namespace.
func (s keyRingIamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.KeyRingIamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KeyRingIamBinding))
	})
	return ret, err
}

// Get retrieves the KeyRingIamBinding from the indexer for a given namespace and name.
func (s keyRingIamBindingNamespaceLister) Get(name string) (*v1alpha1.KeyRingIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("keyringiambinding"), name)
	}
	return obj.(*v1alpha1.KeyRingIamBinding), nil
}
