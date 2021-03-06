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
	v1alpha1 "kubeform.dev/provider-google-api/apis/apigee/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EnvironmentIamPolicyLister helps list EnvironmentIamPolicies.
// All objects returned here must be treated as read-only.
type EnvironmentIamPolicyLister interface {
	// List lists all EnvironmentIamPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EnvironmentIamPolicy, err error)
	// EnvironmentIamPolicies returns an object that can list and get EnvironmentIamPolicies.
	EnvironmentIamPolicies(namespace string) EnvironmentIamPolicyNamespaceLister
	EnvironmentIamPolicyListerExpansion
}

// environmentIamPolicyLister implements the EnvironmentIamPolicyLister interface.
type environmentIamPolicyLister struct {
	indexer cache.Indexer
}

// NewEnvironmentIamPolicyLister returns a new EnvironmentIamPolicyLister.
func NewEnvironmentIamPolicyLister(indexer cache.Indexer) EnvironmentIamPolicyLister {
	return &environmentIamPolicyLister{indexer: indexer}
}

// List lists all EnvironmentIamPolicies in the indexer.
func (s *environmentIamPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.EnvironmentIamPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EnvironmentIamPolicy))
	})
	return ret, err
}

// EnvironmentIamPolicies returns an object that can list and get EnvironmentIamPolicies.
func (s *environmentIamPolicyLister) EnvironmentIamPolicies(namespace string) EnvironmentIamPolicyNamespaceLister {
	return environmentIamPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EnvironmentIamPolicyNamespaceLister helps list and get EnvironmentIamPolicies.
// All objects returned here must be treated as read-only.
type EnvironmentIamPolicyNamespaceLister interface {
	// List lists all EnvironmentIamPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EnvironmentIamPolicy, err error)
	// Get retrieves the EnvironmentIamPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EnvironmentIamPolicy, error)
	EnvironmentIamPolicyNamespaceListerExpansion
}

// environmentIamPolicyNamespaceLister implements the EnvironmentIamPolicyNamespaceLister
// interface.
type environmentIamPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EnvironmentIamPolicies in the indexer for a given namespace.
func (s environmentIamPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EnvironmentIamPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EnvironmentIamPolicy))
	})
	return ret, err
}

// Get retrieves the EnvironmentIamPolicy from the indexer for a given namespace and name.
func (s environmentIamPolicyNamespaceLister) Get(name string) (*v1alpha1.EnvironmentIamPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("environmentiampolicy"), name)
	}
	return obj.(*v1alpha1.EnvironmentIamPolicy), nil
}
