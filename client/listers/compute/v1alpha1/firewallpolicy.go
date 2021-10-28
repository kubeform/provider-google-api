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
	v1alpha1 "kubeform.dev/provider-google-api/apis/compute/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FirewallPolicyLister helps list FirewallPolicies.
// All objects returned here must be treated as read-only.
type FirewallPolicyLister interface {
	// List lists all FirewallPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FirewallPolicy, err error)
	// FirewallPolicies returns an object that can list and get FirewallPolicies.
	FirewallPolicies(namespace string) FirewallPolicyNamespaceLister
	FirewallPolicyListerExpansion
}

// firewallPolicyLister implements the FirewallPolicyLister interface.
type firewallPolicyLister struct {
	indexer cache.Indexer
}

// NewFirewallPolicyLister returns a new FirewallPolicyLister.
func NewFirewallPolicyLister(indexer cache.Indexer) FirewallPolicyLister {
	return &firewallPolicyLister{indexer: indexer}
}

// List lists all FirewallPolicies in the indexer.
func (s *firewallPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.FirewallPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FirewallPolicy))
	})
	return ret, err
}

// FirewallPolicies returns an object that can list and get FirewallPolicies.
func (s *firewallPolicyLister) FirewallPolicies(namespace string) FirewallPolicyNamespaceLister {
	return firewallPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FirewallPolicyNamespaceLister helps list and get FirewallPolicies.
// All objects returned here must be treated as read-only.
type FirewallPolicyNamespaceLister interface {
	// List lists all FirewallPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FirewallPolicy, err error)
	// Get retrieves the FirewallPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FirewallPolicy, error)
	FirewallPolicyNamespaceListerExpansion
}

// firewallPolicyNamespaceLister implements the FirewallPolicyNamespaceLister
// interface.
type firewallPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FirewallPolicies in the indexer for a given namespace.
func (s firewallPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FirewallPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FirewallPolicy))
	})
	return ret, err
}

// Get retrieves the FirewallPolicy from the indexer for a given namespace and name.
func (s firewallPolicyNamespaceLister) Get(name string) (*v1alpha1.FirewallPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("firewallpolicy"), name)
	}
	return obj.(*v1alpha1.FirewallPolicy), nil
}
