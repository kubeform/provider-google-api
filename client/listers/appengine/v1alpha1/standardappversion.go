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
	v1alpha1 "kubeform.dev/provider-google-api/apis/appengine/v1alpha1"
)

// StandardAppVersionLister helps list StandardAppVersions.
// All objects returned here must be treated as read-only.
type StandardAppVersionLister interface {
	// List lists all StandardAppVersions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StandardAppVersion, err error)
	// StandardAppVersions returns an object that can list and get StandardAppVersions.
	StandardAppVersions(namespace string) StandardAppVersionNamespaceLister
	StandardAppVersionListerExpansion
}

// standardAppVersionLister implements the StandardAppVersionLister interface.
type standardAppVersionLister struct {
	indexer cache.Indexer
}

// NewStandardAppVersionLister returns a new StandardAppVersionLister.
func NewStandardAppVersionLister(indexer cache.Indexer) StandardAppVersionLister {
	return &standardAppVersionLister{indexer: indexer}
}

// List lists all StandardAppVersions in the indexer.
func (s *standardAppVersionLister) List(selector labels.Selector) (ret []*v1alpha1.StandardAppVersion, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StandardAppVersion))
	})
	return ret, err
}

// StandardAppVersions returns an object that can list and get StandardAppVersions.
func (s *standardAppVersionLister) StandardAppVersions(namespace string) StandardAppVersionNamespaceLister {
	return standardAppVersionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StandardAppVersionNamespaceLister helps list and get StandardAppVersions.
// All objects returned here must be treated as read-only.
type StandardAppVersionNamespaceLister interface {
	// List lists all StandardAppVersions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StandardAppVersion, err error)
	// Get retrieves the StandardAppVersion from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.StandardAppVersion, error)
	StandardAppVersionNamespaceListerExpansion
}

// standardAppVersionNamespaceLister implements the StandardAppVersionNamespaceLister
// interface.
type standardAppVersionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StandardAppVersions in the indexer for a given namespace.
func (s standardAppVersionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StandardAppVersion, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StandardAppVersion))
	})
	return ret, err
}

// Get retrieves the StandardAppVersion from the indexer for a given namespace and name.
func (s standardAppVersionNamespaceLister) Get(name string) (*v1alpha1.StandardAppVersion, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("standardappversion"), name)
	}
	return obj.(*v1alpha1.StandardAppVersion), nil
}
