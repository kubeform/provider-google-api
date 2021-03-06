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
	v1alpha1 "kubeform.dev/provider-google-api/apis/compute/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RegionURLMapLister helps list RegionURLMaps.
// All objects returned here must be treated as read-only.
type RegionURLMapLister interface {
	// List lists all RegionURLMaps in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RegionURLMap, err error)
	// RegionURLMaps returns an object that can list and get RegionURLMaps.
	RegionURLMaps(namespace string) RegionURLMapNamespaceLister
	RegionURLMapListerExpansion
}

// regionURLMapLister implements the RegionURLMapLister interface.
type regionURLMapLister struct {
	indexer cache.Indexer
}

// NewRegionURLMapLister returns a new RegionURLMapLister.
func NewRegionURLMapLister(indexer cache.Indexer) RegionURLMapLister {
	return &regionURLMapLister{indexer: indexer}
}

// List lists all RegionURLMaps in the indexer.
func (s *regionURLMapLister) List(selector labels.Selector) (ret []*v1alpha1.RegionURLMap, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RegionURLMap))
	})
	return ret, err
}

// RegionURLMaps returns an object that can list and get RegionURLMaps.
func (s *regionURLMapLister) RegionURLMaps(namespace string) RegionURLMapNamespaceLister {
	return regionURLMapNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RegionURLMapNamespaceLister helps list and get RegionURLMaps.
// All objects returned here must be treated as read-only.
type RegionURLMapNamespaceLister interface {
	// List lists all RegionURLMaps in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RegionURLMap, err error)
	// Get retrieves the RegionURLMap from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.RegionURLMap, error)
	RegionURLMapNamespaceListerExpansion
}

// regionURLMapNamespaceLister implements the RegionURLMapNamespaceLister
// interface.
type regionURLMapNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RegionURLMaps in the indexer for a given namespace.
func (s regionURLMapNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RegionURLMap, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RegionURLMap))
	})
	return ret, err
}

// Get retrieves the RegionURLMap from the indexer for a given namespace and name.
func (s regionURLMapNamespaceLister) Get(name string) (*v1alpha1.RegionURLMap, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("regionurlmap"), name)
	}
	return obj.(*v1alpha1.RegionURLMap), nil
}
