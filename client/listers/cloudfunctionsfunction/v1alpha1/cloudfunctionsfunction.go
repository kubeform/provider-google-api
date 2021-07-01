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
	v1alpha1 "kubeform.dev/provider-google-api/apis/cloudfunctionsfunction/v1alpha1"
)

// CloudfunctionsFunctionLister helps list CloudfunctionsFunctions.
// All objects returned here must be treated as read-only.
type CloudfunctionsFunctionLister interface {
	// List lists all CloudfunctionsFunctions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CloudfunctionsFunction, err error)
	// CloudfunctionsFunctions returns an object that can list and get CloudfunctionsFunctions.
	CloudfunctionsFunctions(namespace string) CloudfunctionsFunctionNamespaceLister
	CloudfunctionsFunctionListerExpansion
}

// cloudfunctionsFunctionLister implements the CloudfunctionsFunctionLister interface.
type cloudfunctionsFunctionLister struct {
	indexer cache.Indexer
}

// NewCloudfunctionsFunctionLister returns a new CloudfunctionsFunctionLister.
func NewCloudfunctionsFunctionLister(indexer cache.Indexer) CloudfunctionsFunctionLister {
	return &cloudfunctionsFunctionLister{indexer: indexer}
}

// List lists all CloudfunctionsFunctions in the indexer.
func (s *cloudfunctionsFunctionLister) List(selector labels.Selector) (ret []*v1alpha1.CloudfunctionsFunction, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudfunctionsFunction))
	})
	return ret, err
}

// CloudfunctionsFunctions returns an object that can list and get CloudfunctionsFunctions.
func (s *cloudfunctionsFunctionLister) CloudfunctionsFunctions(namespace string) CloudfunctionsFunctionNamespaceLister {
	return cloudfunctionsFunctionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CloudfunctionsFunctionNamespaceLister helps list and get CloudfunctionsFunctions.
// All objects returned here must be treated as read-only.
type CloudfunctionsFunctionNamespaceLister interface {
	// List lists all CloudfunctionsFunctions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CloudfunctionsFunction, err error)
	// Get retrieves the CloudfunctionsFunction from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CloudfunctionsFunction, error)
	CloudfunctionsFunctionNamespaceListerExpansion
}

// cloudfunctionsFunctionNamespaceLister implements the CloudfunctionsFunctionNamespaceLister
// interface.
type cloudfunctionsFunctionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CloudfunctionsFunctions in the indexer for a given namespace.
func (s cloudfunctionsFunctionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CloudfunctionsFunction, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudfunctionsFunction))
	})
	return ret, err
}

// Get retrieves the CloudfunctionsFunction from the indexer for a given namespace and name.
func (s cloudfunctionsFunctionNamespaceLister) Get(name string) (*v1alpha1.CloudfunctionsFunction, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cloudfunctionsfunction"), name)
	}
	return obj.(*v1alpha1.CloudfunctionsFunction), nil
}
