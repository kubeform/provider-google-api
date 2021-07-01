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
	v1alpha1 "kubeform.dev/provider-google-api/apis/identity/v1alpha1"
)

// PlatformInboundSamlConfigLister helps list PlatformInboundSamlConfigs.
// All objects returned here must be treated as read-only.
type PlatformInboundSamlConfigLister interface {
	// List lists all PlatformInboundSamlConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PlatformInboundSamlConfig, err error)
	// PlatformInboundSamlConfigs returns an object that can list and get PlatformInboundSamlConfigs.
	PlatformInboundSamlConfigs(namespace string) PlatformInboundSamlConfigNamespaceLister
	PlatformInboundSamlConfigListerExpansion
}

// platformInboundSamlConfigLister implements the PlatformInboundSamlConfigLister interface.
type platformInboundSamlConfigLister struct {
	indexer cache.Indexer
}

// NewPlatformInboundSamlConfigLister returns a new PlatformInboundSamlConfigLister.
func NewPlatformInboundSamlConfigLister(indexer cache.Indexer) PlatformInboundSamlConfigLister {
	return &platformInboundSamlConfigLister{indexer: indexer}
}

// List lists all PlatformInboundSamlConfigs in the indexer.
func (s *platformInboundSamlConfigLister) List(selector labels.Selector) (ret []*v1alpha1.PlatformInboundSamlConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PlatformInboundSamlConfig))
	})
	return ret, err
}

// PlatformInboundSamlConfigs returns an object that can list and get PlatformInboundSamlConfigs.
func (s *platformInboundSamlConfigLister) PlatformInboundSamlConfigs(namespace string) PlatformInboundSamlConfigNamespaceLister {
	return platformInboundSamlConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PlatformInboundSamlConfigNamespaceLister helps list and get PlatformInboundSamlConfigs.
// All objects returned here must be treated as read-only.
type PlatformInboundSamlConfigNamespaceLister interface {
	// List lists all PlatformInboundSamlConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PlatformInboundSamlConfig, err error)
	// Get retrieves the PlatformInboundSamlConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PlatformInboundSamlConfig, error)
	PlatformInboundSamlConfigNamespaceListerExpansion
}

// platformInboundSamlConfigNamespaceLister implements the PlatformInboundSamlConfigNamespaceLister
// interface.
type platformInboundSamlConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PlatformInboundSamlConfigs in the indexer for a given namespace.
func (s platformInboundSamlConfigNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PlatformInboundSamlConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PlatformInboundSamlConfig))
	})
	return ret, err
}

// Get retrieves the PlatformInboundSamlConfig from the indexer for a given namespace and name.
func (s platformInboundSamlConfigNamespaceLister) Get(name string) (*v1alpha1.PlatformInboundSamlConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("platforminboundsamlconfig"), name)
	}
	return obj.(*v1alpha1.PlatformInboundSamlConfig), nil
}
