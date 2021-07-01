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
	v1alpha1 "kubeform.dev/provider-google-api/apis/project/v1alpha1"
)

// IamCustomRoleLister helps list IamCustomRoles.
// All objects returned here must be treated as read-only.
type IamCustomRoleLister interface {
	// List lists all IamCustomRoles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IamCustomRole, err error)
	// IamCustomRoles returns an object that can list and get IamCustomRoles.
	IamCustomRoles(namespace string) IamCustomRoleNamespaceLister
	IamCustomRoleListerExpansion
}

// iamCustomRoleLister implements the IamCustomRoleLister interface.
type iamCustomRoleLister struct {
	indexer cache.Indexer
}

// NewIamCustomRoleLister returns a new IamCustomRoleLister.
func NewIamCustomRoleLister(indexer cache.Indexer) IamCustomRoleLister {
	return &iamCustomRoleLister{indexer: indexer}
}

// List lists all IamCustomRoles in the indexer.
func (s *iamCustomRoleLister) List(selector labels.Selector) (ret []*v1alpha1.IamCustomRole, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamCustomRole))
	})
	return ret, err
}

// IamCustomRoles returns an object that can list and get IamCustomRoles.
func (s *iamCustomRoleLister) IamCustomRoles(namespace string) IamCustomRoleNamespaceLister {
	return iamCustomRoleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IamCustomRoleNamespaceLister helps list and get IamCustomRoles.
// All objects returned here must be treated as read-only.
type IamCustomRoleNamespaceLister interface {
	// List lists all IamCustomRoles in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IamCustomRole, err error)
	// Get retrieves the IamCustomRole from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.IamCustomRole, error)
	IamCustomRoleNamespaceListerExpansion
}

// iamCustomRoleNamespaceLister implements the IamCustomRoleNamespaceLister
// interface.
type iamCustomRoleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IamCustomRoles in the indexer for a given namespace.
func (s iamCustomRoleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IamCustomRole, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamCustomRole))
	})
	return ret, err
}

// Get retrieves the IamCustomRole from the indexer for a given namespace and name.
func (s iamCustomRoleNamespaceLister) Get(name string) (*v1alpha1.IamCustomRole, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iamcustomrole"), name)
	}
	return obj.(*v1alpha1.IamCustomRole), nil
}
