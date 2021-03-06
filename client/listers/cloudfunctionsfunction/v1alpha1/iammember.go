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
	v1alpha1 "kubeform.dev/provider-google-api/apis/cloudfunctionsfunction/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IamMemberLister helps list IamMembers.
// All objects returned here must be treated as read-only.
type IamMemberLister interface {
	// List lists all IamMembers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IamMember, err error)
	// IamMembers returns an object that can list and get IamMembers.
	IamMembers(namespace string) IamMemberNamespaceLister
	IamMemberListerExpansion
}

// iamMemberLister implements the IamMemberLister interface.
type iamMemberLister struct {
	indexer cache.Indexer
}

// NewIamMemberLister returns a new IamMemberLister.
func NewIamMemberLister(indexer cache.Indexer) IamMemberLister {
	return &iamMemberLister{indexer: indexer}
}

// List lists all IamMembers in the indexer.
func (s *iamMemberLister) List(selector labels.Selector) (ret []*v1alpha1.IamMember, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamMember))
	})
	return ret, err
}

// IamMembers returns an object that can list and get IamMembers.
func (s *iamMemberLister) IamMembers(namespace string) IamMemberNamespaceLister {
	return iamMemberNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IamMemberNamespaceLister helps list and get IamMembers.
// All objects returned here must be treated as read-only.
type IamMemberNamespaceLister interface {
	// List lists all IamMembers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IamMember, err error)
	// Get retrieves the IamMember from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.IamMember, error)
	IamMemberNamespaceListerExpansion
}

// iamMemberNamespaceLister implements the IamMemberNamespaceLister
// interface.
type iamMemberNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IamMembers in the indexer for a given namespace.
func (s iamMemberNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IamMember, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamMember))
	})
	return ret, err
}

// Get retrieves the IamMember from the indexer for a given namespace and name.
func (s iamMemberNamespaceLister) Get(name string) (*v1alpha1.IamMember, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iammember"), name)
	}
	return obj.(*v1alpha1.IamMember), nil
}
