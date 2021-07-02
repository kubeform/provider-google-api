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
	v1alpha1 "kubeform.dev/provider-google-api/apis/storage/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HmacKeyLister helps list HmacKeys.
// All objects returned here must be treated as read-only.
type HmacKeyLister interface {
	// List lists all HmacKeys in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HmacKey, err error)
	// HmacKeys returns an object that can list and get HmacKeys.
	HmacKeys(namespace string) HmacKeyNamespaceLister
	HmacKeyListerExpansion
}

// hmacKeyLister implements the HmacKeyLister interface.
type hmacKeyLister struct {
	indexer cache.Indexer
}

// NewHmacKeyLister returns a new HmacKeyLister.
func NewHmacKeyLister(indexer cache.Indexer) HmacKeyLister {
	return &hmacKeyLister{indexer: indexer}
}

// List lists all HmacKeys in the indexer.
func (s *hmacKeyLister) List(selector labels.Selector) (ret []*v1alpha1.HmacKey, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HmacKey))
	})
	return ret, err
}

// HmacKeys returns an object that can list and get HmacKeys.
func (s *hmacKeyLister) HmacKeys(namespace string) HmacKeyNamespaceLister {
	return hmacKeyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HmacKeyNamespaceLister helps list and get HmacKeys.
// All objects returned here must be treated as read-only.
type HmacKeyNamespaceLister interface {
	// List lists all HmacKeys in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HmacKey, err error)
	// Get retrieves the HmacKey from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.HmacKey, error)
	HmacKeyNamespaceListerExpansion
}

// hmacKeyNamespaceLister implements the HmacKeyNamespaceLister
// interface.
type hmacKeyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HmacKeys in the indexer for a given namespace.
func (s hmacKeyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HmacKey, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HmacKey))
	})
	return ret, err
}

// Get retrieves the HmacKey from the indexer for a given namespace and name.
func (s hmacKeyNamespaceLister) Get(name string) (*v1alpha1.HmacKey, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("hmackey"), name)
	}
	return obj.(*v1alpha1.HmacKey), nil
}
