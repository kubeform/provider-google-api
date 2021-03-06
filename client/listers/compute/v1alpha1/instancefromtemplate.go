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

// InstanceFromTemplateLister helps list InstanceFromTemplates.
// All objects returned here must be treated as read-only.
type InstanceFromTemplateLister interface {
	// List lists all InstanceFromTemplates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InstanceFromTemplate, err error)
	// InstanceFromTemplates returns an object that can list and get InstanceFromTemplates.
	InstanceFromTemplates(namespace string) InstanceFromTemplateNamespaceLister
	InstanceFromTemplateListerExpansion
}

// instanceFromTemplateLister implements the InstanceFromTemplateLister interface.
type instanceFromTemplateLister struct {
	indexer cache.Indexer
}

// NewInstanceFromTemplateLister returns a new InstanceFromTemplateLister.
func NewInstanceFromTemplateLister(indexer cache.Indexer) InstanceFromTemplateLister {
	return &instanceFromTemplateLister{indexer: indexer}
}

// List lists all InstanceFromTemplates in the indexer.
func (s *instanceFromTemplateLister) List(selector labels.Selector) (ret []*v1alpha1.InstanceFromTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InstanceFromTemplate))
	})
	return ret, err
}

// InstanceFromTemplates returns an object that can list and get InstanceFromTemplates.
func (s *instanceFromTemplateLister) InstanceFromTemplates(namespace string) InstanceFromTemplateNamespaceLister {
	return instanceFromTemplateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InstanceFromTemplateNamespaceLister helps list and get InstanceFromTemplates.
// All objects returned here must be treated as read-only.
type InstanceFromTemplateNamespaceLister interface {
	// List lists all InstanceFromTemplates in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InstanceFromTemplate, err error)
	// Get retrieves the InstanceFromTemplate from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.InstanceFromTemplate, error)
	InstanceFromTemplateNamespaceListerExpansion
}

// instanceFromTemplateNamespaceLister implements the InstanceFromTemplateNamespaceLister
// interface.
type instanceFromTemplateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all InstanceFromTemplates in the indexer for a given namespace.
func (s instanceFromTemplateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.InstanceFromTemplate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InstanceFromTemplate))
	})
	return ret, err
}

// Get retrieves the InstanceFromTemplate from the indexer for a given namespace and name.
func (s instanceFromTemplateNamespaceLister) Get(name string) (*v1alpha1.InstanceFromTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("instancefromtemplate"), name)
	}
	return obj.(*v1alpha1.InstanceFromTemplate), nil
}
