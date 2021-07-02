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
	v1alpha1 "kubeform.dev/provider-google-api/apis/dialogflow/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CxAgentLister helps list CxAgents.
// All objects returned here must be treated as read-only.
type CxAgentLister interface {
	// List lists all CxAgents in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CxAgent, err error)
	// CxAgents returns an object that can list and get CxAgents.
	CxAgents(namespace string) CxAgentNamespaceLister
	CxAgentListerExpansion
}

// cxAgentLister implements the CxAgentLister interface.
type cxAgentLister struct {
	indexer cache.Indexer
}

// NewCxAgentLister returns a new CxAgentLister.
func NewCxAgentLister(indexer cache.Indexer) CxAgentLister {
	return &cxAgentLister{indexer: indexer}
}

// List lists all CxAgents in the indexer.
func (s *cxAgentLister) List(selector labels.Selector) (ret []*v1alpha1.CxAgent, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CxAgent))
	})
	return ret, err
}

// CxAgents returns an object that can list and get CxAgents.
func (s *cxAgentLister) CxAgents(namespace string) CxAgentNamespaceLister {
	return cxAgentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CxAgentNamespaceLister helps list and get CxAgents.
// All objects returned here must be treated as read-only.
type CxAgentNamespaceLister interface {
	// List lists all CxAgents in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CxAgent, err error)
	// Get retrieves the CxAgent from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CxAgent, error)
	CxAgentNamespaceListerExpansion
}

// cxAgentNamespaceLister implements the CxAgentNamespaceLister
// interface.
type cxAgentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CxAgents in the indexer for a given namespace.
func (s cxAgentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CxAgent, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CxAgent))
	})
	return ret, err
}

// Get retrieves the CxAgent from the indexer for a given namespace and name.
func (s cxAgentNamespaceLister) Get(name string) (*v1alpha1.CxAgent, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cxagent"), name)
	}
	return obj.(*v1alpha1.CxAgent), nil
}
