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
	v1alpha1 "kubeform.dev/provider-google-api/apis/compute/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BackendBucketLister helps list BackendBuckets.
// All objects returned here must be treated as read-only.
type BackendBucketLister interface {
	// List lists all BackendBuckets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BackendBucket, err error)
	// BackendBuckets returns an object that can list and get BackendBuckets.
	BackendBuckets(namespace string) BackendBucketNamespaceLister
	BackendBucketListerExpansion
}

// backendBucketLister implements the BackendBucketLister interface.
type backendBucketLister struct {
	indexer cache.Indexer
}

// NewBackendBucketLister returns a new BackendBucketLister.
func NewBackendBucketLister(indexer cache.Indexer) BackendBucketLister {
	return &backendBucketLister{indexer: indexer}
}

// List lists all BackendBuckets in the indexer.
func (s *backendBucketLister) List(selector labels.Selector) (ret []*v1alpha1.BackendBucket, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BackendBucket))
	})
	return ret, err
}

// BackendBuckets returns an object that can list and get BackendBuckets.
func (s *backendBucketLister) BackendBuckets(namespace string) BackendBucketNamespaceLister {
	return backendBucketNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BackendBucketNamespaceLister helps list and get BackendBuckets.
// All objects returned here must be treated as read-only.
type BackendBucketNamespaceLister interface {
	// List lists all BackendBuckets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BackendBucket, err error)
	// Get retrieves the BackendBucket from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BackendBucket, error)
	BackendBucketNamespaceListerExpansion
}

// backendBucketNamespaceLister implements the BackendBucketNamespaceLister
// interface.
type backendBucketNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BackendBuckets in the indexer for a given namespace.
func (s backendBucketNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BackendBucket, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BackendBucket))
	})
	return ret, err
}

// Get retrieves the BackendBucket from the indexer for a given namespace and name.
func (s backendBucketNamespaceLister) Get(name string) (*v1alpha1.BackendBucket, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("backendbucket"), name)
	}
	return obj.(*v1alpha1.BackendBucket), nil
}
