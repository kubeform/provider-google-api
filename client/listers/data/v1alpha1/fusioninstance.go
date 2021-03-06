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
	v1alpha1 "kubeform.dev/provider-google-api/apis/data/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FusionInstanceLister helps list FusionInstances.
// All objects returned here must be treated as read-only.
type FusionInstanceLister interface {
	// List lists all FusionInstances in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FusionInstance, err error)
	// FusionInstances returns an object that can list and get FusionInstances.
	FusionInstances(namespace string) FusionInstanceNamespaceLister
	FusionInstanceListerExpansion
}

// fusionInstanceLister implements the FusionInstanceLister interface.
type fusionInstanceLister struct {
	indexer cache.Indexer
}

// NewFusionInstanceLister returns a new FusionInstanceLister.
func NewFusionInstanceLister(indexer cache.Indexer) FusionInstanceLister {
	return &fusionInstanceLister{indexer: indexer}
}

// List lists all FusionInstances in the indexer.
func (s *fusionInstanceLister) List(selector labels.Selector) (ret []*v1alpha1.FusionInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FusionInstance))
	})
	return ret, err
}

// FusionInstances returns an object that can list and get FusionInstances.
func (s *fusionInstanceLister) FusionInstances(namespace string) FusionInstanceNamespaceLister {
	return fusionInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FusionInstanceNamespaceLister helps list and get FusionInstances.
// All objects returned here must be treated as read-only.
type FusionInstanceNamespaceLister interface {
	// List lists all FusionInstances in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FusionInstance, err error)
	// Get retrieves the FusionInstance from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FusionInstance, error)
	FusionInstanceNamespaceListerExpansion
}

// fusionInstanceNamespaceLister implements the FusionInstanceNamespaceLister
// interface.
type fusionInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FusionInstances in the indexer for a given namespace.
func (s fusionInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FusionInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FusionInstance))
	})
	return ret, err
}

// Get retrieves the FusionInstance from the indexer for a given namespace and name.
func (s fusionInstanceNamespaceLister) Get(name string) (*v1alpha1.FusionInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("fusioninstance"), name)
	}
	return obj.(*v1alpha1.FusionInstance), nil
}
