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
	v1alpha1 "kubeform.dev/provider-google-api/apis/container/v1alpha1"
)

// RegistryLister helps list Registries.
// All objects returned here must be treated as read-only.
type RegistryLister interface {
	// List lists all Registries in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Registry, err error)
	// Registries returns an object that can list and get Registries.
	Registries(namespace string) RegistryNamespaceLister
	RegistryListerExpansion
}

// registryLister implements the RegistryLister interface.
type registryLister struct {
	indexer cache.Indexer
}

// NewRegistryLister returns a new RegistryLister.
func NewRegistryLister(indexer cache.Indexer) RegistryLister {
	return &registryLister{indexer: indexer}
}

// List lists all Registries in the indexer.
func (s *registryLister) List(selector labels.Selector) (ret []*v1alpha1.Registry, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Registry))
	})
	return ret, err
}

// Registries returns an object that can list and get Registries.
func (s *registryLister) Registries(namespace string) RegistryNamespaceLister {
	return registryNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RegistryNamespaceLister helps list and get Registries.
// All objects returned here must be treated as read-only.
type RegistryNamespaceLister interface {
	// List lists all Registries in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Registry, err error)
	// Get retrieves the Registry from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Registry, error)
	RegistryNamespaceListerExpansion
}

// registryNamespaceLister implements the RegistryNamespaceLister
// interface.
type registryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Registries in the indexer for a given namespace.
func (s registryNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Registry, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Registry))
	})
	return ret, err
}

// Get retrieves the Registry from the indexer for a given namespace and name.
func (s registryNamespaceLister) Get(name string) (*v1alpha1.Registry, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("registry"), name)
	}
	return obj.(*v1alpha1.Registry), nil
}
