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
	v1alpha1 "kubeform.dev/provider-google-api/apis/data/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CatalogTagTemplateIamPolicyLister helps list CatalogTagTemplateIamPolicies.
// All objects returned here must be treated as read-only.
type CatalogTagTemplateIamPolicyLister interface {
	// List lists all CatalogTagTemplateIamPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CatalogTagTemplateIamPolicy, err error)
	// CatalogTagTemplateIamPolicies returns an object that can list and get CatalogTagTemplateIamPolicies.
	CatalogTagTemplateIamPolicies(namespace string) CatalogTagTemplateIamPolicyNamespaceLister
	CatalogTagTemplateIamPolicyListerExpansion
}

// catalogTagTemplateIamPolicyLister implements the CatalogTagTemplateIamPolicyLister interface.
type catalogTagTemplateIamPolicyLister struct {
	indexer cache.Indexer
}

// NewCatalogTagTemplateIamPolicyLister returns a new CatalogTagTemplateIamPolicyLister.
func NewCatalogTagTemplateIamPolicyLister(indexer cache.Indexer) CatalogTagTemplateIamPolicyLister {
	return &catalogTagTemplateIamPolicyLister{indexer: indexer}
}

// List lists all CatalogTagTemplateIamPolicies in the indexer.
func (s *catalogTagTemplateIamPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.CatalogTagTemplateIamPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CatalogTagTemplateIamPolicy))
	})
	return ret, err
}

// CatalogTagTemplateIamPolicies returns an object that can list and get CatalogTagTemplateIamPolicies.
func (s *catalogTagTemplateIamPolicyLister) CatalogTagTemplateIamPolicies(namespace string) CatalogTagTemplateIamPolicyNamespaceLister {
	return catalogTagTemplateIamPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CatalogTagTemplateIamPolicyNamespaceLister helps list and get CatalogTagTemplateIamPolicies.
// All objects returned here must be treated as read-only.
type CatalogTagTemplateIamPolicyNamespaceLister interface {
	// List lists all CatalogTagTemplateIamPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CatalogTagTemplateIamPolicy, err error)
	// Get retrieves the CatalogTagTemplateIamPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CatalogTagTemplateIamPolicy, error)
	CatalogTagTemplateIamPolicyNamespaceListerExpansion
}

// catalogTagTemplateIamPolicyNamespaceLister implements the CatalogTagTemplateIamPolicyNamespaceLister
// interface.
type catalogTagTemplateIamPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CatalogTagTemplateIamPolicies in the indexer for a given namespace.
func (s catalogTagTemplateIamPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CatalogTagTemplateIamPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CatalogTagTemplateIamPolicy))
	})
	return ret, err
}

// Get retrieves the CatalogTagTemplateIamPolicy from the indexer for a given namespace and name.
func (s catalogTagTemplateIamPolicyNamespaceLister) Get(name string) (*v1alpha1.CatalogTagTemplateIamPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("catalogtagtemplateiampolicy"), name)
	}
	return obj.(*v1alpha1.CatalogTagTemplateIamPolicy), nil
}
