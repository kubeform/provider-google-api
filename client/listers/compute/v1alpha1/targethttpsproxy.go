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

// TargetHTTPSProxyLister helps list TargetHTTPSProxies.
// All objects returned here must be treated as read-only.
type TargetHTTPSProxyLister interface {
	// List lists all TargetHTTPSProxies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TargetHTTPSProxy, err error)
	// TargetHTTPSProxies returns an object that can list and get TargetHTTPSProxies.
	TargetHTTPSProxies(namespace string) TargetHTTPSProxyNamespaceLister
	TargetHTTPSProxyListerExpansion
}

// targetHTTPSProxyLister implements the TargetHTTPSProxyLister interface.
type targetHTTPSProxyLister struct {
	indexer cache.Indexer
}

// NewTargetHTTPSProxyLister returns a new TargetHTTPSProxyLister.
func NewTargetHTTPSProxyLister(indexer cache.Indexer) TargetHTTPSProxyLister {
	return &targetHTTPSProxyLister{indexer: indexer}
}

// List lists all TargetHTTPSProxies in the indexer.
func (s *targetHTTPSProxyLister) List(selector labels.Selector) (ret []*v1alpha1.TargetHTTPSProxy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TargetHTTPSProxy))
	})
	return ret, err
}

// TargetHTTPSProxies returns an object that can list and get TargetHTTPSProxies.
func (s *targetHTTPSProxyLister) TargetHTTPSProxies(namespace string) TargetHTTPSProxyNamespaceLister {
	return targetHTTPSProxyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TargetHTTPSProxyNamespaceLister helps list and get TargetHTTPSProxies.
// All objects returned here must be treated as read-only.
type TargetHTTPSProxyNamespaceLister interface {
	// List lists all TargetHTTPSProxies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TargetHTTPSProxy, err error)
	// Get retrieves the TargetHTTPSProxy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.TargetHTTPSProxy, error)
	TargetHTTPSProxyNamespaceListerExpansion
}

// targetHTTPSProxyNamespaceLister implements the TargetHTTPSProxyNamespaceLister
// interface.
type targetHTTPSProxyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TargetHTTPSProxies in the indexer for a given namespace.
func (s targetHTTPSProxyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TargetHTTPSProxy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TargetHTTPSProxy))
	})
	return ret, err
}

// Get retrieves the TargetHTTPSProxy from the indexer for a given namespace and name.
func (s targetHTTPSProxyNamespaceLister) Get(name string) (*v1alpha1.TargetHTTPSProxy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("targethttpsproxy"), name)
	}
	return obj.(*v1alpha1.TargetHTTPSProxy), nil
}
