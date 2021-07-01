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
	v1alpha1 "kubeform.dev/provider-google-api/apis/accesscontext/v1alpha1"
)

// ManagerServicePerimeterLister helps list ManagerServicePerimeters.
// All objects returned here must be treated as read-only.
type ManagerServicePerimeterLister interface {
	// List lists all ManagerServicePerimeters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagerServicePerimeter, err error)
	// ManagerServicePerimeters returns an object that can list and get ManagerServicePerimeters.
	ManagerServicePerimeters(namespace string) ManagerServicePerimeterNamespaceLister
	ManagerServicePerimeterListerExpansion
}

// managerServicePerimeterLister implements the ManagerServicePerimeterLister interface.
type managerServicePerimeterLister struct {
	indexer cache.Indexer
}

// NewManagerServicePerimeterLister returns a new ManagerServicePerimeterLister.
func NewManagerServicePerimeterLister(indexer cache.Indexer) ManagerServicePerimeterLister {
	return &managerServicePerimeterLister{indexer: indexer}
}

// List lists all ManagerServicePerimeters in the indexer.
func (s *managerServicePerimeterLister) List(selector labels.Selector) (ret []*v1alpha1.ManagerServicePerimeter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagerServicePerimeter))
	})
	return ret, err
}

// ManagerServicePerimeters returns an object that can list and get ManagerServicePerimeters.
func (s *managerServicePerimeterLister) ManagerServicePerimeters(namespace string) ManagerServicePerimeterNamespaceLister {
	return managerServicePerimeterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ManagerServicePerimeterNamespaceLister helps list and get ManagerServicePerimeters.
// All objects returned here must be treated as read-only.
type ManagerServicePerimeterNamespaceLister interface {
	// List lists all ManagerServicePerimeters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagerServicePerimeter, err error)
	// Get retrieves the ManagerServicePerimeter from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ManagerServicePerimeter, error)
	ManagerServicePerimeterNamespaceListerExpansion
}

// managerServicePerimeterNamespaceLister implements the ManagerServicePerimeterNamespaceLister
// interface.
type managerServicePerimeterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ManagerServicePerimeters in the indexer for a given namespace.
func (s managerServicePerimeterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ManagerServicePerimeter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagerServicePerimeter))
	})
	return ret, err
}

// Get retrieves the ManagerServicePerimeter from the indexer for a given namespace and name.
func (s managerServicePerimeterNamespaceLister) Get(name string) (*v1alpha1.ManagerServicePerimeter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("managerserviceperimeter"), name)
	}
	return obj.(*v1alpha1.ManagerServicePerimeter), nil
}
