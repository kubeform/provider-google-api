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
	v1alpha1 "kubeform.dev/provider-google-api/apis/secret/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ManagerSecretIamMemberLister helps list ManagerSecretIamMembers.
// All objects returned here must be treated as read-only.
type ManagerSecretIamMemberLister interface {
	// List lists all ManagerSecretIamMembers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagerSecretIamMember, err error)
	// ManagerSecretIamMembers returns an object that can list and get ManagerSecretIamMembers.
	ManagerSecretIamMembers(namespace string) ManagerSecretIamMemberNamespaceLister
	ManagerSecretIamMemberListerExpansion
}

// managerSecretIamMemberLister implements the ManagerSecretIamMemberLister interface.
type managerSecretIamMemberLister struct {
	indexer cache.Indexer
}

// NewManagerSecretIamMemberLister returns a new ManagerSecretIamMemberLister.
func NewManagerSecretIamMemberLister(indexer cache.Indexer) ManagerSecretIamMemberLister {
	return &managerSecretIamMemberLister{indexer: indexer}
}

// List lists all ManagerSecretIamMembers in the indexer.
func (s *managerSecretIamMemberLister) List(selector labels.Selector) (ret []*v1alpha1.ManagerSecretIamMember, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagerSecretIamMember))
	})
	return ret, err
}

// ManagerSecretIamMembers returns an object that can list and get ManagerSecretIamMembers.
func (s *managerSecretIamMemberLister) ManagerSecretIamMembers(namespace string) ManagerSecretIamMemberNamespaceLister {
	return managerSecretIamMemberNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ManagerSecretIamMemberNamespaceLister helps list and get ManagerSecretIamMembers.
// All objects returned here must be treated as read-only.
type ManagerSecretIamMemberNamespaceLister interface {
	// List lists all ManagerSecretIamMembers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagerSecretIamMember, err error)
	// Get retrieves the ManagerSecretIamMember from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ManagerSecretIamMember, error)
	ManagerSecretIamMemberNamespaceListerExpansion
}

// managerSecretIamMemberNamespaceLister implements the ManagerSecretIamMemberNamespaceLister
// interface.
type managerSecretIamMemberNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ManagerSecretIamMembers in the indexer for a given namespace.
func (s managerSecretIamMemberNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ManagerSecretIamMember, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagerSecretIamMember))
	})
	return ret, err
}

// Get retrieves the ManagerSecretIamMember from the indexer for a given namespace and name.
func (s managerSecretIamMemberNamespaceLister) Get(name string) (*v1alpha1.ManagerSecretIamMember, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("managersecretiammember"), name)
	}
	return obj.(*v1alpha1.ManagerSecretIamMember), nil
}
