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
	v1alpha1 "kubeform.dev/provider-google-api/apis/recaptcha/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EnterpriseKeyLister helps list EnterpriseKeys.
// All objects returned here must be treated as read-only.
type EnterpriseKeyLister interface {
	// List lists all EnterpriseKeys in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EnterpriseKey, err error)
	// EnterpriseKeys returns an object that can list and get EnterpriseKeys.
	EnterpriseKeys(namespace string) EnterpriseKeyNamespaceLister
	EnterpriseKeyListerExpansion
}

// enterpriseKeyLister implements the EnterpriseKeyLister interface.
type enterpriseKeyLister struct {
	indexer cache.Indexer
}

// NewEnterpriseKeyLister returns a new EnterpriseKeyLister.
func NewEnterpriseKeyLister(indexer cache.Indexer) EnterpriseKeyLister {
	return &enterpriseKeyLister{indexer: indexer}
}

// List lists all EnterpriseKeys in the indexer.
func (s *enterpriseKeyLister) List(selector labels.Selector) (ret []*v1alpha1.EnterpriseKey, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EnterpriseKey))
	})
	return ret, err
}

// EnterpriseKeys returns an object that can list and get EnterpriseKeys.
func (s *enterpriseKeyLister) EnterpriseKeys(namespace string) EnterpriseKeyNamespaceLister {
	return enterpriseKeyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EnterpriseKeyNamespaceLister helps list and get EnterpriseKeys.
// All objects returned here must be treated as read-only.
type EnterpriseKeyNamespaceLister interface {
	// List lists all EnterpriseKeys in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EnterpriseKey, err error)
	// Get retrieves the EnterpriseKey from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EnterpriseKey, error)
	EnterpriseKeyNamespaceListerExpansion
}

// enterpriseKeyNamespaceLister implements the EnterpriseKeyNamespaceLister
// interface.
type enterpriseKeyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EnterpriseKeys in the indexer for a given namespace.
func (s enterpriseKeyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EnterpriseKey, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EnterpriseKey))
	})
	return ret, err
}

// Get retrieves the EnterpriseKey from the indexer for a given namespace and name.
func (s enterpriseKeyNamespaceLister) Get(name string) (*v1alpha1.EnterpriseKey, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("enterprisekey"), name)
	}
	return obj.(*v1alpha1.EnterpriseKey), nil
}
