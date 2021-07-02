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

// BucketIamMemberLister helps list BucketIamMembers.
// All objects returned here must be treated as read-only.
type BucketIamMemberLister interface {
	// List lists all BucketIamMembers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BucketIamMember, err error)
	// BucketIamMembers returns an object that can list and get BucketIamMembers.
	BucketIamMembers(namespace string) BucketIamMemberNamespaceLister
	BucketIamMemberListerExpansion
}

// bucketIamMemberLister implements the BucketIamMemberLister interface.
type bucketIamMemberLister struct {
	indexer cache.Indexer
}

// NewBucketIamMemberLister returns a new BucketIamMemberLister.
func NewBucketIamMemberLister(indexer cache.Indexer) BucketIamMemberLister {
	return &bucketIamMemberLister{indexer: indexer}
}

// List lists all BucketIamMembers in the indexer.
func (s *bucketIamMemberLister) List(selector labels.Selector) (ret []*v1alpha1.BucketIamMember, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BucketIamMember))
	})
	return ret, err
}

// BucketIamMembers returns an object that can list and get BucketIamMembers.
func (s *bucketIamMemberLister) BucketIamMembers(namespace string) BucketIamMemberNamespaceLister {
	return bucketIamMemberNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BucketIamMemberNamespaceLister helps list and get BucketIamMembers.
// All objects returned here must be treated as read-only.
type BucketIamMemberNamespaceLister interface {
	// List lists all BucketIamMembers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BucketIamMember, err error)
	// Get retrieves the BucketIamMember from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BucketIamMember, error)
	BucketIamMemberNamespaceListerExpansion
}

// bucketIamMemberNamespaceLister implements the BucketIamMemberNamespaceLister
// interface.
type bucketIamMemberNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BucketIamMembers in the indexer for a given namespace.
func (s bucketIamMemberNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BucketIamMember, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BucketIamMember))
	})
	return ret, err
}

// Get retrieves the BucketIamMember from the indexer for a given namespace and name.
func (s bucketIamMemberNamespaceLister) Get(name string) (*v1alpha1.BucketIamMember, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("bucketiammember"), name)
	}
	return obj.(*v1alpha1.BucketIamMember), nil
}
