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
	v1alpha1 "kubeform.dev/provider-google-api/apis/vertex/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AiDatasetLister helps list AiDatasets.
// All objects returned here must be treated as read-only.
type AiDatasetLister interface {
	// List lists all AiDatasets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AiDataset, err error)
	// AiDatasets returns an object that can list and get AiDatasets.
	AiDatasets(namespace string) AiDatasetNamespaceLister
	AiDatasetListerExpansion
}

// aiDatasetLister implements the AiDatasetLister interface.
type aiDatasetLister struct {
	indexer cache.Indexer
}

// NewAiDatasetLister returns a new AiDatasetLister.
func NewAiDatasetLister(indexer cache.Indexer) AiDatasetLister {
	return &aiDatasetLister{indexer: indexer}
}

// List lists all AiDatasets in the indexer.
func (s *aiDatasetLister) List(selector labels.Selector) (ret []*v1alpha1.AiDataset, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AiDataset))
	})
	return ret, err
}

// AiDatasets returns an object that can list and get AiDatasets.
func (s *aiDatasetLister) AiDatasets(namespace string) AiDatasetNamespaceLister {
	return aiDatasetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AiDatasetNamespaceLister helps list and get AiDatasets.
// All objects returned here must be treated as read-only.
type AiDatasetNamespaceLister interface {
	// List lists all AiDatasets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AiDataset, err error)
	// Get retrieves the AiDataset from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AiDataset, error)
	AiDatasetNamespaceListerExpansion
}

// aiDatasetNamespaceLister implements the AiDatasetNamespaceLister
// interface.
type aiDatasetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AiDatasets in the indexer for a given namespace.
func (s aiDatasetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AiDataset, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AiDataset))
	})
	return ret, err
}

// Get retrieves the AiDataset from the indexer for a given namespace and name.
func (s aiDatasetNamespaceLister) Get(name string) (*v1alpha1.AiDataset, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("aidataset"), name)
	}
	return obj.(*v1alpha1.AiDataset), nil
}
