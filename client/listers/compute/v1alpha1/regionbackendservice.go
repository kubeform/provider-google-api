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
	v1alpha1 "kubeform.dev/provider-google-api/apis/compute/v1alpha1"
)

// RegionBackendServiceLister helps list RegionBackendServices.
// All objects returned here must be treated as read-only.
type RegionBackendServiceLister interface {
	// List lists all RegionBackendServices in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RegionBackendService, err error)
	// RegionBackendServices returns an object that can list and get RegionBackendServices.
	RegionBackendServices(namespace string) RegionBackendServiceNamespaceLister
	RegionBackendServiceListerExpansion
}

// regionBackendServiceLister implements the RegionBackendServiceLister interface.
type regionBackendServiceLister struct {
	indexer cache.Indexer
}

// NewRegionBackendServiceLister returns a new RegionBackendServiceLister.
func NewRegionBackendServiceLister(indexer cache.Indexer) RegionBackendServiceLister {
	return &regionBackendServiceLister{indexer: indexer}
}

// List lists all RegionBackendServices in the indexer.
func (s *regionBackendServiceLister) List(selector labels.Selector) (ret []*v1alpha1.RegionBackendService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RegionBackendService))
	})
	return ret, err
}

// RegionBackendServices returns an object that can list and get RegionBackendServices.
func (s *regionBackendServiceLister) RegionBackendServices(namespace string) RegionBackendServiceNamespaceLister {
	return regionBackendServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RegionBackendServiceNamespaceLister helps list and get RegionBackendServices.
// All objects returned here must be treated as read-only.
type RegionBackendServiceNamespaceLister interface {
	// List lists all RegionBackendServices in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RegionBackendService, err error)
	// Get retrieves the RegionBackendService from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.RegionBackendService, error)
	RegionBackendServiceNamespaceListerExpansion
}

// regionBackendServiceNamespaceLister implements the RegionBackendServiceNamespaceLister
// interface.
type regionBackendServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RegionBackendServices in the indexer for a given namespace.
func (s regionBackendServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RegionBackendService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RegionBackendService))
	})
	return ret, err
}

// Get retrieves the RegionBackendService from the indexer for a given namespace and name.
func (s regionBackendServiceNamespaceLister) Get(name string) (*v1alpha1.RegionBackendService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("regionbackendservice"), name)
	}
	return obj.(*v1alpha1.RegionBackendService), nil
}
