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
	v1alpha1 "kubeform.dev/provider-google-api/apis/iap/v1alpha1"
)

// WebBackendServiceIamPolicyLister helps list WebBackendServiceIamPolicies.
// All objects returned here must be treated as read-only.
type WebBackendServiceIamPolicyLister interface {
	// List lists all WebBackendServiceIamPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WebBackendServiceIamPolicy, err error)
	// WebBackendServiceIamPolicies returns an object that can list and get WebBackendServiceIamPolicies.
	WebBackendServiceIamPolicies(namespace string) WebBackendServiceIamPolicyNamespaceLister
	WebBackendServiceIamPolicyListerExpansion
}

// webBackendServiceIamPolicyLister implements the WebBackendServiceIamPolicyLister interface.
type webBackendServiceIamPolicyLister struct {
	indexer cache.Indexer
}

// NewWebBackendServiceIamPolicyLister returns a new WebBackendServiceIamPolicyLister.
func NewWebBackendServiceIamPolicyLister(indexer cache.Indexer) WebBackendServiceIamPolicyLister {
	return &webBackendServiceIamPolicyLister{indexer: indexer}
}

// List lists all WebBackendServiceIamPolicies in the indexer.
func (s *webBackendServiceIamPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.WebBackendServiceIamPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WebBackendServiceIamPolicy))
	})
	return ret, err
}

// WebBackendServiceIamPolicies returns an object that can list and get WebBackendServiceIamPolicies.
func (s *webBackendServiceIamPolicyLister) WebBackendServiceIamPolicies(namespace string) WebBackendServiceIamPolicyNamespaceLister {
	return webBackendServiceIamPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WebBackendServiceIamPolicyNamespaceLister helps list and get WebBackendServiceIamPolicies.
// All objects returned here must be treated as read-only.
type WebBackendServiceIamPolicyNamespaceLister interface {
	// List lists all WebBackendServiceIamPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WebBackendServiceIamPolicy, err error)
	// Get retrieves the WebBackendServiceIamPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.WebBackendServiceIamPolicy, error)
	WebBackendServiceIamPolicyNamespaceListerExpansion
}

// webBackendServiceIamPolicyNamespaceLister implements the WebBackendServiceIamPolicyNamespaceLister
// interface.
type webBackendServiceIamPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WebBackendServiceIamPolicies in the indexer for a given namespace.
func (s webBackendServiceIamPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WebBackendServiceIamPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WebBackendServiceIamPolicy))
	})
	return ret, err
}

// Get retrieves the WebBackendServiceIamPolicy from the indexer for a given namespace and name.
func (s webBackendServiceIamPolicyNamespaceLister) Get(name string) (*v1alpha1.WebBackendServiceIamPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("webbackendserviceiampolicy"), name)
	}
	return obj.(*v1alpha1.WebBackendServiceIamPolicy), nil
}
