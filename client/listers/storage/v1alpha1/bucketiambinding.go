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
	v1alpha1 "kubeform.dev/provider-google-api/apis/storage/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BucketIamBindingLister helps list BucketIamBindings.
// All objects returned here must be treated as read-only.
type BucketIamBindingLister interface {
	// List lists all BucketIamBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BucketIamBinding, err error)
	// BucketIamBindings returns an object that can list and get BucketIamBindings.
	BucketIamBindings(namespace string) BucketIamBindingNamespaceLister
	BucketIamBindingListerExpansion
}

// bucketIamBindingLister implements the BucketIamBindingLister interface.
type bucketIamBindingLister struct {
	indexer cache.Indexer
}

// NewBucketIamBindingLister returns a new BucketIamBindingLister.
func NewBucketIamBindingLister(indexer cache.Indexer) BucketIamBindingLister {
	return &bucketIamBindingLister{indexer: indexer}
}

// List lists all BucketIamBindings in the indexer.
func (s *bucketIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.BucketIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BucketIamBinding))
	})
	return ret, err
}

// BucketIamBindings returns an object that can list and get BucketIamBindings.
func (s *bucketIamBindingLister) BucketIamBindings(namespace string) BucketIamBindingNamespaceLister {
	return bucketIamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BucketIamBindingNamespaceLister helps list and get BucketIamBindings.
// All objects returned here must be treated as read-only.
type BucketIamBindingNamespaceLister interface {
	// List lists all BucketIamBindings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BucketIamBinding, err error)
	// Get retrieves the BucketIamBinding from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BucketIamBinding, error)
	BucketIamBindingNamespaceListerExpansion
}

// bucketIamBindingNamespaceLister implements the BucketIamBindingNamespaceLister
// interface.
type bucketIamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BucketIamBindings in the indexer for a given namespace.
func (s bucketIamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BucketIamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BucketIamBinding))
	})
	return ret, err
}

// Get retrieves the BucketIamBinding from the indexer for a given namespace and name.
func (s bucketIamBindingNamespaceLister) Get(name string) (*v1alpha1.BucketIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("bucketiambinding"), name)
	}
	return obj.(*v1alpha1.BucketIamBinding), nil
}
