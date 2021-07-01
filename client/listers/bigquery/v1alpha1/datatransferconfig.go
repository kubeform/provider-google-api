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
	v1alpha1 "kubeform.dev/provider-google-api/apis/bigquery/v1alpha1"
)

// DataTransferConfigLister helps list DataTransferConfigs.
// All objects returned here must be treated as read-only.
type DataTransferConfigLister interface {
	// List lists all DataTransferConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DataTransferConfig, err error)
	// DataTransferConfigs returns an object that can list and get DataTransferConfigs.
	DataTransferConfigs(namespace string) DataTransferConfigNamespaceLister
	DataTransferConfigListerExpansion
}

// dataTransferConfigLister implements the DataTransferConfigLister interface.
type dataTransferConfigLister struct {
	indexer cache.Indexer
}

// NewDataTransferConfigLister returns a new DataTransferConfigLister.
func NewDataTransferConfigLister(indexer cache.Indexer) DataTransferConfigLister {
	return &dataTransferConfigLister{indexer: indexer}
}

// List lists all DataTransferConfigs in the indexer.
func (s *dataTransferConfigLister) List(selector labels.Selector) (ret []*v1alpha1.DataTransferConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataTransferConfig))
	})
	return ret, err
}

// DataTransferConfigs returns an object that can list and get DataTransferConfigs.
func (s *dataTransferConfigLister) DataTransferConfigs(namespace string) DataTransferConfigNamespaceLister {
	return dataTransferConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DataTransferConfigNamespaceLister helps list and get DataTransferConfigs.
// All objects returned here must be treated as read-only.
type DataTransferConfigNamespaceLister interface {
	// List lists all DataTransferConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DataTransferConfig, err error)
	// Get retrieves the DataTransferConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DataTransferConfig, error)
	DataTransferConfigNamespaceListerExpansion
}

// dataTransferConfigNamespaceLister implements the DataTransferConfigNamespaceLister
// interface.
type dataTransferConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DataTransferConfigs in the indexer for a given namespace.
func (s dataTransferConfigNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DataTransferConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataTransferConfig))
	})
	return ret, err
}

// Get retrieves the DataTransferConfig from the indexer for a given namespace and name.
func (s dataTransferConfigNamespaceLister) Get(name string) (*v1alpha1.DataTransferConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("datatransferconfig"), name)
	}
	return obj.(*v1alpha1.DataTransferConfig), nil
}
