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
	v1alpha1 "kubeform.dev/provider-google-api/apis/dataproc/v1alpha1"
)

// JobIamMemberLister helps list JobIamMembers.
// All objects returned here must be treated as read-only.
type JobIamMemberLister interface {
	// List lists all JobIamMembers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.JobIamMember, err error)
	// JobIamMembers returns an object that can list and get JobIamMembers.
	JobIamMembers(namespace string) JobIamMemberNamespaceLister
	JobIamMemberListerExpansion
}

// jobIamMemberLister implements the JobIamMemberLister interface.
type jobIamMemberLister struct {
	indexer cache.Indexer
}

// NewJobIamMemberLister returns a new JobIamMemberLister.
func NewJobIamMemberLister(indexer cache.Indexer) JobIamMemberLister {
	return &jobIamMemberLister{indexer: indexer}
}

// List lists all JobIamMembers in the indexer.
func (s *jobIamMemberLister) List(selector labels.Selector) (ret []*v1alpha1.JobIamMember, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.JobIamMember))
	})
	return ret, err
}

// JobIamMembers returns an object that can list and get JobIamMembers.
func (s *jobIamMemberLister) JobIamMembers(namespace string) JobIamMemberNamespaceLister {
	return jobIamMemberNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// JobIamMemberNamespaceLister helps list and get JobIamMembers.
// All objects returned here must be treated as read-only.
type JobIamMemberNamespaceLister interface {
	// List lists all JobIamMembers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.JobIamMember, err error)
	// Get retrieves the JobIamMember from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.JobIamMember, error)
	JobIamMemberNamespaceListerExpansion
}

// jobIamMemberNamespaceLister implements the JobIamMemberNamespaceLister
// interface.
type jobIamMemberNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all JobIamMembers in the indexer for a given namespace.
func (s jobIamMemberNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.JobIamMember, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.JobIamMember))
	})
	return ret, err
}

// Get retrieves the JobIamMember from the indexer for a given namespace and name.
func (s jobIamMemberNamespaceLister) Get(name string) (*v1alpha1.JobIamMember, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("jobiammember"), name)
	}
	return obj.(*v1alpha1.JobIamMember), nil
}
