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

// WebIamPolicyLister helps list WebIamPolicies.
// All objects returned here must be treated as read-only.
type WebIamPolicyLister interface {
	// List lists all WebIamPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WebIamPolicy, err error)
	// WebIamPolicies returns an object that can list and get WebIamPolicies.
	WebIamPolicies(namespace string) WebIamPolicyNamespaceLister
	WebIamPolicyListerExpansion
}

// webIamPolicyLister implements the WebIamPolicyLister interface.
type webIamPolicyLister struct {
	indexer cache.Indexer
}

// NewWebIamPolicyLister returns a new WebIamPolicyLister.
func NewWebIamPolicyLister(indexer cache.Indexer) WebIamPolicyLister {
	return &webIamPolicyLister{indexer: indexer}
}

// List lists all WebIamPolicies in the indexer.
func (s *webIamPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.WebIamPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WebIamPolicy))
	})
	return ret, err
}

// WebIamPolicies returns an object that can list and get WebIamPolicies.
func (s *webIamPolicyLister) WebIamPolicies(namespace string) WebIamPolicyNamespaceLister {
	return webIamPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WebIamPolicyNamespaceLister helps list and get WebIamPolicies.
// All objects returned here must be treated as read-only.
type WebIamPolicyNamespaceLister interface {
	// List lists all WebIamPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WebIamPolicy, err error)
	// Get retrieves the WebIamPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.WebIamPolicy, error)
	WebIamPolicyNamespaceListerExpansion
}

// webIamPolicyNamespaceLister implements the WebIamPolicyNamespaceLister
// interface.
type webIamPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WebIamPolicies in the indexer for a given namespace.
func (s webIamPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WebIamPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WebIamPolicy))
	})
	return ret, err
}

// Get retrieves the WebIamPolicy from the indexer for a given namespace and name.
func (s webIamPolicyNamespaceLister) Get(name string) (*v1alpha1.WebIamPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("webiampolicy"), name)
	}
	return obj.(*v1alpha1.WebIamPolicy), nil
}
