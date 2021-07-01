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
	v1alpha1 "kubeform.dev/provider-google-api/apis/kms/v1alpha1"
)

// CryptoKeyLister helps list CryptoKeys.
// All objects returned here must be treated as read-only.
type CryptoKeyLister interface {
	// List lists all CryptoKeys in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CryptoKey, err error)
	// CryptoKeys returns an object that can list and get CryptoKeys.
	CryptoKeys(namespace string) CryptoKeyNamespaceLister
	CryptoKeyListerExpansion
}

// cryptoKeyLister implements the CryptoKeyLister interface.
type cryptoKeyLister struct {
	indexer cache.Indexer
}

// NewCryptoKeyLister returns a new CryptoKeyLister.
func NewCryptoKeyLister(indexer cache.Indexer) CryptoKeyLister {
	return &cryptoKeyLister{indexer: indexer}
}

// List lists all CryptoKeys in the indexer.
func (s *cryptoKeyLister) List(selector labels.Selector) (ret []*v1alpha1.CryptoKey, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CryptoKey))
	})
	return ret, err
}

// CryptoKeys returns an object that can list and get CryptoKeys.
func (s *cryptoKeyLister) CryptoKeys(namespace string) CryptoKeyNamespaceLister {
	return cryptoKeyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CryptoKeyNamespaceLister helps list and get CryptoKeys.
// All objects returned here must be treated as read-only.
type CryptoKeyNamespaceLister interface {
	// List lists all CryptoKeys in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CryptoKey, err error)
	// Get retrieves the CryptoKey from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CryptoKey, error)
	CryptoKeyNamespaceListerExpansion
}

// cryptoKeyNamespaceLister implements the CryptoKeyNamespaceLister
// interface.
type cryptoKeyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CryptoKeys in the indexer for a given namespace.
func (s cryptoKeyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CryptoKey, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CryptoKey))
	})
	return ret, err
}

// Get retrieves the CryptoKey from the indexer for a given namespace and name.
func (s cryptoKeyNamespaceLister) Get(name string) (*v1alpha1.CryptoKey, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cryptokey"), name)
	}
	return obj.(*v1alpha1.CryptoKey), nil
}
