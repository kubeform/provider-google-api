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
	v1alpha1 "kubeform.dev/provider-google-api/apis/dataproc/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterIamPolicyLister helps list ClusterIamPolicies.
// All objects returned here must be treated as read-only.
type ClusterIamPolicyLister interface {
	// List lists all ClusterIamPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterIamPolicy, err error)
	// ClusterIamPolicies returns an object that can list and get ClusterIamPolicies.
	ClusterIamPolicies(namespace string) ClusterIamPolicyNamespaceLister
	ClusterIamPolicyListerExpansion
}

// clusterIamPolicyLister implements the ClusterIamPolicyLister interface.
type clusterIamPolicyLister struct {
	indexer cache.Indexer
}

// NewClusterIamPolicyLister returns a new ClusterIamPolicyLister.
func NewClusterIamPolicyLister(indexer cache.Indexer) ClusterIamPolicyLister {
	return &clusterIamPolicyLister{indexer: indexer}
}

// List lists all ClusterIamPolicies in the indexer.
func (s *clusterIamPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterIamPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterIamPolicy))
	})
	return ret, err
}

// ClusterIamPolicies returns an object that can list and get ClusterIamPolicies.
func (s *clusterIamPolicyLister) ClusterIamPolicies(namespace string) ClusterIamPolicyNamespaceLister {
	return clusterIamPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterIamPolicyNamespaceLister helps list and get ClusterIamPolicies.
// All objects returned here must be treated as read-only.
type ClusterIamPolicyNamespaceLister interface {
	// List lists all ClusterIamPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterIamPolicy, err error)
	// Get retrieves the ClusterIamPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterIamPolicy, error)
	ClusterIamPolicyNamespaceListerExpansion
}

// clusterIamPolicyNamespaceLister implements the ClusterIamPolicyNamespaceLister
// interface.
type clusterIamPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterIamPolicies in the indexer for a given namespace.
func (s clusterIamPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterIamPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterIamPolicy))
	})
	return ret, err
}

// Get retrieves the ClusterIamPolicy from the indexer for a given namespace and name.
func (s clusterIamPolicyNamespaceLister) Get(name string) (*v1alpha1.ClusterIamPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusteriampolicy"), name)
	}
	return obj.(*v1alpha1.ClusterIamPolicy), nil
}
