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

// DatasetIamMemberLister helps list DatasetIamMembers.
// All objects returned here must be treated as read-only.
type DatasetIamMemberLister interface {
	// List lists all DatasetIamMembers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DatasetIamMember, err error)
	// DatasetIamMembers returns an object that can list and get DatasetIamMembers.
	DatasetIamMembers(namespace string) DatasetIamMemberNamespaceLister
	DatasetIamMemberListerExpansion
}

// datasetIamMemberLister implements the DatasetIamMemberLister interface.
type datasetIamMemberLister struct {
	indexer cache.Indexer
}

// NewDatasetIamMemberLister returns a new DatasetIamMemberLister.
func NewDatasetIamMemberLister(indexer cache.Indexer) DatasetIamMemberLister {
	return &datasetIamMemberLister{indexer: indexer}
}

// List lists all DatasetIamMembers in the indexer.
func (s *datasetIamMemberLister) List(selector labels.Selector) (ret []*v1alpha1.DatasetIamMember, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DatasetIamMember))
	})
	return ret, err
}

// DatasetIamMembers returns an object that can list and get DatasetIamMembers.
func (s *datasetIamMemberLister) DatasetIamMembers(namespace string) DatasetIamMemberNamespaceLister {
	return datasetIamMemberNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DatasetIamMemberNamespaceLister helps list and get DatasetIamMembers.
// All objects returned here must be treated as read-only.
type DatasetIamMemberNamespaceLister interface {
	// List lists all DatasetIamMembers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DatasetIamMember, err error)
	// Get retrieves the DatasetIamMember from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DatasetIamMember, error)
	DatasetIamMemberNamespaceListerExpansion
}

// datasetIamMemberNamespaceLister implements the DatasetIamMemberNamespaceLister
// interface.
type datasetIamMemberNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DatasetIamMembers in the indexer for a given namespace.
func (s datasetIamMemberNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DatasetIamMember, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DatasetIamMember))
	})
	return ret, err
}

// Get retrieves the DatasetIamMember from the indexer for a given namespace and name.
func (s datasetIamMemberNamespaceLister) Get(name string) (*v1alpha1.DatasetIamMember, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("datasetiammember"), name)
	}
	return obj.(*v1alpha1.DatasetIamMember), nil
}
