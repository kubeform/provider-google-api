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
	v1alpha1 "kubeform.dev/provider-google-api/apis/iap/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AppEngineServiceIamBindingLister helps list AppEngineServiceIamBindings.
// All objects returned here must be treated as read-only.
type AppEngineServiceIamBindingLister interface {
	// List lists all AppEngineServiceIamBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AppEngineServiceIamBinding, err error)
	// AppEngineServiceIamBindings returns an object that can list and get AppEngineServiceIamBindings.
	AppEngineServiceIamBindings(namespace string) AppEngineServiceIamBindingNamespaceLister
	AppEngineServiceIamBindingListerExpansion
}

// appEngineServiceIamBindingLister implements the AppEngineServiceIamBindingLister interface.
type appEngineServiceIamBindingLister struct {
	indexer cache.Indexer
}

// NewAppEngineServiceIamBindingLister returns a new AppEngineServiceIamBindingLister.
func NewAppEngineServiceIamBindingLister(indexer cache.Indexer) AppEngineServiceIamBindingLister {
	return &appEngineServiceIamBindingLister{indexer: indexer}
}

// List lists all AppEngineServiceIamBindings in the indexer.
func (s *appEngineServiceIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.AppEngineServiceIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppEngineServiceIamBinding))
	})
	return ret, err
}

// AppEngineServiceIamBindings returns an object that can list and get AppEngineServiceIamBindings.
func (s *appEngineServiceIamBindingLister) AppEngineServiceIamBindings(namespace string) AppEngineServiceIamBindingNamespaceLister {
	return appEngineServiceIamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AppEngineServiceIamBindingNamespaceLister helps list and get AppEngineServiceIamBindings.
// All objects returned here must be treated as read-only.
type AppEngineServiceIamBindingNamespaceLister interface {
	// List lists all AppEngineServiceIamBindings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AppEngineServiceIamBinding, err error)
	// Get retrieves the AppEngineServiceIamBinding from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AppEngineServiceIamBinding, error)
	AppEngineServiceIamBindingNamespaceListerExpansion
}

// appEngineServiceIamBindingNamespaceLister implements the AppEngineServiceIamBindingNamespaceLister
// interface.
type appEngineServiceIamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AppEngineServiceIamBindings in the indexer for a given namespace.
func (s appEngineServiceIamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AppEngineServiceIamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppEngineServiceIamBinding))
	})
	return ret, err
}

// Get retrieves the AppEngineServiceIamBinding from the indexer for a given namespace and name.
func (s appEngineServiceIamBindingNamespaceLister) Get(name string) (*v1alpha1.AppEngineServiceIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("appengineserviceiambinding"), name)
	}
	return obj.(*v1alpha1.AppEngineServiceIamBinding), nil
}
