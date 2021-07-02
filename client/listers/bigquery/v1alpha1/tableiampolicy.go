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
	v1alpha1 "kubeform.dev/provider-google-api/apis/bigquery/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TableIamPolicyLister helps list TableIamPolicies.
// All objects returned here must be treated as read-only.
type TableIamPolicyLister interface {
	// List lists all TableIamPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TableIamPolicy, err error)
	// TableIamPolicies returns an object that can list and get TableIamPolicies.
	TableIamPolicies(namespace string) TableIamPolicyNamespaceLister
	TableIamPolicyListerExpansion
}

// tableIamPolicyLister implements the TableIamPolicyLister interface.
type tableIamPolicyLister struct {
	indexer cache.Indexer
}

// NewTableIamPolicyLister returns a new TableIamPolicyLister.
func NewTableIamPolicyLister(indexer cache.Indexer) TableIamPolicyLister {
	return &tableIamPolicyLister{indexer: indexer}
}

// List lists all TableIamPolicies in the indexer.
func (s *tableIamPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.TableIamPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TableIamPolicy))
	})
	return ret, err
}

// TableIamPolicies returns an object that can list and get TableIamPolicies.
func (s *tableIamPolicyLister) TableIamPolicies(namespace string) TableIamPolicyNamespaceLister {
	return tableIamPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TableIamPolicyNamespaceLister helps list and get TableIamPolicies.
// All objects returned here must be treated as read-only.
type TableIamPolicyNamespaceLister interface {
	// List lists all TableIamPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TableIamPolicy, err error)
	// Get retrieves the TableIamPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.TableIamPolicy, error)
	TableIamPolicyNamespaceListerExpansion
}

// tableIamPolicyNamespaceLister implements the TableIamPolicyNamespaceLister
// interface.
type tableIamPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TableIamPolicies in the indexer for a given namespace.
func (s tableIamPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TableIamPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TableIamPolicy))
	})
	return ret, err
}

// Get retrieves the TableIamPolicy from the indexer for a given namespace and name.
func (s tableIamPolicyNamespaceLister) Get(name string) (*v1alpha1.TableIamPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("tableiampolicy"), name)
	}
	return obj.(*v1alpha1.TableIamPolicy), nil
}
