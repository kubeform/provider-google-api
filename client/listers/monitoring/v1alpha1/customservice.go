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
	v1alpha1 "kubeform.dev/provider-google-api/apis/monitoring/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CustomServiceLister helps list CustomServices.
// All objects returned here must be treated as read-only.
type CustomServiceLister interface {
	// List lists all CustomServices in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CustomService, err error)
	// CustomServices returns an object that can list and get CustomServices.
	CustomServices(namespace string) CustomServiceNamespaceLister
	CustomServiceListerExpansion
}

// customServiceLister implements the CustomServiceLister interface.
type customServiceLister struct {
	indexer cache.Indexer
}

// NewCustomServiceLister returns a new CustomServiceLister.
func NewCustomServiceLister(indexer cache.Indexer) CustomServiceLister {
	return &customServiceLister{indexer: indexer}
}

// List lists all CustomServices in the indexer.
func (s *customServiceLister) List(selector labels.Selector) (ret []*v1alpha1.CustomService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CustomService))
	})
	return ret, err
}

// CustomServices returns an object that can list and get CustomServices.
func (s *customServiceLister) CustomServices(namespace string) CustomServiceNamespaceLister {
	return customServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CustomServiceNamespaceLister helps list and get CustomServices.
// All objects returned here must be treated as read-only.
type CustomServiceNamespaceLister interface {
	// List lists all CustomServices in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CustomService, err error)
	// Get retrieves the CustomService from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CustomService, error)
	CustomServiceNamespaceListerExpansion
}

// customServiceNamespaceLister implements the CustomServiceNamespaceLister
// interface.
type customServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CustomServices in the indexer for a given namespace.
func (s customServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CustomService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CustomService))
	})
	return ret, err
}

// Get retrieves the CustomService from the indexer for a given namespace and name.
func (s customServiceNamespaceLister) Get(name string) (*v1alpha1.CustomService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("customservice"), name)
	}
	return obj.(*v1alpha1.CustomService), nil
}
