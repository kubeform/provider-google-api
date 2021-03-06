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
	v1alpha1 "kubeform.dev/provider-google-api/apis/network/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServicesEdgeCacheOriginLister helps list ServicesEdgeCacheOrigins.
// All objects returned here must be treated as read-only.
type ServicesEdgeCacheOriginLister interface {
	// List lists all ServicesEdgeCacheOrigins in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServicesEdgeCacheOrigin, err error)
	// ServicesEdgeCacheOrigins returns an object that can list and get ServicesEdgeCacheOrigins.
	ServicesEdgeCacheOrigins(namespace string) ServicesEdgeCacheOriginNamespaceLister
	ServicesEdgeCacheOriginListerExpansion
}

// servicesEdgeCacheOriginLister implements the ServicesEdgeCacheOriginLister interface.
type servicesEdgeCacheOriginLister struct {
	indexer cache.Indexer
}

// NewServicesEdgeCacheOriginLister returns a new ServicesEdgeCacheOriginLister.
func NewServicesEdgeCacheOriginLister(indexer cache.Indexer) ServicesEdgeCacheOriginLister {
	return &servicesEdgeCacheOriginLister{indexer: indexer}
}

// List lists all ServicesEdgeCacheOrigins in the indexer.
func (s *servicesEdgeCacheOriginLister) List(selector labels.Selector) (ret []*v1alpha1.ServicesEdgeCacheOrigin, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServicesEdgeCacheOrigin))
	})
	return ret, err
}

// ServicesEdgeCacheOrigins returns an object that can list and get ServicesEdgeCacheOrigins.
func (s *servicesEdgeCacheOriginLister) ServicesEdgeCacheOrigins(namespace string) ServicesEdgeCacheOriginNamespaceLister {
	return servicesEdgeCacheOriginNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServicesEdgeCacheOriginNamespaceLister helps list and get ServicesEdgeCacheOrigins.
// All objects returned here must be treated as read-only.
type ServicesEdgeCacheOriginNamespaceLister interface {
	// List lists all ServicesEdgeCacheOrigins in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServicesEdgeCacheOrigin, err error)
	// Get retrieves the ServicesEdgeCacheOrigin from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ServicesEdgeCacheOrigin, error)
	ServicesEdgeCacheOriginNamespaceListerExpansion
}

// servicesEdgeCacheOriginNamespaceLister implements the ServicesEdgeCacheOriginNamespaceLister
// interface.
type servicesEdgeCacheOriginNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServicesEdgeCacheOrigins in the indexer for a given namespace.
func (s servicesEdgeCacheOriginNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServicesEdgeCacheOrigin, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServicesEdgeCacheOrigin))
	})
	return ret, err
}

// Get retrieves the ServicesEdgeCacheOrigin from the indexer for a given namespace and name.
func (s servicesEdgeCacheOriginNamespaceLister) Get(name string) (*v1alpha1.ServicesEdgeCacheOrigin, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("servicesedgecacheorigin"), name)
	}
	return obj.(*v1alpha1.ServicesEdgeCacheOrigin), nil
}
