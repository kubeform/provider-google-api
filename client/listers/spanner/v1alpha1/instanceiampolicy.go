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
	v1alpha1 "kubeform.dev/provider-google-api/apis/spanner/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InstanceIamPolicyLister helps list InstanceIamPolicies.
// All objects returned here must be treated as read-only.
type InstanceIamPolicyLister interface {
	// List lists all InstanceIamPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InstanceIamPolicy, err error)
	// InstanceIamPolicies returns an object that can list and get InstanceIamPolicies.
	InstanceIamPolicies(namespace string) InstanceIamPolicyNamespaceLister
	InstanceIamPolicyListerExpansion
}

// instanceIamPolicyLister implements the InstanceIamPolicyLister interface.
type instanceIamPolicyLister struct {
	indexer cache.Indexer
}

// NewInstanceIamPolicyLister returns a new InstanceIamPolicyLister.
func NewInstanceIamPolicyLister(indexer cache.Indexer) InstanceIamPolicyLister {
	return &instanceIamPolicyLister{indexer: indexer}
}

// List lists all InstanceIamPolicies in the indexer.
func (s *instanceIamPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.InstanceIamPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InstanceIamPolicy))
	})
	return ret, err
}

// InstanceIamPolicies returns an object that can list and get InstanceIamPolicies.
func (s *instanceIamPolicyLister) InstanceIamPolicies(namespace string) InstanceIamPolicyNamespaceLister {
	return instanceIamPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InstanceIamPolicyNamespaceLister helps list and get InstanceIamPolicies.
// All objects returned here must be treated as read-only.
type InstanceIamPolicyNamespaceLister interface {
	// List lists all InstanceIamPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InstanceIamPolicy, err error)
	// Get retrieves the InstanceIamPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.InstanceIamPolicy, error)
	InstanceIamPolicyNamespaceListerExpansion
}

// instanceIamPolicyNamespaceLister implements the InstanceIamPolicyNamespaceLister
// interface.
type instanceIamPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all InstanceIamPolicies in the indexer for a given namespace.
func (s instanceIamPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.InstanceIamPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InstanceIamPolicy))
	})
	return ret, err
}

// Get retrieves the InstanceIamPolicy from the indexer for a given namespace and name.
func (s instanceIamPolicyNamespaceLister) Get(name string) (*v1alpha1.InstanceIamPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("instanceiampolicy"), name)
	}
	return obj.(*v1alpha1.InstanceIamPolicy), nil
}
