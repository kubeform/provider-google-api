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
	v1alpha1 "kubeform.dev/provider-google-api/apis/spanner/v1alpha1"
)

// DatabaseIamMemberLister helps list DatabaseIamMembers.
// All objects returned here must be treated as read-only.
type DatabaseIamMemberLister interface {
	// List lists all DatabaseIamMembers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DatabaseIamMember, err error)
	// DatabaseIamMembers returns an object that can list and get DatabaseIamMembers.
	DatabaseIamMembers(namespace string) DatabaseIamMemberNamespaceLister
	DatabaseIamMemberListerExpansion
}

// databaseIamMemberLister implements the DatabaseIamMemberLister interface.
type databaseIamMemberLister struct {
	indexer cache.Indexer
}

// NewDatabaseIamMemberLister returns a new DatabaseIamMemberLister.
func NewDatabaseIamMemberLister(indexer cache.Indexer) DatabaseIamMemberLister {
	return &databaseIamMemberLister{indexer: indexer}
}

// List lists all DatabaseIamMembers in the indexer.
func (s *databaseIamMemberLister) List(selector labels.Selector) (ret []*v1alpha1.DatabaseIamMember, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DatabaseIamMember))
	})
	return ret, err
}

// DatabaseIamMembers returns an object that can list and get DatabaseIamMembers.
func (s *databaseIamMemberLister) DatabaseIamMembers(namespace string) DatabaseIamMemberNamespaceLister {
	return databaseIamMemberNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DatabaseIamMemberNamespaceLister helps list and get DatabaseIamMembers.
// All objects returned here must be treated as read-only.
type DatabaseIamMemberNamespaceLister interface {
	// List lists all DatabaseIamMembers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DatabaseIamMember, err error)
	// Get retrieves the DatabaseIamMember from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DatabaseIamMember, error)
	DatabaseIamMemberNamespaceListerExpansion
}

// databaseIamMemberNamespaceLister implements the DatabaseIamMemberNamespaceLister
// interface.
type databaseIamMemberNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DatabaseIamMembers in the indexer for a given namespace.
func (s databaseIamMemberNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DatabaseIamMember, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DatabaseIamMember))
	})
	return ret, err
}

// Get retrieves the DatabaseIamMember from the indexer for a given namespace and name.
func (s databaseIamMemberNamespaceLister) Get(name string) (*v1alpha1.DatabaseIamMember, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("databaseiammember"), name)
	}
	return obj.(*v1alpha1.DatabaseIamMember), nil
}
