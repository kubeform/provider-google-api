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
	v1alpha1 "kubeform.dev/provider-google-api/apis/cloud/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IdentityGroupLister helps list IdentityGroups.
// All objects returned here must be treated as read-only.
type IdentityGroupLister interface {
	// List lists all IdentityGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IdentityGroup, err error)
	// IdentityGroups returns an object that can list and get IdentityGroups.
	IdentityGroups(namespace string) IdentityGroupNamespaceLister
	IdentityGroupListerExpansion
}

// identityGroupLister implements the IdentityGroupLister interface.
type identityGroupLister struct {
	indexer cache.Indexer
}

// NewIdentityGroupLister returns a new IdentityGroupLister.
func NewIdentityGroupLister(indexer cache.Indexer) IdentityGroupLister {
	return &identityGroupLister{indexer: indexer}
}

// List lists all IdentityGroups in the indexer.
func (s *identityGroupLister) List(selector labels.Selector) (ret []*v1alpha1.IdentityGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IdentityGroup))
	})
	return ret, err
}

// IdentityGroups returns an object that can list and get IdentityGroups.
func (s *identityGroupLister) IdentityGroups(namespace string) IdentityGroupNamespaceLister {
	return identityGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IdentityGroupNamespaceLister helps list and get IdentityGroups.
// All objects returned here must be treated as read-only.
type IdentityGroupNamespaceLister interface {
	// List lists all IdentityGroups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IdentityGroup, err error)
	// Get retrieves the IdentityGroup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.IdentityGroup, error)
	IdentityGroupNamespaceListerExpansion
}

// identityGroupNamespaceLister implements the IdentityGroupNamespaceLister
// interface.
type identityGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IdentityGroups in the indexer for a given namespace.
func (s identityGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IdentityGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IdentityGroup))
	})
	return ret, err
}

// Get retrieves the IdentityGroup from the indexer for a given namespace and name.
func (s identityGroupNamespaceLister) Get(name string) (*v1alpha1.IdentityGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("identitygroup"), name)
	}
	return obj.(*v1alpha1.IdentityGroup), nil
}
