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

// WebTypeAppEngineIamPolicyLister helps list WebTypeAppEngineIamPolicies.
// All objects returned here must be treated as read-only.
type WebTypeAppEngineIamPolicyLister interface {
	// List lists all WebTypeAppEngineIamPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WebTypeAppEngineIamPolicy, err error)
	// WebTypeAppEngineIamPolicies returns an object that can list and get WebTypeAppEngineIamPolicies.
	WebTypeAppEngineIamPolicies(namespace string) WebTypeAppEngineIamPolicyNamespaceLister
	WebTypeAppEngineIamPolicyListerExpansion
}

// webTypeAppEngineIamPolicyLister implements the WebTypeAppEngineIamPolicyLister interface.
type webTypeAppEngineIamPolicyLister struct {
	indexer cache.Indexer
}

// NewWebTypeAppEngineIamPolicyLister returns a new WebTypeAppEngineIamPolicyLister.
func NewWebTypeAppEngineIamPolicyLister(indexer cache.Indexer) WebTypeAppEngineIamPolicyLister {
	return &webTypeAppEngineIamPolicyLister{indexer: indexer}
}

// List lists all WebTypeAppEngineIamPolicies in the indexer.
func (s *webTypeAppEngineIamPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.WebTypeAppEngineIamPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WebTypeAppEngineIamPolicy))
	})
	return ret, err
}

// WebTypeAppEngineIamPolicies returns an object that can list and get WebTypeAppEngineIamPolicies.
func (s *webTypeAppEngineIamPolicyLister) WebTypeAppEngineIamPolicies(namespace string) WebTypeAppEngineIamPolicyNamespaceLister {
	return webTypeAppEngineIamPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WebTypeAppEngineIamPolicyNamespaceLister helps list and get WebTypeAppEngineIamPolicies.
// All objects returned here must be treated as read-only.
type WebTypeAppEngineIamPolicyNamespaceLister interface {
	// List lists all WebTypeAppEngineIamPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WebTypeAppEngineIamPolicy, err error)
	// Get retrieves the WebTypeAppEngineIamPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.WebTypeAppEngineIamPolicy, error)
	WebTypeAppEngineIamPolicyNamespaceListerExpansion
}

// webTypeAppEngineIamPolicyNamespaceLister implements the WebTypeAppEngineIamPolicyNamespaceLister
// interface.
type webTypeAppEngineIamPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WebTypeAppEngineIamPolicies in the indexer for a given namespace.
func (s webTypeAppEngineIamPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WebTypeAppEngineIamPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WebTypeAppEngineIamPolicy))
	})
	return ret, err
}

// Get retrieves the WebTypeAppEngineIamPolicy from the indexer for a given namespace and name.
func (s webTypeAppEngineIamPolicyNamespaceLister) Get(name string) (*v1alpha1.WebTypeAppEngineIamPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("webtypeappengineiampolicy"), name)
	}
	return obj.(*v1alpha1.WebTypeAppEngineIamPolicy), nil
}
