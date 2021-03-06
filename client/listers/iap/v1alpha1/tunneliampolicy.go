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

// TunnelIamPolicyLister helps list TunnelIamPolicies.
// All objects returned here must be treated as read-only.
type TunnelIamPolicyLister interface {
	// List lists all TunnelIamPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TunnelIamPolicy, err error)
	// TunnelIamPolicies returns an object that can list and get TunnelIamPolicies.
	TunnelIamPolicies(namespace string) TunnelIamPolicyNamespaceLister
	TunnelIamPolicyListerExpansion
}

// tunnelIamPolicyLister implements the TunnelIamPolicyLister interface.
type tunnelIamPolicyLister struct {
	indexer cache.Indexer
}

// NewTunnelIamPolicyLister returns a new TunnelIamPolicyLister.
func NewTunnelIamPolicyLister(indexer cache.Indexer) TunnelIamPolicyLister {
	return &tunnelIamPolicyLister{indexer: indexer}
}

// List lists all TunnelIamPolicies in the indexer.
func (s *tunnelIamPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.TunnelIamPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TunnelIamPolicy))
	})
	return ret, err
}

// TunnelIamPolicies returns an object that can list and get TunnelIamPolicies.
func (s *tunnelIamPolicyLister) TunnelIamPolicies(namespace string) TunnelIamPolicyNamespaceLister {
	return tunnelIamPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TunnelIamPolicyNamespaceLister helps list and get TunnelIamPolicies.
// All objects returned here must be treated as read-only.
type TunnelIamPolicyNamespaceLister interface {
	// List lists all TunnelIamPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TunnelIamPolicy, err error)
	// Get retrieves the TunnelIamPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.TunnelIamPolicy, error)
	TunnelIamPolicyNamespaceListerExpansion
}

// tunnelIamPolicyNamespaceLister implements the TunnelIamPolicyNamespaceLister
// interface.
type tunnelIamPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TunnelIamPolicies in the indexer for a given namespace.
func (s tunnelIamPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TunnelIamPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TunnelIamPolicy))
	})
	return ret, err
}

// Get retrieves the TunnelIamPolicy from the indexer for a given namespace and name.
func (s tunnelIamPolicyNamespaceLister) Get(name string) (*v1alpha1.TunnelIamPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("tunneliampolicy"), name)
	}
	return obj.(*v1alpha1.TunnelIamPolicy), nil
}
