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

// NodeTemplateLister helps list NodeTemplates.
// All objects returned here must be treated as read-only.
type NodeTemplateLister interface {
	// List lists all NodeTemplates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NodeTemplate, err error)
	// NodeTemplates returns an object that can list and get NodeTemplates.
	NodeTemplates(namespace string) NodeTemplateNamespaceLister
	NodeTemplateListerExpansion
}

// nodeTemplateLister implements the NodeTemplateLister interface.
type nodeTemplateLister struct {
	indexer cache.Indexer
}

// NewNodeTemplateLister returns a new NodeTemplateLister.
func NewNodeTemplateLister(indexer cache.Indexer) NodeTemplateLister {
	return &nodeTemplateLister{indexer: indexer}
}

// List lists all NodeTemplates in the indexer.
func (s *nodeTemplateLister) List(selector labels.Selector) (ret []*v1alpha1.NodeTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NodeTemplate))
	})
	return ret, err
}

// NodeTemplates returns an object that can list and get NodeTemplates.
func (s *nodeTemplateLister) NodeTemplates(namespace string) NodeTemplateNamespaceLister {
	return nodeTemplateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NodeTemplateNamespaceLister helps list and get NodeTemplates.
// All objects returned here must be treated as read-only.
type NodeTemplateNamespaceLister interface {
	// List lists all NodeTemplates in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NodeTemplate, err error)
	// Get retrieves the NodeTemplate from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.NodeTemplate, error)
	NodeTemplateNamespaceListerExpansion
}

// nodeTemplateNamespaceLister implements the NodeTemplateNamespaceLister
// interface.
type nodeTemplateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NodeTemplates in the indexer for a given namespace.
func (s nodeTemplateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NodeTemplate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NodeTemplate))
	})
	return ret, err
}

// Get retrieves the NodeTemplate from the indexer for a given namespace and name.
func (s nodeTemplateNamespaceLister) Get(name string) (*v1alpha1.NodeTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("nodetemplate"), name)
	}
	return obj.(*v1alpha1.NodeTemplate), nil
}
