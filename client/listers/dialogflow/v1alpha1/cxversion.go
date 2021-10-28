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

// CxVersionLister helps list CxVersions.
// All objects returned here must be treated as read-only.
type CxVersionLister interface {
	// List lists all CxVersions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CxVersion, err error)
	// CxVersions returns an object that can list and get CxVersions.
	CxVersions(namespace string) CxVersionNamespaceLister
	CxVersionListerExpansion
}

// cxVersionLister implements the CxVersionLister interface.
type cxVersionLister struct {
	indexer cache.Indexer
}

// NewCxVersionLister returns a new CxVersionLister.
func NewCxVersionLister(indexer cache.Indexer) CxVersionLister {
	return &cxVersionLister{indexer: indexer}
}

// List lists all CxVersions in the indexer.
func (s *cxVersionLister) List(selector labels.Selector) (ret []*v1alpha1.CxVersion, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CxVersion))
	})
	return ret, err
}

// CxVersions returns an object that can list and get CxVersions.
func (s *cxVersionLister) CxVersions(namespace string) CxVersionNamespaceLister {
	return cxVersionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CxVersionNamespaceLister helps list and get CxVersions.
// All objects returned here must be treated as read-only.
type CxVersionNamespaceLister interface {
	// List lists all CxVersions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CxVersion, err error)
	// Get retrieves the CxVersion from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CxVersion, error)
	CxVersionNamespaceListerExpansion
}

// cxVersionNamespaceLister implements the CxVersionNamespaceLister
// interface.
type cxVersionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CxVersions in the indexer for a given namespace.
func (s cxVersionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CxVersion, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CxVersion))
	})
	return ret, err
}

// Get retrieves the CxVersion from the indexer for a given namespace and name.
func (s cxVersionNamespaceLister) Get(name string) (*v1alpha1.CxVersion, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cxversion"), name)
	}
	return obj.(*v1alpha1.CxVersion), nil
}
