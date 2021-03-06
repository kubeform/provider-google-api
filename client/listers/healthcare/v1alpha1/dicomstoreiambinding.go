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
	v1alpha1 "kubeform.dev/provider-google-api/apis/healthcare/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DicomStoreIamBindingLister helps list DicomStoreIamBindings.
// All objects returned here must be treated as read-only.
type DicomStoreIamBindingLister interface {
	// List lists all DicomStoreIamBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DicomStoreIamBinding, err error)
	// DicomStoreIamBindings returns an object that can list and get DicomStoreIamBindings.
	DicomStoreIamBindings(namespace string) DicomStoreIamBindingNamespaceLister
	DicomStoreIamBindingListerExpansion
}

// dicomStoreIamBindingLister implements the DicomStoreIamBindingLister interface.
type dicomStoreIamBindingLister struct {
	indexer cache.Indexer
}

// NewDicomStoreIamBindingLister returns a new DicomStoreIamBindingLister.
func NewDicomStoreIamBindingLister(indexer cache.Indexer) DicomStoreIamBindingLister {
	return &dicomStoreIamBindingLister{indexer: indexer}
}

// List lists all DicomStoreIamBindings in the indexer.
func (s *dicomStoreIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.DicomStoreIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DicomStoreIamBinding))
	})
	return ret, err
}

// DicomStoreIamBindings returns an object that can list and get DicomStoreIamBindings.
func (s *dicomStoreIamBindingLister) DicomStoreIamBindings(namespace string) DicomStoreIamBindingNamespaceLister {
	return dicomStoreIamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DicomStoreIamBindingNamespaceLister helps list and get DicomStoreIamBindings.
// All objects returned here must be treated as read-only.
type DicomStoreIamBindingNamespaceLister interface {
	// List lists all DicomStoreIamBindings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DicomStoreIamBinding, err error)
	// Get retrieves the DicomStoreIamBinding from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DicomStoreIamBinding, error)
	DicomStoreIamBindingNamespaceListerExpansion
}

// dicomStoreIamBindingNamespaceLister implements the DicomStoreIamBindingNamespaceLister
// interface.
type dicomStoreIamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DicomStoreIamBindings in the indexer for a given namespace.
func (s dicomStoreIamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DicomStoreIamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DicomStoreIamBinding))
	})
	return ret, err
}

// Get retrieves the DicomStoreIamBinding from the indexer for a given namespace and name.
func (s dicomStoreIamBindingNamespaceLister) Get(name string) (*v1alpha1.DicomStoreIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dicomstoreiambinding"), name)
	}
	return obj.(*v1alpha1.DicomStoreIamBinding), nil
}
