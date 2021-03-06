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

// LossPreventionStoredInfoTypeLister helps list LossPreventionStoredInfoTypes.
// All objects returned here must be treated as read-only.
type LossPreventionStoredInfoTypeLister interface {
	// List lists all LossPreventionStoredInfoTypes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LossPreventionStoredInfoType, err error)
	// LossPreventionStoredInfoTypes returns an object that can list and get LossPreventionStoredInfoTypes.
	LossPreventionStoredInfoTypes(namespace string) LossPreventionStoredInfoTypeNamespaceLister
	LossPreventionStoredInfoTypeListerExpansion
}

// lossPreventionStoredInfoTypeLister implements the LossPreventionStoredInfoTypeLister interface.
type lossPreventionStoredInfoTypeLister struct {
	indexer cache.Indexer
}

// NewLossPreventionStoredInfoTypeLister returns a new LossPreventionStoredInfoTypeLister.
func NewLossPreventionStoredInfoTypeLister(indexer cache.Indexer) LossPreventionStoredInfoTypeLister {
	return &lossPreventionStoredInfoTypeLister{indexer: indexer}
}

// List lists all LossPreventionStoredInfoTypes in the indexer.
func (s *lossPreventionStoredInfoTypeLister) List(selector labels.Selector) (ret []*v1alpha1.LossPreventionStoredInfoType, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LossPreventionStoredInfoType))
	})
	return ret, err
}

// LossPreventionStoredInfoTypes returns an object that can list and get LossPreventionStoredInfoTypes.
func (s *lossPreventionStoredInfoTypeLister) LossPreventionStoredInfoTypes(namespace string) LossPreventionStoredInfoTypeNamespaceLister {
	return lossPreventionStoredInfoTypeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LossPreventionStoredInfoTypeNamespaceLister helps list and get LossPreventionStoredInfoTypes.
// All objects returned here must be treated as read-only.
type LossPreventionStoredInfoTypeNamespaceLister interface {
	// List lists all LossPreventionStoredInfoTypes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LossPreventionStoredInfoType, err error)
	// Get retrieves the LossPreventionStoredInfoType from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.LossPreventionStoredInfoType, error)
	LossPreventionStoredInfoTypeNamespaceListerExpansion
}

// lossPreventionStoredInfoTypeNamespaceLister implements the LossPreventionStoredInfoTypeNamespaceLister
// interface.
type lossPreventionStoredInfoTypeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LossPreventionStoredInfoTypes in the indexer for a given namespace.
func (s lossPreventionStoredInfoTypeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LossPreventionStoredInfoType, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LossPreventionStoredInfoType))
	})
	return ret, err
}

// Get retrieves the LossPreventionStoredInfoType from the indexer for a given namespace and name.
func (s lossPreventionStoredInfoTypeNamespaceLister) Get(name string) (*v1alpha1.LossPreventionStoredInfoType, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("losspreventionstoredinfotype"), name)
	}
	return obj.(*v1alpha1.LossPreventionStoredInfoType), nil
}
