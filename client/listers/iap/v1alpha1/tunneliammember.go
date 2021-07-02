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

// TunnelIamMemberLister helps list TunnelIamMembers.
// All objects returned here must be treated as read-only.
type TunnelIamMemberLister interface {
	// List lists all TunnelIamMembers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TunnelIamMember, err error)
	// TunnelIamMembers returns an object that can list and get TunnelIamMembers.
	TunnelIamMembers(namespace string) TunnelIamMemberNamespaceLister
	TunnelIamMemberListerExpansion
}

// tunnelIamMemberLister implements the TunnelIamMemberLister interface.
type tunnelIamMemberLister struct {
	indexer cache.Indexer
}

// NewTunnelIamMemberLister returns a new TunnelIamMemberLister.
func NewTunnelIamMemberLister(indexer cache.Indexer) TunnelIamMemberLister {
	return &tunnelIamMemberLister{indexer: indexer}
}

// List lists all TunnelIamMembers in the indexer.
func (s *tunnelIamMemberLister) List(selector labels.Selector) (ret []*v1alpha1.TunnelIamMember, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TunnelIamMember))
	})
	return ret, err
}

// TunnelIamMembers returns an object that can list and get TunnelIamMembers.
func (s *tunnelIamMemberLister) TunnelIamMembers(namespace string) TunnelIamMemberNamespaceLister {
	return tunnelIamMemberNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TunnelIamMemberNamespaceLister helps list and get TunnelIamMembers.
// All objects returned here must be treated as read-only.
type TunnelIamMemberNamespaceLister interface {
	// List lists all TunnelIamMembers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TunnelIamMember, err error)
	// Get retrieves the TunnelIamMember from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.TunnelIamMember, error)
	TunnelIamMemberNamespaceListerExpansion
}

// tunnelIamMemberNamespaceLister implements the TunnelIamMemberNamespaceLister
// interface.
type tunnelIamMemberNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TunnelIamMembers in the indexer for a given namespace.
func (s tunnelIamMemberNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TunnelIamMember, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TunnelIamMember))
	})
	return ret, err
}

// Get retrieves the TunnelIamMember from the indexer for a given namespace and name.
func (s tunnelIamMemberNamespaceLister) Get(name string) (*v1alpha1.TunnelIamMember, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("tunneliammember"), name)
	}
	return obj.(*v1alpha1.TunnelIamMember), nil
}
