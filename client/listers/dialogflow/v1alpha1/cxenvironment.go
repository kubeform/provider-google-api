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
	v1alpha1 "kubeform.dev/provider-google-api/apis/dialogflow/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CxEnvironmentLister helps list CxEnvironments.
// All objects returned here must be treated as read-only.
type CxEnvironmentLister interface {
	// List lists all CxEnvironments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CxEnvironment, err error)
	// CxEnvironments returns an object that can list and get CxEnvironments.
	CxEnvironments(namespace string) CxEnvironmentNamespaceLister
	CxEnvironmentListerExpansion
}

// cxEnvironmentLister implements the CxEnvironmentLister interface.
type cxEnvironmentLister struct {
	indexer cache.Indexer
}

// NewCxEnvironmentLister returns a new CxEnvironmentLister.
func NewCxEnvironmentLister(indexer cache.Indexer) CxEnvironmentLister {
	return &cxEnvironmentLister{indexer: indexer}
}

// List lists all CxEnvironments in the indexer.
func (s *cxEnvironmentLister) List(selector labels.Selector) (ret []*v1alpha1.CxEnvironment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CxEnvironment))
	})
	return ret, err
}

// CxEnvironments returns an object that can list and get CxEnvironments.
func (s *cxEnvironmentLister) CxEnvironments(namespace string) CxEnvironmentNamespaceLister {
	return cxEnvironmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CxEnvironmentNamespaceLister helps list and get CxEnvironments.
// All objects returned here must be treated as read-only.
type CxEnvironmentNamespaceLister interface {
	// List lists all CxEnvironments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CxEnvironment, err error)
	// Get retrieves the CxEnvironment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CxEnvironment, error)
	CxEnvironmentNamespaceListerExpansion
}

// cxEnvironmentNamespaceLister implements the CxEnvironmentNamespaceLister
// interface.
type cxEnvironmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CxEnvironments in the indexer for a given namespace.
func (s cxEnvironmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CxEnvironment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CxEnvironment))
	})
	return ret, err
}

// Get retrieves the CxEnvironment from the indexer for a given namespace and name.
func (s cxEnvironmentNamespaceLister) Get(name string) (*v1alpha1.CxEnvironment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cxenvironment"), name)
	}
	return obj.(*v1alpha1.CxEnvironment), nil
}