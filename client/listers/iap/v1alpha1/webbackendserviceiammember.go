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
	v1alpha1 "kubeform.dev/provider-google-api/apis/iap/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WebBackendServiceIamMemberLister helps list WebBackendServiceIamMembers.
// All objects returned here must be treated as read-only.
type WebBackendServiceIamMemberLister interface {
	// List lists all WebBackendServiceIamMembers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WebBackendServiceIamMember, err error)
	// WebBackendServiceIamMembers returns an object that can list and get WebBackendServiceIamMembers.
	WebBackendServiceIamMembers(namespace string) WebBackendServiceIamMemberNamespaceLister
	WebBackendServiceIamMemberListerExpansion
}

// webBackendServiceIamMemberLister implements the WebBackendServiceIamMemberLister interface.
type webBackendServiceIamMemberLister struct {
	indexer cache.Indexer
}

// NewWebBackendServiceIamMemberLister returns a new WebBackendServiceIamMemberLister.
func NewWebBackendServiceIamMemberLister(indexer cache.Indexer) WebBackendServiceIamMemberLister {
	return &webBackendServiceIamMemberLister{indexer: indexer}
}

// List lists all WebBackendServiceIamMembers in the indexer.
func (s *webBackendServiceIamMemberLister) List(selector labels.Selector) (ret []*v1alpha1.WebBackendServiceIamMember, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WebBackendServiceIamMember))
	})
	return ret, err
}

// WebBackendServiceIamMembers returns an object that can list and get WebBackendServiceIamMembers.
func (s *webBackendServiceIamMemberLister) WebBackendServiceIamMembers(namespace string) WebBackendServiceIamMemberNamespaceLister {
	return webBackendServiceIamMemberNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WebBackendServiceIamMemberNamespaceLister helps list and get WebBackendServiceIamMembers.
// All objects returned here must be treated as read-only.
type WebBackendServiceIamMemberNamespaceLister interface {
	// List lists all WebBackendServiceIamMembers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WebBackendServiceIamMember, err error)
	// Get retrieves the WebBackendServiceIamMember from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.WebBackendServiceIamMember, error)
	WebBackendServiceIamMemberNamespaceListerExpansion
}

// webBackendServiceIamMemberNamespaceLister implements the WebBackendServiceIamMemberNamespaceLister
// interface.
type webBackendServiceIamMemberNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WebBackendServiceIamMembers in the indexer for a given namespace.
func (s webBackendServiceIamMemberNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WebBackendServiceIamMember, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WebBackendServiceIamMember))
	})
	return ret, err
}

// Get retrieves the WebBackendServiceIamMember from the indexer for a given namespace and name.
func (s webBackendServiceIamMemberNamespaceLister) Get(name string) (*v1alpha1.WebBackendServiceIamMember, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("webbackendserviceiammember"), name)
	}
	return obj.(*v1alpha1.WebBackendServiceIamMember), nil
}
