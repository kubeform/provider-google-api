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
	v1alpha1 "kubeform.dev/provider-google-api/apis/bigtable/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InstanceIamMemberLister helps list InstanceIamMembers.
// All objects returned here must be treated as read-only.
type InstanceIamMemberLister interface {
	// List lists all InstanceIamMembers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InstanceIamMember, err error)
	// InstanceIamMembers returns an object that can list and get InstanceIamMembers.
	InstanceIamMembers(namespace string) InstanceIamMemberNamespaceLister
	InstanceIamMemberListerExpansion
}

// instanceIamMemberLister implements the InstanceIamMemberLister interface.
type instanceIamMemberLister struct {
	indexer cache.Indexer
}

// NewInstanceIamMemberLister returns a new InstanceIamMemberLister.
func NewInstanceIamMemberLister(indexer cache.Indexer) InstanceIamMemberLister {
	return &instanceIamMemberLister{indexer: indexer}
}

// List lists all InstanceIamMembers in the indexer.
func (s *instanceIamMemberLister) List(selector labels.Selector) (ret []*v1alpha1.InstanceIamMember, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InstanceIamMember))
	})
	return ret, err
}

// InstanceIamMembers returns an object that can list and get InstanceIamMembers.
func (s *instanceIamMemberLister) InstanceIamMembers(namespace string) InstanceIamMemberNamespaceLister {
	return instanceIamMemberNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InstanceIamMemberNamespaceLister helps list and get InstanceIamMembers.
// All objects returned here must be treated as read-only.
type InstanceIamMemberNamespaceLister interface {
	// List lists all InstanceIamMembers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InstanceIamMember, err error)
	// Get retrieves the InstanceIamMember from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.InstanceIamMember, error)
	InstanceIamMemberNamespaceListerExpansion
}

// instanceIamMemberNamespaceLister implements the InstanceIamMemberNamespaceLister
// interface.
type instanceIamMemberNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all InstanceIamMembers in the indexer for a given namespace.
func (s instanceIamMemberNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.InstanceIamMember, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InstanceIamMember))
	})
	return ret, err
}

// Get retrieves the InstanceIamMember from the indexer for a given namespace and name.
func (s instanceIamMemberNamespaceLister) Get(name string) (*v1alpha1.InstanceIamMember, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("instanceiammember"), name)
	}
	return obj.(*v1alpha1.InstanceIamMember), nil
}
