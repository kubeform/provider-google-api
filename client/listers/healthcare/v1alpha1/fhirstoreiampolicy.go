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
	v1alpha1 "kubeform.dev/provider-google-api/apis/healthcare/v1alpha1"
)

// FhirStoreIamPolicyLister helps list FhirStoreIamPolicies.
// All objects returned here must be treated as read-only.
type FhirStoreIamPolicyLister interface {
	// List lists all FhirStoreIamPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FhirStoreIamPolicy, err error)
	// FhirStoreIamPolicies returns an object that can list and get FhirStoreIamPolicies.
	FhirStoreIamPolicies(namespace string) FhirStoreIamPolicyNamespaceLister
	FhirStoreIamPolicyListerExpansion
}

// fhirStoreIamPolicyLister implements the FhirStoreIamPolicyLister interface.
type fhirStoreIamPolicyLister struct {
	indexer cache.Indexer
}

// NewFhirStoreIamPolicyLister returns a new FhirStoreIamPolicyLister.
func NewFhirStoreIamPolicyLister(indexer cache.Indexer) FhirStoreIamPolicyLister {
	return &fhirStoreIamPolicyLister{indexer: indexer}
}

// List lists all FhirStoreIamPolicies in the indexer.
func (s *fhirStoreIamPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.FhirStoreIamPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FhirStoreIamPolicy))
	})
	return ret, err
}

// FhirStoreIamPolicies returns an object that can list and get FhirStoreIamPolicies.
func (s *fhirStoreIamPolicyLister) FhirStoreIamPolicies(namespace string) FhirStoreIamPolicyNamespaceLister {
	return fhirStoreIamPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FhirStoreIamPolicyNamespaceLister helps list and get FhirStoreIamPolicies.
// All objects returned here must be treated as read-only.
type FhirStoreIamPolicyNamespaceLister interface {
	// List lists all FhirStoreIamPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FhirStoreIamPolicy, err error)
	// Get retrieves the FhirStoreIamPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FhirStoreIamPolicy, error)
	FhirStoreIamPolicyNamespaceListerExpansion
}

// fhirStoreIamPolicyNamespaceLister implements the FhirStoreIamPolicyNamespaceLister
// interface.
type fhirStoreIamPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FhirStoreIamPolicies in the indexer for a given namespace.
func (s fhirStoreIamPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FhirStoreIamPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FhirStoreIamPolicy))
	})
	return ret, err
}

// Get retrieves the FhirStoreIamPolicy from the indexer for a given namespace and name.
func (s fhirStoreIamPolicyNamespaceLister) Get(name string) (*v1alpha1.FhirStoreIamPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("fhirstoreiampolicy"), name)
	}
	return obj.(*v1alpha1.FhirStoreIamPolicy), nil
}
