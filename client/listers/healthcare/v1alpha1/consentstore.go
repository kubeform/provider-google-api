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

// ConsentStoreLister helps list ConsentStores.
// All objects returned here must be treated as read-only.
type ConsentStoreLister interface {
	// List lists all ConsentStores in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ConsentStore, err error)
	// ConsentStores returns an object that can list and get ConsentStores.
	ConsentStores(namespace string) ConsentStoreNamespaceLister
	ConsentStoreListerExpansion
}

// consentStoreLister implements the ConsentStoreLister interface.
type consentStoreLister struct {
	indexer cache.Indexer
}

// NewConsentStoreLister returns a new ConsentStoreLister.
func NewConsentStoreLister(indexer cache.Indexer) ConsentStoreLister {
	return &consentStoreLister{indexer: indexer}
}

// List lists all ConsentStores in the indexer.
func (s *consentStoreLister) List(selector labels.Selector) (ret []*v1alpha1.ConsentStore, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConsentStore))
	})
	return ret, err
}

// ConsentStores returns an object that can list and get ConsentStores.
func (s *consentStoreLister) ConsentStores(namespace string) ConsentStoreNamespaceLister {
	return consentStoreNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ConsentStoreNamespaceLister helps list and get ConsentStores.
// All objects returned here must be treated as read-only.
type ConsentStoreNamespaceLister interface {
	// List lists all ConsentStores in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ConsentStore, err error)
	// Get retrieves the ConsentStore from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ConsentStore, error)
	ConsentStoreNamespaceListerExpansion
}

// consentStoreNamespaceLister implements the ConsentStoreNamespaceLister
// interface.
type consentStoreNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ConsentStores in the indexer for a given namespace.
func (s consentStoreNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ConsentStore, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConsentStore))
	})
	return ret, err
}

// Get retrieves the ConsentStore from the indexer for a given namespace and name.
func (s consentStoreNamespaceLister) Get(name string) (*v1alpha1.ConsentStore, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("consentstore"), name)
	}
	return obj.(*v1alpha1.ConsentStore), nil
}
