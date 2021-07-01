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

// PlatformTenantInboundSamlConfigLister helps list PlatformTenantInboundSamlConfigs.
// All objects returned here must be treated as read-only.
type PlatformTenantInboundSamlConfigLister interface {
	// List lists all PlatformTenantInboundSamlConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PlatformTenantInboundSamlConfig, err error)
	// PlatformTenantInboundSamlConfigs returns an object that can list and get PlatformTenantInboundSamlConfigs.
	PlatformTenantInboundSamlConfigs(namespace string) PlatformTenantInboundSamlConfigNamespaceLister
	PlatformTenantInboundSamlConfigListerExpansion
}

// platformTenantInboundSamlConfigLister implements the PlatformTenantInboundSamlConfigLister interface.
type platformTenantInboundSamlConfigLister struct {
	indexer cache.Indexer
}

// NewPlatformTenantInboundSamlConfigLister returns a new PlatformTenantInboundSamlConfigLister.
func NewPlatformTenantInboundSamlConfigLister(indexer cache.Indexer) PlatformTenantInboundSamlConfigLister {
	return &platformTenantInboundSamlConfigLister{indexer: indexer}
}

// List lists all PlatformTenantInboundSamlConfigs in the indexer.
func (s *platformTenantInboundSamlConfigLister) List(selector labels.Selector) (ret []*v1alpha1.PlatformTenantInboundSamlConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PlatformTenantInboundSamlConfig))
	})
	return ret, err
}

// PlatformTenantInboundSamlConfigs returns an object that can list and get PlatformTenantInboundSamlConfigs.
func (s *platformTenantInboundSamlConfigLister) PlatformTenantInboundSamlConfigs(namespace string) PlatformTenantInboundSamlConfigNamespaceLister {
	return platformTenantInboundSamlConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PlatformTenantInboundSamlConfigNamespaceLister helps list and get PlatformTenantInboundSamlConfigs.
// All objects returned here must be treated as read-only.
type PlatformTenantInboundSamlConfigNamespaceLister interface {
	// List lists all PlatformTenantInboundSamlConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PlatformTenantInboundSamlConfig, err error)
	// Get retrieves the PlatformTenantInboundSamlConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PlatformTenantInboundSamlConfig, error)
	PlatformTenantInboundSamlConfigNamespaceListerExpansion
}

// platformTenantInboundSamlConfigNamespaceLister implements the PlatformTenantInboundSamlConfigNamespaceLister
// interface.
type platformTenantInboundSamlConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PlatformTenantInboundSamlConfigs in the indexer for a given namespace.
func (s platformTenantInboundSamlConfigNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PlatformTenantInboundSamlConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PlatformTenantInboundSamlConfig))
	})
	return ret, err
}

// Get retrieves the PlatformTenantInboundSamlConfig from the indexer for a given namespace and name.
func (s platformTenantInboundSamlConfigNamespaceLister) Get(name string) (*v1alpha1.PlatformTenantInboundSamlConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("platformtenantinboundsamlconfig"), name)
	}
	return obj.(*v1alpha1.PlatformTenantInboundSamlConfig), nil
}
