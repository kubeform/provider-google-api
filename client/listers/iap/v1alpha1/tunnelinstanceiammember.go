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

// TunnelInstanceIamMemberLister helps list TunnelInstanceIamMembers.
// All objects returned here must be treated as read-only.
type TunnelInstanceIamMemberLister interface {
	// List lists all TunnelInstanceIamMembers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TunnelInstanceIamMember, err error)
	// TunnelInstanceIamMembers returns an object that can list and get TunnelInstanceIamMembers.
	TunnelInstanceIamMembers(namespace string) TunnelInstanceIamMemberNamespaceLister
	TunnelInstanceIamMemberListerExpansion
}

// tunnelInstanceIamMemberLister implements the TunnelInstanceIamMemberLister interface.
type tunnelInstanceIamMemberLister struct {
	indexer cache.Indexer
}

// NewTunnelInstanceIamMemberLister returns a new TunnelInstanceIamMemberLister.
func NewTunnelInstanceIamMemberLister(indexer cache.Indexer) TunnelInstanceIamMemberLister {
	return &tunnelInstanceIamMemberLister{indexer: indexer}
}

// List lists all TunnelInstanceIamMembers in the indexer.
func (s *tunnelInstanceIamMemberLister) List(selector labels.Selector) (ret []*v1alpha1.TunnelInstanceIamMember, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TunnelInstanceIamMember))
	})
	return ret, err
}

// TunnelInstanceIamMembers returns an object that can list and get TunnelInstanceIamMembers.
func (s *tunnelInstanceIamMemberLister) TunnelInstanceIamMembers(namespace string) TunnelInstanceIamMemberNamespaceLister {
	return tunnelInstanceIamMemberNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TunnelInstanceIamMemberNamespaceLister helps list and get TunnelInstanceIamMembers.
// All objects returned here must be treated as read-only.
type TunnelInstanceIamMemberNamespaceLister interface {
	// List lists all TunnelInstanceIamMembers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TunnelInstanceIamMember, err error)
	// Get retrieves the TunnelInstanceIamMember from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.TunnelInstanceIamMember, error)
	TunnelInstanceIamMemberNamespaceListerExpansion
}

// tunnelInstanceIamMemberNamespaceLister implements the TunnelInstanceIamMemberNamespaceLister
// interface.
type tunnelInstanceIamMemberNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TunnelInstanceIamMembers in the indexer for a given namespace.
func (s tunnelInstanceIamMemberNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TunnelInstanceIamMember, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TunnelInstanceIamMember))
	})
	return ret, err
}

// Get retrieves the TunnelInstanceIamMember from the indexer for a given namespace and name.
func (s tunnelInstanceIamMemberNamespaceLister) Get(name string) (*v1alpha1.TunnelInstanceIamMember, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("tunnelinstanceiammember"), name)
	}
	return obj.(*v1alpha1.TunnelInstanceIamMember), nil
}
