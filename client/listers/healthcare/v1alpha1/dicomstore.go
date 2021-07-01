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
	v1alpha1 "kubeform.dev/provider-google-api/apis/healthcare/v1alpha1"
)

// DicomStoreLister helps list DicomStores.
// All objects returned here must be treated as read-only.
type DicomStoreLister interface {
	// List lists all DicomStores in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DicomStore, err error)
	// DicomStores returns an object that can list and get DicomStores.
	DicomStores(namespace string) DicomStoreNamespaceLister
	DicomStoreListerExpansion
}

// dicomStoreLister implements the DicomStoreLister interface.
type dicomStoreLister struct {
	indexer cache.Indexer
}

// NewDicomStoreLister returns a new DicomStoreLister.
func NewDicomStoreLister(indexer cache.Indexer) DicomStoreLister {
	return &dicomStoreLister{indexer: indexer}
}

// List lists all DicomStores in the indexer.
func (s *dicomStoreLister) List(selector labels.Selector) (ret []*v1alpha1.DicomStore, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DicomStore))
	})
	return ret, err
}

// DicomStores returns an object that can list and get DicomStores.
func (s *dicomStoreLister) DicomStores(namespace string) DicomStoreNamespaceLister {
	return dicomStoreNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DicomStoreNamespaceLister helps list and get DicomStores.
// All objects returned here must be treated as read-only.
type DicomStoreNamespaceLister interface {
	// List lists all DicomStores in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DicomStore, err error)
	// Get retrieves the DicomStore from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DicomStore, error)
	DicomStoreNamespaceListerExpansion
}

// dicomStoreNamespaceLister implements the DicomStoreNamespaceLister
// interface.
type dicomStoreNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DicomStores in the indexer for a given namespace.
func (s dicomStoreNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DicomStore, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DicomStore))
	})
	return ret, err
}

// Get retrieves the DicomStore from the indexer for a given namespace and name.
func (s dicomStoreNamespaceLister) Get(name string) (*v1alpha1.DicomStore, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dicomstore"), name)
	}
	return obj.(*v1alpha1.DicomStore), nil
}
